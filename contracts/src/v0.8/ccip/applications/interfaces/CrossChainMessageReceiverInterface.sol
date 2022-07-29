// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../utils/CCIP.sol";
import "../../utils/interfaces/SubscriptionManagerInterface.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRamp should implement this interface.
 */
interface CrossChainMessageReceiverInterface is SubscriptionManagerInterface {
  /**
   * @notice Called by the OffRamp to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMSubscriptionMessage calldata message) external;

  /**
   * @notice Called by the OffRamp to deliver a message
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external;
}
