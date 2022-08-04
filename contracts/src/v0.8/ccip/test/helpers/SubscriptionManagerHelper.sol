// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/subscription/SubscriptionManagerInterface.sol";

contract SubscriptionManagerHelper is SubscriptionManagerInterface {
  address s_manager;

  constructor(address manager) {
    s_manager = manager;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_manager;
  }
}
