// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title ResolverIncentives — pay-per-query micropayment channels (FS §2.3).
/// @notice Rewards resolvers by query volume *without* a gameable "proof of
///         volume". The Functional Spec deferred this feature because an
///         exploit-free volume oracle is hard — so we sidestep the oracle
///         entirely: clients pay for the queries they actually make.
///
///         A client opens a unidirectional channel to a resolver operator with
///         a deposit, then signs off-chain vouchers authorizing a
///         monotonically increasing cumulative amount — one bumped voucher per
///         query, exchanged for the answer. The resolver redeems the latest
///         voucher on-chain to collect what it has earned; only the newest is
///         ever needed. Unspent funds return to the client after a challenge
///         window.
///
///         Why this is not gameable: payment *is* the proof of work, so there
///         is nothing to forge. A resolver inflating volume by paying its own
///         sock-puppet client just moves its own money in a circle and nets
///         zero (minus gas). See docs/incentives.md for the full model and its
///         boundaries (it is a settlement primitive, not a complete tokenomic
///         system).
contract ResolverIncentives {
    struct Channel {
        address client;
        address resolver; // operator address (as advertised in ResolverRegistry)
        uint256 deposit;
        uint256 claimed; // cumulative wei already paid out
        uint64 expiresAt; // client may reclaim the remainder after this
    }

    uint64 public constant MIN_DURATION = 1 hours;

    mapping(bytes32 => Channel) public channels;
    uint256 private _nonce;

    event ChannelOpened(
        bytes32 indexed id,
        address indexed client,
        address indexed resolver,
        uint256 deposit,
        uint64 expiresAt
    );
    event Claimed(bytes32 indexed id, uint256 cumulative, uint256 paid);
    event ChannelClosed(bytes32 indexed id, uint256 refunded);

    error ZeroResolver();
    error NoDeposit();
    error NoChannel();
    error NotResolver();
    error NotClient();
    error BadVoucher();
    error NothingToClaim();
    error NotExpired();
    error TransferFailed();

    /// @notice Open a channel funding `resolver` for at least `duration`.
    function openChannel(
        address resolver,
        uint64 duration
    ) external payable returns (bytes32 id) {
        if (resolver == address(0)) revert ZeroResolver();
        if (msg.value == 0) revert NoDeposit();
        if (duration < MIN_DURATION) duration = MIN_DURATION;

        id = keccak256(abi.encode(msg.sender, resolver, _nonce++));
        uint64 expiresAt = uint64(block.timestamp) + duration;
        channels[id] = Channel(msg.sender, resolver, msg.value, 0, expiresAt);
        emit ChannelOpened(id, msg.sender, resolver, msg.value, expiresAt);
    }

    /// @notice Resolver redeems a client-signed voucher for `cumulative` wei.
    ///         Idempotent in the cumulative sense: only the delta over what was
    ///         already claimed is paid, so replaying an old voucher is a no-op.
    function claim(
        bytes32 id,
        uint256 cumulative,
        bytes calldata clientSig
    ) external {
        Channel storage ch = channels[id];
        if (ch.client == address(0)) revert NoChannel();
        if (msg.sender != ch.resolver) revert NotResolver();
        // Verify against the value the client actually signed, *then* cap the
        // payout at the deposit (capping first would change the digest).
        if (_recoverVoucher(id, cumulative, clientSig) != ch.client) {
            revert BadVoucher();
        }
        uint256 effective = cumulative > ch.deposit ? ch.deposit : cumulative;
        if (effective <= ch.claimed) revert NothingToClaim();
        uint256 pay = effective - ch.claimed;
        ch.claimed = effective;
        (bool ok, ) = payable(ch.resolver).call{value: pay}("");
        if (!ok) revert TransferFailed();
        emit Claimed(id, effective, pay);
    }

    /// @notice Client reclaims the unspent remainder once the channel expires.
    function closeChannel(bytes32 id) external {
        Channel storage ch = channels[id];
        if (ch.client == address(0)) revert NoChannel();
        if (msg.sender != ch.client) revert NotClient();
        if (block.timestamp < ch.expiresAt) revert NotExpired();

        uint256 refund = ch.deposit - ch.claimed;
        delete channels[id];
        if (refund > 0) {
            (bool ok, ) = payable(msg.sender).call{value: refund}("");
            if (!ok) revert TransferFailed();
        }
        emit ChannelClosed(id, refund);
    }

    /// @notice The EIP-191 digest a client signs to authorize `cumulative` wei
    ///         on channel `id`. Bound to this contract so vouchers cannot be
    ///         replayed against a different deployment.
    function voucherDigest(
        bytes32 id,
        uint256 cumulative
    ) public view returns (bytes32) {
        bytes32 inner = keccak256(abi.encode(address(this), id, cumulative));
        return
            keccak256(
                abi.encodePacked("\x19Ethereum Signed Message:\n32", inner)
            );
    }

    function _recoverVoucher(
        bytes32 id,
        uint256 cumulative,
        bytes calldata sig
    ) internal view returns (address) {
        if (sig.length != 65) revert BadVoucher();
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(sig.offset)
            s := calldataload(add(sig.offset, 32))
            v := byte(0, calldataload(add(sig.offset, 64)))
        }
        if (v < 27) v += 27;
        return ecrecover(voucherDigest(id, cumulative), v, r, s);
    }
}
