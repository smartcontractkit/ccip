// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../../utils/CCIP.sol";
import "./BaseOnRampInterface.sol";

interface TollOnRampInterface is BaseOnRampInterface {
  error UnsupportedFeeToken(IERC20 token);

  event CCIPSendRequested(CCIP.EVM2EVMTollEvent message);
  event OnRampConfigSet(OnRampConfig config);

  struct OnRampConfig {
    address router;
    // Fee for sending message taken in this contract
    uint64 relayingFeeJuels;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
  }

  /**
   * @notice Request a message to be sent to the destination chain
   * @param message the EVM2AnyTollMessage containing all message information
   * @param originalSender Original sender of the message if sent by a Router
   * @return The sequence number of the message
   */
  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender) external returns (uint64);

  function getRequiredFee(IERC20 feeToken) external returns (uint256);
}
