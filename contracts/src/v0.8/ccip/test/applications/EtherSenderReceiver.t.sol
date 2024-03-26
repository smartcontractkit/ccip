// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Test, stdError} from "forge-std/Test.sol";

import {EtherSenderReceiver} from "../../applications/EtherSenderReceiver.sol";
import {Router} from "../../Router.sol";
import {WETH9} from "../WETH9.sol";
import {MockARM} from "../mocks/MockARM.sol";

contract EtherSenderReceiverTest is Test {
  EtherSenderReceiver internal s_etherSenderReceiver;
  Router internal s_router;
  WETH9 internal s_weth;

  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;

  function setUp() public {
    vm.startPrank(OWNER);

    s_weth = new WETH9();
    MockARM mockARM = new MockARM();
    s_router = new Router(address(s_weth), address(mockARM));
    s_etherSenderReceiver = new EtherSenderReceiver(address(this));
  }
}

// TODO: write tests.
