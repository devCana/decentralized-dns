package torrent

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

func testEngine(t *testing.T) *Engine {
	t.Helper()
	e, err := New(Config{
		DataDir:    t.TempDir(),
		ListenPort: 0,
		DisableDHT: true,
		Logger:     slog.New(slog.DiscardHandler),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(e.Close)
	return e
}

func TestSeedAndFetchVerifiesSHA256(t *testing.T) {
	seed := testEngine(t)
	fetcher := testEngine(t)

	payload := []byte("<!doctype html><title>ddns</title><p>verified torrent payload</p>")
	file := filepath.Join(t.TempDir(), "index.html")
	if err := os.WriteFile(file, payload, 0o644); err != nil {
		t.Fatal(err)
	}

	infoHash, digest, err := seed.SeedFile(context.Background(), file)
	if err != nil {
		t.Fatal(err)
	}
	want := sha256.Sum256(payload)
	if digest != hex.EncodeToString(want[:]) {
		t.Fatalf("digest = %s, want %x", digest, want)
	}
	if infoHash == "" {
		t.Fatal("empty infohash")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	got, err := fetcher.Fetch(ctx, infoHash, digest, seed.ListenAddrs())
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(payload) {
		t.Fatalf("payload = %q, want %q", got, payload)
	}
}

func TestFetchRetainsForLocalReuse(t *testing.T) {
	seed := testEngine(t)
	fetcher := testEngine(t)

	payload := []byte("<!doctype html><title>retained</title><h1>served twice</h1>")
	file := filepath.Join(t.TempDir(), "site.html")
	if err := os.WriteFile(file, payload, 0o644); err != nil {
		t.Fatal(err)
	}
	infoHash, digest, err := seed.SeedFile(context.Background(), file)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// First fetch uses the seeder's address as an explicit peer.
	if _, err := fetcher.Fetch(ctx, infoHash, digest, seed.ListenAddrs()); err != nil {
		t.Fatalf("first fetch: %v", err)
	}
	if got := fetcher.Stats().Torrents; got != 1 {
		t.Fatalf("retained torrents = %d, want 1", got)
	}

	// Stop the seeder, then fetch again with NO peers. It must still succeed,
	// served from the resolver's retained copy — exactly what lets /web serve a
	// site over DHT-less localhost after the first download.
	seed.Close()
	got, err := fetcher.Fetch(ctx, infoHash, digest, nil)
	if err != nil {
		t.Fatalf("retained re-fetch: %v", err)
	}
	if string(got) != string(payload) {
		t.Fatalf("payload = %q, want %q", got, payload)
	}
}

func TestConcurrentFetchSameInfohash(t *testing.T) {
	seed := testEngine(t)
	fetcher := testEngine(t)

	payload := []byte(strings.Repeat("decentralized-dns ", 2000))
	file := filepath.Join(t.TempDir(), "shared.bin")
	if err := os.WriteFile(file, payload, 0o644); err != nil {
		t.Fatal(err)
	}
	infoHash, digest, err := seed.SeedFile(context.Background(), file)
	if err != nil {
		t.Fatal(err)
	}

	// Several goroutines fetch the same infohash at once. The refcount must keep
	// the shared torrent alive until the last reader finishes — no fetch should
	// have it dropped out from under it.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	const N = 5
	var wg sync.WaitGroup
	errs := make([]error, N)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			got, err := fetcher.Fetch(ctx, infoHash, digest, seed.ListenAddrs())
			if err == nil && !bytes.Equal(got, payload) {
				err = errors.New("payload mismatch")
			}
			errs[i] = err
		}(i)
	}
	wg.Wait()
	for i, err := range errs {
		if err != nil {
			t.Errorf("concurrent fetch %d failed: %v", i, err)
		}
	}
}

func TestFetchRejectsTamperedResource(t *testing.T) {
	seed := testEngine(t)
	fetcher := testEngine(t)

	payload := []byte("original published site bundle")
	file := filepath.Join(t.TempDir(), "site.zip")
	if err := os.WriteFile(file, payload, 0o644); err != nil {
		t.Fatal(err)
	}
	infoHash, _, err := seed.SeedFile(context.Background(), file)
	if err != nil {
		t.Fatal(err)
	}
	wrong := sha256.Sum256([]byte("different bytes anchored on-chain"))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	got, err := fetcher.Fetch(ctx, infoHash, hex.EncodeToString(wrong[:]), seed.ListenAddrs())
	if !errors.Is(err, ErrHashMismatch) {
		t.Fatalf("err = %v, want ErrHashMismatch", err)
	}
	if got != nil {
		t.Fatalf("tampered fetch returned %d bytes", len(got))
	}
}

func TestFetchRejectsBadInputs(t *testing.T) {
	fetcher := testEngine(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := fetcher.Fetch(ctx, "nothex", hex.EncodeToString(make([]byte, sha256.Size)), nil); err == nil {
		t.Fatal("bad infohash accepted")
	}
	if _, err := fetcher.Fetch(ctx, "0123456789012345678901234567890123456789", "bad", nil); err == nil {
		t.Fatal("bad sha accepted")
	}
}
