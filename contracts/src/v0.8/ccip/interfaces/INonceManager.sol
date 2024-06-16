// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @notice Application contracts that intend to receive messages from
/// the router should implement this interface.
interface INonceManager {
  /// @notice Increments the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return The new outbound nonce
  function incrementOutboundNonce(uint64 destChainSelector, bytes calldata sender) external returns (uint64);
}
