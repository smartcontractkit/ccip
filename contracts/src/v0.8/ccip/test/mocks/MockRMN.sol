
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
  mapping(bytes32 => bool) private s_subjectCurses;

  function isCursed() external view override returns (bool) {
    if (s_err.length != 0) {
      revert CustomError(s_err);
    }
    return s_curse;
  }

  function isCursed(bytes32 subject) external view override returns (bool) {
    if (s_err.length != 0) {
      revert CustomError(s_err);
    }
    return s_curse || s_subjectCurses[subject];
  }

  function voteToCurse() external {
    s_curse = true;
  }

  function voteToCurse(bytes32 subject) external {
    s_subjectCurses[subject] = true;
  }

  function ownerUnvoteToCurse(RMN.UnvoteToCurseRecord[] memory) external {
    s_curse = false;
  }

  function ownerUnvoteToCurse(RMN.UnvoteToCurseRecord[] memory, bytes32 subject) external {
    s_subjectCurses[subject] = false;
  }

  function setRevert(bytes memory err) external {
    s_err = err;
  }

  function isBlessed(IRMN.TaggedRoot calldata) external view override returns (bool) {
    return !s_curse;
  }

  function getConfigDetails() external view returns (uint32 version, uint32 blockNumber, RMN.Config memory config) {
    return (s_versionedConfig.configVersion, s_versionedConfig.blockNumber, s_versionedConfig.config);
  }
}
