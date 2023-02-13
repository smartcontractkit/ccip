// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "../onRamp/IEVM2AnyOnRamp.sol";

import {Consumer} from "../../models/Consumer.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IRouterClient {
  error UnsupportedDestinationChain(uint64 destinationChainId);
  error SenderNotAllowed(address sender);
  error InsufficientFeeTokenAmount();
  error InvalidMsgValue();

  /// @notice Request a message to be sent to the destination chain
  /// @param destinationChainId The destination chain ID
  /// @param message The message payload
  /// @return The message ID
  /// @dev Note if msg.value is larger than the required fee (from getFee) we accept
  /// the overpayment with no refund.
  function ccipSend(uint64 destinationChainId, Consumer.EVM2AnyMessage calldata message)
    external
    payable
    returns (bytes32);

  /// @param destinationChainId The destination chain ID
  /// @param message The message payload
  /// @return fee returns guaranteed execution fee for the specified message
  /// delivery to destination chain
  function getFee(uint64 destinationChainId, Consumer.EVM2AnyMessage memory message)
    external
    view
    returns (uint256 fee);

  /**
   * @notice Checks if the given destination chain ID is supported
   * @param chainId The destination chain to check
   * @return supported is true if it is supported, false if not
   */
  function isChainSupported(uint64 chainId) external view returns (bool supported);

  /// @notice Gets the current OnRamp for the destination chain ID
  /// @param chainId chain ID to get ramp details for
  /// @return onRamp
  function getOnRamp(uint64 chainId) external view returns (IEVM2AnyOnRamp);

  /// @notice Gets a list of all supported source chain tokens for a given
  /// destination chain.
  /// @param destChainId The destination chain Id
  /// @return tokens The addresses of all tokens that are supported.
  function getSupportedTokens(uint64 destChainId) external view returns (address[] memory tokens);
}
