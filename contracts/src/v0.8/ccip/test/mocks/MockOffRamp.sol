// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../applications/interfaces/CrossChainMessageReceiverInterface.sol";
import "../../offRamp/interfaces/Any2EVMTollOffRampInterface.sol";

contract MockOffRamp is Any2EVMTollOffRampInterface {
  IERC20 public s_token;

  function deliverMessageTo(CrossChainMessageReceiverInterface recipient, CCIP.Any2EVMTollMessage calldata message)
    external
  {
    recipient.ccipReceive(message);
  }

  function SOURCE_CHAIN_ID() external view returns (uint256) {}

  function CHAIN_ID() external view returns (uint256) {}

  function setRouter(Any2EVMTollOffRampRouterInterface router) external {}

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata) external pure override {
    revert();
  }

  function execute(CCIP.ExecutionReport memory report, bool needFee) external override {}

  function executeSingleMessage(CCIP.Any2EVMTollMessage memory message) external {}

  function setToken(IERC20 token) external {
    s_token = token;
  }

  function TOKEN() external view returns (IERC20) {
    return s_token;
  }

  /// @inheritdoc BaseOffRampInterface
  function getExecutionState(uint64) public pure returns (CCIP.MessageExecutionState) {
    return CCIP.MessageExecutionState.Success;
  }

  /// @inheritdoc BaseOffRampInterface
  function getBlobVerifier() public pure returns (BlobVerifierInterface) {
    return BlobVerifierInterface(address(1));
  }

  /// @inheritdoc BaseOffRampInterface
  function setBlobVerifier(BlobVerifierInterface blobVerifier) public pure {}

  /// @inheritdoc BaseOffRampInterface
  function getConfig() public pure returns (OffRampConfig memory config) {}

  /// @inheritdoc BaseOffRampInterface
  function setConfig(OffRampConfig memory config) public {}
}
