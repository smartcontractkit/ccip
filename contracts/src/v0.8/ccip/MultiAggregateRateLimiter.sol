// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {OwnerIsCreator} from "./../shared/access/OwnerIsCreator.sol";
import {Client} from "./libraries/Client.sol";
import {RateLimiterNoEvents} from "./libraries/RateLimiterNoEvents.sol";
import {USDPriceWith18Decimals} from "./libraries/USDPriceWith18Decimals.sol";

/// @notice The aggregate rate limiter is a wrapper of the token bucket rate limiter
/// which permits rate limiting based on the aggregate value of a group of
/// token transfers, using a price registry to convert to a numeraire asset (e.g. USD).
contract MultiAggregateRateLimiter is OwnerIsCreator {
  using RateLimiterNoEvents for RateLimiterNoEvents.TokenBucket;
  using USDPriceWith18Decimals for uint224;

  error PriceNotFoundForToken(address token);
  error UpdateLengthMismatch();

  event RateLimiterConfigUpdated(uint64 indexed chainSelector, RateLimiterNoEvents.Config config);
  event RateLimiterTokensConsumed(uint64 indexed chainSelector, uint256 tokens);
  event AdminSet(address newAdmin);

  // The address of the token limit admin that has the same permissions as the owner.
  address internal s_admin;

  // Rate limiter token bucket states per chain
  mapping(uint64 chainSelector => RateLimiterNoEvents.TokenBucket rateLimiter) s_rateLimitersByChainSelector;

  /// @notice A collection of rate limiter configuration updates
  struct RateLimiterConfigUpdates {
    uint64[] chainSelectors;
    RateLimiterNoEvents.Config[] rateLimiterConfigs;
  }

  /// @param rateLimiterConfigs The RateLimiterNoEvents.Configs per chain containing the capacity and refill rate
  /// of the bucket
  /// @param admin the admin address to set
  constructor(RateLimiterConfigUpdates memory rateLimiterConfigs, address admin) {
    _applyRateLimiterConfigUpdates(rateLimiterConfigs);
    _setAdmin(admin);
  }

  /// @notice Consumes value from the rate limiter bucket based on the token value given.
  /// @param chainSelector chain selector to apply rate limit to
  /// @param value consumed value
  function _rateLimitValue(uint64 chainSelector, uint256 value) internal {
    s_rateLimitersByChainSelector[chainSelector]._consume(value, address(0));
    emit RateLimiterTokensConsumed(chainSelector, value);
  }

  function _getTokenValue(
    Client.EVMTokenAmount memory tokenAmount,
    IPriceRegistry priceRegistry
  ) internal view returns (uint256) {
    // not fetching validated price, as price staleness is not important for value-based rate limiting
    // we only need to verify the price is not 0
    uint224 pricePerToken = priceRegistry.getTokenPrice(tokenAmount.token).value;
    if (pricePerToken == 0) revert PriceNotFoundForToken(tokenAmount.token);
    return pricePerToken._calcUSDValueFromTokenAmount(tokenAmount.amount);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @param chainSelector chain selector to retrieve state for
  /// @return The token bucket.
  function currentRateLimiterState(uint64 chainSelector) external view returns (RateLimiterNoEvents.TokenBucket memory) {
    return s_rateLimitersByChainSelector[chainSelector]._currentTokenBucketState();
  }

  /// @notice Applies the provided rate limiter config updates.
  /// @param rateLimiterUpdates Rate limiter updates
  /// @dev should only be callable by the owner or token limit admin
  function applyRateLimiterConfigUpdates(RateLimiterConfigUpdates memory rateLimiterUpdates) external onlyAdminOrOwner {
    _applyRateLimiterConfigUpdates(rateLimiterUpdates);
  }

  /// @notice Applies the provided rate limiter config updates.
  /// @param rateLimiterUpdates Rate limiter updates
  function _applyRateLimiterConfigUpdates(RateLimiterConfigUpdates memory rateLimiterUpdates) internal {
    uint256 updateLength = rateLimiterUpdates.chainSelectors.length;
    if (updateLength != rateLimiterUpdates.rateLimiterConfigs.length) {
      revert UpdateLengthMismatch();
    }

    for (uint256 i = 0; i < updateLength; ++i) {
      RateLimiterNoEvents.Config memory configUpdate = rateLimiterUpdates.rateLimiterConfigs[i];
      uint64 chainSelector = rateLimiterUpdates.chainSelectors[i];

      RateLimiterNoEvents.TokenBucket memory tokenBucket = s_rateLimitersByChainSelector[chainSelector];
      uint32 lastUpdated = tokenBucket.lastUpdated;

      if (lastUpdated == 0) {
        // Token bucket needs to be newly added
        s_rateLimitersByChainSelector[chainSelector] = RateLimiterNoEvents.TokenBucket({
          rate: configUpdate.rate,
          capacity: configUpdate.capacity,
          tokens: configUpdate.capacity,
          lastUpdated: uint32(block.timestamp),
          isEnabled: configUpdate.isEnabled
        });
      } else {
        s_rateLimitersByChainSelector[chainSelector]._setTokenBucketConfig(configUpdate);
      }
      emit RateLimiterConfigUpdated(chainSelector, configUpdate);
    }
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Gets the token limit admin address.
  /// @return the token limit admin address.
  function getTokenLimitAdmin() external view returns (address) {
    return s_admin;
  }

  /// @notice Sets the token limit admin address.
  /// @param newAdmin the address of the new admin.
  /// @dev setting this to address(0) indicates there is no active admin.
  function setAdmin(address newAdmin) external onlyAdminOrOwner {
    _setAdmin(newAdmin);
  }

  /// @notice Sets the token limit admin address.
  /// @param newAdmin the address of the new admin.
  /// @dev setting this to address(0) indicates there is no active admin.
  function _setAdmin(address newAdmin) internal {
    s_admin = newAdmin;
    emit AdminSet(newAdmin);
  }

  /// @notice a modifier that allows the owner or the s_tokenLimitAdmin call the functions
  /// it is applied to.
  modifier onlyAdminOrOwner() {
    if (msg.sender != owner() && msg.sender != s_admin) revert RateLimiterNoEvents.OnlyCallableByAdminOrOwner();
    _;
  }
}
