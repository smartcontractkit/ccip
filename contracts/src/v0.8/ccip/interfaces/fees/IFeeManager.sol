// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../models/Client.sol";
import {Internal} from "../../models/Internal.sol";

interface IFeeManager {
  error TokenOrChainNotSupported(address token, uint64 chain);
  error FeeUpdaterNotAllowed(address feeUpdater);
  error OnlyCallableByUpdaterOrOwner();
  error StaleFee(uint32 threshold, uint256 timePassed);
  error InvalidWithdrawalAddress();
  error NullAddressNotAllowed();

  event FeeUpdaterSet(address indexed feeUpdater);
  event FeeUpdaterRemoved(address indexed feeUpdater);
  event GasFeeUpdated(
    address indexed token,
    uint64 indexed destChain,
    uint128 feeTokenBaseUnitsPerUnitGas,
    uint64 timestamp
  );

  struct TimestampedFeeUpdate {
    uint128 feeTokenBaseUnitsPerUnitGas; // --┐
    uint64 timestamp; // ---------------------┘
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
  function updateFees(Internal.FeeUpdate[] memory feeUpdates) external;

  /**
   * @notice Get the `feeTokenBaseUnitsPerUnitGas` for a given source chain token and destination chain ID.
   * @param feeToken The token to get the fee for.
   * @param destChainId The destination chain to get the fee for.
   * @return feeTokenBaseUnitsPerUnitGas The feeTokenBaseUnitsPerUnitGas for the given source chain token and destination chain ID.
   * @dev Example:
   * * The feeToken is WETH,
   * * The destination chain cost per unit of gas is 1 GWEI (1_000_000_000),
   * * The return value from this function is 1_000_000_000
   */
  function getFeeTokenBaseUnitsPerUnitGas(address feeToken, uint64 destChainId)
    external
    view
    returns (uint128 feeTokenBaseUnitsPerUnitGas);

  /**
   * @notice Get the staleness threshold.
   * @return stalenessThreshold The staleness threshold.
   */
  function getStalenessThreshold() external view returns (uint32 stalenessThreshold);

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
