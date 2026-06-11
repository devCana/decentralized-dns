// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title NamespaceDApp — on-chain authority for the Decentralized DNS.
/// @notice Manages namespace registration, renewal, ownership transfer and
///         fee collection (HLD §3.7, UC-1/UC-3). Record storage is layered
///         on top by the record-store extension (issue #3).
/// @dev Domains are keyed by keccak256(name). Storage is intentionally
///      minimal: blockchain storage is expensive (HLD §2.2) — only
///      namespace metadata lives here.
contract NamespaceDApp {
    struct Domain {
        address owner; // current owner; address(0) == never registered
        bytes pubKey; // owner's public key, verified off-chain by resolvers
        uint64 expiry; // unix timestamp after which the name is re-registerable
        uint64 generation; // bumped on every (re-)registration; namespaces records
    }

    /// @notice Registration / renewal period.
    uint256 public constant REGISTRATION_PERIOD = 365 days;

    /// @notice Base yearly fee in wei for names of 5+ characters.
    ///         Shorter names cost more: 1 char x16, 2 x8, 3 x4, 4 x2.
    uint256 public immutable basePrice;

    /// @notice Account allowed to withdraw collected fees (the deployer).
    address public immutable treasurer;

    mapping(bytes32 => Domain) internal _domains;

    event Registered(
        bytes32 indexed nameHash,
        string name,
        address indexed owner,
        bytes pubKey,
        uint64 expiry,
        uint64 generation,
        uint256 feePaid
    );
    event Renewed(bytes32 indexed nameHash, uint64 newExpiry, uint256 feePaid);
    event Transferred(
        bytes32 indexed nameHash,
        address indexed oldOwner,
        address indexed newOwner,
        bytes newPubKey
    );
    event Withdrawn(address indexed to, uint256 amount);

    error InvalidName();
    error InvalidPubKey();
    error NameUnavailable();
    error DomainNotRegistered();
    error DomainExpired();
    error NotDomainOwner();
    error InsufficientFee(uint256 required, uint256 provided);
    error NotTreasurer();
    error ZeroAddress();
    error RefundFailed();
    error WithdrawFailed();

    constructor(uint256 _basePrice) {
        basePrice = _basePrice;
        treasurer = msg.sender;
    }

    // ---------------------------------------------------------------- views

    /// @notice Yearly fee for `name` (length-based pricing).
    function priceOf(string calldata name) public view returns (uint256) {
        uint256 len = bytes(name).length;
        if (len == 0 || len > 63) revert InvalidName();
        if (len == 1) return basePrice * 16;
        if (len == 2) return basePrice * 8;
        if (len == 3) return basePrice * 4;
        if (len == 4) return basePrice * 2;
        return basePrice;
    }

    /// @notice True when `name` can be registered (never owned or expired).
    function available(string calldata name) public view returns (bool) {
        Domain storage d = _domains[keccak256(bytes(name))];
        return d.owner == address(0) || block.timestamp > d.expiry;
    }

    /// @notice Owner of an active domain; address(0) if free or expired.
    function ownerOf(string calldata name) external view returns (address) {
        Domain storage d = _domains[keccak256(bytes(name))];
        if (d.owner == address(0) || block.timestamp > d.expiry) {
            return address(0);
        }
        return d.owner;
    }

    /// @notice Raw domain state (also for expired domains; check `expiry`).
    function getDomain(
        string calldata name
    )
        external
        view
        returns (address owner, bytes memory pubKey, uint64 expiry, uint64 generation)
    {
        Domain storage d = _domains[keccak256(bytes(name))];
        return (d.owner, d.pubKey, d.expiry, d.generation);
    }

    // ------------------------------------------------------------- mutations

    /// @notice Register (or re-register an expired) `name`, paying the fee.
    /// @param name   the domain name: 1-63 chars of [a-z0-9-], no edge hyphen
    /// @param pubKey the owner's public key used by resolvers/clients to
    ///               verify record signatures (kept off-chain by the owner)
    function register(string calldata name, bytes calldata pubKey) external payable {
        _validateName(bytes(name));
        if (pubKey.length == 0 || pubKey.length > 128) revert InvalidPubKey();

        bytes32 h = keccak256(bytes(name));
        Domain storage d = _domains[h];
        if (d.owner != address(0) && block.timestamp <= d.expiry) {
            revert NameUnavailable();
        }

        uint256 price = priceOf(name);
        if (msg.value < price) revert InsufficientFee(price, msg.value);

        d.owner = msg.sender;
        d.pubKey = pubKey;
        d.expiry = uint64(block.timestamp + REGISTRATION_PERIOD);
        // Generation bump logically invalidates any records left behind by a
        // previous owner of the expired name (used by the record store).
        unchecked {
            d.generation += 1;
        }

        emit Registered(h, name, msg.sender, pubKey, d.expiry, d.generation, price);
        _refundExcess(price);
    }

    /// @notice Extend an active domain by one period. Anyone may pay.
    function renew(string calldata name) external payable {
        bytes32 h = keccak256(bytes(name));
        Domain storage d = _domains[h];
        if (d.owner == address(0)) revert DomainNotRegistered();
        if (block.timestamp > d.expiry) revert DomainExpired();

        uint256 price = priceOf(name);
        if (msg.value < price) revert InsufficientFee(price, msg.value);

        d.expiry = uint64(uint256(d.expiry) + REGISTRATION_PERIOD);

        emit Renewed(h, d.expiry, price);
        _refundExcess(price);
    }

    /// @notice Transfer an active domain, atomically rewriting owner and
    ///         pubKey (UC-3). Historical records remain on-chain.
    function transfer(
        string calldata name,
        address newOwner,
        bytes calldata newPubKey
    ) external {
        if (newOwner == address(0)) revert ZeroAddress();
        if (newPubKey.length == 0 || newPubKey.length > 128) revert InvalidPubKey();

        bytes32 h = keccak256(bytes(name));
        Domain storage d = _domains[h];
        if (d.owner == address(0)) revert DomainNotRegistered();
        if (block.timestamp > d.expiry) revert DomainExpired();
        if (msg.sender != d.owner) revert NotDomainOwner();

        address oldOwner = d.owner;
        d.owner = newOwner;
        d.pubKey = newPubKey;

        emit Transferred(h, oldOwner, newOwner, newPubKey);
    }

    /// @notice Withdraw all collected fees to `to`. Treasurer only.
    function withdraw(address payable to) external {
        if (msg.sender != treasurer) revert NotTreasurer();
        if (to == address(0)) revert ZeroAddress();
        uint256 amount = address(this).balance;
        emit Withdrawn(to, amount);
        (bool ok, ) = to.call{value: amount}("");
        if (!ok) revert WithdrawFailed();
    }

    // -------------------------------------------------------------- internal

    /// @dev Names: 1-63 bytes, [a-z0-9-], hyphen not at either edge.
    ///      The loop is bounded by 63 iterations (HLD determinism constraint).
    function _validateName(bytes memory b) internal pure {
        uint256 len = b.length;
        if (len == 0 || len > 63) revert InvalidName();
        if (b[0] == "-" || b[len - 1] == "-") revert InvalidName();
        for (uint256 i = 0; i < len; i++) {
            bytes1 c = b[i];
            bool ok = (c >= "a" && c <= "z") || (c >= "0" && c <= "9") || c == "-";
            if (!ok) revert InvalidName();
        }
    }

    /// @dev Refund any overpayment, after all state changes (CEI pattern).
    function _refundExcess(uint256 price) internal {
        uint256 excess = msg.value - price;
        if (excess > 0) {
            (bool ok, ) = payable(msg.sender).call{value: excess}("");
            if (!ok) revert RefundFailed();
        }
    }

    /// @dev Internal accessor for derived contracts (record store).
    function _domain(bytes32 nameHash) internal view returns (Domain storage) {
        return _domains[nameHash];
    }
}
