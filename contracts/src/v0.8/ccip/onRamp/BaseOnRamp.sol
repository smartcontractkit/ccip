// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPoolRegistry, PoolInterface} from "../pools/TokenPoolRegistry.sol";
import {HealthChecker, AFNInterface} from "../health/HealthChecker.sol";
import {IERC20} from "../pools/PoolCollector.sol";
import {AllowList} from "../access/AllowList.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {BaseOnRampInterface} from "../interfaces/onRamp/BaseOnRampInterface.sol";
import {CCIP} from "../models/Models.sol";

contract BaseOnRamp is BaseOnRampInterface, HealthChecker, TokenPoolRegistry, AllowList, AggregateRateLimiter {
  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable i_chainId;
  // Chain ID of the destination chain (where this contract sends messages)
  uint256 public immutable i_destinationChainId;

  // The last used sequence number. This is zero in the case where no
  // messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  uint64 internal s_sequenceNumber;

  // The current configuration of the onRamp.
  OnRampConfig internal s_config;
  // The router that is allowed to interact with this onRamp.
  address internal s_router;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    address[] memory allowlist,
    AFNInterface afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    address router
  )
    HealthChecker(afn)
    TokenPoolRegistry(tokens, pools)
    AllowList(allowlist)
    AggregateRateLimiter(rateLimiterConfig, tokenLimitsAdmin)
  {
    // TokenPoolRegistry does a check on tokensAndAmounts.length != pools.length
    i_chainId = chainId;
    i_destinationChainId = destinationChainId;
    s_config = config;
    s_router = router;
    s_sequenceNumber = 0;
  }

  /// @inheritdoc BaseOnRampInterface
  function getTokenPool(IERC20 token) external view override returns (PoolInterface) {
    return getPool(token);
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
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts,
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
      CCIP.EVMTokenAndAmount memory ta = tokensAndAmounts[i];
      IERC20 token = IERC20(ta.token);
      PoolInterface pool = getPool(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      pool.lockOrBurn(ta.amount);
    }
  }
}
