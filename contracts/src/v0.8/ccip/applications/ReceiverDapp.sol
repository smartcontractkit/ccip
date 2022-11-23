// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {Any2EVMMessageReceiverInterface} from "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import {Any2EVMOffRampRouterInterface} from "../interfaces/offRamp/Any2EVMOffRampRouterInterface.sol";
import {CCIP} from "../models/Models.sol";
import {IERC20} from "../../vendor/IERC20.sol";

/**
 * @notice Application contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is Any2EVMMessageReceiverInterface, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
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

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMMessage calldata message) external override onlyRouter {
    _handleMessage(message.data, message.destTokensAndAmounts);
  }

  function _handleMessage(bytes memory data, CCIP.EVMTokenAndAmount[] memory tokensAndAmounts) internal {
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

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_router)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
