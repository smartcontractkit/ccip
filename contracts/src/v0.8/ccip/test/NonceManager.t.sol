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
}

contract NonceManager_applyPreviousRampsUpdates is EVM2EVMMultiOnRampSetup {
  function test_SingleRampUpdate() public {
    address prevOnRamp = vm.addr(1);
    NonceManager.PreviousRampsArgs[] memory previousRamps = new NonceManager.PreviousRampsArgs[](1);
    previousRamps[0] = NonceManager.PreviousRampsArgs(DEST_CHAIN_SELECTOR, NonceManager.PreviousRamps(prevOnRamp));

    vm.expectEmit();
    emit NonceManager.PreviousOnRampUpdated(DEST_CHAIN_SELECTOR, prevOnRamp);

    s_nonceManager.applyPreviousRampsUpdates(previousRamps);

    _assertPreviousRampsEqual(s_nonceManager.getPreviousRamps(DEST_CHAIN_SELECTOR), previousRamps[0].prevRamps);
  }

  function test_MultipleRampsUpdates() public {
    address prevOnRamp1 = vm.addr(1);
    address prevOnRamp2 = vm.addr(2);
    NonceManager.PreviousRampsArgs[] memory previousRamps = new NonceManager.PreviousRampsArgs[](2);
    previousRamps[0] = NonceManager.PreviousRampsArgs(DEST_CHAIN_SELECTOR, NonceManager.PreviousRamps(prevOnRamp1));
    previousRamps[1] = NonceManager.PreviousRampsArgs(DEST_CHAIN_SELECTOR + 1, NonceManager.PreviousRamps(prevOnRamp2));

    vm.expectEmit();
    emit NonceManager.PreviousOnRampUpdated(DEST_CHAIN_SELECTOR, prevOnRamp1);
    vm.expectEmit();
    emit NonceManager.PreviousOnRampUpdated(DEST_CHAIN_SELECTOR + 1, prevOnRamp2);

    s_nonceManager.applyPreviousRampsUpdates(previousRamps);

    _assertPreviousRampsEqual(s_nonceManager.getPreviousRamps(DEST_CHAIN_SELECTOR), previousRamps[0].prevRamps);
    _assertPreviousRampsEqual(s_nonceManager.getPreviousRamps(DEST_CHAIN_SELECTOR + 1), previousRamps[1].prevRamps);
  }

  function test_ZeroInput() public {
    vm.recordLogs();
    s_nonceManager.applyPreviousRampsUpdates(new NonceManager.PreviousRampsArgs[](0));

    assertEq(vm.getRecordedLogs().length, 0);
  }

  function test_PreviousRampAlreadySetOnRamp_Revert() public {
    NonceManager.PreviousRampsArgs[] memory previousRamps = new NonceManager.PreviousRampsArgs[](1);
    previousRamps[0] =
      NonceManager.PreviousRampsArgs(DEST_CHAIN_SELECTOR, NonceManager.PreviousRamps(address(vm.addr(1))));

    s_nonceManager.applyPreviousRampsUpdates(previousRamps);

    previousRamps[0] =
      NonceManager.PreviousRampsArgs(DEST_CHAIN_SELECTOR, NonceManager.PreviousRamps(address(vm.addr(2))));

    vm.expectRevert(NonceManager.PreviousRampAlreadySet.selector);
    s_nonceManager.applyPreviousRampsUpdates(previousRamps);
  }

  function _assertPreviousRampsEqual(
    NonceManager.PreviousRamps memory a,
    NonceManager.PreviousRamps memory b
  ) internal pure {
    assertEq(a.prevOnRamp, b.prevOnRamp);
  }
}

// TODO: move upgradability tests for both OnRmap and OffRamp here
