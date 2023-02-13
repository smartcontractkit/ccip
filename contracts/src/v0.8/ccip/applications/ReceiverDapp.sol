// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";

import {CCIPReceiver} from "./CCIPReceiver.sol";

import {Common} from "../models/Common.sol";

import {IERC20} from "../../vendor/IERC20.sol";

/// @title ReceiverDapp - Application contract for receiving messages from the OffRamp on behalf of an EOA
/// @dev For test purposes only, not to be used as an example or production code.
contract ReceiverDapp is CCIPReceiver, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ReceiverDapp 2.0.0";

  constructor(address router) CCIPReceiver(router) {}

  /// @inheritdoc CCIPReceiver
  function _ccipReceive(Common.Any2EVMMessage memory message) internal override {
    _handleMessage(message.data, message.destTokensAndAmounts);
  }

  function _handleMessage(bytes memory data, Common.EVMTokenAndAmount[] memory tokensAndAmounts) internal {
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(data, (address, address));
    for (uint256 i = 0; i < tokensAndAmounts.length; ++i) {
      uint256 amount = tokensAndAmounts[i].amount;
      if (destinationAddress != address(0) && amount != 0) {
        IERC20(tokensAndAmounts[i].token).transfer(destinationAddress, amount);
      }
    }
  }
}
