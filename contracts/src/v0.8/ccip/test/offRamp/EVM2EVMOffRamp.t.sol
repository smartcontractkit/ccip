// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "../../health/HealthChecker.sol";
import "./EVM2EVMOffRampSetup.t.sol";
import "../../router/Router.sol";
import "../helpers/receivers/ConformingReceiver.sol";
import "../helpers/receivers/MaybeRevertMessageReceiverNo165.sol";
import "../helpers/receivers/ReentrancyAbuser.sol";
import "../../interfaces/offRamp/IEVM2EVMOffRamp.sol";

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

  // Revert
  function testTokenConfigMismatchReverts() public {
    vm.expectRevert(OffRampTokenPoolRegistry.InvalidTokenPoolConfig.selector);

    IPool[] memory pools = new IPool[](1);

    IERC20[] memory wrongTokens = new IERC20[](5);
    IEVM2EVMOffRamp.OffRampConfig memory offRampConfig = offRampConfig(s_feeManager, s_mockCommitStore, s_destRouter);
    s_offRamp = new EVM2EVMOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      ON_RAMP_ADDRESS,
      offRampConfig,
      s_afn,
      wrongTokens,
      pools,
      rateLimiterConfig()
    );
  }

  function testZeroOnRampAddressReverts() public {
    IPool[] memory pools = new IPool[](2);
    pools[0] = IPool(s_sourcePools[0]);
    pools[1] = new LockReleaseTokenPool(IERC20(s_sourceTokens[1]));

    vm.expectRevert(IEVM2EVMOffRamp.ZeroAddressNotAllowed.selector);

    IAggregateRateLimiter.RateLimiterConfig memory rateLimiterConfig = IAggregateRateLimiter.RateLimiterConfig({
      rate: 1e20,
      capacity: 1e20,
      admin: TOKEN_LIMIT_ADMIN
    });

    IEVM2EVMOffRamp.OffRampConfig memory offRampConfig = offRampConfig(s_feeManager, s_mockCommitStore, s_destRouter);
    s_offRamp = new EVM2EVMOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      ZERO_ADDRESS,
      offRampConfig,
      s_afn,
      getCastedSourceTokens(),
      pools,
      rateLimiterConfig
    );
  }

  function assertSameConfig(IEVM2EVMOffRamp.OffRampConfig memory a, IEVM2EVMOffRamp.OffRampConfig memory b) public {
    assertEq(address(a.feeManager), address(b.feeManager));
    assertEq(a.executionDelaySeconds, b.executionDelaySeconds);
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.permissionLessExecutionThresholdSeconds, b.permissionLessExecutionThresholdSeconds);
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

  function testRootNotCommittedReverts() public {
    vm.mockCall(address(s_mockCommitStore), abi.encodeWithSelector(ICommitStore.verify.selector), abi.encode(0));
    vm.expectRevert(IEVM2EVMOffRamp.RootNotCommitted.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(ICommitStore.verify.selector),
      abi.encode(BLOCK_TIME)
    );
    vm.expectRevert(IEVM2EVMOffRamp.ManualExecutionNotYetEnabled.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testAlreadyExecutedReverts() public {
    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(_generateBasicMessages());
    s_offRamp.execute(executionReport, false);
    vm.expectRevert(
      abi.encodeWithSelector(IEVM2EVMOffRamp.AlreadyExecuted.selector, executionReport.sequenceNumbers[0])
    );
    s_offRamp.execute(executionReport, false);
  }

  function testInvalidSourceChainReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOffRamp.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Common.EVMTokenAndAmount[] memory newTokens = new Common.EVMTokenAndAmount[](MAX_TOKENS_LENGTH + 1);
    messages[0].tokensAndAmounts = newTokens;
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectRevert(
      abi.encodeWithSelector(IEVM2EVMOffRamp.UnsupportedNumberOfTokens.selector, messages[0].sequenceNumber)
    );
    s_offRamp.execute(report, false);
  }

  function testMessageTooLargeReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(IEVM2EVMOffRamp.MessageTooLarge.selector, MAX_DATA_SIZE, messages[0].data.length)
    );
    s_offRamp.execute(executionReport, false);
  }

  function testUnsupportedTokenReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokensAndAmounts[0] = getCastedDestinationEVMTokenAndAmountsWithZeroAmounts()[0];
    messages[0].feeToken = messages[0].tokensAndAmounts[0].token;
    vm.expectRevert(
      abi.encodeWithSelector(
        IEVM2EVMOffRamp.ExecutionError.selector,
        abi.encodeWithSelector(IEVM2EVMOffRamp.UnsupportedToken.selector, s_destTokens[0])
      )
    );
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }
}

/// @notice #executeSingleMessage
contract EVM2EVMOffRamp_executeSingleMessage is EVM2EVMOffRampSetup {
  function setUp() public virtual override {
    EVM2EVMOffRampSetup.setUp();
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
    message.receiver = address(new ConformingReceiver(address(s_destRouter), s_destFeeToken));
    vm.expectRevert(IEVM2EVMOffRamp.ReceiverError.selector);
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
    vm.expectRevert(IEVM2EVMOffRamp.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(message, false);
  }
}

/// @notice #_report
contract EVM2EVMOffRamp__report is EVM2EVMOffRampSetup {
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

  function testReentrancyManualExecuteFAILS() public {
    uint256 tokenAmount = 1e9;
    IERC20 tokenToAbuse = IERC20(s_destFeeToken);

    // This needs to be deployed before the source chain message is sent
    // because we need the address for the receiver.
    ReentrancyAbuser receiver = new ReentrancyAbuser(address(s_destRouter), s_offRamp);
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

/// @notice #getExecutionState
contract EVM2EVMOffRamp_getExecutionState is EVM2EVMOffRampSetup {
  // Success
  function testSuccess() public {
    // setting the execution state is done with a helper function. This
    // is normally not exposed.
    s_offRamp.setExecutionState(1, Internal.MessageExecutionState.FAILURE);
    s_offRamp.setExecutionState(10, Internal.MessageExecutionState.IN_PROGRESS);
    s_offRamp.setExecutionState(33, Internal.MessageExecutionState.UNTOUCHED);
    s_offRamp.setExecutionState(50, Internal.MessageExecutionState.SUCCESS);

    assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(1)));
    assertEq(uint256(Internal.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(10)));
    assertEq(uint256(Internal.MessageExecutionState.UNTOUCHED), uint256(s_offRamp.getExecutionState(33)));
    assertEq(uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(50)));
  }
}

/// @notice #_releaseOrMintToken internal function
contract EVM2EVMOffRamp__releaseOrMintToken is EVM2EVMOffRampSetup {
  // Success
  function testSuccess() public {
    IERC20 destToken0 = IERC20(s_destTokens[0]);
    uint256 startingBalance = destToken0.balanceOf(OWNER);
    uint256 amount = POOL_BALANCE / 2;
    s_offRamp.releaseOrMintToken(IPool(s_destPools[0]), amount, OWNER);
    assertEq(startingBalance + amount, destToken0.balanceOf(OWNER));
  }

  // Success on BurnMintTokenPool
  function testMintSuccess() public {
    IERC20 destToken1 = IERC20(s_destTokens[1]);
    uint256 startingBalance = destToken1.balanceOf(OWNER);
    uint256 amount = POOL_BALANCE * 2; // amount bigger than balance
    uint256 startingPoolBalance = destToken1.balanceOf(s_destPools[1]);
    s_offRamp.releaseOrMintToken(IPool(s_destPools[1]), amount, OWNER);
    assertEq(startingBalance + amount, destToken1.balanceOf(OWNER));
    // pool balance doesn't change, because tokens were minted
    assertEq(startingPoolBalance, destToken1.balanceOf(s_destPools[1]));
  }

  // Revert
  function testExceedsPoolReverts() public {
    vm.expectRevert("ERC20: transfer amount exceeds balance");
    s_offRamp.releaseOrMintToken(IPool(s_destPools[0]), POOL_BALANCE * 2, OWNER);
  }
}

/// @notice #_releaseOrMintTokens
contract EVM2EVMOffRamp__releaseOrMintTokens is EVM2EVMOffRampSetup {
  // Success
  function testSuccess() public {
    Common.EVMTokenAndAmount[] memory srcTokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    IERC20 dstToken1 = IERC20(s_destTokens[0]);
    uint256 startingBalance = dstToken1.balanceOf(OWNER);
    uint256 amount1 = 100;
    srcTokensAndAmounts[0].amount = 100;

    s_offRamp.releaseOrMintTokens(srcTokensAndAmounts, OWNER);
    assertEq(startingBalance + amount1, dstToken1.balanceOf(OWNER));
  }

  // Revert

  function testUnsupportedTokenReverts() public {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](1);

    vm.expectRevert(abi.encodeWithSelector(IEVM2EVMOffRamp.UnsupportedToken.selector, address(0)));
    s_offRamp.releaseOrMintTokens(tokensAndAmounts, OWNER);
  }
}
