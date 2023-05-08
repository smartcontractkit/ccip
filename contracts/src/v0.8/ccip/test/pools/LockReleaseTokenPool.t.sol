// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {MockERC20} from "../mocks/MockERC20.sol";
import {LockReleaseTokenPool} from "../../pools/LockReleaseTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

contract LockReleaseTokenPoolSetup is BaseTest {
  IERC20 internal s_token;
  LockReleaseTokenPool internal s_lockReleaseTokenPool;
  address s_allowedOnRamp = address(123);
  address s_allowedOffRamp = address(234);

  function setUp() public virtual override {
    BaseTest.setUp();
    s_token = new MockERC20("LINK", "LNK", OWNER, 2**256 - 1);
    s_lockReleaseTokenPool = new LockReleaseTokenPool(s_token, rateLimiterConfig());

    TokenPool.RampUpdate[] memory onRamps = new TokenPool.RampUpdate[](1);
    onRamps[0] = TokenPool.RampUpdate({ramp: s_allowedOnRamp, allowed: true});
    TokenPool.RampUpdate[] memory offRamps = new TokenPool.RampUpdate[](1);
    offRamps[0] = TokenPool.RampUpdate({ramp: s_allowedOffRamp, allowed: true});

    s_lockReleaseTokenPool.applyRampUpdates(onRamps, offRamps);
  }
}

contract LockReleaseTokenPool_releaseOrMint is LockReleaseTokenPoolSetup {
  event TokensConsumed(uint256 tokens);
  event Released(address indexed sender, address indexed recipient, uint256 amount);

  function testReleaseOrMintSuccess(address recipient, uint256 amount) public {
    // Since the owner already has tokens this would break the checks
    vm.assume(recipient != OWNER);
    vm.assume(recipient != address(0));

    // Makes sure the pool always has enough funds
    deal(address(s_token), address(s_lockReleaseTokenPool), amount);
    changePrank(s_allowedOffRamp);

    uint256 capacity = rateLimiterConfig().capacity;
    // Determine if we hit the rate limit or the txs should succeed.
    if (amount > capacity) {
      vm.expectRevert(abi.encodeWithSelector(RateLimiter.ConsumingMoreThanMaxCapacity.selector, capacity, amount));
    } else {
      // Only rate limit if the amount is >0
      if (amount > 0) {
        vm.expectEmit();
        emit TokensConsumed(amount);
      }

      vm.expectEmit();
      emit Released(s_allowedOffRamp, recipient, amount);
    }

    s_lockReleaseTokenPool.releaseOrMint(bytes(""), recipient, amount, SOURCE_CHAIN_ID, bytes(""));
  }
}

contract LockReleaseTokenPool_getProvidedLiquidity is LockReleaseTokenPoolSetup {
  function testGetProvidedLiquiditySuccess(uint256 amount) public {
    s_token.approve(address(s_lockReleaseTokenPool), amount);

    s_lockReleaseTokenPool.addLiquidity(amount);

    assertEq(s_lockReleaseTokenPool.getProvidedLiquidity(OWNER), amount);
  }
}

contract LockReleaseTokenPool_addLiquidity is LockReleaseTokenPoolSetup {
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
