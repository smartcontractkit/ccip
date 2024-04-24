// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {RateLimiterNoEvents} from "../../libraries/RateLimiterNoEvents.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {RateLimiterNoEventsHelper} from "../helpers/RateLimiterNoEventsHelper.sol";

contract RateLimiterNoEventsSetup is BaseTest {
  RateLimiterNoEventsHelper internal s_helper;
  RateLimiterNoEvents.Config internal s_config;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_config = RateLimiterNoEvents.Config({isEnabled: true, rate: 5, capacity: 100});
    s_helper = new RateLimiterNoEventsHelper(s_config);
  }
}

contract RateLimiterNoEvents_constructor is RateLimiterNoEventsSetup {
  function test_Constructor_Success() public view {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);
  }
}

/// @notice #setTokenBucketConfig
contract RateLimiterNoEvents_setTokenBucketConfig is RateLimiterNoEventsSetup {
  function test_SetRateLimiterNoEventsConfig_Success() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);

    s_config =
      RateLimiterNoEvents.Config({isEnabled: true, rate: uint128(rateLimiter.rate * 2), capacity: rateLimiter.capacity * 8});

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
contract RateLimiterNoEvents_currentTokenBucketState is RateLimiterNoEventsSetup {
  function test_CurrentTokenBucketState_Success() public {
    RateLimiterNoEvents.TokenBucket memory bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    s_config = RateLimiterNoEvents.Config({isEnabled: true, rate: uint128(bucket.rate * 2), capacity: bucket.capacity * 8});

    s_helper.setTokenBucketConfig(s_config);

    bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity / 8, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);
  }

  function test_Refill_Success() public {
    RateLimiterNoEvents.TokenBucket memory bucket = s_helper.currentTokenBucketState();
    assertEq(s_config.rate, bucket.rate);
    assertEq(s_config.capacity, bucket.capacity);
    assertEq(s_config.capacity, bucket.tokens);
    assertEq(s_config.isEnabled, bucket.isEnabled);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    s_config = RateLimiterNoEvents.Config({isEnabled: true, rate: uint128(bucket.rate * 2), capacity: bucket.capacity * 8});

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
contract RateLimiterNoEvents_consume is RateLimiterNoEventsSetup {
  address internal s_token = address(100);

  function test_ConsumeAggregateValue_Success() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);

    uint256 requestTokens = 50;

    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);
  }

  function test_ConsumeTokens_Success() public {
    uint256 requestTokens = 50;

    s_helper.consume(requestTokens, s_token);
  }

  function test_Refill_Success() public {
    uint256 requestTokens = 50;

    s_helper.consume(requestTokens, address(0));

    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME, rateLimiter.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(BLOCK_TIME + warpTime);

    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(s_config.capacity - requestTokens * 2 + warpTime * s_config.rate, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
    assertEq(BLOCK_TIME + warpTime, rateLimiter.lastUpdated);
  }

  function test_ConsumeUnlimited_Success() public {
    s_helper.consume(0, address(0));

    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.capacity, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);

    RateLimiterNoEvents.Config memory disableConfig = RateLimiterNoEvents.Config({isEnabled: false, rate: 0, capacity: 0});

    s_helper.setTokenBucketConfig(disableConfig);

    uint256 requestTokens = 50;
    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(disableConfig.capacity, rateLimiter.tokens);
    assertEq(disableConfig.isEnabled, rateLimiter.isEnabled);

    s_helper.setTokenBucketConfig(s_config);

    vm.expectRevert(abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueRateLimitReached.selector, 10, 0));
    s_helper.consume(requestTokens, address(0));

    rateLimiter = s_helper.getRateLimiter();
    assertEq(s_config.rate, rateLimiter.rate);
    assertEq(s_config.capacity, rateLimiter.capacity);
    assertEq(0, rateLimiter.tokens);
    assertEq(s_config.isEnabled, rateLimiter.isEnabled);
  }

  // Reverts

  function test_AggregateValueMaxCapacityExceeded_Revert() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiterNoEvents.AggregateValueMaxCapacityExceeded.selector, rateLimiter.capacity, rateLimiter.capacity + 1
      )
    );
    s_helper.consume(rateLimiter.capacity + 1, address(0));
  }

  function test_TokenMaxCapacityExceeded_Revert() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiterNoEvents.TokenMaxCapacityExceeded.selector, rateLimiter.capacity, rateLimiter.capacity + 1, s_token
      )
    );
    s_helper.consume(rateLimiter.capacity + 1, s_token);
  }

  function test_ConsumingMoreThanUint128_Revert() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 request = uint256(type(uint128).max) + 1;

    vm.expectRevert(
      abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueMaxCapacityExceeded.selector, rateLimiter.capacity, request)
    );
    s_helper.consume(request, address(0));
  }

  function test_AggregateValueRateLimitReached_Revert() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 overLimit = 20;
    uint256 requestTokens1 = rateLimiter.capacity / 2;
    uint256 requestTokens2 = rateLimiter.capacity / 2 + overLimit;

    uint256 waitInSeconds = overLimit / rateLimiter.rate;

    s_helper.consume(requestTokens1, address(0));

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiterNoEvents.AggregateValueRateLimitReached.selector, waitInSeconds, rateLimiter.capacity - requestTokens1
      )
    );
    s_helper.consume(requestTokens2, address(0));
  }

  function test_TokenRateLimitReached_Revert() public {
    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    uint256 overLimit = 20;
    uint256 requestTokens1 = rateLimiter.capacity / 2;
    uint256 requestTokens2 = rateLimiter.capacity / 2 + overLimit;

    uint256 waitInSeconds = overLimit / rateLimiter.rate;

    s_helper.consume(requestTokens1, s_token);

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiterNoEvents.TokenRateLimitReached.selector, waitInSeconds, rateLimiter.capacity - requestTokens1, s_token
      )
    );
    s_helper.consume(requestTokens2, s_token);
  }

  function test_RateLimitReachedOverConsecutiveBlocks_Revert() public {
    uint256 initBlockTime = BLOCK_TIME + 10000;
    vm.warp(initBlockTime);

    RateLimiterNoEvents.TokenBucket memory rateLimiter = s_helper.getRateLimiter();

    s_helper.consume(rateLimiter.capacity, address(0));

    vm.warp(initBlockTime + 1);

    // Over rate limit by 1, force 1 second wait
    uint256 overLimit = 1;

    vm.expectRevert(abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueRateLimitReached.selector, 1, rateLimiter.rate));
    s_helper.consume(rateLimiter.rate + overLimit, address(0));
  }
}
