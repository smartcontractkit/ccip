// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/PoolCollector.sol";
import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../access/OwnerIsCreator.sol";
import "../interfaces/TollOnRampInterface.sol";
import "../interfaces/TollOnRampRouterInterface.sol";

contract EVM2AnyTollOnRampRouter is TollOnRampRouterInterface, TypeAndVersionInterface, OwnerIsCreator, PoolCollector {
  using SafeERC20 for IERC20;

  // destination chain id => OnRampInterface
  mapping(uint256 => TollOnRampInterface) private s_onRamps;

  /// @inheritdoc TollOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64) {
    address sender = msg.sender;
    TollOnRampInterface onRamp = s_onRamps[destinationChainId];
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    if (message.tokens.length != message.amounts.length) revert BaseOnRampInterface.UnsupportedNumberOfTokens();

    uint256 feeTaken = _collectTokens(
      onRamp,
      message.tokens,
      message.amounts,
      message.feeToken,
      message.feeTokenAmount
    );
    message.feeTokenAmount -= feeTaken;

    return onRamp.forwardFromRouter(message, sender);
  }

  /// @inheritdoc TollOnRampRouterInterface
  function setOnRamp(uint256 chainId, TollOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /**
   * @notice Retrieve current and proposed OnRamp details for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint256 chainId) external view returns (TollOnRampInterface) {
    return s_onRamps[chainId];
  }

  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EVM2AnyTollOnRampRouter 1.0.0";
  }
}
