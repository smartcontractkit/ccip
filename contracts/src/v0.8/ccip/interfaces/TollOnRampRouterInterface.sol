// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../utils/CCIP.sol";
import "./TollOnRampInterface.sol";

interface TollOnRampRouterInterface {
  error OnRampAlreadySet(uint256 chainId, TollOnRampInterface onRamp);

  event OnRampSet(uint256 indexed chainId, TollOnRampInterface indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The EVM2AnyTollMessage payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64);
}
