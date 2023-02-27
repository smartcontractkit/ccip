// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IERC20} from "../../../vendor/IERC20.sol";

interface IBurnMintERC20 is IERC20 {
  /// @dev Creates `amount` tokens and assigns them to `account`, increasing
  /// the total supply.
  function mint(address account, uint256 amount) external;

  /// @dev Destroys `amount` tokens from `account`, reducing the
  /// total supply.
  function burn(address account, uint256 amount) external;
}
