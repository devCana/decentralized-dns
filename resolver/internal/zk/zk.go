// Package zk implements the record-commitment proof system (HLD §3.5):
// records anchor a MiMC commitment on-chain; the resolver proves with
// Groth16 (BN254) that the payload it serves hashes to that commitment.
//
// Witness layout: [length, chunk_0 .. chunk_{MaxChunks-1}] where chunks
// are 31-byte big-endian slices of the canonical record message
// (pki.RecordMessage), zero-padded. The length element binds the padding.
// Public input: the MiMC hash of that sequence (= on-chain commitment).
package zk

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	gcmimc "github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

const (
	// ChunkSize is bytes per field element (31 < 254-bit BN254 scalar).
	ChunkSize = 31
	// MaxChunks fixes the circuit size; payloads up to 496 bytes.
	MaxChunks = 16
	// MaxPayload is the largest record message that can be committed.
	MaxPayload = ChunkSize * MaxChunks
)

// ErrPayloadTooLarge is returned for messages beyond MaxPayload; such
// records simply carry no commitment (commitment = 0, no proof).
var ErrPayloadTooLarge = errors.New("record payload exceeds ZK commitment capacity")

// RecordCommitmentCircuit proves MiMC(length, chunks...) == Commitment.
type RecordCommitmentCircuit struct {
	Length     frontend.Variable            `gnark:"length"`
	Chunks     [MaxChunks]frontend.Variable `gnark:"chunks"`
	Commitment frontend.Variable            `gnark:"commitment,public"`
}

// Define implements frontend.Circuit.
func (c *RecordCommitmentCircuit) Define(api frontend.API) error {
	h, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}
	h.Write(c.Length)
	h.Write(c.Chunks[:]...)
	api.AssertIsEqual(h.Sum(), c.Commitment)
	return nil
}

// PackMessage maps msg into the fixed witness layout [len, chunk...].
func PackMessage(msg []byte) ([]*big.Int, error) {
	if len(msg) > MaxPayload {
		return nil, fmt.Errorf("%w: %d > %d bytes", ErrPayloadTooLarge, len(msg), MaxPayload)
	}
	out := make([]*big.Int, 1+MaxChunks)
	out[0] = big.NewInt(int64(len(msg)))
	for i := 0; i < MaxChunks; i++ {
		lo := i * ChunkSize
		chunk := make([]byte, 0, ChunkSize)
		if lo < len(msg) {
			hi := min(lo+ChunkSize, len(msg))
			chunk = msg[lo:hi]
		}
		out[i+1] = new(big.Int).SetBytes(chunk)
	}
	return out, nil
}

// Commitment computes the on-chain MiMC commitment of msg, matching the
// in-circuit hash (one field element per Write).
func Commitment(msg []byte) ([32]byte, error) {
	var zero [32]byte
	elements, err := PackMessage(msg)
	if err != nil {
		return zero, err
	}
	h := gcmimc.NewMiMC()
	for _, el := range elements {
		var fe fr.Element
		fe.SetBigInt(el)
		b := fe.Bytes()
		if _, err := h.Write(b[:]); err != nil {
			return zero, err
		}
	}
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out, nil
}

// assignment builds the full witness for msg.
func assignment(msg []byte) (*RecordCommitmentCircuit, error) {
	elements, err := PackMessage(msg)
	if err != nil {
		return nil, err
	}
	commitment, err := Commitment(msg)
	if err != nil {
		return nil, err
	}
	var c RecordCommitmentCircuit
	c.Length = elements[0]
	for i := 0; i < MaxChunks; i++ {
		c.Chunks[i] = elements[i+1]
	}
	c.Commitment = new(big.Int).SetBytes(commitment[:])
	return &c, nil
}
