// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../health/HealthChecker.sol";

contract HealthCheckerHelper is HealthChecker {
  constructor(IAFN afn) HealthChecker(afn) {}

  function whenHealthyFunction() external whenHealthy {}
}
