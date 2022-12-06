// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {GEConsumer} from "../models/GEConsumer.sol";
import {Common} from "../models/Common.sol";
import {GERouterInterface} from "../interfaces/router/GERouterInterface.sol";
import {Any2EVMMessageReceiverInterface} from "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";

contract PingPongDemo is Any2EVMMessageReceiverInterface, OwnerIsCreator {
  event Ping(uint256 pingPongCount);
  event Pong(uint256 pingPongCount);

  GERouterInterface internal s_router;

  // The chain ID of the counterpart ping pong contract
  uint256 public s_counterpartChainId;
  // The contract address of the counterpart ping pong contract
  address public s_counterpartAddress;

  // Pause ping-ponging
  bool public s_isPaused;

  // The fee token for CCIP billing
  address internal immutable i_feeToken;

  constructor(GERouterInterface router, address feeToken) {
    s_router = router;
    s_isPaused = false;
    i_feeToken = feeToken;
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
    GEConsumer.EVM2AnyGEMessage memory message = GEConsumer.EVM2AnyGEMessage({
      receiver: abi.encode(s_counterpartAddress),
      data: data,
      tokensAndAmounts: new Common.EVMTokenAndAmount[](0),
      extraArgs: GEConsumer._argsToBytes(GEConsumer.EVMExtraArgsV1({gasLimit: 200_000, strict: false})),
      feeToken: i_feeToken
    });
    s_router.ccipSend(s_counterpartChainId, message);
  }

  function ccipReceive(Common.Any2EVMMessage memory message) external override onlyRouter {
    uint256 pingPongCount = abi.decode(message.data, (uint256));
    if (!s_isPaused) {
      _respond(pingPongCount + 1);
    }
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////
  function setRouter(GERouterInterface router) external onlyOwner {
    s_router = router;
  }

  function getRouter() external view returns (GERouterInterface) {
    return s_router;
  }

  function setPaused(bool isPaused) external onlyOwner {
    s_isPaused = isPaused;
  }

  error InvalidRouter(address router);

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_router)) revert InvalidRouter(msg.sender);
    _;
  }
}
