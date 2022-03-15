// SPDX-License-Identifier: MIT

pragma solidity 0.8.12;

import "./TokenPool.sol";

/**
 * @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
 * @dev One token per NativeTokenPool.
 */
contract NativeTokenPool is TokenPool {
  using SafeERC20 for IERC20;

  constructor(
    IERC20 token,
    BucketConfig memory lockConfig,
    BucketConfig memory releaseConfig
  ) TokenPool(token, lockConfig, releaseConfig) {}

  /**
   * @notice Locks the token in the pool
   * @param depositor Token holder address
   * @param amount Amount to lock
   */
  function lockOrBurn(address depositor, uint256 amount) external override whenNotPaused assertLockOrBurn(amount) {
    getToken().safeTransferFrom(depositor, address(this), amount);
    emit Locked(msg.sender, depositor, amount);
  }

  /**
   * @notice Release tokens fromm the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to release
   */
  function releaseOrMint(address recipient, uint256 amount)
    external
    override
    whenNotPaused
    assertMintOrRelease(amount)
  {
    getToken().safeTransfer(recipient, amount);
    emit Released(msg.sender, recipient, amount);
  }
}
