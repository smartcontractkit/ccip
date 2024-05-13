// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMN} from "../../RMN.sol";
import {IRMN} from "../../interfaces/IRMN.sol";
import {OwnerIsCreator} from "./../../../shared/access/OwnerIsCreator.sol";

contract MockRMN is IRMN, OwnerIsCreator {
  error CustomError(bytes err);

  bool private s_curse;
  bytes private s_err;
  RMN.VersionedConfig private s_versionedConfig;

  function isCursed() external view override returns (bool) {
    if (s_err.length != 0) {
      revert CustomError(s_err);
    }
    return s_curse;
  }

  function isCursed(bytes32 /* subject */ ) external view override returns (bool) {
    if (s_err.length != 0) {
      revert CustomError(s_err);
    }
    return s_curse;
  }

  function voteToCurse(bytes32) external {
    s_curse = true;
  }

  function setRevert(bytes memory err) external {
    s_err = err;
  }

  function ownerUnvoteToCurse(RMN.UnvoteToCurseRecord[] memory) external {
    s_curse = false;
  }

  function isBlessed(IRMN.TaggedRoot calldata) external view override returns (bool) {
    return !s_curse;
  }

  function getConfigDetails() external view returns (uint32 version, uint32 blockNumber, RMN.Config memory config) {
    version = s_versionedConfig.configVersion;
    blockNumber = s_versionedConfig.blockNumber;
    config = s_versionedConfig.config;
  }
}
