// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

contract TokenAdminRegistry is ITokenAdminRegistry, ITypeAndVersion, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;

  error OnlyRegistryModule(address sender);
  error OnlyAdministrator(address sender, address token);
  error OnlyPendingAdministrator(address sender, address token);
  error UnsupportedToken(address token);
  error AlreadyRegistered(address token);

  event AdministratorRegistered(address indexed token, address indexed administrator);
  event PoolSet(address indexed token, address indexed previousPool, address indexed newPool);
  event AdministratorTransferRequested(address indexed token, address indexed currentAdmin, address indexed newAdmin);
  event AdministratorTransferred(address indexed token, address indexed newAdmin);
  event DisableReRegistrationSet(address indexed token, bool disabled);
  event RegistryModuleAdded(address indexed module);
  event RegistryModuleRemoved(address indexed module);

  // The struct is packed in a way that optimizes the attributes that are accessed together.
  // solhint-disable-next-line gas-struct-packing
  struct TokenConfig {
    bool isPermissionedAdmin; // ─────────────╮ if true, this administrator has been configured by the CCIP owner
    //                                        │ and it could have elevated permissions.
    bool isRegistered; //                     │ if true, the token is registered in the registry
    bool disableReRegistration; //            │ if true, the token cannot be permissionlessly re-registered
    address administrator; // ────────────────╯ the current administrator of the token
    address pendingAdministrator; //            the address that is pending to become the new owner
    address tokenPool; // the token pool for this token. Can be address(0) if not deployed or not configured.
  }

  string public constant override typeAndVersion = "TokenAdminRegistry 1.5.0-dev";

  // Mapping of token address to token configuration
  mapping(address token => TokenConfig) internal s_tokenConfig;

  // All tokens that have been configured
  EnumerableSet.AddressSet internal s_tokens;

  // Registry modules are allowed to register administrators for tokens
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

  /// @notice Returns a list of tokens that are configured in the token admin registry.
  /// @param startIndex Starting index in list, can be 0 if you want to start from the beginning.
  /// @param maxCount Maximum number of tokens to retrieve. Since the list can be large,
  /// it is recommended to use a paging mechanism to retrieve all tokens. If querying for very
  /// large lists, RPCs can time out. If you want all tokens, use type(uint64).max.
  /// @return tokens List of configured tokens.
  /// @dev The function is paginated to avoid RPC timeouts.
  /// @dev The ordering is guaranteed to remain the same as it is not possible to remove tokens
  /// from s_tokens.
  function getAllConfiguredTokens(uint64 startIndex, uint64 maxCount) external view returns (address[] memory tokens) {
    uint256 numberOfTokens = s_tokens.length();
    if (startIndex >= numberOfTokens) {
      return tokens;
    }
    uint256 count = maxCount;
    if (count + startIndex > numberOfTokens) {
      count = numberOfTokens - startIndex;
    }
    tokens = new address[](count);
    for (uint256 i = 0; i < count; ++i) {
      tokens[i] = s_tokens.at(startIndex + i);
    }

    return tokens;
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

  /// @notice Disables the re-registration of a token.
  /// @param token The token to disable re-registration for.
  /// @param disabled True to disable re-registration, false to enable it.
  function setDisableReRegistration(address token, bool disabled) external onlyTokenAdmin(token) {
    s_tokenConfig[token].disableReRegistration = disabled;

    emit DisableReRegistrationSet(token, disabled);
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
    if (!isRegistryModule(msg.sender)) {
      revert OnlyRegistryModule(msg.sender);
    }
    TokenConfig storage config = s_tokenConfig[localToken];

    if (config.disableReRegistration && config.isRegistered) {
      revert AlreadyRegistered(localToken);
    }

    config.administrator = administrator;
    config.isRegistered = true;
    config.isPermissionedAdmin = false;

    s_tokens.add(localToken);

    emit AdministratorRegistered(localToken, administrator);
  }

  /// @notice Registers a local administrator for a token. This will overwrite any potential current administrator
  /// and set the permissionedAdmin to true.
  /// @param localToken The token to register the administrator for.
  /// @param administrator The address of the new administrator.
  /// @dev Can only be called by the owner.
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

  /// @notice Checks if an address is a registry module.
  /// @param module The address to check.
  /// @return True if the address is a registry module, false otherwise.
  function isRegistryModule(address module) public view returns (bool) {
    return s_RegistryModules.contains(module);
  }

  /// @notice Adds a new registry module to the list of allowed modules.
  /// @param module The module to add.
  function addRegistryModule(address module) external onlyOwner {
    if (s_RegistryModules.add(module)) {
      emit RegistryModuleAdded(module);
    }
  }

  /// @notice Removes a registry module from the list of allowed modules.
  /// @param module The module to remove.
  function removeRegistryModule(address module) external onlyOwner {
    if (s_RegistryModules.remove(module)) {
      emit RegistryModuleRemoved(module);
    }
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Checks if an address is the administrator of the given token.
  modifier onlyTokenAdmin(address token) {
    if (s_tokenConfig[token].administrator != msg.sender) {
      revert OnlyAdministrator(msg.sender, token);
    }
    _;
  }
}
