// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../blobVerifier/interfaces/BlobVerifierInterface.sol";
import "../../../health/interfaces/AFNInterface.sol";
import "../../../offRamp/interfaces/Any2EVMSubscriptionOffRampRouterInterface.sol";
import "../../../offRamp/subscription/Any2EVMSubscriptionOffRamp.sol";
import "../../../offRamp/subscription/Any2EVMSubscriptionOffRampRouter.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../helpers/Any2EVMSubscriptionOffRampHelper.sol";
import "../../helpers/MerkleHelper.sol";
import "../../mocks/MockBlobVerifier.sol";
import "../../TokenSetup.t.sol";

contract Any2EVMSubscriptionOffRampSetup is TokenSetup {
  Any2EVMSubscriptionOffRampHelper s_offRamp;
  Any2EVMSubscriptionOffRampRouter s_router;

  BaseOffRampInterface.OffRampConfig s_offRampConfig;
  SubscriptionInterface.SubscriptionConfig s_subscriptionConfig;
  BlobVerifierInterface s_mockBlobVerifier;
  CrossChainMessageReceiverInterface s_receiver;
  CrossChainMessageReceiverInterface s_secondary_receiver;
  MerkleHelper s_merkleHelper;

  event ExecutionCompleted(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  IERC20 s_destFeeToken;

  uint256 immutable SUBSCRIPTION_BALANCE = 1e7;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_offRampConfig = BaseOffRampInterface.OffRampConfig({
      sourceChainId: SOURCE_CHAIN_ID,
      executionDelaySeconds: 0,
      maxDataSize: 500,
      maxTokensLength: 5,
      permissionLessExecutionThresholdSeconds: 500
    });
    s_destFeeToken = s_destTokens[0];

    s_mockBlobVerifier = new MockBlobVerifier();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    s_merkleHelper = new MerkleHelper();

    deployOffRampAndRouter(s_mockBlobVerifier);
  }

  // This function us re-used in the e2e test as we need a real blob verifier
  // there while we require a mock version for all other tests.
  function deployOffRampAndRouter(BlobVerifierInterface blobVerifier) internal {
    s_offRamp = new Any2EVMSubscriptionOffRampHelper(
      DEST_CHAIN_ID,
      s_offRampConfig,
      blobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      s_sourceTokens,
      s_destPools,
      HEARTBEAT
    );
    s_subscriptionConfig = SubscriptionInterface.SubscriptionConfig(100, 100, s_destFeeToken);
    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMSubscriptionOffRampRouter(offRamps, s_subscriptionConfig);
    s_offRamp.setRouter(s_router);

    createSubscription(s_receiver, s_router, SUBSCRIPTION_BALANCE);
    createSubscription(s_secondary_receiver, s_router, SUBSCRIPTION_BALANCE);
  }

  function getAny2EVMSubscriptionMessageNoTokens(uint64 sequenceNumber, uint64 nonce)
    public
    view
    returns (CCIP.Any2EVMSubscriptionMessage memory)
  {
    IERC20[] memory tokens;
    uint256[] memory amounts;

    return getAny2EVMSubscriptionMessage(sequenceNumber, nonce, tokens, amounts);
  }

  function generateMessagesWithTokens() internal view returns (CCIP.Any2EVMSubscriptionMessage[] memory) {
    CCIP.Any2EVMSubscriptionMessage[] memory messages = new CCIP.Any2EVMSubscriptionMessage[](2);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = getAny2EVMSubscriptionMessage(10, 1, s_sourceTokens, amounts);
    messages[1] = getAny2EVMSubscriptionMessage(11, 2, s_sourceTokens, amounts);
    return messages;
  }

  function getAny2EVMSubscriptionMessage(
    uint64 sequenceNumber,
    uint64 nonce,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal view returns (CCIP.Any2EVMSubscriptionMessage memory) {
    bytes memory data = abi.encode(0);
    return
      CCIP.Any2EVMSubscriptionMessage(
        SOURCE_CHAIN_ID,
        sequenceNumber,
        OWNER,
        address(s_receiver),
        nonce,
        data,
        tokens,
        amounts,
        0
      );
  }

  function _generateBasicMessages() public view returns (CCIP.Any2EVMSubscriptionMessage[] memory) {
    CCIP.Any2EVMSubscriptionMessage[] memory messages = new CCIP.Any2EVMSubscriptionMessage[](1);
    messages[0] = getAny2EVMSubscriptionMessageNoTokens(1, 1);
    return messages;
  }

  function _generateReportFromMessages(CCIP.Any2EVMSubscriptionMessage[] memory messages)
    public
    pure
    returns (CCIP.ExecutionReport memory)
  {
    bytes[] memory encodedMessages = new bytes[](messages.length);
    uint64[] memory sequenceNumbers = new uint64[](messages.length);
    for (uint256 i = 0; i < messages.length; ++i) {
      encodedMessages[i] = abi.encode(messages[i]);
      sequenceNumbers[i] = messages[i].sequenceNumber;
    }

    bytes32[] memory innerProofs = new bytes32[](0);
    bytes32[] memory outerProofs = new bytes32[](0);
    address[] memory tokenPerFeeCoinAddresses = new address[](3);
    uint256[] memory tokenPerFeeCoin = new uint256[](3);
    tokenPerFeeCoin[0] = 1e18;
    tokenPerFeeCoin[1] = 1e18;
    tokenPerFeeCoin[2] = 1e18;

    return
      CCIP.ExecutionReport({
        sequenceNumbers: sequenceNumbers,
        innerProofs: innerProofs,
        innerProofFlagBits: 2**256 - 1,
        outerProofs: outerProofs,
        outerProofFlagBits: 2**256 - 1,
        encodedMessages: encodedMessages,
        tokenPerFeeCoinAddresses: tokenPerFeeCoinAddresses,
        tokenPerFeeCoin: tokenPerFeeCoin
      });
  }

  function _generateMessagesWithTokens() public view returns (CCIP.Any2EVMSubscriptionMessage[] memory) {
    CCIP.Any2EVMSubscriptionMessage[] memory messages = new CCIP.Any2EVMSubscriptionMessage[](2);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = getAny2EVMSubscriptionMessage(1, 1, s_sourceTokens, amounts);
    messages[1] = getAny2EVMSubscriptionMessage(2, 2, s_sourceTokens, amounts);
    return messages;
  }

  function createSubscription(
    SubscriptionManagerInterface receiver,
    Any2EVMSubscriptionOffRampRouterInterface router,
    uint256 funding
  ) public {
    address[] memory senders = new address[](1);
    senders[0] = OWNER;
    s_destFeeToken.approve(address(router), funding);
    router.createSubscription(
      SubscriptionInterface.OffRampSubscription({
        senders: senders,
        receiver: receiver,
        strictSequencing: true,
        balance: funding
      })
    );
  }
}
