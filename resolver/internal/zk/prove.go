package zk

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
)

// Prove generates a Groth16 proof that msg hashes to its MiMC commitment.
func Prove(msg []byte) (groth16.Proof, error) {
	if err := loadArtifacts(); err != nil {
		return nil, err
	}
	assign, err := assignment(msg)
	if err != nil {
		return nil, err
	}
	w, err := frontend.NewWitness(assign, ecc.BN254.ScalarField())
	if err != nil {
		return nil, fmt.Errorf("witness: %w", err)
	}
	proof, err := groth16.Prove(ccs, pk, w)
	if err != nil {
		return nil, fmt.Errorf("prove: %w", err)
	}
	return proof, nil
}

// publicWitness builds the public witness for a commitment.
func publicWitness(commitment [32]byte) (witness.Witness, error) {
	assign := &RecordCommitmentCircuit{Commitment: new(big.Int).SetBytes(commitment[:])}
	return frontend.NewWitness(assign, ecc.BN254.ScalarField(), frontend.PublicOnly())
}

// Verify checks proof against an on-chain commitment with the embedded VK.
func Verify(proof groth16.Proof, commitment [32]byte) error {
	if err := loadArtifacts(); err != nil {
		return err
	}
	pubW, err := publicWitness(commitment)
	if err != nil {
		return err
	}
	return groth16.Verify(proof, vk, pubW)
}

// SolidityCalldata serializes proof for the exported ZKVerifier contract:
// 8 uint256 words (A.X, A.Y, B.X1, B.X0, B.Y1, B.Y0, C.X, C.Y).
func SolidityCalldata(proof groth16.Proof) ([]byte, error) {
	p, ok := proof.(*groth16bn254.Proof)
	if !ok {
		return nil, fmt.Errorf("unexpected proof type %T", proof)
	}
	return p.MarshalSolidity(), nil
}

// ProofFromSolidityCalldata reconstructs a proof from the 256-byte wire
// form (used by ddns-lookup to re-verify resolver-supplied proofs).
func ProofFromSolidityCalldata(data []byte) (groth16.Proof, error) {
	if len(data) != 256 {
		return nil, fmt.Errorf("want 256-byte proof calldata, got %d", len(data))
	}
	var p groth16bn254.Proof
	if err := p.Ar.X.SetBytesCanonical(data[0:32]); err != nil {
		return nil, err
	}
	if err := p.Ar.Y.SetBytesCanonical(data[32:64]); err != nil {
		return nil, err
	}
	if err := p.Bs.X.A1.SetBytesCanonical(data[64:96]); err != nil {
		return nil, err
	}
	if err := p.Bs.X.A0.SetBytesCanonical(data[96:128]); err != nil {
		return nil, err
	}
	if err := p.Bs.Y.A1.SetBytesCanonical(data[128:160]); err != nil {
		return nil, err
	}
	if err := p.Bs.Y.A0.SetBytesCanonical(data[160:192]); err != nil {
		return nil, err
	}
	if err := p.Krs.X.SetBytesCanonical(data[192:224]); err != nil {
		return nil, err
	}
	if err := p.Krs.Y.SetBytesCanonical(data[224:256]); err != nil {
		return nil, err
	}
	return &p, nil
}
