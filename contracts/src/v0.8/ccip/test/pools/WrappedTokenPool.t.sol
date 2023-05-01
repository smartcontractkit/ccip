// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import {WrappedTokenPool} from "../../pools/WrappedTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

contract WrappedTokenPoolSetup is BaseTest {
  WrappedTokenPool internal s_wrappedTokenPool;
  address s_onRamp = address(1234567);
  address s_offRamp = address(7654321);

  function setUp() public virtual override {
    BaseTest.setUp();
    s_wrappedTokenPool = new WrappedTokenPool("Test", "TST", 18, rateLimiterConfig());

    TokenPool.RampUpdate[] memory onRamps = new TokenPool.RampUpdate[](1);
    onRamps[0] = TokenPool.RampUpdate({ramp: address(s_onRamp), allowed: true});
    TokenPool.RampUpdate[] memory offRamps = new TokenPool.RampUpdate[](1);
    offRamps[0] = TokenPool.RampUpdate({ramp: address(s_offRamp), allowed: true});

    s_wrappedTokenPool.applyRampUpdates(onRamps, offRamps);
  }
}

contract WrappedTokenPool_releaseOrMint is WrappedTokenPoolSetup {
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testReleaseOrMintSuccess() public {
    changePrank(s_offRamp);

    vm.expectEmit();
    emit Minted(s_offRamp, s_offRamp, 1);

    s_wrappedTokenPool.releaseOrMint(bytes(""), s_offRamp, 1, SOURCE_CHAIN_ID, bytes(""));

    assertEq(s_wrappedTokenPool.balanceOf(s_offRamp), 1);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.releaseOrMint(bytes(""), OWNER, 1, SOURCE_CHAIN_ID, bytes(""));
  }

  function testNonOffRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(TokenPool.PermissionsError.selector);
    s_wrappedTokenPool.releaseOrMint(bytes(""), OWNER, 1, SOURCE_CHAIN_ID, bytes(""));
  }
}

contract WrappedTokenPool_lockOrBurn is WrappedTokenPoolSetup {
  event Burned(address indexed account, uint256 amount);

  function testLockOrBurnSuccess() public {
    changePrank(s_offRamp);
    s_wrappedTokenPool.releaseOrMint(bytes(""), s_onRamp, 1, SOURCE_CHAIN_ID, bytes(""));
    assertEq(s_wrappedTokenPool.balanceOf(s_onRamp), 1);

    changePrank(s_onRamp);
    s_wrappedTokenPool.transfer(address(s_wrappedTokenPool), 1);

    vm.expectEmit();
    emit Burned(s_onRamp, 1);

    s_wrappedTokenPool.lockOrBurn(s_onRamp, bytes(""), 1, DEST_CHAIN_ID, bytes(""));
    assertEq(s_wrappedTokenPool.balanceOf(s_onRamp), 0);
  }

  function testPausedReverts() public {
    s_wrappedTokenPool.pause();
    changePrank(s_onRamp);
    assertTrue(s_wrappedTokenPool.paused());

    vm.expectRevert("Pausable: paused");
    s_wrappedTokenPool.lockOrBurn(s_onRamp, bytes(""), 1, DEST_CHAIN_ID, bytes(""));
  }

  function testNonOnRampReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(TokenPool.PermissionsError.selector);
    s_wrappedTokenPool.lockOrBurn(s_onRamp, bytes(""), 1, DEST_CHAIN_ID, bytes(""));
  }
}
