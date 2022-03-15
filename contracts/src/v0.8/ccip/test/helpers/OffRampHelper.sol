// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../../ramps/OffRamp.sol";

contract OffRampHelper is OffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint64 executionDelaySeconds,
    uint64 maxTokensLength
  )
    OffRamp(
      sourceChainId,
      chainId,
      sourceTokens,
      pools,
      feeds,
      afn,
      maxTimeWithoutAFNSignal,
      OffRampConfig({
        executionFeeJuels: 1,
        executionDelaySeconds: executionDelaySeconds,
        maxDataSize: 1000,
        maxTokensLength: maxTokensLength
      })
    )
  {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory merkle) external {
    _report(bytes32(0), 0, merkle);
  }
}
