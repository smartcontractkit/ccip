// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../libraries/Client.sol";
import {CCIPClient} from "../external/CCIPClient.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @title PingPongDemo - A simple ping-pong contract for demonstrating cross-chain communication
contract PingPongDemo is CCIPClient {
  using SafeERC20 for IERC20;

  event Ping(uint256 pingPongCount);
  event Pong(uint256 pingPongCount);

  // The chain ID of the counterpart ping pong contract
  uint64 internal s_counterpartChainSelector;

  // The contract address of the counterpart ping pong contract
  address internal s_counterpartAddress;

  // Pause ping-ponging
  bool private s_isPaused;

  // CCIPClient will handle the token approval so there's no need to do it here
  constructor(address router, IERC20 feeToken) CCIPClient(router, feeToken) {}

  function typeAndVersion() external pure virtual override returns (string memory) {
    return "PingPongDemo 1.3.0";
  }

  function startPingPong() external onlyOwner {
    s_isPaused = false;

    // Start the game
    _respond(1);
  }

  function _respond(uint256 pingPongCount) internal virtual {
    if (pingPongCount & 1 == 1) {
      emit Ping(pingPongCount);
    } else {
      emit Pong(pingPongCount);
    }

    bytes memory data = abi.encode(pingPongCount);

    ccipSend({
      destChainSelector: s_counterpartChainSelector, // destChaio
      tokenAmounts: new Client.EVMTokenAmount[](0),
      data: data,
      feeToken: address(s_feeToken)
    });
  }

  /// @notice This function the entrypoint for this contract to process messages.
  /// @param message The message to process.
  /// @dev This example just sends the tokens to the owner of this contracts. More
  /// interesting functions could be implemented.
  /// @dev It has to be external because of the try/catch.
  function processMessage(Client.Any2EVMMessage calldata message)
    external
    override
    onlySelf
    validSender(message.sourceChainSelector, message.sender)
    validChain(message.sourceChainSelector)
  {
    uint256 pingPongCount = abi.decode(message.data, (uint256));
    if (!s_isPaused) {
      _respond(pingPongCount + 1);
    }
  }

  // ================================================================
  // │                     Admin Functions                          │
  // ================================================================

  function setCounterpart(uint64 counterpartChainSelector, address counterpartAddress) external onlyOwner {
    s_counterpartChainSelector = counterpartChainSelector;
    s_counterpartAddress = counterpartAddress;

    // Approve the counterpart contract under validSender
    s_approvedSenders[counterpartChainSelector][abi.encode(counterpartAddress)] = true;

    // Approve the counterpart Chain selector under validChain
    s_chains[counterpartChainSelector] = abi.encode(counterpartAddress);
  }

  function setCounterpartChainSelector(uint64 counterpartChainSelector) external onlyOwner {
    s_counterpartChainSelector = counterpartChainSelector;
  }

  function setCounterpartAddress(address counterpartAddress) external onlyOwner {
    s_counterpartAddress = counterpartAddress;

    s_chains[s_counterpartChainSelector] = abi.encode(counterpartAddress);
  }

  function setPaused(bool pause) external onlyOwner {
    s_isPaused = pause;
  }

  // ================================================================
  // │                      State Management                        │
  // ================================================================

  function getCounterpartChainSelector() external view returns (uint64) {
    return s_counterpartChainSelector;
  }

  function getCounterpartAddress() external view returns (address) {
    return s_counterpartAddress;
  }

  function isPaused() external view returns (bool) {
    return s_isPaused;
  }
}
