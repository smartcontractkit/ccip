// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @dev Struct to hold a merkle root and an interval for a source chain so that an array of these can be passed in the CommitReport.
/// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
struct MerkleRoot {
  uint64 sourceChainSelector; // ──╮ Remote source chain selector that the Merkle Root is scoped to
  uint64 minSeqNr; //              | Minimum sequence number, inclusive
  uint64 maxSeqNr; // ─────────────╯ Maximum sequence number, inclusive
  bytes32 merkleRoot; //             Merkle root covering the interval & source chain messages
  bytes onRampAddress; //            Generic onramp address, to support arbitrary sources; for EVM, use abi.encode
}

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMNRemote {
  /// @notice signature components from RMN nodes
  struct Signature {
    bytes32 r;
    bytes32 s;
  }

  function verify(MerkleRoot[] memory merkleRoots, Signature[] memory signatures) external view;

  /// @notice If there is an active global or legacy curse, this function returns true.
  function isCursed() external view returns (bool);

  /// @notice If there is an active global curse, or an active curse for `subject`, this function returns true.
  /// @param subject To check whether a particular chain is cursed, set to bytes16(uint128(chainSelector)).
  function isCursed(bytes16 subject) external view returns (bool);
}
