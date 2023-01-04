// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {Toll} from "../../models/Toll.sol";
import {TollConsumer} from "../../models/TollConsumer.sol";
import {IEVM2EVMTollOnRamp} from "./IEVM2EVMTollOnRamp.sol";
import {IBaseOnRampRouter} from "./IBaseOnRampRouter.sol";

interface IEVM2AnyTollOnRampRouter is IBaseOnRampRouter {
  error OnRampAlreadySet(uint64 chainId, IEVM2EVMTollOnRamp onRamp);

  event OnRampSet(uint64 indexed chainId, IEVM2EVMTollOnRamp indexed onRamp);

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
  function setOnRamp(uint64 chainId, IEVM2EVMTollOnRamp onRamp) external;

  /**
   * @notice Gets the current OnRamp for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint64 chainId) external view returns (IEVM2EVMTollOnRamp);
}
