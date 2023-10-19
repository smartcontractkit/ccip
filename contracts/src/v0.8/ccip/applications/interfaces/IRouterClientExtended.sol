// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../../interfaces/IRouterClient.sol";

interface IRouterClientExtended is IRouterClient {
    /// @notice Return the configured onramp for specific a destination chain.
    /// @param destChainSelector The destination chain Id to get the onRamp for.
    /// @return The address of the onRamp.
    function getOnRamp(uint64 destChainSelector) external view returns (address);
}
