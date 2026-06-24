// Package contenttype implements Resource Type Validation (Functional Spec
// §2.2): a check that a BitTorrent-served file's actual bytes are consistent
// with the media type its owner registered on-chain. HLD open issue #3 weighs
// validating inside the resolver (local content sniffing) against a trusted
// external endpoint; we take the decentralization-pure option and sniff
// locally with net/http's algorithm, so no third party is trusted.
//
// The same logic runs in two places: the resolver flags (or, in strict mode,
// rejects) a mismatched file at serve time, and ddns-fetch re-checks the bytes
// it downloaded against the resolver-signed content type — a fully trustless
// client-side verdict that needs no cooperation from the resolver.
package contenttype

import (
	"mime"
	"net/http"
	"strings"
)

// Result is the outcome of validating a body against a declared media type.
type Result struct {
	Declared string `json:"declared"` // media type registered on-chain
	Detected string `json:"detected"` // sniffed from the actual bytes
	OK       bool   `json:"ok"`
	Reason   string `json:"reason"`
}

// Validate sniffs body and reports whether it is consistent with the declared
// media type. It is deliberately lenient: content sniffing collapses every
// text-based format (HTML, JS, CSS, JSON, XML, CSV) onto a handful of text/*
// types, so any textual declared type is accepted for textual bytes. Only a
// hard family mismatch — e.g. an "image/png" record serving HTML — fails.
func Validate(declared string, body []byte) Result {
	detected := http.DetectContentType(body)
	res := Result{Declared: declared, Detected: detected}
	dDecl := baseType(declared)
	dDet := baseType(detected)

	switch {
	case dDecl == "" || dDecl == "application/octet-stream":
		// Unspecified or opaque-binary: there is nothing meaningful to check.
		res.OK = true
		res.Reason = "declared type is unspecified; not validated"
	case dDecl == dDet:
		res.OK = true
		res.Reason = "exact media-type match"
	case isText(dDecl) && isText(dDet):
		res.OK = true
		res.Reason = "textual content consistent with declared text type"
	default:
		res.OK = false
		res.Reason = "declared " + dDecl + " but content sniffs as " + dDet
	}
	return res
}

// baseType strips parameters (charset, boundary, …) and lower-cases the media
// type, falling back to a rough split when the value is not RFC-compliant.
func baseType(ct string) string {
	if mt, _, err := mime.ParseMediaType(ct); err == nil {
		return mt
	}
	if i := strings.IndexByte(ct, ';'); i >= 0 {
		ct = ct[:i]
	}
	return strings.ToLower(strings.TrimSpace(ct))
}

// isText reports whether a base media type carries textual content. The set
// covers the text-based application/* and image/svg+xml types that
// http.DetectContentType cannot tell apart from text/plain or text/html, so we
// treat them as one equivalence class to avoid false mismatches.
func isText(base string) bool {
	if strings.HasPrefix(base, "text/") {
		return true
	}
	switch base {
	case "application/javascript", "application/x-javascript", "application/ecmascript",
		"application/json", "application/ld+json",
		"application/xml", "application/xhtml+xml", "image/svg+xml":
		return true
	}
	return false
}
