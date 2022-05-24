// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../pools/PoolCollector.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../access/OwnerIsCreator.sol";

contract OnRampRouter is TypeAndVersionInterface, OwnerIsCreator, PoolCollector {
  using SafeERC20 for IERC20;

  error OnRampAlreadySet(uint256 chainId, OnRampInterface onRamp);

  event OnRampSet(uint256 indexed chainId, OnRampInterface indexed onRamp);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  // destination chain id => OnRampInterface
  mapping(uint256 => OnRampInterface) private s_onRamps;

  /**
   * @notice Request a message to be sent to the destination chain
   * @param payload The message payload
   * @return The sequence number of the message
   */
  function requestCrossChainSend(CCIP.MessagePayload calldata payload) external returns (uint64) {
    address sender = msg.sender;
    OnRampInterface onRamp = s_onRamps[payload.destinationChainId];
    if (address(onRamp) == address(0)) revert OnRampInterface.UnsupportedDestinationChain(payload.destinationChainId);
    if (payload.tokens.length != payload.amounts.length) revert OnRampInterface.UnsupportedNumberOfTokens();
    _collectTokens(onRamp, payload);
    return onRamp.requestCrossChainSend(payload, sender);
  }

  function withdrawAccumulatedFees(
    IERC20 feeToken,
    address recipient,
    uint256 amount
  ) external onlyOwner {
    feeToken.safeTransfer(recipient, amount);
    emit FeesWithdrawn(feeToken, recipient, amount);
  }

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint256 chainId, OnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /**
   * @notice Retrieve current and proposed OnRamp details for the specified chain ID
   * @param chainId Chain ID to get ramp details for
   * @return onRamp
   */
  function getOnRamp(uint256 chainId) external view returns (OnRampInterface) {
    return s_onRamps[chainId];
  }

  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "OnRampRouter 0.0.1";
  }
}
