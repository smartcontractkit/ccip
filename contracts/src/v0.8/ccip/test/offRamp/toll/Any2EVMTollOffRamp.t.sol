// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./Any2EVMTollOffRampSetup.t.sol";
import "../../../offRamp/toll/Any2EVMTollOffRampRouter.sol";
import "../../helpers/receivers/RevertingMessageReceiver.sol";

/// @notice #constructor
contract Any2EVMTollOffRamp_constructor is Any2EVMTollOffRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("Any2EVMTollOffRamp 1.0.0", s_offRamp.typeAndVersion());

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
contract Any2EVMTollOffRamp_setRouter is Any2EVMTollOffRampSetup {
  Any2EVMTollOffRampRouterInterface public s_router;

  event OffRampRouterSet(address indexed router);

  function setUp() public virtual override {
    Any2EVMTollOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  // Success

  function testSuccessNewAddress() public {
    _testSuccess(_generateNewRouter());
  }

  function testSuccessZeroAddress() public {
    _testSuccess(Any2EVMTollOffRampRouterInterface(address(0)));
  }

  function _testSuccess(Any2EVMTollOffRampRouterInterface newRouter) private {
    vm.expectEmit(true, false, false, true);
    emit OffRampRouterSet(address(newRouter));

    s_offRamp.setRouter(newRouter);

    assertEq(address(newRouter), address(s_offRamp.getRouter()));
  }

  // Reverts

  function testOwnerReverts() public {
    Any2EVMTollOffRampRouterInterface newRouter = _generateNewRouter();

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_offRamp.setRouter(newRouter);
  }
}

/// @notice #ccipReceive
contract Any2EVMTollOffRamp_ccipReceive is Any2EVMTollOffRampSetup {
  // Reverts

  function testReverts() public {
    vm.expectRevert();
    s_offRamp.ccipReceive(_generateAny2EVMTollMessageNoTokens(1));
  }
}

/// @notice #execute
contract Any2EVMTollOffRamp_execute is Any2EVMTollOffRampSetup {
  Any2EVMTollOffRampRouterInterface s_router;

  function setUp() public virtual override {
    Any2EVMTollOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  // Success

  function testSingleMessageNoTokensSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a message execution fails, but it does
  // not disrupt the overall execution of the batch
  function testSingleMessageFailureSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateBasicMessages();
    RevertingMessageReceiver newReceiver = new RevertingMessageReceiver();
    messages[0].receiver = address(newReceiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Failure);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[1].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensAndTollSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[1].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), true);
  }

  // Reverts

  function testPausedReverts() public {
    s_offRamp.pause();
    vm.expectRevert("Pausable: paused");
    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
  }

  function testRouterNotSetReverts() public {
    Any2EVMTollOffRampRouterInterface newRouter = MockTollOffRampRouter(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(BaseOffRampInterface.RouterNotSet.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), false);
  }

  function testNoMessagesReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](0);
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

  function testAlreadyExecutedReverts() public {
    CCIP.ExecutionReport memory executionReport = _generateReportFromMessages(_generateBasicMessages());
    s_offRamp.execute(executionReport, false);
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.AlreadyExecuted.selector, executionReport.sequenceNumbers[0])
    );
    s_offRamp.execute(executionReport, false);
  }

  function testInvalidSourceChainReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateBasicMessages();
    IERC20[] memory newTokens = new IERC20[](1);
    newTokens[0] = s_sourceTokens[0];
    messages[0].tokens = newTokens;

    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.UnsupportedNumberOfTokens.selector, messages[0].sequenceNumber)
    );
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testMessageTooLargeReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateBasicMessages();
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

  function testUnsupportedTokenReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokens[0] = s_destTokens[0];
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, s_destTokens[0]));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }
}

/// @notice #executeSingleMessage
contract Any2EVMTollOffRamp_executeSingleMessage is Any2EVMTollOffRampSetup {
  Any2EVMTollOffRampRouterInterface public s_router;

  function setUp() public virtual override {
    Any2EVMTollOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
    changePrank(address(s_offRamp));
  }

  // Success

  function testNoTokensSuccess() public {
    s_offRamp.executeSingleMessage(_generateAny2EVMTollMessageNoTokens(1));
  }

  function testTokensSuccess() public {
    s_offRamp.executeSingleMessage(_generateMessagesWithTokens()[0]);
  }

  // Reverts

  function testMessageSenderReverts() public {
    vm.stopPrank();
    vm.expectRevert(BaseOffRampInterface.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(_generateAny2EVMTollMessageNoTokens(1));
  }

  function testUnsupportedTokenReverts() public {
    CCIP.Any2EVMTollMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokens[0] = s_destTokens[0];
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, s_destTokens[0]));
    s_offRamp.executeSingleMessage(messages[0]);
  }

  function testInvalidReceiverReverts() public {
    CCIP.Any2EVMTollMessage memory message = _generateAny2EVMTollMessageNoTokens(1);
    message.receiver = STRANGER;
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidReceiver.selector, STRANGER));
    s_offRamp.executeSingleMessage(message);
  }
}
