// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyGEOnRamp} from "./IEVM2AnyGEOnRamp.sol";
import {GE} from "../../models/GE.sol";

interface IEVM2EVMGEOnRamp is IEVM2AnyGEOnRamp {
  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);
  error OnlyCallableByOwnerOrFeeAdmin();

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event CCIPSendRequested(GE.EVM2EVMGEMessage message);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external;

  /// @dev Struct to hold the fee configuration for a token
  struct FeeTokenConfig {
    uint96 feeAmount; // ---------┐ Flat fee
    uint64 multiplier; //         | Price multiplier for gas costs
    uint32 destGasOverhead; // ---┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Struct to hold the fee configuration for a token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // -------┐ Token address
    uint64 multiplier; // ---┘ Price multiplier for gas costs
    uint96 feeAmount; // ------┐ Flat fee in LINK
    uint32 destGasOverhead; //-┘ Extra gas charged on top of the gasLimit
  }
}
