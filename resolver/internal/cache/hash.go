package cache

import (
	"golang.org/x/crypto/sha3"
)

// nameHash computes keccak256(name), matching the contract's domain key.
func nameHash(name string) (out [32]byte) {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(name))
	copy(out[:], h.Sum(nil))
	return out
}
