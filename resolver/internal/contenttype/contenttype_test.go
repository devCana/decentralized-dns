package contenttype

import "testing"

func TestValidate(t *testing.T) {
	html := []byte("<!doctype html><title>ddns</title><h1>hello</h1>")
	js := []byte("function main(){ console.log('decentralized'); }\n")
	json := []byte(`{"name":"example","type":"A"}`)
	// Minimal valid PNG header so http.DetectContentType returns image/png.
	png := []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR")

	cases := []struct {
		name     string
		declared string
		body     []byte
		wantOK   bool
	}{
		{"html matches html", "text/html", html, true},
		{"html with charset param", "text/html; charset=utf-8", html, true},
		{"javascript sniffs as text", "application/javascript", js, true},
		{"text/javascript sniffs as text", "text/javascript", js, true},
		{"json sniffs as text", "application/json", json, true},
		{"png matches png", "image/png", png, true},
		{"octet-stream is unvalidated", "application/octet-stream", png, true},
		{"empty declared is unvalidated", "", html, true},
		{"png declared but body is html", "image/png", html, false},
		{"html declared but body is png", "text/html", png, false},
		{"pdf declared but body is html", "application/pdf", html, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Validate(c.declared, c.body)
			if got.OK != c.wantOK {
				t.Fatalf("Validate(%q) OK = %v (%s), want %v", c.declared, got.OK, got.Reason, c.wantOK)
			}
			if got.Detected == "" {
				t.Error("Detected should always be populated")
			}
		})
	}
}

func TestBaseType(t *testing.T) {
	cases := map[string]string{
		"text/html; charset=utf-8":  "text/html",
		"IMAGE/PNG":                 "image/png",
		"application/json":          "application/json",
		"  text/plain  ":            "text/plain",
		"not a media type ; broken": "not a media type",
	}
	for in, want := range cases {
		if got := baseType(in); got != want {
			t.Errorf("baseType(%q) = %q, want %q", in, got, want)
		}
	}
}
