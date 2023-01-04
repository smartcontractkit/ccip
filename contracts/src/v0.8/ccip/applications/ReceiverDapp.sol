// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";
import {IAny2EVMOffRampRouter} from "../interfaces/offRamp/IAny2EVMOffRampRouter.sol";
import {GEConsumer} from "../models/GEConsumer.sol";
import {Common} from "../models/Common.sol";
import {IERC20} from "../../vendor/IERC20.sol";

/**
 * @notice Application contract for receiving messages from the OffRamp on behalf of an EOA
 */
contract ReceiverDapp is IAny2EVMMessageReceiver, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "ReceiverDapp 1.0.0";

  IAny2EVMOffRampRouter public s_router;

  address internal s_manager;

  error InvalidDeliverer(address deliverer);

  constructor(IAny2EVMOffRampRouter router) {
    s_router = router;
    s_manager = msg.sender;
  }

  function setRouter(IAny2EVMOffRampRouter router) public {
    s_router = router;
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function ccipReceive(Common.Any2EVMMessage calldata message) external override onlyRouter {
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

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_router)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
