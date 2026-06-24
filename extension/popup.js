// Popup: query a resolver's /resolve endpoint and verify the resolver's
// ed25519 envelope signature *in the browser* — proving the answer was signed
// by the resolver key you pinned, not by whoever happened to answer.
//
// The signature covers the exact transmitted bytes of the `data` field, so we
// extract that raw substring (never re-serialize, which would change the bytes)
// and verify it with WebCrypto Ed25519. All record fields are rendered as inert
// text (textContent) — never innerHTML — because record values are arbitrary
// on-chain strings set by domain owners and must not be able to inject markup
// into this privileged page.

const DEFAULT_RESOLVER = "http://localhost:8080"

const $ = (id) => document.getElementById(id)

async function settings() {
  const { resolver, resolverKey } = await chrome.storage.sync.get(["resolver", "resolverKey"])
  return {
    base: (resolver || DEFAULT_RESOLVER).replace(/\/+$/, ""),
    pinnedKey: (resolverKey || "").trim().toLowerCase(),
  }
}

// Extract the exact raw substring of a top-level JSON field's {object} value.
// Returns null on anything malformed so a bad envelope fails verification
// rather than mis-extracting.
function extractRaw(text, field) {
  const key = `"${field}":`
  const i = text.indexOf(key)
  if (i < 0) return null
  let j = i + key.length
  if (text[j] !== "{" && text[j] !== "[") return null // must be an object/array
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
    else if (c === "}" || c === "]") {
      if (--depth < 0) return null
      if (depth === 0) { j++; return text.slice(start, j) }
    }
  }
  return null // never balanced
}

const hexToBytes = (h) =>
  Uint8Array.from((h.replace(/^0x/, "").match(/../g) || []).map((b) => parseInt(b, 16)))

async function verifyEnvelopeSig(text, env) {
  if (typeof env.resolver !== "string" || typeof env.signature !== "string") {
    throw new Error("response is not a signed envelope")
  }
  const rawData = extractRaw(text, "data")
  if (!rawData) throw new Error("could not extract signed payload")
  const pub = hexToBytes(env.resolver)
  const sig = hexToBytes(env.signature)
  if (pub.length !== 32 || sig.length !== 64) throw new Error("bad envelope key/sig length")
  const key = await crypto.subtle.importKey("raw", pub, { name: "Ed25519" }, false, ["verify"])
  return crypto.subtle.verify({ name: "Ed25519" }, key, sig, new TextEncoder().encode(rawData))
}

// --- inert DOM rendering (no innerHTML anywhere) ---
function clear() { $("out").textContent = "" }
function line(parts) {
  // parts: array of {text, cls?}
  const div = document.createElement("div")
  for (const p of parts) {
    const span = document.createElement("span")
    if (p.cls) span.className = p.cls
    span.textContent = p.text
    div.appendChild(span)
  }
  $("out").appendChild(div)
}

async function run() {
  clear()
  line([{ text: "resolving…", cls: "k" }])
  try {
    const { base, pinnedKey } = await settings()
    const params = new URLSearchParams({
      name: $("name").value.trim(),
      type: $("type").value.trim() || "A",
      selector: $("selector").value.trim(),
    })
    const resp = await fetch(`${base}/resolve?${params}`)
    const text = await resp.text()
    let env
    try { env = JSON.parse(text) } catch { throw new Error(`HTTP ${resp.status} (not JSON)`) }

    const sigOK = await verifyEnvelopeSig(text, env)
    const keyHex = (env.resolver || "").toLowerCase()
    const pinnedOK = pinnedKey ? keyHex === pinnedKey : null
    const data = env.data || {} // already parsed by JSON.parse(text)

    clear()
    if (!sigOK) {
      line([{ text: "✗ resolver signature INVALID", cls: "bad" }])
    } else if (pinnedOK === false) {
      line([{ text: "⚠ signed, but by an UNTRUSTED key (≠ pinned)", cls: "bad" }])
    } else if (pinnedOK === true) {
      line([{ text: "✓ verified — signature matches your pinned resolver", cls: "ok" }])
    } else {
      line([{ text: "✓ signature valid (key not pinned — set one in settings)", cls: "ok" }])
    }
    line([{ text: "resolver ", cls: "k" }, { text: `${(env.resolver || "").slice(0, 22)}…` }])

    if (!data.found) {
      line([{ text: "no match ", cls: "bad" }, { text: `(${data.error || "no_match"})` }])
    } else {
      const r = data.record || {}
      const fields = (r.fieldNames || []).map((n, i) => `${n}=${(r.fieldValues || [])[i]}`).join(" ")
      line([{ text: "record ", cls: "k" }, { text: `${r.type || "?"} ${fields} (ttl ${r.ttl}s)` }])
      line([{ text: "owner ", cls: "k" }, { text: `${(data.owner || "").slice(0, 22)}…` }])
      line([{ text: "owner sig ", cls: "k" }, { text: data.ownerSigVerified ? "ok (per resolver)" : "unverified" }])
      line([{ text: "cached ", cls: "k" }, { text: String(data.cached) }])
    }
  } catch (e) {
    clear()
    line([{ text: "error: ", cls: "bad" }, { text: e.message }])
  }
}

$("go").addEventListener("click", run)
$("opts").addEventListener("click", (e) => { e.preventDefault(); chrome.runtime.openOptionsPage() })
document.addEventListener("keydown", (e) => { if (e.key === "Enter") run() })
