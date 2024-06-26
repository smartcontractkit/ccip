// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMN} from "../../interfaces/IRMN.sol";

import {RMN} from "../../RMN.sol";
import {RMNProxy} from "../../RMNProxy.sol";
import {MockRMN} from "../mocks/MockRMN.sol";
import {RMNSetup} from "./RMNSetup.t.sol";

contract RMNProxyTest is RMNSetup {
  event RMNSet(address rmn);

  RMNProxy internal s_rmnProxy;

  function setUp() public virtual override {
    RMNSetup.setUp();
    s_rmnProxy = new RMNProxy(address(s_rmn));
  }

  function test_RMNIsCursed_Success() public {
    s_rmnProxy.setRMN(address(s_mockRMN));
    assertFalse(IRMN(address(s_rmnProxy)).isCursed());
    RMN(address(s_rmnProxy)).voteToCurse(bytes32(0));
    assertTrue(IRMN(address(s_rmnProxy)).isCursed());
  }

  function test_RMNIsBlessed_Success() public {
    s_rmnProxy.setRMN(address(s_mockRMN));
    assertTrue(IRMN(address(s_rmnProxy)).isBlessed(IRMN.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
    RMN(address(s_rmnProxy)).voteToCurse(bytes32(0));
    assertFalse(IRMN(address(s_rmnProxy)).isBlessed(IRMN.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
  }

  function test_RMNCallRevertReasonForwarded() public {
    bytes memory err = bytes("revert");
    s_mockRMN.setRevert(err);
    s_rmnProxy.setRMN(address(s_mockRMN));
    vm.expectRevert(abi.encodeWithSelector(MockRMN.CustomError.selector, err));
    IRMN(address(s_rmnProxy)).isCursed();
  }
}
