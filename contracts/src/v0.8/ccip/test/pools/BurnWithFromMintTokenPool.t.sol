// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";
import {BurnWithFromMintTokenPool} from "../../pools/BurnWithFromMintTokenPool.sol";

contract BurnWithFromMintTokenPoolSetup is BurnMintSetup {
  BurnWithFromMintTokenPool internal s_pool;

  function setUp() public virtual override {
    BurnMintSetup.setUp();

    s_pool = new BurnWithFromMintTokenPool(s_burnMintERC677, new address[](0), address(s_mockARM));
    s_burnMintERC677.grantMintAndBurnRoles(address(s_pool));

    applyRampUpdates(address(s_pool));
  }
}

contract BurnWithFromMintTokenPool_lockOrBurn is BurnWithFromMintTokenPoolSetup {
  function testSetupSuccess() public {
    assertEq(address(s_burnMintERC677), address(s_pool.getToken()));
    assertEq(address(s_mockARM), s_pool.getArmProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq(type(uint256).max, s_burnMintERC677.allowance(address(s_pool), address(s_pool)));
    assertEq("BurnWithFromMintTokenPool 1.3.0-dev", s_pool.typeAndVersion());
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

    bytes4 expectedSignature = bytes4(keccak256("burn(address,uint256)"));
    vm.expectCall(address(s_burnMintERC677), abi.encodeWithSelector(expectedSignature, address(s_pool), burnAmount));

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
