// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../../models/Common.sol";
import {GE} from "../../models/GE.sol";
import {GEConsumer} from "../../models/GEConsumer.sol";
import {BaseOnRampInterface} from "./BaseOnRampInterface.sol";

interface EVM2EVMGEOnRampInterface is BaseOnRampInterface {
  error MismatchedFeeToken(address expected, address got);
  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);

  event CCIPSendRequested(GE.EVM2EVMGEMessage message);
  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(DynamicFeeConfig feeConfig);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external;

  function getFee(GEConsumer.EVM2AnyGEMessage calldata message) external view returns (uint256 fee);

  struct DynamicFeeConfig {
    // LINK
    address feeToken;
    // Flat fee in LINK
    uint256 feeAmount;
    // Extra gas charged on top of the gasLimit
    uint256 destGasOverhead;
    // Price multiplier for gas costs
    uint256 multiplier;
    // Gas fee cache contract
    address gasFeeCache;
    // Destination chain ID
    uint256 destChainId;
  }

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(
    GEConsumer.EVM2AnyGEMessage memory message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (bytes32);
}
