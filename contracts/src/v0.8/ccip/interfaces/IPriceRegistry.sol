// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../libraries/Internal.sol";

interface IPriceRegistry {
  /// @notice Value with a timestamp, used for gas prices and token prices.
  struct TimestampedUint128Value {
    uint128 value; // -------┐
    uint128 timestamp; // ---┘
  }

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

  /// @notice Convert a given token amount to target token amount.
  /// @param fromToken The given token address.
  /// @param fromTokenAmount The given token amount.
  /// @param toToken The target token address.
  /// @return toTokenAmount The target token amount.
  function convertTokenAmount(
    address fromToken,
    uint256 fromTokenAmount,
    address toToken
  ) external returns (uint256 toTokenAmount);
}
