// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPriceRegistry} from "../interfaces/prices/IPriceRegistry.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Internal} from "../models/Internal.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice The PriceRegistry contract responsibility is to store the current gas price in USD for a given destination chain,
/// and the price of a token in USD allowing the owner or priceUpdater to update this value.
contract PriceRegistry is IPriceRegistry, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.AddressSet;
  /// @dev The price, in USD, of 1 unit of gas for a given destination chain.
  /// @dev 1e18 is 1 USD. Examples:
  ///     Very Expensive:   1 unit of gas costs 1 USD                  -> 1e18
  ///     Expensive:        1 unit of gas costs 0.1 USD                -> 1e17
  ///     Cheap:            1 unit of gas costs 0.000001 USD           -> 1e12
  mapping(uint64 => TimestampedUint128Value) private s_usdPerUnitGasByDestChainId;

  /// @dev USD per full token, in base units 1e18.
  /// @dev Example:
  ///     1 USDC = 1.00 USD per token -> 1e18
  ///     1 LINK = 5.00 USD per token -> 5e18
  ///     1 ETH = 2,000 USD per token -> 2_000e18
  mapping(address => TimestampedUint128Value) private s_usdPerToken;

  // Price updaters are allowed to update the prices.
  EnumerableSet.AddressSet private s_priceUpdaters;
  // Subset of tokens which prices tracked by this registry which are fee tokens.
  EnumerableSet.AddressSet private s_feeTokens;
  // The amount of time a price can be stale before it is considered invalid.
  uint32 private immutable i_stalenessThreshold;

  constructor(
    Internal.PriceUpdates memory priceUpdates,
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    uint32 stalenessThreshold
  ) {
    _updatePrices(priceUpdates);
    _applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
    _applyFeeTokensUpdates(feeTokens, new address[](0));
    if (stalenessThreshold == 0) revert InvalidStalenessThreshold();
    i_stalenessThreshold = stalenessThreshold;
  }

  // @inheritdoc IPriceRegistry
  function applyPriceUpdatersUpdates(address[] memory priceUpdatersToAdd, address[] memory priceUpdatersToRemove)
    external
    override
    onlyOwner
  {
    _applyPriceUpdatersUpdates(priceUpdatersToAdd, priceUpdatersToRemove);
  }

  // @inheritdoc IPriceRegistry
  function getPriceUpdaters() external view override returns (address[] memory priceUpdaters) {
    priceUpdaters = new address[](s_priceUpdaters.length());
    for (uint256 i = 0; i < s_priceUpdaters.length(); ++i) {
      priceUpdaters[i] = s_priceUpdaters.at(i);
    }
  }

  // @inheritdoc IPriceRegistry
  function applyFeeTokensUpdates(address[] memory feeTokensToAdd, address[] memory feeTokensToRemove)
    external
    override
    onlyOwner
  {
    _applyFeeTokensUpdates(feeTokensToAdd, feeTokensToRemove);
  }

  // @inheritdoc IPriceRegistry
  function getFeeTokens() external view override returns (address[] memory feeTokens) {
    feeTokens = new address[](s_feeTokens.length());
    for (uint256 i = 0; i < s_feeTokens.length(); ++i) {
      feeTokens[i] = s_feeTokens.at(i);
    }
  }

  // @inheritdoc IPriceRegistry
  function updatePrices(Internal.PriceUpdates memory priceUpdates) external override requireUpdaterOrOwner {
    _updatePrices(priceUpdates);
  }

  // @inheritdoc IPriceRegistry
  function getTokenPrice(address token) external view override returns (TimestampedUint128Value memory) {
    return s_usdPerToken[token];
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
  function getFeeTokenBaseUnitsPerUnitGas(address feeToken, uint64 destChainId)
    external
    view
    override
    returns (uint256 feeTokenBaseUnitsPerUnitGas)
  {
    if (!s_feeTokens.contains(feeToken)) revert NotAFeeToken(feeToken);

    TimestampedUint128Value memory gasPrice = s_usdPerUnitGasByDestChainId[destChainId];
    if (gasPrice.timestamp == 0 || gasPrice.value == 0) revert ChainNotSupported(destChainId);
    uint256 timePassed = block.timestamp - gasPrice.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleGasPrice(destChainId, i_stalenessThreshold, timePassed);

    TimestampedUint128Value memory feeTokenPrice = s_usdPerToken[feeToken];
    if (feeTokenPrice.timestamp == 0 || feeTokenPrice.value == 0) revert TokenNotSupported(feeToken);
    timePassed = block.timestamp - feeTokenPrice.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleTokenPrice(feeToken, i_stalenessThreshold, timePassed);

    return (uint256(gasPrice.value) * 1e18) / uint256(feeTokenPrice.value);
  }

  /// @inheritdoc IPriceRegistry
  /// @dev this function assumed that no more than 1e72 dollar, or type(uint240).max, is
  /// sent as payment. If more is sent, the multiplication of feeTokenAmount and feeTokenValue
  /// will overflow. Since there isn't even close to 1e72 dollars in the world economy this is safe.
  /// @dev the result is a uint96 which is can store more than all the link that exists and is
  /// therefore considered safe.
  function convertFeeTokenAmountToLinkAmount(
    address linkToken,
    address feeToken,
    uint256 feeTokenAmount
  ) external view override returns (uint96 linkTokenAmount) {
    if (!s_feeTokens.contains(feeToken)) revert NotAFeeToken(feeToken);

    TimestampedUint128Value memory feeTokenPrice = s_usdPerToken[feeToken];
    if (feeTokenPrice.timestamp == 0 || feeTokenPrice.value == 0) revert TokenNotSupported(feeToken);
    uint256 feeTokenTimePassed = block.timestamp - feeTokenPrice.timestamp;
    if (feeTokenTimePassed > i_stalenessThreshold)
      revert StaleTokenPrice(feeToken, i_stalenessThreshold, feeTokenTimePassed);

    TimestampedUint128Value memory linkTokenPrice = s_usdPerToken[linkToken];
    if (linkTokenPrice.timestamp == 0 || linkTokenPrice.value == 0) revert TokenNotSupported(linkToken);
    uint256 linkTokenTimePassed = block.timestamp - linkTokenPrice.timestamp;
    if (linkTokenTimePassed > i_stalenessThreshold)
      revert StaleTokenPrice(linkToken, i_stalenessThreshold, linkTokenTimePassed);

    /// Example:
    /// feeTokenAmount:   1e18      // 1 ETH
    /// ETH:              2_000e18
    /// LINK:             5e18
    /// return:           1e18 * 2_000e18 / 5e18 = 400e18 (400 LINK)
    return uint96((feeTokenAmount * uint256(feeTokenPrice.value)) / uint256(linkTokenPrice.value));
  }

  /// @inheritdoc IPriceRegistry
  function getStalenessThreshold() external view override returns (uint128) {
    return i_stalenessThreshold;
  }

  /// @notice Adds new priceUpdaters and remove existing ones.
  /// @param priceUpdatersToAdd The addresses of the priceUpdaters that are now allowed
  /// to send fee updates.
  /// @param priceUpdatersToRemove The addresses of the priceUpdaters that are no longer allowed
  /// to send fee updates.
  function _applyPriceUpdatersUpdates(address[] memory priceUpdatersToAdd, address[] memory priceUpdatersToRemove)
    private
  {
    for (uint256 i = 0; i < priceUpdatersToAdd.length; ++i) {
      if (s_priceUpdaters.add(priceUpdatersToAdd[i])) {
        emit PriceUpdaterSet(priceUpdatersToAdd[i]);
      }
    }
    for (uint256 i = 0; i < priceUpdatersToRemove.length; ++i) {
      if (s_priceUpdaters.remove(priceUpdatersToRemove[i])) {
        emit PriceUpdaterRemoved(priceUpdatersToRemove[i]);
      }
    }
  }

  /// @notice Add and remove tokens from feeTokens set.
  /// @param feeTokensToAdd The addresses of the tokens which are now considered fee tokens
  /// and can be used to calculate fees.
  /// @param feeTokensToRemove The addresses of the tokens which are no longer considered feeTokens.
  function _applyFeeTokensUpdates(address[] memory feeTokensToAdd, address[] memory feeTokensToRemove) private {
    for (uint256 i = 0; i < feeTokensToAdd.length; ++i) {
      if (s_feeTokens.add(feeTokensToAdd[i])) {
        emit FeeTokenAdded(feeTokensToAdd[i]);
      }
    }
    for (uint256 i = 0; i < feeTokensToRemove.length; ++i) {
      if (s_feeTokens.remove(feeTokensToRemove[i])) {
        emit FeeTokenRemoved(feeTokensToRemove[i]);
      }
    }
  }

  /// @notice Updates all prices in the priceUpdates struct.
  /// @param priceUpdates The struct containing all the price updates.
  function _updatePrices(Internal.PriceUpdates memory priceUpdates) private {
    for (uint256 i = 0; i < priceUpdates.tokenPriceUpdates.length; ++i) {
      Internal.TokenPriceUpdate memory update = priceUpdates.tokenPriceUpdates[i];
      s_usdPerToken[update.sourceToken] = TimestampedUint128Value({
        value: update.usdPerToken,
        timestamp: uint128(block.timestamp)
      });
      emit UsdPerTokenUpdated(update.sourceToken, update.usdPerToken, block.timestamp);
    }

    if (priceUpdates.destChainId != 0) {
      s_usdPerUnitGasByDestChainId[priceUpdates.destChainId] = TimestampedUint128Value({
        value: priceUpdates.usdPerUnitGas,
        timestamp: uint128(block.timestamp)
      });
      emit UsdPerUnitGasUpdated(priceUpdates.destChainId, priceUpdates.usdPerUnitGas, block.timestamp);
    }
  }

  /// @notice Require that the caller is the owner or a fee updater.
  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_priceUpdaters.contains(msg.sender)) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
