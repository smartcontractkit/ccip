// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC677} from "./IERC677.sol";
import {IERC677Receiver} from "../../interfaces/IERC677Receiver.sol";

import {BurnMintERC20} from "../ERC20/BurnMintERC20.sol";

/// @notice A basic ERC677 compatible token contract with burn and minting roles.
/// @dev The total supply can be limited during deployment.
contract BurnMintERC677 is BurnMintERC20, IERC677 {
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals_,
    uint256 maxSupply_
  ) BurnMintERC20(name, symbol, decimals_, maxSupply_) {}

  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == type(IERC677).interfaceId || super.supportsInterface(interfaceId);
  }

  /// @inheritdoc IERC677
  /// @dev This function has been duplicated from ERC677.sol since functionality cannot be inherited due to
  /// dual imports of ERC20 with BurnMintERC20.sol
  function transferAndCall(address to, uint256 amount, bytes memory data) public returns (bool success) {
    super.transfer(to, amount);
    emit Transfer(msg.sender, to, amount, data);
    if (to.code.length > 0) {
      IERC677Receiver(to).onTokenTransfer(msg.sender, amount, data);
    }
    return true;
  }
}
