// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";

import {Client} from "../../libraries/Client.sol";
import {CCIPReceiverWithACK} from "./CCIPReceiverWithACK.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice CCIPReceiver and CCIPSender cannot be simultaneously imported due to similar parents so CCIPSender functionality has been duplicated
contract CCIPClient is CCIPReceiverWithACK {
  using SafeERC20 for IERC20;

  error InvalidConfig();
  error CannotAcknowledgeUnsentMessage(bytes32);

  constructor(address router, IERC20 feeToken) CCIPReceiverWithACK(router, feeToken) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CCIPClient 1.0.0-dev";
  }

  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes memory data
  ) public payable isValidChain(destChainSelector) returns (bytes32 messageId) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: s_chainConfigs[destChainSelector].recipient,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_chainConfigs[destChainSelector].extraArgsBytes,
      feeToken: address(s_feeToken)
    });

    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      // Transfer the tokens to pay for tokens in tokenAmounts
      IERC20(tokenAmounts[i].token).safeTransferFrom(msg.sender, address(this), tokenAmounts[i].amount);

      // Do not approve the tokens if it is the feeToken, otherwise the approval amount may overflow
      if (tokenAmounts[i].token != address(s_feeToken)) {
        IERC20(tokenAmounts[i].token).safeIncreaseAllowance(i_ccipRouter, tokenAmounts[i].amount);
      }
    }

    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);

    // Additional tokens for fees do not need to be approved to the router since it is already handled by setting s_feeToken
    if (address(s_feeToken) != address(0)) {
      IERC20(s_feeToken).safeTransferFrom(msg.sender, address(this), fee);
    }

    messageId = IRouterClient(i_ccipRouter).ccipSend{value: address(s_feeToken) == address(0) ? fee : 0}(
      destChainSelector, message
    );

    s_messageStatus[messageId] = CCIPReceiverWithACK.MessageStatus.SENT;

    // Since the message was outgoing, and not ACK, reflect this with bytes32(0)
    emit MessageSent(messageId, bytes32(0));

    return messageId;
  }

  /// CCIPReceiver processMessage to make easier to modify
  /// @notice function requres that
  function processMessage(Client.Any2EVMMessage calldata message) external virtual override onlySelf {
    (MessagePayload memory payload) = abi.decode(message.data, (MessagePayload));

    if (payload.messageType == MessageType.OUTGOING) {
      // Insert Processing workflow here.

      // If the message was outgoing, then send an ack response.
      _sendAck(message);
    } else if (payload.messageType == MessageType.ACK) {
      // Decode message into the message-heacder and the messageId to ensure the message is encoded correctly
      (bytes memory messageHeader, bytes32 messageId) = abi.decode(payload.data, (bytes, bytes32));

      // Ensure Ack Message contains proper message header
      if (keccak256(messageHeader) != keccak256(ACK_MESSAGE_HEADER)) revert InvalidAckMessageHeader();

      // Make sure the ACK message was originally sent by this contract
      if (s_messageStatus[messageId] != MessageStatus.SENT) revert CannotAcknowledgeUnsentMessage(messageId);

      // Mark the message has finalized from a proper ack-message.
      s_messageStatus[messageId] = MessageStatus.ACKNOWLEDGED;

      emit MessageAckReceived(messageId);
    }
  }
}
