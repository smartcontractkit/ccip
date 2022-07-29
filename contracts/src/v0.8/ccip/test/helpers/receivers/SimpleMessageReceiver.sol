// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../applications/interfaces/CrossChainMessageReceiverInterface.sol";
import "../../../utils/interfaces/SubscriptionManagerInterface.sol";

contract SimpleMessageReceiver is CrossChainMessageReceiverInterface {
  event MessageReceived(uint256 message);

  address s_manager;

  constructor() {
    s_manager = msg.sender;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_manager;
  }

  function ccipReceive(CCIP.Any2EVMSubscriptionMessage calldata message) external override {
    emit MessageReceived(message.sequenceNumber);
  }

  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external override {
    emit MessageReceived(message.sequenceNumber);
  }
}
