// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IAny2EVMMultiOffRamp} from "./IAny2EVMMultiOffRamp.sol";
import {IRMNBase} from "./IRMNBase.sol";

/// @notice This interface contains the only RMN-related functions that might be used on-chain by other CCIP contracts.
interface IRMNOffChainBlessor is IRMNBase {
  /// @notice A Merkle root tagged with the address of the commit store contract it is destined for.
  struct TaggedRoot {
    uint64 sourceChainSelector;
    address commitStore;
    bytes32 root;
  }

  function isBlessed(IAny2EVMMultiOffRamp.MerkleRoot calldata root) external view returns (bool);
}
