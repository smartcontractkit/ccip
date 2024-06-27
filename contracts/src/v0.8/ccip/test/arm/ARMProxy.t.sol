// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMN} from "src/IRMN.sol";

import {ARMProxy} from "src/ARMProxy.sol";
import {RMN} from "src/RMN.sol";

import {MockRMN} from "test/MockRMN.t.sol";
import {RMNSetup, makeSubjects} from "test/RMNSetup.t.sol";

contract ARMProxyTest is RMNSetup {
  MockRMN internal s_mockRMN;
  ARMProxy internal s_armProxy;

  function setUp() public virtual override {
    RMNSetup.setUp();
    s_mockRMN = new MockRMN();
    s_armProxy = new ARMProxy(address(s_rmn));
  }

  function test_ARMIsCursed_Success() public {
    s_armProxy.setARM(address(s_mockRMN));
    assertFalse(IRMN(address(s_armProxy)).isCursed());
    RMN(address(s_armProxy)).voteToCurse(makeCurseId(0), makeSubjects(0));
    assertTrue(IRMN(address(s_armProxy)).isCursed());
  }

  // FIXME: misleading test that will pass
  function test_ARMIsBlessed_Success() public {
    s_armProxy.setARM(address(s_mockRMN));
    assertTrue(IRMN(address(s_armProxy)).isBlessed(IRMN.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
    RMN(address(s_armProxy)).voteToCurse(makeCurseId(0), makeSubjects(0));
    // depends on the implementation of MockRMN which is inconsistent with ARM
    // in the ARM contract, a vote to curse will not cause all isBlessed calls to return false
    assertFalse(IRMN(address(s_armProxy)).isBlessed(IRMN.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
  }

  function test_ARMCallRevertReasonForwarded() public {
    bytes memory err = bytes("revert");
    s_mockRMN.setRevert(err);
    s_armProxy.setARM(address(s_mockRMN));
    vm.expectRevert(abi.encodeWithSelector(MockRMN.CustomError.selector, err));
    IRMN(address(s_armProxy)).isCursed();
  }
}
