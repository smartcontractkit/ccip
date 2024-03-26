// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {Client} from "./../libraries/Client.sol";
import {CCIPReceiver} from "./CCIPReceiver.sol";
import {IWrappedNative} from "../interfaces/IWrappedNative.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface CCIPRouter {
  function getWrappedNative() external view returns (address);
}

/// @notice A contract that can send raw ether cross-chain using CCIP.
/// Since CCIP only supports ERC-20 token transfers, this contract accepts
/// normal ether, wraps it, and uses CCIP to send it cross-chain.
/// On the receiving side, the wrapped ether is unwrapped and sent to the final receiver.
contract EtherSenderReceiver is CCIPReceiver, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error CCIPReceiveFailed();
  error InvalidReceiver(bytes receiver);
  error InvalidDestinationEOA(bytes destEOA);
  error InvalidTokenAmounts(uint256 gotAmounts);
  error InvalidWethAddress(address want, address got);
  error TokenAmountNotEqualToMsgValue(uint256 gotAmount, uint256 msgValue);
  error InsufficientMsgValue(uint256 gotAmount, uint256 msgValue);
  error InsufficientFee(uint256 gotFee, uint256 fee);
  error AllowanceTooHigh();

  /// @notice The WETH contract address.
  IWrappedNative public immutable i_weth;

  /// @notice the gas limit for the message call on the destination chain, 500,000 should be plenty.
  uint256 public constant MESSAGE_GAS_LIMIT = 500_000;

  string public constant override typeAndVersion = "EtherSenderReceiver 1.0.0-dev";

  constructor(address router) CCIPReceiver(router) {
    i_weth = IWrappedNative(CCIPRouter(router).getWrappedNative());
  }

  /// @param destinationChainSelector The destination chainSelector
  /// @param message The cross-chain CCIP message including data and/or tokens
  /// @return fee returns execution fee for the message
  /// delivery to destination chain, denominated in the feeToken specified in the message.
  /// @dev Reverts with appropriate reason upon invalid message.
  function getFee(
    uint64 destinationChainSelector,
    Client.EVM2AnyMessage memory message
  ) external view returns (uint256 fee) {
    _validateMessage(message);

    // set the gas limit for the call on destination.
    bytes memory extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1(MESSAGE_GAS_LIMIT));
    message.extraArgs = extraArgs;

    return IRouterClient(getRouter()).getFee(destinationChainSelector, message);
  }

  /// @notice Send raw native tokens cross-chain.
  /// @param destinationChainSelector The destination chain selector.
  /// @param message The CCIP message with the following fields correctly set:
  /// - bytes receiver: The _contract_ address on the destination chain that will receive the wrapped ether.
  /// The caller must ensure that this contract address is correct, otherwise funds may be lost forever.
  /// - bytes data: The abi-encoded EOA that will receive the unwrapped ether on the destination chain.
  /// - address feeToken: The fee token address. Must be address(0) for native tokens, or a supported CCIP fee token otherwise (i.e, LINK token).
  /// In the event a feeToken is set, we will transferFrom the caller the fee amount before sending the message, in order to forward
  /// - EVMTokenAmount[] tokenAmounts: The tokenAmounts array must contain a single element with the following fields:
  ///   - address token: The WETH token address.
  ///   - uint256 amount: The amount of ether to send.
  /// There are a couple of cases here that depend on the fee token specified:
  /// 1. If feeToken == address(0), the fee must be included in msg.value. Therefore tokenAmounts[0].amount must be less than msg.value,
  ///    and the difference will be used as the fee.
  /// 2. If feeToken != address(0), the fee is not included in msg.value, and tokenAmounts[0].amount must be equal to msg.value.
  // these fees to the CCIP router.
  /// @return messageId The CCIP message ID.
  function ccipSend(
    uint64 destinationChainSelector,
    Client.EVM2AnyMessage memory message
  ) external payable returns (bytes32) {
    _validateMessage(message);
    _validateFeeToken(message, msg.value);

    // set the gas limit for the call on destination.
    bytes memory extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1(MESSAGE_GAS_LIMIT));
    message.extraArgs = extraArgs;

    // deposit the ether into the weth contract to get the wrapped ether.
    // approve the router to spend the wrapped ether.
    i_weth.deposit{value: message.tokenAmounts[0].amount}();
    i_weth.approve(getRouter(), message.tokenAmounts[0].amount);

    // get the fee from the router now that we have the full message data.
    // if the fee token is not native, we need to transfer the fee to this contract and re-approve it to the router.
    uint256 fee = IRouterClient(getRouter()).getFee(destinationChainSelector, message);
    if (message.feeToken != address(0)) {
      // Note: its not possible to have any leftover tokens in this path because we transferFrom the exact fee that CCIP
      // requires from the caller.
      IERC20(message.feeToken).safeTransferFrom(msg.sender, address(this), fee);
      IERC20(message.feeToken).safeIncreaseAllowance(getRouter(), fee);

      return IRouterClient(getRouter()).ccipSend(destinationChainSelector, message);
    } else {
      // We've already checked that msg.value > tokenAmounts[0].amount, so we can use the difference as the fee.
      // We don't want to keep any excess ether in this contract, so we send over the entire diff as the fee.
      uint256 diff = msg.value - message.tokenAmounts[0].amount;
      if (diff < fee) {
        revert InsufficientFee(diff, fee);
      }

      return IRouterClient(getRouter()).ccipSend{value: diff}(destinationChainSelector, message);
    }
  }

  function _validateMessage(Client.EVM2AnyMessage memory message) private view {
    // receiver and destination EOA addresses must be correctly specified.
    // abi.decode will revert if the bytes in receiver do not decode to an address.
    address receiver = abi.decode(message.receiver, (address));
    if (receiver == address(0)) {
      revert InvalidReceiver(message.receiver);
    }

    address destEOA = abi.decode(message.data, (address));
    if (destEOA == address(0)) {
      revert InvalidDestinationEOA(message.data);
    }

    // Only one tokenAmount is allowed, which is the weth token and amount.
    if (message.tokenAmounts.length != 1) {
      revert InvalidTokenAmounts(message.tokenAmounts.length);
    }

    Client.EVMTokenAmount memory tokenAmount = message.tokenAmounts[0];
    if (tokenAmount.token != address(i_weth)) {
      revert InvalidWethAddress(address(i_weth), tokenAmount.token);
    }
  }

  function _validateFeeToken(Client.EVM2AnyMessage memory message, uint256 msgValue) private pure {
    Client.EVMTokenAmount memory tokenAmount = message.tokenAmounts[0];

    if (message.feeToken == address(0)) {
      // If the fee token is native, the fee must be included in msgValue.
      if (msgValue > tokenAmount.amount) {
        revert InsufficientMsgValue(tokenAmount.amount, msgValue);
      }
    }
  }

  /// @notice Receive the wrapped ether, unwrap it, and send it to the specified EOA in the data field.
  /// @param message The CCIP message containing the wrapped ether amount and the final receiver.
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    // we receive WETH, unwrap it and send it to the final receiver.
    address receiver = abi.decode(message.data, (address));
    i_weth.withdraw(message.destTokenAmounts[0].amount);
    (bool success, ) = payable(receiver).call{value: message.destTokenAmounts[0].amount}("");
    if (!success) {
      revert CCIPReceiveFailed();
    }
  }
}
