// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../interfaces/OffRampInterface.sol";

contract MockOffRamp is OffRampInterface {
  event MessageExecuted(bytes32[] proof, CCIP.Message message, uint256 index);

  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Message calldata message) external {
    recipient.receiveMessage(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external override {
    emit MessageExecuted(proof, message, index);
  }

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }
}
