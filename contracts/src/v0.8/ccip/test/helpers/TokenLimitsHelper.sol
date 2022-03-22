// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../utils/TokenLimits.sol";

contract TokenLimitsHelper {
  using TokenLimits for TokenLimits.TokenBucket;

  TokenLimits.TokenBucket public s_bucket;

  event RemovalSuccess(bool success);

  function constructTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) public {
    s_bucket = TokenLimits.constructTokenBucket(rate, capacity, full);
  }

  function alterCapacity(uint256 newCapacity) public {
    s_bucket.capacity = newCapacity;
  }

  function update() public {
    s_bucket.update();
  }

  function remove(uint256 tokens) public returns (bool removed) {
    removed = s_bucket.remove(tokens);
    emit RemovalSuccess(removed);
  }
}
