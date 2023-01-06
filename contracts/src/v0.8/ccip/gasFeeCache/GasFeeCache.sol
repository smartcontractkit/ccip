// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IGasFeeCache} from "../interfaces/gasFeeCache/IGasFeeCache.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Common} from "../models/Common.sol";
import {GE} from "../models/GE.sol";

contract GasFeeCache is IGasFeeCache, OwnerIsCreator {
  mapping(uint64 => TimestampedFeeUpdate) private s_linkPerUnitGasByDestChainId;
  mapping(address => bool) private s_feeUpdaters;
  uint128 private immutable i_stalenessThreshold;

  constructor(
    GE.FeeUpdate[] memory feeUpdates,
    address[] memory feeUpdaters,
    uint128 stalenessThreshold
  ) {
    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].chainId, feeUpdates[i].linkPerUnitGas, uint128(block.timestamp));
    }

    for (uint256 i = 0; i < feeUpdaters.length; ++i) {
      _setFeedUpdater(feeUpdaters[i]);
    }

    i_stalenessThreshold = stalenessThreshold;
  }

  function setFeeUpdater(address feeUpdater) external onlyOwner {
    _setFeedUpdater(feeUpdater);
  }

  function removeFeeUpdater(address feeUpdater) external onlyOwner {
    _removeFeeUpdater(feeUpdater);
  }

  function updateFees(GE.FeeUpdate[] memory feeUpdates) external requireUpdaterOrOwner {
    uint128 timestamp = uint128(block.timestamp);

    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].chainId, feeUpdates[i].linkPerUnitGas, timestamp);
    }
  }

  function getFee(uint64 destChainId) external view returns (uint128 fee) {
    TimestampedFeeUpdate memory update = s_linkPerUnitGasByDestChainId[destChainId];
    uint256 stalenessThreshold = i_stalenessThreshold;
    uint256 timePassed = block.timestamp - update.timestamp;
    if (timePassed > stalenessThreshold) revert StaleFee(stalenessThreshold, timePassed);

    // Must have an initial price.
    fee = update.linkPerUnitGas;
    if (fee == 0) revert ChainNotSupported(destChainId);
  }

  function getStalenessThreshold() external view returns (uint128) {
    return i_stalenessThreshold;
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  function _setFeedUpdater(address feeUpdater) private {
    if (feeUpdater != address(0)) {
      s_feeUpdaters[feeUpdater] = true;
      emit FeeUpdaterSet(feeUpdater);
    }
  }

  function _removeFeeUpdater(address feeUpdater) private {
    delete s_feeUpdaters[feeUpdater];
    emit FeeUpdaterRemoved(feeUpdater);
  }

  function _updateFee(
    uint64 chainId,
    uint128 linkPerUnitGas,
    uint128 timestamp
  ) private {
    s_linkPerUnitGasByDestChainId[chainId] = TimestampedFeeUpdate({
      linkPerUnitGas: linkPerUnitGas,
      timestamp: timestamp
    });
    emit GasFeeUpdated(chainId, linkPerUnitGas, timestamp);
  }

  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_feeUpdaters[msg.sender]) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
