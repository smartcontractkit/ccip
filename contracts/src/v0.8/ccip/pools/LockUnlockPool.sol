// SPDX-License-Identifier: MIT

pragma solidity ^0.8.6;

import "../../vendor/SafeERC20.sol";
import "../../vendor/Pausable.sol";
import "../interfaces/OnRampInterface.sol";
import "../interfaces/OffRampInterface.sol";
import "../interfaces/PoolInterface.sol";
import "../access/OwnerIsCreator.sol";

contract LockUnlockPool is PoolInterface, OwnerIsCreator, Pausable {
  using SafeERC20 for IERC20;

  IERC20 private immutable s_token;
  mapping(OnRampInterface => bool) private s_onRamps;
  mapping(OffRampInterface => bool) private s_offRamps;

  error PermissionsError();

  constructor(IERC20 token) {
    s_token = token;
  }

  function lockOrBurn(address depositor, uint256 amount) external override whenNotPaused onlyOwnerOrOnRamp {
    getToken().safeTransferFrom(depositor, address(this), amount);
    emit Locked(msg.sender, depositor, amount);
  }

  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused onlyOwnerOrOffRamp {
    getToken().safeTransfer(recipient, amount);
    emit Released(msg.sender, recipient, amount);
  }

  function pause() external override onlyOwner {
    _pause();
  }

  function unpause() external override onlyOwner {
    _unpause();
  }

  function setOnRamp(OnRampInterface onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  function setOffRamp(OffRampInterface offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  function isOnRamp(OnRampInterface onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  function isOffRamp(OffRampInterface offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  function getToken() public view override returns (IERC20 token) {
    return s_token;
  }

  function _validateOwnerOrOnRamp() internal view {
    if (msg.sender != owner() && !isOnRamp(OnRampInterface(msg.sender))) revert PermissionsError();
  }

  function _validateOwnerOrOffRamp() internal view {
    if (msg.sender != owner() && !isOffRamp(OffRampInterface(msg.sender))) revert PermissionsError();
  }

  modifier onlyOwnerOrOnRamp() {
    _validateOwnerOrOnRamp();
    _;
  }

  modifier onlyOwnerOrOffRamp() {
    _validateOwnerOrOffRamp();
    _;
  }
}
