// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../applications/interfaces/CrossChainMessageReceiverInterface.sol";

contract RevertingMessageReceiver is CrossChainMessageReceiverInterface {
  address s_manager;

  constructor() {
    s_manager = msg.sender;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_manager;
  }

  function ccipReceive(CCIP.Any2EVMSubscriptionMessage calldata) external pure override {
    revert();
  }

  function ccipReceive(CCIP.Any2EVMTollMessage calldata) external pure override {
    revert();
  }

  function ccipReceive(CCIP.Any2EVMMOMessage calldata) external pure override {
    revert();
  }
}
