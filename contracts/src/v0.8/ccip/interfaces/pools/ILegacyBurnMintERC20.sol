// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol";

interface ILegacyBurnMintERC20 is IERC20 {
  /// @notice Mints new tokens for a given address.
  /// @param account The address to mint the new tokens to.
  /// @param amount The number of tokens to be minted.
  /// @dev this function increases the total supply.
  function mint(address account, uint256 amount) external;

  /// @notice Burns tokens from the sender.
  /// @param account The address to burn tokens from.
  /// @param amount The number of tokens to be burned.
  /// @dev this function decreases the total supply.
  function burn(address account, uint256 amount) external;

  /// @notice Burns tokens from a given address..
  /// @param account The address to burn tokens from.
  /// @param amount The number of tokens to be burned.
  /// @dev this function decreases the total supply.
  function burnFrom(address account, uint256 amount) external;
}
