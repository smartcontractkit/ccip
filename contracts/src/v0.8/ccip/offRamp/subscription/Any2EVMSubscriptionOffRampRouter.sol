// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../subscription/Subscription.sol";
import "../BaseOffRampRouter.sol";

contract Any2EVMSubscriptionOffRampRouter is BaseOffRampRouter, Subscription, TypeAndVersionInterface {
  string public constant override typeAndVersion = "Any2EVMSubscriptionOffRampRouter 1.0.0";

  constructor(
    BaseOffRampInterface[] memory offRamps,
    SubscriptionInterface.SubscriptionConfig memory subscriptionConfig
  ) BaseOffRampRouter(offRamps) Subscription(subscriptionConfig) {}

  /**
   * @notice Charges a subscription
   * @param receiver Receiver address of the subscription that is to be charged
   * @param sender The sender of the cross chain message that is charging the subscription
   * @param amount The fee amount to be charged
   * @dev should be called from the OffRamp
   */
  function chargeSubscription(
    address receiver,
    address sender,
    uint256 amount
  ) public onlyOffRamp {
    address[] memory allowedSenders = s_subscriptions[receiver].senders;
    for (uint256 i = 0; i < allowedSenders.length; ++i) {
      if (allowedSenders[i] == sender) {
        s_subscriptions[receiver].balance -= amount;
        emit SubscriptionCharged(receiver, amount);
        return;
      }
    }
    revert SenderNotAllowed(sender);
  }
}
