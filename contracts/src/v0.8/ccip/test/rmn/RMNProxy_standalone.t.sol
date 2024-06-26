// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMNProxy} from "../../RMNProxy.sol";
import {Test} from "forge-std/Test.sol";

contract RMNProxyStandaloneTest is Test {
  event RMNSet(address rmn);

  address internal constant EMPTY_ADDRESS = address(0x1);
  address internal constant OWNER_ADDRESS = 0xC0ffeeEeC0fFeeeEc0ffeEeEc0ffEEEEC0FfEEee;
  address internal constant MOCK_RMN_ADDRESS = 0x1337133713371337133713371337133713371337;

  RMNProxy internal s_rmnProxy;

  function setUp() public virtual {
    // needed so that the extcodesize check in RMNProxy.fallback doesn't revert
    vm.etch(MOCK_RMN_ADDRESS, bytes("fake bytecode"));

    vm.prank(OWNER_ADDRESS);
    s_rmnProxy = new RMNProxy(MOCK_RMN_ADDRESS);
  }

  function test_Constructor() public {
    vm.expectEmit();
    emit RMNSet(MOCK_RMN_ADDRESS);
    RMNProxy proxy = new RMNProxy(MOCK_RMN_ADDRESS);
    assertEq(proxy.getRMN(), MOCK_RMN_ADDRESS);
  }

  function test_SetRMN() public {
    vm.expectEmit();
    emit RMNSet(MOCK_RMN_ADDRESS);
    vm.prank(OWNER_ADDRESS);
    s_rmnProxy.setRMN(MOCK_RMN_ADDRESS);
    assertEq(s_rmnProxy.getRMN(), MOCK_RMN_ADDRESS);
  }

  function test_SetRMNzero() public {
    vm.expectRevert(abi.encodeWithSelector(RMNProxy.ZeroAddressNotAllowed.selector));
    vm.prank(OWNER_ADDRESS);
    s_rmnProxy.setRMN(address(0x0));
  }

  function test_RMNCallEmptyContractRevert() public {
    vm.prank(OWNER_ADDRESS);
    s_rmnProxy.setRMN(EMPTY_ADDRESS); // No code at address 1, should revert.
    vm.expectRevert();
    bytes memory b = new bytes(0);
    (bool success,) = address(s_rmnProxy).call(b);
    success;
  }
}
