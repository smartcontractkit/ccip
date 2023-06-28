// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "../../offRamp/EVM2EVMOffRamp.sol";

contract EVM2EVMOffRampHelper is EVM2EVMOffRamp {
  constructor(
    StaticConfig memory staticConfig,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiter.Config memory rateLimiterConfig
  ) EVM2EVMOffRamp(staticConfig, sourceTokens, pools, rateLimiterConfig) {}

  function setExecutionStateHelper(uint64 sequenceNumber, Internal.MessageExecutionState state) public {
    _setExecutionState(sequenceNumber, state);
  }

  function getExecutionStateBitMap(uint64 bitmapIndex) public view returns (uint256) {
    return s_executionStates[bitmapIndex];
  }

  function releaseOrMintTokens(
    Client.EVMTokenAmount[] memory sourceTokenAmounts,
    bytes calldata originalSender,
    address receiver,
    bytes[] calldata offchainTokenData
  ) external returns (Client.EVMTokenAmount[] memory) {
    return _releaseOrMintTokens(sourceTokenAmounts, originalSender, receiver, offchainTokenData);
  }

  function report(bytes calldata executableMessages) external {
    _report(executableMessages);
  }

  function execute(Internal.ExecutionReport memory rep, bool manualExecution) external {
    _execute(rep, manualExecution);
  }

  function metadataHash() external view returns (bytes32) {
    return _metadataHash(Internal.EVM_2_EVM_MESSAGE_HASH);
  }
}
