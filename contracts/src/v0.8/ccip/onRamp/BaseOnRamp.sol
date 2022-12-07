// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {HealthChecker, AFNInterface} from "../health/HealthChecker.sol";
import {IERC20} from "../pools/PoolCollector.sol";
import {AllowList} from "../access/AllowList.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {BaseOnRampInterface, PoolInterface} from "../interfaces/onRamp/BaseOnRampInterface.sol";
import {Common} from "../models/Common.sol";

contract BaseOnRamp is BaseOnRampInterface, HealthChecker, AllowList, AggregateRateLimiter {
  // Chain ID of the source chain (where this contract is deployed)
  uint64 public immutable i_chainId;
  // Chain ID of the destination chain (where this contract sends messages)
  uint64 public immutable i_destinationChainId;

  // The last used sequence number. This is zero in the case where no
  // messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  uint64 internal s_sequenceNumber;

  // The current configuration of the onRamp.
  OnRampConfig internal s_config;
  // The router that is allowed to interact with this onRamp.
  address internal s_router;

  // source token => token pool
  mapping(IERC20 => PoolConfig) private s_poolsBySourceToken;
  IERC20[] private s_sourceTokenList;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    address[] memory allowlist,
    AFNInterface afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    address router
  ) HealthChecker(afn) AllowList(allowlist) AggregateRateLimiter(rateLimiterConfig, tokenLimitsAdmin) {
    i_chainId = chainId;
    i_destinationChainId = destinationChainId;
    s_config = config;
    s_router = router;
    s_sequenceNumber = 0;

    if (tokens.length != pools.length) revert InvalidTokenPoolConfig();
    s_sourceTokenList = tokens;
    // Set new tokens and pools
    for (uint256 i = 0; i < tokens.length; ++i) {
      s_poolsBySourceToken[tokens[i]] = PoolConfig({pool: pools[i], enabled: true});
    }
  }

  /// @inheritdoc BaseOnRampInterface
  function getPoolBySourceToken(IERC20 sourceToken) public view override returns (PoolInterface) {
    PoolConfig memory poolConfig = s_poolsBySourceToken[sourceToken];
    if (poolConfig.enabled) {
      return s_poolsBySourceToken[sourceToken].pool;
    }
    revert UnsupportedToken(sourceToken);
  }

  /// @inheritdoc BaseOnRampInterface
  function getExpectedNextSequenceNumber() external view returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc BaseOnRampInterface
  function setRouter(address router) public onlyOwner {
    s_router = router;
    emit RouterSet(router);
  }

  /// @inheritdoc BaseOnRampInterface
  function getRouter() external view returns (address router) {
    return s_router;
  }

  /// @inheritdoc BaseOnRampInterface
  function setConfig(OnRampConfig calldata config) external onlyOwner {
    s_config = config;
    emit OnRampConfigSet(config);
  }

  /// @inheritdoc BaseOnRampInterface
  function getConfig() external view returns (OnRampConfig memory config) {
    return s_config;
  }

  function addPool(IERC20 token, PoolInterface pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    if (s_poolsBySourceToken[token].enabled) revert PoolAlreadyAdded();

    s_poolsBySourceToken[token] = PoolConfig({pool: pool, enabled: true});
    s_sourceTokenList.push(token);

    emit PoolAdded(token, pool);
  }

  function removePool(IERC20 token, PoolInterface pool) public onlyOwner {
    PoolConfig memory oldConfig = s_poolsBySourceToken[token];
    // Check if the pool exists
    if (address(oldConfig.pool) == address(0)) revert PoolDoesNotExist(token);
    // Sanity check
    if (address(oldConfig.pool) != address(pool)) revert TokenPoolMismatch();

    s_poolsBySourceToken[token].enabled = false;

    emit PoolRemoved(token, pool);
  }

  /**
   * @notice Get all configured source tokens
   * @return Array of configured source tokens
   * @dev this is not very efficient but this method only exists for
   * offchain use so gas does not matter.
   */
  function getPoolTokens() external view returns (IERC20[] memory) {
    uint256 numberOfSupportedTokens = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[s_sourceTokenList[i]].enabled) {
        numberOfSupportedTokens++;
      }
    }

    IERC20[] memory sourceTokens = new IERC20[](numberOfSupportedTokens);
    uint256 j = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[s_sourceTokenList[i]].enabled) {
        sourceTokens[j++] = s_sourceTokenList[i];
      }
    }
    return sourceTokens;
  }

  /**
   * @notice Handles common checks and token locking for forwardFromRouter calls.
   * @dev this function is generic over message types, thereby reducing code duplication.
   * @param dataLength The length of the data field of the message
   * @param gasLimit The gasLimit set in message for destination execution
   * @param tokensAndAmounts The token payload to be sent. They will be locked into pools by this function.
   * @param originalSender The original sender of the message on the router.
   */
  function _handleForwardFromRouter(
    uint256 dataLength,
    uint256 gasLimit,
    Common.EVMTokenAndAmount[] memory tokensAndAmounts,
    address originalSender
  ) internal {
    if (s_router == address(0)) revert RouterNotSet();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_config.maxDataSize)) revert MessageTooLarge(uint256(s_config.maxDataSize), dataLength);
    if (gasLimit > uint256(s_config.maxGasLimit)) revert MessageGasLimitTooHigh();
    uint256 tokenLength = tokensAndAmounts.length;
    if (tokenLength > uint256(s_config.maxTokensLength)) revert UnsupportedNumberOfTokens();

    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    _removeTokens(tokensAndAmounts);

    // Lock all tokens in their corresponding pools
    for (uint256 i = 0; i < tokenLength; ++i) {
      Common.EVMTokenAndAmount memory ta = tokensAndAmounts[i];
      IERC20 token = IERC20(ta.token);
      PoolInterface pool = getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      pool.lockOrBurn(ta.amount);
    }
  }
}
