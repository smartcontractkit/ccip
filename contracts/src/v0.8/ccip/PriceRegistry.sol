// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../shared/interfaces/ITypeAndVersion.sol";
import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {OwnerIsCreator} from "./../shared/access/OwnerIsCreator.sol";
import {AggregatorV3Interface} from "./../shared/interfaces/AggregatorV3Interface.sol";
import {Internal} from "./libraries/Internal.sol";
import {USDPriceWith18Decimals} from "./libraries/USDPriceWith18Decimals.sol";

import {EnumerableSet} from "../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice The PriceRegistry contract responsibility is to store the current gas price in USD for a given destination chain,
/// and the price of a token in USD allowing the owner or priceUpdater to update this value.
contract PriceRegistry is IPriceRegistry, OwnerIsCreator, ITypeAndVersion {
  using EnumerableSet for EnumerableSet.AddressSet;
  using USDPriceWith18Decimals for uint224;

  /// @notice Token price data feed update
  struct TokenPriceFeedUpdate {
    address sourceToken; // Source token to update feed for
    IPriceRegistry.TokenPriceFeedConfig feedConfig; // Feed config update data
  }

  struct EVMDestChainExecConfig {
    Internal.TimestampedPackedUint224 gasPrice; // Gas price for the destination chain
    uint16 maxNumberOfTokensPerMsg; //           │ Maximum number of distinct ERC20 token transferred per message
    uint32 maxDataBytes; //                      │ Maximum payload data size in bytes
    uint32 maxPerMsgGasLimit; //                 │ Maximum gas limit for messages targeting EVMs
    uint32 destGasOverhead; //                   │ Gas charged on top of the gasLimit to cover destination chain costs
    uint16 destGasPerPayloadByte; //             │ Destination chain gas charged for passing each byte of `data` payload to receiver
    uint32 destDataAvailabilityOverheadGas; //   | Extra data availability gas charged on top of the message, e.g. for OCR
    uint16 destGasPerDataAvailabilityByte; //    | Amount of gas to charge per byte of message data that needs availability
    uint16 destDataAvailabilityMultiplierBps; // │ Multiplier for data availability gas, multiples of bps, or 0.0001
    uint32 defaultTokenDestGasOverhead; // ──────╯ Default gas charged to execute the token transfer on the destination chain
    uint32 defaultTokenDestBytesOverhead; // ────╮ Default extra data availability bytes charged per token transfer
    uint64 defaultTxGasLimit; //                 │ Default gas limit for a tx
    uint64 gasMultiplierWeiPerEth; //            │ Multiplier for gas costs, 1e18 based so 11e17 = 10% extra cost.
    bool enforceOutOfOrder;
  }

  struct EVMDestTokenExecConfig {
    uint32 defaultTokenDestGasOverhead; //       │ Default gas charged to execute the token transfer on the destination chain
    //                                           │ Default data availability bytes that are returned from the source pool and sent
    uint32 defaultTokenDestBytesOverhead; //     | to the destination pool. Must be >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
  }

  error TokenNotSupported(address token);
  error ChainNotSupported(uint64 chain);
  error OnlyCallableByUpdaterOrOwner();
  error StaleGasPrice(uint64 destChainSelector, uint256 threshold, uint256 timePassed);
  error StaleTokenPrice(address token, uint256 threshold, uint256 timePassed);
  error InvalidStalenessThreshold();
  error DataFeedValueOutOfUint224Range();

  event PriceUpdaterSet(address indexed priceUpdater);
  event PriceUpdaterRemoved(address indexed priceUpdater);
  event FeeTokenAdded(address indexed feeToken);
  event FeeTokenRemoved(address indexed feeToken);
  event UsdPerUnitGasUpdated(uint64 indexed destChain, uint256 value, uint256 timestamp);
  event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp);
  event PriceFeedPerTokenUpdated(address indexed token, IPriceRegistry.TokenPriceFeedConfig priceFeedConfig);

  string public constant override typeAndVersion = "PriceRegistry 1.6.0-dev";

  /// @dev The gas price per unit of gas for a given destination chain, in USD with 18 decimals.
  /// Multiple gas prices can be encoded into the same value. Each price takes {Internal.GAS_PRICE_BITS} bits.
  /// For example, if Optimism is the destination chain, gas price can include L1 base fee and L2 gas price.
  /// Logic to parse the price components is chain-specific, and should live in OnRamp.
  /// @dev Price of 1e18 is 1 USD. Examples:
  ///     Very Expensive:   1 unit of gas costs 1 USD                  -> 1e18
  ///     Expensive:        1 unit of gas costs 0.1 USD                -> 1e17
  ///     Cheap:            1 unit of gas costs 0.000001 USD           -> 1e12
  mapping(uint64 destChainSelector => EVMDestChainExecConfig) private
    s_evmDestChainExecConfig;
  /// @dev The token transfer fee config that can be set by the owner or fee admin
  mapping(uint64 destChainSelector => mapping(address token => EVMDestTokenExecConfig tokenExecConfig)) internal
    s_evmDestTokenExecConfig;

  /// @dev The price, in USD with 18 decimals, per 1e18 of the smallest token denomination.
  /// @dev Price of 1e18 represents 1 USD per 1e18 token amount.
  ///     1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
  ///     1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
  ///     1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
  mapping(address token => Internal.TimestampedPackedUint224 price) private s_usdPerToken;

  /// @dev Stores the price data feed configurations per token.
  mapping(address token => IPriceRegistry.TokenPriceFeedConfig dataFeedAddress) private s_usdPriceFeedsPerToken;

  // Price updaters are allowed to update the prices.
  EnumerableSet.AddressSet private s_priceUpdaters;
  // Subset of tokens which prices tracked by this registry which are fee tokens.
  EnumerableSet.AddressSet private s_feeTokens;
  // The amount of time a price can be stale before it is considered invalid.
  uint32 private immutable i_stalenessThreshold;

  constructor(
    address[] memory priceUpdaters,
    address[] memory feeTokens,
    uint32 stalenessThreshold,
    TokenPriceFeedUpdate[] memory tokenPriceFeeds
  ) {
    _applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
    _applyFeeTokensUpdates(feeTokens, new address[](0));
    _updateTokenPriceFeeds(tokenPriceFeeds);
    if (stalenessThreshold == 0) revert InvalidStalenessThreshold();
    i_stalenessThreshold = stalenessThreshold;
  }

  // ================================================================
  // │                     Price calculations                       │
  // ================================================================

  // @inheritdoc IPriceRegistry
  function getTokenPrice(address token) public view override returns (Internal.TimestampedPackedUint224 memory) {
    IPriceRegistry.TokenPriceFeedConfig memory priceFeedConfig = s_usdPriceFeedsPerToken[token];
    if (priceFeedConfig.dataFeedAddress == address(0)) {
      return s_usdPerToken[token];
    }

    (uint224 price, uint32 timestamp) = _getTokenPriceFromDataFeed(priceFeedConfig);
    return Internal.TimestampedPackedUint224({value: price, timestamp: timestamp});
  }

  // @inheritdoc IPriceRegistry
  function getValidatedTokenPrice(address token) external view override returns (uint224) {
    return _getValidatedTokenPrice(token);
  }

  // @inheritdoc IPriceRegistry
  function getTokenPrices(address[] calldata tokens)
    external
    view
    override
    returns (Internal.TimestampedPackedUint224[] memory)
  {
    uint256 length = tokens.length;
    Internal.TimestampedPackedUint224[] memory tokenPrices = new Internal.TimestampedPackedUint224[](length);
    for (uint256 i = 0; i < length; ++i) {
      tokenPrices[i] = getTokenPrice(tokens[i]);
    }
    return tokenPrices;
  }

  // @inheritdoc IPriceRegistry
  function getTokenPriceFeedConfig(address token)
    external
    view
    override
    returns (IPriceRegistry.TokenPriceFeedConfig memory)
  {
    return s_usdPriceFeedsPerToken[token];
  }

  /// @notice Get the staleness threshold.
  /// @return stalenessThreshold The staleness threshold.
  function getStalenessThreshold() external view returns (uint128) {
    return i_stalenessThreshold;
  }

  // @inheritdoc IPriceRegistry
  function getDestinationChainGasPrice(uint64 destChainSelector)
    external
    view
    override
    returns (Internal.TimestampedPackedUint224 memory)
  {
    return s_evmDestChainExecConfig[destChainSelector];
  }

  function getEVMExecFee(uint64 destination, Client.EVM2AnyMessage message)
    internal
    view
    returns (uint256)
  {
    // PARSE
    // We know extraArgs are EVM.
    bytes4 extraArgsTag = bytes4(extraArgs);
    Client.EVMExtraArgsV2 memory extraArgs;
    if (extraArgsTag == Client.EVM_EXTRA_ARGS_V2_TAG) {
      extraArgs = abi.decode(extraArgs[4:], (Client.EVMExtraArgsV2));
    } else if (extraArgsTag == Client.EVM_EXTRA_ARGS_V1_TAG) {
      // EVMExtraArgsV1 originally included a second boolean (strict) field which has been deprecated.
      // Clients may still include it but it will be ignored.
      extraArgs = Client.EVMExtraArgsV2({gasLimit: abi.decode(extraArgs[4:], (uint256)), allowOutOfOrderExecution: false});
    }

    // VALIDATE
    EVMDestChainExecConfig config = s_evmDestChainExecConfig[destination];
    if (config.defaultTxGasLimit == 0) { // Could use some explicit enabled bool as well.
      // Chain not supported
      revert ChainNotSupported(destination);
    }
    // Validate the message with various checks
    uint256 numberOfTokens = message.tokenAmounts.length;
    if (message.data.length > uint256(config.maxDataBytes)) {
      revert MessageTooLarge(uint256(config.maxDataBytes), dataLength);
    }
    if (extraArgs.gasLimit > uint256(config.maxPerMsgGasLimit)) revert MessageGasLimitTooHigh();
    if (numberOfTokens > uint256(config.maxNumberOfTokensPerMsg)) revert UnsupportedNumberOfTokens();
    // TODO: Validate receiver.

    // COMPUTE
    uint256 gasEstimate = extraArgs.gasLimit
    uint256 tokenTransferBytesOverhead = 0;
    for (uint256 i = 0; i < numberOfTokens; ++i) {
        // Calculate the gas cost of the token transfer on the destination chain.
        EVMDestTokenExecConfig tokenConfig = s_evmDestTokenExecConfig[destination][message.tokenAmounts[i].token];
        if (tokenConfig.defaultTokenDestBytesOverhead == 0) {
          gasEstimate += config.defaultTokenDestGasOverhead;
          dataSize += config.defaultTokenDestBytesOverhead
        } else {
          gasEstimate += config.defaultTokenDestGasOverhead;
          dataSize += config.defaultTokenDestBytesOverhead
        }
    }
    // Calculate data availability cost in USD with 36 decimals. Data availability cost exists on rollups that need to post
    // transaction calldata onto another storage layer, e.g. Eth mainnet, incurring additional storage gas costs.
    uint256 dataAvailabilityCost = 0;
    // Only calculate data availability cost if data availability multiplier is non-zero.
    // The multiplier should be set to 0 if destination chain does not charge data availability cost.
    if (destChainDynamicConfig.destDataAvailabilityMultiplierBps > 0) {
      dataAvailabilityCost = _getDataAvailabilityCost(
        destChainSelector,
        // Parse the data availability gas price stored in the higher-order 112 bits of the encoded gas price.
        uint112(packedGasPrice >> Internal.GAS_PRICE_BITS),
        message.data.length,
        numberOfTokens,
        tokenTransferBytesOverhead
      );
    }

    // Calculate execution gas fee on destination chain in USD with 36 decimals.
    // We add the message gas limit, the overhead gas, the gas of passing message data to receiver, and token transfer gas together.
    // We then multiply this gas total with the gas multiplier and gas price, converting it into USD with 36 decimals.
    // uint112(packedGasPrice) = executionGasPrice
    uint256 executionCost = uint112(config.gasPrice)
      * (
        gasLimit + config.destGasOverhead
        + (message.data.length * config.destGasPerPayloadByte) + tokenTransferGas
      ) * config.gasMultiplierWeiPerEth;
    return executionCost + dataAvailabilityCost;
  }

  function getExecFee(uint64 destination, Client.EVM2AnyMessage message)
    external
    view
    override
    returns (uint256)
  {
    // ExtraArgs tells us which family of chains we are dealing with.
    bytes4 extraArgsTag = bytes4(extraArgs);
    if (extraArgs.length == 0 || extraArgsTag == Client.EVM_EXTRA_ARGS_V2_TAG || extraArgsTag == Client.EVM_EXTRA_ARGS_V1_TAG) {
        return getEVMExecFee(destination, message);
    }
    // TODO: add other families here
    revert InvalidExtraArgsTag();
  }

  /// @inheritdoc IPriceRegistry
  /// @dev this function assumes that no more than 1e59 dollars are sent as payment.
  /// If more is sent, the multiplication of feeTokenAmount and feeTokenValue will overflow.
  /// Since there isn't even close to 1e59 dollars in the world economy this is safe.
  function convertTokenAmount(
    address fromToken,
    uint256 fromTokenAmount,
    address toToken
  ) external view override returns (uint256) {
    /// Example:
    /// fromTokenAmount:   1e18      // 1 ETH
    /// ETH:               2_000e18
    /// LINK:              5e18
    /// return:            1e18 * 2_000e18 / 5e18 = 400e18 (400 LINK)
    return (fromTokenAmount * _getValidatedTokenPrice(fromToken)) / _getValidatedTokenPrice(toToken);
  }

  /// @notice Gets the token price for a given token and revert if the token is either
  /// not supported or the price is stale.
  /// @param token The address of the token to get the price for
  /// @return the token price
  function _getValidatedTokenPrice(address token) internal view returns (uint224) {
    Internal.TimestampedPackedUint224 memory tokenPrice = getTokenPrice(token);
    if (tokenPrice.timestamp == 0 || tokenPrice.value == 0) revert TokenNotSupported(token);
    uint256 timePassed = block.timestamp - tokenPrice.timestamp;
    if (timePassed > i_stalenessThreshold) revert StaleTokenPrice(token, i_stalenessThreshold, timePassed);
    return tokenPrice.value;
  }

  /// @notice Gets the token price from a data feed address, rebased to the same units as s_usdPerToken
  /// @param priceFeedConfig token data feed configuration with valid data feed address (used to retrieve price & timestamp)
  /// @return value data feed answer value (rebased to s_usdPerToken units)
  /// @return timestamp data feed last updated timestamp
  function _getTokenPriceFromDataFeed(IPriceRegistry.TokenPriceFeedConfig memory priceFeedConfig)
    internal
    view
    returns (uint224 value, uint32 timestamp)
  {
    AggregatorV3Interface dataFeedContract = AggregatorV3Interface(priceFeedConfig.dataFeedAddress);
    (
      /* uint80 roundID */
      ,
      int256 dataFeedAnswer,
      /* uint startedAt */
      ,
      /* uint256 updatedAt */
      ,
      /* uint80 answeredInRound */
    ) = dataFeedContract.latestRoundData();

    if (dataFeedAnswer < 0) {
      revert DataFeedValueOutOfUint224Range();
    }
    uint256 rebasedValue = uint256(dataFeedAnswer);

    // Rebase formula for units in smallest token denomination: usdValue * (1e18 * 1e18) / 1eTokenDecimals
    // feedValue * (10 ** (18 - feedDecimals)) * (10 ** (18 - erc20Decimals))
    // feedValue * (10 ** ((18 - feedDecimals) + (18 - erc20Decimals)))
    // feedValue * (10 ** (36 - feedDecimals - erc20Decimals))
    // feedValue * (10 ** (36 - (feedDecimals + erc20Decimals)))
    // feedValue * (10 ** (36 - excessDecimals))
    // If excessDecimals > 36 => flip it to feedValue / (10 ** (excessDecimals - 36))

    uint8 excessDecimals = dataFeedContract.decimals() + priceFeedConfig.tokenDecimals;

    if (excessDecimals > 36) {
      rebasedValue /= 10 ** (excessDecimals - 36);
    } else {
      rebasedValue *= 10 ** (36 - excessDecimals);
    }

    if (rebasedValue > type(uint224).max) {
      revert DataFeedValueOutOfUint224Range();
    }
    return (uint224(rebasedValue), uint32(block.timestamp));
  }

  // ================================================================
  // │                         Fee tokens                           │
  // ================================================================

  /// @notice Get the list of fee tokens.
  /// @return The tokens set as fee tokens.
  function getFeeTokens() external view returns (address[] memory) {
    return s_feeTokens.values();
  }

  /// @notice Add and remove tokens from feeTokens set.
  /// @param feeTokensToAdd The addresses of the tokens which are now considered fee tokens
  /// and can be used to calculate fees.
  /// @param feeTokensToRemove The addresses of the tokens which are no longer considered feeTokens.
  function applyFeeTokensUpdates(
    address[] memory feeTokensToAdd,
    address[] memory feeTokensToRemove
  ) external onlyOwner {
    _applyFeeTokensUpdates(feeTokensToAdd, feeTokensToRemove);
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

  // ================================================================
  // │                       Price updates                          │
  // ================================================================

  // @inheritdoc IPriceRegistry
  function updatePrices(Internal.PriceUpdates calldata priceUpdates) external override requireUpdaterOrOwner {
    uint256 tokenUpdatesLength = priceUpdates.tokenPriceUpdates.length;

    for (uint256 i = 0; i < tokenUpdatesLength; ++i) {
      Internal.TokenPriceUpdate memory update = priceUpdates.tokenPriceUpdates[i];
      s_usdPerToken[update.sourceToken] =
        Internal.TimestampedPackedUint224({value: update.usdPerToken, timestamp: uint32(block.timestamp)});
      emit UsdPerTokenUpdated(update.sourceToken, update.usdPerToken, block.timestamp);
    }

    uint256 gasUpdatesLength = priceUpdates.gasPriceUpdates.length;

    for (uint256 i = 0; i < gasUpdatesLength; ++i) {
      Internal.GasPriceUpdate memory update = priceUpdates.gasPriceUpdates[i];
      // TODO could validate its an EVM chain
      s_evmDestChainExecConfig[update.destChainSelector].gasPrice =
        Internal.TimestampedPackedUint224({value: update.usdPerUnitGas, timestamp: uint32(block.timestamp)});
      emit UsdPerUnitGasUpdated(update.destChainSelector, update.usdPerUnitGas, block.timestamp);
    }
  }

  /// @notice Updates the USD token price feeds for given tokens
  /// @param tokenPriceFeedUpdates Token price feed updates to apply
  function updateTokenPriceFeeds(TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates) external onlyOwner {
    _updateTokenPriceFeeds(tokenPriceFeedUpdates);
  }

  /// @notice Updates the USD token price feeds for given tokens
  /// @param tokenPriceFeedUpdates Token price feed updates to apply
  function _updateTokenPriceFeeds(TokenPriceFeedUpdate[] memory tokenPriceFeedUpdates) private {
    for (uint256 i; i < tokenPriceFeedUpdates.length; ++i) {
      TokenPriceFeedUpdate memory update = tokenPriceFeedUpdates[i];
      address sourceToken = update.sourceToken;
      IPriceRegistry.TokenPriceFeedConfig memory tokenPriceFeedConfig = update.feedConfig;

      s_usdPriceFeedsPerToken[sourceToken] = tokenPriceFeedConfig;
      emit PriceFeedPerTokenUpdated(sourceToken, tokenPriceFeedConfig);
    }
  }

  // ================================================================
  // │                           Access                             │
  // ================================================================

  /// @notice Get the list of price updaters.
  /// @return The price updaters.
  function getPriceUpdaters() external view returns (address[] memory) {
    return s_priceUpdaters.values();
  }

  /// @notice Adds new priceUpdaters and remove existing ones.
  /// @param priceUpdatersToAdd The addresses of the priceUpdaters that are now allowed
  /// to send fee updates.
  /// @param priceUpdatersToRemove The addresses of the priceUpdaters that are no longer allowed
  /// to send fee updates.
  function applyPriceUpdatersUpdates(
    address[] memory priceUpdatersToAdd,
    address[] memory priceUpdatersToRemove
  ) external onlyOwner {
    _applyPriceUpdatersUpdates(priceUpdatersToAdd, priceUpdatersToRemove);
  }

  /// @notice Adds new priceUpdaters and remove existing ones.
  /// @param priceUpdatersToAdd The addresses of the priceUpdaters that are now allowed
  /// to send fee updates.
  /// @param priceUpdatersToRemove The addresses of the priceUpdaters that are no longer allowed
  /// to send fee updates.
  function _applyPriceUpdatersUpdates(
    address[] memory priceUpdatersToAdd,
    address[] memory priceUpdatersToRemove
  ) private {
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

  /// @notice Require that the caller is the owner or a fee updater.
  modifier requireUpdaterOrOwner() {
    if (msg.sender != owner() && !s_priceUpdaters.contains(msg.sender)) revert OnlyCallableByUpdaterOrOwner();
    _;
  }
}
