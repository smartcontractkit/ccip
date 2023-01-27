// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAFN} from "../interfaces/health/IAFN.sol";
import {IBaseOnRamp} from "../interfaces/onRamp/IBaseOnRamp.sol";
import {IPool} from "../interfaces/pools/IPool.sol";

import {HealthChecker} from "../health/HealthChecker.sol";
import {AllowList} from "../access/AllowList.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {Common} from "../models/Common.sol";

import {IERC20} from "../../vendor/IERC20.sol";

contract BaseOnRamp is IBaseOnRamp, HealthChecker, AllowList, AggregateRateLimiter {
  // Chain ID of the source chain (where this contract is deployed)
  uint64 internal immutable i_chainId;
  // Chain ID of the destination chain (where this contract sends messages)
  uint64 internal immutable i_destinationChainId;

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
  address[] private s_sourceTokenList;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    address[] memory tokens,
    IPool[] memory pools,
    address[] memory allowlist,
    IAFN afn,
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
      s_poolsBySourceToken[IERC20(tokens[i])] = PoolConfig({pool: pools[i], enabled: true});
    }
  }

  /// @inheritdoc IBaseOnRamp
  function getPoolBySourceToken(IERC20 sourceToken) public view virtual override returns (IPool) {
    PoolConfig memory poolConfig = s_poolsBySourceToken[sourceToken];
    if (poolConfig.enabled) {
      return s_poolsBySourceToken[sourceToken].pool;
    }
    revert UnsupportedToken(sourceToken);
  }

  /// @inheritdoc IBaseOnRamp
  function getExpectedNextSequenceNumber() external view override returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc IBaseOnRamp
  function setRouter(address router) public override onlyOwner {
    s_router = router;
    emit RouterSet(router);
  }

  /// @inheritdoc IBaseOnRamp
  function getRouter() external view override returns (address router) {
    return s_router;
  }

  /// @inheritdoc IBaseOnRamp
  function setOnRampConfig(OnRampConfig calldata config) external override onlyOwner {
    s_config = config;
    emit OnRampConfigSet(config);
  }

  /// @inheritdoc IBaseOnRamp
  function getOnRampConfig() external view override returns (OnRampConfig memory config) {
    return s_config;
  }

  /// @inheritdoc IBaseOnRamp
  function getChainId() external view override returns (uint64) {
    return i_chainId;
  }

  /// @inheritdoc IBaseOnRamp
  function getDestinationChainId() external view override returns (uint64) {
    return i_destinationChainId;
  }

  /**
   * @notice Add a new token pool
   * @param token The source token
   * @param pool The pool that will be used
   * @dev This method can only be called by the owner of the contract.
   */
  function addPool(IERC20 token, IPool pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    if (s_poolsBySourceToken[token].enabled) revert PoolAlreadyAdded();

    s_poolsBySourceToken[token] = PoolConfig({pool: pool, enabled: true});
    s_sourceTokenList.push(address(token));

    emit PoolAdded(token, pool);
  }

  /**
   * @notice Remove a token pool
   * @param token The source token
   * @param pool The pool that will be removed
   * @dev This method can only be called by the owner of the contract.
   */
  function removePool(IERC20 token, IPool pool) public onlyOwner {
    PoolConfig memory oldConfig = s_poolsBySourceToken[token];
    // Check if the pool exists
    if (address(oldConfig.pool) == address(0)) revert PoolDoesNotExist(token);
    // Sanity check
    if (address(oldConfig.pool) != address(pool)) revert TokenPoolMismatch();

    s_poolsBySourceToken[token].enabled = false;

    emit PoolRemoved(token, pool);
  }

  /// @inheritdoc IBaseOnRamp
  function getSupportedTokens() public view virtual returns (address[] memory) {
    uint256 numberOfSupportedTokens = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[IERC20(s_sourceTokenList[i])].enabled) {
        numberOfSupportedTokens++;
      }
    }

    address[] memory sourceTokens = new address[](numberOfSupportedTokens);
    uint256 j = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[IERC20(s_sourceTokenList[i])].enabled) {
        sourceTokens[j++] = s_sourceTokenList[i];
      }
    }
    return sourceTokens;
  }

  /// @notice Validate the forwarded message with various checks.
  /// @param dataLength The length of the data field of the message
  /// @param gasLimit The gasLimit set in message for destination execution
  /// @param tokensAndAmounts The token payload to be sent. They will be locked into pools by this function.
  /// @param originalSender The original sender of the message on the router.
  function _validateMessage(
    uint256 dataLength,
    uint256 gasLimit,
    Common.EVMTokenAndAmount[] memory tokensAndAmounts,
    address originalSender
  ) internal {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_config.maxDataSize)) revert MessageTooLarge(uint256(s_config.maxDataSize), dataLength);
    if (gasLimit > uint256(s_config.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (tokensAndAmounts.length > uint256(s_config.maxTokensLength)) revert UnsupportedNumberOfTokens();
    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    _removeTokens(tokensAndAmounts);
  }
}
