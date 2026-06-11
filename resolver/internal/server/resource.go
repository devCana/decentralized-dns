package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/devCana/decentralized-dns/resolver/internal/query"
	bttorrent "github.com/devCana/decentralized-dns/resolver/internal/torrent"
)

const resourceFetchTimeout = 60 * time.Second

// ResourceFetcher is the torrent fetch API needed by the REST resource flow.
type ResourceFetcher interface {
	Fetch(ctx context.Context, infoHashHex, expectedSHAHex string, peers []string) ([]byte, error)
}

func parseResourceQuery(r *http.Request) (query.Query, error) {
	params := r.URL.Query()
	recordType := params.Get("type")
	if recordType == "" {
		recordType = "ResourceRef"
	}
	if recordType != "ResourceRef" {
		return query.Query{}, fmt.Errorf("resource endpoint only resolves ResourceRef records")
	}
	pairs, err := query.ParsePairs(params.Get("selector"))
	if err != nil {
		return query.Query{}, err
	}
	return query.New(params.Get("name"), recordType, pairs)
}

// handleResource serves GET /resource?name=&selector=&peer=. It resolves a
// ResourceRef, verifies the owner signature through HandleQuery, fetches the
// torrent payload, and returns the exact verified bytes signed in headers.
func (s *Server) handleResource(w http.ResponseWriter, r *http.Request) {
	if s.resource == nil {
		writeError(w, http.StatusServiceUnavailable, "resource_engine_unavailable", "resource fetching is not enabled")
		return
	}
	q, err := parseResourceQuery(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_query", err.Error())
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), resourceFetchTimeout)
	defer cancel()

	res, err := s.HandleQuery(ctx, q)
	if err != nil {
		writeError(w, http.StatusBadGateway, "chain_error", err.Error())
		return
	}
	if !res.Found() {
		writeError(w, http.StatusNotFound, "no_match", "ResourceRef record was not found")
		return
	}
	if !res.Result.OwnerSigValid {
		writeError(w, http.StatusBadGateway, "owner_signature_invalid", "ResourceRef owner signature failed verification")
		return
	}

	rec := res.Result.Record
	infoHash, ok := rec.Field("infoHash")
	if !ok || strings.TrimSpace(infoHash) == "" {
		writeError(w, http.StatusBadGateway, "bad_resource_ref", "ResourceRef is missing infoHash")
		return
	}
	sha, ok := rec.Field("sha256")
	if !ok || strings.TrimSpace(sha) == "" {
		writeError(w, http.StatusBadGateway, "bad_resource_ref", "ResourceRef is missing sha256")
		return
	}
	contentType, ok := rec.Field("contentType")
	if !ok || strings.TrimSpace(contentType) == "" {
		contentType = "application/octet-stream"
	}
	if strings.ContainsAny(contentType, "\r\n") {
		writeError(w, http.StatusBadGateway, "bad_resource_ref", "ResourceRef contentType is invalid")
		return
	}

	body, err := s.resource.Fetch(ctx, infoHash, sha, r.URL.Query()["peer"])
	if err != nil {
		switch {
		case errors.Is(err, bttorrent.ErrHashMismatch):
			writeError(w, http.StatusBadGateway, "resource_hash_mismatch", "torrent payload did not match on-chain sha256")
		case errors.Is(err, bttorrent.ErrTooLarge):
			writeError(w, http.StatusBadGateway, "resource_too_large", err.Error())
		case errors.Is(err, context.DeadlineExceeded), errors.Is(err, context.Canceled):
			writeError(w, http.StatusGatewayTimeout, "resource_timeout", err.Error())
		default:
			writeError(w, http.StatusBadGateway, "resource_fetch_failed", err.Error())
		}
		return
	}

	sig := hexutil.Encode(s.identity.Sign(body))
	h := w.Header()
	h.Set("Content-Type", contentType)
	h.Set("X-DDNS-Resolver", s.identity.PublicKeyHex())
	h.Set("X-DDNS-Signature", sig)
	h.Set("X-DDNS-Owner", res.Result.Owner.Hex())
	h.Set("X-DDNS-PubKey", hexutil.Encode(res.Result.PubKey))
	h.Set("X-DDNS-InfoHash", infoHash)
	h.Set("X-DDNS-SHA256", sha)
	h.Set("X-DDNS-OwnerSig-Verified", "true")
	if len(res.Result.ZKProof) > 0 {
		h.Set("X-DDNS-ZKProof", hexutil.Encode(res.Result.ZKProof))
	}
	if rec.TTL > 0 {
		h.Set("Cache-Control", fmt.Sprintf("public, max-age=%d", rec.TTL))
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}
