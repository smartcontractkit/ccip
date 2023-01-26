// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyGEOnRamp} from "./IEVM2AnyGEOnRamp.sol";
import {GE} from "../../models/GE.sol";

interface IEVM2EVMGEOnRamp is IEVM2AnyGEOnRamp {
  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);
  error OnlyCallableByOwnerOrFeeAdmin();

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(DynamicFeeConfig feeConfig);
  event CCIPSendRequested(GE.EVM2EVMGEMessage message);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external;

  struct DynamicFeeConfig {
    address linkToken; // -----┐    LINK token address
    uint96 feeAmount; // ------┘    Flat fee in LINK
    uint64 multiplier; // -----┐    Price multiplier for gas costs
    uint32 destGasOverhead; // |    Extra gas charged on top of the gasLimit
    address feeManager; // ----┘    Fee manager contract
  }
}
