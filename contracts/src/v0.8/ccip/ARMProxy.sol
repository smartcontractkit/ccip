// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {IARM} from "./interfaces/IARM.sol";

import {OwnerIsCreator} from "./../shared/access/OwnerIsCreator.sol";

contract ARMProxy is OwnerIsCreator, TypeAndVersionInterface {
  error ZeroAddressNotAllowed();

  event ARMSet(address arm);

  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ARMProxy 1.0.0";

  // DYNAMIC CONFIG
  address private s_arm;

  constructor(address arm) {
    setARM(arm);
  }

  /// @notice SetARM sets the ARM implementation contract address.
  /// @param arm The address of the arm implementation contract.
  function setARM(address arm) public onlyOwner {
    if (arm == address(0)) revert ZeroAddressNotAllowed();
    s_arm = arm;
    emit ARMSet(arm);
  }

  /// @notice getARM gets the ARM implementation contract address.
  /// @return arm The address of the arm implementation contract.
  function getARM() external view returns (address) {
    return s_arm;
  }

  fallback() external {
    address arm = s_arm;
    assembly {
      // Revert if no contract present at destination address,
      // otherwise call would succeed unintentionally.
      if iszero(extcodesize(arm)) {
        revert(0, 0)
      }
      // This messes with solidity's expectations around memory layout, but it's fine
      // because we always exit execution of this contract inside this assembly block.
      // We overwrite the Solidity scratch pad at memory position 0.
      // calldatacopy(destOffset, offset, size)
      calldatacopy(0, 0, calldatasize())
      // Call the implementation.
      // out and outsize are 0 because we don't know the size yet.
      // We hardcode value to zero.
      let success := call(gas(), arm, 0, 0, calldatasize(), 0, 0)
      // Copy the returned data.
      returndatacopy(0, 0, returndatasize())
      if success {
        return(0, returndatasize())
      }
      // Throw a revert on success=false like solidity does.
      revert(0, returndatasize())
    }
  }
}
