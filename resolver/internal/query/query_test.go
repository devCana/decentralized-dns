package query

import "testing"

func TestValidateName(t *testing.T) {
	for _, good := range []string{"a", "example", "a-b", "x0", "a1-b2-c3"} {
		if err := ValidateName(good); err != nil {
			t.Errorf("ValidateName(%q) = %v", good, err)
		}
	}
	for _, bad := range []string{"", "UPPER", "-lead", "trail-", "und_er", "dot.ted", "spa ce", string(make([]byte, 64))} {
		if err := ValidateName(bad); err == nil {
			t.Errorf("ValidateName(%q) should fail", bad)
		}
	}
}

func TestCanonicalSelector(t *testing.T) {
	got, err := CanonicalSelector(map[string]string{
		"transport": "tcp", "PORT": "25", "service": "smtp",
	})
	if err != nil {
		t.Fatal(err)
	}
	want := "port=25&service=SMTP&transport=TCP"
	if got != want {
		t.Errorf("canonical = %q, want %q", got, want)
	}

	if got, _ := CanonicalSelector(nil); got != "" {
		t.Errorf("empty selector = %q", got)
	}

	// arbitrary keys for dynamic types pass through, lowercased + sorted
	got, err = CanonicalSelector(map[string]string{"zone": "eu", "lang": "en"})
	if err != nil {
		t.Fatal(err)
	}
	if got != "lang=en&zone=eu" {
		t.Errorf("dynamic canonical = %q", got)
	}
}

func TestCanonicalSelectorRejections(t *testing.T) {
	cases := []map[string]string{
		{"port": "0"},
		{"port": "65536"},
		{"port": "abc"},
		{"transport": "SCTP"},
		{"service": "FTP"},
		{"bad key": "v"},
		{"k": "a&b"},
		{"k": ""},
	}
	for _, pairs := range cases {
		if _, err := CanonicalSelector(pairs); err == nil {
			t.Errorf("CanonicalSelector(%v) should fail", pairs)
		}
	}
}

func TestParseSelectorString(t *testing.T) {
	got, err := ParseSelectorString("service=HTTP&port=443&transport=quic")
	if err != nil {
		t.Fatal(err)
	}
	if got != "port=443&service=HTTP&transport=QUIC" {
		t.Errorf("parsed = %q", got)
	}
	if _, err := ParseSelectorString("noequals"); err == nil {
		t.Error("expected error for fragment without '='")
	}
	if _, err := ParseSelectorString("port=1&port=2"); err == nil {
		t.Error("expected error for duplicate key")
	}
	if got, _ := ParseSelectorString("  "); got != "" {
		t.Errorf("blank = %q", got)
	}
}

func TestNew(t *testing.T) {
	q, err := New("example", "SVC", map[string]string{"service": "HTTP"})
	if err != nil {
		t.Fatal(err)
	}
	if q.Selector != "service=HTTP" {
		t.Errorf("selector = %q", q.Selector)
	}
	if _, err := New("BAD", "A", nil); err == nil {
		t.Error("expected name validation error")
	}
	if _, err := New("ok", "bad type!", nil); err == nil {
		t.Error("expected type validation error")
	}
}
