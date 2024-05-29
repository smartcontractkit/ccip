pragma solidity ^0.8.0;

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {CCIPReceiver} from "./CCIPReceiver.sol";
import {CCIPSender} from "./CCIPSender.sol";

import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {Client} from "../libraries/Client.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableMap.sol";

contract CCIPClient is CCIPReceiver {
  using SafeERC20 for IERC20;

  error InvalidConfig();

  event MessageSent(bytes32 messageId);
  event MessageReceived(bytes32 messageId);

  // Current feeToken
  IERC20 public s_feeToken;

  /// @notice You can't import CCIPReceiver and Sender due to similar parents so functionality of CCIPSender is duplicated here
  constructor(address router, IERC20 feeToken) CCIPReceiver(router) {
    s_feeToken = feeToken;
    s_feeToken.approve(address(router), type(uint256).max);
  }

  /// @notice sends data to receiver on dest chain. Assumes address(this) has sufficient native asset.
  function sendDataPayNative(
    uint64 destChainSelector,
    bytes memory receiver,
    bytes memory data
  ) external validChain(destChainSelector) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_chains[destChainSelector],
      feeToken: address(0) // We leave the feeToken empty indicating we'll pay raw native.
    });
    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend{
      value: IRouterClient(i_ccipRouter).getFee(destChainSelector, message)
    }(destChainSelector, message);
    emit MessageSent(messageId);
  }

  /// @notice sends data to receiver on dest chain. Assumes address(this) has sufficient feeToken.
  function sendDataPayFeeToken(
    uint64 destChainSelector,
    bytes memory receiver,
    bytes memory data
  ) external validChain(destChainSelector) {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](0);
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_chains[destChainSelector],
      feeToken: address(s_feeToken)
    });
    // Optional uint256 fee = i_ccipRouter.getFee(destChainSelector, message);
    // Can decide if fee is acceptable.
    // address(this) must have sufficient feeToken or the send will revert.
    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  /// @notice sends data to receiver on dest chain. Assumes address(this) has sufficient native token.
  function sendDataAndTokens(
    uint64 destChainSelector,
    bytes memory receiver,
    bytes memory data,
    Client.EVMTokenAmount[] memory tokenAmounts
  ) external validChain(destChainSelector) {
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_chains[destChainSelector],
      feeToken: address(s_feeToken)
    });
    // Optional uint256 fee = i_ccipRouter.getFee(destChainSelector, message);
    // Can decide if fee is acceptable.
    // address(this) must have sufficient feeToken or the send will revert.
    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  // @notice user sends tokens to a receiver
  // Approvals can be optimized with a whitelist of tokens and inf approvals if desired.
  function sendTokens(
    uint64 destChainSelector,
    bytes memory receiver,
    Client.EVMTokenAmount[] memory tokenAmounts
  ) external validChain(destChainSelector) {
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }
    bytes memory data;
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_chains[destChainSelector],
      feeToken: address(s_feeToken)
    });
    // Optional uint256 fee = i_ccipRouter.getFee(destChainSelector, message);
    // Can decide if fee is acceptable.
    // address(this) must have sufficient feeToken or the send will revert.
    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend(destChainSelector, message);
    emit MessageSent(messageId);
  }

  
}
