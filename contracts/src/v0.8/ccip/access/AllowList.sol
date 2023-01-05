// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAllowList} from "../interfaces/access/IAllowList.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

contract AllowList is IAllowList, OwnerIsCreator {
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

  /// @inheritdoc IAllowList
  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowListEnabledSet(enabled);
  }

  /// @inheritdoc IAllowList
  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  /// @inheritdoc IAllowList
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

  /// @inheritdoc IAllowList
  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }
}
