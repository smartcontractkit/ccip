// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

// This contract has minimal functionality and minimal test coverage. It will be
// improved upon in future tickets.
contract TokenAdminRegistry is ITokenAdminRegistry, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;

  error OnlyRegistryModule(address sender);
  error OnlyAdministrator(address sender, address token);
  error OnlyPendingAdministrator(address sender, address token);
  error AlreadyRegistered(address token, address currentAdministrator);
  error UnsupportedToken(address token);

  event AdministratorRegistered(address indexed token, address indexed administrator);
  event PoolSet(address indexed token, address indexed previousPool, address indexed newPool);
  event AdministratorTransferRequested(address indexed token, address indexed currentAdmin, address indexed newAdmin);
  event AdministratorTransferred(address indexed token, address indexed newAdmin);

  struct TokenConfig {
    address pendingAdministrator; //            the address that is pending to become the new owner
    address administrator; // ────────────────╮ the current administrator of the token
    bool isPermissionedAdmin; //              │ if true, this administrator has been configured by the CCIP owner
    //                                        │ and it could have elevated permissions.
    bool allowPermissionlessReRegistration; //│ if true, the token can be re-registered without the administrator's signature
    bool isRegistered; // ────────────────────╯ if true, the token is registered in the registry
    address tokenPool; // the token pool for this token. Can be address(0) if not deployed or not configured.
  }

  mapping(address token => TokenConfig) internal s_tokenConfig;
  EnumerableSet.AddressSet internal s_tokens;

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
      revert UnsupportedToken(token);
    }

    return pool;
  }

  function getAllConfiguredTokens() external view returns (address[] memory) {
    return s_tokens.values();
  }

  function getTokenConfig(address token) external view returns (TokenConfig memory) {
    return s_tokenConfig[token];
  }

  // ================================================================
  // │                  Administrator functions                     │
  // ================================================================

  /// @notice Sets the pool for a token. Setting the pool to address(0) effectively delists the token
  /// from CCIP. Setting the pool to any other address enables the token on CCIP.
  /// @param token The token to set the pool for.
  /// @param pool The pool to set for the token.
  function setPool(address token, address pool) external onlyTokenAdmin(token) {
    TokenConfig storage config = s_tokenConfig[token];

    address previousPool = config.tokenPool;
    config.tokenPool = pool;

    if (previousPool != pool) {
      emit PoolSet(token, previousPool, pool);
    }
  }

  /// @notice Transfers the administrator role for a token to a new address with a 2-step process.
  /// @param token The token to transfer the administrator role for.
  /// @param newOwner The address to transfer the administrator role to.
  /// @dev The new owner must call `acceptAdminRole` to accept the role.
  function transferAdminRole(address token, address newOwner) external onlyTokenAdmin(token) {
    TokenConfig storage config = s_tokenConfig[token];
    config.pendingAdministrator = newOwner;

    emit AdministratorTransferRequested(token, msg.sender, newOwner);
  }

  /// @notice Accepts the administrator role for a token.
  /// @param token The token to accept the administrator role for.
  /// @dev This function can only be called by the pending administrator.
  function acceptAdminRole(address token) external {
    TokenConfig storage config = s_tokenConfig[token];
    if (config.pendingAdministrator != msg.sender) {
      revert OnlyPendingAdministrator(msg.sender, token);
    }

    config.administrator = msg.sender;
    config.pendingAdministrator = address(0);

    emit AdministratorTransferred(token, msg.sender);
  }

  // ================================================================
  // │                    Administrator config                      │
  // ================================================================

  /// @notice Public getter to check for permissions of an administrator
  function isAdministrator(address localToken, address administrator) public view returns (bool) {
    return s_tokenConfig[localToken].administrator == administrator;
  }

  /// @notice Resisters a new local administrator for a token.
  /// @param localToken The token to register the administrator for.
  /// @param administrator The address of the new administrator.
  /// @dev Can only be called by a registry module.
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
    config.isPermissionedAdmin = false;

    s_tokens.add(localToken);

    emit AdministratorRegistered(localToken, administrator);
  }

  /// @notice Registers a local administrator for a token. This will overwrite any potential current administrator
  /// and set the permissionedAdmin to true.
  function registerAdministratorPermissioned(address localToken, address administrator) external onlyOwner {
    TokenConfig storage config = s_tokenConfig[localToken];

    config.administrator = administrator;
    config.isRegistered = true;
    config.isPermissionedAdmin = true;

    s_tokens.add(localToken);

    emit AdministratorRegistered(localToken, administrator);
  }

  // ================================================================
  // │                      Registry Modules                        │
  // ================================================================

  function addRegistryModule(address module) external onlyOwner {
    s_RegistryModules.add(module);
  }

  modifier onlyTokenAdmin(address token) {
    if (s_tokenConfig[token].administrator != msg.sender) {
      revert OnlyAdministrator(msg.sender, token);
    }
    _;
  }
}
