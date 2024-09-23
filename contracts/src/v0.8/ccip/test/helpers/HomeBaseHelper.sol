// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {HomeBase} from "../../capability/HomeBase.sol";

contract HomeBaseHelper is HomeBase {
  string public constant override typeAndVersion = "HomeBaseHelper 1.6.0-dev";

  uint256 public constant PREFIX = 0x0c0c << (256 - 16);

  constructor(address capabilitiesRegistry) HomeBase(capabilitiesRegistry) {}

  function _validateStaticAndDynamicConfig(bytes memory, bytes memory) internal view override {}

  function _validateDynamicConfig(bytes memory, bytes memory) internal view override {}

  function _getConfigDigestPrefix() internal pure override returns (uint256) {
    return PREFIX;
  }

  function getStoredConfig(
    uint32 donId,
    uint8 pluginType,
    bytes32 configDigest
  ) external view returns (StoredConfig memory, bool ok) {
    return _getStoredConfig(donId, pluginType, configDigest);
  }

  function getPrimaryStoredConfig(uint32 donId, uint8 pluginType) external view returns (StoredConfig memory, bool ok) {
    return _getPrimaryStoredConfig(donId, pluginType);
  }

  function getSecondaryStoredConfig(
    uint32 donId,
    uint8 pluginType
  ) external view returns (StoredConfig memory, bool ok) {
    return _getSecondaryStoredConfig(donId, pluginType);
  }

  function calculateConfigDigest(
    uint32 donId,
    uint8 pluginType,
    bytes memory staticConfig,
    uint32 version
  ) external view returns (bytes32) {
    return _calculateConfigDigest(donId, pluginType, staticConfig, version);
  }
}
