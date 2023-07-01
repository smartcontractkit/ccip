// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import "./ARMSetup.t.sol";
import {ARMProxy} from "../../ARMProxy.sol";
import {IARM} from "../../interfaces/IARM.sol";

contract ARMProxyTest is ARMSetup {
  event ARMSet(address arm);

  ARMProxy s_armProxy;

  function setUp() public virtual override {
    ARMSetup.setUp();
    s_armProxy = new ARMProxy(address(s_arm));
  }

  function testConstructor() public {
    vm.expectEmit();
    emit ARMSet(address(s_mockARM));
    ARMProxy proxy = new ARMProxy(address(s_mockARM));
    assertEq(proxy.getARM(), address(s_mockARM));
  }

  function testSetARM() public {
    vm.expectEmit();
    emit ARMSet(address(s_mockARM));
    s_armProxy.setARM(address(s_mockARM));
    assertEq(s_armProxy.getARM(), address(s_mockARM));
  }

  function testSetARMzero() public {
    vm.expectRevert(abi.encodeWithSelector(ARMProxy.ZeroAddressNotAllowed.selector));
    s_armProxy.setARM(address(0x0));
  }

  function testARMCall_fuzz(bytes memory call, bytes memory ret) public {
    s_armProxy.setARM(address(s_mockARM));
    vm.mockCall(address(s_mockARM), 0, call, ret);
    (bool success, bytes memory result) = address(s_armProxy).call(call);
    assertEq(result, ret);
    assertTrue(success);
    vm.clearMockedCalls();
  }

  function testARMIsCursedSuccess() public {
    s_armProxy.setARM(address(s_mockARM));
    assertFalse(IARM(address(s_armProxy)).isCursed());
    ARM(address(s_armProxy)).voteToCurse(bytes32(0));
    assertTrue(IARM(address(s_armProxy)).isCursed());
  }

  function testARMIsBlessedSuccess() public {
    s_armProxy.setARM(address(s_mockARM));
    assertTrue(IARM(address(s_armProxy)).isBlessed(IARM.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
    ARM(address(s_armProxy)).voteToCurse(bytes32(0));
    assertFalse(IARM(address(s_armProxy)).isBlessed(IARM.TaggedRoot({commitStore: address(0), root: bytes32(0)})));
  }

  function testARMCallEmptyContractRevert() public {
    address emptyAddress = address(1);
    s_armProxy.setARM(emptyAddress); // No code at address 1 should revert.
    vm.expectRevert();
    bytes memory b = new bytes(0);
    address(s_armProxy).call(b);
  }

  function testARMCallRevertReasonForwarded() public {
    bytes memory err = bytes("revert");
    s_mockARM.setRevert(err);
    s_armProxy.setARM(address(s_mockARM));
    vm.expectRevert(abi.encodeWithSelector(MockARM.CustomError.selector, err));
    IARM(address(s_armProxy)).isCursed();
  }
}
