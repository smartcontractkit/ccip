// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../../onRamp/EVM2EVMOnRamp.sol";

contract EVM2EVMOnRampHelper is EVM2EVMOnRamp {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    TokenAndPool[] memory tokensAndPools,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig,
    FeeTokenConfigArgs[] memory feeTokenConfigs,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    NopAndWeight[] memory nopsAndWeights
  )
    EVM2EVMOnRamp(
      staticConfig,
      dynamicConfig,
      tokensAndPools,
      allowlist,
      rateLimiterConfig,
      feeTokenConfigs,
      tokenTransferFeeConfigArgs,
      nopsAndWeights
    )
  {}

  function getMessageExecutionFee(address feeToken, bytes calldata extraArgs) public view returns (uint256) {
    return _getMessageExecutionFee(feeToken, extraArgs);
  }

  function getTokenTransferFee(
    address feeToken,
    Client.EVMTokenAmount[] calldata tokenAmounts
  ) public view returns (uint256) {
    return _getTokenTransferFee(feeToken, tokenAmounts);
  }
}
