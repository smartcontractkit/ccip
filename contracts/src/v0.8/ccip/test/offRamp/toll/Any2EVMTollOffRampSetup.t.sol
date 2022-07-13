// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../blobVerifier/interfaces/BlobVerifierInterface.sol";
import "../../../health/interfaces/AFNInterface.sol";
import "../../../offRamp/interfaces/Any2EVMTollOffRampInterface.sol";
import "../../../offRamp/toll/Any2EVMTollOffRamp.sol";
import "../../mocks/MockBlobVerifier.sol";
import "../../TokenSetup.t.sol";
import "../../helpers/Any2EVMTollOffRampHelper.sol";
import "../../mocks/MockTollOffRampRouter.sol";
import "../../helpers/receivers/SimpleMessageReceiver.sol";

contract Any2EVMTollOffRampSetup is TokenSetup {
  Any2EVMTollOffRampInterface.OffRampConfig internal s_offRampConfig;
  BlobVerifierInterface internal s_mockBlobVerifier;
  CrossChainMessageReceiverInterface internal s_receiver;
  CrossChainMessageReceiverInterface internal s_secondary_receiver;

  Any2EVMTollOffRampHelper internal s_offRamp;

  event ExecutionCompleted(uint64 indexed sequenceNumber, CCIP.MessageExecutionState state);

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_offRampConfig = BaseOffRampInterface.OffRampConfig({
      sourceChainId: SOURCE_CHAIN_ID,
      executionDelaySeconds: 10,
      maxDataSize: 500,
      maxTokensLength: 5,
      permissionLessExecutionThresholdSeconds: 500
    });

    s_mockBlobVerifier = new MockBlobVerifier();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();

    s_offRamp = new Any2EVMTollOffRampHelper(
      DEST_CHAIN_ID,
      s_offRampConfig,
      s_mockBlobVerifier,
      ON_RAMP_ADDRESS,
      s_afn,
      s_sourceTokens,
      s_destPools,
      HEARTBEAT
    );
  }

  function _generateNewRouter() internal returns (Any2EVMTollOffRampRouterInterface newRouter) {
    newRouter = new MockTollOffRampRouter();
    assertTrue(address(newRouter) != address(s_offRamp.getRouter()));
  }

  function _generateAny2EVMTollMessageNoTokens(uint64 sequenceNumber)
    internal
    view
    returns (CCIP.Any2EVMTollMessage memory)
  {
    IERC20[] memory tokens;
    uint256[] memory amounts;

    return _generateAny2EVMTollMessage(sequenceNumber, tokens, amounts);
  }

  function _generateAny2EVMTollMessage(
    uint64 sequenceNumber,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal view returns (CCIP.Any2EVMTollMessage memory) {
    bytes memory data = abi.encode(0);
    return
      CCIP.Any2EVMTollMessage(
        SOURCE_CHAIN_ID,
        sequenceNumber,
        OWNER,
        address(s_receiver),
        data,
        tokens,
        amounts,
        s_sourceTokens[0],
        0,
        0
      );
  }

  function _generateBasicMessages() internal view returns (CCIP.Any2EVMTollMessage[] memory) {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](1);
    messages[0] = _generateAny2EVMTollMessageNoTokens(1);
    return messages;
  }

  function _generateMessagesWithTokens() internal view returns (CCIP.Any2EVMTollMessage[] memory) {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](2);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = _generateAny2EVMTollMessage(10, s_sourceTokens, amounts);
    messages[0].feeTokenAmount = 10;
    messages[1] = _generateAny2EVMTollMessage(11, s_sourceTokens, amounts);
    messages[1].feeTokenAmount = 10;
    return messages;
  }

  function _generateReportFromMessages(CCIP.Any2EVMTollMessage[] memory messages)
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
    address[] memory tokenPerFeeCoinAddresses = new address[](3);
    tokenPerFeeCoinAddresses[0] = address(s_destTokens[0]);
    tokenPerFeeCoinAddresses[1] = address(s_sourceTokens[0]);
    uint256[] memory tokenPerFeeCoin = new uint256[](3);
    tokenPerFeeCoin[0] = 1;
    tokenPerFeeCoin[1] = 1;
    tokenPerFeeCoin[2] = 1;

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
