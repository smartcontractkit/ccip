// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../helpers/receivers/RevertingMessageReceiver.sol";
import "./Any2EVMMOOffRampSetup.t.sol";

/// @notice #constructor
contract Any2EVMMOOffRamp_constructor is Any2EVMMOOffRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("Any2EVMMOOffRamp 1.0.0", s_offRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_offRamp.owner());

    // OffRamp config
    assertEq(SOURCE_CHAIN_ID, s_offRamp.SOURCE_CHAIN_ID());
    assertEq(DEST_CHAIN_ID, s_offRamp.CHAIN_ID());
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
    IERC20[] memory pools = s_offRamp.getPoolTokens();
    assertEq(pools.length, s_sourceTokens.length);
    assertTrue(address(pools[0]) == address(s_sourceTokens[0]));
    assertTrue(address(pools[1]) == address(s_sourceTokens[1]));

    // HealthChecker
    assertEq(HEARTBEAT, s_offRamp.getMaxSecondsWithoutAFNHeartbeat());
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
  }
}

/// @notice #setRouter
contract Any2EVMMOOffRamp_setRouter is Any2EVMMOOffRampSetup {
  // Success

  // Assert that setRouter will set the router to the given router argument.
  function testSuccess() public {
    Any2EVMMOOffRampRouterInterface newRouter = Any2EVMMOOffRampRouter(address(1));
    assertTrue(address(newRouter) != address(s_offRamp.s_router()));
    s_offRamp.setRouter(newRouter);
    assertEq(address(newRouter), address(s_offRamp.s_router()));
  }

  function testZeroRouterSuccess() public {
    Any2EVMMOOffRampRouterInterface newRouter = Any2EVMMOOffRampRouter(address(0));
    assertTrue(address(newRouter) != address(s_offRamp.s_router()));
    s_offRamp.setRouter(newRouter);
    assertEq(address(newRouter), address(s_offRamp.s_router()));
  }

  // Reverts

  function testOwnerReverts() public {
    Any2EVMMOOffRampRouterInterface newRouter = Any2EVMMOOffRampRouter(address(1));

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_offRamp.setRouter(newRouter);
  }
}

/// @notice ccipReceive
contract Any2EVMMOOffRamp_ccipReceive is Any2EVMMOOffRampSetup {
  // Reverts

  function testReverts() public {
    vm.expectRevert();
    s_offRamp.ccipReceive(getAny2EVMMOMessage(1, 1));
  }
}

/// @notice #executeSingleMessage
contract Any2EVMMOOffRamp_executeSingleMessage is Any2EVMMOOffRampSetup {
  // Success

  // Assert that a self call to executeSingleMessage with a valid receiver will succeed.
  function testSuccessNoTokensSuccess() public {
    changePrank(address(s_offRamp));
    s_offRamp.executeSingleMessage(getAny2EVMMOMessage(1, 1));
  }

  // Reverts

  // Asserts that any call to executeSingleMessage will revert when not
  // it's not a self call.
  function testNoSelfCall() public {
    vm.expectRevert(BaseOffRampInterface.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(getAny2EVMMOMessage(1, 1));
  }

  // Assert that any call to executeSingleMessage with an invalid receiver
  // will revert.
  function testInvalidReceiver() public {
    changePrank(address(s_offRamp));
    CCIP.Any2EVMMOMessage memory message = getAny2EVMMOMessage(1, 1);
    message.receiver = STRANGER;
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidReceiver.selector, STRANGER));
    s_offRamp.executeSingleMessage(message);
  }
}

/// @notice #execute
contract Any2EVMMOOffRamp_execute is Any2EVMMOOffRampSetup {
  event SubscriptionCharged(address receiver, uint256 amount);

  // Success

  // Asserts that the nonces and seq nums are stored in the contract
  // and remain incremented after the execute calls.
  function testIncrementsNonceAndSeqNum() public {
    CCIP.Any2EVMMOMessage[] memory messages = new CCIP.Any2EVMMOMessage[](3);
    messages[0] = getAny2EVMMOMessage(1, 1);
    messages[1] = getAny2EVMMOMessage(2, 2);
    messages[2] = getAny2EVMMOMessage(3, 3);

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    messages[0] = getAny2EVMMOMessage(4, 4);
    messages[1] = getAny2EVMMOMessage(5, 5);
    messages[2] = getAny2EVMMOMessage(6, 6);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a properly formed call to execute will succeed.
  function testNoTokensSingleMessageSuccess() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a call to execute succeeds even though the call
  // to execute the individual tx fails. The resulting tx state is set to Failed.
  function testNoTokensSingleMessageFailedCallSuccess() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    RevertingMessageReceiver newReceiver = new RevertingMessageReceiver();
    createSubscription(newReceiver, s_router, SUBSCRIPTION_BALANCE);
    messages[0].receiver = address(newReceiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Failure);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that executing two messages emits the proper events when executed manually.
  function testSuccessWithManualExecution() public {
    CCIP.Any2EVMMOMessage[] memory messages = new CCIP.Any2EVMMOMessage[](2);
    messages[0] = getAny2EVMMOMessage(1, 1);
    messages[1] = getAny2EVMMOMessage(2, 1);
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);
    messages[1].nonce = 1;

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[1].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), true);
  }

  function testChargeSubscriptionSuccess() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    uint256 balancePreTx = s_router.getSubscription(address(s_receiver)).balance;

    // Explicitly do NOT check the data as we don't know how much gas was
    // used. We use 0 as placeholder and only assert that the event was emitted.
    vm.expectEmit(false, false, false, false);
    emit SubscriptionCharged(address(s_receiver), 0);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    uint256 balancePostTx = s_router.getSubscription(address(s_receiver)).balance;

    // No exact gas calculations here because they can change very easily. It just checks
    // that some reasonable amount of gas was taken from the proper subscription.
    assertTrue(balancePreTx - 1e4 > balancePostTx);
    assertTrue(balancePostTx > SUBSCRIPTION_BALANCE - 1e5);
  }

  // Asserts that a failed message can be executed even when the nonce is out
  // of order as long as strict sequencing is off.
  function testStrictSequencingStateFailureSuccess() public {
    SimpleMessageReceiver newReceiver = new SimpleMessageReceiver();
    address[] memory senders = new address[](1);
    senders[0] = OWNER;
    s_destFeeToken.approve(address(s_router), SUBSCRIPTION_BALANCE);
    s_router.createSubscription(
      SubscriptionInterface.OffRampSubscription({
        senders: senders,
        receiver: newReceiver,
        strictSequencing: false,
        balance: SUBSCRIPTION_BALANCE
      })
    );

    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    messages[0].nonce = 100;
    messages[0].receiver = address(newReceiver);

    s_offRamp.setMessageState(messages[0].sequenceNumber, CCIP.MessageExecutionState.Failure);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), true);
  }

  // Reverts

  // Asserts that a call to execute will revert when the router is unset.
  function testNoRouterSetReverts() public {
    Any2EVMMOOffRampRouterInterface newRouter = Any2EVMMOOffRampRouter(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(BaseOffRampInterface.RouterNotSet.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), false);
  }

  function testNoMessagesReverts() public {
    CCIP.Any2EVMMOMessage[] memory messages = new CCIP.Any2EVMMOMessage[](0);
    vm.expectRevert(BaseOffRampInterface.NoMessagesToExecute.selector);
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testRootNotRelayedReverts() public {
    vm.mockCall(
      address(s_mockBlobVerifier),
      abi.encodeWithSelector(BlobVerifierInterface.verify.selector),
      abi.encode(0)
    );
    vm.expectRevert(BaseOffRampInterface.RootNotRelayed.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockBlobVerifier),
      abi.encodeWithSelector(BlobVerifierInterface.verify.selector),
      abi.encode(BLOCK_TIME)
    );
    vm.expectRevert(BaseOffRampInterface.ManualExecutionNotYetEnabled.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testNotNextInSequenceReverts() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    messages[0].nonce = 100;

    vm.expectRevert(abi.encodeWithSelector(Any2EVMMOOffRampInterface.IncorrectNonce.selector, messages[0].nonce));

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a call to execute will revert if a message in the execution report
  // is already executed.
  function testAlreadyExecutedReverts() public {
    CCIP.ExecutionReport memory executionReport = _generateReportFromMessages(_generateBasicMessages());
    s_offRamp.execute(executionReport, false);
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.AlreadyExecuted.selector, executionReport.sequenceNumbers[0])
    );
    s_offRamp.execute(executionReport, false);
  }

  // Asserts that a call to execute will revert when a message has the wrong
  // source chain id.
  function testInvalidSourceChainReverts() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a call to execute will revert when a message has data that
  // exceeds the maximum data length.
  function testMessageDataTooLargeReverts() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    CCIP.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOffRampInterface.MessageTooLarge.selector,
        s_offRampConfig.maxDataSize,
        messages[0].data.length
      )
    );
    s_offRamp.execute(executionReport, false);
  }

  // Asserts that the tx reverts when a subscription is not found.
  function testSubscriptionNotFound() public {
    CCIP.Any2EVMMOMessage[] memory messages = _generateBasicMessages();
    SimpleMessageReceiver new_receiver = new SimpleMessageReceiver();
    messages[0].receiver = address(new_receiver);

    vm.expectRevert(abi.encodeWithSelector(SubscriptionInterface.SubscriptionNotFound.selector, new_receiver));
    s_offRamp.execute(_generateReportFromMessages(messages), true);
  }

  // Asserts that the tx reverts when the balance of the subscription is
  // too low.
  function testBalanceTooLow() public {
    CCIP.Any2EVMMOMessage[] memory messages = new CCIP.Any2EVMMOMessage[](3);
    messages[0] = getAny2EVMMOMessage(1, 1);
    messages[1] = getAny2EVMMOMessage(2, 2);
    messages[2] = getAny2EVMMOMessage(3, 3);

    CCIP.ExecutionReport memory report = _generateReportFromMessages(messages);
    report.tokenPerFeeCoin[0] = 1e25;

    vm.expectRevert(SubscriptionInterface.BalanceTooLow.selector);

    s_offRamp.execute(report, false);
  }
}
