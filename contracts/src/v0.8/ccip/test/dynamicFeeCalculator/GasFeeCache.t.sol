// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {GasFeeCache} from "../../dynamicFeeCalculator/GasFeeCache.sol";

contract GasFeeCacheSetup is TokenSetup {
  GasFeeCache s_gasFeeCache;

  function setUp() public virtual override {
    TokenSetup.setUp();
    CCIP.FeeUpdate[] memory fees = new CCIP.FeeUpdate[](1);
    fees[0] = CCIP.FeeUpdate({chainId: DEST_CHAIN_ID, linkPerUnitGas: 100});
    address[] memory feeUpdaters = new address[](0);

    s_gasFeeCache = new GasFeeCache(fees, feeUpdaters);
  }
}

contract GasFeeCache_getFee is GasFeeCacheSetup {
  function testGetFeeSuccess() public {
    assertEq(100, s_gasFeeCache.getFee(DEST_CHAIN_ID));
  }
}
