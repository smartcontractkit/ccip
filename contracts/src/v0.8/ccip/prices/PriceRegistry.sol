// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPriceRegistry} from "../interfaces/prices/IPriceRegistry.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Internal} from "../models/Internal.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice The PriceRegistry contract responsibility is to store the current gas price in USD for a given destination chain,
/// and the price of a fee token in USD allowing the owner or priceUpdater to update this value.
contract PriceRegistry is IPriceRegistry, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;
  /// @dev The price, in USD, of 1 unit of gas for a given destination chain.
  /// @dev 1e18 is 1 USD. Examples:
  ///   * Very Expensive:   1 unit of gas costs 1 USD                  -> 1e18
  ///   * Expensive:        1 unit of gas costs 0.1 USD                -> 1e17
  ///   * Cheap:            1 unit of gas costs 0.000001 USD           -> 1e12
  mapping(uint64 => TimestampedUint128Value) private s_usdPerUnitGasByDestChainId;

  /// @dev USD per full fee token, in base units 1e18.
  /// @dev Example:
  ///   * 1 USDC = 1.00 USD per token -> 1e18
  ///   * 1 LINK = 5.00 USD per token -> 5e18
  ///   * 1 ETH = 2,000 USD per token -> 2_000e18
  mapping(address => TimestampedUint128Value) private s_usdPerFeeToken;

  // Price updaters are allowed to update the prices.
  EnumerableSet.AddressSet private s_priceUpdaters;
  // The amount of time a price can be stale before it is considered invalid.
  uint32 private immutable i_stalenessThreshold;

  constructor(
    Internal.PriceUpdates memory priceUpdates,
    address[] memory priceUpdaters,
    uint32 stalenessThreshold
  ) {
    _updatePrices(priceUpdates);
    _addPriceUpdaters(priceUpdaters);
    i_stalenessThreshold = stalenessThreshold;
  }

  // @inheritdoc IPriceRegistry
  function addPriceUpdaters(address[] memory priceUpdaters) external override onlyOwner {
    _addPriceUpdaters(priceUpdaters);
  }

  // @inheritdoc IPriceRegistry
  function removePriceUpdaters(address[] memory priceUpdaters) external override onlyOwner {
    _removePriceUpdaters(priceUpdaters);
  }

  // @inheritdoc IPriceRegistry
  function getPriceUpdaters() external view override returns (address[] memory priceUpdaters) {
    priceUpdaters = new address[](s_priceUpdaters.length());
    for (uint256 i = 0; i < s_priceUpdaters.length(); i++) {
      priceUpdaters[i] = s_priceUpdaters.at(i);
    }
  }

  // @inheritdoc IPriceRegistry
  function updatePrices(Internal.PriceUpdates memory priceUpdates) external override requireUpdaterOrOwner {
    _updatePrices(priceUpdates);
  }

  // @inheritdoc IPriceRegistry
  function getFeeTokenPrice(address token) external view override returns (TimestampedUint128Value memory) {
    return s_usdPerFeeToken[token];
  }

  // @inheritdoc IPriceRegistry
  function getDestinationChainGasPrice(uint64 destChainId)
    external
    view
    override
    returns (TimestampedUint128Value memory)
  {
    return s_usdPerUnitGasByDestChainId[destChainId];
  }

  // @inheritdoc IPriceRegistry
  function getFeeTokenBaseUnitsPerUnitGas(address token, uint64 destChainId)
    external
    view
    override
    returns (uint256 feeTokenBaseUnitsPerUnitGas)
  {
    TimestampedUint128Value memory gasPrice = s_usdPerUnitGasByDestChainId[destChainId];
    if (gasPrice.timestamp == 0 || gasPrice.value == 0) revert ChainNotSupported(destChainId);
    uint256 timePassed = block.timestamp - gasPrice.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleGasPrice(destChainId, i_stalenessThreshold, timePassed);

    TimestampedUint128Value memory feeTokenPrice = s_usdPerFeeToken[token];
    if (feeTokenPrice.timestamp == 0 || feeTokenPrice.value == 0) revert TokenNotSupported(token);
    timePassed = block.timestamp - feeTokenPrice.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleTokenPrice(token, i_stalenessThreshold, timePassed);

    return (uint256(gasPrice.value) * 1e18) / uint256(feeTokenPrice.value);
  }

  // @inheritdoc IPriceRegistry
  function convertFeeTokenAmountToLinkAmount(
    address linkToken,
    address feeToken,
    uint256 feeTokenAmount
  ) external view override returns (uint256 linkTokenAmount) {
    TimestampedUint128Value memory feeTokenPrice = s_usdPerFeeToken[feeToken];
    if (feeTokenPrice.timestamp == 0 || feeTokenPrice.value == 0) revert TokenNotSupported(feeToken);
    uint256 feeTokenTimePassed = block.timestamp - feeTokenPrice.timestamp;
    if (feeTokenTimePassed > i_stalenessThreshold)
      revert StaleTokenPrice(feeToken, i_stalenessThreshold, feeTokenTimePassed);

    TimestampedUint128Value memory linkTokenPrice = s_usdPerFeeToken[linkToken];
    if (linkTokenPrice.timestamp == 0 || linkTokenPrice.value == 0) revert TokenNotSupported(linkToken);
    uint256 linkTokenTimePassed = block.timestamp - linkTokenPrice.timestamp;
    if (linkTokenTimePassed > i_stalenessThreshold)
      revert StaleTokenPrice(linkToken, i_stalenessThreshold, linkTokenTimePassed);

    /// Example:
    /// feeTokenAmount:   1e18      // 1 ETH
    /// ETH:              2_000e18
    /// LINK:             5e18
    /// conversionRate:   (2_000e18 * 1e18) / 5e18 = 400e18
    /// return:           (1e18 * 400e18) / 1e18 = 400e18 (400 LINK)
    uint256 conversionRate = (uint256(feeTokenPrice.value) * 1e18) / uint256(linkTokenPrice.value);
    return (feeTokenAmount * conversionRate) / 1e18;
  }

  // @inheritdoc IPriceRegistry
  function getStalenessThreshold() external view override returns (uint128) {
    return i_stalenessThreshold;
  }

  /// @notice Adds new price updaters.
  /// @param priceUpdaters The addresses of the priceUpdaters that are now allowed
  /// to send fee updates.
  function _addPriceUpdaters(address[] memory priceUpdaters) private {
    for (uint256 i = 0; i < priceUpdaters.length; ++i) {
      s_priceUpdaters.add(priceUpdaters[i]);
      emit PriceUpdaterSet(priceUpdaters[i]);
    }
  }

  /// @notice Removes price updaters.
  /// @param priceUpdaters The addresses of the priceUpdaters that are no longer allowed
  /// to send fee updates.
  function _removePriceUpdaters(address[] memory priceUpdaters) private {
    for (uint256 i = 0; i < priceUpdaters.length; ++i) {
      if (s_priceUpdaters.remove(priceUpdaters[i])) {
        emit PriceUpdaterRemoved(priceUpdaters[i]);
      }
    }
  }

  /// @notice Updates all prices in the priceUpdates struct.
  /// @param priceUpdates The struct containing all the price updates.
  function _updatePrices(Internal.PriceUpdates memory priceUpdates) private {
    for (uint256 i = 0; i < priceUpdates.feeTokenPriceUpdates.length; ++i) {
      _updateUsdPerFeeToken(priceUpdates.feeTokenPriceUpdates[i]);
    }
    if (priceUpdates.destChainId != 0) {
      _updateUsdPerGasUnitByDestChainId(priceUpdates.destChainId, priceUpdates.usdPerUnitGas);
    }
  }

  /// @notice Updates the USD per gas unit for a given destination chain.
  /// @param destChainId The destination chain id.
  /// @param usdPerUnitGas The gas price in USD per unit gas.
  function _updateUsdPerGasUnitByDestChainId(uint64 destChainId, uint128 usdPerUnitGas) private {
    s_usdPerUnitGasByDestChainId[destChainId] = TimestampedUint128Value({
      value: usdPerUnitGas,
      timestamp: uint128(block.timestamp)
    });
    emit UsdPerUnitGasUpdated(destChainId, usdPerUnitGas, block.timestamp);
  }

  /// @notice Updates the USD per fee token.
  /// @param update The struct containing the update.
  function _updateUsdPerFeeToken(Internal.FeeTokenPriceUpdate memory update) private {
    s_usdPerFeeToken[update.sourceFeeToken] = TimestampedUint128Value({
      value: update.usdPerFeeToken,
      timestamp: uint128(block.timestamp)
    });
    emit UsdPerFeeTokenUpdated(update.sourceFeeToken, update.usdPerFeeToken, block.timestamp);
  }

  /// @notice Require that the caller is the owner or a fee updater.
  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_priceUpdaters.contains(msg.sender)) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
