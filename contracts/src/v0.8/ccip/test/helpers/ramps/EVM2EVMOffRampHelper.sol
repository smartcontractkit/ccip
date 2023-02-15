// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../offRamp/EVM2EVMOffRamp.sol";

contract EVM2EVMOffRampHelper is EVM2EVMOffRamp {
  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    address onRampAddress,
    OffRampConfig memory offRampConfig,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  ) EVM2EVMOffRamp(sourceChainId, chainId, onRampAddress, offRampConfig, afn, sourceTokens, pools, rateLimiterConfig) {}

  function setExecutionState(uint64 sequenceNumber, Internal.MessageExecutionState state) public {
    s_executedMessages[sequenceNumber] = state;
  }

  function releaseOrMintToken(
    IPool pool,
    uint256 amount,
    address receiver
  ) external {
    _releaseOrMintToken(pool, amount, receiver);
  }

  function releaseOrMintTokens(Client.EVMTokenAmount[] memory sourceTokenAmounts, address receiver)
    external
    returns (Client.EVMTokenAmount[] memory)
  {
    return _releaseOrMintTokens(sourceTokenAmounts, receiver);
  }

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
