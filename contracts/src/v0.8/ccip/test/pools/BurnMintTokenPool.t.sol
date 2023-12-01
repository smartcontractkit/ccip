// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";

contract BurnMintTokenPoolSetup is BurnMintSetup {
  BurnMintTokenPool internal s_pool;

  function setUp() public virtual override {
    BurnMintSetup.setUp();

    s_pool = new BurnMintTokenPool(s_burnMintERC677, new address[](0), address(s_mockARM));
    s_burnMintERC677.grantMintAndBurnRoles(address(s_pool));

    applyRampUpdates(address(s_pool));
  }
}

contract BurnMintTokenPool_lockOrBurn is BurnMintTokenPoolSetup {
  function testSetupSuccess() public {
    assertEq(address(s_burnMintERC677), address(s_pool.getToken()));
    assertEq(address(s_mockARM), s_pool.getArmProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq("BurnMintTokenPool 1.2.0", s_pool.typeAndVersion());
  }

  function testPoolBurnSuccess() public {
    uint256 burnAmount = 20_000e18;

    deal(address(s_burnMintERC677), address(s_pool), burnAmount);
    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), burnAmount);

    vm.startPrank(s_burnMintOnRamp);

    vm.expectEmit();
    emit TokensConsumed(burnAmount);

    vm.expectEmit();
    emit Transfer(address(s_pool), address(0), burnAmount);

    vm.expectEmit();
    emit Burned(address(s_burnMintOnRamp), burnAmount);

    bytes4 expectedSignature = bytes4(keccak256("burn(uint256)"));
    vm.expectCall(address(s_burnMintERC677), abi.encodeWithSelector(expectedSignature, burnAmount));

    s_pool.lockOrBurn(OWNER, bytes(""), burnAmount, 0, bytes(""));

    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), 0);
  }

  // Should not burn tokens if cursed.
  function testPoolBurnRevertNotHealthyReverts() public {
    s_mockARM.voteToCurse(bytes32(0));
    uint256 before = s_burnMintERC677.balanceOf(address(s_pool));
    vm.startPrank(s_burnMintOnRamp);

    vm.expectRevert(EVM2EVMOnRamp.BadARMSignal.selector);
    s_pool.lockOrBurn(OWNER, bytes(""), 1e5, 0, bytes(""));

    assertEq(s_burnMintERC677.balanceOf(address(s_pool)), before);
  }

  function testPermissionsErrorReverts() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(TokenPool.PermissionsError.selector);
    s_pool.lockOrBurn(OWNER, bytes(""), 1, 0, bytes(""));
  }
}

contract BurnMintTokenPool_releaseOrMint is BurnMintTokenPoolSetup {
  function testPoolMintSuccess() public {
    uint256 amount = 1e19;
    vm.startPrank(s_burnMintOffRamp);
    vm.expectEmit();
    emit Transfer(address(0), OWNER, amount);
    s_pool.releaseOrMint(bytes(""), OWNER, amount, 0, bytes(""));
    assertEq(s_burnMintERC677.balanceOf(OWNER), amount);
  }

  function testPoolMintNotHealthyReverts() public {
    // Should not mint tokens if cursed.
    s_mockARM.voteToCurse(bytes32(0));
    uint256 before = s_burnMintERC677.balanceOf(OWNER);
    vm.startPrank(s_burnMintOffRamp);
    vm.expectRevert(EVM2EVMOffRamp.BadARMSignal.selector);
    s_pool.releaseOrMint(bytes(""), OWNER, 1e5, 0, bytes(""));
    assertEq(s_burnMintERC677.balanceOf(OWNER), before);
  }

  function testPermissionsErrorReverts() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(TokenPool.PermissionsError.selector);
    s_pool.releaseOrMint(bytes(""), OWNER, 1, 0, bytes(""));
  }
}
