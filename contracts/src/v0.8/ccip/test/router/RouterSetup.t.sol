// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import {Router} from "../../Router.sol";
import {WETH9} from "../WETH9.sol";

contract RouterSetup is BaseTest {
  Router internal s_sourceRouter;
  Router internal s_destRouter;

  function setUp() public virtual override {
    BaseTest.setUp();

    if (address(s_sourceRouter) == address(0)) {
      WETH9 weth = new WETH9();
      s_sourceRouter = new Router(address(weth));
    }
    if (address(s_destRouter) == address(0)) {
      WETH9 weth = new WETH9();
      s_destRouter = new Router(address(weth));
    }
  }

  function generateReceiverMessage(uint64 chainSelector) internal pure returns (Client.Any2EVMMessage memory) {
    Client.EVMTokenAmount[] memory ta = new Client.EVMTokenAmount[](0);
    return
      Client.Any2EVMMessage({
        messageId: bytes32("a"),
        sourceChainSelector: chainSelector,
        sender: bytes("a"),
        data: bytes("a"),
        destTokenAmounts: ta
      });
  }
}
