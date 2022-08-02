// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/interfaces/PoolInterface.sol";

contract MockPool is PoolInterface {
  // Unique ID used in tests
  uint256 public s_uid;

  constructor(uint256 uid) {
    s_uid = uid;
  }

  function lockOrBurn(uint256 amount) external override {
    emit Locked(msg.sender, amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override {
    emit Released(msg.sender, recipient, amount);
  }

  function getToken() external view override returns (IERC20 pool) {}

  function setLockOrBurnBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external override {}

  function setReleaseOrMintBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external override {}

  function getLockOrBurnBucket() external view override returns (TokenLimits.TokenBucket memory) {
    return TokenLimits.constructTokenBucket(0, 0, true);
  }

  function getReleaseOrMintBucket() external view override returns (TokenLimits.TokenBucket memory) {
    return TokenLimits.constructTokenBucket(0, 0, true);
  }

  function pause() external override {}

  function unpause() external override {}
}
