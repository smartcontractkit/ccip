// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../applications/interfaces/CrossChainMessageReceiverInterface.sol";
import "../../utils/interfaces/SubscriptionInterface.sol";
import "./BaseOffRampRouterInterface.sol";

interface Any2EVMSubscriptionOffRampRouterInterface is SubscriptionInterface, BaseOffRampRouterInterface {
  /**
   * @notice Route the message to its intended receiver contract
   * @param receiver Receiver contract implementing CrossChainMessageReceiverInterface
   * @param message CCIP.Any2EVMSubscriptionMessage struct
   */
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMSubscriptionMessage calldata message)
    external;

  /**
   * @notice Charges a subscription
   * @param receiver Receiver address of the subscription that is to be charged
   * @param sender The sender of the cross chain message that is charging
   *          the subscription
   * @param amount The fee amount to be charged
   * @dev should be called from the OffRamp
   */
  function chargeSubscription(
    address receiver,
    address sender,
    uint256 amount
  ) external;
}
