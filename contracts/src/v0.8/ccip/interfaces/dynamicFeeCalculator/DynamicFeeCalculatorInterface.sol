// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";

interface DynamicFeeCalculatorInterface {
  error MismatchedFeeToken(address expected, address got);

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(DynamicFeeConfig feeConfig);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external;

  function getFee(CCIP.EVM2AnyGEMessage calldata message) external view returns (uint256 fee);

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
}
