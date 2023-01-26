// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {CCIPConsumer} from "../../../applications/CCIPConsumer.sol";
import {EVM2EVMGEOffRamp} from "../../../offRamp/ge/EVM2EVMGEOffRamp.sol";
import {Common} from "../../../models/Common.sol";
import {GE} from "../../../models/GE.sol";

contract ReentrancyAbuser is CCIPConsumer {
  event ReentrancySucceeded();

  bool s_ReentrancyDone = false;
  GE.ExecutionReport s_payload;
  EVM2EVMGEOffRamp s_offRamp;

  constructor(
    address router,
    address feeToken,
    EVM2EVMGEOffRamp offRamp
  ) CCIPConsumer(router, feeToken) {
    s_offRamp = offRamp;
  }

  function setPayload(GE.ExecutionReport calldata payload) public {
    s_payload = payload;
  }

  function _ccipReceive(Common.Any2EVMMessage memory) internal override {
    if (!s_ReentrancyDone) {
      // Could do more rounds but a PoC one is enough
      s_ReentrancyDone = true;
      s_offRamp.manuallyExecute(s_payload);
    } else {
      emit ReentrancySucceeded();
    }
  }
}
