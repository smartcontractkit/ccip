// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyGEOnRamp} from "./IEVM2AnyGEOnRamp.sol";
import {GE} from "../../models/GE.sol";

interface IEVM2EVMGEOnRamp is IEVM2AnyGEOnRamp {
  error MismatchedFeeToken(address expected, address got);
  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(DynamicFeeConfig feeConfig);
  event CCIPSendRequested(GE.EVM2EVMGEMessage message);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external;

  struct DynamicFeeConfig {
    // LINK
    address feeToken;
    // Flat fee in LINK
    uint256 feeAmount;
    // Extra gas charged on top of the gasLimit
    uint256 destGasOverhead;
    // Price multiplier for gas costs
    uint256 multiplier;
    // Fee manager contract
    address feeManager;
    // Destination chain ID
    uint64 destChainId;
  }
}
