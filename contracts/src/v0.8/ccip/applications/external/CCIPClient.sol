// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";

import {Client} from "../../libraries/Client.sol";
import {CCIPReceiverWithACK} from "./CCIPReceiverWithACK.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title CCIPClient
/// @notice This contract implements logic for sending and receiving CCIP Messages. It utilizes CCIPReceiver's defensive patterns by default.
/// @dev CCIPReceiver and CCIPSender cannot be simultaneously imported due to similar parents so CCIPSender functionality has been duplicated
// TODO make CCIPClient inherit from CCIPReceiver
contract CCIPClient is CCIPReceiverWithACK {
  using SafeERC20 for IERC20;

  constructor(address router, IERC20 feeToken) CCIPReceiverWithACK(router, feeToken) {}

  /// @notice sends a message through CCIP to the router
  // TODO really beef up the comments here
  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes memory data
  ) public payable virtual isValidChain(destChainSelector) returns (bytes32 messageId) {
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
        IERC20(tokenAmounts[i].token).safeApprove(s_ccipRouter, tokenAmounts[i].amount);
      }
    }

    uint256 fee = IRouterClient(s_ccipRouter).getFee(destChainSelector, message);

    // Additional tokens for fees do not need to be approved to the router since it is already handled by setting s_feeToken
    if (address(s_feeToken) != address(0)) {
      IERC20(s_feeToken).safeTransferFrom(msg.sender, address(this), fee);
    }

    // TODO comment we only have messageId after calling ccipSend, so brekaing CEI is necessary
    // messageId clac lives in OnRamp, which can be upgradaed, this it should be abstracted away from client impl
    messageId = IRouterClient(s_ccipRouter).ccipSend{value: address(s_feeToken) == address(0) ? fee : 0}(
      destChainSelector, message
    );

    s_messageStatus[messageId] = CCIPReceiverWithACK.MessageStatus.SENT;

    // Since the message was outgoing, and not ACK, reflect this with bytes32(0)
    emit MessageSent(messageId, bytes32(0));

    return messageId;
  }
}
