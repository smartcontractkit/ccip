// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {GE} from "../../models/GE.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {GasFeeCache, IGasFeeCache} from "../../gasFeeCache/GasFeeCache.sol";

contract GasFeeCacheSetup is TokenSetup {
  GasFeeCache s_gasFeeCache;

  function setUp() public virtual override {
    TokenSetup.setUp();
    GE.FeeUpdate[] memory fees = new GE.FeeUpdate[](1);
    fees[0] = GE.FeeUpdate({chainId: DEST_CHAIN_ID, linkPerUnitGas: 100});
    address[] memory feeUpdaters = new address[](0);

    s_gasFeeCache = new GasFeeCache(fees, feeUpdaters, uint128(TWELVE_HOURS));
  }
}

contract GasFeeCache_getFee is GasFeeCacheSetup {
  function testGetFeeSuccess() public {
    assertEq(100, s_gasFeeCache.getFee(DEST_CHAIN_ID));
  }

  function testGetFeeStaleReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(abi.encodeWithSelector(IGasFeeCache.StaleFee.selector, TWELVE_HOURS, diff));
    s_gasFeeCache.getFee(DEST_CHAIN_ID);
  }
}
