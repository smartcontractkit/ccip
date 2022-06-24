// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/CrossChainMessageReceiverInterface.sol";

contract NoStorageMessageReceiver is CrossChainMessageReceiverInterface {
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external override {}
}
