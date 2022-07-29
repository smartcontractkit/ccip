// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../blobVerifier/interfaces/BlobVerifierInterface.sol";
import "../../../health/interfaces/AFNInterface.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../helpers/MerkleHelper.sol";
import "../../mocks/MockBlobVerifier.sol";
import "../../TokenSetup.t.sol";
import "../../../offRamp/mo/Any2EVMMOOffRampRouter.sol";
import "../../helpers/Any2EVMMOOffRampHelper.sol";

contract Any2EVMMOOffRampSetup is TokenSetup {
  Any2EVMMOOffRampHelper s_offRamp;
  Any2EVMMOOffRampRouter s_router;

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
    s_offRamp = new Any2EVMMOOffRampHelper(
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
    s_router = new Any2EVMMOOffRampRouter(offRamps, s_subscriptionConfig);
    s_offRamp.setRouter(s_router);

    createSubscription(s_receiver, s_router, SUBSCRIPTION_BALANCE);
    createSubscription(s_secondary_receiver, s_router, SUBSCRIPTION_BALANCE);
  }

  function getAny2EVMMOMessage(uint64 sequenceNumber, uint64 nonce)
    internal
    view
    returns (CCIP.Any2EVMMOMessage memory)
  {
    bytes memory data = abi.encode(0);
    return CCIP.Any2EVMMOMessage(SOURCE_CHAIN_ID, sequenceNumber, OWNER, address(s_receiver), nonce, data, 0);
  }

  function _generateBasicMessages() public view returns (CCIP.Any2EVMMOMessage[] memory) {
    CCIP.Any2EVMMOMessage[] memory messages = new CCIP.Any2EVMMOMessage[](1);
    messages[0] = getAny2EVMMOMessage(1, 1);
    return messages;
  }

  function _generateReportFromMessages(CCIP.Any2EVMMOMessage[] memory messages)
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

  function createSubscription(
    SubscriptionManagerInterface receiver,
    Any2EVMMOOffRampRouterInterface router,
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
