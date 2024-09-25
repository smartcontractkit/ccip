// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {BurnMintERC20} from "../../shared/token/ERC20/BurnMintERC20.sol";

/// @notice A basic ERC20 compatible token contract with burn and minting roles.
/// @dev The total supply can be limited during deployment.
contract FactoryBurnMintERC20 is BurnMintERC20 {
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals_,
    uint256 maxSupply_,
    uint256 preMint_,
    address newOwner_
  ) BurnMintERC20(name, symbol, decimals_, maxSupply_) {
    i_decimals = decimals_;
    i_maxSupply = maxSupply_;

    s_ccipAdmin = newOwner_;

    // Mint the initial supply to the new Owner, saving gas by not calling if the mint amount is zero
    if (preMint_ != 0) _mint(newOwner_, preMint_);

    // Grant the deployer the minter and burner roles. This contract is expected to be deployed by a factory
    // contract that will transfer ownership to the correct address after deployment, so granting minting and burning
    // privileges here saves gas by not requiring two transactions.
    grantMintRole(newOwner_);
    grantBurnRole(newOwner_);
  }
}
