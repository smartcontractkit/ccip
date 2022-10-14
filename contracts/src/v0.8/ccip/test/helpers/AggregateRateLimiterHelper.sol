// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../rateLimiter/AggregateRateLimiter.sol";

contract AggregateRateLimiterHelper is AggregateRateLimiter {
  constructor(RateLimiterConfig memory config, address tokenLimitsAdmin)
    AggregateRateLimiter(config, tokenLimitsAdmin)
  {}

  function removeTokens(address[] memory tokens, uint256[] memory amounts) public {
    _removeTokens(tokens, amounts);
  }
}
