// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./blobVerifier/BlobVerifierSetup.t.sol";
import "./offRamp/toll/TollOffRampSetup.t.sol";
import "./helpers/MerkleHelper.sol";
import "../onRamp/toll/EVM2AnyTollOnRampRouter.sol";
import "../onRamp/toll/EVM2EVMTollOnRamp.sol";
import "./onRamp/toll/OnRampSetup.t.sol";

contract E2ETest is OnRampSetup, BlobVerifierSetup, TollOffRampSetup {
  EVM2AnyTollOnRampRouter public s_onRampRouter;
  EVM2EVMTollOnRamp public s_onRamp;
  Any2EVMTollOffRamp s_offRamp;
  TollOffRampRouterInterface s_router;

  MerkleHelper s_merkleHelper;

  function setUp() public virtual override(OnRampSetup, BlobVerifierSetup, TollOffRampSetup) {
    OnRampSetup.setUp();
    BlobVerifierSetup.setUp();
    TollOffRampSetup.setUp();

    s_onRampRouter = new EVM2AnyTollOnRampRouter();
    s_merkleHelper = new MerkleHelper();

    address[] memory allowList;

    s_onRamp = new EVM2EVMTollOnRamp(
      s_sourceChainId,
      s_destChainId,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      allowList,
      s_afn,
      1e18,
      TollOnRampInterface.OnRampConfig({
        router: address(s_onRampRouter),
        relayingFeeJuels: 0,
        maxDataSize: 50,
        maxTokensLength: 3
      })
    );

    s_onRampRouter.setOnRamp(s_destChainId, s_onRamp);

    s_offRamp = new Any2EVMTollOffRamp(
      s_sourceChainId,
      s_offRampConfig,
      s_blobVerifier,
      s_onRampAddress,
      s_afn,
      s_sourceTokens,
      s_destPools,
      2**20
    );
    TollOffRampInterface[] memory offRamps = new TollOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  function testE2ENoTokenNoFee() public {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](3);
    messages[0] = parseEventToDestChainMessage(sendRequest(1));
    messages[1] = parseEventToDestChainMessage(sendRequest(2));
    messages[2] = parseEventToDestChainMessage(sendRequest(3));

    bytes32[] memory hashedMessages = new bytes32[](3);
    hashedMessages[0] = keccak256(bytes.concat(hex"00", abi.encode(messages[0])));
    hashedMessages[1] = keccak256(bytes.concat(hex"00", abi.encode(messages[1])));
    hashedMessages[2] = keccak256(bytes.concat(hex"00", abi.encode(messages[2])));

    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(messages[0].sequenceNumber, messages[2].sequenceNumber);

    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = s_merkleHelper.getMerkleRoot(hashedMessages);

    address[] memory onRamps = new address[](1);
    onRamps[0] = s_config.onRamps[0];

    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: merkleRoots[0]
    });

    s_blobVerifier.report(abi.encode(report));
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_blobVerifier.verify(merkleRoots, proofs, 2**2 - 1, proofs, 2**2 - 1);
    assertEq(s_blockTime, timestamp);

    // We change the block time so when execute would e.g. use the current
    // block time instead of the relayed block time the value would be
    // incorrect in the checks below.
    vm.warp(s_blockTime + 2000);

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);

    CCIP.ExecutionResult[] memory executionResult = s_offRamp.execute(executionReport, false);

    assertEq(executionResult.length, messages.length);
    assertEq(executionResult[0].sequenceNumber, messages[0].sequenceNumber);
    assertTrue(executionResult[0].state == CCIP.MessageExecutionState.Success);
    assertEq(executionResult[0].timestampRelayed, s_blockTime);
    assertEq(executionResult[1].sequenceNumber, messages[1].sequenceNumber);
    assertTrue(executionResult[1].state == CCIP.MessageExecutionState.Success);
    assertEq(executionResult[1].timestampRelayed, s_blockTime);
    assertEq(executionResult[2].sequenceNumber, messages[2].sequenceNumber);
    assertTrue(executionResult[2].state == CCIP.MessageExecutionState.Success);
    assertEq(executionResult[2].timestampRelayed, s_blockTime);
  }

  function sendRequest(uint64 expectedSeqNum) public returns (CCIP.EVM2EVMTollEvent memory) {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.receiver = address(s_receiver);
    CCIP.EVM2EVMTollEvent memory tollEvent = messageToEvent(message, expectedSeqNum);
    vm.expectEmit(true, false, false, true);

    emit CCIPSendRequested(tollEvent);

    s_onRampRouter.ccipSend(s_destChainId, message);

    return tollEvent;
  }

  function parseEventToDestChainMessage(CCIP.EVM2EVMTollEvent memory sendEvent)
    public
    pure
    returns (CCIP.Any2EVMTollMessage memory)
  {
    return
      CCIP.Any2EVMTollMessage({
        sourceChainId: sendEvent.sourceChainId,
        sequenceNumber: sendEvent.sequenceNumber,
        sender: sendEvent.sender,
        receiver: sendEvent.receiver,
        data: sendEvent.data,
        tokens: sendEvent.tokens,
        amounts: sendEvent.amounts,
        feeToken: sendEvent.feeToken,
        feeTokenAmount: sendEvent.feeTokenAmount,
        gasLimit: sendEvent.gasLimit
      });
  }
}
