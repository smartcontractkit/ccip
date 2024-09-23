// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {HomeBase} from "../../capability/HomeBase.sol";

contract HomeBaseHelper is HomeBase {
  string public constant override typeAndVersion = "HomeBaseHelper 1.6.0-dev";

  uint256 public constant PREFIX = 0x0c0c << (256 - 16);

  function _validateStaticAndDynamicConfig(bytes memory, bytes memory) internal view override {}

  function _validateDynamicConfig(bytes memory, bytes memory) internal view override {}

  function _getConfigDigestPrefix() internal pure override returns (uint256) {
    return PREFIX;
  }

  function getStoredConfig(
    bytes32 pluginKey,
    bytes32 configDigest
  ) external view returns (StoredConfig memory, bool ok) {
    return _getStoredConfig(pluginKey, configDigest);
  }

  function getPrimaryStoredConfig(bytes32 pluginKey) external view returns (StoredConfig memory, bool ok) {
    return _getPrimaryStoredConfig(pluginKey);
  }

  function getSecondaryStoredConfig(bytes32 pluginKey) external view returns (StoredConfig memory, bool ok) {
    return _getSecondaryStoredConfig(pluginKey);
  }

  function calculateConfigDigest(
    bytes32 pluginKey,
    bytes memory staticConfig,
    uint32 version
  ) external view returns (bytes32) {
    return _calculateConfigDigest(pluginKey, staticConfig, version);
  }
}
