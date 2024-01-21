// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {TokenPoolHelper} from "../helpers/TokenPoolHelper.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BurnMintERC677} from "../../../shared/token/ERC677/BurnMintERC677.sol";
import {RouterSetup} from "../router/RouterSetup.t.sol";
import {Router} from "../../Router.sol";

contract TokenPoolSetup is RouterSetup {
  IERC20 internal s_token;
  TokenPoolHelper internal s_tokenPool;

  function setUp() public virtual override {
    RouterSetup.setUp();
    s_token = new BurnMintERC677("LINK", "LNK", 18, 0);
    deal(address(s_token), OWNER, type(uint256).max);

    s_tokenPool = new TokenPoolHelper(s_token, new address[](0), address(s_mockARM), address(s_sourceRouter));
  }
}

contract TokenPool_constructor is TokenPoolSetup {
  // Reverts
  function testZeroAddressNotAllowedReverts() public {
    vm.expectRevert(TokenPool.ZeroAddressNotAllowed.selector);

    s_tokenPool = new TokenPoolHelper(
      IERC20(address(0)),
      new address[](0),
      address(s_mockARM),
      address(s_sourceRouter)
    );
  }
}

contract TokenPool_applyRampUpdates is TokenPoolSetup {
  event ChainAdded(uint64 chainSelector, RateLimiter.Config rateLimiterConfig);
  event ChainRemoved(uint64 chainSelector);

  function chainSelectorsFromUpdates(TokenPool.ChainUpdate[] memory updates) public pure returns (uint64[] memory) {
    uint64[] memory chainSelectors = new uint64[](updates.length);
    for (uint256 i = 0; i < updates.length; i++) {
      chainSelectors[i] = updates[i].chainSelector;
    }
    return chainSelectors;
  }

  function assertState(TokenPool.ChainUpdate[] memory chainUpdates) public {
    //assertEq(s_tokenPool.getSupportedChains(), chainSelectorsFromUpdates(chainUpdates));
    for (uint256 i = 0; i < chainUpdates.length; ++i) {
      assertTrue(s_tokenPool.isSupportedChain(chainUpdates[i].chainSelector));
      RateLimiter.TokenBucket memory bkt = s_tokenPool.currentOutboundRateLimiterState(chainUpdates[i].chainSelector);
      assertEq(bkt.capacity, chainUpdates[i].rateLimiterConfig.capacity);
      assertEq(bkt.rate, chainUpdates[i].rateLimiterConfig.rate);
      assertEq(bkt.isEnabled, chainUpdates[i].rateLimiterConfig.isEnabled);

      bkt = s_tokenPool.currentInboundRateLimiterState(chainUpdates[i].chainSelector);
      assertEq(bkt.capacity, chainUpdates[i].rateLimiterConfig.capacity);
      assertEq(bkt.rate, chainUpdates[i].rateLimiterConfig.rate);
      assertEq(bkt.isEnabled, chainUpdates[i].rateLimiterConfig.isEnabled);
    }
  }

  function testApplyRampUpdatesSuccess() public {
    // Create on and offramps.
    RateLimiter.Config memory rateLimit1 = RateLimiter.Config({isEnabled: true, capacity: 100e28, rate: 1e15});
    RateLimiter.Config memory rateLimit2 = RateLimiter.Config({isEnabled: true, capacity: 100e27, rate: 1e14});
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](2);
    chainUpdates[0] = TokenPool.ChainUpdate({chainSelector: 1, allowed: true, rateLimiterConfig: rateLimit1});
    chainUpdates[1] = TokenPool.ChainUpdate({chainSelector: 2, allowed: true, rateLimiterConfig: rateLimit2});

    // Assert configuration is applied
    vm.expectEmit();
    emit ChainAdded(chainUpdates[0].chainSelector, chainUpdates[0].rateLimiterConfig);
    vm.expectEmit();
    emit ChainAdded(chainUpdates[1].chainSelector, chainUpdates[1].rateLimiterConfig);
    s_tokenPool.applyChainUpdates(chainUpdates);
    // on1: rateLimit1, on2: rateLimit2, off1: rateLimit1, off2: rateLimit3
    assertState(chainUpdates);

    // Removing an non-existent chain should revert
    TokenPool.ChainUpdate[] memory chainRemoves = new TokenPool.ChainUpdate[](1);
    uint64 strangerChainSelector = 120938;
    chainRemoves[0] = TokenPool.ChainUpdate({
      chainSelector: strangerChainSelector,
      allowed: false,
      rateLimiterConfig: rateLimit1
    });
    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, strangerChainSelector));
    s_tokenPool.applyChainUpdates(chainRemoves);
    // State remains
    assertState(chainUpdates);

    // Can remove a chain
    chainRemoves[0].chainSelector = 1;

    vm.expectEmit();
    emit ChainRemoved(chainRemoves[0].chainSelector);

    s_tokenPool.applyChainUpdates(chainRemoves);

    // State updated, only chain 2 remains
    TokenPool.ChainUpdate[] memory singleChainConfigured = new TokenPool.ChainUpdate[](1);
    singleChainConfigured[0] = chainUpdates[1];
    assertState(singleChainConfigured);

    // Cannot reset already configured ramp
    vm.expectRevert(
      abi.encodeWithSelector(TokenPool.ChainAlreadyExists.selector, singleChainConfigured[0].chainSelector)
    );
    s_tokenPool.applyChainUpdates(singleChainConfigured);
  }

  // Reverts

  function testOnlyCallableByOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_tokenPool.applyChainUpdates(new TokenPool.ChainUpdate[](0));
  }
}

contract TokenPool_setOnRampRateLimiterConfig is TokenPoolSetup {
  event ConfigChanged(RateLimiter.Config);
  event ChainConfigured(uint64 chainSelector, RateLimiter.Config);

  uint64 internal s_remoteChainSelector;

  function setUp() public virtual override {
    TokenPoolSetup.setUp();
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](1);
    s_remoteChainSelector = 123124;
    chainUpdates[0] = TokenPool.ChainUpdate({
      chainSelector: s_remoteChainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdates);
  }

  function testFuzz_SetRateLimiterConfigSuccess(uint128 capacity, uint128 rate, uint32 newTime) public {
    // Bucket updates only work on increasing time
    vm.assume(newTime >= block.timestamp);
    vm.warp(newTime);

    uint256 oldTokens = s_tokenPool.currentOutboundRateLimiterState(s_remoteChainSelector).tokens;

    RateLimiter.Config memory newConfig = RateLimiter.Config({isEnabled: true, capacity: capacity, rate: rate});

    vm.expectEmit();
    emit ConfigChanged(newConfig);
    vm.expectEmit();
    emit ChainConfigured(s_remoteChainSelector, newConfig);

    s_tokenPool.setChainRateLimiterConfig(s_remoteChainSelector, newConfig);

    uint256 expectedTokens = RateLimiter._min(newConfig.capacity, oldTokens);

    RateLimiter.TokenBucket memory bucket = s_tokenPool.currentOutboundRateLimiterState(s_remoteChainSelector);
    assertEq(bucket.capacity, newConfig.capacity);
    assertEq(bucket.rate, newConfig.rate);
    assertEq(bucket.tokens, expectedTokens);
    assertEq(bucket.lastUpdated, newTime);
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    s_tokenPool.setChainRateLimiterConfig(s_remoteChainSelector, rateLimiterConfig());
  }

  function testNonExistentRampReverts() public {
    uint64 wrongChainSelector = 9084102894;

    vm.expectRevert(abi.encodeWithSelector(TokenPool.NonExistentChain.selector, wrongChainSelector));
    s_tokenPool.setChainRateLimiterConfig(wrongChainSelector, rateLimiterConfig());
  }
}

contract TokenPool_onlyOnRamp is TokenPoolSetup {
  function test_onlyOnRampSuccess() public {
    uint64 chainSelector = 13377;
    address onRamp = makeAddr("onRamp");

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: chainSelector, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    vm.startPrank(onRamp);

    s_tokenPool.onlyOnRampModifier(chainSelector);
  }

  function test_ChainNotAllowedReverts() public {
    uint64 chainSelector = 13377;
    address onRamp = makeAddr("onRamp");

    vm.startPrank(onRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, chainSelector));
    s_tokenPool.onlyOnRampModifier(chainSelector);

    vm.startPrank(OWNER);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    Router.OnRamp[] memory onRampUpdates = new Router.OnRamp[](1);
    onRampUpdates[0] = Router.OnRamp({destChainSelector: chainSelector, onRamp: onRamp});
    s_sourceRouter.applyRampUpdates(onRampUpdates, new Router.OffRamp[](0), new Router.OffRamp[](0));

    vm.startPrank(onRamp);
    // Should succeed now that we've added the chain
    s_tokenPool.onlyOnRampModifier(chainSelector);

    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: false,
      rateLimiterConfig: rateLimiterConfig()
    });

    vm.startPrank(OWNER);
    s_tokenPool.applyChainUpdates(chainUpdate);

    vm.startPrank(onRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, chainSelector));
    s_tokenPool.onlyOffRampModifier(chainSelector);
  }

  function test_CallerIsNotARampOnRouterReverts() public {
    uint64 chainSelector = 13377;
    address onRamp = makeAddr("onRamp");

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    vm.startPrank(onRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotARampOnRouter.selector, onRamp));

    s_tokenPool.onlyOnRampModifier(chainSelector);
  }
}

contract TokenPool_onlyOffRamp is TokenPoolSetup {
  function test_onlyOffRampSuccess() public {
    uint64 chainSelector = 13377;
    address offRamp = makeAddr("onRamp");

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: chainSelector, offRamp: offRamp});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    vm.startPrank(offRamp);

    s_tokenPool.onlyOffRampModifier(chainSelector);
  }

  function test_ChainNotAllowedReverts() public {
    uint64 chainSelector = 13377;
    address offRamp = makeAddr("onRamp");

    vm.startPrank(offRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, chainSelector));
    s_tokenPool.onlyOffRampModifier(chainSelector);

    vm.startPrank(OWNER);

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    Router.OffRamp[] memory offRampUpdates = new Router.OffRamp[](1);
    offRampUpdates[0] = Router.OffRamp({sourceChainSelector: chainSelector, offRamp: offRamp});
    s_sourceRouter.applyRampUpdates(new Router.OnRamp[](0), new Router.OffRamp[](0), offRampUpdates);

    vm.startPrank(offRamp);
    // Should succeed now that we've added the chain
    s_tokenPool.onlyOffRampModifier(chainSelector);

    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: false,
      rateLimiterConfig: rateLimiterConfig()
    });

    vm.startPrank(OWNER);
    s_tokenPool.applyChainUpdates(chainUpdate);

    vm.startPrank(offRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, chainSelector));
    s_tokenPool.onlyOffRampModifier(chainSelector);
  }

  function test_CallerIsNotARampOnRouterReverts() public {
    uint64 chainSelector = 13377;
    address offRamp = makeAddr("offRamp");

    TokenPool.ChainUpdate[] memory chainUpdate = new TokenPool.ChainUpdate[](1);
    chainUpdate[0] = TokenPool.ChainUpdate({
      chainSelector: chainSelector,
      allowed: true,
      rateLimiterConfig: rateLimiterConfig()
    });
    s_tokenPool.applyChainUpdates(chainUpdate);

    vm.startPrank(offRamp);

    vm.expectRevert(abi.encodeWithSelector(TokenPool.CallerIsNotARampOnRouter.selector, offRamp));

    s_tokenPool.onlyOffRampModifier(chainSelector);
  }
}

contract TokenPoolWithAllowListSetup is TokenPoolSetup {
  address[] internal s_allowedSenders;

  function setUp() public virtual override {
    TokenPoolSetup.setUp();

    s_allowedSenders.push(STRANGER);
    s_allowedSenders.push(DUMMY_CONTRACT_ADDRESS);

    s_tokenPool = new TokenPoolHelper(s_token, s_allowedSenders, address(s_mockARM), address(s_sourceRouter));
  }
}

/// @notice #getAllowListEnabled
contract TokenPoolWithAllowList_getAllowListEnabled is TokenPoolWithAllowListSetup {
  function testGetAllowListEnabledSuccess() public {
    assertTrue(s_tokenPool.getAllowListEnabled());
  }
}

/// @notice #getAllowList
contract TokenPoolWithAllowList_getAllowList is TokenPoolWithAllowListSetup {
  function testGetAllowListSuccess() public {
    address[] memory setAddresses = s_tokenPool.getAllowList();
    assertEq(2, setAddresses.length);
    assertEq(s_allowedSenders[0], setAddresses[0]);
    assertEq(s_allowedSenders[1], setAddresses[1]);
  }
}

/// @notice #setAllowList
contract TokenPoolWithAllowList_applyAllowListUpdates is TokenPoolWithAllowListSetup {
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);

  function testSetAllowListSuccess() public {
    address[] memory newAddresses = new address[](2);
    newAddresses[0] = address(1);
    newAddresses[1] = address(2);

    for (uint256 i = 0; i < 2; ++i) {
      vm.expectEmit();
      emit AllowListAdd(newAddresses[i]);
    }

    s_tokenPool.applyAllowListUpdates(new address[](0), newAddresses);
    address[] memory setAddresses = s_tokenPool.getAllowList();

    assertEq(s_allowedSenders[0], setAddresses[0]);
    assertEq(s_allowedSenders[1], setAddresses[1]);
    assertEq(address(1), setAddresses[2]);
    assertEq(address(2), setAddresses[3]);

    // address(2) exists noop, add address(3), remove address(1)
    newAddresses = new address[](2);
    newAddresses[0] = address(2);
    newAddresses[1] = address(3);

    address[] memory removeAddresses = new address[](1);
    removeAddresses[0] = address(1);

    vm.expectEmit();
    emit AllowListRemove(address(1));

    vm.expectEmit();
    emit AllowListAdd(address(3));

    s_tokenPool.applyAllowListUpdates(removeAddresses, newAddresses);
    setAddresses = s_tokenPool.getAllowList();

    assertEq(s_allowedSenders[0], setAddresses[0]);
    assertEq(s_allowedSenders[1], setAddresses[1]);
    assertEq(address(2), setAddresses[2]);
    assertEq(address(3), setAddresses[3]);

    // remove all from allowList
    for (uint256 i = 0; i < setAddresses.length; ++i) {
      vm.expectEmit();
      emit AllowListRemove(setAddresses[i]);
    }

    s_tokenPool.applyAllowListUpdates(setAddresses, new address[](0));
    setAddresses = s_tokenPool.getAllowList();

    assertEq(0, setAddresses.length);
  }

  function testSetAllowListSkipsZeroSuccess() public {
    uint256 setAddressesLength = s_tokenPool.getAllowList().length;

    address[] memory newAddresses = new address[](1);
    newAddresses[0] = address(0);

    s_tokenPool.applyAllowListUpdates(new address[](0), newAddresses);
    address[] memory setAddresses = s_tokenPool.getAllowList();

    assertEq(setAddresses.length, setAddressesLength);
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    address[] memory newAddresses = new address[](2);
    s_tokenPool.applyAllowListUpdates(new address[](0), newAddresses);
  }
}
