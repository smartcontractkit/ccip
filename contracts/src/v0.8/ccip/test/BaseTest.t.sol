// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "forge-std/Test.sol";
import "./mocks/MockAFN.sol";

contract BaseTest is Test {
  uint256 internal constant SOURCE_CHAIN_ID = 1;
  uint256 internal constant DEST_CHAIN_ID = 2;
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant STRANGER = 0x1111111111111111111111111111111111111111;
  uint256 internal constant BLOCK_TIME = 1234567890;

  address internal constant ON_RAMP_ADDRESS = 0x11118e64e1FB0c487f25dD6D3601FF6aF8d32E4e;

  AFNInterface internal s_afn;

  function setUp() public virtual {
    // Set the sender to OWNER permanently
    changePrank(OWNER);

    // Set the block time to a constant known value
    vm.warp(BLOCK_TIME);

    s_afn = new MockAFN();
  }
}
