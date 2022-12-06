// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../../models/Common.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRampRouter should implement this interface.
 */
interface Any2EVMMessageReceiverInterface {
  /**
   * @notice Called by the OffRampRouter to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(Common.Any2EVMMessage calldata message) external;
}
