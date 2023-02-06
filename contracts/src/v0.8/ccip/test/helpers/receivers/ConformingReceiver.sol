// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {CCIPConsumer} from "../../../applications/CCIPConsumer.sol";
import {Common} from "../../../models/Common.sol";

contract ConformingReceiver is CCIPConsumer {
  event MessageReceived();

  constructor(address router, address feeToken) CCIPConsumer(router) {}

  function _ccipReceive(Common.Any2EVMMessage memory) internal virtual override {
    emit MessageReceived();
  }
}
