package server

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/devCana/decentralized-dns/resolver/internal/cache"
	"github.com/devCana/decentralized-dns/resolver/internal/chain"
	"github.com/devCana/decentralized-dns/resolver/internal/config"
)

// fakeChain is an in-memory ChainReader for handler tests.
type fakeChain struct {
	resolveCalls int
	lastSelector string
	records      map[string]chain.ResolveResult // key: name|type|selector
	domains      map[string]chain.Domain
	types        []string
}

func key(name, typ, sel string) string { return name + "|" + typ + "|" + sel }

func (f *fakeChain) Resolve(_ context.Context, name, recordType, selector string) (*chain.ResolveResult, error) {
	f.resolveCalls++
	f.lastSelector = selector
	if res, ok := f.records[key(name, recordType, selector)]; ok {
		return &res, nil
	}
	return &chain.ResolveResult{}, nil // exists=false, zero owner
}

func (f *fakeChain) GetDomain(_ context.Context, name string) (*chain.Domain, error) {
	d := f.domains[name]
	return &d, nil
}

func (f *fakeChain) ListRecords(_ context.Context, name string) ([]chain.Record, error) {
	var out []chain.Record
	for _, r := range f.records {
		out = append(out, r.Record)
	}
	return out, nil
}

func (f *fakeChain) ListTypes(_ context.Context) ([]string, error) { return f.types, nil }

func (f *fakeChain) ChainHead(_ context.Context) (uint64, error) { return 7, nil }

var owner = common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")

func seededFake() *fakeChain {
	rec := chain.Record{
		Type: "A", Selector: "", FieldNames: []string{"address"},
		FieldVals: []string{"93.184.216.34"}, TTL: 3600, Generation: 1,
		OwnerSig: []byte{0xab}, Exists: true,
	}
	svc := chain.Record{
		Type: "SVC", Selector: "port=443&service=HTTP&transport=QUIC",
		FieldNames: []string{"target", "service", "transport", "port"},
		FieldVals:  []string{"web.example", "HTTP", "QUIC", "443"},
		TTL:        600, Generation: 1, OwnerSig: []byte{0xcd}, Exists: true,
	}
	return &fakeChain{
		records: map[string]chain.ResolveResult{
			key("example", "A", ""):             {Record: rec, Owner: owner, PubKey: []byte{0x04, 0x01}},
			key("example", "SVC", svc.Selector): {Record: svc, Owner: owner, PubKey: []byte{0x04, 0x01}},
		},
		domains: map[string]chain.Domain{
			"example": {Owner: owner, PubKey: []byte{0x04, 0x01}, Expiry: 1<<62 - 1, Generation: 1},
		},
		types: []string{"A", "AAAA", "MX", "SVC", "ResourceRef"},
	}
}

func newTestServer(t *testing.T, fc ChainReader, rps, burst int) *Server {
	t.Helper()
	c, err := cache.New[*chain.ResolveResult](64)
	if err != nil {
		t.Fatal(err)
	}
	s := &Server{
		cfg:   &config.Config{RateRPS: rps, RateBurst: burst},
		log:   slog.New(slog.DiscardHandler),
		chain: fc,
		cache: c,
		mux:   http.NewServeMux(),
	}
	s.registerRoutes()
	return s
}

func get(t *testing.T, s *Server, url string) (*httptest.ResponseRecorder, map[string]any) {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	s.mux.ServeHTTP(w, req)
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON from %s: %v\n%s", url, err, w.Body.String())
	}
	return w, body
}

func TestResolveReadThroughAndCache(t *testing.T) {
	fc := seededFake()
	s := newTestServer(t, fc, 100, 100)

	w, body := get(t, s, "/resolve?name=example&type=A")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d: %v", w.Code, body)
	}
	if body["found"] != true || body["cached"] != false {
		t.Fatalf("first hit: %v", body)
	}
	rec := body["record"].(map[string]any)
	if rec["fieldValues"].([]any)[0] != "93.184.216.34" {
		t.Errorf("record = %v", rec)
	}
	if body["owner"] != owner.Hex() {
		t.Errorf("owner = %v", body["owner"])
	}

	// second hit must come from cache
	_, body = get(t, s, "/resolve?name=example&type=A")
	if body["cached"] != true {
		t.Errorf("second hit not cached: %v", body)
	}
	if fc.resolveCalls != 1 {
		t.Errorf("resolveCalls = %d, want 1", fc.resolveCalls)
	}
}

func TestResolveSelectorNormalization(t *testing.T) {
	fc := seededFake()
	s := newTestServer(t, fc, 100, 100)

	// explicit params, mixed case, unsorted -> canonical selector
	w, body := get(t, s, "/resolve?name=example&type=SVC&transport=quic&service=http&port=443")
	if w.Code != http.StatusOK || body["found"] != true {
		t.Fatalf("status=%d body=%v", w.Code, body)
	}
	if fc.lastSelector != "port=443&service=HTTP&transport=QUIC" {
		t.Errorf("selector sent to chain = %q", fc.lastSelector)
	}

	// raw selector string form works too
	_, body = get(t, s, "/resolve?name=example&type=SVC&selector=service%3Dhttp%26port%3D443%26transport%3Dquic")
	if body["found"] != true {
		t.Errorf("raw selector form: %v", body)
	}
}

func TestResolveTypedNoMatch(t *testing.T) {
	s := newTestServer(t, seededFake(), 100, 100)
	w, body := get(t, s, "/resolve?name=example&type=MX")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if body["found"] != false || body["error"] != "no_match" {
		t.Errorf("body = %v", body)
	}
	if _, hasRecord := body["record"]; hasRecord {
		t.Error("no_match must not carry a record")
	}
}

func TestResolveValidation(t *testing.T) {
	s := newTestServer(t, seededFake(), 100, 100)
	urls := []string{
		"/resolve",                                                // missing name+type
		"/resolve?name=UPPER&type=A",                              // bad name
		"/resolve?name=example&type=bad%20type",                   // bad type
		"/resolve?name=example&type=SVC&port=99999",               // bad port
		"/resolve?name=example&type=SVC&transport=SCTP",           // bad transport
		"/resolve?name=example&type=SVC&selector=port%3D1&port=2", // dup key
	}
	for _, u := range urls {
		w, body := get(t, s, u)
		if w.Code != http.StatusBadRequest || body["error"] != "invalid_query" {
			t.Errorf("%s: status=%d body=%v", u, w.Code, body)
		}
	}
}

func TestRateLimit429(t *testing.T) {
	s := newTestServer(t, seededFake(), 1, 2) // 1 rps, burst 2

	codes := []int{}
	for range 4 {
		w, _ := get(t, s, "/resolve?name=example&type=A")
		codes = append(codes, w.Code)
	}
	if codes[0] != http.StatusOK || codes[1] != http.StatusOK {
		t.Fatalf("burst should pass: %v", codes)
	}
	if codes[2] != http.StatusTooManyRequests || codes[3] != http.StatusTooManyRequests {
		t.Fatalf("expected 429s after burst: %v", codes)
	}

	// healthz is exempt from limiting
	w, _ := get(t, s, "/healthz")
	if w.Code == http.StatusTooManyRequests {
		t.Error("healthz must not be rate limited")
	}
}

func TestDomainEndpoint(t *testing.T) {
	s := newTestServer(t, seededFake(), 100, 100)

	w, body := get(t, s, "/domains/example")
	if w.Code != http.StatusOK || body["active"] != true {
		t.Fatalf("status=%d body=%v", w.Code, body)
	}
	if len(body["records"].([]any)) != 2 {
		t.Errorf("records = %v", body["records"])
	}

	w, body = get(t, s, "/domains/ghost")
	if w.Code != http.StatusNotFound || body["error"] != "not_registered" {
		t.Errorf("unregistered: status=%d body=%v", w.Code, body)
	}

	w, _ = get(t, s, "/domains/BAD!")
	if w.Code != http.StatusBadRequest {
		t.Errorf("bad name: status=%d", w.Code)
	}
}

func TestTypesEndpoint(t *testing.T) {
	s := newTestServer(t, seededFake(), 100, 100)
	w, body := get(t, s, "/types")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	if len(body["types"].([]any)) != 5 {
		t.Errorf("types = %v", body["types"])
	}
}
