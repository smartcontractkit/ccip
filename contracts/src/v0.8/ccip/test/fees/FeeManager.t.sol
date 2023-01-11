// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";

import {GE} from "../../models/GE.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {FeeManager} from "../../fees/FeeManager.sol";

contract FeeManagerSetup is TokenSetup {
  FeeManager s_feeManager;

  function setUp() public virtual override {
    TokenSetup.setUp();
    GE.FeeUpdate[] memory fees = new GE.FeeUpdate[](1);
    fees[0] = GE.FeeUpdate({chainId: DEST_CHAIN_ID, linkPerUnitGas: 100});
    address[] memory feeUpdaters = new address[](0);

    s_feeManager = new FeeManager(fees, feeUpdaters, uint128(TWELVE_HOURS));
  }
}

contract FeeManager_getFee is FeeManagerSetup {
  function testGetFeeSuccess() public {
    assertEq(100, s_feeManager.getFee(DEST_CHAIN_ID));
  }

  function testGetFeeStaleReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(abi.encodeWithSelector(IFeeManager.StaleFee.selector, TWELVE_HOURS, diff));
    s_feeManager.getFee(DEST_CHAIN_ID);
  }
}
