// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/ge/EVM2EVMGEOffRamp.sol";

contract EVM2EVMGEOffRampHelper is EVM2EVMGEOffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    GEOffRampConfig memory offRampConfig,
    address onRampAddress,
    CommitStoreInterface commitStore,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    IERC20 feeToken
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
      rateLimiterConfig,
      tokenLimitsAdmin,
      feeToken
    )
  {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }

  function execute(GE.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(GE.EVM_2_EVM_GE_MESSAGE_HASH);
  }
}
