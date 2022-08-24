// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./EVM2EVMSubscriptionOffRampSetup.t.sol";
import "../../helpers/receivers/MaybeRevertMessageReceiver.sol";

/// @notice #constructor
contract EVM2EVMSubscriptionOffRampRouter_constructor is EVM2EVMSubscriptionOffRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("Any2EVMSubscriptionOffRampRouter 1.0.0", s_router.typeAndVersion());

    // owner
    assertEq(OWNER, s_router.owner());

    // router config
    BaseOffRampInterface[] memory configuredOffRamps = s_router.getOffRamps();
    assertEq(1, configuredOffRamps.length);
    assertEq(address(s_offRamp), address(configuredOffRamps[0]));
    assertTrue(s_router.isOffRamp(s_offRamp));

    SubscriptionInterface.SubscriptionConfig memory configured = s_router.getSubscriptionConfig();

    SubscriptionInterface.SubscriptionConfig memory expected = subscriptionConfig(s_destFeeToken);

    assertEq(expected.setSubscriptionSenderDelay, configured.setSubscriptionSenderDelay);
    assertEq(expected.withdrawalDelay, configured.withdrawalDelay);
    assertEq(address(expected.feeToken), address(configured.feeToken));
  }
}

/// @notice #routeMessage
contract EVM2EVMSubscriptionOffRampRouter_routeMessage is EVM2EVMSubscriptionOffRampSetup {
  event MessageReceived(uint256 sequenceNumber);

  MaybeRevertMessageReceiver s_revertingReceiver;

  function setUp() public virtual override {
    EVM2EVMSubscriptionOffRampSetup.setUp();
    changePrank(address(s_offRamp));

    s_revertingReceiver = new MaybeRevertMessageReceiver(true);
  }

  // Success
  function testSuccess() public {
    CCIP.Any2EVMMessage memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageNoTokens(1, 1)
    );

    vm.expectEmit(false, false, false, true);
    emit MessageReceived(message.sequenceNumber);

    assertTrue(s_router.routeMessage(message));
  }

  function testMessageFailureReturnsFalseSuccess() public {
    CCIP.Any2EVMMessage memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageNoTokens(1, 1)
    );
    message.receiver = address(s_revertingReceiver);
    assertFalse(s_router.routeMessage(message));
  }

  function testNotEnoughMessageGasLimitReturnsFalseSuccess() public {
    CCIP.Any2EVMMessage memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageNoTokens(1, 1)
    );
    message.gasLimit = 1;
    assertFalse(s_router.routeMessage(message));
  }

  function testMessageFailureReturnsFalse() public {
    CCIP.Any2EVMMessage memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageNoTokens(1, 1)
    );
    message.receiver = address(s_revertingReceiver);
    assertFalse(s_router.routeMessage(message));
  }

  // Reverts
  function testMustCallFromOffRampReverts() public {
    changePrank(OWNER);
    vm.expectRevert(
      abi.encodeWithSelector(Any2EVMOffRampRouterInterface.MustCallFromOffRamp.selector, BaseOffRampInterface(OWNER))
    );
    s_router.routeMessage(_convertSubscriptionToGeneralMessage(_generateAny2EVMSubscriptionMessageNoTokens(1, 1)));
  }
}

/// @notice #chargeSubscription
contract EVM2EVMSubscriptionOffRampRouter_chargeSubscription is EVM2EVMSubscriptionOffRampSetup {
  event SubscriptionCharged(address receiver, uint256 amount);

  function setUp() public virtual override {
    EVM2EVMSubscriptionOffRampSetup.setUp();
    changePrank(address(s_offRamp));
  }

  // Success
  function testSuccess() public {
    address receiver = address(s_receiver);
    uint256 preBalance = s_router.getSubscription(receiver).balance;

    vm.expectEmit(false, false, false, true);
    emit SubscriptionCharged(receiver, SUBSCRIPTION_BALANCE);

    s_router.chargeSubscription(receiver, OWNER, SUBSCRIPTION_BALANCE);

    assertEq(preBalance - SUBSCRIPTION_BALANCE, s_router.getSubscription(receiver).balance);
  }

  // Reverts
  function testMustCallFromOffRampReverts() public {
    changePrank(OWNER);
    vm.expectRevert(abi.encodeWithSelector(Any2EVMOffRampRouterInterface.MustCallFromOffRamp.selector, OWNER));
    s_router.chargeSubscription(address(s_receiver), OWNER, SUBSCRIPTION_BALANCE);
  }

  function testBalanceTooLowReverts() public {
    vm.expectRevert(stdError.arithmeticError);
    s_router.chargeSubscription(address(s_receiver), OWNER, SUBSCRIPTION_BALANCE + 1);
  }

  function testSenderNotAllowedReverts() public {
    vm.expectRevert(abi.encodeWithSelector(Any2EVMOffRampRouterInterface.SenderNotAllowed.selector, STRANGER));
    s_router.chargeSubscription(address(s_receiver), STRANGER, SUBSCRIPTION_BALANCE);
  }
}
