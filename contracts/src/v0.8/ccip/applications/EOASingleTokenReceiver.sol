// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../ramps/SingleTokenOffRamp.sol";
import "../../interfaces/TypeAndVersionInterface.sol";

/**
 * @notice Appliation contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract EOASingleTokenReceiver is CrossChainMessageReceiverInterface, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  SingleTokenOffRamp public immutable OFF_RAMP;

  error InvalidDeliverer(address deliverer);

  constructor(SingleTokenOffRamp offRamp) {
    OFF_RAMP = offRamp;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function receiveMessage(CCIP.Message calldata message) external override {
    if (msg.sender != address(OFF_RAMP)) revert InvalidDeliverer(msg.sender);
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(message.payload.data, (address, address));
    if (destinationAddress != address(0) && message.payload.amounts[0] != 0) {
      OFF_RAMP.TOKEN().transfer(destinationAddress, message.payload.amounts[0]);
    }
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EOASingleTokenReceiver 1.1.0";
  }
}
