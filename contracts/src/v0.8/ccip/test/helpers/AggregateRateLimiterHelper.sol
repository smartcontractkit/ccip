// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../AggregateRateLimiter.sol";

contract AggregateRateLimiterHelper is AggregateRateLimiter {
  constructor(RateLimiter.Config memory config) AggregateRateLimiter(config) {}

  function rateLimitValue(Client.EVMTokenAmount[] memory tokenAmounts) public {
    _rateLimitValue(tokenAmounts);
  }
}
