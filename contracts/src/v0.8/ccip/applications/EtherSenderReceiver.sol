// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {CCIPReceiver} from "./CCIPReceiver.sol";
import {Client} from "./../libraries/Client.sol";
import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
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
/// @notice This contract only supports chains where the wrapped native contract
/// is the WETH contract (i.e not WMATIC, or WAVAX, etc.). This is because the
/// receiving contract will always unwrap the ether using it's local wrapped native contract.
/// @dev This contract is both a sender and a receiver. This same contract can be
/// deployed on source and destination chains to facilitate cross-chain ether transfers
/// and act as a sender and a receiver.
/// @dev This contract is intentionally ownerless and permissionless. This contract
/// will never hold any excess funds, native or otherwise.
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

  string public constant override typeAndVersion = "EtherSenderReceiver 1.0.0-dev";

  /// @notice The WETH contract address.
  /// This matches the wrapped ether contract set on the CCIP router by construction.
  IWrappedNative public immutable i_weth;

  /// @notice the gas limit for the message call on the destination chain, 500,000 should be plenty.
  /// @dev This won't vary on L2's, callbacks are always provided so-called "L2 gas".
  uint256 public constant MESSAGE_GAS_LIMIT = 250_000;

  /// @param router The CCIP router address.
  constructor(address router) CCIPReceiver(router) {
    i_weth = IWrappedNative(CCIPRouter(router).getWrappedNative());
  }

  /// @notice Need this in order to unwrap correctly.
  receive() external payable {}

  /// @notice Get the fee for sending a message to a destination chain.
  /// This is mirrored from the router for convenience, construct the appropriate
  /// message and get it's fee.
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

    message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1(MESSAGE_GAS_LIMIT));

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
    _validateFeeToken(message);

    message.extraArgs = Client._argsToBytes(Client.EVMExtraArgsV1(MESSAGE_GAS_LIMIT));

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
      // We've already checked that msg.value > tokenAmounts[0].amount, so there is no overflow/underflow risk.
      // We don't want to keep any excess ether in this contract, so we send over the entire diff as the fee.
      uint256 diff = msg.value - message.tokenAmounts[0].amount;
      if (diff < fee) {
        revert InsufficientFee(diff, fee);
      }

      return IRouterClient(getRouter()).ccipSend{value: diff}(destinationChainSelector, message);
    }
  }

  function _validateMessage(Client.EVM2AnyMessage memory message) private view {
    // receiver is already checked by ccip, so don't need to duplicate that check.
    // destination EOA address must be correctly specified in message.data.
    // abi.decode will revert if the bytes in receiver do not decode to an address.
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

  function _validateFeeToken(Client.EVM2AnyMessage memory message) private view {
    uint256 tokenAmount = message.tokenAmounts[0].amount;

    if (message.feeToken == address(0)) {
      // If the fee token is native, the fee must be included in msg.value.
      if (msg.value <= tokenAmount) {
        revert InsufficientMsgValue(tokenAmount, msg.value);
      }
    } else {
      // If the fee token is NOT native, then the token amount must be equal to msg.value.
      // This is done to ensure that there is no leftover ether in this contract.
      if (msg.value != tokenAmount) {
        revert TokenAmountNotEqualToMsgValue(tokenAmount, msg.value);
      }
    }
  }

  /// @notice Receive the wrapped ether, unwrap it, and send it to the specified EOA in the data field.
  /// @param message The CCIP message containing the wrapped ether amount and the final receiver.
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    // we receive WETH, unwrap it and send it to the final receiver.
    // The code below should never revert if the message being is valid according
    // to the above _validateMessage and _validateFeeToken functions.

    // decode the receiver from the message data.
    address receiver = abi.decode(message.data, (address));

    // withdraw the WETH received from the token pool.
    uint256 tokenAmount = message.destTokenAmounts[0].amount;
    i_weth.withdraw(tokenAmount);

    // it is possible that the below call may fail if receiver.code.length > 0 and the contract
    // doesn't e.g have a receive() or a fallback() function.
    (bool success, ) = payable(receiver).call{value: tokenAmount}("");
    if (!success) {
      // We have a few options here:
      // 1. Revert: this is bad generally because it may mean that these tokens are stuck.
      // 2. Store the tokens in a mapping and allow the user to withdraw them with another tx.
      // 3. Send weth to the receiver address.
      // We opt for (3) here because at least the receiver will have the funds and can unwrap them if needed.
      // However it is worth noting that if receiver is actually a contract AND the contract _cannot_ withdraw
      // the WETH, then the WETH will be stuck in this contract.
      i_weth.deposit{value: tokenAmount}();
      i_weth.transfer(receiver, tokenAmount);
    }
  }
}
