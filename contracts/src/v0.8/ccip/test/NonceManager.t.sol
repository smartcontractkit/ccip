// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {NonceManager} from "../NonceManager.sol";
import {EVM2EVMMultiOnRampSetup} from "./onRamp/EVM2EVMMultiOnRampSetup.t.sol";

contract NonceManagerTest_incrementOutboundNonce is EVM2EVMMultiOnRampSetup {
  function test_incrementOutboundNonce_Success() public {
    vm.startPrank(address(s_onRamp));
    bytes memory sender = abi.encode(address(this));

    assertEq(s_nonceManager.getOutboundNonce(DEST_CHAIN_SELECTOR, sender), 0);

    uint64 outboundNonce = s_nonceManager.incrementOutboundNonce(DEST_CHAIN_SELECTOR, sender);
    assertEq(outboundNonce, 1);
  }

  // TODO: move upgradability tests for both OnRmap and OffRamp here
}

contract NonceManager_applyRampUpdates is EVM2EVMMultiOnRampSetup {
  function test_applyRampUpdates_Success() public {
    address newOnRamp = vm.addr(1);
    address prevOnRamp = vm.addr(2);
    NonceManager.PreviousRamp[] memory prevOnRamps = new NonceManager.PreviousRamp[](1);
    prevOnRamps[0] = NonceManager.PreviousRamp(DEST_CHAIN_SELECTOR, prevOnRamp);

    vm.expectEmit();
    emit NonceManager.OnRampUpdated(newOnRamp);
    vm.expectEmit();
    emit NonceManager.PreviousOnRampUpdated(DEST_CHAIN_SELECTOR, prevOnRamp);

    s_nonceManager.applyRampUpdates(newOnRamp, prevOnRamps);

    assertEq(s_nonceManager.getOnRamp(), newOnRamp);
    assertEq(s_nonceManager.getPrevOnRamp(DEST_CHAIN_SELECTOR), prevOnRamp);
  }

  function test_applyRampUpdatesOnlyOnRampUpdate_Success() public {
    address newOnRamp = vm.addr(1);

    vm.expectEmit();
    emit NonceManager.OnRampUpdated(newOnRamp);

    s_nonceManager.applyRampUpdates(newOnRamp, new NonceManager.PreviousRamp[](0));

    assertEq(s_nonceManager.getOnRamp(), newOnRamp);
  }

  function test_InvalidRampUpdatePreviousOnRampAlreadySet_Revert() public {
    NonceManager.PreviousRamp[] memory prevOnRamps = new NonceManager.PreviousRamp[](1);
    prevOnRamps[0] = NonceManager.PreviousRamp(DEST_CHAIN_SELECTOR, address(vm.addr(1)));

    s_nonceManager.applyRampUpdates(address(0), prevOnRamps);

    prevOnRamps[0] = NonceManager.PreviousRamp(DEST_CHAIN_SELECTOR, address(vm.addr(2)));

    vm.expectRevert(NonceManager.InvalidRampUpdate.selector);
    s_nonceManager.applyRampUpdates(address(0), prevOnRamps);
  }
}
