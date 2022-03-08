// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../interfaces/OffRampInterface.sol";

contract SimpleMessageReceiver is CrossChainMessageReceiverInterface {
  CCIP.Message public s_message;

  event MessageReceived(CCIP.Message message);

  function receiveMessage(CCIP.Message calldata message) external override {
    s_message = message;
    emit MessageReceived(message);
  }
}
