// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {RateLimiter} from "../../models/RateLimiter.sol";

contract RateLimiterHelper {
  using RateLimiter for RateLimiter.TokenBucket;

  RateLimiter.TokenBucket internal s_rateLimiter;

  constructor(RateLimiter.Config memory config) {
    s_rateLimiter = RateLimiter.TokenBucket({
      rate: config.rate,
      capacity: config.capacity,
      tokens: config.capacity,
      lastUpdated: uint40(block.timestamp),
      isEnabled: config.isEnabled
    });
  }

  function consume(uint256 requestTokens) external {
    s_rateLimiter._consume(requestTokens);
  }

  function currentTokenBucketState() external view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  function setTokenBucketConfig(RateLimiter.Config memory config) external {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  function getRateLimiter() external view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter;
  }
}
