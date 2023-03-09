// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IBurnMintERC20} from "../interfaces/pools/IBurnMintERC20.sol";

import {TokenPool} from "./TokenPool.sol";

/// @notice This pool mints and burns a 3rd-party token.
contract BurnMintTokenPool is TokenPool {
  constructor(IBurnMintERC20 token) TokenPool(token) {
    token.approve(address(this), 2**256 - 1);
  }

  /// @notice Burn the token in the pool
  /// @param amount Amount to burn
  function lockOrBurn(uint256 amount, address) external override whenNotPaused onlyOnRamp {
    IBurnMintERC20(address(i_token)).burn(address(this), amount);
    emit Burned(msg.sender, amount);
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @param recipient Recipient address
  /// @param amount Amount to mint
  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused onlyOffRamp {
    IBurnMintERC20(address(i_token)).mint(recipient, amount);
    emit Minted(msg.sender, recipient, amount);
  }
}
