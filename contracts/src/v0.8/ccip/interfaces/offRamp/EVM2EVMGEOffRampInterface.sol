// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {GasFeeCacheInterface} from "../dynamicFeeCalculator/GasFeeCacheInterfaceInterface.sol";
import {CCIP} from "../../models/Models.sol";

interface EVM2EVMGEOffRampInterface {
  error UnauthorizedGasPriceUpdate();
  error AlreadyAttempted(uint64 seqNum);

  event GEOffRampConfigChanged(GEOffRampConfig);
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);
  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    CCIP.MessageExecutionState state
  );

  struct GEOffRampConfig {
    uint256 gasOverhead;
    GasFeeCacheInterface gasFeeCache;
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
  function getSenderNonce(address sender) external view returns (uint256 nonce);

  /**
   * @notice Returns the current balance for a given NOP.
   */
  function getNopBalance(address nop) external view returns (uint256 balance);

  /**
   * @notice Returns the current config.
   */
  function getGEConfig() external view returns (GEOffRampConfig memory);

  /**
   * @notice Sets a new config.
   */
  function setGEConfig(GEOffRampConfig memory config) external;
}
