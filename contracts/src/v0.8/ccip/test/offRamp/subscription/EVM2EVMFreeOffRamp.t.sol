// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "./EVM2EVMFreeOffRampSetup.t.sol";

/// @notice #constructor
contract EVM2EVMFreeOffRamp_constructor is EVM2EVMFreeOffRampSetup {
  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMSubscriptionOffRamp 1.0.0", s_offRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_offRamp.owner());

    // OffRamp config
    (uint256 source, uint256 dest) = s_offRamp.getChainIDs();
    assertEq(SOURCE_CHAIN_ID, source);
    assertEq(DEST_CHAIN_ID, dest);
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
    IERC20[] memory pools = s_offRamp.getPoolTokens();
    assertEq(pools.length, s_sourceTokens.length);
    assertTrue(address(pools[0]) == address(s_sourceTokens[0]));
    assertTrue(address(pools[1]) == address(s_sourceTokens[1]));

    // HealthChecker
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
  }
}

/// @notice #setRouter
contract EVM2EVMFreeOffRamp_setRouter is EVM2EVMFreeOffRampSetup {
  // Success

  // Assert that setRouter will set the router to the given router argument.
  function testSuccess() public {
    Any2EVMOffRampRouterInterface newRouter = Any2EVMSubscriptionOffRampRouter(address(1));
    assertTrue(address(newRouter) != address(s_offRamp.getRouter()));
    s_offRamp.setRouter(newRouter);
    assertEq(address(newRouter), address(s_offRamp.getRouter()));
  }

  function testZeroRouterSuccess() public {
    Any2EVMOffRampRouterInterface newRouter = Any2EVMSubscriptionOffRampRouter(address(0));
    assertTrue(address(newRouter) != address(s_offRamp.getRouter()));
    s_offRamp.setRouter(newRouter);
    assertEq(address(newRouter), address(s_offRamp.getRouter()));
  }

  // Reverts

  function testOwnerReverts() public {
    Any2EVMOffRampRouterInterface newRouter = Any2EVMSubscriptionOffRampRouter(address(1));

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_offRamp.setRouter(newRouter);
  }
}

/// @notice ccipReceive
contract EVM2EVMFreeOffRamp_ccipReceive is EVM2EVMFreeOffRampSetup {
  // Reverts

  function testReverts() public {
    vm.expectRevert();
    s_offRamp.ccipReceive(_convertSubscriptionToGeneralMessage(_generateAny2EVMSubscriptionMessageNoTokens(1, 1)));
  }
}

/// @notice #executeSingleMessage
contract EVM2EVMFreeOffRamp_executeSingleMessage is EVM2EVMFreeOffRampSetup {
  // Success

  // Assert that a self call to executeSingleMessage with a valid receiver will succeed.
  function testSuccessNoTokensSuccess() public {
    changePrank(address(s_offRamp));
    s_offRamp.executeSingleMessage(
      _convertSubscriptionToGeneralMessage(_generateAny2EVMSubscriptionMessageNoTokens(1, 1))
    );
  }

  // Assert that any call to executeSingleMessage with an EOA will still succeed
  function testNonContractSuccess() public {
    changePrank(address(s_offRamp));
    CCIP.Any2EVMMessageFromSender memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageNoTokens(1, 1)
    );
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message);
  }

  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  // Assert that any call to executeSingleMessage with an EOA will still succeed with tokens
  function testNonContractWithTokensSuccess() public {
    changePrank(address(s_offRamp));
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    vm.expectEmit(true, true, false, true);
    emit Released(address(s_offRamp), STRANGER, amounts[0]);
    vm.expectEmit(true, true, false, true);
    emit Minted(address(s_offRamp), STRANGER, amounts[1]);
    CCIP.Any2EVMMessageFromSender memory message = _convertSubscriptionToGeneralMessage(
      _generateAny2EVMSubscriptionMessageWithTokens(1, 1, amounts)
    );
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message);
  }

  // Reverts

  // Asserts that any call to executeSingleMessage will revert when not
  // it's not a self call.
  function testNoSelfCall() public {
    vm.expectRevert(BaseOffRampInterface.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(
      _convertSubscriptionToGeneralMessage(_generateAny2EVMSubscriptionMessageNoTokens(1, 1))
    );
  }
}

/// @notice #execute
contract EVM2EVMFreeOffRamp_execute is EVM2EVMFreeOffRampSetup {
  event SubscriptionCharged(address receiver, uint256 amount);

  // Success

  // Asserts that the nonces and seq nums are stored in the contract
  // and remain incremented after the execute calls.
  function testIncrementsNonceAndSeqNum() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = new CCIP.EVM2EVMSubscriptionMessage[](3);
    messages[0] = _generateAny2EVMSubscriptionMessageNoTokens(1, 1);
    messages[1] = _generateAny2EVMSubscriptionMessageNoTokens(2, 2);
    messages[2] = _generateAny2EVMSubscriptionMessageNoTokens(3, 3);

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    messages[0] = _generateAny2EVMSubscriptionMessageNoTokens(4, 4);
    messages[1] = _generateAny2EVMSubscriptionMessageNoTokens(5, 5);
    messages[2] = _generateAny2EVMSubscriptionMessageNoTokens(6, 6);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a properly formed call to execute will succeed.
  function testNoTokensSingleMessageSuccess() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a call to execute succeeds even though the call
  // to execute the individual tx fails. The resulting tx state is set to Failed.
  function testNoTokensSingleMessageFailedCallSuccess() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateBasicMessages();
    MaybeRevertMessageReceiver newReceiver = new MaybeRevertMessageReceiver(true);
    messages[0].receiver = address(newReceiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.FAILURE);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that executing two messages emits the proper events.
  function testWithTokensSuccess() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);
    messages[1].nonce = 1;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[1].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that executing two messages emits the proper events when executed manually.
  function testWithTokensSuccessWithManualExecution() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);
    messages[1].nonce = 1;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[1].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), true);
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
        receiver: SubscriptionManagerInterface(address(newReceiver)),
        strictSequencing: false,
        balance: SUBSCRIPTION_BALANCE
      })
    );

    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateBasicMessages();
    messages[0].nonce = 100;
    messages[0].receiver = address(newReceiver);

    s_offRamp.setMessageState(messages[0].sequenceNumber, CCIP.MessageExecutionState.FAILURE);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), true);
  }

  // Reverts

  // Asserts that a call to execute will revert when the router is unset.
  function testNoRouterSetReverts() public {
    Any2EVMOffRampRouterInterface newRouter = Any2EVMSubscriptionOffRampRouter(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(BaseOffRampInterface.RouterNotSet.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), false);
  }

  function testNoMessagesReverts() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = new CCIP.EVM2EVMSubscriptionMessage[](0);
    vm.expectRevert(BaseOffRampInterface.NoMessagesToExecute.selector);
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testRootNotCommittedReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(CommitStoreInterface.verify.selector),
      abi.encode(0)
    );
    vm.expectRevert(BaseOffRampInterface.RootNotCommitted.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(CommitStoreInterface.verify.selector),
      abi.encode(BLOCK_TIME)
    );
    vm.expectRevert(BaseOffRampInterface.ManualExecutionNotYetEnabled.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
    vm.clearMockedCalls();
  }

  function testUnsupportedTokenReverts() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateMessagesWithTokens();
    address unknownToken = address(50);
    messages[0].tokensAndAmounts[0].token = unknownToken;

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, unknownToken));

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
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a call to execute will revert when a message has data that
  // exceeds the maximum data length.
  function testMessageDataTooLargeReverts() public {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    CCIP.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.MessageTooLarge.selector, MAX_DATA_SIZE, messages[0].data.length)
    );
    s_offRamp.execute(executionReport, false);
  }
}

/// @notice #_report
contract EVM2EVMFreeOffRamp__report is EVM2EVMFreeOffRampSetup {
  // Asserts that execute is called with the proper arguments.
  function testSuccess() public {
    CCIP.ExecutionReport memory report = _generateReportFromMessages(_generateBasicMessages());

    vm.expectCall(address(s_offRamp), abi.encodeCall(s_offRamp.execute, (report, false)));

    s_offRamp.report(abi.encode(report));
  }
}
