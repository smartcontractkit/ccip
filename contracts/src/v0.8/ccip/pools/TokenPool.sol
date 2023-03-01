// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IPool} from "../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";

/// @notice Base abstract class with common functions for all token pools.
abstract contract TokenPool is IPool, OwnerIsCreator, Pausable {
  IERC20 internal immutable i_token;
  mapping(address => bool) internal s_onRamps;
  mapping(address => bool) internal s_offRamps;

  constructor(IERC20 token) {
    if (address(token) == address(0)) revert NullAddressNotAllowed();

    i_token = token;
  }

  /// @inheritdoc IPool
  function pause() external override onlyOwner {
    _pause();
  }

  /// @inheritdoc IPool
  function unpause() external override onlyOwner {
    _unpause();
  }

  /// @notice Set an onRamp's permissions.
  /// @dev Only callable by the owner.
  /// @param onRamp The onRamp contract address.
  /// @param permission Whether or not the onRamp has onRamp permissions on this contract.
  function setOnRamp(address onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  /// @notice Set an offRamp's permissions.
  /// @dev Only callable by the owner.
  /// @param offRamp The offRamp contract address.
  /// @param permission Whether or not the offRamp has offRamp permissions on this contract.
  function setOffRamp(address offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  /// @notice Checks whether something is a permissioned onRamp on this contract.
  /// @return true if the given address is a permissioned onRamp.
  function isOnRamp(address onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  /// @notice Checks whether something is a permissioned offRamp on this contract.
  /// @return true is the given address is a permissioned offRamp.
  function isOffRamp(address offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  /// @inheritdoc IPool
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /// @notice Checks whether the msg.sender is a permissioned onRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOnRamp() {
    if (!isOnRamp(msg.sender)) revert PermissionsError();
    _;
  }

  /// @notice Checks whether the msg.sender is a permissioned offRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier onlyOffRamp() {
    if (!isOffRamp(msg.sender)) revert PermissionsError();
    _;
  }
}
