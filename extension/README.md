# ddns browser extension

A small Manifest V3 extension that brings decentralized DNS into a standard browser —
the deferred "native browser integration" nice-to-have, in the form that actually works
without OS-level protocol hooks.

It does two things:

1. **Omnibox resolution.** Type `ddns` + space + a name in the address bar (e.g.
   `ddns example`) to open `‹resolver›/web/‹name›` — the resolver's [`/web` gateway](../README.md#resolver-rest-api),
   which resolves the domain's HTTP `ResourceRef`, verifies it on-chain (owner signature +
   SHA-256 + ZK), and renders it.
2. **In-browser verification.** The toolbar popup queries `‹resolver›/resolve` and verifies
   the resolver's **ed25519 envelope signature in the browser** with WebCrypto — proving
   the answer really came from that resolver identity, not the page. It extracts the exact
   signed bytes of the response (rather than re-serializing, which would change them) and
   verifies them directly.

## Install (unpacked)

1. Open `chrome://extensions` (Chrome/Edge/Brave).
2. Enable **Developer mode**.
3. **Load unpacked** → select this `extension/` directory.
4. Click the extension's **Options** and set your resolver URL (default
   `http://localhost:8080`). Discover public resolvers on-chain with `ddns resolvers`.

Try it: run `make demo` (or any local resolver), then type `ddns example` in the address
bar, or open the popup and click **Resolve & verify**.

## Notes & scope

- Requires a Chromium-based browser with WebCrypto **Ed25519** support (Chrome/Edge 137+).
- The `/web` gateway trusts the resolver to do the on-chain verification server-side; the
  popup adds a client-side check of the resolver's *identity signature*. Full owner-signature
  and ZK re-verification (as the `ddns-lookup` CLI does) is heavier and left to the CLI.
- A true `ddns://` URL scheme would need a native protocol handler / OS registration, which
  remains out of scope; this extension delivers the same outcome through the omnibox.
