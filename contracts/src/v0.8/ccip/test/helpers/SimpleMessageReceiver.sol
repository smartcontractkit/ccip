// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";

contract SimpleMessageReceiver is CrossChainMessageReceiverInterface {
  event MessageReceived(uint256 message);

  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external override {
    emit MessageReceived(message.sequenceNumber);
  }
}
