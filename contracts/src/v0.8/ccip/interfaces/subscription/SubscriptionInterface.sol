// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC20} from "../../../vendor/IERC20.sol";
import {SubscriptionManagerInterface} from "./SubscriptionManagerInterface.sol";

interface SubscriptionInterface {
  error SubscriptionAlreadyExists();
  error DelayNotPassedYet(uint256 allowedBy);
  error AddressMismatch(address expected, address got);
  error AmountMismatch(uint256 expected, uint256 got);
  error BalanceTooLow();
  error SubscriptionNotFound(address receiver);
  error InvalidManager();
  error FundingAmountNotPositive();

  event SubscriptionCreated(address receiver);
  event SubscriptionFunded(address receiver, uint256 amount);
  event PreparedSetSenders(address receiver, address[] newSenders);
  event NewSendersSet(address receiver, address[] newSenders);
  event PreparedWithdrawalRequest(address receiver, uint256 amount);
  event WithdrawalProcessed(address receiver, uint256 amount);
  event SubscriptionCharged(address receiver, uint256 amount);

  struct OffRampSubscription {
    address[] senders;
    SubscriptionManagerInterface receiver;
    bool strictSequencing;
    uint256 balance;
  }

  struct PreparedWithdrawal {
    uint256 amount;
    uint256 timestamp;
  }

  struct PreparedNewSenders {
    address[] newSenders;
    uint256 timestamp;
  }

  struct SubscriptionConfig {
    uint32 setSubscriptionSenderDelay;
    uint32 withdrawalDelay;
    IERC20 feeToken;
  }

  /**
   * @notice Gets the supported fee tokens
   * @return The supported fee tokens
   */
  function getSupportedTokensForExecutionFee() external returns (address[] memory);

  /**
   * @notice Gets the subscription corresponding to the given receiver
   * @param receiver The receiver for which to get the subscription
   * @return The subscription belonging to the receiver
   */
  function getSubscription(address receiver) external view returns (OffRampSubscription memory);

  /**
   * @notice Gets the fee token
   */
  function getFeeToken() external returns (IERC20);

  /**
   * @notice Creates a new subscription if one doesn't already exist for the
   *          given receiver
   * @param subscription The OffRampSubscription to be created
   */
  function createSubscription(OffRampSubscription memory subscription) external;

  /**
   * @notice Increases the balance of an existing subscription. The tokens
   *          need to be approved before making this call.
   * @param receiver Indicated which subscription to fund
   * @param amount The amount to fund the subscription
   */
  function fundSubscription(address receiver, uint256 amount) external;

  /**
   * @notice Indicates the desire to change the senders property on an
   *          existing subscription. This process can be completed after
   *          a set delay by calling `setSubscriptionSenders`. Calling
   *          this function again overwrites any existing prepared senders.
   * @param receiver Indicated which subscription to modify
   * @param newSenders The new sender addresses
   */
  function prepareSetSubscriptionSenders(address receiver, address[] memory newSenders) external;

  /**
   * @notice Finalizes a call to prepareSetSubscriptionSenders and actually
   *          modify the subscription.
   * @param receiver Indicated which subscription to modify
   * @param newSenders The new sender addresses, these are checked against the
   *          addresses previously given in the prepare step.
   */
  function setSubscriptionSenders(address receiver, address[] memory newSenders) external;

  /**
   * @notice Indicates the desire to withdrawal funds from a subscription
   *        This process can be completed after a set delay by calling
   *        `withdrawal`. Calling this function again overwrites any existing
   *        prepared withdrawal.
   * @param receiver Indicated which subscription to withdrawal from
   * @param amount The amount to withdrawal
   */
  function prepareWithdrawal(address receiver, uint256 amount) external;

  /**
   * @notice Completes the withdrawal previously initiated by calling
   *          `prepareWithdrawal`. This will send the token to the
   *          sender of this transaction.
   * @param receiver Indicated which subscription to withdrawal from
   * @param amount The amount to withdrawal
   */
  function withdrawal(address receiver, uint256 amount) external;

  /**
   * @notice Gets the current subscription configuration.
   * @return the current configuration
   */
  function getSubscriptionConfig() external view returns (SubscriptionConfig memory);
}
