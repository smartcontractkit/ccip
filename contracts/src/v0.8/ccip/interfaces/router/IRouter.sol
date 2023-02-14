// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "../onRamp/IEVM2AnyOnRamp.sol";
import {IRouterClient} from "./IRouterClient.sol";

import {Client} from "../../models/Client.sol";
import {Client} from "../../models/Client.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IRouter is IRouterClient {
  error OnlyOffRamp();

  event OnRampSet(uint64 indexed destChainId, address onRamp);
  event OffRampAdded(uint64 indexed sourceChainId, address offRamp);
  event OffRampRemoved(uint64 indexed sourceChainId, address offRamp);

  function getWrappedNative() external view returns (address);

  function setWrappedNative(address wrappedNative) external;

  struct OnRampUpdate {
    uint64 destChainId;
    address onRamp;
  }
  struct OffRampUpdate {
    uint64 sourceChainId;
    address[] offRamps;
  }

  /// @notice Set applies a set of ingress and egress config updates.
  /// @dev only callable by owner
  function applyRampUpdates(OnRampUpdate[] memory onRampUpdates, OffRampUpdate[] memory offRampUpdates) external;

  // @notice Get a list of offramps for a source chain.
  function getOffRamps(uint64 sourceChainId) external view returns (address[] memory);

  // @notice Get the onramp for a destination chain.
  function getOnRamp(uint64 destChainId) external view returns (address);

  /// @notice Route the message to its intended receiver contract
  /// @param message Client.Any2EVMMessage struct
  /// @param manualExecution bool to indicate manual instead of DON execution
  function routeMessage(
    Client.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external returns (bool success);
}
