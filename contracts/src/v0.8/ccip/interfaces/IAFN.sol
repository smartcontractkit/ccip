// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice This interface contains the only AFN-related functions that might be used on-chain by other CCIP contracts.
interface IAFN {
  struct TaggedRoot {
    address commitStore;
    bytes32 root;
  }

  /// @notice Callers MUST NOT cache the return value as a blessed tagged root could become unblessed.
  function isBlessed(TaggedRoot calldata taggedRoot) external view returns (bool);

  function isCursed() external view returns (bool);
}
