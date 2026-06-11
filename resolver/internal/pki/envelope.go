package pki

import (
	"crypto/ed25519"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Envelope wraps a response body with the resolver's identity signature.
// The signature covers the exact Data bytes as transmitted, so clients
// verify without any re-canonicalization.
type Envelope struct {
	Data      json.RawMessage `json:"data"`
	Resolver  string          `json:"resolver"`  // 0x-hex ed25519 public key
	Signature string          `json:"signature"` // 0x-hex ed25519 sig over Data
}

// SealEnvelope marshals v and signs the resulting bytes.
func (id *Identity) SealEnvelope(v any) (*Envelope, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return &Envelope{
		Data:      data,
		Resolver:  id.PublicKeyHex(),
		Signature: hexutil.Encode(id.Sign(data)),
	}, nil
}

// VerifyEnvelope checks the resolver signature over env.Data. Callers
// compare env.Resolver against their trusted resolver list (bootstrap).
func VerifyEnvelope(env *Envelope) error {
	pub, err := hexutil.Decode(env.Resolver)
	if err != nil || len(pub) != ed25519.PublicKeySize {
		return fmt.Errorf("bad resolver key %q", env.Resolver)
	}
	sig, err := hexutil.Decode(env.Signature)
	if err != nil || len(sig) != ed25519.SignatureSize {
		return fmt.Errorf("bad signature encoding %q", env.Signature)
	}
	if !ed25519.Verify(ed25519.PublicKey(pub), env.Data, sig) {
		return errors.New("envelope signature invalid")
	}
	return nil
}
