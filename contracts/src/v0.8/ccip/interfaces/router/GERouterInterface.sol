// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";
import {BaseOnRampRouterInterface} from "../onRamp/BaseOnRampRouterInterface.sol";
import {EVM2EVMGEOnRampInterface} from "../onRamp/EVM2EVMGEOnRampInterface.sol";
import {Any2EVMOffRampRouterInterface} from "../offRamp/Any2EVMOffRampRouterInterface.sol";

interface GERouterInterface is BaseOnRampRouterInterface, Any2EVMOffRampRouterInterface {
  error OnRampAlreadySet(uint256 chainId, EVM2EVMGEOnRampInterface onRamp);

  event OnRampSet(uint256 indexed chainId, EVM2EVMGEOnRampInterface indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyGEMessage calldata message) external returns (uint64);

  function getFee(uint256 destinationChainId, CCIP.EVM2AnyGEMessage memory message) external view returns (uint256 fee);

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint256 chainId, EVM2EVMGEOnRampInterface onRamp) external;

  /**
   * @notice Gets the current OnRamp for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint256 chainId) external view returns (EVM2EVMGEOnRampInterface);
}
