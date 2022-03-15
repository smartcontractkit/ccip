// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../../ramps/OnRamp.sol";

contract OnRampHelper is OnRamp {
  constructor(
    uint256 chainId,
    uint256[] memory destinationChainIds,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint64 maxTokensLength,
    uint64 maxDataSize,
    uint64 relayingFeeJuels
  )
    OnRamp(
      chainId,
      destinationChainIds,
      tokens,
      pools,
      feeds,
      allowlist,
      afn,
      maxTimeWithoutAFNSignal,
      OnRampConfig({maxTokensLength: maxTokensLength, maxDataSize: maxDataSize, relayingFeeJuels: relayingFeeJuels})
    )
  {}

  function publicCalculateFee(IERC20 feeToken) external view returns (uint256) {
    return _calculateFee(feeToken);
  }
}
