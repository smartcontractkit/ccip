// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAggregateRateLimiter} from "../interfaces/rateLimiter/IAggregateRateLimiter.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Client} from "../models/Client.sol";

import {IERC20} from "../../vendor/IERC20.sol";

contract AggregateRateLimiter is IAggregateRateLimiter, OwnerIsCreator {
  // The address of the token limit admin that has the same permissions as the owner.
  address private s_tokenLimitAdmin;

  // A mapping of token => tokenPrice
  mapping(IERC20 => uint256) private s_priceByToken;
  // The tokens that have a set price
  IERC20[] private s_allowedTokens;

  // The token bucket object that contains the bucket state.
  TokenBucket private s_tokenBucket;

  /// @param config The RateLimiterConfig containing the capacity and refill rate
  /// of the bucket, plus the admin address.
  constructor(RateLimiterConfig memory config) {
    s_tokenLimitAdmin = config.admin;
    s_tokenBucket = TokenBucket({
      rate: config.rate,
      capacity: config.capacity,
      tokens: config.capacity,
      lastUpdated: block.timestamp
    });
  }

  /// @inheritdoc IAggregateRateLimiter
  function getTokenLimitAdmin() public view returns (address) {
    return s_tokenLimitAdmin;
  }

  /// @inheritdoc IAggregateRateLimiter
  function setTokenLimitAdmin(address newAdmin) public onlyOwner {
    s_tokenLimitAdmin = newAdmin;
  }

  /// @inheritdoc IAggregateRateLimiter
  function calculateCurrentTokenBucketState() public view returns (TokenBucket memory) {
    TokenBucket memory bucket = s_tokenBucket;

    // We update the bucket to reflect the status at the exact time of the
    // call. This means to might need to refill a part of the bucket based
    // on the time that has passed since the last update.
    uint256 difference = block.timestamp - bucket.lastUpdated;

    // Overflow doesn't happen here because bucket.rate is <= type(uint208).max
    // leaving 48 bits for the time difference. 2 ** 48 seconds = 9 million years.
    bucket.tokens = _min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = block.timestamp;
    return bucket;
  }

  /// @inheritdoc IAggregateRateLimiter
  function setRateLimiterConfig(RateLimiterConfig memory config) public requireAdminOrOwner {
    // First update the bucket to make sure the proper rate is used for all the time
    // up until the config change.
    _update(s_tokenBucket);

    s_tokenBucket.capacity = config.capacity;
    s_tokenBucket.rate = config.rate;
    s_tokenBucket.tokens = _min(config.capacity, s_tokenBucket.tokens);

    emit ConfigChanged(config.capacity, config.rate);
  }

  /// @inheritdoc IAggregateRateLimiter
  function getPricesForTokens(IERC20[] memory tokens) public view returns (uint256[] memory prices) {
    uint256 numberOfTokens = tokens.length;
    prices = new uint256[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      prices[i] = s_priceByToken[tokens[i]];
    }

    return prices;
  }

  /// @inheritdoc IAggregateRateLimiter
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

  /// @notice _removeTokens removes the given token values from the pool, lowering the
  /// value allowed to be transferred for subsequent calls. It will use the
  /// s_priceByToken mapping to determine value in a standardised unit.
  /// @param tokenAmounts The tokenAmounts that are send across the bridge. All
  /// of the tokens need to have a corresponding price set in s_priceByToken.
  /// @dev Reverts when a token price is not found or when the tx value exceeds the
  /// amount allowed in the bucket.
  /// @dev Will only remove and therefore emit removal of value if the value is > 0.
  function _removeTokens(Client.EVMTokenAmount[] memory tokenAmounts) internal {
    uint256 value = 0;
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      uint256 pricePerToken = s_priceByToken[IERC20(tokenAmounts[i].token)];
      if (pricePerToken == 0) revert PriceNotFoundForToken(tokenAmounts[i].token);
      value += pricePerToken * tokenAmounts[i].amount;
    }

    // If there is no value to remove skip this step to reduce gas usage
    if (value > 0) {
      // Refill the bucket if possible
      // This mutates s_tokenBucket in storage
      _update(s_tokenBucket);

      if (s_tokenBucket.capacity < value) revert ValueExceedsCapacity(s_tokenBucket.capacity, value);
      if (s_tokenBucket.tokens < value) {
        // Seconds wait required until the bucket is refilled enough to accept this value
        uint256 waitInSeconds = (value - s_tokenBucket.tokens) / s_tokenBucket.rate;
        revert ValueExceedsAllowedThreshold(waitInSeconds);
      }

      s_tokenBucket.tokens -= value;
      emit TokensRemovedFromBucket(value);
    }
  }

  function _update(TokenBucket storage bucket) internal {
    // Return if there's nothing to update
    if (bucket.tokens == bucket.capacity || bucket.lastUpdated == block.timestamp) return;
    // Revert if the tokens in the bucket exceed its capacity
    if (bucket.tokens > bucket.capacity) revert BucketOverfilled();
    uint256 difference = block.timestamp - bucket.lastUpdated;
    bucket.tokens = _min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = block.timestamp;
  }

  /// @notice Return the smallest of two integers
  /// @param a first int
  /// @param b second int
  /// @return smallest
  function _min(uint256 a, uint256 b) internal pure returns (uint256) {
    return a < b ? a : b;
  }

  /// @notice a modifier that allows the owner or the s_tokenLimitAdmin call the functions
  /// it is applied to.
  modifier requireAdminOrOwner() {
    if (msg.sender != owner() && msg.sender != s_tokenLimitAdmin) revert OnlyCallableByAdminOrOwner();
    _;
  }
}
