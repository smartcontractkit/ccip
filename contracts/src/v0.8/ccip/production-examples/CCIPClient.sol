pragma solidity ^0.8.0;

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {CCIPReceiverWithACK} from "./CCIPReceiverWithACK.sol";

import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {Client} from "../libraries/Client.sol";

contract CCIPClient is CCIPReceiverWithACK {
  using SafeERC20 for IERC20;

  error InvalidConfig();
  error CannotAcknowledgeUnsentMessage(bytes32);

  /// @notice You can't import CCIPReceiver and Sender due to similar parents so functionality of CCIPSender is duplicated here
  constructor(address router, IERC20 feeToken) CCIPReceiverWithACK(router, feeToken) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "CCIPReceiverWithACK 1.0.0-dev";
  }

  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes memory data,
    address feeToken
  ) public payable validChain(destChainSelector) returns (bytes32 messageId) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: s_chains[destChainSelector],
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_extraArgsBytes[destChainSelector],
      feeToken: feeToken
    });

    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);

    bool sendingFeeTokenNormally;

    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      // Transfer the tokens to pay for tokens in tokenAmounts
      IERC20(tokenAmounts[i].token).safeTransferFrom(msg.sender, address(this), tokenAmounts[i].amount);

      // If they are sending the feeToken through, and its the same as the ack fee token, and also paying in it, then you don't need to approve
      // it at all cause its already set as type(uint).max. You can't use safeIncreaseAllowance() either cause it will overflow the token allowance
      if (tokenAmounts[i].token == feeToken && feeToken != address(0) && feeToken == address(s_feeToken)) {
        sendingFeeTokenNormally = true;
        IERC20(tokenAmounts[i].token).safeTransferFrom(msg.sender, address(this), fee);
      }
      // If they're not sending the fee token, then go ahead and approve
      else {
        IERC20(tokenAmounts[i].token).safeApprove(i_ccipRouter, tokenAmounts[i].amount);
      }
    }

    // Since the fee token was already set in the ReceiverWithACK parent, we don't need to approve it to spend, only to ensure we have enough
    // funds for the transfer
    if (!sendingFeeTokenNormally && feeToken == address(s_feeToken) && feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), fee);
    } else if (feeToken == address(0) && msg.value < fee) {
      revert IRouterClient.InsufficientFeeTokenAmount();
    }

    messageId =
      IRouterClient(i_ccipRouter).ccipSend{value: feeToken == address(0) ? fee : 0}(destChainSelector, message);

    s_messageStatus[messageId] = CCIPReceiverWithACK.MessageStatus.SENT;

    // Since the message was outgoing, and not ACK, use bytes32(0) to reflect this
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
      // Decode message into the magic-bytes and the messageId to ensure the message is encoded correctly
      (bytes memory magicBytes, bytes32 messageId) = abi.decode(payload.data, (bytes, bytes32));

      // Ensure Ack Message contains proper magic-bytes
      if (keccak256(magicBytes) != keccak256(ACKMESSAGEMAGICBYTES)) revert InvalidMagicBytes();

      // Make sure the ACK message was originally sent by this contract
      if (s_messageStatus[messageId] != MessageStatus.SENT) revert CannotAcknowledgeUnsentMessage(messageId);

      // Mark the message has finalized from a proper ack-message.
      s_messageStatus[messageId] = MessageStatus.ACKNOWLEDGED;

      emit MessageAckReceived(messageId);
    }
  }
}
