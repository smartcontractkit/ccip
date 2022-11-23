// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRampRouter should implement this interface.
 */
interface Any2EVMMessageReceiverInterface {
  /**
   * @notice Called by the OffRampRouter to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMMessage calldata message) external;
}
