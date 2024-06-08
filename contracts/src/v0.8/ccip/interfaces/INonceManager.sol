// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface INonceManager {
  function incrementOutboundNonce(uint64 destChainSelector, bytes calldata sender) external returns (uint64);
}
