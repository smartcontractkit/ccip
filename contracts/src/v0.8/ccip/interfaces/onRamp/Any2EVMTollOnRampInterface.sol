// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../models/Models.sol";
import "./BaseOnRampInterface.sol";

interface Any2EVMTollOnRampInterface is BaseOnRampInterface {
  error UnsupportedFeeToken(IERC20 token);

  event CCIPSendRequested(CCIP.EVM2EVMTollEvent message);

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender) external returns (uint64);

  /**
   * @notice Get the required fee for a specific fee token
   * @param feeToken token to get the fee for
   * @return fee uint256
   */
  function getRequiredFee(IERC20 feeToken) external returns (uint256);
}
