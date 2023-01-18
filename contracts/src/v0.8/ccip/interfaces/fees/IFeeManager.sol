// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../../models/Common.sol";
import {GE} from "../../models/GE.sol";

interface IFeeManager {
  error TokenOrChainNotSupported(address token, uint64 chain);
  error FeeUpdaterNotAllowed(address feeUpdater);
  error OnlyCallableByUpdaterOrOwner();
  error StaleFee(uint256 threshold, uint256 timePassed);
  error InvalidWithdrawalAddress();
  error NullAddressNotAllowed();

  event FeeUpdaterSet(address indexed feeUpdater);
  event FeeUpdaterRemoved(address indexed feeUpdater);
  event GasFeeUpdated(address indexed token, uint64 indexed destChain, uint128 linkPerUnitGas, uint128 timestamp);

  struct TimestampedFeeUpdate {
    uint128 linkPerUnitGas;
    uint128 timestamp;
  }

  /**
   * @notice Set a fee updater.
   * @param feeUpdater The address of the fee updater.
   */
  function setFeeUpdater(address feeUpdater) external;

  /**
   * @notice Remove a fee updater.
   * @param feeUpdater The address of the fee updater.
   */
  function removeFeeUpdater(address feeUpdater) external;

  /**
   * @notice Update the fee for a given token and destination chain.
   * @param feeUpdates The fee updates to apply.
   */
  function updateFees(GE.FeeUpdate[] memory feeUpdates) external;

  /**
   * @notice Get the fee for a given token and destination chain.
   * @param token The token to get the fee for.
   * @param destChainId The destination chain to get the fee for.
   * @return fee The fee for the given token and destination chain.
   */
  function getFee(address token, uint64 destChainId) external view returns (uint128 fee);

  /**
   * @notice Get the staleness threshold.
   * @return stalenessThreshold The staleness threshold.
   */
  function getStalenessThreshold() external view returns (uint128 stalenessThreshold);

  /**
   * @notice Withdraw a specified amount of any token from the contract.
   * @param token The token to withdraw.
   * @param to The address to send the withdrawn tokens to.
   * @param amount The amount of tokens to withdraw.
   */
  function withdrawToken(
    address token,
    address to,
    uint256 amount
  ) external;
}
