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
  BaseOffRampInterface[] internal s_offRamps;

  function setUp() public virtual override {
    BaseTest.setUp();
    MockOffRamp offRamp = new MockOffRamp();
    s_offRamps.push(offRamp);
    s_router = new Any2EVMTollOffRampRouter(s_offRamps);
  }

  function _generateMockMessage(address receiver) internal pure returns (Internal.Any2EVMMessageFromSender memory) {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](0);
    address[] memory pools = new address[](0);
    return (
      Internal.Any2EVMMessageFromSender({
        sourceChainId: SOURCE_CHAIN_ID,
        sender: abi.encode(STRANGER),
        receiver: receiver,
        data: abi.encode(0),
        destTokensAndAmounts: tokensAndAmounts,
        destPools: pools,
        gasLimit: GAS_LIMIT
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
contract EVM2EVMTollOffRampRouter_routeMessage is EVM2EVMTollOffRampRouterSetup {
  MaybeRevertMessageReceiver internal s_revertingReceiver;
  Any2EVMMessageReceiverInterface internal s_receiver;

  function setUp() public virtual override {
    EVM2EVMTollOffRampRouterSetup.setUp();

    s_revertingReceiver = new MaybeRevertMessageReceiver(true);
    s_receiver = new SimpleMessageReceiver();
  }

  // Success

  function testSuccess() public {
    Internal.Any2EVMMessageFromSender memory message = _generateMockMessage(address(s_receiver));
    changePrank(address(s_offRamps[0]));
    vm.expectEmit(false, false, false, true);
    emit MessageReceived();

    s_router.routeMessage(message);
  }

  function testMessageFailureReturnsFalseSuccess() public {
    Internal.Any2EVMMessageFromSender memory message = _generateMockMessage(address(s_revertingReceiver));
    changePrank(address(s_offRamps[0]));
    assertFalse(s_router.routeMessage(message));
  }

  function testNotEnoughMessageGasLimitReturnsFalseSuccess() public {
    Internal.Any2EVMMessageFromSender memory message = _generateMockMessage(address(s_receiver));
    message.gasLimit = 1;
    changePrank(address(s_offRamps[0]));
    assertFalse(s_router.routeMessage(message));
  }

  // Reverts

  function testMustCallFromOffRampReverts() public {
    Internal.Any2EVMMessageFromSender memory message = _generateMockMessage(STRANGER);
    vm.expectRevert(
      abi.encodeWithSelector(Any2EVMOffRampRouterInterface.MustCallFromOffRamp.selector, BaseOffRampInterface(OWNER))
    );
    s_router.routeMessage(message);
  }

  function testZeroAddressReceiverReverts() public {
    Internal.Any2EVMMessageFromSender memory message = _generateMockMessage(address(0));
    changePrank(address(s_offRamps[0]));
    vm.expectRevert();
    s_router.routeMessage(message);
  }
}
