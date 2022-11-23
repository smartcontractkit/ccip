// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";
import {BaseOnRampInterface} from "./BaseOnRampInterface.sol";
import {DynamicFeeCalculatorInterface} from "../dynamicFeeCalculator/DynamicFeeCalculatorInterface.sol";

interface EVM2EVMGEOnRampInterface is BaseOnRampInterface, DynamicFeeCalculatorInterface {
  event CCIPSendRequested(CCIP.EVM2EVMGEMessage message);

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(
    CCIP.EVM2AnyGEMessage memory message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (uint64);
}
