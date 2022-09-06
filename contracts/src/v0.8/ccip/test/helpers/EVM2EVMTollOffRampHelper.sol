// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/toll/EVM2EVMTollOffRamp.sol";

contract EVM2EVMTollOffRampHelper is EVM2EVMTollOffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    address onRampAddress,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    EVM2EVMTollOffRamp(
      sourceChainId,
      chainId,
      offRampConfig,
      blobVerifier,
      onRampAddress,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig,
      tokenLimitsAdmin
    )
  {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
