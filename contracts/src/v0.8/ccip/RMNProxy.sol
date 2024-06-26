// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "./../shared/access/OwnerIsCreator.sol";

/// @notice The RMNProxy serves to allow CCIP contracts
/// to point to a static address for RMN queries, which saves gas
/// since each contract need not store an RMN address in storage. That way
/// we can add RMN queries along many code paths for increased defense in depth
/// with minimal additional cost.
contract RMNProxy is OwnerIsCreator, ITypeAndVersion {
  error ZeroAddressNotAllowed();

  event RMNSet(address rmn);

  // STATIC CONFIG
  string public constant override typeAndVersion = "RMNProxy 1.5.0-dev";

  // DYNAMIC CONFIG
  address private s_rmn;

  constructor(address rmn) {
    setRMN(rmn);
  }

  /// @notice SetRMN sets the RMN implementation contract address.
  /// @param rmn The address of the rmn implementation contract.
  function setRMN(address rmn) public onlyOwner {
    if (rmn == address(0)) revert ZeroAddressNotAllowed();
    s_rmn = rmn;
    emit RMNSet(rmn);
  }

  /// @notice getRMN gets the RMN implementation contract address.
  /// @return rmn The address of the rmn implementation contract.
  function getRMN() external view returns (address) {
    return s_rmn;
  }

  // We use a fallback function instead of explicit implementations of the functions
  // defined in IRMN.sol to preserve compatibility with future additions to the IRMN
  // interface. Calling IRMN interface methods in RMNProxy should be transparent, i.e.
  // their input/output behaviour should be identical to calling the proxied s_rmn
  // contract directly. (If s_rmn doesn't point to a contract, we always revert.)
  // solhint-disable-next-line payable-fallback, no-complex-fallback
  fallback() external {
    address rmn = s_rmn;
    // solhint-disable-next-line no-inline-assembly
    assembly {
      // Revert if no contract present at destination address, otherwise call
      // might succeed unintentionally.
      if iszero(extcodesize(rmn)) { revert(0, 0) }
      // We use memory starting at zero, overwriting anything that might already
      // be stored there. This messes with Solidity's expectations around memory
      // layout, but it's fine because we always exit execution of this contract
      // inside this assembly block, i.e. we don't cede control to code generated
      // by the Solidity compiler that might have expectations around memory
      // layout.
      // Copy calldatasize() bytes from calldata offset 0 to memory offset 0.
      calldatacopy(0, 0, calldatasize())
      // Call the underlying RMN implementation. out and outsize are 0 because
      // we don't know the size yet. We hardcode value to zero.
      let success := call(gas(), rmn, 0, 0, calldatasize(), 0, 0)
      // Copy the returned data.
      returndatacopy(0, 0, returndatasize())
      // Pass through successful return or revert and associated data.
      if success { return(0, returndatasize()) }
      revert(0, returndatasize())
    }
  }
}
