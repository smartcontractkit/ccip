// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {Types} from "../../capability/libraries/Types.sol";
import {CCIPConfig} from "../../capability/CCIPConfig.sol";

contract CCIPConfigHelper is CCIPConfig {
  constructor(address capabilitiesRegistry) CCIPConfig(capabilitiesRegistry) {}

  function stateFromConfigLength(uint256 configLength) public pure returns (Types.ConfigState) {
    return _stateFromConfigLength(configLength);
  }

  function validateConfigStateTransition(Types.ConfigState currentState, Types.ConfigState newState) public pure {
    _validateConfigStateTransition(currentState, newState);
  }

  function validateConfigTransition(
    Types.OCR3ConfigWithMeta[] memory currentConfig,
    Types.OCR3ConfigWithMeta[] memory newConfigWithMeta
  ) public pure {
    _validateConfigTransition(currentConfig, newConfigWithMeta);
  }

  function computeNewConfigWithMeta(
    uint32 donId,
    Types.OCR3ConfigWithMeta[] memory currentConfig,
    Types.OCR3Config[] memory newConfig,
    Types.ConfigState currentState,
    Types.ConfigState newState
  ) public view returns (Types.OCR3ConfigWithMeta[] memory) {
    return _computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, newState);
  }

  function groupByPluginType(Types.OCR3Config[] memory ocr3Configs)
    public
    pure
    returns (Types.OCR3Config[] memory commitConfigs, Types.OCR3Config[] memory execConfigs)
  {
    return _groupByPluginType(ocr3Configs);
  }

  function computeConfigDigest(
    uint32 donId,
    uint64 configCount,
    Types.OCR3Config memory ocr3Config
  ) public pure returns (bytes32) {
    return _computeConfigDigest(donId, configCount, ocr3Config);
  }

  function validateConfig(Types.OCR3Config memory cfg) public view {
    _validateConfig(cfg);
  }

  function updatePluginConfig(uint32 donId, Internal.OCRPluginType pluginType, Types.OCR3Config[] memory newConfig) public {
    _updatePluginConfig(donId, pluginType, newConfig);
  }
}
