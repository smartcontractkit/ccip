// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/EVM2EVMOffRamp.sol";

contract EVM2EVMOffRampHelper is EVM2EVMOffRamp {
  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    OffRampConfig memory offRampConfig,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  )
    EVM2EVMOffRamp(
      sourceChainId,
      chainId,
      offRampConfig,
      onRampAddress,
      commitStore,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig
    )
  {}

  function report(bytes memory executableMessages) external {
    _report(executableMessages);
  }

  function execute(Internal.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(Internal.EVM_2_EVM_MESSAGE_HASH);
  }
}
