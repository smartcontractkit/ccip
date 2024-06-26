// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Contract interface that allows managing sender nonces
interface INonceManager {
  /// @notice Increments the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The sender address
  /// @return The new outbound nonce
  function getIncrementedOutboundNonce(uint64 destChainSelector, address sender) external returns (uint64);

  /// @notice Increments the inbound nonce for the given sender on the given source chain
  /// @param sourceChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return True if the nonce was incremented, false otherwise
  function incrementInboundNonce(
    uint64 sourceChainSelector,
    uint64 expectedNonce,
    bytes calldata sender
  ) external returns (bool);
}
