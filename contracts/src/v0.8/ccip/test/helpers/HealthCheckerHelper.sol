// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../health/HealthChecker.sol";

contract HealthCheckerHelper is HealthChecker {
  constructor(AFNInterface afn, uint256 maxTimeWithoutAFNSignal) HealthChecker(afn, maxTimeWithoutAFNSignal) {}

  function whenHealthyFunction() external whenHealthy {}
}
