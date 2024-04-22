// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";
import {IGetCCIPAdmin} from "./IGetCCIPAdmin.sol";
import {IOwner} from "./IOwner.sol";

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {TokenAdminRegistry} from "../TokenAdminRegistry.sol";

contract AdminRegistryModuleOwnerCustom is ITypeAndVersion, OwnerIsCreator {
  error CanOnlySelfRegister(address admin, address token);

  event AdministratorRegistered(address indexed token, address indexed administrator);

  string public constant override typeAndVersion = "AdminRegistryModuleOwnerCustom 1.5.0-dev";

  // The TokenAdminRegistry contract
  TokenAdminRegistry internal s_tokenAdminRegistry;

  constructor(address tokenAdminRegistry) {
    s_tokenAdminRegistry = TokenAdminRegistry(tokenAdminRegistry);
  }

  /// @notice Registers the admin of the token using the `getCCIPAdmin` method.
  /// @param token The token to register the admin for.
  /// @dev The caller must be the admin returned by the `getCCIPAdmin` method.
  function registerAdminGetCCIPAdmin(address token) external {
    _registerAdmin(token, IGetCCIPAdmin(token).getCCIPAdmin());
  }

  /// @notice Registers the admin of the token using the `owner` method.
  /// @param token The token to register the admin for.
  /// @dev The caller must be the admin returned by the `owner` method.
  function registerAdminOwner(address token) external {
    _registerAdmin(token, IOwner(token).owner());
  }

  /// @notice Registers the admin of the token to msg.sender given that the
  /// admin is equal to msg.sender.
  /// @param token The token to register the admin for.
  /// @param admin The caller must be the admin.
  function _registerAdmin(address token, address admin) internal {
    if (admin != msg.sender) {
      revert CanOnlySelfRegister(admin, token);
    }

    s_tokenAdminRegistry.registerAdministrator(token, admin);

    emit AdministratorRegistered(token, admin);
  }

  function getCCIPAdmin() external view returns (address) {
    return owner();
  }
}
