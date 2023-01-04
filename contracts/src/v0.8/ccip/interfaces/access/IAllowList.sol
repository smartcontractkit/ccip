// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IAllowList {
  error SenderNotAllowed(address sender);

  event AllowListSet(address[] allowlist);
  event AllowListEnabledSet(bool enabled);

  /**
   * @notice Enables or disabled the allowList functionality.
   * @param enabled Signals whether the allowlist should be enabled.
   */
  function setAllowlistEnabled(bool enabled) external;

  /**
   * @notice Gets whether the allowList functionality is enabled.
   * @return true is enabled, false if not.
   */
  function getAllowlistEnabled() external view returns (bool);

  /**
   * @notice Sets the allowed addresses.
   * @param allowlist The new allowed addresses.
   */
  function setAllowlist(address[] calldata allowlist) external;

  /**
   * @notice Gets the allowed addresses.
   * @return The allowed addresses.
   */
  function getAllowlist() external view returns (address[] memory);
}
