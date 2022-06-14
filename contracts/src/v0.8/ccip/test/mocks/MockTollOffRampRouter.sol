// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../../interfaces/TollOffRampRouterInterface.sol";

contract MockTollOffRampRouter is TollOffRampRouterInterface {
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage calldata message)
    external
  {}
}
