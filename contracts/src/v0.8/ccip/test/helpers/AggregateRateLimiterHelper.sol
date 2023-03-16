// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../AggregateRateLimiter.sol";

contract AggregateRateLimiterHelper is AggregateRateLimiter {
  constructor(RateLimiterConfig memory config) AggregateRateLimiter(config) {}

  function removeTokens(Client.EVMTokenAmount[] memory tokenAmounts) public {
    _removeTokens(tokenAmounts);
  }
}
