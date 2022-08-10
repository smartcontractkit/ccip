// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../mocks/MockBlobVerifier.sol";
import "../../mocks/MockTollOffRampRouter.sol";
import "../../helpers/EVM2EVMTollOffRampHelper.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";
import "../../TokenSetup.t.sol";

contract EVM2EVMTollOffRampSetup is TokenSetup {
  BlobVerifierInterface internal s_mockBlobVerifier;
  Any2EVMMessageReceiverInterface internal s_receiver;
  Any2EVMMessageReceiverInterface internal s_secondary_receiver;

  EVM2EVMTollOffRampHelper internal s_offRamp;

  uint256 internal constant EXECUTION_FEE_AMOUNT = 1e18;

  event ExecutionStateChanged(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  function setUp() public virtual override {
    TokenSetup.setUp();

    s_mockBlobVerifier = new MockBlobVerifier();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    s_offRamp = new EVM2EVMTollOffRampHelper(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      offRampConfig(),
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      s_sourceTokens,
      s_destPools,
      HEARTBEAT
    );

    NativeTokenPool(address(s_destPools[0])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
    NativeTokenPool(address(s_destPools[1])).setOffRamp(BaseOffRampInterface(address(s_offRamp)), true);
  }

  function _generateNewRouter() internal returns (Any2EVMOffRampRouterInterface newRouter) {
    newRouter = new MockTollOffRampRouter();
    assertTrue(address(newRouter) != address(s_offRamp.getRouter()));
  }

  function _convertTollToGeneralMessage(CCIP.EVM2EVMTollMessage memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory message)
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

  function _generateAny2EVMTollMessageNoTokens(uint64 sequenceNumber)
    internal
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    IERC20[] memory tokens;
    uint256[] memory amounts;

    return _generateAny2EVMTollMessage(sequenceNumber, tokens, amounts);
  }

  function _generateAny2EVMTollMessageWithTokens(uint64 sequenceNumber, uint256[] memory amounts)
    internal
    view
    returns (CCIP.EVM2EVMTollMessage memory)
  {
    return _generateAny2EVMTollMessage(sequenceNumber, s_sourceTokens, amounts);
  }

  function _generateAny2EVMTollMessage(
    uint64 sequenceNumber,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal view returns (CCIP.EVM2EVMTollMessage memory) {
    bytes memory data = abi.encode(0);
    return
      CCIP.EVM2EVMTollMessage(
        SOURCE_CHAIN_ID,
        sequenceNumber,
        OWNER,
        address(s_receiver),
        data,
        tokens,
        amounts,
        s_sourceTokens[0],
        EXECUTION_FEE_AMOUNT,
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
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = _generateAny2EVMTollMessage(10, s_sourceTokens, amounts);
    messages[0].feeToken = s_sourceTokens[0];
    messages[0].feeTokenAmount = EXECUTION_FEE_AMOUNT;
    messages[1] = _generateAny2EVMTollMessage(11, s_sourceTokens, amounts);
    messages[1].feeToken = s_sourceTokens[0];
    messages[1].feeTokenAmount = EXECUTION_FEE_AMOUNT;
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
