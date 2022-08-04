// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/applications/Any2EVMMessageReceiverInterface.sol";

contract RevertingMessageReceiver is Any2EVMMessageReceiverInterface {
  address s_manager;

  constructor() {
    s_manager = msg.sender;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_manager;
  }

  function ccipReceive(CCIP.Any2EVMMessage calldata) external pure override {
    revert();
  }
}
