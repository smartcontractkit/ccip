// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/pools/IPool.sol";

contract MockPool is IPool {
  // Unique ID used in tests
  uint256 public s_uid;

  constructor(uint256 uid) {
    s_uid = uid;
  }

  function lockOrBurn(uint256 amount, address) external override {
    emit Locked(msg.sender, amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override {
    emit Released(msg.sender, recipient, amount);
  }

  function applyRampUpdates(RampUpdate[] memory, RampUpdate[] memory) public {}

  function getToken() external view override returns (IERC20 pool) {}

  function pause() external override {}

  function unpause() external override {}
}
