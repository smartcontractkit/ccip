// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Any2EVMMOOffRampRouterInterface.sol";

interface Any2EVMMOOffRampInterface is BaseOffRampInterface {
  error IncorrectNonce(uint64 nonce);

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner should be able to call this function
   */
  function setRouter(Any2EVMMOOffRampRouterInterface router) external;

  /**
   * @notice ccipReceive implements the receive function to create a collision
   *          if some other method happens to hash to the same signature
   */
  function ccipReceive(CCIP.Any2EVMMOMessage calldata message) external;

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param manualExecution Whether or not it is manual or DON execution
   */
  function execute(CCIP.ExecutionReport memory report, bool manualExecution) external;

  /**
   * @notice Execute a single message
   * @param message The Any2EVMSubscriptionMessage message that will be executed
   * @dev this can only be called by the contract itself. It is part of the
   *       Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMMOMessage memory message) external;
}
