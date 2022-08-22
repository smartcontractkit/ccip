// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import "../interfaces/offRamp/Any2EVMOffRampRouterInterface.sol";

/**
 * @notice Application contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is Any2EVMMessageReceiverInterface, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "ReceiverDapp 1.0.0";

  Any2EVMOffRampRouterInterface public immutable i_router;

  address s_manager;

  error InvalidDeliverer(address deliverer);

  constructor(Any2EVMOffRampRouterInterface router) {
    i_router = router;
    s_manager = msg.sender;
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
    handleMessage(message.data, message.tokens, message.amounts);
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
    if (msg.sender != address(i_router)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
