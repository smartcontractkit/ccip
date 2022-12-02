// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "../../../health/HealthChecker.sol";
import "./EVM2EVMGEOffRampSetup.t.sol";
import "../../../router/GERouter.sol";

/// @notice #constructor
contract EVM2EVMGEOffRamp_constructor is EVM2EVMGEOffRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMGEOffRamp 1.0.0", s_offRamp.typeAndVersion());

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
contract EVM2EVMGEOffRamp_setRouter is EVM2EVMGEOffRampSetup {
  Any2EVMOffRampRouterInterface public s_router;

  event OffRampRouterSet(address indexed router, uint256 sourceChainId, address onRampAddress);

  function setUp() public virtual override {
    EVM2EVMGEOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new GERouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  function _generateNewRouter() internal returns (GERouter newRouter) {
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](0);
    newRouter = new GERouter(offRamps);
  }

  // Success

  function testSuccessNewAddress() public {
    _testSuccess(_generateNewRouter());
  }

  function testSuccessZeroAddress() public {
    _testSuccess(Any2EVMOffRampRouterInterface(address(0)));
  }

  function _testSuccess(Any2EVMOffRampRouterInterface newRouter) private {
    vm.expectEmit(true, false, false, true);
    emit OffRampRouterSet(address(newRouter), SOURCE_CHAIN_ID, ON_RAMP_ADDRESS);

    s_offRamp.setRouter(newRouter);

    assertEq(address(newRouter), address(s_offRamp.getRouter()));
  }

  // Reverts

  function testOwnerReverts() public {
    Any2EVMOffRampRouterInterface newRouter = _generateNewRouter();

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_offRamp.setRouter(newRouter);
  }
}

/// @notice #ccipReceive
contract EVM2EVMGEOffRamp_ccipReceive is EVM2EVMGEOffRampSetup {
  // Reverts

  function testReverts() public {
    CCIP.Any2EVMMessageFromSender memory message = _convertGEToGeneralMessage(_generateAny2EVMGEMessageNoTokens(1));
    vm.expectRevert();
    s_offRamp.ccipReceive(message);
  }
}

/// @notice #execute
contract EVM2EVMGEOffRamp_execute is EVM2EVMGEOffRampSetup {
  Any2EVMOffRampRouterInterface s_router;

  function setUp() public virtual override {
    EVM2EVMGEOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new GERouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  // Success

  function testSingleMessageNoTokensSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    messages[0].nonce++;
    messages[0].sequenceNumber++;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSkippedIncorrectNonceSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();

    messages[0].nonce++;

    vm.expectEmit(false, false, false, true);
    emit SkippedIncorrectNonce(messages[0].nonce, messages[0].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSkippedIncorrectNonceStillExecutesSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateMessagesWithTokens();

    messages[1].nonce++;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit SkippedIncorrectNonce(messages[1].nonce, messages[1].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Asserts that a message execution fails, but it does
  // not disrupt the overall execution of the batch
  function testSingleMessageFailureSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();
    MaybeRevertMessageReceiver newReceiver = new MaybeRevertMessageReceiver(true);
    messages[0].receiver = address(newReceiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.FAILURE);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[1].sequenceNumber, messages[1].messageId, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensAndGESuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[1].messageId, CCIP.MessageExecutionState.SUCCESS);

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
    Any2EVMOffRampRouterInterface newRouter = GERouter(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(BaseOffRampInterface.RouterNotSet.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), false);
  }

  function testRootNotCommittedReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(CommitStoreInterface.verify.selector),
      abi.encode(0)
    );
    vm.expectRevert(BaseOffRampInterface.RootNotCommitted.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(CommitStoreInterface.verify.selector),
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
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();
    CCIP.EVMTokenAndAmount[] memory newTokens = new CCIP.EVMTokenAndAmount[](MAX_TOKENS_LENGTH + 1);
    messages[0].tokensAndAmounts = newTokens;
    CCIP.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.UnsupportedNumberOfTokens.selector, messages[0].sequenceNumber)
    );
    s_offRamp.execute(report, false);
  }

  function testMessageTooLargeReverts() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    CCIP.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(BaseOffRampInterface.MessageTooLarge.selector, MAX_DATA_SIZE, messages[0].data.length)
    );
    s_offRamp.execute(executionReport, false);
  }

  function testUnsupportedTokenReverts() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokensAndAmounts[0] = getCastedDestinationEVMTokenAndAmountsWithZeroAmounts()[0];
    messages[0].feeToken = messages[0].tokensAndAmounts[0].token;
    messages[0].feeTokenAmount = COMMIT_FEE_JUELS;
    vm.expectRevert(abi.encodeWithSelector(BaseOffRampInterface.UnsupportedToken.selector, s_destTokens[0]));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }
}

/// @notice #executeSingleMessage
contract EVM2EVMGEOffRamp_executeSingleMessage is EVM2EVMGEOffRampSetup {
  Any2EVMOffRampRouterInterface public s_router;

  function setUp() public virtual override {
    EVM2EVMGEOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new GERouter(offRamps);
    s_offRamp.setRouter(s_router);
    changePrank(address(s_offRamp));
  }

  // Success

  function testNoTokensSuccess() public {
    s_offRamp.executeSingleMessage(_convertGEToGeneralMessage(_generateAny2EVMGEMessageNoTokens(1)));
  }

  function testTokensSuccess() public {
    s_offRamp.executeSingleMessage(_convertGEToGeneralMessage(_generateMessagesWithTokens()[0]));
  }

  function testNonContractSuccess() public {
    CCIP.Any2EVMMessageFromSender memory message = _convertGEToGeneralMessage(_generateAny2EVMGEMessageNoTokens(1));
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message);
  }

  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testNonContractWithTokensSuccess() public {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    vm.expectEmit(true, true, false, true);
    emit Released(address(s_offRamp), STRANGER, amounts[0]);
    vm.expectEmit(true, true, false, true);
    emit Minted(address(s_offRamp), STRANGER, amounts[1]);
    CCIP.Any2EVMMessageFromSender memory message = _convertGEToGeneralMessage(
      _generateAny2EVMGEMessageWithTokens(1, amounts)
    );
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message);
  }

  // Reverts

  function testMessageSenderReverts() public {
    vm.stopPrank();
    CCIP.Any2EVMMessageFromSender memory message = _convertGEToGeneralMessage(_generateAny2EVMGEMessageNoTokens(1));
    vm.expectRevert(BaseOffRampInterface.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(message);
  }
}

/// @notice #_report
contract EVM2EVMGEOffRamp__report is EVM2EVMGEOffRampSetup {
  Any2EVMOffRampRouterInterface s_router;

  function setUp() public virtual override {
    EVM2EVMGEOffRampSetup.setUp();
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new GERouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  // Asserts that execute completes
  function testReportSuccess() public {
    CCIP.EVM2EVMGEMessage[] memory messages = _generateBasicMessages();
    CCIP.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectEmit(true, true, false, false);
    emit ExecutionStateChanged(messages[0].sequenceNumber, messages[0].messageId, CCIP.MessageExecutionState.SUCCESS);

    s_offRamp.report(abi.encode(report));
  }
}
