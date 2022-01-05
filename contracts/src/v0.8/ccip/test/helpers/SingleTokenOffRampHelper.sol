// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../../ramps/SingleTokenOffRamp.sol";

contract SingleTokenOffRampHelper is SingleTokenOffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 destinationChainId,
    IERC20 token,
    PoolInterface pool,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint256 executionDelaySeconds
  )
    SingleTokenOffRamp(
      sourceChainId,
      destinationChainId,
      token,
      pool,
      tokenBucketRate,
      tokenBucketCapacity,
      afn,
      maxTimeWithoutAFNSignal,
      executionDelaySeconds
    )
  {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory merkle) external {
    _report(bytes32(0), 0, merkle);
  }
}
