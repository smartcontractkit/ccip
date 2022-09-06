pragma solidity ^0.8.0;

import "../../offRamp/BaseOffRamp.sol";

contract BaseOffRampHelper is BaseOffRamp {
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
    BaseOffRamp(
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

  function setExecutionState(uint64 sequenceNumber, CCIP.MessageExecutionState state) public {
    s_executedMessages[sequenceNumber] = state;
  }

  function releaseOrMintToken(
    PoolInterface pool,
    uint256 amount,
    address receiver
  ) external {
    _releaseOrMintToken(pool, amount, receiver);
  }

  function releaseOrMintTokens(
    PoolInterface[] memory tokens,
    uint256[] memory amounts,
    address receiver
  ) external {
    _releaseOrMintTokens(tokens, amounts, receiver);
  }

  function verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] memory outerProofs,
    uint256 outerProofFlagBits
  ) external returns (uint256, uint256) {
    return _verifyMessages(hashedLeaves, innerProofs, innerProofFlagBits, outerProofs, outerProofFlagBits);
  }

  function getPool_helper(IERC20 token) external view returns (PoolInterface pool) {
    return _getPool(token);
  }
}
