// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";

interface DynamicFeeCalculatorInterface {
  error MismatchedFeeToken(address expected, address got);

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(DynamicFeeConfig feeConfig);

  function setFeeAdmin(address feeAdmin) external;

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external;

  function getFee(CCIP.EVM2AnyGEMessage calldata message) external returns (uint256 fee);

  struct DynamicFeeConfig {
    address feeToken;
    uint256 feeAmount;
    uint256 destGasOverhead;
    uint256 multiplier;
    address gasFeeCache;
    uint256 destChainId;
  }
}
