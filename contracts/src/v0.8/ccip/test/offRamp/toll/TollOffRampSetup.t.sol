// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "../../../interfaces/TollOffRampInterface.sol";
import "../../../interfaces/BlobVerifierInterface.sol";
import "../../../interfaces/AFNInterface.sol";
import "../../../offRamp/toll/Any2EVMTollOffRamp.sol";
import "../../../utils/CCIP.sol";
import "../../mocks/MockBlobVerifier.sol";
import "../../helpers/SimpleMessageReceiver.sol";
import "../../TokenSetup.t.sol";

contract TollOffRampSetup is TokenSetup {
  TollOffRampInterface.OffRampConfig s_offRampConfig;
  BlobVerifierInterface s_mockBlobVerifier;
  CrossChainMessageReceiverInterface s_receiver;
  CrossChainMessageReceiverInterface s_secondary_receiver;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_offRampConfig = TollOffRampInterface.OffRampConfig({
      sourceChainId: s_sourceChainId,
      executionDelaySeconds: 0,
      maxDataSize: 500,
      maxTokensLength: 5
    });

    s_mockBlobVerifier = new MockBlobVerifier();
    s_receiver = new SimpleMessageReceiver();
    s_secondary_receiver = new SimpleMessageReceiver();
  }

  function getAny2EVMTollMessageNoTokens(uint64 sequenceNumber) public view returns (CCIP.Any2EVMTollMessage memory) {
    IERC20[] memory tokens;
    uint256[] memory amounts;

    return getAny2EVMTollMessage(sequenceNumber, tokens, amounts);
  }

  function getAny2EVMTollMessage(
    uint64 sequenceNumber,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) public view returns (CCIP.Any2EVMTollMessage memory) {
    bytes memory data = abi.encode(0);
    return
      CCIP.Any2EVMTollMessage(
        s_sourceChainId,
        sequenceNumber,
        s_owner,
        address(s_receiver),
        data,
        tokens,
        amounts,
        s_sourceTokens[0],
        0,
        0
      );
  }

  function getBasicMessages() public view returns (CCIP.Any2EVMTollMessage[] memory) {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](1);
    messages[0] = getAny2EVMTollMessageNoTokens(1);
    return messages;
  }

  function getMessagesWithTokens() public view returns (CCIP.Any2EVMTollMessage[] memory) {
    CCIP.Any2EVMTollMessage[] memory messages = new CCIP.Any2EVMTollMessage[](2);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 1000;
    amounts[1] = 50;
    messages[0] = getAny2EVMTollMessage(10, s_sourceTokens, amounts);
    messages[0].feeTokenAmount = 10;
    messages[1] = getAny2EVMTollMessage(11, s_sourceTokens, amounts);
    messages[1].feeTokenAmount = 10;
    return messages;
  }

  function createReportFromMessages(CCIP.Any2EVMTollMessage[] memory messages)
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
    address[] memory tokenPerFeeCoinAddresses = new address[](2);
    uint256[] memory tokenPerFeeCoin = new uint256[](2);

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
