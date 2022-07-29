// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../blobVerifier/BlobVerifierSetup.t.sol";
import "../onRamp/subscription/EVM2EVMSubscriptionOnRampSetup.t.sol";
import "../offRamp/subscription/Any2EVMSubscriptionOffRampSetup.t.sol";

contract E2E_subscription is EVM2EVMSubscriptionOnRampSetup, BlobVerifierSetup, Any2EVMSubscriptionOffRampSetup {
  function setUp()
    public
    virtual
    override(EVM2EVMSubscriptionOnRampSetup, BlobVerifierSetup, Any2EVMSubscriptionOffRampSetup)
  {
    EVM2EVMSubscriptionOnRampSetup.setUp();
    BlobVerifierSetup.setUp();
    Any2EVMSubscriptionOffRampSetup.setUp();

    // This overwrites the setup done in Any2EVMSubscriptionOffRampSetup because
    // we need to use a real blob verifier and not a mock.
    deployOffRampAndRouter(s_blobVerifier);
  }

  function testSuccessWithTokens() public {
    uint256 balance0Pre = s_sourceTokens[0].balanceOf(OWNER);
    uint256 balance1Pre = s_sourceTokens[1].balanceOf(OWNER);
    uint256 subscriptionBalance = s_onRampRouter.getBalance(OWNER);

    CCIP.Any2EVMSubscriptionMessage[] memory messages = new CCIP.Any2EVMSubscriptionMessage[](3);
    messages[0] = parseEventToDestChainMessage(sendRequest(getTokenMessage(), 1));
    messages[1] = parseEventToDestChainMessage(sendRequest(getTokenMessage(), 2));
    messages[2] = parseEventToDestChainMessage(sendRequest(getTokenMessage(), 3));

    // Asserts that the tokens have been sent and the fee has been paid.
    assertEq(balance0Pre - messages.length * TOKEN_AMOUNT_0, s_sourceTokens[0].balanceOf(OWNER));
    assertEq(balance1Pre - messages.length * TOKEN_AMOUNT_1, s_sourceTokens[1].balanceOf(OWNER));
    assertEq(subscriptionBalance - messages.length * s_onRampRouter.getFee(), s_onRampRouter.getBalance(OWNER));

    _relayAndExecute(messages);
  }

  function testSuccessWithoutTokens() public {
    uint256 subscriptionBalance = s_onRampRouter.getBalance(OWNER);

    CCIP.Any2EVMSubscriptionMessage[] memory messages = new CCIP.Any2EVMSubscriptionMessage[](3);
    messages[0] = parseEventToDestChainMessage(sendRequest(getEmptyMessage(), 1));
    messages[1] = parseEventToDestChainMessage(sendRequest(getEmptyMessage(), 2));
    messages[2] = parseEventToDestChainMessage(sendRequest(getEmptyMessage(), 3));

    // Asserts that the tokens have been sent and the fee has been paid.
    assertEq(subscriptionBalance - messages.length * s_onRampRouter.getFee(), s_onRampRouter.getBalance(OWNER));

    _relayAndExecute(messages);
  }

  function _relayAndExecute(CCIP.Any2EVMSubscriptionMessage[] memory messages) internal {
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
    assertEq(BLOCK_TIME, timestamp);

    // We change the block time so when execute would e.g. use the current
    // block time instead of the relayed block time the value would be
    // incorrect in the checks below.
    vm.warp(BLOCK_TIME + 2000);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[0].sequenceNumber, CCIP.MessageExecutionState.Success);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[1].sequenceNumber, CCIP.MessageExecutionState.Success);

    vm.expectEmit(false, false, false, true);
    emit ExecutionCompleted(messages[2].sequenceNumber, CCIP.MessageExecutionState.Success);

    s_offRamp.execute(_generateReportFromMessages(messages), false);
  }

  function sendRequest(CCIP.EVM2AnySubscriptionMessage memory message, uint64 expectedSeqNum)
    public
    returns (CCIP.EVM2EVMSubscriptionEvent memory)
  {
    if (message.amounts.length != 0) {
      s_sourceTokens[0].approve(address(s_onRampRouter), TOKEN_AMOUNT_0);
      s_sourceTokens[1].approve(address(s_onRampRouter), TOKEN_AMOUNT_1);
    }

    message.receiver = address(s_receiver);
    uint64 expectedNonce = expectedSeqNum;
    if (expectedSeqNum == 3) {
      expectedNonce = 1;
      message.receiver = address(s_secondary_receiver);
    }
    CCIP.EVM2EVMSubscriptionEvent memory subscriptionEvent = messageToEvent(message, expectedSeqNum, expectedNonce);

    vm.expectEmit(false, false, false, true);
    emit CCIPSendRequested(subscriptionEvent);

    s_onRampRouter.ccipSend(DEST_CHAIN_ID, message);

    return subscriptionEvent;
  }

  function parseEventToDestChainMessage(CCIP.EVM2EVMSubscriptionEvent memory sendEvent)
    public
    pure
    returns (CCIP.Any2EVMSubscriptionMessage memory)
  {
    return
      CCIP.Any2EVMSubscriptionMessage({
        sourceChainId: sendEvent.sourceChainId,
        sequenceNumber: sendEvent.sequenceNumber,
        sender: sendEvent.sender,
        receiver: sendEvent.receiver,
        nonce: sendEvent.nonce,
        data: sendEvent.data,
        tokens: sendEvent.tokens,
        amounts: sendEvent.amounts,
        gasLimit: sendEvent.gasLimit
      });
  }
}
