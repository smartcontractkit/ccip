// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @dev Struct to hold a merkle root and an interval for a source chain so that an array of these can be passed in the CommitReport.
/// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
struct MerkleRoot {
  uint64 sourceChainSelector; // Remote source chain selector that the Merkle Root is scoped to
  bytes onRampAddress; // generic, to support arbitrary sources; for EVM2EVM, use abi.encode
  uint64 minSeqNr; // Minimum sequence number, inclusive
  uint64 maxSeqNr; // Maximum sequence number, inclusive
  bytes32 merkleRoot; // Merkle root covering the interval & source chain messages
}

struct Signature {
  bytes32 r;
  bytes32 s;
}

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMNRemote {
  function verify(MerkleRoot[] memory merkleRoots, Signature[] memory signatures) external view;
}
