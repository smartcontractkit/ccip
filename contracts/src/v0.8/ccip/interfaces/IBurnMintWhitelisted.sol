// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

interface IBurnMintWhitelisted is IERC20 {
  // @notice Burns the specified amount of tokens from the caller's account.
  function burnFromWhitelistedContract(uint256 amount) external;

  // @notice Mints the specified amount of tokens from the caller's account.
  function mintFromWhitelistedContract(uint256 amount) external;
}
