// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../../vendor/SafeERC20.sol";
import "../../utils/Subscription.sol";
import "../interfaces/Any2EVMSubscriptionOffRampRouterInterface.sol";
import "../BaseOffRampRouter.sol";

contract Any2EVMSubscriptionOffRampRouter is
  BaseOffRampRouter,
  Any2EVMSubscriptionOffRampRouterInterface,
  Subscription,
  TypeAndVersionInterface
{
  string public constant override typeAndVersion = "Any2EVMSubscriptionOffRampRouter 1.0.0";

  constructor(
    BaseOffRampInterface[] memory offRamps,
    SubscriptionInterface.SubscriptionConfig memory subscriptionConfig
  ) BaseOffRampRouter(offRamps) Subscription(subscriptionConfig) {}

  /// @inheritdoc Any2EVMSubscriptionOffRampRouterInterface
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMSubscriptionMessage calldata message)
    external
    override
    onlyOffRamp
  {
    try receiver.ccipReceive(message) {} catch (bytes memory reason) {
      // TODO: use RouterResults and exact gas
      revert MessageFailure(message.sequenceNumber, reason);
    }
  }

  /// @inheritdoc Any2EVMSubscriptionOffRampRouterInterface
  function chargeSubscription(
    address receiver,
    address sender,
    uint256 amount
  ) public onlyOffRamp {
    OffRampSubscription memory subscription = s_subscriptions[receiver];
    if (subscription.balance < amount) {
      revert BalanceTooLow();
    }
    address[] memory allowedSenders = subscription.senders;
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
