// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IBaseOnRampRouter} from "../onRamp/IBaseOnRampRouter.sol";
import {IEVM2AnyOnRamp} from "../onRamp/IEVM2AnyOnRamp.sol";
import {IAny2EVMOffRampRouter} from "../offRamp/IAny2EVMOffRampRouter.sol";
import {IBaseOnRampRouter} from "../onRamp/IBaseOnRampRouter.sol";

import {Consumer} from "../../models/Consumer.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IRouter is IBaseOnRampRouter, IAny2EVMOffRampRouter {
  /// @notice This error is thrown when trying to set the same on-ramp for a
  /// destination chain
  /// @param chainId chain ID of the destination chain where the on-ramp exists
  /// @param onRamp The existing on-ramp
  error OnRampAlreadySet(uint64 chainId, IEVM2AnyOnRamp onRamp);

  /// @notice This event is emitted when an on-ramp is set for a destination
  /// chain ID
  /// @param chainId chain ID of the destination chain where the on-ramp was set
  /// @param onRamp On-ramp to use for that destination chain
  event OnRampSet(uint64 indexed chainId, IEVM2AnyOnRamp indexed onRamp);

  /// @notice Request a message to be sent to the destination chain
  /// @param destinationChainId The destination chain ID
  /// @param message The message payload
  /// @return The message ID
  function ccipSend(uint64 destinationChainId, Consumer.EVM2AnyMessage calldata message) external returns (bytes32);

  /// @param destinationChainId The destination chain ID
  /// @param message The message payload
  /// @return fee returns guaranteed execution fee for the specified message
  /// delivery to destination chain
  function getFee(uint64 destinationChainId, Consumer.EVM2AnyMessage memory message)
    external
    view
    returns (uint256 fee);

  /// @notice Set chainId => onRamp mapping
  /// @dev only callable by owner
  /// @param chainId destination chain ID
  /// @param onRamp OnRamp to use for that destination chain
  function setOnRamp(uint64 chainId, IEVM2AnyOnRamp onRamp) external;

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
