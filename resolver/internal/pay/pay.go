// Package pay implements the off-chain micropayment vouchers a client signs to
// pay a resolver per query (FS §2.3, ResolverIncentives). A voucher authorizes
// a monotonically increasing cumulative amount on a channel; the resolver
// redeems the latest one on-chain. The signed digest is bound to the contract
// address and channel id, so a voucher cannot be replayed elsewhere.
//
// The digest must match ResolverIncentives.voucherDigest exactly:
//
//	inner  = keccak256(abi.encode(contract, id, cumulative))   // 3 × 32 bytes
//	digest = keccak256("\x19Ethereum Signed Message:\n32" || inner)
package pay

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// innerHash is keccak256(abi.encode(contract, id, cumulative)).
func innerHash(contract common.Address, id [32]byte, cumulative *big.Int) []byte {
	buf := make([]byte, 0, 96)
	buf = append(buf, common.LeftPadBytes(contract.Bytes(), 32)...)
	buf = append(buf, id[:]...)
	buf = append(buf, common.LeftPadBytes(cumulative.Bytes(), 32)...)
	return crypto.Keccak256(buf)
}

// Digest returns the EIP-191 voucher digest the contract verifies.
func Digest(contract common.Address, id [32]byte, cumulative *big.Int) []byte {
	return accounts.TextHash(innerHash(contract, id, cumulative))
}

// SignVoucher produces a 65-byte voucher signature (v in {27,28}) authorizing
// `cumulative` wei on channel `id` of `contract`.
func SignVoucher(key *ecdsa.PrivateKey, contract common.Address, id [32]byte, cumulative *big.Int) ([]byte, error) {
	sig, err := crypto.Sign(Digest(contract, id, cumulative), key)
	if err != nil {
		return nil, err
	}
	sig[64] += 27 // align with eth_sign / the contract's v handling
	return sig, nil
}

// RecoverVoucher returns the address that signed a voucher (mirrors the
// contract's ecrecover), used to validate a voucher before relying on it.
func RecoverVoucher(contract common.Address, id [32]byte, cumulative *big.Int, sig []byte) (common.Address, error) {
	s := make([]byte, len(sig))
	copy(s, sig)
	if len(s) == 65 && s[64] >= 27 {
		s[64] -= 27
	}
	pub, err := crypto.SigToPub(Digest(contract, id, cumulative), s)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}
