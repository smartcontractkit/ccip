// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../fees/IFeeManager.sol";
import {ICommitStore} from "../ICommitStore.sol";

import {Internal} from "../../models/Internal.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IEVM2EVMOffRamp {
  error UnauthorizedGasPriceUpdate();
  error AlreadyAttempted(uint64 sequenceNumber);
  error AlreadyExecuted(uint64 sequenceNumber);
  error ZeroAddressNotAllowed();
  error ExecutionError(bytes error);
  error InvalidSourceChain(uint64 sourceChainId);
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error ManualExecutionNotYetEnabled();
  error RootNotCommitted();
  error InvalidOffRampConfig(OffRampConfig config);
  error UnsupportedToken(IERC20 token);
  error CanOnlySelfCall();
  error ReceiverError();

  // sourceChainId and onRamp are needed by Atlas, to track onramp <-> offramp -> router relationship
  event OffRampConfigChanged(OffRampConfig config, uint64 sourceChainId, address onRamp);
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);
  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state
  );

  /// @notice Returns the current execution state of a message based on its
  ///          sequenceNumber.
  function getExecutionState(uint64 sequenceNumber) external view returns (Internal.MessageExecutionState);

  /// @notice Returns the the current nonce for a receiver.
  function getSenderNonce(address sender) external view returns (uint64 nonce);

  // since OffRampConfig is part of OffRampConfigChanged event, if changing it, we should update the ABI on Atlas
  struct OffRampConfig {
    address feeManager;
    // The waiting time before manual execution is enabled
    uint32 permissionLessExecutionThresholdSeconds;
    // execution delay in seconds
    uint64 executionDelaySeconds;
    address router;
    // maximum payload data size
    uint64 maxDataSize;
    address commitStore;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
  }

  /// @notice Returns the current config.
  function getOffRampConfig() external view returns (OffRampConfig memory);

  /// @notice Sets a new config.
  function setOffRampConfig(OffRampConfig memory config) external;

  /// @notice Manually execute a message
  /// @param report Internal.ExecutionReport
  function manuallyExecute(Internal.ExecutionReport memory report) external;
}
