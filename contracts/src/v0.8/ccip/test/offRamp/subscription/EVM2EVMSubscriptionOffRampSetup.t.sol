// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../blobVerifier/interfaces/BlobVerifierInterface.sol";
import "../../../health/interfaces/AFNInterface.sol";
import "../../../offRamp/interfaces/Any2EVMOffRampRouterInterface.sol";
import "../../../offRamp/subscription/EVM2EVMSubscriptionOffRamp.sol";
import "../../../offRamp/subscription/Any2EVMSubscriptionOffRampRouter.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../helpers/EVM2EVMSubscriptionOffRampHelper.sol";
import "../../helpers/MerkleHelper.sol";
import "../../mocks/MockBlobVerifier.sol";
import "../../TokenSetup.t.sol";

contract EVM2EVMSubscriptionOffRampSetup is TokenSetup {
  EVM2EVMSubscriptionOffRampHelper s_offRamp;
  Any2EVMSubscriptionOffRampRouter s_router;

  BaseOffRampInterface.OffRampConfig s_offRampConfig;
  SubscriptionInterface.SubscriptionConfig s_subscriptionConfig;
  BlobVerifierInterface s_mockBlobVerifier;
  SimpleMessageReceiver s_receiver;
  Any2EVMMessageReceiverInterface s_secondary_receiver;
  MerkleHelper s_merkleHelper;

  event ExecutionStateChanged(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  IERC20 s_destFeeToken;

  uint256 internal constant SUBSCRIPTION_BALANCE = 1e18;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_offRampConfig = BaseOffRampInterface.OffRampConfig({
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

    _deployOffRampAndRouter(s_mockBlobVerifier);
  }

  // This function us re-used in the e2e test as we need a real blob verifier
  // there while we require a mock version for all other tests.
  function _deployOffRampAndRouter(BlobVerifierInterface blobVerifier) internal {
    s_offRamp = new EVM2EVMSubscriptionOffRampHelper(
      SOURCE_CHAIN_ID,
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

    NativeTokenPool(address(s_destPools[0])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);

    _createSubscription(SubscriptionManagerInterface(address(s_receiver)), s_router, SUBSCRIPTION_BALANCE, true);
    _createSubscription(
      SubscriptionManagerInterface(address(s_secondary_receiver)),
      s_router,
      SUBSCRIPTION_BALANCE,
      true
    );
  }

  function _generateAny2EVMSubscriptionMessageNoTokens(uint64 sequenceNumber, uint64 nonce)
    internal
    view
    returns (CCIP.EVM2EVMSubscriptionMessage memory)
  {
    IERC20[] memory tokens;
    uint256[] memory amounts;

    return _generateAny2EVMSubscriptionMessage(sequenceNumber, nonce, tokens, amounts);
  }

  function _generateAny2EVMSubscriptionMessageWithTokens(
    uint64 sequenceNumber,
    uint64 nonce,
    uint256[] memory amounts
  ) internal view returns (CCIP.EVM2EVMSubscriptionMessage memory) {
    return _generateAny2EVMSubscriptionMessage(sequenceNumber, nonce, s_sourceTokens, amounts);
  }

  function _generateMessagesWithTokens() internal view returns (CCIP.EVM2EVMSubscriptionMessage[] memory) {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = new CCIP.EVM2EVMSubscriptionMessage[](2);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = _generateAny2EVMSubscriptionMessage(1, 1, s_sourceTokens, amounts);
    messages[1] = _generateAny2EVMSubscriptionMessage(2, 2, s_sourceTokens, amounts);
    return messages;
  }

  function _generateAny2EVMSubscriptionMessage(
    uint64 sequenceNumber,
    uint64 nonce,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal view returns (CCIP.EVM2EVMSubscriptionMessage memory) {
    bytes memory data = abi.encode(0);
    return
      CCIP.EVM2EVMSubscriptionMessage(
        SOURCE_CHAIN_ID,
        sequenceNumber,
        OWNER,
        address(s_receiver),
        nonce,
        data,
        tokens,
        amounts,
        GAS_LIMIT
      );
  }

  function _convertSubscriptionToGeneralMessage(CCIP.EVM2EVMSubscriptionMessage memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory)
  {
    return
      CCIP.Any2EVMMessage({
        sourceChainId: original.sourceChainId,
        sequenceNumber: original.sequenceNumber,
        sender: abi.encode(original.sender),
        receiver: original.receiver,
        data: original.data,
        tokens: original.tokens,
        amounts: original.amounts,
        gasLimit: original.gasLimit
      });
  }

  function _generateBasicMessages() internal view returns (CCIP.EVM2EVMSubscriptionMessage[] memory) {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = new CCIP.EVM2EVMSubscriptionMessage[](1);
    messages[0] = _generateAny2EVMSubscriptionMessageNoTokens(1, 1);
    return messages;
  }

  function _generateReportFromMessages(CCIP.EVM2EVMSubscriptionMessage[] memory messages)
    internal
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
    tokenPerFeeCoin[0] = TOKENS_PER_FEE_COIN;
    tokenPerFeeCoin[1] = TOKENS_PER_FEE_COIN;
    tokenPerFeeCoin[2] = TOKENS_PER_FEE_COIN;

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

  function _createSubscription(
    SubscriptionManagerInterface receiver,
    Any2EVMOffRampRouterInterface router,
    uint256 funding,
    bool strictSequencing
  ) internal {
    address[] memory senders = new address[](1);
    senders[0] = OWNER;
    s_destFeeToken.approve(address(router), funding);
    Subscription(address(router)).createSubscription(
      SubscriptionInterface.OffRampSubscription({
        senders: senders,
        receiver: receiver,
        strictSequencing: strictSequencing,
        balance: funding
      })
    );
  }
}
