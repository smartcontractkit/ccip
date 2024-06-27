// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMN} from "src/RMN.sol";
import {IRMN} from "src/IRMN.sol";
import {OwnerIsCreator} from "src/OwnerIsCreator.sol";

contract MockRMN is IRMN, OwnerIsCreator {
  error CustomError(bytes err);

  bool private s_curse;
  bytes private s_err;
  RMN.VersionedConfig private s_versionedConfig;

  function isCursed() public view override returns (bool) {
    if (s_err.length != 0) {
      revert CustomError(s_err);
    }
    return s_curse;
  }

  function isCursed(bytes16) external view override returns (bool) {
    return isCursed();
  }

  function voteToCurse(bytes32) public {
    s_curse = true;
  }

  function voteToCurse(bytes16 curseId, bytes16[] memory) external {
    voteToCurse(curseId);
  }

  function setRevert(bytes memory err) external {
    s_err = err;
  }

  function ownerUnvoteToCurse(RMN.OwnerUnvoteToCurseRequest[] memory) external {
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
