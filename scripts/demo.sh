#!/usr/bin/env bash
#
# End-to-end demo (HLD §4.4): boots a local Hardhat chain, deploys the
# contracts, seeds a sample domain, starts the Go resolver, then exercises the
# owner CLI and the verifying client CLIs against it — including a
# BitTorrent-backed ResourceRef publish + fetch with on-chain hash verification.
#
# Everything runs locally and is torn down on exit. Requires: node 22+, go 1.25+.
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CONTRACTS="$ROOT/contracts"
RESOLVER="$ROOT/resolver"
WORK="$(mktemp -d)"
# Hardhat account #1 (alice) — owns the demo domain "example".
ALICE_PK="0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
RPC_URL="http://127.0.0.1:8545"
REST="http://127.0.0.1:8080"
PUBLISH_PORT=42100

PIDS=()
cleanup() {
  echo
  echo "== cleaning up =="
  for pid in "${PIDS[@]:-}"; do kill "$pid" 2>/dev/null || true; done
  wait 2>/dev/null || true
  rm -rf "$WORK"
}
trap cleanup EXIT

step() { printf '\n\033[1;36m== %s ==\033[0m\n' "$1"; }

wait_for() { # url, name
  for _ in $(seq 1 60); do
    if curl -fsS "$1" >/dev/null 2>&1; then return 0; fi
    sleep 0.5
  done
  echo "timed out waiting for $2 ($1)" >&2
  exit 1
}

step "1/7  install + compile contracts"
( cd "$CONTRACTS" && npm install --silent && npx hardhat compile )

step "2/7  start local Hardhat chain"
( cd "$CONTRACTS" && npx hardhat node --hostname 127.0.0.1 >"$WORK/chain.log" 2>&1 ) &
PIDS+=($!)
wait_for "$RPC_URL" "hardhat node" || true
# eth_chainId probe (curl on a JSON-RPC POST endpoint)
for _ in $(seq 1 60); do
  if curl -fsS -X POST "$RPC_URL" -H 'content-type: application/json' \
      --data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId"}' >/dev/null 2>&1; then break; fi
  sleep 0.5
done

step "3/7  deploy + seed contracts"
( cd "$CONTRACTS" && npx hardhat run scripts/deploy.ts --network localhost )
( cd "$CONTRACTS" && npx hardhat run scripts/seed.ts --network localhost )

NAMESPACE="$(node -pe 'require("'"$CONTRACTS"'/deployments/localhost.json").contracts.NamespaceDApp')"
REGISTRY="$(node -pe 'require("'"$CONTRACTS"'/deployments/localhost.json").contracts.RecordSchemaRegistry')"
echo "NamespaceDApp=$NAMESPACE"
echo "RecordSchemaRegistry=$REGISTRY"

step "4/7  build Go binaries"
( cd "$RESOLVER" && go build -o "$WORK/bin/" ./cmd/resolver ./cmd/ddns ./cmd/ddns-lookup ./cmd/ddns-fetch )
export PATH="$WORK/bin:$PATH"

step "5/7  start the resolver"
RPC_URL="$RPC_URL" CONTRACT_ADDRESS="$NAMESPACE" REGISTRY_ADDRESS="$REGISTRY" \
  RESOLVER_KEYSTORE="$WORK/resolver.key" DATA_DIR="$WORK/data" ALLOW_PEER_HINTS=true \
  "$WORK/bin/resolver" >"$WORK/resolver.log" 2>&1 &
PIDS+=($!)
wait_for "$REST/healthz" "resolver"
echo "resolver healthy: $(curl -fsS "$REST/healthz")"

export DDNS_PRIVATE_KEY="$ALICE_PK"
export DDNS_DEPLOYMENTS="$CONTRACTS/deployments/localhost.json"
export DDNS_RESOLVER="$REST"

step "6/7  resolve + verify seeded records (ddns-lookup)"
ddns-lookup example A
echo
ddns-lookup example SVC --selector "service=SMTP&transport=TCP&port=25"

step "6b   owner CLI writes a new record (ddns set) and we resolve it"
ddns set example MX --ttl 300 host=mail.example priority=10
echo
ddns-lookup example MX

step "7/7  publish a static file + fetch it verified over BitTorrent"
echo '<!doctype html><title>Decentralized DNS</title><h1>served from BitTorrent, verified on-chain</h1>' >"$WORK/site.html"
ddns publish-resource example "$WORK/site.html" --selector "service=HTTP" \
  --bt-port "$PUBLISH_PORT" --data-dir "$WORK/seed" --seconds 30 >"$WORK/publish.log" 2>&1 &
PIDS+=($!)
sleep 5
echo "fetching resource through the resolver (verifying SHA-256 + provenance)..."
if ddns-fetch example --selector "service=HTTP" --peer "127.0.0.1:$PUBLISH_PORT" \
    --timeout 60s -o "$WORK/fetched.html"; then
  echo "--- fetched, verified file ---"
  cat "$WORK/fetched.html"
else
  echo "(resource fetch did not complete in time — see $WORK/publish.log; the on-chain anchor + record resolution above are the core demo)"
fi

step "demo complete"
echo "All responses above were verified client-side: resolver identity signature,"
echo "owner record signature, ZK commitment proof, and (for the file) SHA-256."
