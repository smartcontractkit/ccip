// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import "forge-std/Test.sol";

import {IBridge} from "../../../pools/liquidity/interfaces/IBridge.sol";
import {ILiquidityContainer} from "../../../pools/liquidity/interfaces/ILiquidityContainer.sol";

import {LockReleaseTokenPool} from "../../../pools/LockReleaseTokenPool.sol";
import {LiquidityManager} from "../../../pools/liquidity/LiquidityManager.sol";

import {ERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract LiquidityManagerSetup is Test {
  event LiquidityTransferred(
    uint64 indexed fromChainSelector,
    uint64 indexed toChainSelector,
    address indexed to,
    uint256 amount
  );

  IERC20 public s_token;
  LiquidityManager internal s_liquidityManager;
  LockReleaseTokenPool internal s_lockReleaseTokenPool;
  uint64 internal immutable i_localChainSelector = 1234;

  function setUp() external {
    s_token = new ERC20("Test", "TEST");

    s_lockReleaseTokenPool = new LockReleaseTokenPool(s_token, new address[](0), address(1), true);
    s_liquidityManager = new LiquidityManager(s_token, i_localChainSelector, s_lockReleaseTokenPool);

    s_lockReleaseTokenPool.setLiquidityManager(address(s_liquidityManager));
  }
}

contract LiquidityManager_rebalanceLiquidity is LiquidityManagerSetup {
  function test_rebalanceLiquiditySuccess() external {
    uint256 amount = 12345679;
    uint64 destChainSelector = 123;
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);

    LiquidityManager.CrossChainLiquidityManagerArgs[]
      memory args = new LiquidityManager.CrossChainLiquidityManagerArgs[](1);
    args[0] = LiquidityManager.CrossChainLiquidityManagerArgs({
      destLiquidityManager: address(s_liquidityManager),
      bridge: IBridge(address(0)),
      destChainSelector: destChainSelector,
      enabled: true
    });
    s_liquidityManager.setCrossChainLiquidityManager(args);

    vm.expectEmit();

    emit LiquidityTransferred(i_localChainSelector, destChainSelector, address(s_liquidityManager), amount);

    s_liquidityManager.rebalanceLiquidity(destChainSelector, amount);
  }

  // Reverts

  function test_InvalidDestinationChainReverts() external {
    uint256 amount = 12345679;
    uint64 destChainSelector = 123;
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);

    vm.expectRevert(abi.encodeWithSelector(LiquidityManager.InvalidDestinationChain.selector, destChainSelector));

    s_liquidityManager.rebalanceLiquidity(destChainSelector, amount);
  }

  function test_InsufficientLiquidityReverts() external {
    uint256 amount = 1245;

    vm.expectRevert(abi.encodeWithSelector(LiquidityManager.InsufficientLiquidity.selector, amount, 0));

    s_liquidityManager.rebalanceLiquidity(0, amount);
  }
}
