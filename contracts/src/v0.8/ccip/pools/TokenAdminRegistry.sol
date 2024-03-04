// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

contract TokenAdminRegistry is ITokenAdminRegistry, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;

  error OnlyRegistryModule(address sender);
  error OnlyAdministrator(address sender, address localToken);
  error AlreadyRegistered(address localToken, address currentAdministrator);
  error TokenHasNoPool(address token);

  event AdministratorRegistered(address indexed localToken, address indexed administrator);

  struct TokenConfig {
    address administrator; // ────────────────╮ the current administrator of the token
    bool isPermissionedAdmin; //              │ if true, this administrator has been configured by the CCIP owner
    //                                        │ and it could have elevated permissions.
    bool allowPermissionlessReRegistration; //│ if true, the token can be re-registered without the administrator's signature
    bool isRegistered; // ────────────────────╯ if true, the token is registered in the registry
    address tokenPool; // the token pool for this token. Can be address(0) if not deployed or not configured.
  }

  mapping(address token => TokenConfig) internal s_tokenConfig;

  EnumerableSet.AddressSet internal s_RegistryModules;

  /// @notice Returns all pools for the given tokens.
  /// @dev Will return address(0) for tokens that do not have a pool.
  function getPools(address[] calldata tokens) external view returns (address[] memory) {
    address[] memory pools = new address[](tokens.length);
    for (uint256 i = 0; i < tokens.length; ++i) {
      pools[i] = s_tokenConfig[tokens[i]].tokenPool;
    }
    return pools;
  }

  /// @inheritdoc ITokenAdminRegistry
  function getPool(address token) external view returns (address) {
    address pool = s_tokenConfig[token].tokenPool;
    if (pool == address(0)) {
      revert TokenHasNoPool(token);
    }

    return pool;
  }

  /// @notice Public getter to check for permissions of an administrator
  function isAdministrator(address localToken, address administrator) public view returns (bool) {
    return s_tokenConfig[localToken].administrator == administrator;
  }

  // ================================================================
  // │                    Administrator config                      │
  // ================================================================

  /// @notice Resisters a new local administrator for a token.
  function registerAdministrator(address localToken, address administrator) external {
    // Only allow permissioned registry modules to register administrators
    if (!s_RegistryModules.contains(msg.sender)) {
      revert OnlyRegistryModule(msg.sender);
    }
    TokenConfig storage config = s_tokenConfig[localToken];

    if (config.isRegistered && !config.allowPermissionlessReRegistration) {
      revert AlreadyRegistered(localToken, config.administrator);
    }

    // If the token is not registered yet, or if re-registration is permitted, register the new administrator
    config.administrator = administrator;
    config.isRegistered = true;

    emit AdministratorRegistered(localToken, administrator);
  }

  /// @notice Registers a local administrator for a token. This will overwrite any potential current administrator
  /// and set the permissionedAdmin to true.
  function registerAdministratorPermissioned(address localToken, address administrator) external onlyOwner {
    TokenConfig storage config = s_tokenConfig[localToken];

    config.administrator = administrator;
    config.isRegistered = true;
    config.isPermissionedAdmin = true;

    emit AdministratorRegistered(localToken, administrator);
  }
}
