// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface BaseOnRampRouterInterface {
  error UnsupportedDestinationChain(uint64 destinationChainId);

  /**
   * @notice Checks if the given destination chain ID is supported
   * @param chainId The destination chain to check
   * @return supported is true if it is supported, false if not
   */
  function isChainSupported(uint64 chainId) external view returns (bool supported);
}
