// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../mocks/MockOffRamp.sol";
import "../helpers/ramps/BaseOffRampRouterHelper.sol";
import "../TokenSetup.t.sol";

contract BaseOffRampRouterSetup is TokenSetup {
  BaseOffRampRouterHelper s_router;
  IBaseOffRamp[] s_offRamps;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_offRamps = new IBaseOffRamp[](2);
    s_offRamps[0] = IBaseOffRamp(address(10));
    s_offRamps[1] = IBaseOffRamp(address(11));
    s_router = new BaseOffRampRouterHelper(s_offRamps);
  }
}

// TODO _callWithExactGas

/// @notice #addOffRamp
contract BaseOffRampRouter_addOffRamp is BaseOffRampRouterSetup {
  IBaseOffRamp internal s_newOffRamp;

  event OffRampAdded(IBaseOffRamp indexed offRamp);

  function setUp() public virtual override {
    BaseOffRampRouterSetup.setUp();

    s_newOffRamp = new MockOffRamp();
  }

  // Success

  function testSuccess() public {
    assertFalse(s_router.isOffRamp(s_newOffRamp));
    uint256 lengthBefore = s_router.getOffRamps().length;

    vm.expectEmit(true, false, false, true);
    emit OffRampAdded(s_newOffRamp);
    s_router.addOffRamp(s_newOffRamp);

    assertTrue(s_router.isOffRamp(s_newOffRamp));
    assertEq(lengthBefore + 1, s_router.getOffRamps().length);
  }

  // Reverts

  function testOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_router.addOffRamp(s_newOffRamp);
  }

  function testAlreadyConfiguredReverts() public {
    IBaseOffRamp existingOffRamp = s_offRamps[0];
    vm.expectRevert(abi.encodeWithSelector(IAny2EVMOffRampRouter.AlreadyConfigured.selector, existingOffRamp));
    s_router.addOffRamp(existingOffRamp);
  }

  function testZeroAddressReverts() public {
    vm.expectRevert(IAny2EVMOffRampRouter.InvalidAddress.selector);
    s_router.addOffRamp(IBaseOffRamp(address(0)));
  }
}

/// @notice #removeOffRamp
contract BaseOffRampRouter_removeOffRamp is BaseOffRampRouterSetup {
  event OffRampRemoved(IBaseOffRamp indexed offRamp);

  // Success

  function testSuccess() public {
    uint256 lengthBefore = s_router.getOffRamps().length;

    vm.expectEmit(true, false, false, true);
    emit OffRampRemoved(s_offRamps[0]);
    s_router.removeOffRamp(s_offRamps[0]);

    assertFalse(s_router.isOffRamp(s_offRamps[0]));
    assertEq(lengthBefore - 1, s_router.getOffRamps().length);
  }

  // Reverts

  function testOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_router.removeOffRamp(s_offRamps[0]);
  }

  function testNoOffRampsReverts() public {
    s_router.removeOffRamp(s_offRamps[0]);
    s_router.removeOffRamp(s_offRamps[1]);

    assertEq(0, s_router.getOffRamps().length);

    vm.expectRevert(IAny2EVMOffRampRouter.NoOffRampsConfigured.selector);
    s_router.removeOffRamp(s_offRamps[0]);
  }

  function testOffRampNotAllowedReverts() public {
    IBaseOffRamp newRamp = new MockOffRamp();
    vm.expectRevert(abi.encodeWithSelector(IAny2EVMOffRampRouter.OffRampNotAllowed.selector, newRamp));
    s_router.removeOffRamp(newRamp);
  }
}

/// @notice #getOffRamps
contract BaseOffRampRouter_getOffRamps is BaseOffRampRouterSetup {
  // Success
  function testSuccess() public {
    IBaseOffRamp[] memory offRamps = s_router.getOffRamps();
    assertEq(2, offRamps.length);
    assertEq(address(s_offRamps[0]), address(offRamps[0]));
    assertEq(address(s_offRamps[1]), address(offRamps[1]));
  }
}

/// @notice #isOffRamp
contract BaseOffRampRouter_isOffRamp is BaseOffRampRouterSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_router.isOffRamp(s_offRamps[0]));
    assertTrue(s_router.isOffRamp(s_offRamps[1]));
    assertFalse(s_router.isOffRamp(IBaseOffRamp(address(1))));
  }
}
