// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import {WrappedTokenPool} from "../../pools/WrappedTokenPool.sol";

contract WrappedTokenPoolSetup is BaseTest {
  WrappedTokenPool internal s_wrappedTokenPool;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_wrappedTokenPool = new WrappedTokenPool("Test", "TST");
  }
}

contract WrappedTokenPool_releaseOrMint is WrappedTokenPoolSetup {
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testReleaseOrMintSuccess() public {
    vm.expectEmit(true, true, true, true);
    emit Minted(OWNER, OWNER, 1);
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
    assertEq(s_wrappedTokenPool.balanceOf(OWNER), 1);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
  }

  function testNonOwnerOrOffRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(IPool.PermissionsError.selector);
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
  }
}

contract WrappedTokenPool_lockOrBurn is WrappedTokenPoolSetup {
  event Burned(address indexed account, uint256 amount);

  function testLockOrBurnSuccess() public {
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
    assertEq(s_wrappedTokenPool.balanceOf(OWNER), 1);

    s_wrappedTokenPool.transfer(address(s_wrappedTokenPool), 1);

    vm.expectEmit(true, true, true, true);
    emit Burned(OWNER, 1);
    s_wrappedTokenPool.lockOrBurn(1, OWNER);
    assertEq(s_wrappedTokenPool.balanceOf(OWNER), 0);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.lockOrBurn(1, OWNER);
  }

  function testNonOwnerOrOnRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(IPool.PermissionsError.selector);
    s_wrappedTokenPool.lockOrBurn(1, OWNER);
  }
}
