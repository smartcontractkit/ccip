// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Vm} from "forge-std/Vm.sol";

import {Client} from "../../libraries/Client.sol";
import {Internal} from "../../libraries/Internal.sol";
import {RateLimiterNoEvents} from "../../libraries/RateLimiterNoEvents.sol";
import {MultiAggregateRateLimiter} from "../../validators/MultiAggregateRateLimiter.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {MultiAggregateRateLimiterHelper} from "../helpers/MultiAggregateRateLimiterHelper.sol";
import {PriceRegistrySetup} from "../priceRegistry/PriceRegistry.t.sol";

import {stdError} from "forge-std/Test.sol";

contract MultiAggregateRateLimiterSetup is BaseTest, PriceRegistrySetup {
  MultiAggregateRateLimiterHelper internal s_rateLimiter;

  address internal immutable TOKEN = 0x21118E64E1fB0c487F25Dd6d3601FF6af8D32E4e;
  uint224 internal constant TOKEN_PRICE = 4e18;

  uint64 internal constant CHAIN_SELECTOR_1 = 5009297550715157269;
  uint64 internal constant CHAIN_SELECTOR_2 = 4949039107694359620;

  RateLimiterNoEvents.Config internal RATE_LIMITER_CONFIG_1 =
    RateLimiterNoEvents.Config({isEnabled: true, rate: 5, capacity: 100});
  RateLimiterNoEvents.Config internal RATE_LIMITER_CONFIG_2 =
    RateLimiterNoEvents.Config({isEnabled: true, rate: 10, capacity: 200});

  address internal immutable MOCK_OFFRAMP = address(1111);
  address internal immutable MOCK_ONRAMP = address(1112);

  function setUp() public virtual override(BaseTest, PriceRegistrySetup) {
    BaseTest.setUp();
    PriceRegistrySetup.setUp();

    Internal.PriceUpdates memory priceUpdates = getSingleTokenPriceUpdateStruct(TOKEN, TOKEN_PRICE);
    s_priceRegistry.updatePrices(priceUpdates);

    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](2);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });
    configUpdates[1] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_2,
      rateLimiterConfig: RATE_LIMITER_CONFIG_2
    });

    address[] memory authorizedCallers = new address[](2);
    authorizedCallers[0] = MOCK_OFFRAMP;
    authorizedCallers[1] = MOCK_ONRAMP;

    s_rateLimiter =
      new MultiAggregateRateLimiterHelper(configUpdates, ADMIN, address(s_priceRegistry), authorizedCallers);
  }

  function _assertConfigWithTokenBucketEquality(
    RateLimiterNoEvents.Config memory config,
    RateLimiterNoEvents.TokenBucket memory tokenBucket
  ) internal pure {
    assertEq(config.rate, tokenBucket.rate);
    assertEq(config.capacity, tokenBucket.capacity);
    assertEq(config.capacity, tokenBucket.tokens);
    assertEq(config.isEnabled, tokenBucket.isEnabled);
  }

  function _assertTokenBucketEquality(
    RateLimiterNoEvents.TokenBucket memory tokenBucketA,
    RateLimiterNoEvents.TokenBucket memory tokenBucketB
  ) internal pure {
    assertEq(tokenBucketA.rate, tokenBucketB.rate);
    assertEq(tokenBucketA.capacity, tokenBucketB.capacity);
    assertEq(tokenBucketA.tokens, tokenBucketB.tokens);
    assertEq(tokenBucketA.isEnabled, tokenBucketB.isEnabled);
  }
}

/// @notice #constructor
contract MultiAggregateRateLimiter_constructor is MultiAggregateRateLimiterSetup {
  function test_ConstructorNoAuthorizedCallers_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](0);
    address[] memory authorizedCallers = new address[](0);

    vm.recordLogs();
    s_rateLimiter =
      new MultiAggregateRateLimiterHelper(configUpdates, ADMIN, address(s_priceRegistry), authorizedCallers);

    // AdminSet + PriceRegistrySet
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 2);

    assertEq(ADMIN, s_rateLimiter.getTokenLimitAdmin());
    assertEq(OWNER, s_rateLimiter.owner());
    assertEq(address(s_priceRegistry), s_rateLimiter.getPriceRegistry());
  }

  function test_ConstructorNoConfigs_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](0);
    address[] memory authorizedCallers = new address[](2);
    authorizedCallers[0] = MOCK_OFFRAMP;
    authorizedCallers[1] = MOCK_ONRAMP;

    vm.recordLogs();
    s_rateLimiter =
      new MultiAggregateRateLimiterHelper(configUpdates, ADMIN, address(s_priceRegistry), authorizedCallers);

    // AdminSet + PriceRegistrySet + 2 authorized caller sets
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 4);

    assertEq(ADMIN, s_rateLimiter.getTokenLimitAdmin());
    assertEq(OWNER, s_rateLimiter.owner());
    assertEq(address(s_priceRegistry), s_rateLimiter.getPriceRegistry());
  }

  function test_Constructor_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](2);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });
    configUpdates[1] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_2,
      rateLimiterConfig: RATE_LIMITER_CONFIG_2
    });

    address[] memory authorizedCallers = new address[](2);
    authorizedCallers[0] = MOCK_OFFRAMP;
    authorizedCallers[1] = MOCK_ONRAMP;

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(CHAIN_SELECTOR_1, RATE_LIMITER_CONFIG_1);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(CHAIN_SELECTOR_2, RATE_LIMITER_CONFIG_2);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.PriceRegistrySet(address(s_priceRegistry));

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(MOCK_OFFRAMP);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(MOCK_ONRAMP);

    s_rateLimiter =
      new MultiAggregateRateLimiterHelper(configUpdates, ADMIN, address(s_priceRegistry), authorizedCallers);

    assertEq(ADMIN, s_rateLimiter.getTokenLimitAdmin());
    assertEq(OWNER, s_rateLimiter.owner());
    assertEq(address(s_priceRegistry), s_rateLimiter.getPriceRegistry());

    RateLimiterNoEvents.TokenBucket memory bucketSrcChain1 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    _assertConfigWithTokenBucketEquality(RATE_LIMITER_CONFIG_1, bucketSrcChain1);
    assertEq(BLOCK_TIME, bucketSrcChain1.lastUpdated);

    RateLimiterNoEvents.TokenBucket memory bucketSrcChain2 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_2);
    _assertConfigWithTokenBucketEquality(RATE_LIMITER_CONFIG_2, bucketSrcChain2);
    assertEq(BLOCK_TIME, bucketSrcChain2.lastUpdated);
  }
}

/// @notice #getTokenLimitAdmin
contract MultiAggregateRateLimiter_getTokenLimitAdmin is MultiAggregateRateLimiterSetup {
  function test_GetTokenLimitAdmin_Success() public view {
    assertEq(ADMIN, s_rateLimiter.getTokenLimitAdmin());
  }
}

/// @notice #setAdmin
contract MultiAggregateRateLimiter_setAdmin is MultiAggregateRateLimiterSetup {
  function test_Owner_Success() public {
    vm.expectEmit();
    emit MultiAggregateRateLimiter.AdminSet(STRANGER);

    s_rateLimiter.setAdmin(STRANGER);
    assertEq(STRANGER, s_rateLimiter.getTokenLimitAdmin());
  }

  function test_Admin_Success() public {
    vm.startPrank(ADMIN);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AdminSet(STRANGER);

    s_rateLimiter.setAdmin(STRANGER);
    assertEq(STRANGER, s_rateLimiter.getTokenLimitAdmin());
  }

  // Reverts

  function test_OnlyOwnerOrAdmin_Revert() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(RateLimiterNoEvents.OnlyCallableByAdminOrOwner.selector);

    s_rateLimiter.setAdmin(STRANGER);
  }
}

/// @notice #setPriceRegistry
contract MultiAggregateRateLimiter_setPriceRegistry is MultiAggregateRateLimiterSetup {
  function test_Owner_Success() public {
    address newAddress = address(42);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.PriceRegistrySet(newAddress);

    s_rateLimiter.setPriceRegistry(newAddress);
    assertEq(newAddress, s_rateLimiter.getPriceRegistry());
  }

  function test_Admin_Success() public {
    vm.startPrank(ADMIN);

    address newAddress = address(42);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.PriceRegistrySet(newAddress);

    s_rateLimiter.setPriceRegistry(newAddress);
    assertEq(newAddress, s_rateLimiter.getPriceRegistry());
  }

  // Reverts

  function test_OnlyOwnerOrAdmin_Revert() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(RateLimiterNoEvents.OnlyCallableByAdminOrOwner.selector);

    s_rateLimiter.setPriceRegistry(STRANGER);
  }

  function test_ZeroAddress_Revert() public {
    vm.expectRevert(MultiAggregateRateLimiter.ZeroAddressNotAllowed.selector);
    s_rateLimiter.setPriceRegistry(address(0));
  }
}

/// @notice #setPriceRegistry
contract MultiAggregateRateLimiter_setAuthorizedCallers is MultiAggregateRateLimiterSetup {
  function test_OnlyAdd_Success() public {
    address[] memory addedCallers = new address[](2);
    addedCallers[0] = address(42);
    addedCallers[1] = address(43);

    address[] memory removedCallers = new address[](0);

    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));
    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[1]));

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(addedCallers[0]);
    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(addedCallers[1]);

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertTrue(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));
    assertTrue(s_rateLimiter.isAuthorizedCaller(addedCallers[1]));
  }

  function test_OnlyRemove_Success() public {
    address[] memory addedCallers = new address[](0);

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = MOCK_OFFRAMP;

    assertTrue(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerRemoved(removedCallers[0]);

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertFalse(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));
  }

  function test_AddAndRemove_Success() public {
    address[] memory addedCallers = new address[](2);
    addedCallers[0] = address(42);
    addedCallers[1] = address(43);

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = MOCK_OFFRAMP;

    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));
    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[1]));
    assertTrue(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(addedCallers[0]);
    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(addedCallers[1]);
    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerRemoved(removedCallers[0]);

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertTrue(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));
    assertTrue(s_rateLimiter.isAuthorizedCaller(addedCallers[1]));
    assertFalse(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));
  }

  function test_AddThenRemove_Success() public {
    address[] memory addedCallers = new address[](1);
    addedCallers[0] = address(42);

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = address(42);

    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));

    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerAdded(addedCallers[0]);
    vm.expectEmit();
    emit MultiAggregateRateLimiter.AuthorizedCallerRemoved(addedCallers[0]);

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertFalse(s_rateLimiter.isAuthorizedCaller(addedCallers[0]));
  }

  function test_SkipRemove_Success() public {
    address[] memory addedCallers = new address[](0);

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = address(42);

    vm.recordLogs();
    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertFalse(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));

    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);
  }

  function test_ChangedByAdmin_Success() public {
    vm.startPrank(ADMIN);

    address[] memory addedCallers = new address[](0);

    address[] memory removedCallers = new address[](1);
    removedCallers[0] = MOCK_OFFRAMP;

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );

    assertFalse(s_rateLimiter.isAuthorizedCaller(removedCallers[0]));
  }

  // Reverts

  function test_OnlyOwnerOrAdmin_Revert() public {
    vm.startPrank(STRANGER);
    vm.expectRevert(RateLimiterNoEvents.OnlyCallableByAdminOrOwner.selector);

    address[] memory addedCallers = new address[](2);
    addedCallers[0] = address(42);
    addedCallers[1] = address(43);

    address[] memory removedCallers = new address[](0);

    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );
  }

  function test_ZeroAddressAdd_Revert() public {
    address[] memory addedCallers = new address[](1);
    addedCallers[0] = address(0);
    address[] memory removedCallers = new address[](0);

    vm.expectRevert(MultiAggregateRateLimiter.ZeroAddressNotAllowed.selector);
    s_rateLimiter.applyAuthorizedCallerUpdates(
      MultiAggregateRateLimiter.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: removedCallers})
    );
  }
}

/// @notice #getTokenBucket
contract MultiAggregateRateLimiter_getTokenBucket is MultiAggregateRateLimiterSetup {
  function test_GetTokenBucket_Success() public view {
    RateLimiterNoEvents.TokenBucket memory bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    assertEq(RATE_LIMITER_CONFIG_1.rate, bucket.rate);
    assertEq(RATE_LIMITER_CONFIG_1.capacity, bucket.capacity);
    assertEq(RATE_LIMITER_CONFIG_1.capacity, bucket.tokens);
    assertEq(BLOCK_TIME, bucket.lastUpdated);
  }

  function test_Refill_Success() public {
    RATE_LIMITER_CONFIG_1.capacity = RATE_LIMITER_CONFIG_1.capacity * 2;

    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });

    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    RateLimiterNoEvents.TokenBucket memory bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);

    assertEq(RATE_LIMITER_CONFIG_1.rate, bucket.rate);
    assertEq(RATE_LIMITER_CONFIG_1.capacity, bucket.capacity);
    assertEq(RATE_LIMITER_CONFIG_1.capacity / 2, bucket.tokens);
    assertEq(BLOCK_TIME, bucket.lastUpdated);

    uint256 warpTime = 4;
    vm.warp(BLOCK_TIME + warpTime);

    bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);

    assertEq(RATE_LIMITER_CONFIG_1.rate, bucket.rate);
    assertEq(RATE_LIMITER_CONFIG_1.capacity, bucket.capacity);
    assertEq(RATE_LIMITER_CONFIG_1.capacity / 2 + warpTime * RATE_LIMITER_CONFIG_1.rate, bucket.tokens);
    assertEq(BLOCK_TIME + warpTime, bucket.lastUpdated);

    vm.warp(BLOCK_TIME + warpTime * 100);

    // Bucket overflow
    bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    assertEq(RATE_LIMITER_CONFIG_1.capacity, bucket.tokens);
  }

  // Reverts

  function test_TimeUnderflow_Revert() public {
    vm.warp(BLOCK_TIME - 1);

    vm.expectRevert(stdError.arithmeticError);
    s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
  }
}

/// @notice #applyRateLimiterConfigUpdates
contract MultiAggregateRateLimiter_applyRateLimiterConfigUpdates is MultiAggregateRateLimiterSetup {
  function test_ZeroConfigs_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](0);

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);
  }

  function test_SingleConfig_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1 + 1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(
      configUpdates[0].chainSelector, configUpdates[0].rateLimiterConfig
    );

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 1);

    RateLimiterNoEvents.TokenBucket memory bucket1 =
      s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);
    _assertConfigWithTokenBucketEquality(configUpdates[0].rateLimiterConfig, bucket1);
    assertEq(BLOCK_TIME, bucket1.lastUpdated);
  }

  function test_SingleConfigByAdmin_Success() public {
    vm.startPrank(ADMIN);
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1 + 1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(
      configUpdates[0].chainSelector, configUpdates[0].rateLimiterConfig
    );

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 1);

    RateLimiterNoEvents.TokenBucket memory bucket1 =
      s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);
    _assertConfigWithTokenBucketEquality(configUpdates[0].rateLimiterConfig, bucket1);
    assertEq(BLOCK_TIME, bucket1.lastUpdated);
  }

  function test_MultipleConfigs_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](3);

    for (uint64 i; i < configUpdates.length; ++i) {
      configUpdates[i] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
        chainSelector: CHAIN_SELECTOR_1 + i + 1,
        rateLimiterConfig: RateLimiterNoEvents.Config({isEnabled: true, rate: 5 + i, capacity: 100 + i})
      });

      vm.expectEmit();
      emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(
        configUpdates[i].chainSelector, configUpdates[i].rateLimiterConfig
      );
    }

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, configUpdates.length);

    for (uint256 i; i < configUpdates.length; ++i) {
      RateLimiterNoEvents.TokenBucket memory bucket =
        s_rateLimiter.currentRateLimiterState(configUpdates[i].chainSelector);
      _assertConfigWithTokenBucketEquality(configUpdates[i].rateLimiterConfig, bucket);
      assertEq(BLOCK_TIME, bucket.lastUpdated);
    }
  }

  function test_UpdateExistingConfig_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_2
    });

    RateLimiterNoEvents.TokenBucket memory bucket1 =
      s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);

    // Capacity equals tokens
    assertEq(bucket1.capacity, bucket1.tokens);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(
      configUpdates[0].chainSelector, configUpdates[0].rateLimiterConfig
    );

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    vm.warp(BLOCK_TIME + 1);
    bucket1 = s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);
    assertEq(BLOCK_TIME + 1, bucket1.lastUpdated);

    // Tokens < capacity since capacity doubled
    assertTrue(bucket1.capacity != bucket1.tokens);
  }

  function test_UpdateExistingConfigWithNoDifference_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });

    RateLimiterNoEvents.TokenBucket memory bucketPreUpdate =
      s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterConfigUpdated(
      configUpdates[0].chainSelector, configUpdates[0].rateLimiterConfig
    );

    vm.recordLogs();
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    vm.warp(BLOCK_TIME + 1);
    RateLimiterNoEvents.TokenBucket memory bucketPostUpdate =
      s_rateLimiter.currentRateLimiterState(configUpdates[0].chainSelector);
    _assertTokenBucketEquality(bucketPreUpdate, bucketPostUpdate);
    assertEq(BLOCK_TIME + 1, bucketPostUpdate.lastUpdated);
  }

  // Reverts
  function test_ZeroChainSelector_Revert() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] =
      MultiAggregateRateLimiter.RateLimiterConfigArgs({chainSelector: 0, rateLimiterConfig: RATE_LIMITER_CONFIG_1});

    vm.expectRevert(MultiAggregateRateLimiter.ZeroChainSelectorNotAllowed.selector);
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);
  }

  function test_OnlyCallableByAdminOrOwner_Revert() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1 + 1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });
    vm.startPrank(STRANGER);

    vm.expectRevert(RateLimiterNoEvents.OnlyCallableByAdminOrOwner.selector);
    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);
  }
}

// /// @notice #_rateLimitValue
contract MultiAggregateRateLimiter__rateLimitValue is MultiAggregateRateLimiterSetup {
  function test_RateLimitValue_Success_gas() public {
    vm.pauseGasMetering();
    // start from blocktime that does not equal rate limiter init timestamp
    vm.warp(BLOCK_TIME + 1);

    // 15 (tokens) * 4 (price) * 2 (number of times) > 100 (capacity)
    uint256 numberOfTokens = 15;
    uint256 value = (numberOfTokens * TOKEN_PRICE) / 1e18;

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_1, value);

    vm.resumeGasMetering();
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_1, value);
    vm.pauseGasMetering();

    // Get the updated bucket status
    RateLimiterNoEvents.TokenBucket memory bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    // Assert the proper value has been taken out of the bucket
    assertEq(bucket.capacity - value, bucket.tokens);

    // Since value * 2 > bucket.capacity we cannot take it out twice.
    // Expect a revert when we try, with a wait time.
    uint256 waitTime = 4;
    vm.expectRevert(
      abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueRateLimitReached.selector, waitTime, bucket.tokens)
    );
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_1, value);

    // Move the block time forward by 10 so the bucket refills by 10 * rate
    vm.warp(BLOCK_TIME + 1 + waitTime);

    // The bucket has filled up enough so we can take out more tokens
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_1, value);
    bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    assertEq(bucket.capacity - value + waitTime * RATE_LIMITER_CONFIG_1.rate - value, bucket.tokens);
    vm.resumeGasMetering();
  }

  function test_RateLimitValueDifferentChainSelectors_Success() public {
    vm.pauseGasMetering();
    // start from blocktime that does not equal rate limiter init timestamp
    vm.warp(BLOCK_TIME + 1);

    // 15 (tokens) * 4 (price) * 2 (number of times) > 100 (capacity)
    uint256 numberOfTokens = 15;
    uint256 value = (numberOfTokens * TOKEN_PRICE) / 1e18;

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_1, value);

    vm.resumeGasMetering();
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_1, value);
    vm.pauseGasMetering();

    // Get the updated bucket status
    RateLimiterNoEvents.TokenBucket memory bucket1 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    RateLimiterNoEvents.TokenBucket memory bucket2 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_2);

    // Assert the proper value has been taken out of the bucket
    assertEq(bucket1.capacity - value, bucket1.tokens);
    // CHAIN_SELECTOR_2 should remain unchanged
    assertEq(bucket2.capacity, bucket2.tokens);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_2, value);

    vm.resumeGasMetering();
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_2, value);
    vm.pauseGasMetering();

    bucket1 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    bucket2 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_2);

    assertEq(bucket2.capacity - value, bucket2.tokens);
    // CHAIN_SELECTOR_1 should remain unchanged
    assertEq(bucket1.capacity - value, bucket1.tokens);
  }

  // Reverts

  function test_AggregateValueMaxCapacityExceeded_Revert() public {
    RateLimiterNoEvents.TokenBucket memory bucket = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);

    uint256 numberOfTokens = 100;
    uint256 value = (numberOfTokens * TOKEN_PRICE) / 1e18;

    vm.expectRevert(
      abi.encodeWithSelector(
        RateLimiterNoEvents.AggregateValueMaxCapacityExceeded.selector,
        bucket.capacity,
        (numberOfTokens * TOKEN_PRICE) / 1e18
      )
    );
    s_rateLimiter.rateLimitValue(CHAIN_SELECTOR_1, value);
  }
}

/// @notice #_getTokenValue
contract MultiAggregateRateLimiter__getTokenValue is MultiAggregateRateLimiterSetup {
  function test_GetTokenValue_Success() public view {
    uint256 numberOfTokens = 10;
    Client.EVMTokenAmount memory tokenAmount = Client.EVMTokenAmount({token: TOKEN, amount: 10});
    uint256 value = s_rateLimiter.getTokenValue(tokenAmount);
    assertEq(value, (numberOfTokens * TOKEN_PRICE) / 1e18);
  }

  // Reverts
  function test_NoTokenPrice_Reverts() public {
    address tokenWithNoPrice = makeAddr("Token with no price");
    Client.EVMTokenAmount memory tokenAmount = Client.EVMTokenAmount({token: tokenWithNoPrice, amount: 10});

    vm.expectRevert(abi.encodeWithSelector(MultiAggregateRateLimiter.PriceNotFoundForToken.selector, tokenWithNoPrice));
    s_rateLimiter.getTokenValue(tokenAmount);
  }
}

/// @notice #updateRateLimitTokens
contract MultiAggregateRateLimiter_updateRateLimitTokens is MultiAggregateRateLimiterSetup {
  function setUp() public virtual override {
    super.setUp();

    // Clear rate limit tokens state
    MultiAggregateRateLimiter.RateLimitToken[] memory remove =
      new MultiAggregateRateLimiter.RateLimitToken[](s_sourceTokens.length);
    for (uint256 i = 0; i < s_sourceTokens.length; ++i) {
      remove[i] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[i], destToken: s_destTokens[i]});
    }
    s_rateLimiter.updateRateLimitTokens(remove, new MultiAggregateRateLimiter.RateLimitToken[](0));
  }

  function test_UpdateRateLimitTokens_Success() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory adds = new MultiAggregateRateLimiter.RateLimitToken[](2);
    adds[0] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[0], destToken: s_destTokens[0]});
    adds[1] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[1], destToken: s_destTokens[1]});

    for (uint256 i = 0; i < adds.length; ++i) {
      vm.expectEmit();
      emit MultiAggregateRateLimiter.TokenAggregateRateLimitAdded(adds[i].sourceToken, adds[i].destToken);
    }

    s_rateLimiter.updateRateLimitTokens(new MultiAggregateRateLimiter.RateLimitToken[](0), adds);

    (address[] memory sourceTokens, address[] memory destTokens) = s_rateLimiter.getAllRateLimitTokens();

    for (uint256 i = 0; i < adds.length; ++i) {
      assertEq(adds[i].sourceToken, sourceTokens[i]);
      assertEq(adds[i].destToken, destTokens[i]);
    }
  }

  function test_UpdateRateLimitTokens_AddsAndRemoves_Success() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory adds = new MultiAggregateRateLimiter.RateLimitToken[](2);
    adds[0] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[0], destToken: s_destTokens[0]});
    adds[1] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[1], destToken: s_destTokens[1]});

    MultiAggregateRateLimiter.RateLimitToken[] memory removes = new MultiAggregateRateLimiter.RateLimitToken[](1);
    removes[0] = adds[0];

    for (uint256 i = 0; i < adds.length; ++i) {
      vm.expectEmit();
      emit MultiAggregateRateLimiter.TokenAggregateRateLimitAdded(adds[i].sourceToken, adds[i].destToken);
    }

    s_rateLimiter.updateRateLimitTokens(removes, adds);

    for (uint256 i = 0; i < removes.length; ++i) {
      vm.expectEmit();
      emit MultiAggregateRateLimiter.TokenAggregateRateLimitRemoved(removes[i].sourceToken, removes[i].destToken);
    }

    s_rateLimiter.updateRateLimitTokens(removes, new MultiAggregateRateLimiter.RateLimitToken[](0));

    (address[] memory sourceTokens, address[] memory destTokens) = s_rateLimiter.getAllRateLimitTokens();

    assertEq(1, sourceTokens.length);
    assertEq(adds[1].sourceToken, sourceTokens[0]);

    assertEq(1, destTokens.length);
    assertEq(adds[1].destToken, destTokens[0]);
  }

  function test_UpdateRateLimitTokensByAdmin_Success() public {
    vm.startPrank(ADMIN);

    MultiAggregateRateLimiter.RateLimitToken[] memory adds = new MultiAggregateRateLimiter.RateLimitToken[](2);
    adds[0] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[0], destToken: s_destTokens[0]});
    adds[1] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[1], destToken: s_destTokens[1]});

    for (uint256 i = 0; i < adds.length; ++i) {
      vm.expectEmit();
      emit MultiAggregateRateLimiter.TokenAggregateRateLimitAdded(adds[i].sourceToken, adds[i].destToken);
    }

    s_rateLimiter.updateRateLimitTokens(new MultiAggregateRateLimiter.RateLimitToken[](0), adds);

    (address[] memory sourceTokens, address[] memory destTokens) = s_rateLimiter.getAllRateLimitTokens();

    for (uint256 i = 0; i < adds.length; ++i) {
      assertEq(adds[i].sourceToken, sourceTokens[i]);
      assertEq(adds[i].destToken, destTokens[i]);
    }
  }

  // Reverts

  function test_ZeroSourceToken_Revert() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory adds = new MultiAggregateRateLimiter.RateLimitToken[](1);
    adds[0] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: address(0), destToken: s_destTokens[0]});

    vm.expectRevert(MultiAggregateRateLimiter.ZeroAddressNotAllowed.selector);
    s_rateLimiter.updateRateLimitTokens(new MultiAggregateRateLimiter.RateLimitToken[](0), adds);
  }

  function test_ZeroDestToken_Revert() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory adds = new MultiAggregateRateLimiter.RateLimitToken[](1);
    adds[0] = MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_destTokens[0], destToken: address(0)});

    vm.expectRevert(MultiAggregateRateLimiter.ZeroAddressNotAllowed.selector);
    s_rateLimiter.updateRateLimitTokens(new MultiAggregateRateLimiter.RateLimitToken[](0), adds);
  }

  function test_NonOwner_Revert() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory addsAndRemoves = new MultiAggregateRateLimiter.RateLimitToken[](4);

    vm.startPrank(STRANGER);

    vm.expectRevert(RateLimiterNoEvents.OnlyCallableByAdminOrOwner.selector);
    s_rateLimiter.updateRateLimitTokens(addsAndRemoves, addsAndRemoves);
  }
}

/// @notice #validateIncomingMessage
contract MultiAggregateRateLimiter_validateIncomingMessage is MultiAggregateRateLimiterSetup {
  address internal immutable MOCK_RECEIVER = address(1113);

  function setUp() public virtual override {
    super.setUp();

    MultiAggregateRateLimiter.RateLimitToken[] memory tokensToAdd =
      new MultiAggregateRateLimiter.RateLimitToken[](s_sourceTokens.length);
    for (uint224 i = 0; i < s_sourceTokens.length; ++i) {
      tokensToAdd[i] =
        MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[i], destToken: s_destTokens[i]});

      Internal.PriceUpdates memory priceUpdates =
        getSingleTokenPriceUpdateStruct(s_destTokens[i], TOKEN_PRICE * (i + 1));
      s_priceRegistry.updatePrices(priceUpdates);
    }
    s_rateLimiter.updateRateLimitTokens(new MultiAggregateRateLimiter.RateLimitToken[](0), tokensToAdd);
  }

  function test_ValidateMessageWithNoTokens_Success() public {
    vm.startPrank(MOCK_OFFRAMP);

    vm.recordLogs();
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessageNoTokens(CHAIN_SELECTOR_1));

    // No consumed rate limit events
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);
  }

  function test_ValidateMessageWithTokens_Success() public {
    vm.startPrank(MOCK_OFFRAMP);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 3});
    tokenAmounts[1] = Client.EVMTokenAmount({token: s_destTokens[1], amount: 1});

    // 3 tokens * TOKEN_PRICE + 1 token * (2 * TOKEN_PRICE)
    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_1, (5 * TOKEN_PRICE) / 1e18);

    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));
  }

  function test_ValidateMessageWithDisabledRateLimitToken_Success() public {
    MultiAggregateRateLimiter.RateLimitToken[] memory tokensToRemove = new MultiAggregateRateLimiter.RateLimitToken[](1);
    tokensToRemove[0] =
      MultiAggregateRateLimiter.RateLimitToken({sourceToken: s_sourceTokens[1], destToken: s_destTokens[1]});
    s_rateLimiter.updateRateLimitTokens(tokensToRemove, new MultiAggregateRateLimiter.RateLimitToken[](0));

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 5});
    tokenAmounts[1] = Client.EVMTokenAmount({token: s_destTokens[1], amount: 1});

    vm.startPrank(MOCK_OFFRAMP);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_1, (5 * TOKEN_PRICE) / 1e18);

    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));
  }

  function test_ValidateMessageWithRateLimitDisabled_Success() public {
    MultiAggregateRateLimiter.RateLimiterConfigArgs[] memory configUpdates =
      new MultiAggregateRateLimiter.RateLimiterConfigArgs[](1);
    configUpdates[0] = MultiAggregateRateLimiter.RateLimiterConfigArgs({
      chainSelector: CHAIN_SELECTOR_1,
      rateLimiterConfig: RATE_LIMITER_CONFIG_1
    });
    configUpdates[0].rateLimiterConfig.isEnabled = false;

    s_rateLimiter.applyRateLimiterConfigUpdates(configUpdates);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 1000});
    tokenAmounts[1] = Client.EVMTokenAmount({token: s_destTokens[1], amount: 50});

    vm.startPrank(MOCK_OFFRAMP);
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));

    // No consumed rate limit events
    Vm.Log[] memory logEntries = vm.getRecordedLogs();
    assertEq(logEntries.length, 0);
  }

  function test_ValidateMessageWithTokensOnDifferentChains_Success() public {
    vm.startPrank(MOCK_OFFRAMP);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 2});
    tokenAmounts[1] = Client.EVMTokenAmount({token: s_destTokens[1], amount: 1});

    // 2 tokens * (TOKEN_PRICE) + 1 token * (2 * TOKEN_PRICE)
    uint256 totalValue = (4 * TOKEN_PRICE) / 1e18;

    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));

    // Chain 1 changed
    RateLimiterNoEvents.TokenBucket memory bucketChain1 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    assertEq(bucketChain1.capacity - totalValue, bucketChain1.tokens);

    // Chain 2 unchanged
    RateLimiterNoEvents.TokenBucket memory bucketChain2 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_2);
    assertEq(bucketChain2.capacity, bucketChain2.tokens);

    vm.expectEmit();
    emit MultiAggregateRateLimiter.RateLimiterTokensConsumed(CHAIN_SELECTOR_2, (4 * TOKEN_PRICE) / 1e18);

    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_2, tokenAmounts));

    // Chain 1 unchanged
    bucketChain1 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_1);
    assertEq(bucketChain1.capacity - totalValue, bucketChain1.tokens);

    // Chain 2 changed
    bucketChain2 = s_rateLimiter.currentRateLimiterState(CHAIN_SELECTOR_2);
    assertEq(bucketChain2.capacity - totalValue, bucketChain2.tokens);
  }

  function test_ValidateMessageWithRateLimitReset_Success() public {
    vm.startPrank(MOCK_OFFRAMP);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 20});

    // Remaining capacity: 100 -> 20
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));

    // Cannot fit 80 rate limit value (need to wait at least 12 blocks, current capacity is 20)
    vm.expectRevert(abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueRateLimitReached.selector, 12, 20));
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));

    // Remaining capacity: 20 -> 35 (need to wait 9 more blocks)
    vm.warp(BLOCK_TIME + 3);
    vm.expectRevert(abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueRateLimitReached.selector, 9, 35));
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));

    // Remaining capacity: 35 -> 80 (can fit exactly 80)
    vm.warp(BLOCK_TIME + 12);
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));
  }

  // Reverts

  function test_ValidateMessageWithRateLimitExceeded_Revert() public {
    vm.startPrank(MOCK_OFFRAMP);

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_destTokens[0], amount: 80});
    tokenAmounts[1] = Client.EVMTokenAmount({token: s_destTokens[1], amount: 30});

    uint256 totalValue = (80 * TOKEN_PRICE + 2 * (30 * TOKEN_PRICE)) / 1e18;
    vm.expectRevert(
      abi.encodeWithSelector(RateLimiterNoEvents.AggregateValueMaxCapacityExceeded.selector, 100, totalValue)
    );
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessage(CHAIN_SELECTOR_1, tokenAmounts));
  }

  function test_ValidateMessageFromUnauthorizedCaller_Revert() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(MultiAggregateRateLimiter.UnauthorizedCaller.selector, STRANGER));
    s_rateLimiter.validateIncomingMessage(_generateAny2EVMMessageNoTokens(CHAIN_SELECTOR_1));
  }

  function _generateAny2EVMMessageNoTokens(uint64 sourceChainSelector)
    internal
    pure
    returns (Client.Any2EVMMessage memory)
  {
    return _generateAny2EVMMessage(sourceChainSelector, new Client.EVMTokenAmount[](0));
  }

  function _generateAny2EVMMessage(
    uint64 sourceChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts
  ) internal pure returns (Client.Any2EVMMessage memory) {
    return Client.Any2EVMMessage({
      messageId: keccak256(bytes("messageId")),
      sourceChainSelector: sourceChainSelector,
      sender: abi.encode(OWNER),
      data: abi.encode(0),
      destTokenAmounts: tokenAmounts
    });
  }
}
