// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../utils/CCIP.sol";
import "../../utils/interfaces/SubscriptionManagerInterface.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRamp should implement this interface.
 */
interface Any2EVMMessageReceiverInterface is SubscriptionManagerInterface {
  /**
   * @notice Called by the OffRamp to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMMessage calldata message) external;
}
