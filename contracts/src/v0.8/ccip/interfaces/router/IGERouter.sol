// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IBaseOnRampRouter} from "../onRamp/IBaseOnRampRouter.sol";
import {IEVM2AnyGEOnRamp} from "../onRamp/IEVM2AnyGEOnRamp.sol";
import {IAny2EVMOffRampRouter} from "../offRamp/IAny2EVMOffRampRouter.sol";
import {IBaseOnRampRouter} from "../onRamp/IBaseOnRampRouter.sol";

import {GEConsumer} from "../../models/GEConsumer.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IGERouter is IBaseOnRampRouter, IAny2EVMOffRampRouter {
  error OnRampAlreadySet(uint64 chainId, IEVM2AnyGEOnRamp onRamp);

  event OnRampSet(uint64 indexed chainId, IEVM2AnyGEOnRamp indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage calldata message) external returns (bytes32);

  function getFee(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message)
    external
    view
    returns (uint256 fee);

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint64 chainId, IEVM2AnyGEOnRamp onRamp) external;

  /**
   * @notice Gets the current OnRamp for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint64 chainId) external view returns (IEVM2AnyGEOnRamp);

  /**
   * @notice Gets a list of all supported source chain tokens for a given
   *  destination chain.
   * @param destChainId The destination chain Id
   * @return tokens The addresses of all tokens that are supported.
   */
  function getSupportedTokens(uint64 destChainId) external view returns (address[] memory tokens);
}
