// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../../models/Internal.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IEVM2EVMOffRamp {
  error AlreadyAttempted(uint64 sequenceNumber);
  error AlreadyExecuted(uint64 sequenceNumber);
  error ZeroAddressNotAllowed();
  error ExecutionError(bytes error);
  error InvalidSourceChain(uint64 sourceChainId);
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error ManualExecutionNotYetEnabled();
  error RootNotCommitted();
  error InvalidOffRampConfig(DynamicConfig config);
  error UnsupportedToken(IERC20 token);
  error CanOnlySelfCall();
  error ReceiverError();
  error EmptyReport();

  // sourceChainId and onRamp are needed by Atlas, to track onramp <-> offramp -> router relationship
  event DynamicConfigSet(DynamicConfig config, uint64 sourceChainId, address onRamp);
  event StaticConfigSet(StaticConfig);
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);
  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state
  );

  /// @notice Returns the current execution state of a message based on its sequenceNumber.
  /// @param sequenceNumber The sequence number of the message to get the execution state for
  /// @return The current execution state of the message
  function getExecutionState(uint64 sequenceNumber) external view returns (Internal.MessageExecutionState);

  /// @notice Returns the the current nonce for a receiver.
  /// @param sender The sender address
  /// @return nonce The nonce value belonging to the sender address.
  function getSenderNonce(address sender) external view returns (uint64 nonce);

  /// @notice Static offRamp config
  struct StaticConfig {
    address commitStore; // --┐  CommitStore address on the destination chain
    uint64 chainId; // -------┘  Destination chain Id
    uint64 sourceChainId; // -┐  Source chain Id
    address onRamp; // -------┘  OnRamp address on the source chain
  }

  /// @notice Returns the static config.
  /// @dev This function will always return the same struct as the contents is static and can never change.
  function getStaticConfig() external view returns (StaticConfig memory);

  /// @notice Dynamic offRamp config
  /// @dev since OffRampConfig is part of OffRampConfigChanged event, if changing it, we should update the ABI on Atlas
  struct DynamicConfig {
    uint32 permissionLessExecutionThresholdSeconds; // -┐ Waiting time before manual execution is enabled
    uint64 executionDelaySeconds; //                    | Execution delay in seconds
    address router; // ---------------------------------┘ Router address
    uint32 maxDataSize; // --------┐ Maximum payload data size
    uint16 maxTokensLength; // ----┘ Maximum number of distinct ERC20 tokens that can be sent per message
  }

  /// @notice Returns the current dynamic config.
  /// @return The current config.
  function getDynamicConfig() external view returns (DynamicConfig memory);

  /// @notice Sets a new dynamic config.
  /// @param config The new config
  function setDynamicConfig(DynamicConfig memory config) external;

  /// @notice Manually execute a message.
  /// @param report Internal.ExecutionReport.
  function manuallyExecute(Internal.ExecutionReport memory report) external;
}
