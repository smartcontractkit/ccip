// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../models/Models.sol";
import "./Any2EVMOffRampRouterInterface.sol";
import "../BlobVerifierInterface.sol";

interface BaseOffRampInterface {
  error AlreadyExecuted(uint64 sequenceNumber);
  error ExecutionError();
  error InvalidSourceChain(uint256 sourceChainId);
  error NoMessagesToExecute();
  error ManualExecutionNotYetEnabled();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error RouterNotSet();
  error RootNotRelayed();
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error TokenAndAmountMisMatch();
  error UnsupportedToken(IERC20 token);
  error CanOnlySelfCall();
  error ReceiverError();
  error MissingFeeCoinPrice(address feeCoin);
  error InsufficientFeeAmount(uint256 sequenceNumber, uint256 expectedFeeTokens, uint256 feeTokenAmount);
  error IncorrectNonce(uint64 nonce);

  event ExecutionStateChanged(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);
  event OffRampRouterSet(address indexed router);
  event OffRampConfigSet(OffRampConfig config);

  struct OffRampConfig {
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
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner should be able to call this function
   */
  function setRouter(Any2EVMOffRampRouterInterface router) external;

  /**
   * @notice get the current router
   * @return Any2EVMOffRampRouterInterface
   */
  function getRouter() external view returns (Any2EVMOffRampRouterInterface);

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param manualExecution Whether or not it is manual or DON execution
   */
  function execute(CCIP.ExecutionReport memory report, bool manualExecution) external;

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
