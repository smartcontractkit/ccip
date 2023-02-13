// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {CCIPReceiver} from "../../../applications/CCIPReceiver.sol";
import {Common} from "../../../models/Common.sol";

contract ConformingReceiver is CCIPReceiver {
  event MessageReceived();

  constructor(address router, address feeToken) CCIPReceiver(router) {}

  function _ccipReceive(Common.Any2EVMMessage memory) internal virtual override {
    emit MessageReceived();
  }
}
