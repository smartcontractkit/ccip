// SPDX-License-Identifier: MIT

pragma solidity 0.8.13;

import "../../pools/TokenPool.sol";

contract TokenPoolHelper is TokenPool {
  event LockOrBurn(uint256 amount);
  event ReleaseOrMint(address indexed recipient, uint256 amount);
  event AssertionPassed();

  constructor(
    IERC20 token,
    uint256 lockBucketRate,
    uint256 lockBucketCapacity,
    uint256 releaseBucketRate,
    uint256 releaseBucketCapacity
  )
    TokenPool(
      token,
      BucketConfig({rate: lockBucketRate, capacity: lockBucketCapacity}),
      BucketConfig({rate: releaseBucketRate, capacity: releaseBucketCapacity})
    )
  {}

  function lockOrBurn(uint256 amount) external override {
    emit LockOrBurn(amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override {
    emit ReleaseOrMint(recipient, amount);
  }

  function assertLockOrBurnModifier(uint256 amount) external assertLockOrBurn(amount) {
    emit AssertionPassed();
  }

  function assertMintOrReleaseModifier(uint256 amount) external assertMintOrRelease(amount) {
    emit AssertionPassed();
  }
}
