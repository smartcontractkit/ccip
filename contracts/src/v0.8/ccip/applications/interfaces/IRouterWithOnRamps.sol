// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";

interface IRouterWithOnRamps is IRouterClient {
  /// @notice Return the configured onRamp for a specific destination chain.
  /// This is not needed for CCIP send and receive.
  /// It is only used in the long-running PingPongDemo to minimize manual funding ops.
  /// @param destChainSelector The destination chain Id to get the onRamp for.
  /// @return onRamp The address of the onRamp.
  function getOnRamp(uint64 destChainSelector) external view returns (address onRamp);
}
