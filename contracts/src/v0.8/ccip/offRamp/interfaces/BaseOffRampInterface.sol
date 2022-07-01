// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../utils/CCIP.sol";
import "../../blobVerifier/interfaces/BlobVerifierInterface.sol";

interface BaseOffRampInterface {
  error AlreadyExecuted(uint64 sequenceNumber);
  error CanOnlySelfCall();
  error ExecutionError(uint64 sequenceNumber, bytes reason);
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);
  error NoMessagesToExecute();
  error ManualExecutionNotYetEnabled();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error RouterNotSet();
  error RootNotRelayed();
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error TokenAndAmountMisMatch();
  error UnsupportedToken(IERC20 token);

  event ExecutionCompleted(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);
  event OffRampRouterSet(address indexed router);
  event OffRampConfigSet(OffRampConfig config);

  struct OffRampConfig {
    // The ID of the source chain
    uint256 sourceChainId;
    // execution delay in seconds
    uint64 executionDelaySeconds;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
    // The waiting time before manual execution is enabled
    uint32 permissionLessExecutionThresholdSeconds;
  }

  /**
   * @notice Returns the current execution state of a message based on its
   *          sequenceNumber.
   */
  function getExecutionState(uint64 sequenceNumber) external view returns (CCIP.MessageExecutionState);

  /**
   * @notice Returns the current blob verifier.
   */
  function getBlobVerifier() external view returns (BlobVerifierInterface);

  /**
   * @notice Updates the blobVerifier.
   * @param blobVerifier The new blobVerifier
   */
  function setBlobVerifier(BlobVerifierInterface blobVerifier) external;

  /**
   * @notice Returns the current config.
   */
  function getConfig() external view returns (OffRampConfig memory);

  /**
   * @notice Sets a new config.
   */
  function setConfig(OffRampConfig memory config) external;
}
