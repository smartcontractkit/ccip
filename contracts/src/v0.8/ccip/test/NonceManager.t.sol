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

// TODO: move upgradability tests for both OnRmap and OffRamp here
