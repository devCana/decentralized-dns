package torrent

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"log/slog"
	"os"
	"path/filepath"
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
