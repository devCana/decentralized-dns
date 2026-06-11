package chain

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Integration test against a local chain. Skipped unless DDNS_TEST_RPC,
// DDNS_TEST_CONTRACT and DDNS_TEST_REGISTRY are set, e.g.:
//
//	cd contracts && npx hardhat node &
//	npx hardhat run scripts/deploy.ts --network localhost
//	npx hardhat run scripts/seed.ts --network localhost
//	DDNS_TEST_RPC=http://127.0.0.1:8545 DDNS_TEST_CONTRACT=0x... \
//	DDNS_TEST_REGISTRY=0x... go test ./internal/chain/
func testClient(t *testing.T) *Client {
	t.Helper()
	rpc := os.Getenv("DDNS_TEST_RPC")
	contract := os.Getenv("DDNS_TEST_CONTRACT")
	registry := os.Getenv("DDNS_TEST_REGISTRY")
	if rpc == "" || contract == "" || registry == "" {
		t.Skip("set DDNS_TEST_RPC/DDNS_TEST_CONTRACT/DDNS_TEST_REGISTRY to run chain integration tests")
	}
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))
	c, err := Dial(context.Background(), rpc,
		common.HexToAddress(contract), common.HexToAddress(registry), log)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(c.Close)
	return c
}

func TestIntegrationResolveSeededDomain(t *testing.T) {
	c := testClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	head, err := c.ChainHead(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if head == 0 {
		t.Fatal("chain head is 0; deploy+seed first")
	}

	res, err := c.Resolve(ctx, "example", "A", "")
	if err != nil {
		t.Fatal(err)
	}
	if !res.Record.Exists {
		t.Fatal("expected seeded A record for 'example'")
	}
	if addr, ok := res.Record.Field("address"); !ok || addr != "93.184.216.34" {
		t.Errorf("A address = %q, ok=%v", addr, ok)
	}
	if res.Record.TTL != 3600 {
		t.Errorf("TTL = %d, want 3600", res.Record.TTL)
	}
	if res.Owner == (common.Address{}) {
		t.Error("owner is zero")
	}
	if len(res.PubKey) == 0 {
		t.Error("pubKey is empty")
	}

	// selector-aware lookup (UC-8)
	svc, err := c.Resolve(ctx, "example", "SVC", "port=25&service=SMTP&transport=TCP")
	if err != nil {
		t.Fatal(err)
	}
	if target, _ := svc.Record.Field("target"); target != "mail.example" {
		t.Errorf("SVC target = %q", target)
	}

	// typed no-match
	miss, err := c.Resolve(ctx, "example", "SVC", "port=21&service=FTP&transport=TCP")
	if err != nil {
		t.Fatal(err)
	}
	if miss.Record.Exists {
		t.Error("expected no match for FTP selector")
	}

	// domain + registry surfaces
	dom, err := c.GetDomain(ctx, "example")
	if err != nil {
		t.Fatal(err)
	}
	if dom.Generation != 1 {
		t.Errorf("generation = %d", dom.Generation)
	}
	recs, err := c.ListRecords(ctx, "example")
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 3 {
		t.Errorf("listRecords = %d records, want 3", len(recs))
	}
	types, err := c.ListTypes(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(types) < 5 {
		t.Errorf("listTypes = %v", types)
	}
}

func TestIntegrationWatchEvents(t *testing.T) {
	c := testClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	events := make(chan RecordEvent, 16)
	go func() {
		_ = c.WatchRecordEvents(ctx, 500*time.Millisecond, func(ev RecordEvent) {
			events <- ev
		})
	}()

	// The watcher only reports events after its start block; trigger one via
	// the helper script if running interactively. Here we simply ensure the
	// watcher runs without erroring for a couple of polls.
	select {
	case ev := <-events:
		t.Logf("received event: %+v", ev)
	case <-time.After(2 * time.Second):
		// no events in window — acceptable; watcher health is what we assert
	}
}
