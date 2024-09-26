// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {CCIPHome} from "../../capability/CCIPHome.sol";

contract CCIPHomeHelper is CCIPHome {
  constructor(address capabilitiesRegistry) CCIPHome(capabilitiesRegistry) {}

  function validateConfig(OCR3Config memory cfg) external view {
    return _validateConfig(cfg);
  }

  function ensureInRegistry(bytes32[] memory p2pIds) external view {
    return _ensureInRegistry(p2pIds);
  }
}
