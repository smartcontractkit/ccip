// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../../pools/TokenPool.sol";

contract TokenPoolHelper is TokenPool {
  event LockOrBurn(uint256 amount);
  event ReleaseOrMint(address indexed recipient, uint256 amount);
  event AssertionPassed();

  constructor(
    IERC20 token,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig
  ) TokenPool(token, allowlist, rateLimiterConfig) {}

  function lockOrBurn(
    address,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override returns (bytes memory) {
    emit LockOrBurn(amount);
    return "";
  }

  function releaseOrMint(bytes memory, address receiver, uint256 amount, uint64, bytes memory) external override {
    emit ReleaseOrMint(receiver, amount);
  }
}
