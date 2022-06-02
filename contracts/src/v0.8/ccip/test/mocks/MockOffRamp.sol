// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../interfaces/TollOffRampInterface.sol";

contract MockOffRamp is TollOffRampInterface {
  event MessageExecuted(CCIP.ExecutionReport report, bool needFee);

  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Any2EVMTollMessage calldata message)
    external
  {
    recipient.receiveMessage(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function executeTransaction(CCIP.ExecutionReport memory report, bool needFee) external override {
    emit MessageExecuted(report, needFee);
  }

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }
}
