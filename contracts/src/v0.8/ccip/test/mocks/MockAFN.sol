// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAFN} from "../../interfaces/IAFN.sol";

contract MockAFN is IAFN {
  bool public s_badSignal;

  function badSignalReceived() external view override returns (bool) {
    return s_badSignal;
  }

  function voteToBlessRoots(
    bytes32[] calldata // rootsWithOrigin
  ) external override {
    s_badSignal = false;
  }

  function voteBad() external override {
    s_badSignal = true;
  }

  function recoverFromBadSignal() external override {
    s_badSignal = false;
  }

  function isBlessed(bytes32) external view override returns (bool) {
    return !s_badSignal;
  }

  function getWeightThresholds() external override returns (uint256 blessing, uint256 badSignal) {}

  function getParticipants() external override returns (address[] memory) {}

  function getVotesToBlessRoot(bytes32 root) public view override returns (uint256) {}

  function getWeightByParticipant(address) external view override returns (uint256) {}

  function hasVotedToBlessRoot(address participant, bytes32 root) public view override returns (bool) {}

  function getConfigVersion() external view override returns (uint256) {}

  function getBadVotersAndVotes() external view override returns (address[] memory voters, uint256 votes) {}

  function hasVotedBad(address participant) external view override returns (bool) {}

  function setAFNConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external override {
    // nothing
  }
}
