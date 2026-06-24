#!/usr/bin/env bash
# Regenerates Go bindings for the contracts (HLD §3.4: abigen-generated
# bindings wrapped by the resolver's BlockchainClient).
# Usage: scripts/gen-bindings.sh   (from the repo root)
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OUT="$ROOT/resolver/internal/chain/bindings"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

(cd "$ROOT/contracts" && DO_NOT_TRACK=1 npx hardhat compile --quiet)

mkdir -p "$OUT"
for C in NamespaceDApp RecordSchemaRegistry ResolverRegistry ResolverIncentives; do
  node -e "
    const a = require('$ROOT/contracts/artifacts/contracts/$C.sol/$C.json');
    const fs = require('fs');
    fs.writeFileSync('$TMP/$C.abi', JSON.stringify(a.abi));
    fs.writeFileSync('$TMP/$C.bin', a.bytecode);
  "
  SNAKE=$(echo "$C" | sed -E 's/([a-z0-9])([A-Z])/\1_\2/g' | tr '[:upper:]' '[:lower:]')
  (cd "$ROOT/resolver" && go run github.com/ethereum/go-ethereum/cmd/abigen \
    --abi "$TMP/$C.abi" --bin "$TMP/$C.bin" \
    --pkg bindings --type "$C" --out "internal/chain/bindings/$SNAKE.go")
  echo "generated resolver/internal/chain/bindings/$SNAKE.go"
done
