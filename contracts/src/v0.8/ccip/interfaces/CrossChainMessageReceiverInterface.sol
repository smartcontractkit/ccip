// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../utils/CCIP.sol";

/**
 * @notice Application contracts that intend to receive messages from
 * the OffRamp should implement this interface.
 */
interface CrossChainMessageReceiverInterface {
  /**
   * @notice Called by the OffRamp to deliver a message
   * @param message CCIP Message
   */
  function receiveMessage(CCIP.Message calldata message) external;
}
