// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../mocks/MockCommitStore.sol";
import "../../mocks/MockTollOffRampRouter.sol";
import "../../helpers/ramps/EVM2EVMTollOffRampHelper.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../TokenSetup.t.sol";

contract EVM2EVMTollOffRampSetup is TokenSetup {
  CommitStoreInterface internal s_mockCommitStore;
  Any2EVMMessageReceiverInterface internal s_receiver;
  Any2EVMMessageReceiverInterface internal s_secondary_receiver;

  EVM2EVMTollOffRampHelper internal s_offRamp;

  uint256 internal constant EXECUTION_FEE_AMOUNT = 1e18;

  event ExecutionStateChanged(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_mockCommitStore = new MockCommitStore();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    deployOffRamp(s_mockCommitStore);
  }

  function deployOffRamp(CommitStoreInterface commitStore) internal {
    s_offRamp = new EVM2EVMTollOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      offRampConfig(),
      commitStore,
      s_afn,
      getCastedSourceTokens(),
      getCastedDestinationPools(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN
    );

    s_offRamp.setPrices(getCastedDestinationTokens(), getTokenPrices());

    NativeTokenPool(address(s_destPools[0])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
  }

  function _generateNewRouter() internal returns (Any2EVMOffRampRouterInterface newRouter) {
    newRouter = new MockTollOffRampRouter();
    assertTrue(address(newRouter) != address(s_offRamp.getRouter()));
  }

  function _convertTollToGeneralMessage(CCIP.EVM2EVMTollMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory message)
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

  function _generateAny2EVMTollMessageNoTokens(uint64 sequenceNumber)
    internal
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    return _generateAny2EVMTollMessage(sequenceNumber, getCastedSourceEVMTokenAndAmountsWithZeroAmounts());
  }

  function _generateAny2EVMTollMessageWithTokens(uint64 sequenceNumber, uint256[] memory amounts)
    internal
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    for (uint256 i = 0; i < tokensAndAmounts.length; ++i) {
      tokensAndAmounts[i].amount = amounts[i];
    }
    return _generateAny2EVMTollMessage(sequenceNumber, tokensAndAmounts);
  }

  function _generateAny2EVMTollMessage(uint64 sequenceNumber, CCIP.EVMTokenAndAmount[] memory tokensAndAmounts)
    internal
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    bytes memory data = abi.encode(0);
    CCIP.EVMTokenAndAmount memory feeToken = CCIP.EVMTokenAndAmount({token: tokensAndAmounts[0].token, amount: EXECUTION_FEE_AMOUNT});
    return
      CCIP.EVM2EVMTollMessage(
        SOURCE_CHAIN_ID,
        sequenceNumber,
        OWNER,
        address(s_receiver),
        data,
        tokensAndAmounts,
        feeToken,
        GAS_LIMIT
      );
  }

  function _generateBasicMessages() internal view returns (CCIP.EVM2EVMTollMessage[] memory) {
    CCIP.EVM2EVMTollMessage[] memory messages = new CCIP.EVM2EVMTollMessage[](1);
    messages[0] = _generateAny2EVMTollMessageNoTokens(1);
    return messages;
  }

  function _generateMessagesWithTokens() internal view returns (CCIP.EVM2EVMTollMessage[] memory) {
    CCIP.EVM2EVMTollMessage[] memory messages = new CCIP.EVM2EVMTollMessage[](2);
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = getCastedSourceEVMTokenAndAmountsWithZeroAmounts();
    CCIP.EVMTokenAndAmount memory feeToken = tokensAndAmounts[0];
    feeToken.amount = EXECUTION_FEE_AMOUNT;
    tokensAndAmounts[0].amount = 1e18;
    tokensAndAmounts[1].amount = 5e18;
    messages[0] = _generateAny2EVMTollMessage(10, tokensAndAmounts);
    messages[0].feeTokenAndAmount = feeToken;
    messages[1] = _generateAny2EVMTollMessage(11, tokensAndAmounts);
    messages[1].feeTokenAndAmount = feeToken;
    return messages;
  }

  function _generateReportFromMessages(CCIP.EVM2EVMTollMessage[] memory messages)
    internal
    view
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
    address[] memory tokenPerFeeCoinAddresses = new address[](1);
    // The first destination token is the fee token
    tokenPerFeeCoinAddresses[0] = address(s_destTokens[0]);
    uint256[] memory tokenPerFeeCoin = new uint256[](1);
    tokenPerFeeCoin[0] = TOKENS_PER_FEE_COIN;

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
}
