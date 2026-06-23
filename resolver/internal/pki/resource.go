package pki

import (
	"bytes"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/binary"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	errBadResolverKey = errors.New("pki: bad resolver public key encoding")
	errBadResourceSig = errors.New("pki: resource signature invalid")
)

// ResourceManifest builds the canonical bytes the resolver signs for a
// /resource response. It binds the served body (by its SHA-256) to every piece
// of provenance the resolver asserts in X-DDNS-* headers — owner, pubKey,
// infoHash, on-chain sha256, content type, owner-signature verdict and ZK
// proof — so a man-in-the-middle cannot rewrite any of them without
// invalidating the resolver signature. Length-prefixed fields keep the
// encoding unambiguous; both the resolver and ddns-fetch build it identically.
func ResourceManifest(owner, pubKeyHex, infoHash, sha256Hex, contentType string, ownerSigVerified bool, zkProofHex string, body []byte) []byte {
	var buf bytes.Buffer
	buf.WriteString("ddns-resource-v1")
	writeStr := func(s string) {
		var l [4]byte
		binary.BigEndian.PutUint32(l[:], uint32(len(s)))
		buf.Write(l[:])
		buf.WriteString(s)
	}
	writeStr(owner)
	writeStr(pubKeyHex)
	writeStr(infoHash)
	writeStr(sha256Hex)
	writeStr(contentType)
	if ownerSigVerified {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}
	writeStr(zkProofHex)
	sum := sha256.Sum256(body)
	buf.Write(sum[:])
	return buf.Bytes()
}

// VerifyResourceSignature checks that sigHex is the resolver identity's
// signature over the manifest derived from these fields and body. Used by
// ddns-fetch to authenticate a /resource download end-to-end.
func VerifyResourceSignature(resolverPubHex, sigHex string, owner, pubKeyHex, infoHash, sha256Hex, contentType string, ownerSigVerified bool, zkProofHex string, body []byte) error {
	pub, err := hexutil.Decode(resolverPubHex)
	if err != nil || len(pub) != ed25519.PublicKeySize {
		return errBadResolverKey
	}
	sig, err := hexutil.Decode(sigHex)
	if err != nil || len(sig) != ed25519.SignatureSize {
		return errBadResourceSig
	}
	manifest := ResourceManifest(owner, pubKeyHex, infoHash, sha256Hex, contentType, ownerSigVerified, zkProofHex, body)
	if !ed25519.Verify(ed25519.PublicKey(pub), manifest, sig) {
		return errBadResourceSig
	}
	return nil
}
