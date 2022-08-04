// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

/**
 * @notice This library enables token rate limiting using a `TokenBucket`.
 * The bucket holds the number of tokens that can be transferred at any
 * given time. It has:
 *  - capacity: maximum number of tokens possible
 *  - rate: rate at which the bucket refills per second
 *  - tokens: current number of tokens in the bucket
 *  - lastUpdated: timestamp of the last refill
 */
library TokenLimits {
  // Token Bucket used for rate limiting
  struct TokenBucket {
    uint256 rate;
    uint256 capacity;
    uint256 tokens;
    uint256 lastUpdated;
  }

  error TimeError();
  error BucketOverfilled();

  /**
   * @notice Create a fresh token bucket
   * @param rate Refill rate
   * @param capacity Maximum capacity of the bucket
   * @return tokenBucket
   */
  function constructTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) internal view returns (TokenBucket memory) {
    uint256 tokens = full ? capacity : 0;
    return TokenBucket({rate: rate, capacity: capacity, tokens: tokens, lastUpdated: block.timestamp});
  }

  /**
   * @notice Remove tokens from the buck if possible.
   * @dev This acts upon a storage variable in the calling contract.
   * @param bucket token bucket (MUST BE STORAGE)
   * @param tokens number of tokens
   * @return tokens removed (true if removed, false otherwise)
   */
  function remove(TokenBucket storage bucket, uint256 tokens) internal returns (bool) {
    // Refill the bucket if possible
    update(bucket);
    // Remove tokens if available in bucket
    if (bucket.tokens < tokens) return false;
    bucket.tokens -= tokens;
    return true;
  }

  /**
   * @notice Update the tokens in the bucket
   * @dev Uses the `rate` and block timestamp to refill the bucket. The bucket will not start
   * refilling within the same block that it was removed from.
   * @dev This acts upon a storage variable in the calling contract.
   * @param bucket token bucket (MUST BE STORAGE)
   */
  function update(TokenBucket storage bucket) internal {
    // Revert if the tokens in the bucket exceed its capacity
    if (bucket.tokens > bucket.capacity) revert BucketOverfilled();
    // Return if there's nothing to update
    if (bucket.tokens == bucket.capacity) return;
    uint256 timeNow = block.timestamp;
    if (timeNow < bucket.lastUpdated) revert TimeError();
    uint256 difference = timeNow - bucket.lastUpdated;
    bucket.tokens = min(bucket.capacity, bucket.tokens + difference * bucket.rate);
    bucket.lastUpdated = timeNow;
  }

  /**
   * @notice Return the smallest of two integers
   * @param a first int
   * @param b second int
   * @return smallest
   */
  function min(uint256 a, uint256 b) private pure returns (uint256) {
    return a < b ? a : b;
  }
}
