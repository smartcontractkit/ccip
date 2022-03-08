// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../../interfaces/OffRampInterface.sol";

contract MockOffRamp is OffRampInterface {
  event MessageExecuted(bytes32[] path, uint256 index, CCIP.Message message, bool needFee);

  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Message calldata message) external {
    recipient.receiveMessage(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function executeTransaction(
    CCIP.Message memory message,
    CCIP.MerkleProof memory proof,
    bool needFee
  ) external override {
    emit MessageExecuted(proof.path, proof.index, message, needFee);
  }

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }
}
