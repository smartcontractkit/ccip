// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IAny2EVMMultiOffRamp {
  /// @notice a sequenceNumber interval
  /// @dev RMN depends on this struct, if changing, please notify the RMN maintainers.
  struct Interval {
    uint64 min; // ───╮ Minimum sequence number, inclusive
    uint64 max; // ───╯ Maximum sequence number, inclusive
  }

  /// @dev Struct to hold a merkle root and an interval for a source chain so that an array of these can be passed in the CommitReport.
  struct MerkleRoot {
    uint64 sourceChainSelector; // Remote source chain selector that the Merkle Root is scoped to
    Interval interval; // Report interval of the merkle root
    bytes32 merkleRoot; // Merkle root covering the interval & source chain messages
    bytes[] rmnSignatures; // RMN signatures for root blessing
  }
}
