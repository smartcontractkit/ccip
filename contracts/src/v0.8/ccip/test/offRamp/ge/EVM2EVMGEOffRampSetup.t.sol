// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../TokenSetup.t.sol";
import {GasFeeCacheSetup} from "../../gasFeeCache/GasFeeCache.t.sol";
import {MockCommitStore, CommitStoreInterface} from "../../mocks/MockCommitStore.sol";
import {SimpleMessageReceiver, Any2EVMMessageReceiverInterface} from "../../helpers/receivers/SimpleMessageReceiver.sol";
import {EVM2EVMGEOffRampInterface} from "../../../interfaces/offRamp/EVM2EVMGEOffRampInterface.sol";
import {GE} from "../../../models/GE.sol";
import {Common} from "../../../models/Common.sol";
import {GasFeeCacheInterface} from "../../../gasFeeCache/GasFeeCache.sol";
import {EVM2EVMGEOffRampHelper} from "../../helpers/ramps/EVM2EVMGEOffRampHelper.sol";

contract EVM2EVMGEOffRampSetup is TokenSetup, GasFeeCacheSetup {
  CommitStoreInterface internal s_mockCommitStore;
  Any2EVMMessageReceiverInterface internal s_receiver;
  Any2EVMMessageReceiverInterface internal s_secondary_receiver;

  EVM2EVMGEOffRampHelper internal s_offRamp;

  uint256 internal constant EXECUTION_FEE_AMOUNT = 1e18;

  event ExecutionStateChanged(
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    Internal.MessageExecutionState state
  );
  event SkippedIncorrectNonce(uint64 indexed nonce, address indexed sender);

  function setUp() public virtual override(TokenSetup, GasFeeCacheSetup) {
    TokenSetup.setUp();
    GasFeeCacheSetup.setUp();

    s_mockCommitStore = new MockCommitStore();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    deployOffRamp(s_mockCommitStore, s_gasFeeCache);
  }

  function deployOffRamp(CommitStoreInterface commitStore, GasFeeCacheInterface gasFeeCache) internal {
    s_offRamp = new EVM2EVMGEOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      _generateGEOffRampConfig(gasFeeCache),
      ON_RAMP_ADDRESS,
      commitStore,
      s_afn,
      getCastedSourceTokens(),
      getCastedDestinationPools(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN,
      IERC20(s_destFeeToken)
    );

    s_offRamp.setPrices(getCastedDestinationTokens(), getTokenPrices());
    s_gasFeeCache.setFeeUpdater(address(s_offRamp));

    NativeTokenPool(address(s_destPools[0])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
  }

  function _generateGEOffRampConfig(GasFeeCacheInterface gasFeeCache)
    public
    pure
    returns (EVM2EVMGEOffRampInterface.GEOffRampConfig memory)
  {
    return
      EVM2EVMGEOffRampInterface.GEOffRampConfig({
        gasOverhead: 5e5,
        gasFeeCache: gasFeeCache,
        executionDelaySeconds: EXECUTION_DELAY_SECONDS,
        maxDataSize: MAX_DATA_SIZE,
        maxTokensLength: MAX_TOKENS_LENGTH,
        permissionLessExecutionThresholdSeconds: PERMISSION_LESS_EXECUTION_THRESHOLD_SECONDS
      });
  }

  function _convertGEToGeneralMessage(GE.EVM2EVMGEMessage memory original)
    internal
    view
    returns (Internal.Any2EVMMessageFromSender memory message)
  {
    uint256 numberOfTokens = original.tokensAndAmounts.length;
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = s_offRamp.getPoolBySourceToken(IERC20(original.tokensAndAmounts[i].token));
      destPools[i] = address(pool);
      destTokensAndAmounts[i].token = address(pool.getToken());
      destTokensAndAmounts[i].amount = original.tokensAndAmounts[i].amount;
    }

    return
      Internal.Any2EVMMessageFromSender({
        sourceChainId: original.sourceChainId,
        sender: abi.encode(original.sender),
        receiver: original.receiver,
        data: original.data,
        destTokensAndAmounts: destTokensAndAmounts,
        destPools: destPools,
        gasLimit: original.gasLimit
      });
  }

  function _generateAny2EVMGEMessageNoTokens(uint64 sequenceNumber) internal view returns (GE.EVM2EVMGEMessage memory) {
    return _generateAny2EVMGEMessage(sequenceNumber, getCastedSourceEVMTokenAndAmountsWithZeroAmounts());
  }

  function _generateAny2EVMGEMessageWithTokens(uint64 sequenceNumber, uint256[] memory amounts)
    internal
    view
    returns (GE.EVM2EVMGEMessage memory)
  {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    for (uint256 i = 0; i < tokensAndAmounts.length; ++i) {
      tokensAndAmounts[i].amount = amounts[i];
    }
    return _generateAny2EVMGEMessage(sequenceNumber, tokensAndAmounts);
  }

  function _generateAny2EVMGEMessage(uint64 sequenceNumber, Common.EVMTokenAndAmount[] memory tokensAndAmounts)
    internal
    view
    returns (GE.EVM2EVMGEMessage memory)
  {
    bytes memory data = abi.encode(0);
    GE.EVM2EVMGEMessage memory message = GE.EVM2EVMGEMessage({
      sequenceNumber: sequenceNumber,
      feeTokenAmount: EXECUTION_FEE_AMOUNT,
      sender: OWNER,
      nonce: sequenceNumber,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainId: SOURCE_CHAIN_ID,
      receiver: address(s_receiver),
      data: data,
      tokensAndAmounts: tokensAndAmounts,
      feeToken: tokensAndAmounts[0].token,
      messageId: ""
    });
    message.messageId = GE._hash(
      message,
      keccak256(abi.encode(GE.EVM_2_EVM_GE_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );

    return message;
  }

  function _generateBasicMessages() internal view returns (GE.EVM2EVMGEMessage[] memory) {
    GE.EVM2EVMGEMessage[] memory messages = new GE.EVM2EVMGEMessage[](1);
    messages[0] = _generateAny2EVMGEMessageNoTokens(1);
    return messages;
  }

  function _generateMessagesWithTokens() internal view returns (GE.EVM2EVMGEMessage[] memory) {
    GE.EVM2EVMGEMessage[] memory messages = new GE.EVM2EVMGEMessage[](2);
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    tokensAndAmounts[0].amount = 1e18;
    tokensAndAmounts[1].amount = 5e18;
    messages[0] = _generateAny2EVMGEMessage(1, tokensAndAmounts);
    messages[0].feeTokenAmount = EXECUTION_FEE_AMOUNT;
    messages[1] = _generateAny2EVMGEMessage(2, tokensAndAmounts);
    messages[1].feeTokenAmount = EXECUTION_FEE_AMOUNT;
    return messages;
  }

  function _generateReportFromMessages(GE.EVM2EVMGEMessage[] memory messages)
    internal
    view
    returns (GE.ExecutionReport memory)
  {
    bytes[] memory encodedMessages = new bytes[](messages.length);
    uint64[] memory sequenceNumbers = new uint64[](messages.length);
    for (uint256 i = 0; i < messages.length; ++i) {
      encodedMessages[i] = abi.encode(messages[i]);
      sequenceNumbers[i] = messages[i].sequenceNumber;
    }

    bytes32[] memory innerProofs = new bytes32[](0);
    bytes32[] memory outerProofs = new bytes32[](0);
    address[] memory tokenPerFeeCoinAddresses = new address[](1);
    // The first destination token is the fee token
    tokenPerFeeCoinAddresses[0] = address(s_destTokens[0]);
    uint256[] memory tokenPerFeeCoin = new uint256[](1);
    tokenPerFeeCoin[0] = TOKENS_PER_FEE_COIN;

    GE.FeeUpdate[] memory feeUpdates = new GE.FeeUpdate[](0);

    return
      GE.ExecutionReport({
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
