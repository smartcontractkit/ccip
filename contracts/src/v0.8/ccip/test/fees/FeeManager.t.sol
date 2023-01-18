// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";

import {GE} from "../../models/GE.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {FeeManager} from "../../fees/FeeManager.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

contract FeeManagerSetup is TokenSetup {
  FeeManager s_feeManager;

  function setUp() public virtual override {
    TokenSetup.setUp();
    GE.FeeUpdate[] memory fees = new GE.FeeUpdate[](1);
    fees[0] = GE.FeeUpdate({sourceFeeToken: s_sourceTokens[0], destChainId: DEST_CHAIN_ID, linkPerUnitGas: 100});
    address[] memory feeUpdaters = new address[](0);

    s_feeManager = new FeeManager(fees, feeUpdaters, uint128(TWELVE_HOURS));
  }
}

contract FeeManager_getFee is FeeManagerSetup {
  function testGetFeeSuccess() public {
    assertEq(100, s_feeManager.getFee(s_sourceTokens[0], DEST_CHAIN_ID));
  }

  function testUnsupportedTokenReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, s_sourceTokens[1], DEST_CHAIN_ID)
    );
    s_feeManager.getFee(s_sourceTokens[1], DEST_CHAIN_ID);
  }

  function testUnsupportedChainReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, s_sourceTokens[0], DEST_CHAIN_ID + 1)
    );
    s_feeManager.getFee(s_sourceTokens[0], DEST_CHAIN_ID + 1);
  }

  function testGetFeeStaleReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(abi.encodeWithSelector(IFeeManager.StaleFee.selector, TWELVE_HOURS, diff));
    s_feeManager.getFee(s_sourceTokens[0], DEST_CHAIN_ID);
  }
}

contract FeeManager_withdrawToken is FeeManagerSetup {
  IERC20 internal s_token;

  function setUp() public virtual override {
    FeeManagerSetup.setUp();
    s_token = IERC20(s_sourceTokens[0]);
    changePrank(OWNER);
    s_token.transfer(address(s_feeManager), 100);
  }

  function testWithdrawTokenSuccess() public {
    s_feeManager.withdrawToken(address(s_token), address(this), 100);

    assertEq(0, s_token.balanceOf(address(s_feeManager)));
    assertEq(100, s_token.balanceOf(address(this)));
  }

  function testNonOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    s_feeManager.withdrawToken(address(s_token), address(this), 100);
  }

  function testInvalidWithdrawalReverts() public {
    vm.expectRevert(IFeeManager.InvalidWithdrawalAddress.selector);
    s_feeManager.withdrawToken(address(s_token), address(0), 100);
  }
}
