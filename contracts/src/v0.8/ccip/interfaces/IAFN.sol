// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice This interface contains *only* the functions that might be used on-chain by other CCIP contracts.
interface IAFN {
  function badSignalReceived() external view returns (bool);

  function isBlessed(bytes32 taggedRootHash) external view returns (bool);
}
