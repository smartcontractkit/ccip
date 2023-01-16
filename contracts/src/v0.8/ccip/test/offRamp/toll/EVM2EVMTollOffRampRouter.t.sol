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
  IBaseOffRamp[] internal s_offRamps;

  function setUp() public virtual override {
    BaseTest.setUp();
    MockOffRamp offRamp = new MockOffRamp();
    s_offRamps.push(offRamp);
    s_router = new Any2EVMTollOffRampRouter(s_offRamps);
  }

  function _generateMockMessage() internal pure returns (Common.Any2EVMMessage memory) {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](0);
    address[] memory pools = new address[](0);
    return (
      Common.Any2EVMMessage({
        sourceChainId: SOURCE_CHAIN_ID,
        sender: abi.encode(STRANGER),
        data: abi.encode(0),
        destTokensAndAmounts: tokensAndAmounts
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
    IBaseOffRamp[] memory configuredOffRamps = s_router.getOffRamps();
    assertEq(s_offRamps.length, configuredOffRamps.length);
    for (uint256 i = 0; i < s_offRamps.length; ++i) {
      IBaseOffRamp testOffRamp = s_offRamps[i];
      assertEq(address(testOffRamp), address(configuredOffRamps[i]));
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
