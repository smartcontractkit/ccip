// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "../onRamp/IEVM2AnyOnRamp.sol";

import {Consumer} from "../../models/Consumer.sol";
import {Common} from "../../models/Common.sol";

import {IERC20} from "../../../vendor/IERC20.sol";
import {IRouterClient} from "./IRouterClient.sol";

interface IRouter is IRouterClient {
  error NoOffRampsConfigured();
  error MustCallFromOffRamp(address sender);
  error InvalidAddress();
  error OffRampNotAllowed(address offRamp);
  error AlreadyConfigured(address offRamp);

  event OffRampAdded(address indexed offRamp);
  event OffRampRemoved(address indexed offRamp);

  /**
   * @notice Owner can add an offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to add
   */
  function addOffRamp(address offRamp) external;

  /**
   * @notice Owner can remove a specific offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to remove
   */
  function removeOffRamp(address offRamp) external;

  /**
   * @notice Gets all configured offRamps.
   * @return offRamps The offRamp that are configured.
   */
  function getOffRamps() external view returns (address[] memory offRamps);

  function getWrappedNative() external view returns (address);

  function setWrappedNative(address wrappedNative) external;

  /**
   * @notice Returns whether the given offRamp is set to be allowed
   * @param offRamp The offRamp to check.
   * @return allowed True if the offRamp is allowed, false if not.
   */
  function isOffRamp(address offRamp) external view returns (bool allowed);

  /**
   * @notice Route the message to its intended receiver contract
   * @param message Common.Any2EVMMessage struct
   * @param manualExecution bool to indicate manual instead of DON execution
   */
  function routeMessage(
    Common.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external returns (bool success);

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

  /// @notice Set chainId => onRamp mapping
  /// @dev only callable by owner
  /// @param chainId destination chain ID
  /// @param onRamp OnRamp to use for that destination chain
  function setOnRamp(uint64 chainId, IEVM2AnyOnRamp onRamp) external;
}
