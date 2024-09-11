// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMNV2 {
  /// @notice signature components from RMN nodes
  struct Signature {
    bytes32 r;
    bytes32 s;
  }

  /// @notice Verifies signatures of RMN nodes, on dest lane updates as provided in the CommitReport
  /// @param offrampAddress is not inferred by msg.sender, in case the call is made through ARMProxy
  /// @param destLaneUpdates must be well formed, and is a representation of the CommitReport received from the oracles
  /// @param signatures must be sorted in ascending order by signer address
  /// @dev Will revert if verification fails
  function verify(
    address offRampAddress,
    Internal.MerkleRoot[] memory destLaneUpdates,
    Signature[] memory signatures
  ) external view;

  /// @notice If there is an active global or legacy curse, this function returns true.
  function isCursed() external view returns (bool);

  /// @notice If there is an active global curse, or an active curse for `subject`, this function returns true.
  /// @param subject To check whether a particular chain is cursed, set to bytes16(uint128(chainSelector)).
  function isCursed(bytes16 subject) external view returns (bool);
}
