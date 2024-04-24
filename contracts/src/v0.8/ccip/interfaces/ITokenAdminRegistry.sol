// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface ITokenAdminRegistry {
  /// @notice Returns the pool for the given token.
  function getPool(address token) external view returns (address);

  /// @notice Returns every token that has been configured through a permissoned method.
  function getPermissionedTokens() external view returns (address[] memory);
}
