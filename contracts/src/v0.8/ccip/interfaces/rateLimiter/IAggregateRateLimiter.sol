// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC20} from "../../../vendor/IERC20.sol";

interface IAggregateRateLimiter {
  error OnlyCallableByAdminOrOwner();
  error TokensAndPriceLengthMismatch();
  error ValueExceedsAllowedThreshold(uint256 waitInSeconds);
  error ValueExceedsCapacity(uint256 capacity, uint256 requested);
  error PriceNotFoundForToken(address token);
  error AddressCannotBeZero();
  error BucketOverfilled();
  error RefillRateTooHigh();

  event ConfigChanged(uint256 capacity, uint256 rate);
  event TokensRemovedFromBucket(uint256 tokens);
  event TokenPriceChanged(address token, uint256 newPrice);

  struct TokenBucket {
    uint256 rate;
    uint256 capacity;
    uint256 tokens;
    uint256 lastUpdated;
  }

  struct RateLimiterConfig {
    uint256 rate;
    uint256 capacity;
  }

  /**
   * @notice Gets the token limit admin address
   */
  function getTokenLimitAdmin() external view returns (address);

  /**
   * @notice Sets the token limit admin address
   * @param newAdmin the address of the new admin.
   */
  function setTokenLimitAdmin(address newAdmin) external;

  /**
   * @notice Gets the token bucket with it's values for the block it was
   *          requested at.
   * @return The token bucket.
   */
  function calculateCurrentTokenBucketState() external view returns (TokenBucket memory);

  /**
   * @notice Sets the rate limited config.
   * @param config The new rate limiter config.
   * @dev should only be callable by the owner or token limit admin.
   * @dev the max rate is uint208.max
   */
  function setRateLimiterConfig(RateLimiterConfig memory config) external;

  /**
   * @notice Gets the set prices for the given IERC20s.
   * @param tokens The tokens to get the price of.
   * @return prices The current prices of the token.
   */
  function getPricesForTokens(IERC20[] memory tokens) external view returns (uint256[] memory prices);

  /**
   * @notice Sets the prices of the given IERC20 tokens to the given prices.
   * @param tokens The tokens for which the price will be set.
   * @param prices The new prices of the given tokens.
   * @dev if any previous prices were set for a number of given tokens, these
   *        will be overwritten. Previously set prices for tokens that are
   *        not present in subsequent setPrices calls will *not* be reset
   *        to zero but will be left unchanged.
   * @dev should only be callable by the owner or token limit admin.
   */
  function setPrices(IERC20[] memory tokens, uint256[] memory prices) external;
}
