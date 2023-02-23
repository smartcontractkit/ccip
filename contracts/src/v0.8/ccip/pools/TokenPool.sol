// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IPool} from "../interfaces/pools/IPool.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";

/// @notice Base abstract class with common functions for all token pools
abstract contract TokenPool is IPool, OwnerIsCreator, Pausable {
  IERC20 internal immutable i_token;
  mapping(address => bool) internal s_onRamps;
  mapping(address => bool) internal s_offRamps;

  constructor(IERC20 token) {
    if (address(token) == address(0)) revert NullAddressNotAllowed();

    i_token = token;
  }

  /// @notice Pause the pool
  /// @dev Only callable by the owner
  function pause() external override onlyOwner {
    _pause();
  }

  /// @notice Unpause the pool
  /// @dev Only callable by the owner
  function unpause() external override onlyOwner {
    _unpause();
  }

  /// @notice Set an onRamp's permissions
  /// @dev Only callable by the owner
  /// @param onRamp The onRamp
  /// @param permission Whether or not the onRamp has onRamp permissions on this contract
  function setOnRamp(address onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  /// @notice Set an offRamp's permissions
  /// @dev Only callable by the owner
  /// @param offRamp The offRamp
  /// @param permission Whether or not the offRamp has offRamp permissions on this contract
  function setOffRamp(address offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  /// @notice Checks whether something is a permissioned onRamp on this contract
  /// @return boolean
  function isOnRamp(address onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  /// @notice Checks whether something is a permissioned offRamp on this contract
  /// @return boolean
  function isOffRamp(address offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  /// @notice Gets the underlying token
  /// @return token
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /// @notice Checks whether the msg.sender is either the owner, or a permissioned onRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier validateOwnerOrOnRamp() {
    if (msg.sender != owner() && !isOnRamp(address(msg.sender))) revert PermissionsError();
    _;
  }

  /// @notice Checks whether the msg.sender is either the owner, or a permissioned offRamp on this contract
  /// @dev Reverts with a PermissionsError if check fails
  modifier validateOwnerOrOffRamp() {
    if (msg.sender != owner() && !isOffRamp(address(msg.sender))) revert PermissionsError();
    _;
  }
}
