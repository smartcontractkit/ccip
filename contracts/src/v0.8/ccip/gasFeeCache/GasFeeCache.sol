// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../models/Common.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {GasFeeCacheInterface} from "../interfaces/gasFeeCache/GasFeeCacheInterface.sol";
import {GE} from "../models/GE.sol";

contract GasFeeCache is GasFeeCacheInterface, OwnerIsCreator {
  mapping(uint256 => uint256) internal s_linkPerUnitGasByDestChainId;
  mapping(address => bool) internal s_feeUpdaters;

  constructor(GE.FeeUpdate[] memory feeUpdates, address[] memory feeUpdaters) {
    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      s_linkPerUnitGasByDestChainId[feeUpdates[i].chainId] = feeUpdates[i].linkPerUnitGas;
    }

    for (uint256 i = 0; i < feeUpdaters.length; ++i) {
      s_feeUpdaters[feeUpdaters[i]] = true;
      emit FeeUpdaterSet(feeUpdaters[i]);
    }
  }

  function setFeeUpdater(address feeUpdater) external onlyOwner {
    if (feeUpdater != address(0)) {
      s_feeUpdaters[feeUpdater] = true;
      emit FeeUpdaterSet(feeUpdater);
    }
  }

  function removeFeeUpdater(address feeUpdater) external onlyOwner {
    delete s_feeUpdaters[feeUpdater];
    emit FeeUpdaterRemoved(feeUpdater);
  }

  function updateFees(GE.FeeUpdate[] memory feeUpdates) external requireUpdaterOrOwner {
    if (!s_feeUpdaters[msg.sender]) revert FeeUpdaterNotAllowed(msg.sender);

    uint256 numberOfFeeUpdates = feeUpdates.length;
    for (uint256 i = 0; i < numberOfFeeUpdates; ++i) {
      s_linkPerUnitGasByDestChainId[feeUpdates[i].chainId] = feeUpdates[i].linkPerUnitGas;
      emit GasFeeUpdated(feeUpdates[i].chainId, feeUpdates[i].linkPerUnitGas);
    }
  }

  function getFee(uint256 destChainId) external view returns (uint256 fee) {
    fee = s_linkPerUnitGasByDestChainId[destChainId];
    // Must have an initial price.
    if (fee == 0) revert ChainNotSupported(destChainId);
  }

  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_feeUpdaters[msg.sender]) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
