# Resolver Incentives — pay-per-query micropayment channels

The Functional Spec lists a **resolver incentive mechanism** as a nice-to-have and
explicitly defers it: *"Designing an exploit-free economic model requires substantial
time."* The hard part is the obvious naive design — paying resolvers for **claimed query
volume** — is trivially gamed: a resolver just claims it served a billion queries, and
there is no honest oracle for "work done" in a permissionless network.

This implementation sidesteps the oracle entirely.

## The idea: payment *is* the proof of work

Instead of proving volume after the fact, **clients pay for the queries they make**, as
they make them, over a unidirectional payment channel:

1. **Open.** A client opens a channel to a resolver operator with a deposit
   (`ResolverIncentives.openChannel`, `ddns channel-open`).
2. **Pay per query.** For each query, the client hands the resolver an off-chain
   **voucher** — a signature authorizing a *cumulative* amount that ticks up by the
   per-query price (`ddns voucher`). Only the newest voucher matters; it supersedes all
   earlier ones.
3. **Claim.** The resolver redeems the latest voucher on-chain whenever it likes, in a
   single transaction, collecting everything authorized so far (`claim`,
   `ddns channel-claim`). Redeeming old vouchers is a no-op.
4. **Close.** After the channel's expiry the client reclaims any unspent deposit
   (`closeChannel`, `ddns channel-close`).

A resolver that serves more clients signs-and-redeems more vouchers and earns more — the
reward tracks real, paid query volume, with no claimable counter to inflate.

## Why it isn't gameable

- **No volume oracle to forge.** There is no "I served N queries" number anywhere; the
  only thing the contract honours is a client signature over money the client agreed to
  pay. Payment is the proof.
- **Self-dealing nets zero.** A resolver spinning up a sock-puppet "client" to pay itself
  just moves its own ETH in a circle and burns gas. There is nothing to mint.
- **Replay-proof and deployment-bound.** A voucher is an EIP-191 signature over
  `(this contract, channel id, cumulative)`. It can't be replayed on another channel or a
  different deployment, and a stale voucher pays nothing because claims are cumulative.
- **Bounded counterparty risk.** The client signs voucher *N+1* only after it receives the
  answer to query *N*, so at worst it is one micro-priced query "ahead" of the resolver;
  the resolver is symmetrically never more than one query of unpaid work exposed. Neither
  side can steal the deposit — funds only move on a valid client signature, and the
  remainder is always reclaimable by the client after expiry.

It composes with discovery (HLD issue 7): a client finds a resolver via the
`ResolverRegistry`, then opens a channel to that operator address.

## What this is and isn't

This is a **working, exploit-resistant settlement primitive** — the standard
unidirectional payment channel, the same construction used by production state-channel and
micropayment systems — wired end to end through the CLI and verified on-chain (a Go-signed
voucher is redeemed by the Solidity `claim`).

It is deliberately **not** a full tokenomic economy. Out of scope, and where a production
system would go next:

- **No protocol token or inflationary block reward.** Resolvers are paid by the clients
  they serve, not by minting.
- **No staking / slashing.** Sybil resistance for *answers* already comes from the PKI
  (every response is verified against on-chain signatures and ZK proofs), so a fake
  resolver earns nothing; a deposit-and-slash layer for *availability* SLAs would be a
  separate addition, naturally hung off the `ResolverRegistry`.
- **Unidirectional and single-hop.** No bidirectional balance updates, multi-hop routing,
  or watchtowers.
- **Liveness, not fairness arbitration.** There is no on-chain dispute over "did the
  resolver actually answer" — the one-query exposure window makes it not worth arbitrating
  at micropayment sizes.

In short: it resolves the deferred nice-to-have at the right altitude for this project — a
real mechanism that pays resolvers for genuine query volume without anything to exploit —
while being honest that a complete incentive *economy* is a larger design.
