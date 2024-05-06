// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IARM} from "../interfaces/IARM.sol";
import {IPool} from "../interfaces/IPool.sol";
import {IRouter} from "../interfaces/IRouter.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {IERC165} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice Base abstract class with common functions for all token pools.
/// A token pool serves as isolated place for holding tokens and token specific logic
/// that may execute as tokens move across the bridge.
abstract contract TokenPool is IPool, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error CallerIsNotARampOnRouter(address caller);
  error ZeroAddressNotAllowed();
  error SenderNotAllowed(address sender);
  error AllowListNotEnabled();
  error NonExistentChain(uint64 remoteChainSelector);
  error ChainNotAllowed(uint64 remoteChainSelector);
  error BadARMSignal();
  error ChainAlreadyExists(uint64 chainSelector);
  error InvalidSourcePoolAddress(bytes sourcePoolAddress);

  event Locked(address indexed sender, uint256 amount);
  event Burned(address indexed sender, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);
  event ChainAdded(
    uint64 remoteChainSelector,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig
  );
  event ChainConfigured(
    uint64 remoteChainSelector,
    RateLimiter.Config outboundRateLimiterConfig,
    RateLimiter.Config inboundRateLimiterConfig
  );
  event ChainRemoved(uint64 remoteChainSelector);
  event RemotePoolSet(uint64 indexed remoteChainSelector, bytes previousPoolAddress, bytes remotePoolAddress);
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event RouterUpdated(address oldRouter, address newRouter);

  struct ChainUpdate {
    uint64 remoteChainSelector; // ──╮ Remote chain selector
    bool allowed; // ────────────────╯ Whether the chain is allowed
    bytes remotePoolAddress; //        Address of the remote pool
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain
  }

  struct RemoteChainConfig {
    RateLimiter.TokenBucket outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain
    RateLimiter.TokenBucket inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain
    bytes remotePoolAddress; // ────╮ Address of the remote pool
    bool allowed; // ───────────────╯ Whether the chain is allowed
  }

  /// @dev The bridgeable token that is managed by this pool.
  IERC20 internal immutable i_token;
  /// @dev The address of the arm proxy
  address internal immutable i_armProxy;
  /// @dev The immutable flag that indicates if the pool is access-controlled.
  bool internal immutable i_allowlistEnabled;
  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can
  /// move tokens.
  EnumerableSet.AddressSet internal s_allowList;
  /// @dev The address of the router
  IRouter internal s_router;
  /// @dev A set of allowed chain selectors. We want the allowlist to be enumerable to
  /// be able to quickly determine (without parsing logs) who can access the pool.
  /// @dev The chain selectors are in uin256 format because of the EnumerableSet implementation.
  EnumerableSet.UintSet internal s_remoteChainSelectors;
  mapping(uint64 remoteChainSelector => RemoteChainConfig) internal s_remoteChainConfigs;

  constructor(IERC20 token, address[] memory allowlist, address armProxy, address router) {
    if (address(token) == address(0) || router == address(0)) revert ZeroAddressNotAllowed();
    i_token = token;
    i_armProxy = armProxy;
    s_router = IRouter(router);

    // Pool can be set as permissioned or permissionless at deployment time only to save hot-path gas.
    i_allowlistEnabled = allowlist.length > 0;
    if (i_allowlistEnabled) {
      _applyAllowListUpdates(new address[](0), allowlist);
    }
  }

  /// @notice Get ARM proxy address
  /// @return armProxy Address of arm proxy
  function getArmProxy() public view returns (address armProxy) {
    return i_armProxy;
  }

  /// @notice Gets the IERC20 token that this pool can lock or burn.
  /// @return token The IERC20 token representation.
  function getToken() public view returns (IERC20 token) {
    return i_token;
  }

  /// @notice Gets the pool's Router
  /// @return router The pool's Router
  function getRouter() public view returns (address router) {
    return address(s_router);
  }

  /// @notice Sets the pool's Router
  /// @param newRouter The new Router
  function setRouter(address newRouter) public onlyOwner {
    if (newRouter == address(0)) revert ZeroAddressNotAllowed();
    address oldRouter = address(s_router);
    s_router = IRouter(newRouter);

    emit RouterUpdated(oldRouter, newRouter);
  }

  /// @notice Signals which version of the pool interface is supported
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == Pool.CCIP_POOL_V1 || interfaceId == type(IPool).interfaceId
      || interfaceId == type(IERC165).interfaceId;
  }

  // Validate if the sending pool is allowed
  function _validateSourceCaller(uint64 remoteChainSelector, bytes memory sourcePoolAddress) internal view {
    if (keccak256(sourcePoolAddress) != keccak256(getRemotePool(remoteChainSelector))) {
      revert InvalidSourcePoolAddress(sourcePoolAddress);
    }
  }

  // ================================================================
  // │                     Chain permissions                        │
  // ================================================================

  /// @notice Gets the pool address on the remote chain.
  /// @param remoteChainSelector Destination chain selector.
  function getRemotePool(uint64 remoteChainSelector) public view returns (bytes memory) {
    return s_remoteChainConfigs[remoteChainSelector].remotePoolAddress;
  }

  /// @notice Sets the remote pool address for a given chain selector.
  /// @param remoteChainSelector The remote chain selector for which the remote pool address is being set.
  /// @param remotePoolAddress The address of the remote pool.
  function setRemotePool(uint64 remoteChainSelector, bytes calldata remotePoolAddress) external onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);

    bytes memory prevAddress = s_remoteChainConfigs[remoteChainSelector].remotePoolAddress;
    s_remoteChainConfigs[remoteChainSelector].remotePoolAddress = remotePoolAddress;

    emit RemotePoolSet(remoteChainSelector, prevAddress, remotePoolAddress);
  }

  /// @inheritdoc IPool
  function isSupportedChain(uint64 remoteChainSelector) public view returns (bool) {
    return s_remoteChainSelectors.contains(remoteChainSelector);
  }

  /// @notice Get list of allowed chains
  /// @return list of chains.
  function getSupportedChains() public view returns (uint64[] memory) {
    uint256[] memory uint256ChainSelectors = s_remoteChainSelectors.values();
    uint64[] memory chainSelectors = new uint64[](uint256ChainSelectors.length);
    for (uint256 i = 0; i < uint256ChainSelectors.length; ++i) {
      chainSelectors[i] = uint64(uint256ChainSelectors[i]);
    }

    return chainSelectors;
  }

  /// @notice Sets the permissions for a list of chains selectors. Actual senders for these chains
  /// need to be allowed on the Router to interact with this pool.
  /// @dev Only callable by the owner
  /// @param chains A list of chains and their new permission status & rate limits. Rate limits
  /// are only used when the chain is being added through `allowed` being true.
  function applyChainUpdates(ChainUpdate[] calldata chains) external virtual onlyOwner {
    for (uint256 i = 0; i < chains.length; ++i) {
      ChainUpdate memory update = chains[i];
      RateLimiter._validateTokenBucketConfig(update.outboundRateLimiterConfig, !update.allowed);
      RateLimiter._validateTokenBucketConfig(update.inboundRateLimiterConfig, !update.allowed);

      if (update.allowed) {
        // If the chain already exists, revert
        if (!s_remoteChainSelectors.add(update.remoteChainSelector)) {
          revert ChainAlreadyExists(update.remoteChainSelector);
        }

        s_remoteChainConfigs[update.remoteChainSelector] = RemoteChainConfig({
          outboundRateLimiterConfig: RateLimiter.TokenBucket({
            rate: update.outboundRateLimiterConfig.rate,
            capacity: update.outboundRateLimiterConfig.capacity,
            tokens: update.outboundRateLimiterConfig.capacity,
            lastUpdated: uint32(block.timestamp),
            isEnabled: update.outboundRateLimiterConfig.isEnabled
          }),
          inboundRateLimiterConfig: RateLimiter.TokenBucket({
            rate: update.inboundRateLimiterConfig.rate,
            capacity: update.inboundRateLimiterConfig.capacity,
            tokens: update.inboundRateLimiterConfig.capacity,
            lastUpdated: uint32(block.timestamp),
            isEnabled: update.inboundRateLimiterConfig.isEnabled
          }),
          remotePoolAddress: update.remotePoolAddress,
          allowed: update.allowed
        });

        emit ChainAdded(update.remoteChainSelector, update.outboundRateLimiterConfig, update.inboundRateLimiterConfig);
      } else {
        // If the chain doesn't exist, revert
        if (!s_remoteChainSelectors.remove(update.remoteChainSelector)) {
          revert NonExistentChain(update.remoteChainSelector);
        }

        delete s_remoteChainConfigs[update.remoteChainSelector];

        emit ChainRemoved(update.remoteChainSelector);
      }
    }
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Consumes outbound rate limiting capacity in this pool
  function _consumeOutboundRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._consume(amount, address(i_token));
  }

  /// @notice Consumes inbound rate limiting capacity in this pool
  function _consumeInboundRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._consume(amount, address(i_token));
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function getCurrentOutboundRateLimiterState(uint64 remoteChainSelector)
    external
    view
    returns (RateLimiter.TokenBucket memory)
  {
    return s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._currentTokenBucketState();
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function getCurrentInboundRateLimiterState(uint64 remoteChainSelector)
    external
    view
    returns (RateLimiter.TokenBucket memory)
  {
    return s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._currentTokenBucketState();
  }

  /// @notice Sets the chain rate limiter config.
  /// @param remoteChainSelector The remote chain selector for which the rate limits apply.
  /// @param outboundConfig The new outbound rate limiter config, meaning the onRamp rate limits for the given chain.
  /// @param inboundConfig The new inbound rate limiter config, meaning the offRamp rate limits for the given chain.
  function setChainRateLimiterConfig(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) external virtual onlyOwner {
    _setRateLimitConfig(remoteChainSelector, outboundConfig, inboundConfig);
  }

  function _setRateLimitConfig(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) internal {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);
    RateLimiter._validateTokenBucketConfig(outboundConfig, false);
    s_remoteChainConfigs[remoteChainSelector].outboundRateLimiterConfig._setTokenBucketConfig(outboundConfig);
    RateLimiter._validateTokenBucketConfig(inboundConfig, false);
    s_remoteChainConfigs[remoteChainSelector].inboundRateLimiterConfig._setTokenBucketConfig(inboundConfig);
    emit ChainConfigured(remoteChainSelector, outboundConfig, inboundConfig);
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned onRamp for the given chain on the Router.
  function _onlyOnRamp(uint64 remoteChainSelector) internal view {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!(msg.sender == s_router.getOnRamp(remoteChainSelector))) revert CallerIsNotARampOnRouter(msg.sender);
  }

  /// @notice Checks whether remote chain selector is configured on this contract, and if the msg.sender
  /// is a permissioned offRamp for the given chain on the Router.
  function _onlyOffRamp(uint64 remoteChainSelector) internal view {
    if (!isSupportedChain(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!s_router.isOffRamp(remoteChainSelector, msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
  }

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  function _checkAllowList(address sender) internal view {
    if (i_allowlistEnabled && !s_allowList.contains(sender)) revert SenderNotAllowed(sender);
  }

  /// @notice Gets whether the allowList functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return i_allowlistEnabled;
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  function getAllowList() external view returns (address[] memory) {
    return s_allowList.values();
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  /// @dev allowListing will be removed before public launch
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    if (!i_allowlistEnabled) revert AllowListNotEnabled();

    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowList.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowList.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }

  /// @notice Ensure that there is no active curse.
  modifier whenHealthy() {
    if (IARM(i_armProxy).isCursed()) revert BadARMSignal();
    _;
  }
}
