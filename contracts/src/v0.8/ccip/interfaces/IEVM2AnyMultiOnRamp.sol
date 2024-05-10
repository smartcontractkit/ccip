// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRampClient} from "./IEVM2AnyOnRampClient.sol";

interface IEVM2AnyMultiOnRamp is IEVM2AnyOnRampClient {
  /// @notice Gets the next sequence number to be used in the onRamp
  /// @param destChainSelector The destination chain selector
  /// @return the next sequence number to be used
  function getExpectedNextSequenceNumber(uint64 destChainSelector) external view returns (uint64);

  /// @notice Get the next nonce for a given sender
  /// @param destChainSelector The destination chain selector
  /// @param sender The sender to get the nonce for
  /// @return nonce The next nonce for the sender
  function getSenderNonce(uint64 destChainSelector, address sender) external view returns (uint64 nonce);
}
