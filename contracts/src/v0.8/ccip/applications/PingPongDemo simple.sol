// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {IERC20} from "../../vendor/IERC20.sol";

interface CCIPRouterInterface {
  struct Message {
    bytes receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    uint256 gasLimit;
  }

  function ccipSend(uint256 destinationChainId, Message memory message) external;
}

interface CCIPReceiverInterface {
  struct ReceivedMessage {
    uint256 sourceChainId;
    bytes sender;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
  }

  function ccipReceive(ReceivedMessage memory message) external;
}

contract PingPongDemo is CCIPReceiverInterface, OwnerIsCreator {
  event Ping(uint256 pingPongCount);
  event Pong(uint256 pingPongCount);

  CCIPRouterInterface internal s_receivingRouter;
  CCIPRouterInterface internal s_sendingRouter;

  // The chain ID of the counterpart ping pong contract
  uint256 public s_counterpartChainId;
  // The contract address of the counterpart ping pong contract
  address public s_counterpartAddress;

  constructor(CCIPRouterInterface receivingRouter, CCIPRouterInterface sendingRouter) {
    s_receivingRouter = receivingRouter;
    s_sendingRouter = sendingRouter;
  }

  function setCounterpart(uint256 counterpartChainId, address counterpartAddress) external onlyOwner {
    s_counterpartChainId = counterpartChainId;
    s_counterpartAddress = counterpartAddress;
  }

  function startPingPong() external onlyOwner {
    _respond(1);
  }

  function _respond(uint256 pingPongCount) private {
    if (pingPongCount & 1 == 1) {
      // odd pingPongCount
      emit Ping(pingPongCount);
    } else {
      // even pingPongCount
      emit Pong(pingPongCount);
    }

    bytes memory data = abi.encode(pingPongCount);
    CCIPRouterInterface.Message memory message = CCIPRouterInterface.Message({
      receiver: abi.encode(s_counterpartAddress),
      data: data,
      tokens: new IERC20[](0),
      amounts: new uint256[](0),
      gasLimit: 200_000
    });
    s_sendingRouter.ccipSend(s_counterpartChainId, message);
  }

  function ccipReceive(ReceivedMessage memory message) external override onlyRouter {
    require(message.sourceChainId == s_counterpartChainId, "wrong source chain");
    require(abi.decode(message.sender, (address)) == s_counterpartAddress, "wrong sender");

    uint256 pingPongCount = abi.decode(message.data, (uint256));
    _respond(pingPongCount + 1);
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  function setRouters(CCIPRouterInterface receivingRouter, CCIPRouterInterface sendingRouter) external onlyOwner {
    s_receivingRouter = receivingRouter;
    s_sendingRouter = sendingRouter;
  }

  function getRouters() external view returns (CCIPRouterInterface, CCIPRouterInterface) {
    return (s_receivingRouter, s_sendingRouter);
  }

  function getSubscriptionManager() external view returns (address) {
    return owner();
  }

  error InvalidRouter(address router);

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_receivingRouter)) revert InvalidRouter(msg.sender);
    _;
  }
}
