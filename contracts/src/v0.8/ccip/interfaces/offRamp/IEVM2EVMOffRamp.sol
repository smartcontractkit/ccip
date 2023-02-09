// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../fees/IFeeManager.sol";

import {Internal} from "../../models/Internal.sol";
import {Internal} from "../../models/Internal.sol";

interface IEVM2EVMOffRamp {
  error UnauthorizedGasPriceUpdate();
  error AlreadyAttempted(uint64 seqNum);

  event OffRampConfigChanged(OffRampConfig);
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);
  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state
  );

  struct OffRampConfig {
    IFeeManager feeManager;
    // The waiting time before manual execution is enabled
    uint32 permissionLessExecutionThresholdSeconds;
    // execution delay in seconds
    uint64 executionDelaySeconds;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
  }

  /**
   * @notice Returns the the current nonce for a receiver.
   */
  function getSenderNonce(address sender) external view returns (uint64 nonce);

  /**
   * @notice Returns the current config.
   */
  function getOffRampConfig() external view returns (OffRampConfig memory);

  /**
   * @notice Sets a new config.
   */
  function setOffRampConfig(OffRampConfig memory config) external;

  /**
   * @notice Manually execute a message
   * @param report Internal.ExecutionReport
   */
  function manuallyExecute(Internal.ExecutionReport memory report) external;
}
