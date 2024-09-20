// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPHome} from "../../capability/CCIPHome.sol";
import {Internal} from "../../libraries/Internal.sol";

contract CCIPHomeHelper is CCIPHome {
  constructor(address capabilitiesRegistry) CCIPHome(capabilitiesRegistry) {}

  function stateFromConfigLength(uint256 configLength) public pure returns (CCIPHome.ConfigState) {
    return _stateFromConfigLength(configLength);
  }

  function validateConfigStateTransition(CCIPHome.ConfigState currentState, CCIPHome.ConfigState newState) public pure {
    _validateConfigStateTransition(currentState, newState);
  }

  function validateConfigTransition(
    CCIPHome.OCR3ConfigWithMeta[] memory currentConfig,
    CCIPHome.OCR3ConfigWithMeta[] memory newConfigWithMeta
  ) public pure {
    _validateConfigTransition(currentConfig, newConfigWithMeta);
  }

  function computeNewConfigWithMeta(
    uint32 donId,
    CCIPHome.OCR3ConfigWithMeta[] memory currentConfig,
    CCIPHome.OCR3Config[] memory newConfig,
    CCIPHome.ConfigState currentState,
    CCIPHome.ConfigState newState
  ) public view returns (CCIPHome.OCR3ConfigWithMeta[] memory) {
    return _computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, newState);
  }

  function groupByPluginType(
    CCIPHome.OCR3Config[] memory ocr3Configs
  ) public pure returns (CCIPHome.OCR3Config[] memory commitConfigs, CCIPHome.OCR3Config[] memory execConfigs) {
    return _groupByPluginType(ocr3Configs);
  }

  function computeConfigDigest(
    uint32 donId,
    uint64 configCount,
    CCIPHome.OCR3Config memory ocr3Config
  ) public pure returns (bytes32) {
    return _computeConfigDigest(donId, configCount, ocr3Config);
  }

  function validateConfig(CCIPHome.OCR3Config memory cfg) public view {
    _validateConfig(cfg);
  }

  function updatePluginConfig(
    uint32 donId,
    Internal.OCRPluginType pluginType,
    CCIPHome.OCR3Config[] memory newConfig
  ) public {
    _updatePluginConfig(donId, pluginType, newConfig);
  }
}
