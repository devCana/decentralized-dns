// Package chain wraps go-ethereum's ethclient plus the abigen-generated
// contract bindings behind a domain-specific API (HLD §3.4): typed reads,
// RPC retry/back-off, and record-event watching for cache invalidation.
package chain

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/devCana/decentralized-dns/resolver/internal/chain/bindings"
)

// Record is the resolver-level view of an on-chain record.
// FieldNames/FieldValues keep on-chain order — the canonical payload for
// signatures and ZK commitments depends on it.
type Record struct {
	Type       string   `json:"type"`
	Selector   string   `json:"selector"`
	FieldNames []string `json:"fieldNames"`
	FieldVals  []string `json:"fieldValues"`
	TTL        uint32   `json:"ttl"`
	Generation uint64   `json:"generation"`
	OwnerSig   []byte   `json:"ownerSig"`
	Commitment [32]byte `json:"commitment"`
	Exists     bool     `json:"exists"`
}

// Field returns the value of a named field and whether it is present.
func (r *Record) Field(name string) (string, bool) {
	for i, n := range r.FieldNames {
		if n == name {
			return r.FieldVals[i], true
		}
	}
	return "", false
}

// ResolveResult bundles a record with the domain identity needed to verify
// the owner signature (single RPC round-trip via the contract's resolve()).
type ResolveResult struct {
	Record Record         `json:"record"`
	Owner  common.Address `json:"owner"`
	PubKey []byte         `json:"pubKey"`
	// OwnerSigValid is set by the resolver at cache-fill time (UC-5):
	// whether Record.OwnerSig verifies against Owner/PubKey.
	OwnerSigValid bool `json:"ownerSigValid"`
}

// Domain mirrors NamespaceDApp.getDomain.
type Domain struct {
	Owner      common.Address `json:"owner"`
	PubKey     []byte         `json:"pubKey"`
	Expiry     uint64         `json:"expiry"`
	Generation uint64         `json:"generation"`
}

// EventKind discriminates RecordEvent.
type EventKind string

const (
	EventRecordSet     EventKind = "record_set"
	EventRecordRemoved EventKind = "record_removed"
	EventTransferred   EventKind = "transferred"
	EventRegistered    EventKind = "registered"
)

// RecordEvent is a normalized contract event used for cache invalidation
// (HLD §3.3 proactive invalidation). Transferred events carry only the
// name hash; consumers map hashes back to cached names.
type RecordEvent struct {
	Kind       EventKind
	Name       string // empty for Transferred (hash only)
	NameHash   [32]byte
	RecordType string // set/removed only
	Selector   string // set/removed only
	Block      uint64
}

func fromBinding(r bindings.NamespaceDAppRecord) Record {
	return Record{
		Type:       r.RecordType,
		Selector:   r.Selector,
		FieldNames: r.FieldNames,
		FieldVals:  r.FieldValues,
		TTL:        r.Ttl,
		Generation: r.Generation,
		OwnerSig:   r.OwnerSig,
		Commitment: r.Commitment,
		Exists:     r.Exists,
	}
}
