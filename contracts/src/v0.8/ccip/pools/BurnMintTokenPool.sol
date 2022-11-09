// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPool} from "./TokenPool.sol";
import {IBurnMintERC20} from "./IBurnMintERC20.sol";

/**
 * @notice This pool mints and burns a 3rd-party token.
 */
contract BurnMintTokenPool is TokenPool {
  constructor(IBurnMintERC20 token) TokenPool(token) {}

  /**
   * @notice Burn the token in the pool
   * @param amount Amount to burn
   */
  function lockOrBurn(uint256 amount) external override whenNotPaused assertLockOrBurn {
    IBurnMintERC20(address(i_token)).burn(address(this), amount);
    emit Burned(msg.sender, amount);
  }

  /**
   * @notice Mint tokens from the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to mint
   */
  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused assertMintOrRelease {
    IBurnMintERC20(address(i_token)).mint(recipient, amount);
    emit Minted(msg.sender, recipient, amount);
  }
}
