// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {IERC20} from "../../vendor/IERC20.sol";

struct EVMTokenAndAmount {
  address token;
  uint256 amount;
}

interface CCIPRouterInterface {
  struct Message {
    bytes receiver;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
    bytes extraArgs;
  }

  function ccipSend(uint256 destinationChainId, Message memory message) external returns (uint64);
}

interface CCIPReceiverInterface {
  struct ReceivedMessage {
    uint256 sourceChainId;
    bytes sender;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
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

  // Pause ping-ponging
  bool public s_isPaused;

  constructor(CCIPRouterInterface receivingRouter, CCIPRouterInterface sendingRouter) {
    s_receivingRouter = receivingRouter;
    s_sendingRouter = sendingRouter;
    s_isPaused = false;
  }

  function setCounterpart(uint256 counterpartChainId, address counterpartAddress) external onlyOwner {
    s_counterpartChainId = counterpartChainId;
    s_counterpartAddress = counterpartAddress;
  }

  function startPingPong() external onlyOwner {
    s_isPaused = false;
    _respond(1);
  }

  function _respond(uint256 pingPongCount) private {
    if (pingPongCount & 1 == 1) {
      emit Ping(pingPongCount);
    } else {
      emit Pong(pingPongCount);
    }

    bytes memory data = abi.encode(pingPongCount);
    CCIPRouterInterface.Message memory message = CCIPRouterInterface.Message({
      receiver: abi.encode(s_counterpartAddress),
      data: data,
      tokensAndAmounts: new EVMTokenAndAmount[](0),
      extraArgs: _toBytes(EVMExtraArgsV1({gasLimit: 200_000}))
    });
    s_sendingRouter.ccipSend(s_counterpartChainId, message);
  }

  function ccipReceive(ReceivedMessage memory message) external override onlyRouter {
    uint256 pingPongCount = abi.decode(message.data, (uint256));
    if (!s_isPaused) {
      _respond(pingPongCount + 1);
    }
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  struct EVMExtraArgsV1 {
    uint256 gasLimit;
  }

  bytes4 public constant EVM_EXTRA_ARGS_V1_TAG = 0x97a657c9;

  function _toBytes(EVMExtraArgsV1 memory extraArgs) internal pure returns (bytes memory bts) {
    return bytes.concat(EVM_EXTRA_ARGS_V1_TAG, abi.encode(extraArgs.gasLimit));
  }

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

  function setPaused(bool isPaused) external onlyOwner {
    s_isPaused = isPaused;
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
