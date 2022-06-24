// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/toll/Any2EVMTollOffRamp.sol";

contract Any2EVMTollOffRampHelper is Any2EVMTollOffRamp {
  constructor(
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    address onRampAddress,
    AFNInterface afn,
    // TODO token limiter
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    uint256 maxTimeWithoutAFNSignal
  )
    Any2EVMTollOffRamp(
      chainId,
      offRampConfig,
      blobVerifier,
      onRampAddress,
      afn,
      sourceTokens,
      pools,
      maxTimeWithoutAFNSignal
    )
  {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
