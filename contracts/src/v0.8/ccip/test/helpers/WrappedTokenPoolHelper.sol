// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/WrappedTokenPool.sol";

contract WrappedTokenPoolHelper is WrappedTokenPool {
  constructor(
    string memory name,
    string memory symbol,
    uint256 burnBucketRate,
    uint256 burnBucketCapacity,
    uint256 mintBucketRate,
    uint256 mintBucketCapacity
  ) WrappedTokenPool(name, symbol) {}

  function mint(address account, uint256 amount) public onlyOwner {
    _mint(account, amount);
  }
}
