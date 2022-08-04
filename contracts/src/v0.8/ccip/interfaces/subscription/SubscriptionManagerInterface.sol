// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface SubscriptionManagerInterface {
  /**
   * @notice Gets the subscription manager.
   * @return the current subscription manager.
   */
  function getSubscriptionManager() external view returns (address);
}
