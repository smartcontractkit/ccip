// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../models/Client.sol";
import {Internal} from "../../models/Internal.sol";

interface IPriceRegistry {
  error TokenNotSupported(address token);
  error NotAFeeToken(address token);
  error ChainNotSupported(uint64 chain);
  error OnlyCallableByUpdaterOrOwner();
  error StaleGasPrice(uint64 destChainId, uint256 threshold, uint256 timePassed);
  error StaleTokenPrice(address token, uint256 threshold, uint256 timePassed);
  error InvalidStalenessThreshold();

  event PriceUpdaterSet(address indexed priceUpdater);
  event PriceUpdaterRemoved(address indexed priceUpdater);
  event FeeTokenAdded(address indexed feeToken);
  event FeeTokenRemoved(address indexed feeToken);
  event UsdPerUnitGasUpdated(uint64 indexed destChain, uint256 value, uint256 timestamp);
  event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);

  /// @notice Value with a timestamp, used for gas prices and token prices.
  struct TimestampedUint128Value {
    uint128 value; // -------┐
    uint128 timestamp; // ---┘
  }

  /// @notice Adds new priceUpdaters and remove existing ones.
  /// @param priceUpdatersToAdd The addresses of the priceUpdaters that are now allowed
  /// to send fee updates.
  /// @param priceUpdatersToRemove The addresses of the priceUpdaters that are no longer allowed
  /// to send fee updates.
  function applyPriceUpdatersUpdates(address[] memory priceUpdatersToAdd, address[] memory priceUpdatersToRemove)
    external;

  /// @notice Get the list of price updaters.
  /// @return priceUpdaters The price updaters.
  function getPriceUpdaters() external view returns (address[] memory priceUpdaters);

  /// @notice Add and remove tokens from feeTokens set.
  /// @param feeTokensToAdd The addresses of the tokens which are now considered fee tokens
  /// and can be used to calculate fees.
  /// @param feeTokensToRemove The addresses of the tokens which are no longer considered feeTokens.
  function applyFeeTokensUpdates(address[] memory feeTokensToAdd, address[] memory feeTokensToRemove) external;

  /// @notice Get the list of fee tokens.
  /// @return feeTokens The tokens set as fee tokens.
  function getFeeTokens() external view returns (address[] memory feeTokens);

  /// @notice Update the price for given tokens and destination chain.
  /// @param priceUpdates The price updates to apply.
  function updatePrices(Internal.PriceUpdates memory priceUpdates) external;

  /// @notice Get the `tokenPrice` for a given token.
  /// @param token The token to get the price for.
  /// @return tokenPrice The tokenPrice for the given token.
  function getTokenPrice(address token) external view returns (TimestampedUint128Value memory);

  /// @notice Get the `gasPrice` for a given destination chain ID.
  /// @param destChainId The destination chain to get the price for.
  /// @return gasPrice The gasPrice for the given destination chain ID.
  function getDestinationChainGasPrice(uint64 destChainId) external view returns (TimestampedUint128Value memory);

  /// @notice Get the `feeTokenBaseUnitsPerUnitGas` for a given source chain token and destination chain ID.
  /// @param feeToken The source token to get the fee for. Must be a feeToken.
  /// @param destChainId The destination chain to get the fee for.
  /// @return feeTokenBaseUnitsPerUnitGas The feeTokenBaseUnitsPerUnitGas for the given source chain token and destination chain ID.
  /// @dev Example:
  /// * The feeToken is WETH,
  /// * The destination chain cost per unit of gas is 1 GWEI (1_000_000_000),
  /// * The return value from this function is 1_000_000_000
  function getFeeTokenBaseUnitsPerUnitGas(address feeToken, uint64 destChainId)
    external
    view
    returns (uint256 feeTokenBaseUnitsPerUnitGas);

  /// @notice Convert a fee token amount to a link token amount.
  /// @param linkToken The link token address.
  /// @param feeToken The fee token address.
  /// @param feeTokenAmount The fee token amount.
  /// @return linkTokenAmount The link token amount.
  function convertFeeTokenAmountToLinkAmount(
    address linkToken,
    address feeToken,
    uint256 feeTokenAmount
  ) external returns (uint256 linkTokenAmount);

  /// @notice Get the staleness threshold.
  /// @return stalenessThreshold The staleness threshold.
  function getStalenessThreshold() external view returns (uint128 stalenessThreshold);
}
