// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../utils/CCIP.sol";
import "../interfaces/PoolInterface.sol";
import "../interfaces/AFNInterface.sol";

interface OnRampInterface {
  error TokenMismatch();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error UnsupportedFeeToken(IERC20 token);
  error SenderNotAllowed(address sender);
  error UnsupportedDestinationChain(uint256 destinationChainId);

  event CrossChainSendRequested(CCIP.Message message);
  event AllowlistEnabledSet(bool enabled);
  event AllowlistSet(address[] allowlist);
  event NewTokenBucketConstructed(uint256 rate, uint256 capacity, bool full);
  event PayloadConfigSet(uint256 maxDataSize, uint256 maxTokensLength);
  event RelayingFeeLinkSet(uint256 fee);
  event FeeCharged(address from, address to, uint256 fee);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param payload The message payload
   * @return The sequence number of the message
   */
  function requestCrossChainSend(CCIP.MessagePayload calldata payload) external returns (uint256);
}
