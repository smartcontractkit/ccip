// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "./interfaces/SubscriptionManagerInterface.sol";
import "./interfaces/SubscriptionInterface.sol";

contract Subscription is SubscriptionInterface {
  using Address for address;
  using SafeERC20 for IERC20;

  // The subscription config
  SubscriptionConfig s_config;

  // A mapping from receiver to subscription
  mapping(address => OffRampSubscription) internal s_subscriptions;
  // A mapping from receiver to a prepared withdrawal
  mapping(address => PreparedWithdrawal) internal s_preparedWithdrawals;
  // A mapping from receiver to a prepared sender change
  mapping(address => PreparedNewSenders) internal s_preparedNewSenders;

  constructor(SubscriptionConfig memory config) {
    s_config = config;
  }

  /// @inheritdoc SubscriptionInterface
  function getSupportedTokensForExecutionFee() external view returns (address[] memory) {
    address[] memory supportedTokens = new address[](1);
    supportedTokens[0] = address(s_config.feeToken);
    return supportedTokens;
  }

  /// @inheritdoc SubscriptionInterface
  function getSubscription(address receiver) public view returns (OffRampSubscription memory) {
    return s_subscriptions[receiver];
  }

  /// @inheritdoc SubscriptionInterface
  function getFeeToken() public view returns (IERC20) {
    return s_config.feeToken;
  }

  /// @inheritdoc SubscriptionInterface
  function createSubscription(OffRampSubscription memory subscription)
    external
    onlySubscriptionManager(address(subscription.receiver))
  {
    address receiver = address(subscription.receiver);
    if (address(s_subscriptions[receiver].receiver) != address(0)) {
      revert SubscriptionAlreadyExists();
    }
    s_subscriptions[receiver] = subscription;

    if (subscription.balance > 0) {
      s_config.feeToken.safeTransferFrom(msg.sender, address(this), subscription.balance);
    }

    emit SubscriptionCreated(receiver);
  }

  /// @inheritdoc SubscriptionInterface
  function fundSubscription(address receiver, uint256 amount) external {
    if (amount <= 0) {
      revert FundingAmountNotPositive();
    }
    s_subscriptions[receiver].balance += amount;
    s_config.feeToken.safeTransferFrom(msg.sender, address(this), amount);

    emit SubscriptionFunded(receiver, amount);
  }

  /// @inheritdoc SubscriptionInterface
  function prepareSetSubscriptionSenders(address receiver, address[] memory newSenders)
    external
    onlySubscriptionManager(receiver)
  {
    s_preparedNewSenders[receiver] = PreparedNewSenders({
      newSenders: newSenders,
      timestamp: block.timestamp + s_config.setSubscriptionSenderDelay
    });

    emit PreparedSetSenders(receiver, newSenders);
  }

  /// @inheritdoc SubscriptionInterface
  function setSubscriptionSenders(address receiver, address[] memory newSenders)
    external
    onlySubscriptionManager(receiver)
  {
    PreparedNewSenders memory prepared = s_preparedNewSenders[receiver];
    if (prepared.timestamp > block.timestamp) {
      revert DelayNotPassedYet(prepared.timestamp);
    }

    for (uint256 i = 0; i < newSenders.length; ++i) {
      if (newSenders[i] != prepared.newSenders[i]) {
        revert AddressMismatch(prepared.newSenders[i], newSenders[i]);
      }
    }
    s_subscriptions[receiver].senders = newSenders;

    delete s_preparedNewSenders[receiver];
    emit NewSendersSet(receiver, newSenders);
  }

  /// @inheritdoc SubscriptionInterface
  function prepareWithdrawal(address receiver, uint256 amount) external onlySubscriptionManager(receiver) {
    if (amount > s_subscriptions[receiver].balance) {
      revert BalanceTooLow();
    }
    s_preparedWithdrawals[receiver] = PreparedWithdrawal({
      amount: amount,
      timestamp: block.timestamp + s_config.withdrawalDelay
    });

    emit PreparedWithdrawalRequest(receiver, amount);
  }

  /// @inheritdoc SubscriptionInterface
  function withdrawal(address receiver, uint256 amount) external onlySubscriptionManager(address(receiver)) {
    PreparedWithdrawal memory prepared = s_preparedWithdrawals[receiver];
    if (prepared.timestamp > block.timestamp) {
      revert DelayNotPassedYet(prepared.timestamp);
    }
    if (prepared.amount != amount) {
      revert AmountMismatch(prepared.amount, amount);
    }
    if (amount > s_subscriptions[receiver].balance) {
      revert BalanceTooLow();
    }
    s_subscriptions[receiver].balance -= amount;
    delete s_preparedWithdrawals[receiver];
    s_config.feeToken.safeTransfer(msg.sender, amount);

    emit WithdrawalProcessed(receiver, amount);
  }

  /// @inheritdoc SubscriptionInterface
  function getSubscriptionConfig() external view returns (SubscriptionConfig memory) {
    return s_config;
  }

  modifier onlySubscriptionManager(address contractAddress) {
    SubscriptionManagerInterface subscriptionManager = SubscriptionManagerInterface(contractAddress);
    if (subscriptionManager.getSubscriptionManager() != msg.sender) {
      revert InvalidManager();
    }
    _;
  }
}
