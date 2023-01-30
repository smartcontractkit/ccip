// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../rateLimiter/AggregateRateLimiter.sol";

contract AggregateRateLimiterHelper is AggregateRateLimiter {
  constructor(RateLimiterConfig memory config) AggregateRateLimiter(config) {}

  function removeTokens(Common.EVMTokenAndAmount[] memory tokensAndAmounts) public {
    _removeTokens(tokensAndAmounts);
  }
}
