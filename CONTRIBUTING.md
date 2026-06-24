# Contributing

Thanks for taking a look. This is a research/showcase project, but it's built to
real engineering standards and contributions are welcome.

## Layout

| Path         | What lives there                                             |
| ------------ | ------------------------------------------------------------ |
| `contracts/` | Solidity dApp + Hardhat tests (`npx hardhat test`)           |
| `resolver/`  | Go resolver server, CLIs, and unit tests                     |
| `docs/`      | Functional spec and high-level design (with Mermaid diagrams)|
| `scripts/`   | `demo.sh` end-to-end walkthrough                             |

## Prerequisites

- Go 1.25+
- Node 22+

## Development loop

Everything is wired through the `Makefile`:

```bash
make build          # compile contracts + Go binaries
make test           # hardhat tests + go test
make race           # go test under the race detector (CI runs this)
make cover          # resolver coverage report
make fmt            # gofmt -w (CI fails on unformatted code)
make demo           # full local end-to-end demo (chain + resolver + CLIs)
```

## Before opening a PR

CI runs on every push and pull request. To match it locally:

1. `make fmt` — the CI `gofmt` gate fails on any unformatted file.
2. `cd resolver && go vet ./...` — must be clean.
3. `make race` — the full Go suite must pass under `-race`.
4. `make contracts-test` — all Hardhat tests must pass.

New behavior should come with tests. The cryptographic core (`resolver/internal/pki`,
`resolver/internal/zk`) is held to negative-path coverage too: a change to a signed
message format should add a tamper case proving the old signature is rejected.

## Commit style

Conventional-commit prefixes (`feat:`, `fix:`, `docs:`, `chore:`) with a short,
imperative subject. Keep commits logically grouped.
