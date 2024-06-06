pragma solidity ^0.8.0;

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {CCIPReceiverWithACK, CCIPReceiver} from "./CCIPReceiverWithACK.sol";
import {CCIPSender} from "./CCIPSender.sol";
import {CCIPClientBase} from "./CCIPClientBase.sol";

import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {Client} from "../libraries/Client.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableMap.sol";

contract CCIPClient is CCIPReceiverWithACK {
  using SafeERC20 for IERC20;

  error InvalidConfig();

  /// @notice You can't import CCIPReceiver and Sender due to similar parents so functionality of CCIPSender is duplicated here
  constructor(address router, IERC20 feeToken) CCIPReceiverWithACK(router, feeToken) {}

  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes calldata data,
    address feeToken
  ) public payable validChain(destChainSelector) {

    // TODO: Decide whether workflow should assume contract is funded with tokens to send already
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }

    CCIPClientBase.Chain memory chainInfo = s_chains[destChainSelector];

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: chainInfo.recipient,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: chainInfo.extraArgsBytes,
      feeToken: feeToken
    });

    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);

    // Transfer fee token from sender and approve router to pay for message
    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), fee);
      IERC20(feeToken).safeApprove(i_ccipRouter, fee);
    }

    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend{
      value: feeToken == address(0) ? fee : 0
    } (destChainSelector, message);

    emit MessageSent(messageId);
  }
}
