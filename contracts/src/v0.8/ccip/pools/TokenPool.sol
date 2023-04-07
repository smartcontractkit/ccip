// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IPool} from "../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../OwnerIsCreator.sol";
import {RateLimiter} from "../models/RateLimiter.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice Base abstract class with common functions for all token pools.
abstract contract TokenPool is IPool, OwnerIsCreator, Pausable {
  using EnumerableSet for EnumerableSet.AddressSet;
  using RateLimiter for RateLimiter.TokenBucket;

  // The immutable token that belongs to this pool.
  IERC20 internal immutable i_token;
  // A set of allowed onRamps.
  EnumerableSet.AddressSet private s_onRamps;
  // A set of allowed offRamps.
  EnumerableSet.AddressSet private s_offRamps;
  // The token bucket object that contains the bucket state.
  RateLimiter.TokenBucket private s_rateLimiter;

  constructor(IERC20 token, RateLimiter.Config memory rateLimiterConfig) {
    if (address(token) == address(0)) revert NullAddressNotAllowed();

    s_rateLimiter = RateLimiter.TokenBucket({
      rate: rateLimiterConfig.rate,
      capacity: rateLimiterConfig.capacity,
      tokens: rateLimiterConfig.capacity,
      lastUpdated: uint40(block.timestamp),
      isEnabled: rateLimiterConfig.isEnabled
    });

    i_token = token;
  }

  /// @inheritdoc IPool
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /// @notice Checks whether something is a permissioned onRamp on this contract.
  /// @return true if the given address is a permissioned onRamp.
  function isOnRamp(address onRamp) public view returns (bool) {
    return s_onRamps.contains(onRamp);
  }

  /// @notice Checks whether something is a permissioned offRamp on this contract.
  /// @return true is the given address is a permissioned offRamp.
  function isOffRamp(address offRamp) public view returns (bool) {
    return s_offRamps.contains(offRamp);
  }

  /// @notice Sets permissions for all on and offRamps.
  /// @dev Only callable by the owner
  /// @param onRamps A list of onRamps and their new permission status
  /// @param offRamps A list of offRamps and their new permission status
  function applyRampUpdates(RampUpdate[] memory onRamps, RampUpdate[] memory offRamps) public onlyOwner {
    for (uint256 i = 0; i < onRamps.length; ++i) {
      RampUpdate memory update = onRamps[i];
      if (update.allowed) {
        s_onRamps.add(update.ramp);
      } else {
        s_onRamps.remove(update.ramp);
      }

      emit OnRampAllowanceSet(onRamps[i].ramp, onRamps[i].allowed);
    }

    for (uint256 i = 0; i < offRamps.length; ++i) {
      RampUpdate memory update = offRamps[i];
      if (update.allowed) {
        s_offRamps.add(update.ramp);
      } else {
        s_offRamps.remove(update.ramp);
      }

      emit OffRampAllowanceSet(offRamps[i].ramp, offRamps[i].allowed);
    }
  }

  // ================================================================
  // |                        Rate limiting                         |
  // ================================================================

  /// @notice Consumes rate limiting capacity in this pool
  function _consumeRateLimit(uint256 amount) internal {
    s_rateLimiter._consume(amount);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentTokenBucketState() public view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  /// @notice Sets the rate limited config.
  /// @param config The new rate limiter config.
  /// @dev should only be callable by the owner or token limit admin.
  function setRateLimiterConfig(RateLimiter.Config memory config) public onlyOwner {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  // ================================================================
  // |                           Access                             |
  // ================================================================

  /// @notice Checks whether the msg.sender is a permissioned onRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOnRamp() {
    if (!isOnRamp(msg.sender)) revert PermissionsError();
    _;
  }

  /// @notice Checks whether the msg.sender is a permissioned offRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOffRamp() {
    if (!isOffRamp(msg.sender)) revert PermissionsError();
    _;
  }

  /// @inheritdoc IPool
  function pause() external override onlyOwner {
    _pause();
  }

  /// @inheritdoc IPool
  function unpause() external override onlyOwner {
    _unpause();
  }
}
