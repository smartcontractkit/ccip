// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Client} from "../../../libraries/Client.sol";
import {CCIPReceiverBasic} from "./CCIPReceiverBasic.sol";

contract ConformingReceiver is CCIPReceiverBasic {
  event MessageReceived();

  constructor(address router, address feeToken) CCIPReceiverBasic(router) {}

  function _ccipReceive(Client.Any2EVMMessage memory) internal virtual override {
    emit MessageReceived();
  }
}
