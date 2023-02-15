// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IFeeManager} from "../interfaces/fees/IFeeManager.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Client} from "../models/Client.sol";
import {Internal} from "../models/Internal.sol";
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
  uint32 private immutable i_stalenessThreshold;

  constructor(
    Internal.FeeUpdate[] memory feeUpdates,
    address[] memory feeUpdaters,
    uint32 stalenessThreshold
  ) {
    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].sourceFeeToken, feeUpdates[i].destChainId, feeUpdates[i].feeTokenBaseUnitsPerUnitGas);
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
  function updateFees(Internal.FeeUpdate[] memory feeUpdates) external override requireUpdaterOrOwner {
    for (uint256 i = 0; i < feeUpdates.length; ++i) {
      _updateFee(feeUpdates[i].sourceFeeToken, feeUpdates[i].destChainId, feeUpdates[i].feeTokenBaseUnitsPerUnitGas);
    }
  }

  // @inheritdoc IFeeManager
  function getFeeTokenBaseUnitsPerUnitGas(address token, uint64 destChainId)
    external
    view
    override
    returns (uint128 feeTokenBaseUnitsPerUnitGas)
  {
    TimestampedFeeUpdate memory update = s_tokenPerUnitGasByDestChainId[token][destChainId];
    if (update.timestamp == 0 || update.feeTokenBaseUnitsPerUnitGas == 0)
      revert TokenOrChainNotSupported(token, destChainId);
    uint256 timePassed = block.timestamp - update.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleFee(i_stalenessThreshold, timePassed);

    return update.feeTokenBaseUnitsPerUnitGas;
  }

  // @inheritdoc IFeeManager
  function getStalenessThreshold() external view override returns (uint32) {
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

  /**
   * @notice Set a new fee updater.
   * @param feeUpdater The address of the feeUpdater that is now allowed
   * to send fee updates.
   */
  function _setFeedUpdater(address feeUpdater) private {
    if (feeUpdater != address(0)) {
      s_feeUpdaters[feeUpdater] = true;
      emit FeeUpdaterSet(feeUpdater);
    }
  }

  /**
   * @notice Remove a fee updater.
   * @param feeUpdater The address of the feeUpdater that is no longer allowed
   * to send fee updates.
   */
  function _removeFeeUpdater(address feeUpdater) private {
    delete s_feeUpdaters[feeUpdater];
    emit FeeUpdaterRemoved(feeUpdater);
  }

  /**
   * @notice Update the fee for a given fee token and destination chain.
   * @param token The source chain token that is used for payments.
   * @param destinationChainId The destination chain id.
   * @param feeTokenBaseUnitsPerUnitGas The cost of destination chain gas in Link tokens.
   */
  function _updateFee(
    address token,
    uint64 destinationChainId,
    uint128 feeTokenBaseUnitsPerUnitGas
  ) private {
    if (token == address(0)) revert NullAddressNotAllowed();
    s_tokenPerUnitGasByDestChainId[token][destinationChainId] = TimestampedFeeUpdate({
      feeTokenBaseUnitsPerUnitGas: feeTokenBaseUnitsPerUnitGas,
      timestamp: uint64(block.timestamp)
    });
    emit GasFeeUpdated(token, destinationChainId, feeTokenBaseUnitsPerUnitGas, uint64(block.timestamp));
  }

  /**
   * @notice Require that the caller is the owner or a fee updater.
   */
  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_feeUpdaters[msg.sender]) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
