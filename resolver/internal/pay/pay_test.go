package pay

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVoucherSignRecoverRoundTrip(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	client := crypto.PubkeyToAddress(key.PublicKey)
	contract := common.HexToAddress("0x000000000000000000000000000000000000c0de")
	var id [32]byte
	copy(id[:], []byte("channel-one"))
	amount := big.NewInt(1234567890)

	sig, err := SignVoucher(key, contract, id, amount)
	if err != nil {
		t.Fatal(err)
	}
	if len(sig) != 65 {
		t.Fatalf("sig length = %d, want 65", len(sig))
	}
	if v := sig[64]; v != 27 && v != 28 {
		t.Errorf("v = %d, want 27/28", v)
	}

	got, err := RecoverVoucher(contract, id, amount, sig)
	if err != nil {
		t.Fatal(err)
	}
	if got != client {
		t.Errorf("recovered %s, want %s", got, client)
	}

	// Any change to the authorized amount must invalidate the voucher.
	other, _ := RecoverVoucher(contract, id, big.NewInt(1234567891), sig)
	if other == client {
		t.Error("voucher verified against a different amount")
	}
	// Same for a different contract address (replay protection).
	elsewhere, _ := RecoverVoucher(common.HexToAddress("0xdead"), id, amount, sig)
	if elsewhere == client {
		t.Error("voucher verified against a different contract")
	}
}

func TestRecoverVoucherRejectsHighS(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	contract := common.HexToAddress("0x000000000000000000000000000000000000c0de")
	var id [32]byte
	amount := big.NewInt(100)

	sig, err := SignVoucher(key, contract, id, amount) // canonical low-s
	if err != nil {
		t.Fatal(err)
	}
	// Malleate into the equivalent high-s signature: s' = N - s, flip v.
	n := crypto.S256().Params().N
	sPrime := new(big.Int).Sub(n, new(big.Int).SetBytes(sig[32:64]))
	mal := make([]byte, 65)
	copy(mal, sig[:32])
	copy(mal[32:64], common.LeftPadBytes(sPrime.Bytes(), 32))
	mal[64] = sig[64] ^ 1

	if _, err := RecoverVoucher(contract, id, amount, mal); err == nil {
		t.Error("non-canonical high-s voucher must be rejected")
	}
}
