pragma solidity ^0.8.0;

import {TokenPool} from "../../pools/TokenPool.sol";
import {IERC20} from "../../../vendor/IERC20.sol";

contract CustomTokenPool is TokenPool {
  event SynthBurned(uint256 amount);
  event SynthMinted(uint256 amount);

  constructor(IERC20 token) TokenPool(token) {}

  /**
   * @notice Locks the token in the pool
   * @param amount Amount to lock
   */
  function lockOrBurn(uint256 amount, address) external override whenNotPaused validateOwnerOrOnRamp {
    emit SynthBurned(amount);
  }

  /**
   * @notice Release tokens from the pool to the recipient
   * @param amount Amount to release
   */
  function releaseOrMint(address, uint256 amount) external override whenNotPaused validateOwnerOrOffRamp {
    emit SynthMinted(amount);
  }
}
