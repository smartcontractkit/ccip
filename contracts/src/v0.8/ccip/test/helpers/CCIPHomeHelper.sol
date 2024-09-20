// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {CCIPHome} from "../../capability/CCIPHome.sol";
import {Internal} from "../../libraries/Internal.sol";

contract CCIPHomeHelper is CCIPHome {
  constructor(address capabilitiesRegistry) CCIPHome(capabilitiesRegistry) {}

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
    CCIPHome.OCR3Config[] memory newConfig
  ) public view returns (CCIPHome.OCR3ConfigWithMeta[] memory) {
    return _computeNewConfigWithMeta(donId, currentConfig, newConfig);
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
