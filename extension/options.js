const DEFAULT_RESOLVER = "http://localhost:8080"
const input = document.getElementById("resolver")
const saved = document.getElementById("saved")

chrome.storage.sync.get("resolver").then(({ resolver }) => {
  input.value = resolver || DEFAULT_RESOLVER
})

document.getElementById("save").addEventListener("click", async () => {
  const resolver = input.value.trim() || DEFAULT_RESOLVER
  await chrome.storage.sync.set({ resolver })
  saved.textContent = "saved ✓"
  setTimeout(() => (saved.textContent = ""), 1500)
})
