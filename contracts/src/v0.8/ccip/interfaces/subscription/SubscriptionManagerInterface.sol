// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface SubscriptionManagerInterface {
  /**
   * @notice Gets the subscription manager who is allowed to create/update
   * the subscription for this receiver contract.
   * @return the current subscription manager.
   */
  function getSubscriptionManager() external view returns (address);
}
