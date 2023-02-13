// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "../../health/HealthChecker.sol";
import "./EVM2EVMOffRampSetup.t.sol";
import "../../router/Router.sol";
import "../helpers/receivers/ConformingReceiver.sol";
import "../helpers/receivers/MaybeRevertMessageReceiverNo165.sol";
import "../helpers/receivers/ReentrancyAbuser.sol";
import "../fees/WETH9.sol";
import {IBaseOffRamp} from "../../interfaces/offRamp/IBaseOffRamp.sol";

/// @notice #constructor
contract EVM2EVMOffRamp_constructor is EVM2EVMOffRampSetup {
  // Success

  function testSuccess() public {
    // typeAndVersion
    assertEq("EVM2EVMOffRamp 1.0.0", s_offRamp.typeAndVersion());

    // owner
    assertEq(OWNER, s_offRamp.owner());

    // OffRamp config
    (uint64 source, uint64 dest) = s_offRamp.getChainIDs();
    assertEq(SOURCE_CHAIN_ID, source);
    assertEq(DEST_CHAIN_ID, dest);
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
    IERC20[] memory pools = s_offRamp.getSupportedTokens();
    assertEq(pools.length, s_sourceTokens.length);
    assertTrue(address(pools[0]) == address(s_sourceTokens[0]));
    assertTrue(address(pools[1]) == address(s_sourceTokens[1]));

    // HealthChecker
    assertEq(address(s_afn), address(s_offRamp.getAFN()));
  }
}

/// @notice #setRouter
contract EVM2EVMOffRamp_setRouter is EVM2EVMOffRampSetup {
  IRouter public s_router;

  event OffRampRouterSet(address indexed router, uint64 sourceChainId, address onRampAddress);

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    address[] memory offRamps = new address[](1);
    offRamps[0] = address(s_offRamp);
    WETH9 weth = new WETH9();
    s_router = new Router(offRamps, address(weth));
    s_offRamp.setRouter(s_router);
  }

  function _generateNewRouter() internal returns (Router newRouter) {
    address[] memory offRamps = new address[](0);
    WETH9 weth = new WETH9();
    newRouter = new Router(offRamps, address(weth));
  }

  // Success

  function testSuccessNewAddress() public {
    _testSuccess(_generateNewRouter());
  }

  function testSuccessZeroAddress() public {
    _testSuccess(IRouter(address(0)));
  }

  function _testSuccess(IRouter newRouter) private {
    vm.expectEmit(true, false, false, true);
    emit OffRampRouterSet(address(newRouter), SOURCE_CHAIN_ID, ON_RAMP_ADDRESS);

    s_offRamp.setRouter(newRouter);

    assertEq(address(newRouter), address(s_offRamp.getRouter()));
  }

  // Reverts

  function testOwnerReverts() public {
    IRouter newRouter = _generateNewRouter();

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_offRamp.setRouter(newRouter);
  }
}

/// @notice #ccipReceive
contract EVM2EVMOffRamp_ccipReceive is EVM2EVMOffRampSetup {
  // Reverts

  function testReverts() public {
    Common.Any2EVMMessage memory message = _convertToGeneralMessage(_generateAny2EVMMessageNoTokens(1));
    vm.expectRevert();
    s_offRamp.ccipReceive(message);
  }
}

/// @notice #execute
contract EVM2EVMOffRamp_execute is EVM2EVMOffRampSetup {
  IRouter s_router;

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    address[] memory offRamps = new address[](1);
    offRamps[0] = address(s_offRamp);
    WETH9 weth = new WETH9();
    s_router = new Router(offRamps, address(weth));
    s_offRamp.setRouter(s_router);
  }

  // Success

  function testSingleMessageNoTokensSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    messages[0].nonce++;
    messages[0].sequenceNumber++;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSkippedIncorrectNonceSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    messages[0].nonce++;

    vm.expectEmit(false, false, false, true);
    emit SkippedIncorrectNonce(messages[0].nonce, messages[0].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSkippedIncorrectNonceStillExecutesSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();

    messages[1].nonce++;

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit(false, false, false, true);
    emit SkippedIncorrectNonce(messages[1].nonce, messages[1].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Send a message to a contract that does not implement the CCIPReceiver interface
  // This should execute successfully.
  function testSingleMessageToNonCCIPReceiverSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    MaybeRevertMessageReceiverNo165 newReceiver = new MaybeRevertMessageReceiverNo165(true);
    messages[0].receiver = address(newReceiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensSuccess_gas() public {
    vm.pauseGasMetering();
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[1].sequenceNumber,
      messages[1].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.resumeGasMetering();
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testTwoMessagesWithTokensAndGESuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit(false, false, false, true);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[1].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

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
    IRouter newRouter = Router(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(IBaseOffRamp.RouterNotSet.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), false);
  }

  function testRootNotCommittedReverts() public {
    vm.mockCall(address(s_mockCommitStore), abi.encodeWithSelector(ICommitStore.verify.selector), abi.encode(0));
    vm.expectRevert(IBaseOffRamp.RootNotCommitted.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(ICommitStore.verify.selector),
      abi.encode(BLOCK_TIME)
    );
    vm.expectRevert(IBaseOffRamp.ManualExecutionNotYetEnabled.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testAlreadyExecutedReverts() public {
    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(_generateBasicMessages());
    s_offRamp.execute(executionReport, false);
    vm.expectRevert(abi.encodeWithSelector(IBaseOffRamp.AlreadyExecuted.selector, executionReport.sequenceNumbers[0]));
    s_offRamp.execute(executionReport, false);
  }

  function testInvalidSourceChainReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IBaseOffRamp.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Common.EVMTokenAndAmount[] memory newTokens = new Common.EVMTokenAndAmount[](MAX_TOKENS_LENGTH + 1);
    messages[0].tokensAndAmounts = newTokens;
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectRevert(
      abi.encodeWithSelector(IBaseOffRamp.UnsupportedNumberOfTokens.selector, messages[0].sequenceNumber)
    );
    s_offRamp.execute(report, false);
  }

  function testMessageTooLargeReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(IBaseOffRamp.MessageTooLarge.selector, MAX_DATA_SIZE, messages[0].data.length)
    );
    s_offRamp.execute(executionReport, false);
  }

  function testUnsupportedTokenReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokensAndAmounts[0] = getCastedDestinationEVMTokenAndAmountsWithZeroAmounts()[0];
    messages[0].feeToken = messages[0].tokensAndAmounts[0].token;
    vm.expectRevert(
      abi.encodeWithSelector(
        IBaseOffRamp.ExecutionError.selector,
        abi.encodeWithSelector(IBaseOffRamp.UnsupportedToken.selector, s_destTokens[0])
      )
    );
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }
}

/// @notice #executeSingleMessage
contract EVM2EVMOffRamp_executeSingleMessage is EVM2EVMOffRampSetup {
  IRouter public s_router;

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    address[] memory offRamps = new address[](1);
    offRamps[0] = address(s_offRamp);
    s_router = new Router(offRamps, address(1));
    s_offRamp.setRouter(s_router);
    changePrank(address(s_offRamp));
  }

  // Success

  function testNoTokensSuccess() public {
    s_offRamp.executeSingleMessage(_generateAny2EVMMessageNoTokens(1), false);
  }

  function testTokensSuccess() public {
    s_offRamp.executeSingleMessage(_generateMessagesWithTokens()[0], false);
  }

  function testNonContractSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message, false);
  }

  event MessageReceived();

  function testLowGasLimitManualExecutionSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    message.gasLimit = 1;
    message.receiver = address(new ConformingReceiver(address(s_router), s_destFeeToken));
    vm.expectRevert(IBaseOffRamp.ReceiverError.selector);
    s_offRamp.executeSingleMessage(message, false);
    vm.expectEmit(false, false, false, false);
    emit MessageReceived();
    s_offRamp.executeSingleMessage(message, true);
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
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageWithTokens(1, amounts);
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message, false);
  }

  // Reverts

  function testMessageSenderReverts() public {
    vm.stopPrank();
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    vm.expectRevert(IBaseOffRamp.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(message, false);
  }
}

/// @notice #_report
contract EVM2EVMOffRamp__report is EVM2EVMOffRampSetup {
  IRouter s_router;

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    address[] memory offRamps = new address[](1);
    offRamps[0] = address(s_offRamp);
    s_router = new Router(offRamps, address(1));
    s_offRamp.setRouter(s_router);
  }

  // Asserts that execute completes
  function testReportSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectEmit(true, true, false, false);
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );
    s_offRamp.report(abi.encode(report));
  }
}

/// @notice #manuallyExecute
contract EVM2EVMOffRamp_manuallyExecute is EVM2EVMOffRampSetup {
  event ReentrancySucceeded();

  IRouter s_router;

  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
    address[] memory offRamps = new address[](1);
    offRamps[0] = address(s_offRamp);
    s_router = new Router(offRamps, address(1));
    s_offRamp.setRouter(s_router);
  }

  function testReentrancyManualExecuteFAILS() public {
    uint256 tokenAmount = 1e9;
    IERC20 tokenToAbuse = IERC20(s_destFeeToken);

    // This needs to be deployed before the source chain message is sent
    // because we need the address for the receiver.
    ReentrancyAbuser receiver = new ReentrancyAbuser(address(s_router), s_offRamp);
    uint256 balancePre = tokenToAbuse.balanceOf(address(receiver));

    // For this test any message will be flagged as correct by the
    // commitStore. In a real scenario the abuser would have to actually
    // send the message that they want to replay.
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0].tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    messages[0].tokensAndAmounts[0] = Common.EVMTokenAndAmount({token: s_sourceFeeToken, amount: tokenAmount});
    messages[0].receiver = address(receiver);

    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    // sets the report to be repeated on the ReentrancyAbuser to be able to replay
    receiver.setPayload(report);

    s_offRamp.manuallyExecute(report);

    // The first entry should be fine and triggers the second entry. This one fails
    // but since it's an inner tx of the first one it is caught in the cry-catch.
    // This failure of the inner tx flags the outer tx as failed.
    assertEq(
      uint256(s_offRamp.getExecutionState(messages[0].sequenceNumber)),
      uint256(Internal.MessageExecutionState.FAILURE)
    );

    // Since the tx failed we don't release the tokens
    assertEq(tokenToAbuse.balanceOf(address(receiver)), balancePre);
  }
}
