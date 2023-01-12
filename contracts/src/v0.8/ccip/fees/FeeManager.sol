// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../interfaces/fees/IFeeManager.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Common} from "../models/Common.sol";
import {GE} from "../models/GE.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {SafeERC20} from "../../vendor/SafeERC20.sol";

/**
 * @notice The FeeManager contract has 2 responsibilities:
 * 1. It stores the current fee for a given destination chain, allowing the owner or feeUpdater to update this value.
 * 2. Store tokens other than LINK that can be used to pay for gas, and which the owner can withdraw.
 */
contract FeeManager is IFeeManager, OwnerIsCreator {
  using SafeERC20 for IERC20;

  // token => destChainId => feeUpdate
  mapping(address => mapping(uint64 => TimestampedFeeUpdate)) private s_tokenPerUnitGasByDestChainId;
  // feeUpdater => isFeeUpdater
  mapping(address => bool) private s_feeUpdaters;
  // The amount of time a fee can be stale before it is considered invalid.
  uint128 private immutable i_stalenessThreshold;

  constructor(
    GE.FeeUpdate[] memory feeUpdates,
    address[] memory feeUpdaters,
    uint128 stalenessThreshold
  ) {
    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].token, feeUpdates[i].chainId, feeUpdates[i].linkPerUnitGas, uint128(block.timestamp));
    }

    for (uint256 i = 0; i < feeUpdaters.length; ++i) {
      _setFeedUpdater(feeUpdaters[i]);
    }

    i_stalenessThreshold = stalenessThreshold;
  }

  // @inheritdoc IFeeManager
  function setFeeUpdater(address feeUpdater) external override onlyOwner {
    _setFeedUpdater(feeUpdater);
  }

  // @inheritdoc IFeeManager
  function removeFeeUpdater(address feeUpdater) external override onlyOwner {
    _removeFeeUpdater(feeUpdater);
  }

  // @inheritdoc IFeeManager
  function updateFees(GE.FeeUpdate[] memory feeUpdates) external override requireUpdaterOrOwner {
    uint128 timestamp = uint128(block.timestamp);

    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].token, feeUpdates[i].chainId, feeUpdates[i].linkPerUnitGas, timestamp);
    }
  }

  // @inheritdoc IFeeManager
  function getFee(address token, uint64 destChainId) external view override returns (uint128 fee) {
    TimestampedFeeUpdate memory update = s_tokenPerUnitGasByDestChainId[token][destChainId];
    if (update.timestamp == 0 || update.linkPerUnitGas == 0) revert TokenOrChainNotSupported(token, destChainId);
    uint256 stalenessThreshold = i_stalenessThreshold;
    uint256 timePassed = block.timestamp - update.timestamp;
    if (timePassed > stalenessThreshold) revert StaleFee(stalenessThreshold, timePassed);

    return update.linkPerUnitGas;
  }

  // @inheritdoc IFeeManager
  function getStalenessThreshold() external view override returns (uint128) {
    return i_stalenessThreshold;
  }

  // @inheritdoc IFeeManager
  function withdrawToken(
    address token,
    address to,
    uint256 amount
  ) external override onlyOwner {
    if (to == address(0)) revert InvalidWithdrawalAddress();
    IERC20(token).safeTransfer(to, amount);
  }

  /////////////////////////////////////////////////////////////////////
  // Plumbing
  /////////////////////////////////////////////////////////////////////

  /**
   * @notice Set a new fee updater.
   */
  function _setFeedUpdater(address feeUpdater) private {
    if (feeUpdater != address(0)) {
      s_feeUpdaters[feeUpdater] = true;
      emit FeeUpdaterSet(feeUpdater);
    }
  }

  /**
   * @notice Remove a fee updater.
   */
  function _removeFeeUpdater(address feeUpdater) private {
    delete s_feeUpdaters[feeUpdater];
    emit FeeUpdaterRemoved(feeUpdater);
  }

  /**
   * @notice Update the fee for a given token and destination chain.
   */
  function _updateFee(
    address token,
    uint64 chainId,
    uint128 linkPerUnitGas,
    uint128 timestamp
  ) private {
    s_tokenPerUnitGasByDestChainId[token][chainId] = TimestampedFeeUpdate({
      linkPerUnitGas: linkPerUnitGas,
      timestamp: timestamp
    });
    emit GasFeeUpdated(token, chainId, linkPerUnitGas, timestamp);
  }

  /**
   * @notice Require that the caller is the owner or a fee updater.
   */
  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_feeUpdaters[msg.sender]) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
