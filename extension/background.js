// Omnibox integration: typing `ddns <name>` in the address bar opens the
// resolver's /web gateway for that name — the verified decentralized site,
// rendered in a standard browser with no native protocol handler needed.

const DEFAULT_RESOLVER = "http://localhost:8080"

async function resolverBase() {
  const { resolver } = await chrome.storage.sync.get("resolver")
  return (resolver || DEFAULT_RESOLVER).replace(/\/+$/, "")
}

function escapeXml(s) {
  return s.replace(/[<>&'"]/g, (c) =>
    ({ "<": "&lt;", ">": "&gt;", "&": "&amp;", "'": "&apos;", '"': "&quot;" }[c]),
  )
}

chrome.omnibox.setDefaultSuggestion({
  description: "Resolve a decentralized name and open its verified site",
})

chrome.omnibox.onInputChanged.addListener(async (text, suggest) => {
  const name = text.trim()
  if (!name) return
  const base = await resolverBase()
  suggest([
    {
      content: name,
      description: `Open <url>${escapeXml(base)}/web/${escapeXml(name)}</url> — verified decentralized site`,
    },
  ])
})

chrome.omnibox.onInputEntered.addListener(async (text, disposition) => {
  const name = encodeURIComponent(text.trim())
  if (!name) return
  const url = `${await resolverBase()}/web/${name}`
  if (disposition === "currentTab") {
    chrome.tabs.update({ url })
  } else {
    chrome.tabs.create({ url, active: disposition === "newForegroundTab" })
  }
})
