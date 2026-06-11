// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/// @notice Interface used by NamespaceDApp to validate records at write time.
interface IRecordSchemaRegistry {
    function typeExists(string calldata typeName) external view returns (bool);

    function mandatoryFields(
        bytes32 typeHash
    ) external view returns (string[] memory);
}

/// @title RecordSchemaRegistry — dynamic record-type schema registry.
/// @notice Lets any participant declare new record types with mandatory and
///         optional fields (HLD §3.7, UC-9; open issue 4 resolved as
///         permissionless declaration). NamespaceDApp consults this registry
///         on every setRecord to enforce mandatory fields.
contract RecordSchemaRegistry is IRecordSchemaRegistry {
    struct FieldSpec {
        string name;
        bool mandatory;
    }

    struct Schema {
        string name;
        bool exists;
        FieldSpec[] fields;
    }

    uint256 public constant MAX_FIELDS = 16;
    uint256 public constant MAX_FIELD_NAME_LENGTH = 32;
    uint256 public constant MAX_TYPE_NAME_LENGTH = 32;

    mapping(bytes32 => Schema) private _schemas;
    string[] private _typeNames;

    event TypeDeclared(
        bytes32 indexed typeHash,
        string name,
        address indexed declarer,
        uint256 mandatoryCount,
        uint256 optionalCount
    );

    error TypeAlreadyExists();
    error UnknownType();
    error InvalidTypeName();
    error InvalidFieldName();
    error TooManyFields();
    error DuplicateField();

    constructor() {
        // Built-in record types (HLD §3.7: A/AAAA-equivalents, MX-equivalents,
        // Resource References, SVC for extended port/transport/service queries).
        _declare("A", _arr("address"), new string[](0));
        _declare("AAAA", _arr("address"), new string[](0));
        _declare("MX", _arr2("host", "priority"), new string[](0));
        _declare("SVC", _arr3("target", "service", "transport"), _arr("port"));
        _declare(
            "ResourceRef",
            _arr3("infoHash", "sha256", "contentType"),
            new string[](0)
        );
    }

    // ---------------------------------------------------------------- writes

    /// @notice Declare a new record type. Permissionless; first come, first
    ///         served. Schemas are immutable once declared.
    function declareType(
        string calldata name,
        string[] calldata mandatory,
        string[] calldata optional
    ) external {
        _declare(name, mandatory, optional);
    }

    // ----------------------------------------------------------------- views

    function typeExists(string calldata typeName) external view returns (bool) {
        return _schemas[keccak256(bytes(typeName))].exists;
    }

    /// @notice Full field list (mandatory first) of a declared type.
    function getSchema(
        string calldata typeName
    ) external view returns (FieldSpec[] memory) {
        Schema storage s = _schemas[keccak256(bytes(typeName))];
        if (!s.exists) revert UnknownType();
        return s.fields;
    }

    function mandatoryFields(
        bytes32 typeHash
    ) external view returns (string[] memory) {
        Schema storage s = _schemas[typeHash];
        if (!s.exists) revert UnknownType();
        uint256 count = 0;
        for (uint256 i = 0; i < s.fields.length; i++) {
            if (s.fields[i].mandatory) count++;
        }
        string[] memory out = new string[](count);
        uint256 j = 0;
        for (uint256 i = 0; i < s.fields.length; i++) {
            if (s.fields[i].mandatory) out[j++] = s.fields[i].name;
        }
        return out;
    }

    /// @notice All declared type names (for discovery / tooling).
    function listTypes() external view returns (string[] memory) {
        return _typeNames;
    }

    // -------------------------------------------------------------- internal

    function _declare(
        string memory name,
        string[] memory mandatory,
        string[] memory optional
    ) private {
        _validateTypeName(bytes(name));
        if (mandatory.length + optional.length > MAX_FIELDS) revert TooManyFields();

        bytes32 h = keccak256(bytes(name));
        Schema storage s = _schemas[h];
        if (s.exists) revert TypeAlreadyExists();

        s.name = name;
        s.exists = true;
        for (uint256 i = 0; i < mandatory.length; i++) {
            _validateFieldName(bytes(mandatory[i]));
            s.fields.push(FieldSpec({name: mandatory[i], mandatory: true}));
        }
        for (uint256 i = 0; i < optional.length; i++) {
            _validateFieldName(bytes(optional[i]));
            s.fields.push(FieldSpec({name: optional[i], mandatory: false}));
        }
        _checkNoDuplicates(s.fields);
        _typeNames.push(name);

        emit TypeDeclared(h, name, msg.sender, mandatory.length, optional.length);
    }

    /// @dev Type names: 1-32 bytes of [A-Za-z0-9_-].
    function _validateTypeName(bytes memory b) private pure {
        if (b.length == 0 || b.length > MAX_TYPE_NAME_LENGTH) revert InvalidTypeName();
        for (uint256 i = 0; i < b.length; i++) {
            bytes1 c = b[i];
            bool ok = (c >= "a" && c <= "z") ||
                (c >= "A" && c <= "Z") ||
                (c >= "0" && c <= "9") ||
                c == "-" ||
                c == "_";
            if (!ok) revert InvalidTypeName();
        }
    }

    function _validateFieldName(bytes memory b) private pure {
        if (b.length == 0 || b.length > MAX_FIELD_NAME_LENGTH) revert InvalidFieldName();
    }

    /// @dev O(n^2) duplicate scan, bounded by MAX_FIELDS^2 = 256 iterations.
    function _checkNoDuplicates(FieldSpec[] storage fields) private view {
        for (uint256 i = 0; i < fields.length; i++) {
            bytes32 hi = keccak256(bytes(fields[i].name));
            for (uint256 j = i + 1; j < fields.length; j++) {
                if (hi == keccak256(bytes(fields[j].name))) revert DuplicateField();
            }
        }
    }

    // Constructor helpers (Solidity has no string[] literals).
    function _arr(string memory a) private pure returns (string[] memory out) {
        out = new string[](1);
        out[0] = a;
    }

    function _arr2(
        string memory a,
        string memory b
    ) private pure returns (string[] memory out) {
        out = new string[](2);
        out[0] = a;
        out[1] = b;
    }

    function _arr3(
        string memory a,
        string memory b,
        string memory c
    ) private pure returns (string[] memory out) {
        out = new string[](3);
        out[0] = a;
        out[1] = b;
        out[2] = c;
    }
}
