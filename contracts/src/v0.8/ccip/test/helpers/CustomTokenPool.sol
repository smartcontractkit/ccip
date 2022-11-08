pragma solidity ^0.8.0;

import {TokenPool, SafeERC20, IERC20} from "../../pools/TokenPool.sol";

contract CustomTokenPool is TokenPool {
  using SafeERC20 for IERC20;
  event SynthBurned(uint256 amount);
  event SynthMinted(uint256 amount);

  constructor(IERC20 token) TokenPool(token) {}

  /**
   * @notice Locks the token in the pool
   * @param amount Amount to lock
   */
  function lockOrBurn(uint256 amount) external override whenNotPaused assertLockOrBurn {
    emit SynthBurned(amount);
  }

  /**
   * @notice Release tokens fromm the pool to the recipient
   * @param amount Amount to release
   */
  function releaseOrMint(address, uint256 amount) external override whenNotPaused assertMintOrRelease {
    emit SynthMinted(amount);
  }
}
