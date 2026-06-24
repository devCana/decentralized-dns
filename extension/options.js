const DEFAULT_RESOLVER = "http://localhost:8080"
const resolver = document.getElementById("resolver")
const resolverKey = document.getElementById("resolverKey")
const saved = document.getElementById("saved")

chrome.storage.sync.get(["resolver", "resolverKey"]).then((s) => {
  resolver.value = s.resolver || DEFAULT_RESOLVER
  resolverKey.value = s.resolverKey || ""
})

document.getElementById("save").addEventListener("click", async () => {
  const key = resolverKey.value.trim()
  if (key && !/^0x[0-9a-fA-F]{64}$/.test(key)) {
    saved.textContent = "key must be 0x + 64 hex"
    return
  }
  await chrome.storage.sync.set({
    resolver: resolver.value.trim() || DEFAULT_RESOLVER,
    resolverKey: key,
  })
  saved.textContent = "saved ✓"
  setTimeout(() => (saved.textContent = ""), 1500)
})
