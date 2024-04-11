// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../../AggregateRateLimiter.sol";

contract AggregateRateLimiterHelper is AggregateRateLimiter {
  constructor(RateLimiter.Config memory config) AggregateRateLimiter(config) {}

  function rateLimitValue(uint256 value) public {
    _rateLimitValue(value);
  }
}
