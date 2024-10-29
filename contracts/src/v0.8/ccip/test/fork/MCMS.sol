// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract MCMS {
  struct Call {
    address target;
    uint256 value;
    bytes data;
  }

  function scheduleBatch(Call[] calldata calls, bytes32 predecessor, bytes32 salt, uint256 delay) external {}
}
