// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../mocks/MockOffRamp.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../BaseTest.t.sol";
import "../../../offRamp/toll/Any2EVMTollOffRampRouter.sol";
import "../../helpers/receivers/MaybeRevertMessageReceiver.sol";

contract EVM2EVMTollOffRampRouterSetup is BaseTest {
  event MessageReceived();

  Any2EVMTollOffRampRouter internal s_router;
  address[] internal s_offRamps;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_offRamps = new address[](2);
    s_offRamps[0] = address(10);
    s_offRamps[1] = address(11);
    s_router = new Any2EVMTollOffRampRouter(s_offRamps);
  }

  function _generateMockMessage() internal pure returns (Common.Any2EVMMessage memory) {
    return (
      Common.Any2EVMMessage({
        sourceChainId: SOURCE_CHAIN_ID,
        sender: abi.encode(STRANGER),
        data: abi.encode(0),
        destTokensAndAmounts: new Common.EVMTokenAndAmount[](0)
      })
    );
  }
}

/// @notice #constructor
contract EVM2EVMTollOffRampRouter_constructor is EVM2EVMTollOffRampRouterSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("Any2EVMTollOffRampRouter 1.0.0", s_router.typeAndVersion());

    // owner
    assertEq(OWNER, s_router.owner());

    // router config
    address[] memory configuredOffRamps = s_router.getOffRamps();
    assertEq(s_offRamps.length, configuredOffRamps.length);
    for (uint256 i = 0; i < s_offRamps.length; ++i) {
      address testOffRamp = s_offRamps[i];
      assertEq(testOffRamp, configuredOffRamps[i]);
      assertTrue(s_router.isOffRamp(testOffRamp));
    }
  }
}

/// @notice #routeMessage
contract EVM2EVMTollOffRampRouter_routeMessage is EVM2EVMTollOffRampRouterSetup {
  MaybeRevertMessageReceiver internal s_revertingReceiver;
  IAny2EVMMessageReceiver internal s_receiver;

  function setUp() public virtual override {
    EVM2EVMTollOffRampRouterSetup.setUp();

    s_revertingReceiver = new MaybeRevertMessageReceiver(true);
    s_receiver = new SimpleMessageReceiver();
  }

  // Success

  function testSuccess() public {
    Common.Any2EVMMessage memory message = _generateMockMessage();
    changePrank(address(s_offRamps[0]));
    vm.expectEmit(false, false, false, true);
    emit MessageReceived();

    s_router.routeMessage(message, false, GAS_LIMIT, address(s_receiver));
  }

  function testMessageFailureReturnsFalseSuccess() public {
    Common.Any2EVMMessage memory message = _generateMockMessage();
    changePrank(address(s_offRamps[0]));
    assertFalse(s_router.routeMessage(message, false, GAS_LIMIT, address(s_revertingReceiver)));
  }

  function testNotEnoughMessageGasLimitReturnsFalseSuccess() public {
    Common.Any2EVMMessage memory message = _generateMockMessage();
    changePrank(address(s_offRamps[0]));
    assertFalse(s_router.routeMessage(message, false, 1, address(s_receiver)));
  }

  // Reverts

  function testMustCallFromOffRampReverts() public {
    Common.Any2EVMMessage memory message = _generateMockMessage();
    vm.expectRevert(abi.encodeWithSelector(IAny2EVMOffRampRouter.MustCallFromOffRamp.selector, IBaseOffRamp(OWNER)));
    s_router.routeMessage(message, false, GAS_LIMIT, STRANGER);
  }

  function testZeroAddressReceiverReverts() public {
    Common.Any2EVMMessage memory message = _generateMockMessage();
    changePrank(address(s_offRamps[0]));
    vm.expectRevert();
    s_router.routeMessage(message, false, GAS_LIMIT, address(0));
  }
}

/// @notice #addOffRamp
contract EVM2EVMTollOffRampRouter_addOffRamp is EVM2EVMTollOffRampRouterSetup {
  address internal s_newOffRamp;

  event OffRampAdded(address indexed offRamp);

  function setUp() public virtual override {
    EVM2EVMTollOffRampRouterSetup.setUp();

    s_newOffRamp = address(new MockOffRamp());
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
    address existingOffRamp = s_offRamps[0];
    vm.expectRevert(abi.encodeWithSelector(IAny2EVMOffRampRouter.AlreadyConfigured.selector, existingOffRamp));
    s_router.addOffRamp(existingOffRamp);
  }

  function testZeroAddressReverts() public {
    vm.expectRevert(IAny2EVMOffRampRouter.InvalidAddress.selector);
    s_router.addOffRamp(address(0));
  }
}

/// @notice #removeOffRamp
contract EVM2EVMTollOffRampRouter_removeOffRamp is EVM2EVMTollOffRampRouterSetup {
  event OffRampRemoved(address indexed offRamp);

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
    address newRamp = address(new MockOffRamp());
    vm.expectRevert(abi.encodeWithSelector(IAny2EVMOffRampRouter.OffRampNotAllowed.selector, newRamp));
    s_router.removeOffRamp(newRamp);
  }
}

/// @notice #getOffRamps
contract EVM2EVMTollOffRampRouter_getOffRamps is EVM2EVMTollOffRampRouterSetup {
  // Success
  function testSuccess() public {
    address[] memory offRamps = s_router.getOffRamps();
    assertEq(2, offRamps.length);
    assertEq(address(s_offRamps[0]), address(offRamps[0]));
    assertEq(address(s_offRamps[1]), address(offRamps[1]));
  }
}

/// @notice #isOffRamp
contract EVM2EVMTollOffRampRouter_isOffRamp is EVM2EVMTollOffRampRouterSetup {
  // Success
  function testSuccess() public {
    assertTrue(s_router.isOffRamp(s_offRamps[0]));
    assertTrue(s_router.isOffRamp(s_offRamps[1]));
    assertFalse(s_router.isOffRamp(address(1)));
  }
}
