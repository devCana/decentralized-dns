// Package pki implements the two key domains of HLD §3.3: the resolver's
// ed25519 identity (signs every response envelope) and owner secp256k1
// record signatures verified against the on-chain pubKey (UC-5).
package pki

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Identity is the resolver's ed25519 signing identity.
type Identity struct {
	priv ed25519.PrivateKey
}

// LoadOrCreateIdentity reads a 32-byte hex seed from path, generating and
// persisting a fresh one (0600) on first boot.
func LoadOrCreateIdentity(path string) (*Identity, error) {
	data, err := os.ReadFile(path)
	switch {
	case err == nil:
		seed, decErr := hex.DecodeString(strings.TrimPrefix(strings.TrimSpace(string(data)), "0x"))
		if decErr != nil || len(seed) != ed25519.SeedSize {
			return nil, fmt.Errorf("keystore %s: want %d-byte hex ed25519 seed", path, ed25519.SeedSize)
		}
		return &Identity{priv: ed25519.NewKeyFromSeed(seed)}, nil
	case os.IsNotExist(err):
		seed := make([]byte, ed25519.SeedSize)
		if _, err := rand.Read(seed); err != nil {
			return nil, err
		}
		if dir := filepath.Dir(path); dir != "." {
			if err := os.MkdirAll(dir, 0o700); err != nil {
				return nil, err
			}
		}
		if err := os.WriteFile(path, []byte(hex.EncodeToString(seed)+"\n"), 0o600); err != nil {
			return nil, err
		}
		return &Identity{priv: ed25519.NewKeyFromSeed(seed)}, nil
	default:
		return nil, err
	}
}

// Sign signs message with the resolver identity key.
func (id *Identity) Sign(message []byte) []byte { return ed25519.Sign(id.priv, message) }

// PublicKey returns the 32-byte ed25519 public key.
func (id *Identity) PublicKey() ed25519.PublicKey {
	return id.priv.Public().(ed25519.PublicKey)
}

// PublicKeyHex returns the public key as 0x-hex.
func (id *Identity) PublicKeyHex() string {
	return "0x" + hex.EncodeToString(id.PublicKey())
}
