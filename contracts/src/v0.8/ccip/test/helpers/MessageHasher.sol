// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../../libraries/Internal.sol";

/// @notice MessageHasher is a contract that provides a function to hash an Any2EVMRampMessage.
contract MessageHasher {
  function hash(Internal.Any2EVMRampMessage memory message, bytes memory onRamp) public pure returns (bytes32) {
    return Internal._hash(message, onRamp);
  }
}
