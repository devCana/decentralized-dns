// Package torrent wraps anacrolix/torrent behind the BitTorrentEngine API
// from HLD §3.6: resolvers seed published resources and fetch resources by
// infohash, re-hashing every payload end-to-end (SHA-256) against the
// on-chain digest before anything is served (UC-10 tamper rejection).
package torrent

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/bencode"
	"github.com/anacrolix/torrent/metainfo"
)

// MaxFetchBytes caps fetched resources; published site bundles are small
// zips, and the resolver must not be lured into buffering arbitrary blobs.
const MaxFetchBytes = 256 << 20 // 256 MiB

// ErrHashMismatch is returned when swarm-delivered content does not hash to
// the expected on-chain SHA-256. Content is discarded, never served (UC-10).
var ErrHashMismatch = errors.New("torrent: content hash does not match expected SHA-256")

// ErrTooLarge is returned when the announced torrent exceeds MaxFetchBytes.
var ErrTooLarge = errors.New("torrent: resource exceeds maximum fetch size")

// Config tunes an Engine.
type Config struct {
	DataDir    string       // where seeded/fetched payloads live
	ListenPort int          // TCP/uTP listen port (0 = random)
	DisableDHT bool         // true for local/e2e setups using explicit peers
	Logger     *slog.Logger // optional; defaults to slog.Default()
}

// Engine seeds and fetches static resources over BitTorrent.
type Engine struct {
	client *torrent.Client
	log    *slog.Logger
}

// New starts a BitTorrent client. Close the engine to stop seeding.
func New(cfg Config) (*Engine, error) {
	if cfg.Logger == nil {
		cfg.Logger = slog.Default()
	}
	if err := os.MkdirAll(cfg.DataDir, 0o755); err != nil {
		return nil, fmt.Errorf("torrent: data dir: %w", err)
	}
	cc := torrent.NewDefaultClientConfig()
	cc.DataDir = cfg.DataDir
	cc.ListenPort = cfg.ListenPort
	cc.NoDHT = cfg.DisableDHT
	cc.Seed = true
	cc.Logger.SetHandlers() // silence anacrolix's own logging; we use slog
	client, err := torrent.NewClient(cc)
	if err != nil {
		return nil, fmt.Errorf("torrent: client: %w", err)
	}
	return &Engine{client: client, log: cfg.Logger}, nil
}

// Close stops seeding and releases the listen sockets.
func (e *Engine) Close() {
	e.client.Close()
	<-e.client.Closed()
}

// ListenAddrs returns dialable addresses of this engine for explicit peer
// wiring (local compose networks have no public DHT). Unspecified hosts
// are rewritten to 127.0.0.1.
func (e *Engine) ListenAddrs() []string {
	var out []string
	seen := map[string]bool{}
	for _, a := range e.client.ListenAddrs() {
		host, port, err := net.SplitHostPort(a.String())
		if err != nil {
			continue
		}
		if ip := net.ParseIP(host); ip == nil || ip.IsUnspecified() {
			host = "127.0.0.1"
		}
		hp := net.JoinHostPort(host, port)
		if !seen[hp] {
			seen[hp] = true
			out = append(out, hp)
		}
	}
	return out
}

// SeedFile makes path available to the swarm. It returns the torrent
// infohash and the file's SHA-256, both hex — exactly what a ResourceRef
// record anchors on-chain. The file must stay in place while seeding.
func (e *Engine) SeedFile(path string) (infoHash, sha string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		f.Close()
		return "", "", err
	}
	f.Close()

	info := metainfo.Info{PieceLength: 256 << 10}
	if err := info.BuildFromFilePath(path); err != nil {
		return "", "", fmt.Errorf("torrent: build metainfo: %w", err)
	}
	infoBytes, err := bencode.Marshal(info)
	if err != nil {
		return "", "", err
	}
	mi := &metainfo.MetaInfo{InfoBytes: infoBytes}
	t, err := e.client.AddTorrent(mi)
	if err != nil {
		return "", "", fmt.Errorf("torrent: add: %w", err)
	}
	<-t.GotInfo()
	ih := t.InfoHash().HexString()
	digest := hex.EncodeToString(h.Sum(nil))
	e.log.Info("seeding resource", "infoHash", ih, "sha256", digest, "bytes", t.Length())
	return ih, digest, nil
}

// Fetch downloads the torrent identified by infoHashHex, re-hashes the
// payload and compares it to expectedSHAHex before returning it. On any
// mismatch the data is dropped and ErrHashMismatch returned — tampered
// content is never served (UC-10). peers lists explicit host:port seeds
// for networks without DHT; pass nil to rely on DHT discovery.
func (e *Engine) Fetch(ctx context.Context, infoHashHex, expectedSHAHex string, peers []string) ([]byte, error) {
	var ih metainfo.Hash
	if err := ih.FromHexString(infoHashHex); err != nil {
		return nil, fmt.Errorf("torrent: bad infohash %q: %w", infoHashHex, err)
	}
	expected, err := hex.DecodeString(expectedSHAHex)
	if err != nil || len(expected) != sha256.Size {
		return nil, fmt.Errorf("torrent: bad expected sha256 %q", expectedSHAHex)
	}

	t, isNew := e.client.AddTorrentInfoHashWithStorage(ih, nil)
	if isNew {
		defer t.Drop()
	}
	for _, p := range peers {
		host, portStr, err := net.SplitHostPort(p)
		if err != nil {
			return nil, fmt.Errorf("torrent: bad peer %q: %w", p, err)
		}
		addr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(host, portStr))
		if err != nil {
			return nil, fmt.Errorf("torrent: bad peer %q: %w", p, err)
		}
		t.AddPeers([]torrent.PeerInfo{{Addr: addr}})
	}

	select {
	case <-t.GotInfo():
	case <-ctx.Done():
		return nil, fmt.Errorf("torrent: metadata for %s: %w", infoHashHex, ctx.Err())
	}
	if t.Length() > MaxFetchBytes {
		return nil, fmt.Errorf("%w: %d bytes", ErrTooLarge, t.Length())
	}

	r := t.NewReader()
	defer r.Close()
	buf := bytes.NewBuffer(make([]byte, 0, t.Length()))
	done := make(chan error, 1)
	go func() {
		_, err := io.Copy(buf, r)
		done <- err
	}()
	select {
	case err := <-done:
		if err != nil {
			return nil, fmt.Errorf("torrent: download %s: %w", infoHashHex, err)
		}
	case <-ctx.Done():
		return nil, fmt.Errorf("torrent: download %s: %w", infoHashHex, ctx.Err())
	}

	sum := sha256.Sum256(buf.Bytes())
	if !bytes.Equal(sum[:], expected) {
		e.log.Warn("tampered resource rejected",
			"infoHash", infoHashHex, "expected", expectedSHAHex, "got", hex.EncodeToString(sum[:]))
		return nil, ErrHashMismatch
	}
	e.log.Info("resource fetched and verified", "infoHash", infoHashHex, "bytes", buf.Len())
	return buf.Bytes(), nil
}

// Stats summarizes swarm state for the admin dashboard.
type Stats struct {
	Torrents    int   `json:"torrents"`
	TotalPeers  int   `json:"totalPeers"`
	BytesShared int64 `json:"bytesShared"`
}

// Stats reports the number of active torrents and connected peers.
func (e *Engine) Stats() Stats {
	s := Stats{}
	for _, t := range e.client.Torrents() {
		s.Torrents++
		st := t.Stats()
		s.TotalPeers += st.ActivePeers
		s.BytesShared += st.BytesWrittenData.Int64()
	}
	return s
}
