// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {AuthorizedCallers} from "../../access/AuthorizedCallers.sol";
import {BaseTest} from "../BaseTest.t.sol";

contract AuthorizedCallers_setUp is BaseTest {
  address[] s_callers;

  AuthorizedCallers s_authorizedCallers;

  function setUp() public override {
    super.setUp();
    s_callers.push(vm.addr(1));
    s_callers.push(vm.addr(2));

    s_authorizedCallers = new AuthorizedCallers(s_callers);
  }

  function _assertAuthorizedCallersEqual(address[] memory a, address[] memory b) internal {
    assertEq(a.length, b.length);

    for (uint256 i = 0; i < a.length; ++i) {
      assertEq(a[i], b[i]);
    }
  }
}

contract AuthorizedCallers_constructor is AuthorizedCallers_setUp {
  event AuthorizedCallerAdded(address caller);

  function test_constructor_Success() public {
    for (uint256 i = 0; i < s_callers.length; ++i) {
      vm.expectEmit();
      emit AuthorizedCallerAdded(s_callers[i]);
    }

    s_authorizedCallers = new AuthorizedCallers(s_callers);

    _assertAuthorizedCallersEqual(s_callers, s_authorizedCallers.getAllAuthorizedCallers());
  }

  function test_ZeroAddressNotAllowed_Revert() public {
    s_callers[0] = address(0);

    vm.expectRevert(AuthorizedCallers.ZeroAddressNotAllowed.selector);

    new AuthorizedCallers(s_callers);
  }
}

contract AuthorizedCallers_applyAuthorizedCallerUpdates is AuthorizedCallers_setUp {
  event AuthorizedCallerAdded(address caller);
  event AuthorizedCallerRemoved(address caller);

  function test_applyAuthorizedCallerUpdates_Success() public {
    address[] memory addedCallers = new address[](1);
    addedCallers[0] = vm.addr(3);

    for (uint256 i = 0; i < s_callers.length; ++i) {
      vm.expectEmit();
      emit AuthorizedCallerRemoved(s_callers[i]);
    }

    vm.expectEmit();
    emit AuthorizedCallerAdded(vm.addr(3));

    AuthorizedCallers.AuthorizedCallerArgs memory authorizedCallerArgs =
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: addedCallers, removedCallers: s_callers});

    s_authorizedCallers.applyAuthorizedCallerUpdates(authorizedCallerArgs);

    _assertAuthorizedCallersEqual(addedCallers, s_authorizedCallers.getAllAuthorizedCallers());
  }

  function test_OnlyCallableByOwner_Revert() public {
    vm.stopPrank();

    AuthorizedCallers.AuthorizedCallerArgs memory authorizedCallerArgs =
      AuthorizedCallers.AuthorizedCallerArgs({addedCallers: new address[](0), removedCallers: new address[](0)});

    vm.expectRevert("Only callable by owner");

    s_authorizedCallers.applyAuthorizedCallerUpdates(authorizedCallerArgs);
  }

  function test_ZeroAddressNotAllowed_Revert() public {
    s_callers[0] = address(0);

    vm.expectRevert(AuthorizedCallers.ZeroAddressNotAllowed.selector);

    new AuthorizedCallers(s_callers);
  }
}
