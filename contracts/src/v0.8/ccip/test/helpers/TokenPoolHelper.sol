// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/TokenPool.sol";

contract TokenPoolHelper is TokenPool {
  event LockOrBurn(uint256 amount);
  event ReleaseOrMint(address indexed recipient, uint256 amount);
  event AssertionPassed();

  constructor(IERC20 token) TokenPool(token) {}

  function lockOrBurn(uint256 amount) external override {
    emit LockOrBurn(amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override {
    emit ReleaseOrMint(recipient, amount);
  }
}
