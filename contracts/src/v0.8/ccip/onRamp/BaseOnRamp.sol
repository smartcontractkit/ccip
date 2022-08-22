// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "../pools/TokenPoolRegistry.sol";
import "../health/HealthChecker.sol";
import "../priceFeedRegistry/PriceFeedRegistry.sol";
import "../pools/PoolCollector.sol";
import "../access/AllowList.sol";

contract BaseOnRamp is
  BaseOnRampInterface,
  HealthChecker,
  TokenPoolRegistry,
  PriceFeedRegistry,
  PoolCollector,
  AllowList
{
  using SafeERC20 for IERC20;

  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable i_chainId;
  // Chain ID of the destination chain (where this contract sends messages)
  uint256 public immutable i_destinationChainId;

  // The last used sequence number. This is zero in the case where no
  // messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  uint64 internal s_sequenceNumber;

  // The current configuration of the onRamp.
  OnRampConfig s_config;
  // The router that is allowed to interact with this onRamp.
  address internal s_router;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    OnRampConfig memory config,
    address router
  ) HealthChecker(afn) TokenPoolRegistry(tokens, pools) PriceFeedRegistry(tokens, feeds) AllowList(allowlist) {
    // TokenPoolRegistry does a check on tokens.length != pools.length
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
   * @param tokens The tokens to be sent. They will be locked into pools by this function.
   * @param amounts The amounts corresponding to the tokens.
   * @param originalSender The original sender of the message on the router.
   */
  function handleForwardFromRouter(
    uint256 dataLength,
    IERC20[] memory tokens,
    uint256[] memory amounts,
    address originalSender
  ) internal {
    if (s_router == address(0)) revert RouterNotSet();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_config.maxDataSize)) revert MessageTooLarge(uint256(s_config.maxDataSize), dataLength);
    uint256 tokenLength = tokens.length;
    if (tokenLength > uint256(s_config.maxTokensLength) || tokenLength != amounts.length)
      revert UnsupportedNumberOfTokens();

    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    // Lock all tokens in their corresponding pools
    for (uint256 i = 0; i < tokenLength; ++i) {
      IERC20 token = tokens[i];
      PoolInterface pool = getPool(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      pool.lockOrBurn(amounts[i]);
    }
  }
}
