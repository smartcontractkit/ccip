// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import "../../offRamp/BaseOffRampRouter.sol";
import "../mocks/MockOffRamp.sol";

contract BaseOffRampRouterSetup is TokenSetup {
  BaseOffRampRouter s_router;
  BaseOffRampInterface[] s_offRamps;

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_offRamps = new BaseOffRampInterface[](2);
    s_offRamps[0] = BaseOffRampInterface(address(10));
    s_offRamps[1] = BaseOffRampInterface(address(11));
    s_router = new BaseOffRampRouter(s_offRamps);
  }
}

/// @notice #addOffRamp
contract BaseOffRampRouter_addOffRamp is BaseOffRampRouterSetup {
  BaseOffRampInterface internal s_newOffRamp;

  event OffRampAdded(BaseOffRampInterface indexed offRamp);

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
    BaseOffRampInterface existingOffRamp = s_offRamps[0];
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampRouterInterface.AlreadyConfigured.selector, existingOffRamp));
    s_router.addOffRamp(existingOffRamp);
  }

  function testZeroAddressReverts() public {
    vm.expectRevert(BaseOffRampRouterInterface.InvalidAddress.selector);
    s_router.addOffRamp(BaseOffRampInterface(address(0)));
  }
}

/// @notice #removeOffRamp
contract BaseOffRampRouter_removeOffRamp is BaseOffRampRouterSetup {
  event OffRampRemoved(BaseOffRampInterface indexed offRamp);

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

    vm.expectRevert(BaseOffRampRouterInterface.NoOffRampsConfigured.selector);
    s_router.removeOffRamp(s_offRamps[0]);
  }

  function testOffRampNotAllowedReverts() public {
    BaseOffRampInterface newRamp = new MockOffRamp();
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampRouterInterface.OffRampNotAllowed.selector, newRamp));
    s_router.removeOffRamp(newRamp);
  }
}

/// @notice #getOffRamps
contract BaseOffRampRouter_getOffRamps is BaseOffRampRouterSetup {
  // Success
  function testSuccess() public {
    BaseOffRampInterface[] memory offRamps = s_router.getOffRamps();
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
    assertFalse(s_router.isOffRamp(BaseOffRampInterface(address(1))));
  }
}
