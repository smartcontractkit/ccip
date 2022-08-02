// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/interfaces/Any2EVMOffRampRouterInterface.sol";

contract MockTollOffRampRouter is Any2EVMOffRampRouterInterface {
  function routeMessage(CCIP.Any2EVMMessage calldata message) external override returns (bool success) {}

  function addOffRamp(BaseOffRampInterface) external {}

  function removeOffRamp(BaseOffRampInterface) external {}

  function getOffRamps() external pure returns (BaseOffRampInterface[] memory) {
    return new BaseOffRampInterface[](0);
  }

  function isOffRamp(BaseOffRampInterface) external pure returns (bool allowed) {
    return true;
  }
}
