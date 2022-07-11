// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../BaseTest.t.sol";
import "../../../offRamp/toll/Any2EVMTollOffRampRouter.sol";
import "../../mocks/MockOffRamp.sol";
import "../../helpers/receivers/RevertingMessageReceiver.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";

contract Any2EVMTollOffRampRouterSetup is BaseTest {
  event MessageReceived(uint256 sequenceNumber);

  Any2EVMTollOffRampRouter internal s_router;
  BaseOffRampInterface[] internal s_offRamps;

  function setUp() public virtual override {
    BaseTest.setUp();
    MockOffRamp offRamp = new MockOffRamp();
    s_offRamps.push(offRamp);
    s_router = new Any2EVMTollOffRampRouter(s_offRamps);
  }

  function _generateMockMessage(address receiver)
    internal
    pure
    returns (CrossChainMessageReceiverInterface, CCIP.Any2EVMTollMessage memory)
  {
    IERC20[] memory tokens = new IERC20[](0);
    uint256[] memory amounts = new uint256[](0);
    return (
      CrossChainMessageReceiverInterface(receiver),
      CCIP.Any2EVMTollMessage({
        sourceChainId: SOURCE_CHAIN_ID,
        sequenceNumber: 1,
        sender: STRANGER,
        receiver: receiver,
        data: abi.encode(0),
        tokens: tokens,
        amounts: amounts,
        feeToken: IERC20(address(0)),
        feeTokenAmount: 0,
        gasLimit: 0
      })
    );
  }
}

/// @notice #constructor
contract Any2EVMTollOffRampRouter_constructor is Any2EVMTollOffRampRouterSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("Any2EVMTollOffRampRouter 1.0.0", s_router.typeAndVersion());

    // owner
    assertEq(OWNER, s_router.owner());

    // router config
    BaseOffRampInterface[] memory configuredOffRamps = s_router.getOffRamps();
    assertEq(s_offRamps.length, configuredOffRamps.length);
    for (uint256 i = 0; i < s_offRamps.length; ++i) {
      BaseOffRampInterface testOffRamp = s_offRamps[i];
      assertEq(address(testOffRamp), address(configuredOffRamps[i]));
      assertTrue(s_router.isOffRamp(testOffRamp));
    }
  }
}

/// @notice #routeMessage
contract Any2EVMTollOffRampRouter_routeMessage is Any2EVMTollOffRampRouterSetup {
  CrossChainMessageReceiverInterface internal s_revertingReceiver;
  CrossChainMessageReceiverInterface internal s_receiver;

  function setUp() public virtual override {
    Any2EVMTollOffRampRouterSetup.setUp();

    s_revertingReceiver = new RevertingMessageReceiver();
    s_receiver = new SimpleMessageReceiver();
  }

  // Success

  function testSuccess() public {
    (CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage memory message) = _generateMockMessage(
      address(s_receiver)
    );
    changePrank(address(s_offRamps[0]));
    vm.expectEmit(false, false, false, true);
    emit MessageReceived(message.sequenceNumber);

    s_router.routeMessage(receiver, message);
  }

  // Reverts

  function testMustCallFromOffRampReverts() public {
    (CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage memory message) = _generateMockMessage(
      STRANGER
    );
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampRouterInterface.MustCallFromOffRamp.selector, BaseOffRampInterface(OWNER))
    );
    s_router.routeMessage(receiver, message);
  }

  function testZeroAddressReceiverReverts() public {
    (CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage memory message) = _generateMockMessage(
      address(0)
    );
    changePrank(address(s_offRamps[0]));
    vm.expectRevert();
    s_router.routeMessage(receiver, message);
  }

  function testReceiveReverts() public {
    (CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage memory message) = _generateMockMessage(
      address(s_revertingReceiver)
    );
    changePrank(address(s_offRamps[0]));
    bytes memory reason;
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampRouterInterface.MessageFailure.selector, message.sequenceNumber, reason)
    );
    s_router.routeMessage(receiver, message);
  }
}
