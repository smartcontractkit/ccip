// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {DynamicFeeCalculatorInterface} from "../interfaces/dynamicFeeCalculator/DynamicFeeCalculatorInterface.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {CCIP} from "../models/Models.sol";
import {GasFeeCacheInterface} from "../interfaces/dynamicFeeCalculator/GasFeeCacheInterfaceInterface.sol";

contract DynamicFeeCalculator is DynamicFeeCalculatorInterface, OwnerIsCreator {
  using CCIP for bytes;

  DynamicFeeConfig internal s_feeConfig;
  address internal s_feeAdmin;

  constructor(uint256, DynamicFeeConfig memory feeConfig) {
    s_feeConfig = feeConfig;
    emit FeeConfigSet(feeConfig);
  }

  function getFee(CCIP.EVM2AnyGEMessage calldata message) public returns (uint256 fee) {
    if (s_feeConfig.feeToken != message.feeToken) revert MismatchedFeeToken(s_feeConfig.feeToken, message.feeToken);
    uint256 gasLimit = message.extraArgs._fromBytes().gasLimit;
    uint256 gasFee = GasFeeCacheInterface(s_feeConfig.gasFeeCache).getFee(s_feeConfig.destChainId);

    return
      s_feeConfig.feeAmount + // Flat fee
      ((gasLimit + s_feeConfig.destGasOverhead) * gasFee * s_feeConfig.multiplier) / // Total gas reversed for tx
      1 ether; // latest gas reported gas fee with a safety margin
  }

  function setFeeAdmin(address feeAdmin) external onlyOwner {
    s_feeAdmin = feeAdmin;
    emit FeeAdminSet(feeAdmin);
  }

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external onlyOwner {
    s_feeConfig = feeConfig;
    emit FeeConfigSet(feeConfig);
  }
}
