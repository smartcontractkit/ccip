// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";
import {EVM2EVMOffRamp} from "../../../offRamp/EVM2EVMOffRamp.sol";

contract ReentrancyAbuser is CCIPReceiver {
  event ReentrancySucceeded();

  bool internal s_ReentrancyDone = false;
  Internal.ExecutionReport internal s_payload;
  EVM2EVMOffRamp internal s_offRamp;

  constructor(address router, EVM2EVMOffRamp offRamp) CCIPReceiver(router) {
    s_offRamp = offRamp;
  }

  function setPayload(Internal.ExecutionReport calldata payload) public {
    s_payload = payload;
  }

  function _ccipReceive(Client.Any2EVMMessage memory) internal override {
    // Use original message gas limits in manual execution
    uint256 numMsgs = s_payload.messages.length;
    EVM2EVMOffRamp.GasLimitOverride[] memory gasOverrides = new EVM2EVMOffRamp.GasLimitOverride[](numMsgs);

    if (!s_ReentrancyDone) {
      // Could do more rounds but a PoC one is enough
      s_ReentrancyDone = true;
      s_offRamp.manuallyExecute(s_payload, gasOverrides);
    } else {
      emit ReentrancySucceeded();
    }
  }
}
