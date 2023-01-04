// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/toll/EVM2EVMTollOffRamp.sol";

contract EVM2EVMTollOffRampHelper is EVM2EVMTollOffRamp {
  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    OffRampConfig memory offRampConfig,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    EVM2EVMTollOffRamp(
      sourceChainId,
      chainId,
      offRampConfig,
      onRampAddress,
      commitStore,
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

  function execute(Toll.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(Toll.EVM_2_EVM_TOLL_MESSAGE_HASH);
  }
}
