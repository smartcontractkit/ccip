// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/ge/EVM2EVMGEOffRamp.sol";

contract EVM2EVMGEOffRampHelper is EVM2EVMGEOffRamp {
  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    GEOffRampConfig memory offRampConfig,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  )
    EVM2EVMGEOffRamp(
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

  function execute(GE.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(GE.EVM_2_EVM_GE_MESSAGE_HASH);
  }
}
