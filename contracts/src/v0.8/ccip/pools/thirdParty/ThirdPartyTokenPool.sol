// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IPool} from "../../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {ConfirmedBridgeRoleWithProposal} from "./ConfirmedBridgeRoleWithProposal.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";
import {IERC165} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/introspection/IERC165.sol";
import {EnumerableSet} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/structs/EnumerableSet.sol";

/// @notice Allows owner to grant a bridge permission to specify multiple burn mint addresses.
/// @dev Rate limits are chainSelector -> rateLimits.
abstract contract ThirdPartyTokenPool is IPool, OwnerIsCreator, ConfirmedBridgeRoleWithProposal, IERC165 {
  using EnumerableSet for EnumerableSet.AddressSet;
  using RateLimiter for RateLimiter.TokenBucket;

  error ZeroAddressNotAllowed();
  error NonExistentCaller(address caller);
  error CallerAlreadyExists(address caller);

  event LockOrBurnCallerAdded(address caller, uint64 destChainSelector);
  event LockOrBurnCallerRemoved(address caller);
  event ReleaseOrMintCallerAdded(address caller, uint64 sourceChainSelector);
  event ReleaseOrMintCallerRemoved(address caller);
  event LockOrBurnRateLimitConfigured(uint64 destChainSelector, RateLimiter.Config rateLimiterConfig);
  event ReleaseOrMintRateLimitConfigured(uint64 sourceChainSelector, RateLimiter.Config rateLimiterConfig);

  struct CallerUpdate {
    address caller;
    uint64 chainSelector;
    bool allowed;
  }

  /// @dev The bridgeable token that is managed by this pool.
  IERC20 internal immutable i_token;

  /// @dev A set of allowed lockOrBurn callers. We want the whitelist to be enumerable to
  /// be able to quickly determine (without parsing logs) who can lockOrBurn.
  EnumerableSet.AddressSet internal s_lockOrBurnCallers;
  /// @dev Inbound rate limits. This allows per destination chain rate limiting
  /// (e.g. issuers may trust chains to varying degrees and prefer different limits)
  mapping(uint64 => RateLimiter.TokenBucket) internal s_lockOrBurnLimits;

  /// @dev A set of allowed releaseOrMint callers.
  EnumerableSet.AddressSet internal s_releaseOrMintCallers;
  /// @dev Outbound rate limits. Corresponds to the inbound rate limit for the pool
  /// on the remote chain.
  mapping(uint64 => RateLimiter.TokenBucket) internal s_releaseOrMintLimits;

  constructor(IERC20 token) {
    if (address(token) == address(0)) revert ZeroAddressNotAllowed();
    i_token = token;
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
  // │                     Caller permissions                       │
  // ================================================================

  /// @notice Checks whether an address can call lockOrBurn.
  /// @return true if the given address is a permissioned lockOrBurn caller.
  function isLockOrBurnCaller(address caller) public view returns (bool) {
    return s_lockOrBurnCallers.contains(caller);
  }

  /// @notice Checks whether an address can call releaseOrMint.
  /// @return true if the given address is a permissioned releaseOrMint caller.
  function isReleaseOrMintCaller(address caller) public view returns (bool) {
    return s_releaseOrMintCallers.contains(caller);
  }

  /// @notice Get list of permissioned lockOrBurn callers.
  /// @return list of lockOrBurn callers.
  function getLockOrBurnCallers() public view returns (address[] memory) {
    return s_lockOrBurnCallers.values();
  }

  /// @notice Get list of permissioned releaseOrMint callers.
  /// @return list of releaseOrMint callers.
  function getReleaseOrMintCallers() public view returns (address[] memory) {
    return s_releaseOrMintCallers.values();
  }

  /// @notice Sets permissions for all lockOrBurn and releaseOrMint callers.
  /// @dev Only callable by the owner
  /// @param lockOrBurnCallers A list of permissioned lockOrBurn callers.
  /// @param releaseOrMintCallers A list of permissioned releaseOrMint callers.
  function applyCallerUpdates(
    CallerUpdate[] calldata lockOrBurnCallers,
    CallerUpdate[] calldata releaseOrMintCallers
  ) external virtual onlyOwnerOrBridge {
    _applyCallerUpdates(lockOrBurnCallers, releaseOrMintCallers);
  }

  function _applyCallerUpdates(
    CallerUpdate[] calldata lockOrBurnCallers,
    CallerUpdate[] calldata releaseOrMintCallers
  ) internal onlyOwnerOrBridge {
    for (uint256 i = 0; i < lockOrBurnCallers.length; ++i) {
      CallerUpdate memory update = lockOrBurnCallers[i];

      if (update.allowed) {
        if (s_lockOrBurnCallers.add(update.caller)) {
          emit LockOrBurnCallerAdded(update.caller, update.chainSelector);
        } else {
          revert CallerAlreadyExists(update.caller);
        }
      } else {
        if (s_lockOrBurnCallers.remove(update.caller)) {
          emit LockOrBurnCallerRemoved(update.caller);
        } else {
          // Cannot remove a non-existent releaseOrMint caller.
          revert NonExistentCaller(update.caller);
        }
      }
    }

    for (uint256 i = 0; i < releaseOrMintCallers.length; ++i) {
      CallerUpdate memory update = releaseOrMintCallers[i];

      if (update.allowed) {
        if (s_releaseOrMintCallers.add(update.caller)) {
          emit ReleaseOrMintCallerAdded(update.caller, update.chainSelector);
        } else {
          revert CallerAlreadyExists(update.caller);
        }
      } else {
        if (s_releaseOrMintCallers.remove(update.caller)) {
          emit ReleaseOrMintCallerRemoved(update.caller);
        } else {
          // Cannot remove a non-existent releaseOrMint caller.
          revert NonExistentCaller(update.caller);
        }
      }
    }
  }

  // ================================================================
  // │                        Rate limiting                         │
  // ================================================================

  /// @notice Consumes lockOrBurn rate limiting capacity in this pool
  function _consumeLockOrBurnRateLimit(uint64 destChainSelector, uint256 amount) internal {
    s_lockOrBurnLimits[destChainSelector]._consume(amount, address(i_token));
  }

  /// @notice Consumes releaseOrMint rate limiting capacity in this pool
  function _consumeReleaseOrMintRateLimit(uint64 sourceChainSelector, uint256 amount) internal {
    s_releaseOrMintLimits[sourceChainSelector]._consume(amount, address(i_token));
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentLockOrBurnRateLimiterState(
    uint64 destChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_lockOrBurnLimits[destChainSelector]._currentTokenBucketState();
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentReleaseOrMintRateLimiterState(
    uint64 sourceChainSelector
  ) external view returns (RateLimiter.TokenBucket memory) {
    return s_releaseOrMintLimits[sourceChainSelector]._currentTokenBucketState();
  }

  /// @notice Sets the lockOrBurn rate limited config for a lane.
  /// @param config The new rate limiter config.
  function setLockOrBurnRateLimiterConfig(
    uint64 destChainSelector,
    RateLimiter.Config memory config
  ) external onlyOwner {
    s_lockOrBurnLimits[destChainSelector]._setTokenBucketConfig(config);
    emit LockOrBurnRateLimitConfigured(destChainSelector, config);
  }

  /// @notice Sets the releaseOrMint rate limited config.
  /// @param config The new rate limiter config.
  function setReleaseOrMintRateLimiterConfig(
    uint64 sourceChainSelector,
    RateLimiter.Config memory config
  ) external onlyOwner {
    s_releaseOrMintLimits[sourceChainSelector]._setTokenBucketConfig(config);
    emit ReleaseOrMintRateLimitConfigured(sourceChainSelector, config);
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Checks whether the msg.sender is a permissioned lockOrBurn caller on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyLockOrBurnCaller() {
    if (!isLockOrBurnCaller(msg.sender)) revert PermissionsError();
    _;
  }

  /// @notice Checks whether the msg.sender is a permissioned releaseOrMint caller on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyReleaseOrMintCaller() {
    if (!isReleaseOrMintCaller(msg.sender)) revert PermissionsError();
    _;
  }
}
