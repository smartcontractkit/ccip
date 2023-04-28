// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OwnerIsCreator} from "./OwnerIsCreator.sol";
import {Client} from "./libraries/Client.sol";
import {RateLimiter} from "./libraries/RateLimiter.sol";
import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {IERC20} from "../vendor/IERC20.sol";

contract AggregateRateLimiter is OwnerIsCreator {
  using RateLimiter for RateLimiter.TokenBucket;

  error PriceNotFoundForToken(address token);
  event AdminSet(address newAdmin);

  // The address of the token limit admin that has the same permissions as the owner.
  address internal s_admin;

  // The token bucket object that contains the bucket state.
  RateLimiter.TokenBucket private s_rateLimiter;

  /// @param config The RateLimiter.Config containing the capacity and refill rate
  /// of the bucket, plus the admin address.
  constructor(RateLimiter.Config memory config) {
    s_rateLimiter = RateLimiter.TokenBucket({
      rate: config.rate,
      capacity: config.capacity,
      tokens: config.capacity,
      lastUpdated: uint40(block.timestamp),
      isEnabled: config.isEnabled
    });
  }

  function _rateLimitValue(Client.EVMTokenAmount[] memory tokenAmounts, IPriceRegistry priceRegistry) internal {
    uint256 numberOfTokens = tokenAmounts.length;

    uint256 value = 0;
    for (uint256 i = 0; i < numberOfTokens; ++i) {
      uint256 pricePerToken = priceRegistry.getTokenPrice(tokenAmounts[i].token).value;
      if (pricePerToken == 0) revert PriceNotFoundForToken(tokenAmounts[i].token);
      value += tokenAmounts[i].amount * pricePerToken;
    }

    s_rateLimiter._consume(value);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentRateLimiterState() public view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  /// @notice Sets the rate limited config.
  /// @param config The new rate limiter config.
  /// @dev should only be callable by the owner or token limit admin.
  function setRateLimiterConfig(RateLimiter.Config memory config) public requireAdminOrOwner {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  // ================================================================
  // |                           Access                             |
  // ================================================================

  /// @notice Gets the token limit admin address.
  /// @return the token limit admin address.
  function getTokenLimitAdmin() public view returns (address) {
    return s_admin;
  }

  /// @notice Sets the token limit admin address.
  /// @param newAdmin the address of the new admin.
  /// @dev setting this to address(0) indicates there is no active admin.
  function setAdmin(address newAdmin) public requireAdminOrOwner {
    s_admin = newAdmin;
    emit AdminSet(newAdmin);
  }

  /// @notice a modifier that allows the owner or the s_tokenLimitAdmin call the functions
  /// it is applied to.
  modifier requireAdminOrOwner() {
    if (msg.sender != owner() && msg.sender != s_admin) revert RateLimiter.OnlyCallableByAdminOrOwner();
    _;
  }
}
