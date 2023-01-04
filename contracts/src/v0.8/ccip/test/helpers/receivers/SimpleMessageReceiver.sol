// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/applications/IAny2EVMMessageReceiver.sol";

contract SimpleMessageReceiver is IAny2EVMMessageReceiver {
  event MessageReceived();

  address private immutable i_manager;

  constructor() {
    i_manager = msg.sender;
  }

  function ccipReceive(Common.Any2EVMMessage calldata) external override {
    emit MessageReceived();
  }

  function getManager() external view returns (address) {
    return i_manager;
  }
}
