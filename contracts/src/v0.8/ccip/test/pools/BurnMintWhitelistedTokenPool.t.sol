// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {TokenPool} from "../../pools/TokenPool.sol";
import {BurnMintSetup} from "./BurnMintSetup.t.sol";
import {GravitaDebtToken} from "./BurnMintWhitelistedToken.sol";
import {BurnMintWhitelistedTokenPool} from "../../pools/BurnMintWhitelistedTokenPool.sol";
import {IBurnMintWhitelisted} from "../../interfaces/IBurnMintWhitelisted.sol";

contract BurnMintWhitelistedTokenPoolSetup is BurnMintSetup {
  BurnMintWhitelistedTokenPool internal s_pool;
  GravitaDebtToken internal s_token;

  function setUp() public virtual override {
    BurnMintSetup.setUp();

    s_token = new GravitaDebtToken(makeAddr("lz"));
    s_pool = new BurnMintWhitelistedTokenPool(IBurnMintWhitelisted(address(s_token)), new address[](0), address(s_mockARM));
    s_token.addWhitelist(address(s_pool));

    applyRampUpdates(address(s_pool));
  }
}

contract BurnMintWhitelistedTokenPool_mintBurn is BurnMintWhitelistedTokenPoolSetup {
  function testSetupSuccess() public {
    assertEq(address(s_token), address(s_pool.getToken()));
    assertEq(address(s_mockARM), s_pool.getArmProxy());
    assertEq(false, s_pool.getAllowListEnabled());
    assertEq("BurnMintTokenPool 1.4.0-dev", s_pool.typeAndVersion());
  }

  function testPoolMintBurnSuccess() public {
    uint256 amount = 20_000e18;

    vm.startPrank(s_burnMintOffRamp);

    vm.expectEmit();
    emit Transfer(address(0), address(s_pool), amount);
    vm.expectEmit();
    emit Transfer(address(s_pool), OWNER, amount);
    vm.expectEmit();
    emit Minted(address(s_burnMintOffRamp), OWNER, amount);

    s_pool.releaseOrMint(bytes(""), OWNER, amount, 0, bytes(""));
    assertEq(s_token.balanceOf(OWNER), amount);
    assert(s_token.totalSupply() == amount);

    // Mint some tokens for the pool so that we can burn them later.
    s_pool.releaseOrMint(bytes(""), address(s_pool), amount, 0, bytes(""));
    assertEq(s_token.balanceOf(address(s_pool)), amount);

    vm.startPrank(s_burnMintOnRamp);

    vm.expectEmit();
    emit TokensConsumed(amount);

    vm.expectEmit();
    emit Transfer(address(s_pool), address(0), amount);
    vm.expectEmit();
    emit Burned(address(s_burnMintOnRamp), amount);

    bytes4 expectedSignature = bytes4(keccak256("burnFromWhitelistedContract(uint256)"));
    vm.expectCall(address(s_token), abi.encodeWithSelector(expectedSignature, amount));

    s_pool.lockOrBurn(OWNER, bytes(""), amount, 0, bytes(""));

    assertEq(s_token.balanceOf(address(s_pool)), 0);
  }

  // Should not burn tokens if cursed.
  function testPoolBurnRevertNotHealthyReverts() public {
    s_mockARM.voteToCurse(bytes32(0));
    uint256 before = s_token.balanceOf(address(s_pool));
    vm.startPrank(s_burnMintOnRamp);

    vm.expectRevert(EVM2EVMOnRamp.BadARMSignal.selector);
    s_pool.lockOrBurn(OWNER, bytes(""), 1e5, 0, bytes(""));

    assertEq(s_token.balanceOf(address(s_pool)), before);
  }

  function testPermissionsErrorReverts() public {
    vm.startPrank(STRANGER);

    vm.expectRevert(TokenPool.PermissionsError.selector);
    s_pool.lockOrBurn(OWNER, bytes(""), 1, 0, bytes(""));
  }
}
