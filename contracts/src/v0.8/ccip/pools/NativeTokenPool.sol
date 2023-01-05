// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {SafeERC20} from "../../vendor/SafeERC20.sol";

/**
 * @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
 * @dev One token per NativeTokenPool.
 */
contract NativeTokenPool is TokenPool {
  using SafeERC20 for IERC20;

  constructor(IERC20 token) TokenPool(token) {}

  /**
   * @notice Locks the token in the pool
   * @param amount Amount to lock
   */
  function lockOrBurn(uint256 amount) external override whenNotPaused assertLockOrBurn {
    emit Locked(msg.sender, amount);
  }

  /**
   * @notice Release tokens from the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to release
   */
  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused assertMintOrRelease {
    getToken().safeTransfer(recipient, amount);
    emit Released(msg.sender, recipient, amount);
  }
}
