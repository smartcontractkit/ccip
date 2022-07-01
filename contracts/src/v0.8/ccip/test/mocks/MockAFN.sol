// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../health/interfaces/AFNInterface.sol";

contract MockAFN is AFNInterface {
  Heartbeat public s_lastHeartbeat;
  bool public s_badSignal;

  constructor() {
    s_lastHeartbeat = Heartbeat({round: 1, timestamp: block.timestamp, committeeVersion: 1});
  }

  function setTimestamp(uint64 newTimestamp) external {
    s_lastHeartbeat.timestamp = newTimestamp;
  }

  function hasBadSignal() external view override returns (bool) {
    return s_badSignal;
  }

  function getLastHeartbeat() external view override returns (Heartbeat memory) {
    return s_lastHeartbeat;
  }

  function voteGood(
    uint256 /*round*/
  ) external override {
    s_badSignal = false;
  }

  function voteBad() external override {
    s_badSignal = true;
  }

  function recover() external override {
    s_badSignal = false;
  }

  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external override {
    // nothing
  }
}
