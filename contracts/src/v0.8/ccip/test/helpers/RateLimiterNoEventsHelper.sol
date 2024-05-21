// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RateLimiterNoEvents} from "../../libraries/RateLimiterNoEvents.sol";

contract RateLimiterNoEventsHelper {
  using RateLimiterNoEvents for RateLimiterNoEvents.TokenBucket;

  RateLimiterNoEvents.TokenBucket internal s_rateLimiter;

  constructor(RateLimiterNoEvents.Config memory config) {
    s_rateLimiter = RateLimiterNoEvents.TokenBucket({
      rate: config.rate,
      capacity: config.capacity,
      tokens: config.capacity,
      lastUpdated: uint32(block.timestamp),
      isEnabled: config.isEnabled
    });
  }

  function consume(uint256 requestTokens, address tokenAddress) external {
    s_rateLimiter._consume(requestTokens, tokenAddress);
  }

  function currentTokenBucketState() external view returns (RateLimiterNoEvents.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  function setTokenBucketConfig(RateLimiterNoEvents.Config memory config) external {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  function getRateLimiter() external view returns (RateLimiterNoEvents.TokenBucket memory) {
    return s_rateLimiter;
  }
}
