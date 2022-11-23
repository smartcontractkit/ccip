// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {GasFeeCacheInterface} from "../dynamicFeeCalculator/GasFeeCacheInterfaceInterface.sol";

interface EVM2EVMGEOffRampInterface {
  error UnauthorizedGasPriceUpdate();
  error AlreadyAttempted(uint64 seqNum);

  event GEOffRampConfigChanged(GEOffRampConfig);

  struct GEOffRampConfig {
    address feeTokenDest;
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
   * @notice Returns the current config.
   */
  function getGEConfig() external view returns (GEOffRampConfig memory);

  /**
   * @notice Sets a new config.
   */
  function setGEConfig(GEOffRampConfig memory config) external;
}
