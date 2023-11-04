// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";

/// @notice A contract with helpers for bridge role 2-phase transfers.
contract ConfirmedBridgeRoleWithProposal is OwnerIsCreator {
  error PermissionsError();
  error CannotTransferToSelf();

  event BridgeRoleTransferRequested(address indexed from, address indexed to);
  event BridgeRoleTransferred(address indexed from, address indexed to);

  /// @dev Address of the bridge role that can set lockOrBurn and releaseOrMint callers.
  address private s_bridge;

  /// @dev Address being proposed to transfer the bridge role to. Can be address(0) if no ongoign proposal.
  address private s_pendingBridge;

  /// @notice Allows a bridge to propose transferring bridge role to a new address, pending.
  /// @param to The proposed new bridge address.
  function transferBridgeRole(address to) public onlyBridge {
    _transferBridgeRole(to);
  }

  /// @notice Allows a bridge role transfer to be completed by the recipient.
  function acceptBridgeRole() external {
    if (msg.sender != s_pendingBridge) revert PermissionsError();

    address oldBridge = s_bridge;
    s_bridge = msg.sender;
    s_pendingBridge = address(0);

    emit BridgeRoleTransferred(oldBridge, msg.sender);
  }

  /// @notice Get the current bridge role.
  /// @return Address of the current bridge.
  function bridge() public view returns (address) {
    return s_bridge;
  }

  /// @notice validate, transfer ownership, and emit relevant events.
  /// @param to The proposed new bridge address.
  function _transferBridgeRole(address to) private {
    if (to == msg.sender) revert CannotTransferToSelf();

    s_pendingBridge = to;

    emit BridgeRoleTransferRequested(s_bridge, to);
  }

  /// @notice Checks if the msg.sender is the specified bridge address.
  /// @dev Reverts with a PermissionsError if check fails.
  modifier onlyBridge() {
    if (msg.sender != bridge()) revert PermissionsError();
    _;
  }

  /// @notice Checks if the msg.sender is either the owner of the contract, or the specified bridge address.
  /// @dev Reverts with a PermissionsError if check fails.
  modifier onlyOwnerOrBridge() {
    if (msg.sender != owner() && msg.sender != bridge()) revert PermissionsError();
    _;
  }
}
