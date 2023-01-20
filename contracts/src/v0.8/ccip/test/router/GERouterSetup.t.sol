// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import {GERouter} from "../../router/GERouter.sol";

contract GERouterSetup is BaseTest {
  GERouter internal s_sourceRouter;

  function setUp() public virtual override {
    BaseTest.setUp();

    address[] memory offRamps = new address[](0);
    s_sourceRouter = new GERouter(offRamps);
  }
}
