// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/PoolCollector.sol";
import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../access/OwnerIsCreator.sol";
import "../interfaces/Any2EVMTollOnRampInterface.sol";
import "../interfaces/Any2EVMTollOnRampRouterInterface.sol";

contract EVM2AnyTollOnRampRouter is
  Any2EVMTollOnRampRouterInterface,
  TypeAndVersionInterface,
  OwnerIsCreator,
  PoolCollector
{
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2AnyTollOnRampRouter 1.0.0";

  // destination chain id => OnRampInterface
  mapping(uint256 => Any2EVMTollOnRampInterface) private s_onRamps;

  /// @inheritdoc Any2EVMTollOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64) {
    Any2EVMTollOnRampInterface onRamp = s_onRamps[destinationChainId];
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    if (message.tokens.length != message.amounts.length) revert BaseOnRampInterface.UnsupportedNumberOfTokens();

    message.feeTokenAmount -= _chargeFee(onRamp, message.feeToken, message.feeTokenAmount);
    _collectTokens(onRamp, message.tokens, message.amounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc Any2EVMTollOnRampRouterInterface
  function setOnRamp(uint256 chainId, Any2EVMTollOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /**
   * @notice Retrieve current and proposed OnRamp details for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint256 chainId) external view returns (Any2EVMTollOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
