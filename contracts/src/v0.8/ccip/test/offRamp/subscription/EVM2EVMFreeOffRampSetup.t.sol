// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../mocks/MockCommitStore.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../helpers/MerkleHelper.sol";
import "../../TokenSetup.t.sol";
import "../../helpers/ramps/EVM2EVMFreeOffRampHelper.sol";

contract EVM2EVMFreeOffRampSetup is TokenSetup {
  EVM2EVMFreeOffRampHelper internal s_offRamp;
  Any2EVMSubscriptionOffRampRouter internal s_router;

  CommitStoreInterface internal s_mockCommitStore;
  SimpleMessageReceiver internal s_receiver;
  Any2EVMMessageReceiverInterface internal s_secondary_receiver;
  MerkleHelper internal s_merkleHelper;

  event ExecutionStateChanged(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  IERC20 internal s_destFeeToken;

  uint256 internal constant SUBSCRIPTION_BALANCE = 1e18;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_destFeeToken = IERC20(s_destTokens[0]);

    s_mockCommitStore = new MockCommitStore();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    s_merkleHelper = new MerkleHelper();

    _deployOffRampAndRouter(s_mockCommitStore);
  }

  // This function us re-used in the e2e test as we need a real commitStore
  // there while we require a mock version for all other tests.
  function _deployOffRampAndRouter(CommitStoreInterface commitStore) internal {
    s_offRamp = new EVM2EVMFreeOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      offRampConfig(),
      ON_RAMP_ADDRESS,
      commitStore,
      s_afn,
      getCastedSourceTokens(),
      getCastedDestinationPools(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN
    );

    s_offRamp.setPrices(getCastedDestinationTokens(), getTokenPrices());

    BaseOffRampInterface[] memory offRamps = new BaseOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMSubscriptionOffRampRouter(offRamps, subscriptionConfig(s_destFeeToken));
    s_offRamp.setRouter(s_router);

    NativeTokenPool(address(s_destPools[0])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
  }

  function _generateAny2EVMSubscriptionMessageNoTokens(uint64 sequenceNumber, uint64 nonce)
    internal
    view
    returns (CCIP.EVM2EVMSubscriptionMessage memory)
  {
    return _generateAny2EVMSubscriptionMessage(sequenceNumber, nonce, new CCIP.EVMTokenAndAmount[](0));
  }

  function _generateAny2EVMSubscriptionMessageWithTokens(
    uint64 sequenceNumber,
    uint64 nonce,
    uint256[] memory amounts
  ) internal view returns (CCIP.EVM2EVMSubscriptionMessage memory) {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    for (uint256 i = 0; i < tokensAndAmounts.length; i++) {
      tokensAndAmounts[i].amount = amounts[i];
    }
    return _generateAny2EVMSubscriptionMessage(sequenceNumber, nonce, tokensAndAmounts);
  }

  function _generateMessagesWithTokens() internal view returns (CCIP.EVM2EVMSubscriptionMessage[] memory) {
    CCIP.EVM2EVMSubscriptionMessage[] memory messages = new CCIP.EVM2EVMSubscriptionMessage[](2);
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = 1000;
    tokensAndAmounts[1].amount = 50;
    messages[0] = _generateAny2EVMSubscriptionMessage(1, 1, tokensAndAmounts);
    messages[1] = _generateAny2EVMSubscriptionMessage(2, 2, tokensAndAmounts);
    return messages;
  }

  function _generateAny2EVMSubscriptionMessage(
    uint64 sequenceNumber,
    uint64 nonce,
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts
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
        tokensAndAmounts,
        GAS_LIMIT
      );
  }

  function _convertSubscriptionToGeneralMessage(CCIP.EVM2EVMSubscriptionMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory)
  {
    uint256 numberOfTokens = original.tokensAndAmounts.length;
    CCIP.EVMTokenAndAmount[] memory destTokensAndAmounts = new CCIP.EVMTokenAndAmount[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = s_offRamp.getPool(IERC20(original.tokensAndAmounts[i].token));
      destPools[i] = address(pool);
      destTokensAndAmounts[i].token = address(pool.getToken());
      destTokensAndAmounts[i].amount = original.tokensAndAmounts[i].amount;
    }

    return
      CCIP.Any2EVMMessageFromSender({
        sourceChainId: original.sourceChainId,
        sender: abi.encode(original.sender),
        receiver: original.receiver,
        data: original.data,
        destTokensAndAmounts: destTokensAndAmounts,
        destPools: destPools,
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
    address[] memory tokenPerFeeCoinAddresses = new address[](0);
    uint256[] memory tokenPerFeeCoin = new uint256[](1);
    tokenPerFeeCoin[0] = TOKENS_PER_FEE_COIN;

    CCIP.FeeUpdate[] memory feeUpdates = new CCIP.FeeUpdate[](0);

    return
      CCIP.ExecutionReport({
        sequenceNumbers: sequenceNumbers,
        innerProofs: innerProofs,
        innerProofFlagBits: 2**256 - 1,
        outerProofs: outerProofs,
        outerProofFlagBits: 2**256 - 1,
        encodedMessages: encodedMessages,
        tokenPerFeeCoinAddresses: tokenPerFeeCoinAddresses,
        tokenPerFeeCoin: tokenPerFeeCoin,
        feeUpdates: feeUpdates
      });
  }
}
