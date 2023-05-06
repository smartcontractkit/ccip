// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/receivers/MaybeRevertMessageReceiver.sol";
import "./EVM2EVMOffRampSetup.t.sol";
import "../../Router.sol";
import "../helpers/receivers/ConformingReceiver.sol";
import "../helpers/receivers/MaybeRevertMessageReceiverNo165.sol";
import "../helpers/receivers/ReentrancyAbuser.sol";
import {AFN} from "../../AFN.sol";
import "../../offRamp/EVM2EVMOffRamp.sol";
import "../mocks/MockCommitStore.sol";
import "../ocr/OCR2Base.t.sol";

/// @notice #constructor
contract EVM2EVMOffRamp_constructor is EVM2EVMOffRampSetup {
  event ConfigSet(EVM2EVMOffRamp.StaticConfig staticConfig, EVM2EVMOffRamp.DynamicConfig dynamicConfig);
  event PoolAdded(address token, address pool);

  function testConstructorSuccess() public {
    EVM2EVMOffRamp.StaticConfig memory staticConfig = EVM2EVMOffRamp.StaticConfig({
      commitStore: address(s_mockCommitStore),
      chainId: DEST_CHAIN_ID,
      sourceChainId: SOURCE_CHAIN_ID,
      onRamp: ON_RAMP_ADDRESS
    });
    EVM2EVMOffRamp.DynamicConfig memory dynamicConfig = generateDynamicOffRampConfig(
      address(s_destRouter),
      address(s_priceRegistry),
      address(s_mockAFN)
    );
    IERC20[] memory sourceTokens = getCastedSourceTokens();
    IPool[] memory castedPools = getCastedDestinationPools();

    vm.expectEmit();
    emit PoolAdded(address(sourceTokens[0]), address(castedPools[0]));

    s_offRamp = new EVM2EVMOffRampHelper(staticConfig, sourceTokens, castedPools, rateLimiterConfig());

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    s_offRamp.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );

    // Static config
    EVM2EVMOffRamp.StaticConfig memory gotStaticConfig = s_offRamp.getStaticConfig();
    assertEq(staticConfig.commitStore, gotStaticConfig.commitStore);
    assertEq(staticConfig.sourceChainId, gotStaticConfig.sourceChainId);
    assertEq(staticConfig.chainId, gotStaticConfig.chainId);
    assertEq(staticConfig.onRamp, gotStaticConfig.onRamp);

    // Dynamic config
    EVM2EVMOffRamp.DynamicConfig memory gotDynamicConfig = s_offRamp.getDynamicConfig();
    _assertSameConfig(dynamicConfig, gotDynamicConfig);

    // Pools & tokens
    IERC20[] memory pools = s_offRamp.getSupportedTokens();
    assertEq(pools.length, s_sourceTokens.length);
    assertTrue(address(pools[0]) == address(s_sourceTokens[0]));
    assertTrue(address(pools[1]) == address(s_sourceTokens[1]));
    assertEq(address(s_offRamp.getPoolByDestToken(IERC20(s_destTokens[0]))), address(s_destPools[0]));

    (uint32 configCount, uint32 blockNumber, bytes32 configDigest) = s_offRamp.latestConfigDetails();
    assertEq(1, configCount);
    assertEq(block.number, blockNumber);
    assertEq(0x0001285aaf79ffad0ffc300acb5d2eb12258ae419b5691dda1fdf7ef0343541f, configDigest);

    // OffRamp initial values
    assertEq("EVM2EVMOffRamp 1.0.0", s_offRamp.typeAndVersion());
    assertEq(OWNER, s_offRamp.owner());
  }

  // Revert
  function testTokenConfigMismatchReverts() public {
    vm.expectRevert(EVM2EVMOffRamp.InvalidTokenPoolConfig.selector);

    IPool[] memory pools = new IPool[](1);

    IERC20[] memory wrongTokens = new IERC20[](5);
    s_offRamp = new EVM2EVMOffRampHelper(
      EVM2EVMOffRamp.StaticConfig({
        commitStore: address(s_mockCommitStore),
        chainId: DEST_CHAIN_ID,
        sourceChainId: SOURCE_CHAIN_ID,
        onRamp: ON_RAMP_ADDRESS
      }),
      wrongTokens,
      pools,
      rateLimiterConfig()
    );
  }

  function testZeroOnRampAddressReverts() public {
    IPool[] memory pools = new IPool[](2);
    pools[0] = IPool(s_sourcePools[0]);
    pools[1] = new LockReleaseTokenPool(IERC20(s_sourceTokens[1]), rateLimiterConfig());

    vm.expectRevert(EVM2EVMOffRamp.ZeroAddressNotAllowed.selector);

    RateLimiter.Config memory rateLimiterConfig = RateLimiter.Config({isEnabled: true, rate: 1e20, capacity: 1e20});

    s_offRamp = new EVM2EVMOffRampHelper(
      EVM2EVMOffRamp.StaticConfig({
        commitStore: address(s_mockCommitStore),
        chainId: DEST_CHAIN_ID,
        sourceChainId: SOURCE_CHAIN_ID,
        onRamp: ZERO_ADDRESS
      }),
      getCastedSourceTokens(),
      pools,
      rateLimiterConfig
    );
  }

  function testCommitStoreAlreadyInUseReverts() public {
    s_mockCommitStore.setExpectedNextSequenceNumber(2);

    vm.expectRevert(EVM2EVMOffRamp.CommitStoreAlreadyInUse.selector);

    s_offRamp = new EVM2EVMOffRampHelper(
      EVM2EVMOffRamp.StaticConfig({
        commitStore: address(s_mockCommitStore),
        chainId: DEST_CHAIN_ID,
        sourceChainId: SOURCE_CHAIN_ID,
        onRamp: ON_RAMP_ADDRESS
      }),
      getCastedSourceTokens(),
      getCastedDestinationPools(),
      rateLimiterConfig()
    );
  }
}

contract EVM2EVMOffRamp_setDynamicConfig is EVM2EVMOffRampSetup {
  // OffRamp event
  event ConfigSet(EVM2EVMOffRamp.StaticConfig staticConfig, EVM2EVMOffRamp.DynamicConfig dynamicConfig);

  function testSetDynamicConfigSuccess() public {
    EVM2EVMOffRamp.StaticConfig memory staticConfig = s_offRamp.getStaticConfig();
    EVM2EVMOffRamp.DynamicConfig memory dynamicConfig = generateDynamicOffRampConfig(
      USER_3,
      address(s_priceRegistry),
      address(s_mockAFN)
    );
    bytes memory onchainConfig = abi.encode(dynamicConfig);

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    vm.expectEmit();
    uint32 configCount = 1;
    emit ConfigSet(
      uint32(block.number),
      getBasicConfigDigest(address(s_offRamp), s_f, configCount, onchainConfig),
      configCount + 1,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    s_offRamp.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    EVM2EVMOffRamp.DynamicConfig memory newConfig = s_offRamp.getDynamicConfig();
    _assertSameConfig(dynamicConfig, newConfig);
  }

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    EVM2EVMOffRamp.DynamicConfig memory dynamicConfig = generateDynamicOffRampConfig(
      USER_3,
      address(s_priceRegistry),
      address(1)
    );

    vm.expectRevert("Only callable by owner");

    s_offRamp.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );
  }

  function testRouterZeroAddressReverts() public {
    EVM2EVMOffRamp.DynamicConfig memory dynamicConfig = generateDynamicOffRampConfig(
      ZERO_ADDRESS,
      ZERO_ADDRESS,
      address(1)
    );

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOffRamp.InvalidOffRampConfig.selector, dynamicConfig));

    s_offRamp.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );
  }
}

contract EVM2EVMOffRamp_metadataHash is EVM2EVMOffRampSetup {
  function testMetadataHashSuccess() public {
    bytes32 h = s_offRamp.metadataHash();
    assertEq(
      h,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );
  }
}

/// @notice #ccipReceive
contract EVM2EVMOffRamp_ccipReceive is EVM2EVMOffRampSetup {
  // Reverts

  function testReverts() public {
    Client.Any2EVMMessage memory message = _convertToGeneralMessage(_generateAny2EVMMessageNoTokens(1));
    vm.expectRevert();
    s_offRamp.ccipReceive(message);
  }
}

/// @notice #execute
contract EVM2EVMOffRamp_execute is EVM2EVMOffRampSetup {
  function testSingleMessageNoTokensSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    s_offRamp.execute(_generateReportFromMessages(messages), false);

    messages[0].nonce++;
    messages[0].sequenceNumber++;

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    uint64 nonceBefore = s_offRamp.getSenderNonce(messages[0].sender);
    s_offRamp.execute(_generateReportFromMessages(messages), false);
    assertGt(s_offRamp.getSenderNonce(messages[0].sender), nonceBefore);
  }

  function testStrictFailureSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    messages[0].strict = true;
    messages[0].receiver = address(s_reverting_receiver);

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.FAILURE
    );
    // Nonce should not increment on a strict revert.
    assertEq(uint64(0), s_offRamp.getSenderNonce(address(OWNER)));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
    assertEq(uint64(0), s_offRamp.getSenderNonce(address(OWNER)));
  }

  function testStrictUntouchedToSuccessSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    messages[0].strict = true;
    messages[0].receiver = address(s_receiver);

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );
    // Nonce should increment on a strict untouched -> success.
    assertEq(uint64(0), s_offRamp.getSenderNonce(address(OWNER)));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
    assertEq(uint64(1), s_offRamp.getSenderNonce(address(OWNER)));
  }

  function testSkippedIncorrectNonceSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    messages[0].nonce++;

    vm.expectEmit();
    emit SkippedIncorrectNonce(messages[0].nonce, messages[0].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSkippedIncorrectNonceStillExecutesSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();

    messages[1].nonce++;

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit();
    emit SkippedIncorrectNonce(messages[1].nonce, messages[1].sender);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  // Send a message to a contract that does not implement the CCIPReceiver interface
  // This should execute successfully.
  function testSingleMessageToNonCCIPReceiverSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    MaybeRevertMessageReceiverNo165 newReceiver = new MaybeRevertMessageReceiverNo165(true);
    messages[0].receiver = address(newReceiver);

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testSingleMessagesNoTokensSuccess_gas() public {
    vm.pauseGasMetering();
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.resumeGasMetering();
    s_offRamp.execute(report, false);
  }

  function testTwoMessagesWithTokensSuccess_gas() public {
    vm.pauseGasMetering();
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[1].sequenceNumber,
      messages[1].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.resumeGasMetering();
    s_offRamp.execute(report, false);
  }

  function testTwoMessagesWithTokensAndGESuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[0].sequenceNumber,
      messages[0].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    vm.expectEmit();
    emit ExecutionStateChanged(
      messages[1].sequenceNumber,
      messages[1].messageId,
      Internal.MessageExecutionState.SUCCESS
    );

    assertEq(uint64(0), s_offRamp.getSenderNonce(OWNER));
    s_offRamp.execute(_generateReportFromMessages(messages), true);
    assertEq(uint64(2), s_offRamp.getSenderNonce(OWNER));
  }

  // Reverts

  function testPausedReverts() public {
    s_mockCommitStore.pause();
    vm.expectRevert("Pausable: paused");
    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
  }

  function testUnhealthyReverts() public {
    s_mockAFN.voteToCurse(0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
    vm.expectRevert(EVM2EVMOffRamp.BadAFNSignal.selector);
    s_offRamp.execute(_generateReportFromMessages(_generateMessagesWithTokens()), true);
  }

  function testUnexpectedTokenDataReverts() public {
    Internal.ExecutionReport memory report = _generateReportFromMessages(_generateBasicMessages());
    report.offchainTokenData = new bytes[][](report.encodedMessages.length + 1);

    vm.expectRevert(EVM2EVMOffRamp.UnexpectedTokenData.selector);

    s_offRamp.execute(report, false);
  }

  function testEmptyReportReverts() public {
    vm.expectRevert(EVM2EVMOffRamp.EmptyReport.selector);
    s_offRamp.execute(
      Internal.ExecutionReport({
        sequenceNumbers: new uint64[](0),
        proofs: new bytes32[](0),
        proofFlagBits: 2**256 - 1,
        encodedMessages: new bytes[](0),
        offchainTokenData: new bytes[][](0)
      }),
      true
    );
  }

  function testRootNotCommittedReverts() public {
    vm.mockCall(address(s_mockCommitStore), abi.encodeWithSelector(ICommitStore.verify.selector), abi.encode(0));
    vm.expectRevert(EVM2EVMOffRamp.RootNotCommitted.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testManualExecutionNotYetEnabledReverts() public {
    vm.mockCall(
      address(s_mockCommitStore),
      abi.encodeWithSelector(ICommitStore.verify.selector),
      abi.encode(BLOCK_TIME)
    );
    vm.expectRevert(EVM2EVMOffRamp.ManualExecutionNotYetEnabled.selector);

    s_offRamp.execute(_generateReportFromMessages(_generateBasicMessages()), true);
    vm.clearMockedCalls();
  }

  function testAlreadyExecutedReverts() public {
    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(_generateBasicMessages());
    s_offRamp.execute(executionReport, false);
    vm.expectRevert(
      abi.encodeWithSelector(EVM2EVMOffRamp.AlreadyExecuted.selector, executionReport.sequenceNumbers[0])
    );
    s_offRamp.execute(executionReport, false);
  }

  function testInvalidSourceChainReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0].sourceChainId = SOURCE_CHAIN_ID + 1;

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOffRamp.InvalidSourceChain.selector, SOURCE_CHAIN_ID + 1));
    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function testUnsupportedNumberOfTokensReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Client.EVMTokenAmount[] memory newTokens = new Client.EVMTokenAmount[](MAX_TOKENS_LENGTH + 1);
    messages[0].tokenAmounts = newTokens;
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectRevert(
      abi.encodeWithSelector(EVM2EVMOffRamp.UnsupportedNumberOfTokens.selector, messages[0].sequenceNumber)
    );
    s_offRamp.execute(report, false);
  }

  function testTokenDataMismatchReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    report.offchainTokenData[0] = new bytes[](messages[0].tokenAmounts.length + 1);

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOffRamp.TokenDataMismatch.selector, messages[0].sequenceNumber));
    s_offRamp.execute(report, false);
  }

  function testMessageTooLargeReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    Internal.ExecutionReport memory executionReport = _generateReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(EVM2EVMOffRamp.MessageTooLarge.selector, MAX_DATA_SIZE, messages[0].data.length)
    );
    s_offRamp.execute(executionReport, false);
  }

  function testUnsupportedTokenReverts() public {
    Internal.EVM2EVMMessage[] memory messages = _generateMessagesWithTokens();
    messages[0].tokenAmounts[0] = getCastedDestinationEVMTokenAmountsWithZeroAmounts()[0];
    messages[0].feeToken = messages[0].tokenAmounts[0].token;
    vm.expectRevert(
      abi.encodeWithSelector(
        EVM2EVMOffRamp.ExecutionError.selector,
        abi.encodeWithSelector(EVM2EVMOffRamp.UnsupportedToken.selector, s_destTokens[0])
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

  function testNoTokensSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
  }

  function testTokensSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateMessagesWithTokens()[0];
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
  }

  function testNonContractSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
  }

  event MessageReceived();

  function testLowGasLimitManualExecutionSuccess() public {
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    message.gasLimit = 1;
    message.receiver = address(new ConformingReceiver(address(s_destRouter), s_destFeeToken));
    vm.expectRevert(EVM2EVMOffRamp.ReceiverError.selector);
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
    vm.expectEmit();
    emit MessageReceived();
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), true);
  }

  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function testNonContractWithTokensSuccess() public {
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    vm.expectEmit();
    emit Released(address(s_offRamp), STRANGER, amounts[0]);
    vm.expectEmit();
    emit Minted(address(s_offRamp), STRANGER, amounts[1]);
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageWithTokens(1, amounts);
    message.receiver = STRANGER;
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
  }

  // Reverts

  function testMessageSenderReverts() public {
    vm.stopPrank();
    Internal.EVM2EVMMessage memory message = _generateAny2EVMMessageNoTokens(1);
    vm.expectRevert(EVM2EVMOffRamp.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(message, new bytes[](message.tokenAmounts.length), false);
  }
}

/// @notice #_report
contract EVM2EVMOffRamp__report is EVM2EVMOffRampSetup {
  // Asserts that execute completes
  function testReportSuccess() public {
    Internal.EVM2EVMMessage[] memory messages = _generateBasicMessages();
    Internal.ExecutionReport memory report = _generateReportFromMessages(messages);

    vm.expectEmit();
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
    messages[0].tokenAmounts = new Client.EVMTokenAmount[](1);
    messages[0].tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceFeeToken, amount: tokenAmount});
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
  mapping(uint64 => Internal.MessageExecutionState) s_differentialExecutionState;

  function testDifferentialSuccess(uint16[500] memory seqNums, uint8[500] memory values) public {
    for (uint256 i = 0; i < seqNums.length; ++i) {
      // Only use the first three slots. This makes sure existing slots get overwritten
      // as the tests uses 500 sequence numbers.
      uint16 seqNum = seqNums[i] % 386;
      Internal.MessageExecutionState state = Internal.MessageExecutionState(values[i] % 4);
      s_differentialExecutionState[seqNum] = state;
      s_offRamp.setExecutionStateHelper(seqNum, state);
      assertEq(uint256(state), uint256(s_offRamp.getExecutionState(seqNum)));
    }

    for (uint256 i = 0; i < seqNums.length; ++i) {
      uint16 seqNum = seqNums[i] % 386;
      Internal.MessageExecutionState expectedState = s_differentialExecutionState[seqNum];
      assertEq(uint256(expectedState), uint256(s_offRamp.getExecutionState(seqNum)));
    }
  }

  function test_GetExecutionStateSuccess() public {
    s_offRamp.setExecutionStateHelper(0, Internal.MessageExecutionState.FAILURE);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3);

    s_offRamp.setExecutionStateHelper(1, Internal.MessageExecutionState.FAILURE);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3 + (3 << 2));

    s_offRamp.setExecutionStateHelper(1, Internal.MessageExecutionState.IN_PROGRESS);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3 + (1 << 2));

    s_offRamp.setExecutionStateHelper(2, Internal.MessageExecutionState.FAILURE);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3 + (1 << 2) + (3 << 4));

    s_offRamp.setExecutionStateHelper(127, Internal.MessageExecutionState.IN_PROGRESS);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3 + (1 << 2) + (3 << 4) + (1 << 254));

    s_offRamp.setExecutionStateHelper(128, Internal.MessageExecutionState.SUCCESS);
    assertEq(s_offRamp.getExecutionStateBitMap(0), 3 + (1 << 2) + (3 << 4) + (1 << 254));
    assertEq(s_offRamp.getExecutionStateBitMap(1), 2);

    assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(0)));
    assertEq(uint256(Internal.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(1)));
    assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(2)));
    assertEq(uint256(Internal.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(127)));
    assertEq(uint256(Internal.MessageExecutionState.SUCCESS), uint256(s_offRamp.getExecutionState(128)));
  }

  function testFillExecutionStateSuccess() public {
    for (uint64 i = 0; i < 384; ++i) {
      s_offRamp.setExecutionStateHelper(i, Internal.MessageExecutionState.FAILURE);
    }

    for (uint64 i = 0; i < 384; ++i) {
      assertEq(uint256(Internal.MessageExecutionState.FAILURE), uint256(s_offRamp.getExecutionState(i)));
    }

    for (uint64 i = 0; i < 3; ++i) {
      assertEq(type(uint256).max, s_offRamp.getExecutionStateBitMap(i));
    }

    for (uint64 i = 0; i < 384; ++i) {
      s_offRamp.setExecutionStateHelper(i, Internal.MessageExecutionState.IN_PROGRESS);
    }

    for (uint64 i = 0; i < 384; ++i) {
      assertEq(uint256(Internal.MessageExecutionState.IN_PROGRESS), uint256(s_offRamp.getExecutionState(i)));
    }

    for (uint64 i = 0; i < 3; ++i) {
      // 0x555... == 0b101010101010.....
      assertEq(
        0x5555555555555555555555555555555555555555555555555555555555555555,
        s_offRamp.getExecutionStateBitMap(i)
      );
    }
  }
}

/// @notice #_releaseOrMintTokens
contract EVM2EVMOffRamp__releaseOrMintTokens is EVM2EVMOffRampSetup {
  function test_releaseOrMintTokensSuccess() public {
    Client.EVMTokenAmount[] memory srcTokenAmounts = getCastedSourceEVMTokenAmountsWithZeroAmounts();
    IERC20 dstToken1 = IERC20(s_destTokens[0]);
    uint256 startingBalance = dstToken1.balanceOf(OWNER);
    uint256 amount1 = 100;
    srcTokenAmounts[0].amount = 100;

    s_offRamp.releaseOrMintTokens(srcTokenAmounts, abi.encode(OWNER), OWNER, new bytes[](srcTokenAmounts.length));
    assertEq(startingBalance + amount1, dstToken1.balanceOf(OWNER));
  }

  // Revert

  function testUnsupportedTokenReverts() public {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOffRamp.UnsupportedToken.selector, address(0)));
    s_offRamp.releaseOrMintTokens(tokenAmounts, bytes(""), OWNER, new bytes[](0));
  }
}

contract EVM2EVMOffRamp_applyPoolUpdates is EVM2EVMOffRampSetup {
  event PoolAdded(address token, address pool);
  event PoolRemoved(address token, address pool);

  function testApplyPoolUpdatesSuccess() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1)), rateLimiterConfig()))
    });

    vm.expectEmit();
    emit PoolAdded(adds[0].token, adds[0].pool);

    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);

    assertEq(adds[0].pool, address(s_offRamp.getPoolBySourceToken(IERC20(adds[0].token))));

    vm.expectEmit();
    emit PoolRemoved(adds[0].token, adds[0].pool);

    s_offRamp.applyPoolUpdates(adds, new Internal.PoolUpdate[](0));

    vm.expectRevert(abi.encodeWithSelector(EVM2EVMOffRamp.UnsupportedToken.selector, adds[0].token));
    s_offRamp.getPoolBySourceToken(IERC20(adds[0].token));
  }

  // Reverts
  function testOnlyCallableByOwnerReverts() public {
    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");

    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), new Internal.PoolUpdate[](0));
  }

  function testPoolAlreadyExistsReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](2);
    adds[0] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1)), rateLimiterConfig()))
    });
    adds[1] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1)), rateLimiterConfig()))
    });

    vm.expectRevert(EVM2EVMOffRamp.PoolAlreadyAdded.selector);

    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);
  }

  function testInvalidTokenPoolConfigReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({token: address(0), pool: address(2)});

    vm.expectRevert(EVM2EVMOffRamp.InvalidTokenPoolConfig.selector);

    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);

    adds[0] = Internal.PoolUpdate({token: address(1), pool: address(0)});

    vm.expectRevert(EVM2EVMOffRamp.InvalidTokenPoolConfig.selector);

    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);
  }

  function testPoolDoesNotExistReverts() public {
    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1)), rateLimiterConfig()))
    });

    vm.expectRevert(EVM2EVMOffRamp.PoolDoesNotExist.selector);

    s_offRamp.applyPoolUpdates(removes, new Internal.PoolUpdate[](0));
  }

  function testTokenPoolMismatchReverts() public {
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1)), rateLimiterConfig()))
    });
    s_offRamp.applyPoolUpdates(new Internal.PoolUpdate[](0), adds);

    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0] = Internal.PoolUpdate({
      token: address(1),
      pool: address(new LockReleaseTokenPool(IERC20(address(1000)), rateLimiterConfig()))
    });

    vm.expectRevert(EVM2EVMOffRamp.TokenPoolMismatch.selector);

    s_offRamp.applyPoolUpdates(removes, adds);
  }
}

contract EVM2EVMOffRamp_getDestinationToken is EVM2EVMOffRampSetup {
  function testGetDestinationTokenSuccess() public {
    address expectedToken = address(IPool(s_destPools[0]).getToken());
    address actualToken = address(s_offRamp.getDestinationToken(IERC20(s_sourceTokens[0])));

    assertEq(expectedToken, actualToken);

    expectedToken = address(IPool(s_destPools[1]).getToken());
    actualToken = address(s_offRamp.getDestinationToken(IERC20(s_sourceTokens[1])));

    assertEq(expectedToken, actualToken);
  }
}

contract EVM2EVMOffRamp_getDestinationTokens is EVM2EVMOffRampSetup {
  function testGetDestinationTokensSuccess() public {
    IERC20[] memory actualTokens = s_offRamp.getDestinationTokens();

    for (uint256 i = 0; i < actualTokens.length; ++i) {
      assertEq(address(s_destTokens[i]), address(actualTokens[i]));
    }
  }
}

contract EVM2EVMOffRamp_afn is EVM2EVMOffRampSetup {
  function testAFN() public {
    assertEq(s_offRamp.isAFNHealthy(), true);
    s_mockAFN.voteToCurse(0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
    assertEq(s_offRamp.isAFNHealthy(), false);
    AFN.UnvoteToCurseRecord[] memory records = new AFN.UnvoteToCurseRecord[](1);
    records[0] = AFN.UnvoteToCurseRecord({curseVoteAddr: OWNER, cursesHash: bytes32(uint256(0)), forceUnvote: true});
    s_mockAFN.ownerUnvoteToCurse(records);
    assertEq(s_offRamp.isAFNHealthy(), true);
  }
}
