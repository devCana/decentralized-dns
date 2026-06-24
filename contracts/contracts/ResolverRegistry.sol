// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @title ResolverRegistry — on-chain directory of resolvers (HLD open issue 7).
/// @notice Resolves the **resolver bootstrap** problem: with no legacy DNS to
///         lean on, a fresh client needs a way to discover resolver endpoints
///         and the ed25519 identity keys it should pin. Operators announce
///         `{ed25519 pubKey, endpoint URL}` keyed by their wallet address;
///         clients enumerate the directory, dial an endpoint, and verify the
///         resolver's response envelope against the announced key.
///
///         Registration is permissionless, so the registry is a discovery
///         *hint*, not an authority: a client still pins keys and verifies
///         every signature itself, which is exactly what neutralises a Sybil
///         flood of fake resolvers (open issue 8) — a forged resolver cannot
///         produce a valid owner signature or ZK proof for a record.
contract ResolverRegistry {
    struct Resolver {
        bytes32 pubKey; // ed25519 public key (32 bytes)
        string endpoint; // base URL, e.g. https://resolver.example:8080
        uint64 updatedAt; // block timestamp of the last announce
        bool active;
    }

    uint256 public constant MAX_ENDPOINT_LENGTH = 256;

    mapping(address => Resolver) private _resolvers;
    address[] private _operators; // enumeration; may include revoked entries

    event ResolverAnnounced(
        address indexed operator,
        bytes32 pubKey,
        string endpoint
    );
    event ResolverRevoked(address indexed operator);

    error EmptyPubKey();
    error EmptyEndpoint();
    error EndpointTooLong();
    error NotRegistered();

    /// @notice Announce (or update) the calling operator's resolver.
    function announce(bytes32 pubKey, string calldata endpoint) external {
        if (pubKey == bytes32(0)) revert EmptyPubKey();
        bytes memory e = bytes(endpoint);
        if (e.length == 0) revert EmptyEndpoint();
        if (e.length > MAX_ENDPOINT_LENGTH) revert EndpointTooLong();

        Resolver storage r = _resolvers[msg.sender];
        if (r.updatedAt == 0) {
            _operators.push(msg.sender); // first-ever announce → enumerate
        }
        r.pubKey = pubKey;
        r.endpoint = endpoint;
        r.updatedAt = uint64(block.timestamp);
        r.active = true;
        emit ResolverAnnounced(msg.sender, pubKey, endpoint);
    }

    /// @notice Revoke the calling operator's resolver (marks it inactive).
    function revoke() external {
        Resolver storage r = _resolvers[msg.sender];
        if (!r.active) revert NotRegistered();
        r.active = false;
        emit ResolverRevoked(msg.sender);
    }

    /// @notice Look up a single operator's resolver entry.
    function getResolver(
        address operator
    )
        external
        view
        returns (
            bytes32 pubKey,
            string memory endpoint,
            uint64 updatedAt,
            bool active
        )
    {
        Resolver storage r = _resolvers[operator];
        return (r.pubKey, r.endpoint, r.updatedAt, r.active);
    }

    /// @notice Count of operators ever registered (active or revoked).
    function operatorCount() external view returns (uint256) {
        return _operators.length;
    }

    /// @notice Operator address at an enumeration index.
    function operatorAt(uint256 i) external view returns (address) {
        return _operators[i];
    }

    /// @notice Every currently-active resolver, for client discovery.
    function activeResolvers()
        external
        view
        returns (
            address[] memory operators,
            bytes32[] memory pubKeys,
            string[] memory endpoints
        )
    {
        uint256 n;
        for (uint256 i = 0; i < _operators.length; i++) {
            if (_resolvers[_operators[i]].active) n++;
        }
        operators = new address[](n);
        pubKeys = new bytes32[](n);
        endpoints = new string[](n);
        uint256 j;
        for (uint256 i = 0; i < _operators.length; i++) {
            address op = _operators[i];
            Resolver storage r = _resolvers[op];
            if (r.active) {
                operators[j] = op;
                pubKeys[j] = r.pubKey;
                endpoints[j] = r.endpoint;
                j++;
            }
        }
    }
}
