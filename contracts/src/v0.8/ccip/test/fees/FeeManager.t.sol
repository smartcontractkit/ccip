// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";
import {IRouter} from "../../interfaces/router/IRouter.sol";

import {Internal} from "../../models/Internal.sol";
import {TokenSetup} from "../TokenSetup.t.sol";
import {FeeManager} from "../../fees/FeeManager.sol";
import {RouterSetup} from "../router/RouterSetup.t.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

contract FeeManagerSetup is TokenSetup, RouterSetup {
  FeeManager s_sourceFeeManager;
  FeeManager s_destFeeManager;

  function setUp() public virtual override(TokenSetup, RouterSetup) {
    TokenSetup.setUp();
    RouterSetup.setUp();
    if (address(s_sourceFeeManager) == address(0)) {
      Internal.FeeUpdate[] memory feeUpdates = new Internal.FeeUpdate[](2);
      feeUpdates[0] = Internal.FeeUpdate({
        sourceFeeToken: s_sourceTokens[0],
        destChainId: DEST_CHAIN_ID,
        feeTokenBaseUnitsPerUnitGas: 100
      });
      feeUpdates[1] = Internal.FeeUpdate({
        sourceFeeToken: IRouter(s_sourceRouter).getWrappedNative(),
        destChainId: DEST_CHAIN_ID,
        feeTokenBaseUnitsPerUnitGas: 101
      });
      address[] memory feeUpdaters = new address[](0);
      s_sourceFeeManager = new FeeManager(feeUpdates, feeUpdaters, TWELVE_HOURS);
    }
    if (address(s_destFeeManager) == address(0)) {
      Internal.FeeUpdate[] memory feeUpdates = new Internal.FeeUpdate[](2);
      feeUpdates[0] = Internal.FeeUpdate({
        sourceFeeToken: s_sourceTokens[0],
        destChainId: SOURCE_CHAIN_ID,
        feeTokenBaseUnitsPerUnitGas: 100
      });
      feeUpdates[1] = Internal.FeeUpdate({
        sourceFeeToken: IRouter(s_destRouter).getWrappedNative(),
        destChainId: SOURCE_CHAIN_ID,
        feeTokenBaseUnitsPerUnitGas: 101
      });
      address[] memory feeUpdaters = new address[](0);
      s_destFeeManager = new FeeManager(feeUpdates, feeUpdaters, TWELVE_HOURS);
    }
  }
}

contract FeeManager_getFeeTokenBaseUnitsPerUnitGas is FeeManagerSetup {
  function testGetFeeSuccess() public {
    assertEq(100, s_sourceFeeManager.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID));
  }

  function testUnsupportedTokenReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, s_sourceTokens[1], DEST_CHAIN_ID)
    );
    s_sourceFeeManager.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[1], DEST_CHAIN_ID);
  }

  function testUnsupportedChainReverts() public {
    vm.expectRevert(
      abi.encodeWithSelector(IFeeManager.TokenOrChainNotSupported.selector, s_sourceTokens[0], DEST_CHAIN_ID + 1)
    );
    s_sourceFeeManager.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID + 1);
  }

  function testGetFeeStaleReverts() public {
    uint256 diff = TWELVE_HOURS + 1;
    vm.warp(block.timestamp + diff);
    vm.expectRevert(abi.encodeWithSelector(IFeeManager.StaleFee.selector, TWELVE_HOURS, diff));
    s_sourceFeeManager.getFeeTokenBaseUnitsPerUnitGas(s_sourceTokens[0], DEST_CHAIN_ID);
  }
}

contract FeeManager_withdrawToken is FeeManagerSetup {
  IERC20 internal s_token;

  function setUp() public virtual override {
    FeeManagerSetup.setUp();
    s_token = IERC20(s_sourceTokens[0]);
    changePrank(OWNER);
    s_token.transfer(address(s_sourceFeeManager), 100);
  }

  function testWithdrawTokenSuccess() public {
    s_sourceFeeManager.withdrawToken(address(s_token), address(this), 100);

    assertEq(0, s_token.balanceOf(address(s_sourceFeeManager)));
    assertEq(100, s_token.balanceOf(address(this)));
  }

  function testNonOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    s_sourceFeeManager.withdrawToken(address(s_token), address(this), 100);
  }

  function testInvalidWithdrawalReverts() public {
    vm.expectRevert(IFeeManager.InvalidWithdrawalAddress.selector);
    s_sourceFeeManager.withdrawToken(address(s_token), address(0), 100);
  }
}
