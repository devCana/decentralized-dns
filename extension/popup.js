// Popup: query a resolver's /resolve endpoint and verify the resolver's
// ed25519 envelope signature *in the browser* — so the popup proves the answer
// really came from the resolver identity, without trusting the page.
//
// The signature covers the exact transmitted bytes of the `data` field, so we
// extract that raw substring (never re-serialize, which would change the bytes)
// and verify it with WebCrypto Ed25519. Full owner-signature + ZK
// re-verification is what the `ddns-lookup` CLI does; the popup focuses on
// resolver-identity verification, which is what a browser realistically needs.

const DEFAULT_RESOLVER = "http://localhost:8080"

const $ = (id) => document.getElementById(id)

async function resolverBase() {
  const { resolver } = await chrome.storage.sync.get("resolver")
  return (resolver || DEFAULT_RESOLVER).replace(/\/+$/, "")
}

// Extract the exact raw substring of a top-level JSON field's {object} value.
function extractRaw(text, field) {
  const key = `"${field}":`
  const i = text.indexOf(key)
  if (i < 0) return null
  let j = i + key.length
  const start = j
  let depth = 0, inStr = false, esc = false
  for (; j < text.length; j++) {
    const c = text[j]
    if (inStr) {
      if (esc) esc = false
      else if (c === "\\") esc = true
      else if (c === '"') inStr = false
    } else if (c === '"') inStr = true
    else if (c === "{" || c === "[") depth++
    else if (c === "}" || c === "]") { if (--depth === 0) { j++; break } }
  }
  return text.slice(start, j)
}

const hexToBytes = (h) =>
  Uint8Array.from((h.replace(/^0x/, "").match(/../g) || []).map((b) => parseInt(b, 16)))

async function verifyEnvelope(text, env) {
  const rawData = extractRaw(text, "data")
  if (!rawData) throw new Error("response is not a signed envelope")
  const pub = hexToBytes(env.resolver)
  const sig = hexToBytes(env.signature)
  if (pub.length !== 32 || sig.length !== 64) throw new Error("bad envelope key/sig length")
  const key = await crypto.subtle.importKey("raw", pub, { name: "Ed25519" }, false, ["verify"])
  return crypto.subtle.verify({ name: "Ed25519" }, key, sig, new TextEncoder().encode(rawData))
}

function render(html) { $("out").innerHTML = html }

async function run() {
  render('<span class="k">resolving…</span>')
  try {
    const base = await resolverBase()
    const params = new URLSearchParams({
      name: $("name").value.trim(),
      type: $("type").value.trim() || "A",
      selector: $("selector").value.trim(),
    })
    const resp = await fetch(`${base}/resolve?${params}`)
    const text = await resp.text()
    let env
    try { env = JSON.parse(text) } catch { throw new Error(`HTTP ${resp.status}: ${text.slice(0, 120)}`) }

    const verified = await verifyEnvelope(text, env)
    const data = env.data || {} // already parsed by JSON.parse(text) above

    const lines = []
    lines.push(
      verified
        ? `<span class="ok">✓ resolver signature verified</span>`
        : `<span class="bad">✗ resolver signature INVALID</span>`,
    )
    lines.push(`<span class="k">resolver</span> ${env.resolver.slice(0, 18)}…`)
    if (!data.found) {
      lines.push(`<span class="bad">no match</span> (${data.error || "no_match"})`)
    } else {
      const r = data.record || {}
      const fields = (r.fieldNames || []).map((n, i) => `${n}=${(r.fieldValues || [])[i]}`).join(" ")
      lines.push(`<span class="k">record</span> ${r.type} ${fields} (ttl ${r.ttl}s)`)
      lines.push(`<span class="k">owner</span> ${(data.owner || "").slice(0, 18)}…`)
      lines.push(
        `<span class="k">owner sig</span> ${data.ownerSigVerified ? "ok (per resolver)" : "unverified"}`,
      )
      lines.push(`<span class="k">cached</span> ${data.cached}`)
    }
    render(lines.join("\n"))
  } catch (e) {
    render(`<span class="bad">error:</span> ${e.message}`)
  }
}

$("go").addEventListener("click", run)
$("opts").addEventListener("click", (e) => { e.preventDefault(); chrome.runtime.openOptionsPage() })
document.addEventListener("keydown", (e) => { if (e.key === "Enter") run() })
