// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";
import {Client} from "../../libraries/Client.sol";
import {CCIPReceiver} from "./CCIPReceiver.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {EnumerableMap} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableMap.sol";

/// @title CCIPReceiverWithACK
contract CCIPReceiverWithACK is CCIPReceiver {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.Bytes32ToUintMap;

  error InvalidAckMessageHeader();
  error MessageAlreadyAcknowledged(bytes32 messageId);

  event MessageAckSent(bytes32 incomingMessageId);
  event MessageSent(bytes32 indexed incomingMessageId, bytes32 indexed ACKMessageId);
  event MessageAckReceived(bytes32);
  event FeeTokenModified(address indexed oldToken, address indexed newToken);

  enum MessageType {
    OUTGOING,
    ACK
  }

  enum MessageStatus {
    QUIET,
    SENT,
    ACKNOWLEDGED
  }

  struct MessagePayload {
    bytes version;
    bytes data;
    MessageType messageType;
  }

  string public constant ACK_MESSAGE_HEADER = "MESSAGE_ACKNOWLEDGED_";

  // Current feeToken
  IERC20 public s_feeToken;

  mapping(bytes32 messageId => MessageStatus status) public s_messageStatus;

  constructor(address router, IERC20 feeToken) CCIPReceiver(router) {
    s_feeToken = feeToken;

    // If fee token is in LINK, then approve router to transfer
    if (address(feeToken) != address(0)) {
      feeToken.safeIncreaseAllowance(router, type(uint256).max);
    }
  }

  function modifyFeeToken(address token) external onlyOwner {
    // If the current fee token is not-native, zero out the allowance to the router for safety
    if (address(s_feeToken) != address(0)) {
      s_feeToken.safeApprove(getRouter(), 0);
    }

    address oldFeeToken = address(s_feeToken);
    s_feeToken = IERC20(token);

    // Approve the router to spend the new fee token
    if (token != address(0)) {
      s_feeToken.safeIncreaseAllowance(getRouter(), type(uint256).max);
    }

    emit FeeTokenModified(oldFeeToken, token);
  }

  /// @notice The entrypoint for the CCIP router to call. This function should never revert, all errors should be handled internally in this contract.
  /// @param message The message to process.
  /// @dev Extremely important to ensure only router calls this.
  function ccipReceive(Client.Any2EVMMessage calldata message)
    public
    override
    onlyRouter
    isValidSender(message.sourceChainSelector, message.sender)
    isValidChain(message.sourceChainSelector)
  {
    try this.processMessage(message) {}
    catch (bytes memory err) {
      s_failedMessages.set(message.messageId, uint256(ErrorCode.FAILED));
      s_messageContents[message.messageId] = message;

      // Don't revert so CCIPRouter doesn't revert. Emit event instead.
      // The message can be retried later without having to do manual execution of CCIP.
      emit MessageFailed(message.messageId, err);
      return;
    }

    emit MessageSucceeded(message.messageId);
  }

  /// @notice Application-specific logic for incoming ccip-messages.
  /// @dev Function does NOT require the status of an incoming ACK be "sent" because this implementation does not send, only receives
  /// @dev Any MessageType encoding is implemented by the sender-contract, and is not natively part of CCIP-messages.
  function processMessage(Client.Any2EVMMessage calldata message) external virtual override onlySelf {
    (MessagePayload memory payload) = abi.decode(message.data, (MessagePayload));

    if (payload.messageType == MessageType.OUTGOING) {
      // Insert Processing workflow here.

      // If the message was outgoing on the source chain, then send an ack response.
      _sendAck(message);
    } else if (payload.messageType == MessageType.ACK) {
      // Decode message into the message header and the messageId to ensure the message is encoded correctly
      (string memory messageHeader, bytes32 messageId) = abi.decode(payload.data, (string, bytes32));

      // Ensure Ack Message contains proper message header. Must abi.encode() before hashing since its of the string type
      if (keccak256(abi.encode(messageHeader)) != keccak256(abi.encode(ACK_MESSAGE_HEADER))) {
        revert InvalidAckMessageHeader();
      }

      // Make sure the ACK message has not already been acknowledged
      if (s_messageStatus[messageId] == MessageStatus.ACKNOWLEDGED) revert MessageAlreadyAcknowledged(messageId);

      // Mark the message has finalized from a proper ack-message.
      s_messageStatus[messageId] = MessageStatus.ACKNOWLEDGED;

      emit MessageAckReceived(messageId);
    }
  }

  /// @notice Sends the acknowledgement message back through CCIP to original sender contract
  function _sendAck(Client.Any2EVMMessage calldata incomingMessage) internal {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);

    // Build the outgoing ACK message, with no tokens, with data being the concatenation of the acknowledgement header and incoming-messageId
    Client.EVM2AnyMessage memory outgoingMessage = Client.EVM2AnyMessage({
      receiver: incomingMessage.sender,
      data: abi.encode(ACK_MESSAGE_HEADER, incomingMessage.messageId),
      tokenAmounts: tokenAmounts,
      extraArgs: s_chainConfigs[incomingMessage.sourceChainSelector].extraArgsBytes,
      feeToken: address(s_feeToken)
    });

    uint256 feeAmount = IRouterClient(s_ccipRouter).getFee(incomingMessage.sourceChainSelector, outgoingMessage);

    bytes32 ACKmessageId = IRouterClient(s_ccipRouter).ccipSend{
      value: address(s_feeToken) == address(0) ? feeAmount : 0
    }(incomingMessage.sourceChainSelector, outgoingMessage);

    emit MessageSent(incomingMessage.messageId, ACKmessageId);
  }
}
