// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";

contract LockReleaseTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  LockReleaseTokenPool internal s_lockReleaseTokenPool;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new MockERC20("LINK", "LNK", OWNER, 2**256 - 1);
    s_lockReleaseTokenPool = new LockReleaseTokenPool(s_token);
  }
}

contract LockReleaseTokenPool_getProvidedLiquidity is LockReleaseTokenPoolSetup {
  // Success
  function testGetProvidedLiquiditySuccess(uint256 amount) public {
    s_token.approve(address(s_lockReleaseTokenPool), amount);

    s_lockReleaseTokenPool.addLiquidity(amount);

    assertEq(s_lockReleaseTokenPool.getProvidedLiquidity(OWNER), amount);
  }
}

contract LockReleaseTokenPool_addLiquidity is LockReleaseTokenPoolSetup {
  // Success
  function testAddLiquiditySuccess(uint256 amount) public {
    uint256 balancePre = s_token.balanceOf(OWNER);
    s_token.approve(address(s_lockReleaseTokenPool), amount);

    s_lockReleaseTokenPool.addLiquidity(amount);

    assertEq(s_token.balanceOf(OWNER), balancePre - amount);
    assertEq(s_token.balanceOf(address(s_lockReleaseTokenPool)), amount);
  }

  // Reverts

  function testExceedsAllowance(uint256 amount) public {
    vm.assume(amount > 0);
    vm.expectRevert("ERC20: transfer amount exceeds allowance");
    s_lockReleaseTokenPool.addLiquidity(amount);
  }
}

contract LockReleaseTokenPool_removeLiquidity is LockReleaseTokenPoolSetup {
  // Success
  function testRemoveLiquiditySuccess(uint256 amount) public {
    uint256 balancePre = s_token.balanceOf(OWNER);
    s_token.approve(address(s_lockReleaseTokenPool), amount);
    s_lockReleaseTokenPool.addLiquidity(amount);

    s_lockReleaseTokenPool.removeLiquidity(amount);

    assertEq(s_token.balanceOf(OWNER), balancePre);
  }

  // Reverts
  function testWithdrawalTooHighReverts(uint256 balance, uint256 withdrawal) public {
    vm.assume(balance < withdrawal);
    s_token.approve(address(s_lockReleaseTokenPool), balance);
    s_lockReleaseTokenPool.addLiquidity(balance);

    vm.expectRevert(LockReleaseTokenPool.WithdrawalTooHigh.selector);
    s_lockReleaseTokenPool.removeLiquidity(withdrawal);
  }

  function testInsufficientLiquidityReverts() public {
    uint256 maxUint256 = 2**256 - 1;
    s_token.approve(address(s_lockReleaseTokenPool), maxUint256);
    s_lockReleaseTokenPool.addLiquidity(maxUint256);

    changePrank(address(s_lockReleaseTokenPool));
    s_token.transfer(OWNER, maxUint256);
    changePrank(OWNER);

    vm.expectRevert(LockReleaseTokenPool.InsufficientLiquidity.selector);
    s_lockReleaseTokenPool.removeLiquidity(1);
  }
}
