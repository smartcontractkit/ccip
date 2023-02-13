// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IRouter} from "../../interfaces/router/IRouter.sol";

import "../TokenSetup.t.sol";
import {Router} from "../../router/Router.sol";
import "../fees/WETH9.sol";

contract RouterSetup is BaseTest {
  IRouter internal s_sourceRouter;
  IRouter internal s_destRouter;

  function setUp() public virtual override {
    BaseTest.setUp();

    if (address(s_sourceRouter) == address(0)) {
      address[] memory offRamps = new address[](0);
      WETH9 weth = new WETH9();
      s_sourceRouter = new Router(offRamps, address(weth));
    }
    if (address(s_destRouter) == address(0)) {
      address[] memory offRamps = new address[](0);
      WETH9 weth = new WETH9();
      s_destRouter = new Router(offRamps, address(weth));
    }
  }
}
