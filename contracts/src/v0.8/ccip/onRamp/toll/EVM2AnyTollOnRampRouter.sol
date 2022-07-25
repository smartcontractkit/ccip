// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../pools/PoolCollector.sol";
import "../interfaces/Any2EVMTollOnRampRouterInterface.sol";

contract EVM2AnyTollOnRampRouter is
  Any2EVMTollOnRampRouterInterface,
  TypeAndVersionInterface,
  OwnerIsCreator,
  PoolCollector
{
  string public constant override typeAndVersion = "EVM2AnyTollOnRampRouter 1.0.0";

  // destination chain id => OnRampInterface
  mapping(uint256 => Any2EVMTollOnRampInterface) private s_onRamps;

  /// @inheritdoc Any2EVMTollOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64) {
    // Find and put the correct onRamp on the stack.
    Any2EVMTollOnRampInterface onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    if (message.tokens.length != message.amounts.length) revert BaseOnRampInterface.UnsupportedNumberOfTokens();

    // Charge the fee and subtract that amount from the feeTokenAmount. This will revert if
    // the given feeTokenAmount is too low for the needed fee.
    message.feeTokenAmount -= _chargeFee(onRamp, message.feeToken, message.feeTokenAmount);
    // Transfer the tokens to the token pools.
    _collectTokens(onRamp, message.tokens, message.amounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc Any2EVMTollOnRampRouterInterface
  function setOnRamp(uint256 chainId, Any2EVMTollOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc Any2EVMTollOnRampRouterInterface
  function getOnRamp(uint256 chainId) external view returns (Any2EVMTollOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
