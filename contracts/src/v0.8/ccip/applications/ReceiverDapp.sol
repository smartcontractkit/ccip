// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../interfaces/TollOffRampRouterInterface.sol";

/**
 * @notice Application contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is CrossChainMessageReceiverInterface, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  TollOffRampRouterInterface public immutable ROUTER;
  IERC20 public immutable TOKEN;

  error InvalidDeliverer(address deliverer);

  constructor(TollOffRampRouterInterface router, IERC20 token) {
    ROUTER = router;
    TOKEN = token;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external override {
    if (msg.sender != address(ROUTER)) revert InvalidDeliverer(msg.sender);
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(message.data, (address, address));
    for (uint256 i = 0; i < message.tokens.length; i++) {
      uint256 amount = message.amounts[i];
      if (destinationAddress != address(0) && amount != 0) {
        TOKEN.transfer(destinationAddress, amount);
      }
    }
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "ReceiverDapp 0.0.1";
  }
}
