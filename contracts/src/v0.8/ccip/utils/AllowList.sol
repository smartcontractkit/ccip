// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "./interfaces/AllowListInterface.sol";
import "../access/OwnerIsCreator.sol";

contract AllowList is AllowListInterface, OwnerIsCreator {
  using Address for address;

  // Whether the allowlist is enabled
  bool internal s_allowlistEnabled;
  // List of allowed addresses
  address[] internal s_allowList;
  // Addresses that are allowed to send messages
  mapping(address => bool) internal s_allowed;

  constructor(address[] memory allowlist) {
    if (allowlist.length > 0) {
      s_allowlistEnabled = true;
      s_allowList = allowlist;
    }
    for (uint256 i = 0; i < allowlist.length; ++i) {
      s_allowed[allowlist[i]] = true;
    }
  }

  /// @inheritdoc AllowListInterface
  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowListEnabledSet(enabled);
  }

  /// @inheritdoc AllowListInterface
  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  /// @inheritdoc AllowListInterface
  function setAllowlist(address[] calldata allowlist) external onlyOwner {
    // Remove existing allowlist
    address[] memory existingList = s_allowList;
    for (uint256 i = 0; i < existingList.length; ++i) {
      s_allowed[existingList[i]] = false;
    }

    // Set the new allowlist
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; ++i) {
      s_allowed[allowlist[i]] = true;
    }
    emit AllowListSet(allowlist);
  }

  /// @inheritdoc AllowListInterface
  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }
}
