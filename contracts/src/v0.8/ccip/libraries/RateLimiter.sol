// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Client} from "./Client.sol";

library RateLimiter {
  error BucketOverfilled();
  error ConsumingMoreThanMaxCapacity(uint256 capacity, uint256 requested);
  error RateLimitReached(uint256 waitInSeconds);
  error OnlyCallableByAdminOrOwner();

  event TokensConsumed(uint256 tokens);
  event ConfigChanged(Config config);

  struct TokenBucket {
    uint256 capacity; // Maximum number of tokens that can be in the bucket.
    uint256 tokens; // Current number of tokens that are in the bucket.
    uint208 rate; // // -----┐ Number of tokens per second that the bucket is refilled.
    uint40 lastUpdated; //   | Timestamp of the last token update.
    bool isEnabled; // ------┘ Indication whether the rate limiting is enabled or not
  }

  struct Config {
    bool isEnabled; // ---┐ Indicated whether the rate limiting is enabled
    uint208 rate; // -----┘ We only allow a refill rate of uint208 so we don't have to deal
    // with any overflows for the next ~9 million years. Any sensible rate is way below this value.
    uint256 capacity;
  }

  function _refill(TokenBucket storage bucket) private {
    // Return if there's nothing to update
    if (bucket.tokens == bucket.capacity || bucket.lastUpdated == block.timestamp) return;
    // Revert if the tokens in the bucket exceed its capacity
    if (bucket.tokens > bucket.capacity) revert BucketOverfilled();
    uint256 difference = block.timestamp - bucket.lastUpdated;
    bucket.tokens = _min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = uint40(block.timestamp);
  }

  /// @notice _consume removes the given tokens from the pool, lowering the
  /// rate tokens allowed to be consumed for subsequent calls.
  /// @param requestTokens The total tokens to be consumed from the bucket.
  /// @dev Reverts when requestTokens exceeds bucket capacity or available tokens in the bucket
  /// @dev emits removal of requestTokens if requestTokens is > 0
  function _consume(TokenBucket storage bucket, uint256 requestTokens) internal {
    // If there is no value to remove or rate limiting is turned off, skip this step to reduce gas usage
    if (!bucket.isEnabled || requestTokens == 0) {
      return;
    }

    // Refill the bucket if possible, this mutates bucket in storage
    _refill(bucket);

    if (bucket.capacity < requestTokens) revert ConsumingMoreThanMaxCapacity(bucket.capacity, requestTokens);
    if (bucket.tokens < requestTokens)
      // Seconds wait required until the bucket is refilled enough to accept this value
      revert RateLimitReached((requestTokens - bucket.tokens) / bucket.rate);

    bucket.tokens -= requestTokens;
    emit TokensConsumed(requestTokens);
  }

  /// @notice Gets the token bucket with its values for the block it was requested at.
  /// @return The token bucket.
  function _currentTokenBucketState(TokenBucket memory bucket) internal view returns (TokenBucket memory) {
    // We update the bucket to reflect the status at the exact time of the
    // call. This means to might need to refill a part of the bucket based
    // on the time that has passed since the last update.
    uint256 difference = block.timestamp - bucket.lastUpdated;

    // Overflow doesn't happen here because bucket.rate is <= type(uint208).max
    // leaving 48 bits for the time difference. 2 ** 48 seconds = 9 million years.
    bucket.tokens = _min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = uint40(block.timestamp);
    return bucket;
  }

  /// @notice Sets the rate limited config.
  /// @param bucket The token bucket
  /// @param config The new config
  function _setTokenBucketConfig(TokenBucket storage bucket, Config memory config) internal {
    // First update the bucket to make sure the proper rate is used for all the time
    // up until the config change.
    _refill(bucket);

    bucket.isEnabled = config.isEnabled;
    bucket.capacity = config.capacity;
    bucket.rate = config.rate;
    bucket.tokens = _min(config.capacity, bucket.tokens);
    emit ConfigChanged(config);
  }

  /// @notice Return the smallest of two integers
  /// @param a first int
  /// @param b second int
  /// @return smallest
  function _min(uint256 a, uint256 b) internal pure returns (uint256) {
    return a < b ? a : b;
  }
}
