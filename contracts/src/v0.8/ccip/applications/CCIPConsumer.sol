// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";
import {IRouter} from "../interfaces/router/IRouter.sol";
import {IERC165} from "../../vendor/IERC165.sol";

import {Consumer} from "../models/Consumer.sol";
import {Common} from "../models/Common.sol";

/// @title CCIPConsumer - Base contract for CCIP applications that can both send and receive messages.
abstract contract CCIPConsumer is IAny2EVMMessageReceiver, IERC165 {
  IRouter private immutable i_router;

  constructor(address router) {
    if (router == address(0)) revert InvalidRouter(address(0));
    i_router = IRouter(router);
  }

  /**
   * @notice IERC165 supports an interfaceId
   * @param interfaceId The interfaceId to check
   * @return true if the interfaceId is supported
   */
  function supportsInterface(bytes4 interfaceId) public pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  /// @inheritdoc IAny2EVMMessageReceiver
  function ccipReceive(Common.Any2EVMMessage calldata message) external override onlyRouter {
    _ccipReceive(message);
  }

  /**
   * @notice Override this function in your implementation.
   * @param message Any2EVMMessage
   */
  function _ccipReceive(Common.Any2EVMMessage memory message) internal virtual;

  /**
   * @notice Request a message to be sent to the destination chain
   * @dev Internal - Accessible by inheriting contracts
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return messageId assigned to message
   */
  function _ccipSend(uint64 destinationChainId, Consumer.EVM2AnyMessage memory message)
    internal
    returns (bytes32 messageId)
  {
    return i_router.ccipSend(destinationChainId, message);
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  /**
   * @notice Return the current router
   * @return i_router address
   */
  function getRouter() public view returns (address) {
    return address(i_router);
  }

  error InvalidRouter(address router);

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(i_router)) revert InvalidRouter(msg.sender);
    _;
  }
}
