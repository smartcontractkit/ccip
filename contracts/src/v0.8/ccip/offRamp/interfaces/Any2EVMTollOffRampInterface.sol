// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/Any2EVMTollOffRampRouterInterface.sol";
import "./BaseOffRampInterface.sol";

interface Any2EVMTollOffRampInterface is BaseOffRampInterface {
  error MissingFeeCoinPrice(address feeCoin);
  error InsufficientFeeAmount(uint256 sequenceNumber);

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner should be able to call this function
   */
  function setRouter(Any2EVMTollOffRampRouterInterface router) external;

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external;

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param needFee Whether or not the executor requires a fee
   */
  function execute(CCIP.ExecutionReport memory report, bool needFee) external;

  /**
   * @notice Execute a single message
   * @param message The Any2EVMTollMessage message that will be executed
   * @dev this can only be called by the contract itself. It is part of the
   *       Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMTollMessage memory message) external;
}
