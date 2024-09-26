// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

interface IRegistryModuleOwnerCustom {
  /// @notice Registers the admin of the token using the `getCCIPAdmin` method.
  /// @param token The token to register the admin for.
  /// @dev The caller must be the admin returned by the `getCCIPAdmin` method.
  function registerAdminViaGetCCIPAdmin(address token) external;

  /// @notice Registers the admin of the token using the `owner` method.
  /// @param token The token to register the admin for.
  /// @dev The caller must be the admin returned by the `owner` method.
  function registerAdminViaOwner(address token) external;
}
