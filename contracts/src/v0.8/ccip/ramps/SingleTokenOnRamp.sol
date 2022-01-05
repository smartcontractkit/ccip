// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/OnRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../utils/TokenLimits.sol";
import "../health/HealthChecker.sol";

/**
 * @notice An implementation of an On Ramp, which enables just a single token to be
 * used in the protocol.
 */
contract SingleTokenOnRamp is OnRampInterface, TypeAndVersionInterface, HealthChecker {
  using TokenLimits for TokenLimits.TokenBucket;

  // Chain ID of the destination chain. This is sent in the request to the DON.
  uint256 public immutable DESTINATION_CHAIN_ID;
  // Address of the token on the destination chain. This is sent in the request to the DON.
  IERC20 public immutable DESTINATION_TOKEN;

  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable CHAIN_ID;
  // Token pool responsible for managing the TOKEN.
  PoolInterface public immutable POOL;
  // Token that this ramp enables to be sent using the protocol.
  IERC20 public immutable TOKEN;

  // Whether the allowlist is enabled
  bool private s_allowlistEnabled;
  // Addresses that are allowed to send messages
  mapping(address => bool) private s_allowed;
  // List of allowed addresses
  address[] private s_allowList;
  // Simple incremental nonce.
  uint256 private s_sequenceNumber;
  // Token bucket for token rate limiting
  TokenLimits.TokenBucket private s_tokenBucket;

  constructor(
    uint256 sourceChainId,
    IERC20 sourceToken,
    PoolInterface sourcePool,
    uint256 destinationChainId,
    IERC20 destinationToken,
    address[] memory allowlist,
    bool enableAllowlist,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal
  ) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    if (sourcePool.getToken() != sourceToken) revert TokenMismatch();
    CHAIN_ID = sourceChainId;
    TOKEN = sourceToken;
    POOL = sourcePool;
    DESTINATION_CHAIN_ID = destinationChainId;
    DESTINATION_TOKEN = destinationToken;
    s_sequenceNumber = 1;
    s_allowlistEnabled = enableAllowlist;
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    s_tokenBucket = TokenLimits.constructTokenBucket(tokenBucketRate, tokenBucketCapacity, true);
  }

  /**
   * @notice Send a message to the remote chain
   * @dev tokens must be of length 1 and be the token allowed by this contract
   * @dev amounts must also be of length 1, be greater than zero, and approve() must have already
   * been called on the token using the POOL address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param payload Message struct to send
   */
  function requestCrossChainSend(CCIP.MessagePayload memory payload)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint256)
  {
    address sender = msg.sender;
    if (s_allowlistEnabled && !s_allowed[sender]) revert SenderNotAllowed(sender);
    // Check that inputs are correct
    if (payload.tokens.length != 1 || payload.amounts.length != 1) revert UnsupportedNumberOfTokens();
    if (payload.tokens[0] != TOKEN) revert UnsupportedToken(TOKEN, payload.tokens[0]);
    // This step will be a mapping filled with a loop in future when more than one token is suported.
    IERC20[] memory mappedRemoteTokens = new IERC20[](1);
    mappedRemoteTokens[0] = DESTINATION_TOKEN;
    payload.tokens = mappedRemoteTokens;
    // Check that sending these tokens falls within the bucket limits
    if (!s_tokenBucket.remove(payload.amounts[0])) revert ExceedsTokenLimit(s_tokenBucket.tokens, payload.amounts[0]);
    // Store in pool
    POOL.lockOrBurn(sender, payload.amounts[0]);
    // Emit message request
    CCIP.Message memory message = CCIP.Message({
      sequenceNumber: s_sequenceNumber,
      sourceChainId: CHAIN_ID,
      destinationChainId: DESTINATION_CHAIN_ID,
      sender: sender,
      payload: payload
    });
    emit CrossChainSendRequested(message);
    s_sequenceNumber++;
    return message.sequenceNumber;
  }

  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowlistEnabledSet(enabled);
  }

  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  function setAllowlist(address[] calldata allowlist) external onlyOwner {
    // Remove existing allowlist
    address[] memory existingList = s_allowList;
    for (uint256 i = 0; i < existingList.length; i++) {
      s_allowed[existingList[i]] = false;
    }

    // Set the new allowlist
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    emit AllowlistSet(allowlist);
  }

  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }

  function configureTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external onlyOwner {
    s_tokenBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewTokenBucketConstructed(rate, capacity, full);
  }

  function getTokenBucket() external view returns (TokenLimits.TokenBucket memory) {
    return s_tokenBucket;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "SingleTokenOnRamp 1.1.0";
  }
}
