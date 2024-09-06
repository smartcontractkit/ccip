// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {RMNRemote} from "../../rmn/RMNRemote.sol";
import {BaseTest} from "../BaseTest.t.sol";

contract RMNRemoteTest is BaseTest {
  RMNRemote public s_rmnRemote;

  function setUp() public virtual override {
    super.setUp();
    s_rmnRemote = new RMNRemote(1);
  }
}

contract RMNRemote_constructor is RMNRemoteTest {
  function test_constructor_success() public {
    assertEq(s_rmnRemote.getChainSelector(), 1);
  }
}
