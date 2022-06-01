// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface OffRampRouterInterface {
  /**
   * @notice Route the message to its intended receiver contract
   * @param receiver Receiver contract implementing CrossChainMessageReceiverInterface
   * @param message CCIP.Message struct
   */
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.AnyToEVMTollMessage calldata message)
    external;
}
