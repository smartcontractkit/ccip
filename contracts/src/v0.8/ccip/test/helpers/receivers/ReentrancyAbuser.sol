// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";
import {EVM2EVMOffRamp} from "../../../offRamp/EVM2EVMOffRamp.sol";
import {Client} from "../../../libraries/Client.sol";
import {Internal} from "../../../libraries/Internal.sol";

contract ReentrancyAbuser is CCIPReceiver {
  event ReentrancySucceeded();

  bool s_ReentrancyDone = false;
  Internal.ExecutionReport s_payload;
  EVM2EVMOffRamp s_offRamp;

  constructor(address router, EVM2EVMOffRamp offRamp) CCIPReceiver(router) {
    s_offRamp = offRamp;
  }

  function setPayload(Internal.ExecutionReport calldata payload) public {
    s_payload = payload;
  }

  function _ccipReceive(Client.Any2EVMMessage memory) internal override {
    if (!s_ReentrancyDone) {
      // Could do more rounds but a PoC one is enough
      s_ReentrancyDone = true;
      s_offRamp.manuallyExecute(s_payload);
    } else {
      emit ReentrancySucceeded();
    }
  }
}
