// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../libraries/Client.sol";
import {CCIPClient} from "./CCIPClient.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title CCIPClientWithACK
/// @notice This contract implements logic for sending and receiving CCIP Messages, as well as responding to incoming messages with an ACK-response pattern. It utilizes CCIPReceiver's defensive patterns by default.
/// @dev ccipSend functionality has been inherited from CCIPClient.sol, and _sendACK() from CCIPReceiverWithACK, so only processMessage() must be overridden to enable full functionality for processing incoming messages for ACK's
contract CCIPClientWithACK is CCIPClient {
  using SafeERC20 for IERC20;

  error CannotAcknowledgeUnsentMessage(bytes32);

  constructor(address router, IERC20 feeToken) CCIPClient(router, feeToken) {}

  /// @notice Implementation of arbitrary logic to be executed when a CCIP message is received
  /// @dev is only invoked by self on CCIPReceive, and should implement arbitrary dapp-specific logic
  function processMessage(Client.Any2EVMMessage calldata message) external virtual override onlySelf {
    (MessagePayload memory payload) = abi.decode(message.data, (MessagePayload));

    if (payload.messageType == MessageType.OUTGOING) {
      // Insert Processing workflow here.

      // If the message was outgoing, then send an ack response.
      _sendAck(message);
    } else if (payload.messageType == MessageType.ACK) {
      // Decode message into the message-header and the messageId to ensure the message is encoded correctly
      (string memory messageHeader, bytes32 messageId) = abi.decode(payload.data, (string, bytes32));

      // Ensure Ack Message contains proper message header
      if (keccak256(abi.encode(messageHeader)) != keccak256(abi.encode(ACK_MESSAGE_HEADER))) {
        revert InvalidAckMessageHeader();
      }

      // Make sure the ACK message was originally sent by this contract
      if (s_messageStatus[messageId] != MessageStatus.SENT) revert CannotAcknowledgeUnsentMessage(messageId);

      // Mark the message has finalized from a proper ack-message.
      s_messageStatus[messageId] = MessageStatus.ACKNOWLEDGED;

      emit MessageAckReceived(messageId);
    }
  }
}
