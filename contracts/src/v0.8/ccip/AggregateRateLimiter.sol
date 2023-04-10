// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OwnerIsCreator} from "./OwnerIsCreator.sol";
import {Client} from "./libraries/Client.sol";
import {RateLimiter} from "./libraries/RateLimiter.sol";

import {IERC20} from "../vendor/IERC20.sol";

contract AggregateRateLimiter is OwnerIsCreator {
  using RateLimiter for RateLimiter.TokenBucket;

  error TokensAndPriceLengthMismatch();
  error PriceNotFoundForToken(address token);
  error AddressCannotBeZero();

  event TokenPriceChanged(address token, uint256 newPrice);

  // The address of the token limit admin that has the same permissions as the owner.
  address internal s_admin;

  // A mapping of token => tokenPrice
  mapping(IERC20 => uint256) private s_priceByToken;
  // The tokens that have a set price
  IERC20[] private s_allowedTokens;

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

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function currentTokenBucketState() public view returns (RateLimiter.TokenBucket memory) {
    return s_rateLimiter._currentTokenBucketState();
  }

  function _rateLimitValue(Client.EVMTokenAmount[] memory tokenAmounts) internal {
    uint256 value = 0;
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      uint256 pricePerToken = s_priceByToken[IERC20(tokenAmounts[i].token)];
      if (pricePerToken == 0) revert PriceNotFoundForToken(tokenAmounts[i].token);
      value += pricePerToken * tokenAmounts[i].amount;
    }

    s_rateLimiter._consume(value);
  }

  /// @notice Sets the rate limited config.
  /// @param config The new rate limiter config.
  /// @dev should only be callable by the owner or token limit admin.
  function setRateLimiterConfig(RateLimiter.Config memory config) public requireAdminOrOwner {
    s_rateLimiter._setTokenBucketConfig(config);
  }

  /// @notice Gets the set prices for the given IERC20s.
  /// @param tokens The tokens to get the price of.
  /// @return prices The current prices of the token.
  function getPricesForTokens(IERC20[] memory tokens) public view returns (uint256[] memory prices) {
    uint256 numberOfTokens = tokens.length;
    prices = new uint256[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      prices[i] = s_priceByToken[tokens[i]];
    }

    return prices;
  }

  /// @notice Sets the prices of the given IERC20 tokens to the given prices.
  /// @param tokens The tokens for which the price will be set.
  /// @param prices The new prices of the given tokens.
  /// @dev if any previous prices were set for a number of given tokens, these will
  /// be overwritten. Previously set prices for tokens that are not present in subsequent
  /// setPrices calls will *not* be reset to zero but will be left unchanged.
  /// @dev should only be callable by the owner or token limit admin.
  function setPrices(IERC20[] memory tokens, uint256[] memory prices) public requireAdminOrOwner {
    uint256 newTokenLength = tokens.length;
    if (newTokenLength != prices.length) revert TokensAndPriceLengthMismatch();

    // Remove all old entries
    uint256 setTokensLength = s_allowedTokens.length;
    for (uint256 i = 0; i < setTokensLength; ++i) {
      delete s_priceByToken[s_allowedTokens[i]];
    }

    for (uint256 i = 0; i < newTokenLength; ++i) {
      IERC20 token = tokens[i];
      if (token == IERC20(address(0))) revert AddressCannotBeZero();
      s_priceByToken[token] = prices[i];
      emit TokenPriceChanged(address(token), prices[i]);
    }

    s_allowedTokens = tokens;
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
  function setAdmin(address newAdmin) public requireAdminOrOwner {
    s_admin = newAdmin;
  }

  /// @notice a modifier that allows the owner or the s_tokenLimitAdmin call the functions
  /// it is applied to.
  modifier requireAdminOrOwner() {
    if (msg.sender != owner() && msg.sender != s_admin) revert RateLimiter.OnlyCallableByAdminOrOwner();
    _;
  }
}
