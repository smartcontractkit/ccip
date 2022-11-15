// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/subscription/EVM2EVMFreeOffRamp.sol";

contract EVM2EVMFreeOffRampHelper is EVM2EVMFreeOffRamp {
  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    EVM2EVMFreeOffRamp(
      sourceChainId,
      chainId,
      offRampConfig,
      blobVerifier,
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

  function execute(CCIP.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function setMessageState(uint64 sequenceNumber, CCIP.MessageExecutionState state) external {
    s_executedMessages[sequenceNumber] = state;
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(CCIP.EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH);
  }
}
