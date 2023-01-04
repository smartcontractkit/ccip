// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/offRamp/IAny2EVMOffRampRouter.sol";

contract MockTollOffRampRouter is IAny2EVMOffRampRouter {
  function routeMessage(Internal.Any2EVMMessageFromSender calldata message, bool manualExecution)
    external
    override
    returns (bool success)
  {}

  function addOffRamp(IBaseOffRamp) external {}

  function removeOffRamp(IBaseOffRamp) external {}

  function getOffRamps() external pure returns (IBaseOffRamp[] memory) {
    return new IBaseOffRamp[](0);
  }

  function isOffRamp(IBaseOffRamp) external pure returns (bool allowed) {
    return true;
  }
}
