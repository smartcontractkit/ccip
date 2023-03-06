// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import {WrappedTokenPool} from "../../pools/WrappedTokenPool.sol";

contract WrappedTokenPoolSetup is BaseTest {
  WrappedTokenPool internal s_wrappedTokenPool;
  address s_onRamp = address(1234567);
  address s_offRamp = address(7654321);

  function setUp() public virtual override {
    BaseTest.setUp();
    s_wrappedTokenPool = new WrappedTokenPool("Test", "TST");

    IPool.RampUpdate[] memory onRamps = new IPool.RampUpdate[](1);
    onRamps[0] = IPool.RampUpdate({ramp: address(s_onRamp), allowed: true});
    IPool.RampUpdate[] memory offRamps = new IPool.RampUpdate[](1);
    offRamps[0] = IPool.RampUpdate({ramp: address(s_offRamp), allowed: true});

    s_wrappedTokenPool.applyRampUpdates(onRamps, offRamps);
  }
}

contract WrappedTokenPool_releaseOrMint is WrappedTokenPoolSetup {
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testReleaseOrMintSuccess() public {
    changePrank(s_offRamp);
    vm.expectEmit(true, true, true, true);
    emit Minted(s_offRamp, s_offRamp, 1);
    s_wrappedTokenPool.releaseOrMint(s_offRamp, 1);
    assertEq(s_wrappedTokenPool.balanceOf(s_offRamp), 1);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
  }

  function testNonOffRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(IPool.PermissionsError.selector);
    s_wrappedTokenPool.releaseOrMint(OWNER, 1);
  }
}

contract WrappedTokenPool_lockOrBurn is WrappedTokenPoolSetup {
  event Burned(address indexed account, uint256 amount);

  function testLockOrBurnSuccess() public {
    changePrank(s_offRamp);
    s_wrappedTokenPool.releaseOrMint(s_onRamp, 1);
    assertEq(s_wrappedTokenPool.balanceOf(s_onRamp), 1);

    changePrank(s_onRamp);
    s_wrappedTokenPool.transfer(address(s_wrappedTokenPool), 1);

    vm.expectEmit(true, true, true, true);
    emit Burned(s_onRamp, 1);

    s_wrappedTokenPool.lockOrBurn(1, s_onRamp);
    assertEq(s_wrappedTokenPool.balanceOf(s_onRamp), 0);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    changePrank(s_onRamp);
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.lockOrBurn(1, OWNER);
  }

  function testNonOnRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(IPool.PermissionsError.selector);
    s_wrappedTokenPool.lockOrBurn(1, OWNER);
  }
}
