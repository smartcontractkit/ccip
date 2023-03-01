// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "../onRamp/IEVM2AnyOnRamp.sol";
import {IRouterClient} from "./IRouterClient.sol";

import {Client} from "../../models/Client.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IRouter is IRouterClient {
  error OnlyOffRamp();

  event OnRampSet(uint64 indexed destChainId, address onRamp);
  event OffRampAdded(uint64 indexed sourceChainId, address offRamp);
  event OffRampRemoved(uint64 indexed sourceChainId, address offRamp);

  /// @notice Gets the wrapped representation of the native fee coin.
  /// @return The address of the ERC20 wrapped native.
  function getWrappedNative() external view returns (address);

  /// @notice Sets a new wrapped native token.
  /// @param wrappedNative The address of the new wrapped native ERC20 token.
  function setWrappedNative(address wrappedNative) external;

  struct OnRampUpdate {
    uint64 destChainId; // --┐  Destination chain Id.
    address onRamp; // ------┘  OnRamp address that is allowed to use this router.
  }
  struct OffRampUpdate {
    uint64 sourceChainId; //    Source chain Id.
    address[] offRamps; //      List of offRamps that are allowed to use this router.
  }

  /// @notice Set applies a set of ingress and egress config updates.
  /// @dev only callable by owner.
  function applyRampUpdates(OnRampUpdate[] memory onRampUpdates, OffRampUpdate[] memory offRampUpdates) external;

  /// @notice Get a list of offRamps for a source chain.
  function getOffRamps(uint64 sourceChainId) external view returns (address[] memory);

  /// @notice Get the onramp for a destination chain.
  /// @param destChainId The destination chain Id to get the onRamp for.
  /// @return The address of the onRamp.
  function getOnRamp(uint64 destChainId) external view returns (address);

  /// @notice Route the message to its intended receiver contract.
  /// @param message Client.Any2EVMMessage struct.
  /// @param manualExecution bool to indicate manual instead of DON execution.
  /// @param receiver The address of the receiver of the CCIP message.
  /// @dev if the receiver is a contracts that signals support for CCIP execution through EIP-165.
  /// the contract is called. If not, only tokens are transferred.
  /// @return success A boolean value indicating whether the ccip message was received without errors.
  function routeMessage(
    Client.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external returns (bool success);
}
