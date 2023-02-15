// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../models/Client.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the router should implement this interface.
 */
interface IAny2EVMMessageReceiver {
  /**
   * @notice Called by the OffRampRouter to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(Client.Any2EVMMessage calldata message) external;
}
