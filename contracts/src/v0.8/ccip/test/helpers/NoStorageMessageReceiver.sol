// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";

contract NoStorageMessageReceiver is CrossChainMessageReceiverInterface {
  function receiveMessage(CCIP.AnyToEVMTollMessage calldata message) external override {}
}
