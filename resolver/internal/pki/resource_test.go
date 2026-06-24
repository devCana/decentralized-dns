package pki

import (
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// manifestFields mirrors the provenance the resolver asserts for a /resource
// response. Each field is bound into the signed manifest, so flipping any one
// of them must invalidate the resolver signature.
type manifestFields struct {
	owner, pubKeyHex, infoHash, sha256Hex, contentType string
	ownerSigVerified                                   bool
	zkProofHex                                         string
	body                                               []byte
}

func sampleFields() manifestFields {
	return manifestFields{
		owner:            "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		pubKeyHex:        "0x04bfcab2...",
		infoHash:         "f726b0293905877f472749b3cb8699f3f8035b58",
		sha256Hex:        "7cd536fc074e48203a111fc4ed5f570813bfdd85b435339fd93181e80d3d35be",
		contentType:      "text/html; charset=utf-8",
		ownerSigVerified: true,
		zkProofHex:       "0xabcdef",
		body:             []byte("<h1>served from BitTorrent, verified on-chain</h1>"),
	}
}

func (f manifestFields) manifest() []byte {
	return ResourceManifest(f.owner, f.pubKeyHex, f.infoHash, f.sha256Hex, f.contentType, f.ownerSigVerified, f.zkProofHex, f.body)
}

func (f manifestFields) verify(resolverPubHex, sigHex string) error {
	return VerifyResourceSignature(resolverPubHex, sigHex, f.owner, f.pubKeyHex, f.infoHash, f.sha256Hex, f.contentType, f.ownerSigVerified, f.zkProofHex, f.body)
}

func newIdentity(t *testing.T) *Identity {
	t.Helper()
	id, err := LoadOrCreateIdentity(filepath.Join(t.TempDir(), "resolver.key"))
	if err != nil {
		t.Fatal(err)
	}
	return id
}

func TestResourceManifestBindsEveryField(t *testing.T) {
	base := sampleFields().manifest()

	// Each mutation must produce a different manifest — i.e. every field is
	// actually committed, so a MITM cannot rewrite headers without detection.
	mutations := map[string]func(*manifestFields){
		"owner":            func(f *manifestFields) { f.owner = "0xdeadbeef" },
		"pubKey":           func(f *manifestFields) { f.pubKeyHex = "0x04ffff" },
		"infoHash":         func(f *manifestFields) { f.infoHash = "0000000000000000000000000000000000000000" },
		"sha256":           func(f *manifestFields) { f.sha256Hex = "00" + sampleFields().sha256Hex[2:] },
		"contentType":      func(f *manifestFields) { f.contentType = "application/octet-stream" },
		"ownerSigVerified": func(f *manifestFields) { f.ownerSigVerified = false },
		"zkProof":          func(f *manifestFields) { f.zkProofHex = "0x000000" },
		"body":             func(f *manifestFields) { f.body = []byte("tampered body") },
	}
	for name, mutate := range mutations {
		f := sampleFields()
		mutate(&f)
		if string(f.manifest()) == string(base) {
			t.Errorf("mutating %q did not change the manifest (field not bound)", name)
		}
	}
}

func TestResourceManifestDeterministic(t *testing.T) {
	if string(sampleFields().manifest()) != string(sampleFields().manifest()) {
		t.Error("manifest is not deterministic for identical inputs")
	}
}

// TestResourceManifestNoFieldConfusion guards the length-prefixed encoding:
// shifting a byte across a field boundary (here, ""+"ab" vs "a"+"b") must not
// collide, or an attacker could smuggle data between fields.
func TestResourceManifestNoFieldConfusion(t *testing.T) {
	a := ResourceManifest("", "ab", "h", "s", "c", true, "z", nil)
	b := ResourceManifest("a", "b", "h", "s", "c", true, "z", nil)
	if string(a) == string(b) {
		t.Error("ambiguous encoding: adjacent fields collide across a boundary")
	}
}

func TestVerifyResourceSignatureRoundTrip(t *testing.T) {
	id := newIdentity(t)
	f := sampleFields()
	sigHex := hexutil.Encode(id.Sign(f.manifest()))

	if err := f.verify(id.PublicKeyHex(), sigHex); err != nil {
		t.Fatalf("valid resource signature rejected: %v", err)
	}

	// A signature from a different resolver identity must not verify.
	other := newIdentity(t)
	if err := f.verify(other.PublicKeyHex(), sigHex); err == nil {
		t.Error("signature accepted under the wrong resolver key")
	}
}

func TestVerifyResourceSignatureRejectsTamper(t *testing.T) {
	id := newIdentity(t)
	signed := sampleFields()
	sigHex := hexutil.Encode(id.Sign(signed.manifest()))
	pub := id.PublicKeyHex()

	// Verifying the original signature against any single mutated field must
	// fail: the verifier rebuilds the manifest and the hash won't match.
	mutations := map[string]func(*manifestFields){
		"owner":            func(f *manifestFields) { f.owner = "0xdeadbeef" },
		"pubKey":           func(f *manifestFields) { f.pubKeyHex = "0x04ffff" },
		"infoHash":         func(f *manifestFields) { f.infoHash = "ffffffffffffffffffffffffffffffffffffffff" },
		"sha256":           func(f *manifestFields) { f.sha256Hex = "00" + signed.sha256Hex[2:] },
		"contentType":      func(f *manifestFields) { f.contentType = "application/octet-stream" },
		"ownerSigVerified": func(f *manifestFields) { f.ownerSigVerified = false },
		"zkProof":          func(f *manifestFields) { f.zkProofHex = "0xbad" },
		"body":             func(f *manifestFields) { f.body = []byte("tampered") },
	}
	for name, mutate := range mutations {
		f := sampleFields()
		mutate(&f)
		if err := f.verify(pub, sigHex); err == nil {
			t.Errorf("tampered %q field accepted", name)
		}
	}
}

func TestVerifyResourceSignatureRejectsMalformedInputs(t *testing.T) {
	id := newIdentity(t)
	f := sampleFields()
	sigHex := hexutil.Encode(id.Sign(f.manifest()))

	cases := map[string]struct{ pub, sig string }{
		"non-hex resolver key":      {"not-hex", sigHex},
		"resolver key wrong length": {"0x1234", sigHex},
		"non-hex signature":         {id.PublicKeyHex(), "not-hex"},
		"signature wrong length":    {id.PublicKeyHex(), "0xabcd"},
		"empty signature":           {id.PublicKeyHex(), ""},
	}
	for name, c := range cases {
		if err := f.verify(c.pub, c.sig); err == nil {
			t.Errorf("%s accepted", name)
		}
	}
}
