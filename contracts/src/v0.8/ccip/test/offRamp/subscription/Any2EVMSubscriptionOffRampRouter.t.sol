// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./Any2EVMSubscriptionOffRampSetup.t.sol";
import "../../helpers/receivers/RevertingMessageReceiver.sol";

/// @notice #constructor
contract Any2EVMSubscriptionOffRampRouter_constructor is Any2EVMSubscriptionOffRampSetup {
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

    SubscriptionInterface.SubscriptionConfig memory subscriptionConfig = s_router.getSubscriptionConfig();

    assertEq(s_subscriptionConfig.setSubscriptionSenderDelay, subscriptionConfig.setSubscriptionSenderDelay);
    assertEq(s_subscriptionConfig.withdrawalDelay, subscriptionConfig.withdrawalDelay);
    assertEq(address(s_subscriptionConfig.feeToken), address(subscriptionConfig.feeToken));
  }
}

/// @notice #routeMessage
contract Any2EVMSubscriptionOffRampRouter_routeMessage is Any2EVMSubscriptionOffRampSetup {
  event MessageReceived(uint256 sequenceNumber);

  CrossChainMessageReceiverInterface s_revertingReceiver;

  function setUp() public virtual override {
    Any2EVMSubscriptionOffRampSetup.setUp();
    changePrank(address(s_offRamp));

    s_revertingReceiver = new RevertingMessageReceiver();
  }

  // Success
  function testSuccess() public {
    CCIP.Any2EVMSubscriptionMessage memory message = getAny2EVMSubscriptionMessageNoTokens(1, 1);

    vm.expectEmit(false, false, false, true);
    emit MessageReceived(message.sequenceNumber);

    s_router.routeMessage(s_receiver, message);
  }

  // Reverts
  function testMustCallFromOffRampReverts() public {
    changePrank(OWNER);
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampRouterInterface.MustCallFromOffRamp.selector, BaseOffRampInterface(OWNER))
    );
    s_router.routeMessage(s_receiver, getAny2EVMSubscriptionMessageNoTokens(1, 1));
  }

  function testZeroAddressReceiverReverts() public {
    vm.expectRevert();

    s_router.routeMessage(CrossChainMessageReceiverInterface(address(0)), getAny2EVMSubscriptionMessageNoTokens(1, 1));
  }

  function testMessageFailureReverts() public {
    CCIP.Any2EVMSubscriptionMessage memory message = getAny2EVMSubscriptionMessageNoTokens(1, 1);
    message.receiver = address(s_revertingReceiver);
    bytes memory reason;
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampRouterInterface.MessageFailure.selector, message.sequenceNumber, reason)
    );
    s_router.routeMessage(s_revertingReceiver, message);
  }
}

/// @notice #chargeSubscription
contract Any2EVMSubscriptionOffRampRouter_chargeSubscription is Any2EVMSubscriptionOffRampSetup {
  event SubscriptionCharged(address receiver, uint256 amount);

  function setUp() public virtual override {
    Any2EVMSubscriptionOffRampSetup.setUp();
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
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampRouterInterface.MustCallFromOffRamp.selector, OWNER));
    s_router.chargeSubscription(address(s_receiver), OWNER, SUBSCRIPTION_BALANCE);
  }

  function testBalanceTooLowReverts() public {
    vm.expectRevert(SubscriptionInterface.BalanceTooLow.selector);
    s_router.chargeSubscription(address(s_receiver), OWNER, SUBSCRIPTION_BALANCE + 1);
  }

  function testSenderNotAllowedReverts() public {
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampRouterInterface.SenderNotAllowed.selector, STRANGER));
    s_router.chargeSubscription(address(s_receiver), STRANGER, SUBSCRIPTION_BALANCE);
  }
}
