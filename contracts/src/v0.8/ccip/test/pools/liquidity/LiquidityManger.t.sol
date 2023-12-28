// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "forge-std/Test.sol";

import {IBridgeAdapter, IL1BridgeAdapter} from "../../../pools/liquidity/interfaces/IBridge.sol";
import {ILiquidityContainer} from "../../../pools/liquidity/interfaces/ILiquidityContainer.sol";

import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {LiquidityManager} from "../../../pools/liquidity/LiquidityManager.sol";
import {MockL1BridgeAdapter} from "../../mocks/MockBridgeAdapter.sol";

import {ERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract LiquidityManagerSetup is Test {
  event Transfer(address indexed from, address indexed to, uint256 value);
  event Approval(address indexed owner, address indexed spender, uint256 value);
  event LiquidityTransferred(
    uint64 indexed fromChainSelector,
    uint64 indexed toChainSelector,
    address indexed to,
    uint256 amount
  );

  IERC20 public s_token;
  LiquidityManager internal s_liquidityManager;
  LockReleaseTokenPool internal s_lockReleaseTokenPool;
  MockL1BridgeAdapter internal s_bridgeAdapter;
  uint64 internal immutable i_localChainSelector = 1234;

  address internal immutable i_l2Token = address(0x22222222222222222222222222);

  function setUp() external {
    s_token = new ERC20("Test", "TEST");

    s_bridgeAdapter = new MockL1BridgeAdapter(s_token);
    s_lockReleaseTokenPool = new LockReleaseTokenPool(s_token, new address[](0), address(1), true);
    s_liquidityManager = new LiquidityManager(s_token, i_localChainSelector, s_lockReleaseTokenPool);

    s_lockReleaseTokenPool.setLiquidityManager(address(s_liquidityManager));
  }
}

contract LiquidityManager_rebalanceLiquidity is LiquidityManagerSetup {
  function test_rebalanceLiquiditySuccess() external {
    uint256 amount = 12345679;
    uint64 remoteChainSelector = 123;
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);

    LiquidityManager.CrossChainLiquidityManagerArgs[]
      memory args = new LiquidityManager.CrossChainLiquidityManagerArgs[](1);
    args[0] = LiquidityManager.CrossChainLiquidityManagerArgs({
      remoteLiquidityManager: address(s_liquidityManager),
      localBridge: s_bridgeAdapter,
      remoteToken: i_l2Token,
      remoteChainSelector: remoteChainSelector,
      enabled: true
    });
    s_liquidityManager.setCrossChainLiquidityManager(args);

    vm.expectEmit();
    emit Transfer(address(s_lockReleaseTokenPool), address(s_liquidityManager), amount);

    vm.expectEmit();
    emit Approval(address(s_liquidityManager), address(s_bridgeAdapter), amount);

    vm.expectEmit();
    emit Transfer(address(s_liquidityManager), address(s_bridgeAdapter), amount);

    vm.expectEmit();
    emit LiquidityTransferred(i_localChainSelector, remoteChainSelector, address(s_liquidityManager), amount);

    s_liquidityManager.rebalanceLiquidity(remoteChainSelector, amount);

    assertEq(s_token.balanceOf(address(s_liquidityManager)), 0);
    assertEq(s_token.balanceOf(address(s_bridgeAdapter)), amount);
    assertEq(s_token.allowance(address(s_liquidityManager), address(s_bridgeAdapter)), 0);
  }

  /// @notice this test sets up a circular system where the liquidity container of
  /// the local Liquidity manager is the bridge adapter of the remote liquidity manager
  /// and the other way around for the remote liquidity manager. This allows us to
  /// rebalance funds between the two liquidity managers on the same chain.
  function test_rebalanceBetweenPoolsSuccess() external {
    uint256 amount = 12345670;
    uint64 remoteChainSelector = 123;

    s_liquidityManager = new LiquidityManager(s_token, i_localChainSelector, s_bridgeAdapter);

    MockL1BridgeAdapter mockRemoteBridgeAdapter = new MockL1BridgeAdapter(s_token);
    LiquidityManager mockRemoteLiquidityManager = new LiquidityManager(
      s_token,
      remoteChainSelector,
      mockRemoteBridgeAdapter
    );

    LiquidityManager.CrossChainLiquidityManagerArgs[]
      memory args = new LiquidityManager.CrossChainLiquidityManagerArgs[](1);
    args[0] = LiquidityManager.CrossChainLiquidityManagerArgs({
      remoteLiquidityManager: address(mockRemoteLiquidityManager),
      localBridge: mockRemoteBridgeAdapter,
      remoteToken: address(s_token),
      remoteChainSelector: remoteChainSelector,
      enabled: true
    });

    s_liquidityManager.setCrossChainLiquidityManager(args);

    args[0] = LiquidityManager.CrossChainLiquidityManagerArgs({
      remoteLiquidityManager: address(s_liquidityManager),
      localBridge: s_bridgeAdapter,
      remoteToken: address(s_token),
      remoteChainSelector: i_localChainSelector,
      enabled: true
    });

    mockRemoteLiquidityManager.setCrossChainLiquidityManager(args);

    deal(address(s_token), address(s_bridgeAdapter), amount);

    s_liquidityManager.rebalanceLiquidity(remoteChainSelector, amount);

    assertEq(s_token.balanceOf(address(s_bridgeAdapter)), 0);
    assertEq(s_token.balanceOf(address(mockRemoteBridgeAdapter)), amount);
    assertEq(s_token.allowance(address(s_liquidityManager), address(s_bridgeAdapter)), 0);

    mockRemoteLiquidityManager.rebalanceLiquidity(i_localChainSelector, amount);

    assertEq(s_token.balanceOf(address(s_bridgeAdapter)), amount);
    assertEq(s_token.balanceOf(address(mockRemoteBridgeAdapter)), 0);

    // Assert partial rebalancing works correctly
    s_liquidityManager.rebalanceLiquidity(remoteChainSelector, amount / 2);

    assertEq(s_token.balanceOf(address(s_bridgeAdapter)), amount / 2);
    assertEq(s_token.balanceOf(address(mockRemoteBridgeAdapter)), amount / 2);
  }

  // Reverts

  function test_InsufficientLiquidityReverts() external {
    uint256 amount = 1245;

    vm.expectRevert(abi.encodeWithSelector(LiquidityManager.InsufficientLiquidity.selector, amount, 0));

    s_liquidityManager.rebalanceLiquidity(0, amount);
  }

  function test_InvalidRemoteChainReverts() external {
    uint256 amount = 12345679;
    uint64 remoteChainSelector = 123;
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);

    vm.expectRevert(abi.encodeWithSelector(LiquidityManager.InvalidRemoteChain.selector, remoteChainSelector));

    s_liquidityManager.rebalanceLiquidity(remoteChainSelector, amount);
  }
}
