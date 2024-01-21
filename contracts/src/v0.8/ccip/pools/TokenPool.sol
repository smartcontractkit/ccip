// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IPool} from "../interfaces/pools/IPool.sol";
import {IARM} from "../interfaces/IARM.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {IERC165} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";
import {Router} from "../Router.sol";

/// @notice Base abstract class with common functions for all token pools.
/// A token pool serves as isolated place for holding tokens and token specific logic
/// that may execute as tokens move across the bridge.
abstract contract TokenPool is IPool, OwnerIsCreator, IERC165 {
  using EnumerableSet for EnumerableSet.AddressSet;
  using EnumerableSet for EnumerableSet.UintSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error CallerIsNotARampOnRouter(address caller);
  error ZeroAddressNotAllowed();
  error SenderNotAllowed(address sender);
  error AllowListNotEnabled();
  error NonExistentChain(uint64 chainSelector);
  error ChainNotAllowed(uint64 chainSelector);
  error BadARMSignal();
  error ChainAlreadyExists(uint64 chainSelector);

  event Locked(address indexed sender, uint256 amount);
  event Burned(address indexed sender, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);
  event ChainAdded(uint64 chainSelector, RateLimiter.Config rateLimiterConfig);
  event ChainConfigured(uint64 chainSelector, RateLimiter.Config rateLimiterConfig);
  event ChainRemoved(uint64 chainSelector);
  event OffRampConfigured(address offRamp, RateLimiter.Config rateLimiterConfig);
  event AllowListAdd(address sender);
  event AllowListRemove(address sender);

  struct ChainUpdate {
    uint64 chainSelector;
    bool allowed;
    RateLimiter.Config rateLimiterConfig;
  }

  /// @dev The bridgeable token that is managed by this pool.
  IERC20 internal immutable i_token;
  /// @dev The address of the arm proxy
  address internal immutable i_armProxy;

  // TODO: do we want this immutable?
  Router internal immutable i_router;
  /// @dev The immutable flag that indicates if the pool is access-controlled.
  bool internal immutable i_allowlistEnabled;
  /// @dev A set of addresses allowed to trigger lockOrBurn as original senders.
  /// Only takes effect if i_allowlistEnabled is true.
  /// This can be used to ensure only token-issuer specified addresses can
  /// move tokens.
  EnumerableSet.AddressSet internal s_allowList;

  /// @dev A set of allowed onRamps. We want the whitelist to be enumerable to
  /// be able to quickly determine (without parsing logs) who can access the pool.
  EnumerableSet.UintSet internal s_remoteChains;
  /// @dev Inbound rate limits. This allows per destination chain
  /// token issuer specified rate limiting (e.g. issuers may trust chains to varying
  /// degrees and prefer different limits)
  mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket) internal s_inboundRateLimits;
  /// @dev Outbound rate limits. Corresponds to the inbound rate limit for the pool
  /// on the remote chain.
  mapping(uint64 remoteChainSelector => RateLimiter.TokenBucket) internal s_outboundRateLimits;

  constructor(IERC20 token, address[] memory allowlist, address armProxy, address router) {
    if (address(token) == address(0) || router == address(0)) revert ZeroAddressNotAllowed();
    i_token = token;
    i_armProxy = armProxy;
    i_router = Router(router);

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

  /// @inheritdoc IPool
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == type(IPool).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  // ================================================================
  // │                     Chain permissions                        │
  // ================================================================

  /// @notice Checks whether a chain selector is permissioned on this contract.
  /// @return true if the given chain selector is a permissioned remote chain.
  function isSupportedChain(uint64 remoteChainSelector) public view returns (bool) {
    return s_remoteChains.contains(remoteChainSelector);
  }

  /// @notice Get chain whitelist
  /// @return list of chains.
  function getSupportedChains() public view returns (uint64[] memory) {
    uint256[] memory uint256ChainSelectors = s_remoteChains.values();
    uint64[] memory chainSelectors = new uint64[](uint256ChainSelectors.length);
    for (uint256 i = 0; i < uint256ChainSelectors.length; ++i) {
      chainSelectors[i] = uint64(uint256ChainSelectors[i]);
    }

    return chainSelectors;
  }

  /// @notice Sets permissions for all on and chains.
  /// @dev Only callable by the owner
  /// @param chains A list of chains and their new permission status/rate limits
  function applyChainUpdates(ChainUpdate[] calldata chains) external virtual onlyOwner {
    _applyChainUpdates(chains);
  }

  function _applyChainUpdates(ChainUpdate[] calldata chains) internal onlyOwner {
    for (uint256 i = 0; i < chains.length; ++i) {
      ChainUpdate memory update = chains[i];
      if (update.allowed) {
        if (s_remoteChains.add(update.chainSelector)) {
          s_inboundRateLimits[update.chainSelector] = RateLimiter.TokenBucket({
            rate: update.rateLimiterConfig.rate,
            capacity: update.rateLimiterConfig.capacity,
            tokens: update.rateLimiterConfig.capacity,
            lastUpdated: uint32(block.timestamp),
            isEnabled: update.rateLimiterConfig.isEnabled
          });
          s_outboundRateLimits[update.chainSelector] = RateLimiter.TokenBucket({
            rate: update.rateLimiterConfig.rate,
            capacity: update.rateLimiterConfig.capacity,
            tokens: update.rateLimiterConfig.capacity,
            lastUpdated: uint32(block.timestamp),
            isEnabled: update.rateLimiterConfig.isEnabled
          });
          emit ChainAdded(update.chainSelector, update.rateLimiterConfig);
        } else {
          revert ChainAlreadyExists(update.chainSelector);
        }
      } else {
        if (s_remoteChains.remove(update.chainSelector)) {
          delete s_inboundRateLimits[update.chainSelector];
          delete s_outboundRateLimits[update.chainSelector];
          emit ChainRemoved(update.chainSelector);
        } else {
          // Cannot remove a non-existent onRamp.
          revert NonExistentChain(update.chainSelector);
        }
      }
    }
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Consumes outbound rate limiting capacity in this pool
  function _consumeOnRampRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_outboundRateLimits[remoteChainSelector]._consume(amount, address(i_token));
  }

  /// @notice Consumes inbound rate limiting capacity in this pool
  function _consumeOffRampRateLimit(uint64 remoteChainSelector, uint256 amount) internal {
    s_inboundRateLimits[remoteChainSelector]._consume(amount, address(i_token));
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentOnRampRateLimiterState(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_outboundRateLimits[remoteChainSelector]._currentTokenBucketState();
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentOffRampRateLimiterState(
    uint64 remoteChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_inboundRateLimits[remoteChainSelector]._currentTokenBucketState();
  }

  /// @notice Sets the onramp rate limited config.
  /// @param config The new rate limiter config.
  function setChainRateLimiterConfig(uint64 remoteChainSelector, RateLimiter.Config memory config) external onlyOwner {
    if (!isSupportedChain(remoteChainSelector)) revert NonExistentChain(remoteChainSelector);
    s_inboundRateLimits[remoteChainSelector]._setTokenBucketConfig(config);
    s_outboundRateLimits[remoteChainSelector]._setTokenBucketConfig(config);
    emit ChainConfigured(remoteChainSelector, config);
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Checks whether the msg.sender is a permissioned onRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOnRamp(uint64 remoteChainSelector) {
    if (!s_remoteChains.contains(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!(msg.sender == i_router.getOnRamp(remoteChainSelector))) revert CallerIsNotARampOnRouter(msg.sender);
    _;
  }

  /// @notice Checks whether the msg.sender is a permissioned offRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOffRamp(uint64 remoteChainSelector) {
    if (!s_remoteChains.contains(remoteChainSelector)) revert ChainNotAllowed(remoteChainSelector);
    if (!i_router.isOffRamp(remoteChainSelector, msg.sender)) revert CallerIsNotARampOnRouter(msg.sender);
    _;
  }

  // ================================================================
  // │                          Allowlist                           │
  // ================================================================

  modifier checkAllowList(address sender) {
    if (i_allowlistEnabled && !s_allowList.contains(sender)) revert SenderNotAllowed(sender);
    _;
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
