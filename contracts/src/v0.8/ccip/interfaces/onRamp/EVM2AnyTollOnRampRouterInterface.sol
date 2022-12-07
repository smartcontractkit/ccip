// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Toll} from "../../models/Toll.sol";
import {TollConsumer} from "../../models/TollConsumer.sol";
import {EVM2EVMTollOnRampInterface} from "./EVM2EVMTollOnRampInterface.sol";
import {BaseOnRampRouterInterface} from "./BaseOnRampRouterInterface.sol";

interface EVM2AnyTollOnRampRouterInterface is BaseOnRampRouterInterface {
  error OnRampAlreadySet(uint64 chainId, EVM2EVMTollOnRampInterface onRamp);

  event OnRampSet(uint64 indexed chainId, EVM2EVMTollOnRampInterface indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint64 destinationChainId, TollConsumer.EVM2AnyTollMessage calldata message)
    external
    returns (uint64);

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint64 chainId, EVM2EVMTollOnRampInterface onRamp) external;

  /**
   * @notice Gets the current OnRamp for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint64 chainId) external view returns (EVM2EVMTollOnRampInterface);
}
