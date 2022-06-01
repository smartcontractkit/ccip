// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";

contract SimpleMessageReceiver is CrossChainMessageReceiverInterface {
  CCIP.AnyToEVMTollMessage public s_message;

  event MessageReceived(CCIP.AnyToEVMTollMessage message);

  function receiveMessage(CCIP.AnyToEVMTollMessage calldata message) external override {
    s_message = message;
    emit MessageReceived(message);
  }

  /**
   * @dev Exposes s_message fully as the public variable does not include arrays
   * so it will be missing the properties `tokens` and `amounts`.
   * https://docs.soliditylang.org/en/v0.8.2/contracts.html#getter-functions
   */
  function getMessage() external view returns (CCIP.AnyToEVMTollMessage memory msg) {
    return s_message;
  }
}
