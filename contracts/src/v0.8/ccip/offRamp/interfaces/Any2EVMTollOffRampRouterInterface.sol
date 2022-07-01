// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../applications/interfaces/CrossChainMessageReceiverInterface.sol";
import "../interfaces/Any2EVMTollOffRampInterface.sol";
import "./BaseOffRampRouterInterface.sol";

interface Any2EVMTollOffRampRouterInterface is BaseOffRampRouterInterface {
  /**
   * @notice Route the message to its intended receiver contract
   * @param receiver Receiver contract implementing CrossChainMessageReceiverInterface
   * @param message CCIP.Any2EVMTollMessage struct
   */
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage calldata message) external;
}
