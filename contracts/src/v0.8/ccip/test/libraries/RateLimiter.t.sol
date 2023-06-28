// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {BaseTest} from "../BaseTest.t.sol";
import {RateLimiterHelper} from "../helpers/RateLimiterHelper.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";

contract RateLimiterSetup is BaseTest {
  RateLimiterHelper internal s_helper;
  RateLimiter.Config internal s_config;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_config = RateLimiter.Config({isEnabled: true, rate: 5, capacity: 100});
    s_helper = new RateLimiterHelper(s_config);
  }
}

contract RateLimiter_constructor is RateLimiterSetup {
  function testConstructorSuccess() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);
  }
}

/// @notice #setTokenBucketConfig
contract RateLimiter_setTokenBucketConfig is RateLimiterSetup {
  event ConfigChanged(RateLimiter.Config config);

  function testSetRateLimiterConfigSuccess() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);

    s_config = RateLimiter.Config({
      isEnabled: true,
      rate: uint128(rateLimiter.rate * 2),
      capacity: rateLimiter.capacity * 8
    });

    vm.expectEmit();
    emit ConfigChanged(s_config);

    s_helper.setTokenBucketConfig(s_config);

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity / 8, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);
  }
}

/// @notice #currentTokenBucketState
contract RateLimiter_currentTokenBucketState is RateLimiterSetup {
  function testCurrentTokenBucketStateSuccess() public {
    RateLimiter.TokenBucket memory bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    s_config = RateLimiter.Config({isEnabled: true, rate: uint128(bucket.rate * 2), capacity: bucket.capacity * 8});

    s_helper.setTokenBucketConfig(s_config);

    bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity / 8, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);
  }

  function testRefillSuccess() public {
    RateLimiter.TokenBucket memory bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    s_config = RateLimiter.Config({isEnabled: true, rate: uint128(bucket.rate * 2), capacity: bucket.capacity * 8});

    s_helper.setTokenBucketConfig(s_config);

    bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity / 8, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(BLOCK_TIME + warpTime);

    bucket = s_helper.currentTokenBucketState();

    assertEq(s_config.capacity / 8 + warpTime * s_config.rate, bucket.tokens);

    vm.warp(BLOCK_TIME + warpTime * 100);

    // Bucket overflow
    bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.capacity, bucket.tokens);
  }
}

/// @notice #consume
contract RateLimiter_consume is RateLimiterSetup {
  event TokensConsumed(uint256 tokens);

  function testConsumeSuccess() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);

    uint256 requestTokens = 50;

    vm.expectEmit();
    emit TokensConsumed(requestTokens);

    s_helper.consume(requestTokens);

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);
  }

  function testRefillSuccess() public {
    uint256 requestTokens = 50;

    vm.expectEmit();
    emit TokensConsumed(requestTokens);

    s_helper.consume(requestTokens);

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(BLOCK_TIME + warpTime);

    vm.expectEmit();
    emit TokensConsumed(requestTokens);

    s_helper.consume(requestTokens);

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens * 2 + warpTime * s_config.rate, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME + warpTime, rateLimiter.lastUpdated);
  }

  function testConsumeUnlimitedSuccess() public {
    s_helper.consume(0);

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);

    RateLimiter.Config memory disableConfig = RateLimiter.Config({isEnabled: false, rate: 5, capacity: 100});

    s_helper.setTokenBucketConfig(disableConfig);

    uint256 requestTokens = 50;
    s_helper.consume(requestTokens);

    rateLimiter = s_helper.getRateLimiter();
    assertEq(disableConfig.capacity, rateLimiter.tokens);
    assertEq(disableConfig.isEnabled, rateLimiter.isEnabled);

    s_helper.setTokenBucketConfig(s_config);

    s_helper.consume(requestTokens);

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
  }

  // Reverts

  function testConsumingMoreThanMaxCapacityReverts() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiter.ConsumingMoreThanMaxCapacity.selector,
        rateLimiter.capacity,
        rateLimiter.capacity + 1
      )
    );
    s_helper.consume(rateLimiter.capacity + 1);
  }

  function testConsumingMoreThanUint128Reverts() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 request = uint256(type(uint128).max) + 1;

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiter.ConsumingMoreThanMaxCapacity.selector, rateLimiter.capacity, request)
    );
    s_helper.consume(request);
  }

  function testRateLimitReachedReverts() public {
    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 overLimit = 20;
    uint256 requestTokens1 = rateLimiter.capacity / 2;
    uint256 requestTokens2 = rateLimiter.capacity / 2 + overLimit;

    uint256 waitInSeconds = overLimit / rateLimiter.rate;

    s_helper.consume(requestTokens1);

    vm.expectRevert(abi.encodeWithSelector(RateLimiter.RateLimitReached.selector, waitInSeconds));
    s_helper.consume(requestTokens2);
  }

  function testRateLimitReachedOverConsecutiveBlocksReverts() public {
    uint256 initBlockTime = BLOCK_TIME + 10000;
    vm.warp(initBlockTime);

    RateLimiter.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    vm.expectEmit();
    emit TokensConsumed(rateLimiter.capacity);

    s_helper.consume(rateLimiter.capacity);

    vm.warp(initBlockTime + 1);

    // Over rate limit by 1, force 1 second wait
    uint256 overLimit = 1;

    vm.expectRevert(abi.encodeWithSelector(RateLimiter.RateLimitReached.selector, 1));
    s_helper.consume(rateLimiter.rate + overLimit);
  }
}
