// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../../utils/CCIP.sol";
import "./TollOnRampInterface.sol";
import "./BaseOnRampRouterInterface.sol";

interface TollOnRampRouterInterface is BaseOnRampRouterInterface {
  error OnRampAlreadySet(uint256 chainId, TollOnRampInterface onRamp);

  event OnRampSet(uint256 indexed chainId, TollOnRampInterface indexed onRamp);

  /**
   * @notice Request a message to be sent to the destination chain
   * @param destinationChainId The destination chain ID
   * @param message The message payload
   * @return The sequence number assigned to message
   */
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64);

  /**
   * @notice Set chainId => onRamp mapping
   * @dev only callable by owner
   * @param chainId destination chain ID
   * @param onRamp OnRamp to use for that destination chain
   */
  function setOnRamp(uint256 chainId, TollOnRampInterface onRamp) external;
}
