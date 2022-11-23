// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/applications/Any2EVMMessageReceiverInterface.sol";

contract SimpleMessageReceiver is Any2EVMMessageReceiverInterface {
  event MessageReceived();

  address public s_manager;

  constructor() {
    s_manager = msg.sender;
  }

  function ccipReceive(CCIP.Any2EVMMessage calldata) external override {
    emit MessageReceived();
  }
}
