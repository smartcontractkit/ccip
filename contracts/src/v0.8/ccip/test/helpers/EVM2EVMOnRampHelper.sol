// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../../onRamp/EVM2EVMOnRamp.sol";
import {IgnoreContractSize} from "./IgnoreContractSize.sol";

contract EVM2EVMOnRampHelper is EVM2EVMOnRamp, IgnoreContractSize {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    Internal.PoolUpdate[] memory tokensAndPools,
    RateLimiter.Config memory rateLimiterConfig,
    FeeTokenConfigArgs[] memory feeTokenConfigs,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    NopAndWeight[] memory nopsAndWeights
  )
    EVM2EVMOnRamp(
      staticConfig,
      dynamicConfig,
      tokensAndPools,
      rateLimiterConfig,
      feeTokenConfigs,
      tokenTransferFeeConfigArgs,
      nopsAndWeights
    )
  {}

  function getDataAvailabilityCostUSD(
    uint112 dataAvailabilityGasPrice,
    uint256 messageDataLength,
    uint256 numberOfTokens,
    uint32 tokenTransferBytesOverhead
  ) external view returns (uint256) {
    return
      _getDataAvailabilityCostUSD(
        dataAvailabilityGasPrice,
        messageDataLength,
        numberOfTokens,
        tokenTransferBytesOverhead
      );
  }

  function getTokenTransferCost(
    address feeToken,
    uint224 feeTokenPrice,
    Client.EVMTokenAmount[] calldata tokenAmounts,
    FeeTokenConfig memory feeTokenConfig
  ) external view returns (uint256, uint32, uint32) {
    return _getTokenTransferCost(feeToken, feeTokenPrice, tokenAmounts, feeTokenConfig);
  }
}
