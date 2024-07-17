// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

library CCIPEnumerableMap {
    using EnumerableSet for EnumerableSet.Bytes32Set;

    struct Bytes32ToBytesMap {
        EnumerableSet.Bytes32Set _keys;
        mapping(bytes32 => bytes) _values;
    }

    function set(
        Bytes32ToBytesMap storage map,
        bytes32 key,
        bytes memory value
    ) internal returns (bool) {
        map._values[key] = value;
        return map._keys.add(key);
    }

    function remove(Bytes32ToBytesMap storage map, bytes32 key) internal returns (bool) {
        delete map._values[key];
        return map._keys.remove(key);
    }

    function contains(Bytes32ToBytesMap storage map, bytes32 key) internal view returns (bool) {
        return map._keys.contains(key);
    }

    function length(Bytes32ToBytesMap storage map) internal view returns (uint256) {
        return map._keys.length();
    }

    function at(Bytes32ToBytesMap storage map, uint256 index) internal view returns (bytes32, bytes memory) {
        bytes32 key = map._keys.at(index);
        return (key, map._values[key]);
    }

    function tryGet(Bytes32ToBytesMap storage map, bytes32 key) internal view returns (bool, bytes memory) {
        bytes memory value = map._values[key];
        if (value.length == 0) {
            return (contains(map, key), bytes(""));
        } else {
            return (true, value);
        }
    }

    function get(Bytes32ToBytesMap storage map, bytes32 key) internal view returns (bytes memory) {
        bytes memory value = map._values[key];
        require(value.length != 0 || contains(map, key), "EnumerableMap: nonexistent key");
        return value;
    }

    function get(
        Bytes32ToBytesMap storage map,
        bytes32 key,
        string memory errorMessage
    ) internal view returns (bytes memory) {
        bytes memory value = map._values[key];
        require(value.length != 0 || contains(map, key), errorMessage);
        return value;
    }
}