// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/interfaces/Any2EVMTollOffRampRouterInterface.sol";

contract MockTollOffRampRouter is Any2EVMTollOffRampRouterInterface {
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage calldata message)
    external
  {}

  function addOffRamp(BaseOffRampInterface) external {}

  function removeOffRamp(BaseOffRampInterface) external {}

  function getOffRamps() external pure returns (BaseOffRampInterface[] memory) {
    return new BaseOffRampInterface[](0);
  }

  function isOffRamp(BaseOffRampInterface) external pure returns (bool allowed) {
    return true;
  }
}
