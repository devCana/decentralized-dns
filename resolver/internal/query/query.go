// Package query defines the normalized internal Query struct produced by
// every QueryAPI front end (REST, UDP) and the selector canonicalization
// rules shared with the owner CLI (HLD §3.2, §4.1.2).
package query

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Query is the normalized form passed to ResolverServer.HandleQuery.
type Query struct {
	Name     string // domain name, contract rules: [a-z0-9-]{1,63}, no edge hyphen
	Type     string // record type, e.g. "A", "SVC", "ResourceRef"
	Selector string // canonical selector string ("" when none)
}

var (
	nameRe          = regexp.MustCompile(`^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`)
	typeRe          = regexp.MustCompile(`^[A-Za-z0-9_-]{1,32}$`)
	selKey          = regexp.MustCompile(`^[a-z0-9_]{1,32}$`)
	selVal          = regexp.MustCompile(`^[^&=\s]{1,64}$`)
	knownTransports = map[string]bool{"UDP": true, "TCP": true, "QUIC": true}
	knownServices   = map[string]bool{"HTTP": true, "SMTP": true}
)

// ValidateName checks the domain-name grammar (mirrors the contract).
func ValidateName(name string) error {
	if !nameRe.MatchString(name) {
		return fmt.Errorf("invalid name %q: want 1-63 chars of [a-z0-9-], no edge hyphen", name)
	}
	return nil
}

// ValidateType checks the record-type grammar (mirrors the registry).
func ValidateType(t string) error {
	if !typeRe.MatchString(t) {
		return fmt.Errorf("invalid record type %q: want 1-32 chars of [A-Za-z0-9_-]", t)
	}
	return nil
}

// CanonicalSelector normalizes selector pairs into the canonical on-chain
// form: keys lowercased, known keys validated (port 1-65535, transport
// UDP/TCP/QUIC, service HTTP/SMTP with uppercased values), sorted by key
// and joined with '&' (e.g. "port=25&service=SMTP&transport=TCP").
func CanonicalSelector(pairs map[string]string) (string, error) {
	if len(pairs) == 0 {
		return "", nil
	}
	keys := make([]string, 0, len(pairs))
	canon := make(map[string]string, len(pairs))
	for k, v := range pairs {
		k = strings.ToLower(strings.TrimSpace(k))
		v = strings.TrimSpace(v)
		if !selKey.MatchString(k) {
			return "", fmt.Errorf("invalid selector key %q", k)
		}
		switch k {
		case "port":
			n, err := strconv.Atoi(v)
			if err != nil || n < 1 || n > 65535 {
				return "", fmt.Errorf("invalid port %q", v)
			}
			v = strconv.Itoa(n)
		case "transport":
			v = strings.ToUpper(v)
			if !knownTransports[v] {
				return "", fmt.Errorf("invalid transport %q: want UDP, TCP or QUIC", v)
			}
		case "service":
			v = strings.ToUpper(v)
			if !knownServices[v] {
				return "", fmt.Errorf("invalid service %q: want HTTP or SMTP", v)
			}
		default:
			if !selVal.MatchString(v) {
				return "", fmt.Errorf("invalid selector value %q for key %q", v, k)
			}
		}
		if _, dup := canon[k]; dup {
			return "", fmt.Errorf("duplicate selector key %q", k)
		}
		canon[k] = v
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, len(keys))
	for i, k := range keys {
		parts[i] = k + "=" + canon[k]
	}
	return strings.Join(parts, "&"), nil
}

// ParsePairs splits a raw "k=v&k2=v2" selector string into a pair map
// (keys lowercased, duplicates rejected). Validation happens later in
// CanonicalSelector.
func ParsePairs(raw string) (map[string]string, error) {
	pairs := map[string]string{}
	if strings.TrimSpace(raw) == "" {
		return pairs, nil
	}
	for _, part := range strings.Split(raw, "&") {
		k, v, ok := strings.Cut(part, "=")
		if !ok {
			return nil, fmt.Errorf("invalid selector fragment %q: want k=v", part)
		}
		if _, dup := pairs[strings.ToLower(k)]; dup {
			return nil, fmt.Errorf("duplicate selector key %q", k)
		}
		pairs[strings.ToLower(k)] = v
	}
	return pairs, nil
}

// ParseSelectorString splits a raw "k=v&k2=v2" selector and canonicalizes it.
func ParseSelectorString(raw string) (string, error) {
	pairs, err := ParsePairs(raw)
	if err != nil {
		return "", err
	}
	return CanonicalSelector(pairs)
}

// New validates and normalizes the parts into a Query.
func New(name, recordType string, selectorPairs map[string]string) (Query, error) {
	if err := ValidateName(name); err != nil {
		return Query{}, err
	}
	if err := ValidateType(recordType); err != nil {
		return Query{}, err
	}
	sel, err := CanonicalSelector(selectorPairs)
	if err != nil {
		return Query{}, err
	}
	return Query{Name: name, Type: recordType, Selector: sel}, nil
}
