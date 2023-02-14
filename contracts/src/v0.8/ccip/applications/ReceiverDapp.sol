// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";

import {CCIPReceiver} from "./CCIPReceiver.sol";

import {Client} from "../models/Client.sol";

import {IERC20} from "../../vendor/IERC20.sol";

/// @title ReceiverDapp - Application contract for receiving messages from the OffRamp on behalf of an EOA
/// @dev For test purposes only, not to be used as an example or production code.
contract ReceiverDapp is CCIPReceiver, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ReceiverDapp 2.0.0";

  constructor(address router) CCIPReceiver(router) {}

  /// @inheritdoc CCIPReceiver
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
    _handleMessage(message.data, message.destTokenAmounts);
  }

  function _handleMessage(bytes memory data, Client.EVMTokenAmount[] memory tokenAmounts) internal {
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(data, (address, address));
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      uint256 amount = tokenAmounts[i].amount;
      if (destinationAddress != address(0) && amount != 0) {
        IERC20(tokenAmounts[i].token).transfer(destinationAddress, amount);
      }
    }
  }
}
