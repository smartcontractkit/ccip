// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../interfaces/OffRampRouterInterface.sol";

/**
 * @notice Appliation contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is CrossChainMessageReceiverInterface, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  OffRampRouterInterface public immutable ROUTER;
  IERC20 public immutable TOKEN;

  error InvalidDeliverer(address deliverer);

  constructor(OffRampRouterInterface router, IERC20 token) {
    ROUTER = router;
    TOKEN = token;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function receiveMessage(CCIP.Message calldata message) external override {
    if (msg.sender != address(ROUTER)) revert InvalidDeliverer(msg.sender);
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(message.payload.data, (address, address));
    for (uint256 i = 0; i < message.payload.tokens.length; i++) {
      uint256 amount = message.payload.amounts[i];
      if (destinationAddress != address(0) && amount != 0) {
        TOKEN.transfer(destinationAddress, amount);
      }
    }
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "ReceiverDapp 0.0.1";
  }
}
