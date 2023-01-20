// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/offRamp/IAny2EVMOffRampRouter.sol";

contract MockTollOffRampRouter is IAny2EVMOffRampRouter {
  function routeMessage(
    Common.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external override returns (bool success) {}

  function addOffRamp(address) external {}

  function removeOffRamp(address) external {}

  function getOffRamps() external pure returns (address[] memory) {
    return new address[](0);
  }

  function isOffRamp(address) external pure returns (bool allowed) {
    return true;
  }
}
