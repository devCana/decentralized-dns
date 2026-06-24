package zk

import (
	"bytes"
	"embed"
	"fmt"
	"sync"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
)

// Artifacts are produced by `make zk-setup` (resolver/cmd/zkgen) and
// committed. The Groth16 setup is a single-party dev ceremony — NOT a
// production-safe trusted setup; see the README security model.
//
//go:embed all:artifacts
var artifactsFS embed.FS

var (
	loadOnce sync.Once
	loadErr  error
	ccs      constraint.ConstraintSystem
	pk       groth16.ProvingKey
	vk       groth16.VerifyingKey
)

// loadArtifacts lazily deserializes the embedded circuit and keys.
func loadArtifacts() error {
	loadOnce.Do(func() {
		ccsBytes, err := artifactsFS.ReadFile("artifacts/record_v1.ccs")
		if err != nil {
			loadErr = fmt.Errorf("zk artifacts missing (run `make zk-setup`): %w", err)
			return
		}
		ccs = groth16.NewCS(ecc.BN254)
		if _, err := ccs.ReadFrom(bytes.NewReader(ccsBytes)); err != nil {
			loadErr = fmt.Errorf("read ccs: %w", err)
			return
		}
		pkBytes, err := artifactsFS.ReadFile("artifacts/record_v1.pk")
		if err != nil {
			loadErr = err
			return
		}
		pk = groth16.NewProvingKey(ecc.BN254)
		if _, err := pk.ReadFrom(bytes.NewReader(pkBytes)); err != nil {
			loadErr = fmt.Errorf("read pk: %w", err)
			return
		}
		vkBytes, err := artifactsFS.ReadFile("artifacts/record_v1.vk")
		if err != nil {
			loadErr = err
			return
		}
		vk = groth16.NewVerifyingKey(ecc.BN254)
		if _, err := vk.ReadFrom(bytes.NewReader(vkBytes)); err != nil {
			loadErr = fmt.Errorf("read vk: %w", err)
			return
		}
	})
	return loadErr
}

// VerifyingKey exposes the embedded VK (ddns-lookup offline verification).
func VerifyingKey() (groth16.VerifyingKey, error) {
	if err := loadArtifacts(); err != nil {
		return nil, err
	}
	return vk, nil
}
