// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPool} from "./TokenPool.sol";

import {ERC20} from "../../vendor/ERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";

/**
 * @notice This pool mints and burns its own tokens, representing a wrapped form of the native token
 * on a source chain - similar to WBTC.
 */
contract WrappedTokenPool is TokenPool, ERC20 {
  constructor(string memory name, string memory symbol) TokenPool(IERC20(address(this))) ERC20(name, symbol) {}

  /**
   * @notice Burn the token in the pool
   * @param amount Amount to burn
   */
  function lockOrBurn(uint256 amount) external override whenNotPaused assertLockOrBurn {
    _burn(address(this), amount);
    emit Burned(msg.sender, amount);
  }

  /**
   * @notice Mint tokens from the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to mint
   */
  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused assertMintOrRelease {
    _mint(recipient, amount);
    emit Minted(msg.sender, recipient, amount);
  }
}
