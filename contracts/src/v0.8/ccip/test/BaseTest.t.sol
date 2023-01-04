// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "forge-std/Test.sol";
import "./mocks/MockAFN.sol";
import "./StructFactory.sol";

contract BaseTest is Test, StructFactory {
  IAFN internal s_afn;

  function setUp() public virtual {
    // Set the sender to OWNER permanently
    changePrank(OWNER);

    // Set the block time to a constant known value
    vm.warp(BLOCK_TIME);

    s_afn = new MockAFN();
  }
}
