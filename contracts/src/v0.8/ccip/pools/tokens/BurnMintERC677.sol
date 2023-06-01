// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../interfaces/pools/IBurnMintERC20.sol";

import {ERC677} from "./ERC677.sol";

import {ERC20Burnable} from "../../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/extensions/ERC20Burnable.sol";
import {AccessControlEnumerable} from "../../../vendor/openzeppelin-solidity/v4.8.0/access/AccessControlEnumerable.sol";

contract BurnMintERC677 is IBurnMintERC20, ERC677, ERC20Burnable, AccessControlEnumerable {
  bytes32 private constant MINTER_ROLE = keccak256("MINTER_ROLE");
  bytes32 private constant BURNER_ROLE = keccak256("BURNER_ROLE");

  uint8 private immutable i_decimals;

  constructor(string memory name, string memory symbol, uint8 decimals_) ERC677(name, symbol) {
    i_decimals = decimals_;
    _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
  }

  function decimals() public view virtual override returns (uint8) {
    return i_decimals;
  }

  // ================================================================
  // |                      Burning & minting                       |
  // ================================================================

  /// @inheritdoc ERC20Burnable
  function burn(uint256 amount) public override(IBurnMintERC20, ERC20Burnable) onlyRole(BURNER_ROLE) {
    super.burn(amount);
  }

  /// @inheritdoc ERC20Burnable
  function burnFrom(
    address account,
    uint256 amount
  ) public override(IBurnMintERC20, ERC20Burnable) onlyRole(BURNER_ROLE) {
    super.burnFrom(account, amount);
  }

  /// @notice Creates `amount` tokens and assigns them to `account`, increasing
  /// the total supply.
  /// @dev Emits a {Transfer} event with `from` set to the zero address.
  /// @dev `account` cannot be the zero address.
  function mint(address account, uint256 amount) external override onlyRole(MINTER_ROLE) {
    _mint(account, amount);
  }

  // ================================================================
  // |                            Roles                             |
  // ================================================================

  /// @notice grants both mint and burn roles to `account`.
  /// @dev uses the public `grantRole` function internally to manage
  /// role granting permissions.
  function grantMintAndBurnRoles(address account) external {
    grantRole(MINTER_ROLE, account);
    grantRole(BURNER_ROLE, account);
  }

  /// @notice Returns the minter role hash
  function getMinterRole() external pure returns (bytes32) {
    return MINTER_ROLE;
  }

  /// @notice Returns the burner role hash
  function getBurnerRole() external pure returns (bytes32) {
    return BURNER_ROLE;
  }
}
