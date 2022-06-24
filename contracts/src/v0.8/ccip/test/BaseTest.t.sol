// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "forge-std/Test.sol";
import "./mocks/MockAFN.sol";

contract BaseTest is Test {
  uint256 constant s_sourceChainId = 1;
  uint256 constant s_destChainId = 2;
  address constant s_owner = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  uint256 constant s_blockTime = 1234567890;

  address constant s_onRampAddress = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;

  AFNInterface s_afn;

  function setUp() public virtual {
    // Stop any running pranks
    // This is needed when a contract inherits from multiple contracts that
    // each inherit BaseTest as only one prank can be running at the same
    // time.
    vm.stopPrank();
    // Set the sender to s_owner permanently
    vm.startPrank(s_owner);
    // Set the block time to a constant known value
    vm.warp(s_blockTime);

    s_afn = new MockAFN();
  }
}
