// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {Any2EVMMessageReceiverInterface} from "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import {Any2EVMOffRampRouterInterface} from "../interfaces/offRamp/Any2EVMOffRampRouterInterface.sol";
import {CCIP, IERC20} from "../models/Models.sol";

/**
 * @notice Application contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is Any2EVMMessageReceiverInterface, TypeAndVersionInterface {
  string public constant override typeAndVersion = "ReceiverDapp 1.0.0";

  Any2EVMOffRampRouterInterface public s_router;

  address internal s_manager;

  error InvalidDeliverer(address deliverer);

  constructor(Any2EVMOffRampRouterInterface router) {
    s_router = router;
    s_manager = msg.sender;
  }

  function setRouter(Any2EVMOffRampRouterInterface router) public {
    s_router = router;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_manager;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMMessage calldata message) external override onlyRouter {
    handleMessage(message.data, message.destTokens, message.amounts);
  }

  function handleMessage(
    bytes memory data,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal {
    (
      ,
      /* address originalSender */
      address destinationAddress
    ) = abi.decode(data, (address, address));
    for (uint256 i = 0; i < tokens.length; ++i) {
      uint256 amount = amounts[i];
      if (destinationAddress != address(0) && amount != 0) {
        tokens[i].transfer(destinationAddress, amount);
      }
    }
  }

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_router)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
