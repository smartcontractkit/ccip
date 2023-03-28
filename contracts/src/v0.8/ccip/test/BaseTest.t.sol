// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "forge-std/Test.sol";
import "./mocks/MockAFN.sol";
import "./StructFactory.sol";

contract BaseTest is Test, StructFactory {
  MockAFN internal s_mockAFN;

  function setUp() public virtual {
    // Set the sender to OWNER permanently
    changePrank(OWNER);
    deal(OWNER, 1e20);

    // Set the block time to a constant known value
    vm.warp(BLOCK_TIME);

    s_mockAFN = new MockAFN();
  }
}
