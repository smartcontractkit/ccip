// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IAFN} from "../interfaces/IAFN.sol";
import {IPriceRegistry} from "../interfaces/IPriceRegistry.sol";
import {IEVM2AnyOnRamp} from "../interfaces/IEVM2AnyOnRamp.sol";

import {AggregateRateLimiter} from "../AggregateRateLimiter.sol";
import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";
import {EnumerableMapAddresses} from "../../libraries/internal/EnumerableMapAddresses.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableMap.sol";

/// @notice The onRamp is a contract that handles fee logic, NOP payments,
/// token support and an allowList. It will always be deployed 1:1:1 with a
/// commitStore and offRamp contract. These three contracts together form a
/// `lane`. A lane is an upgradable set of contracts within the non-upgradable
/// routers and are always deployed as complete set, even during upgrades.
/// This means an upgrade to an onRamp will require redeployment of the
/// commitStore and offRamp as well.
contract EVM2EVMOnRamp is IEVM2AnyOnRamp, Pausable, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.AddressToUintMap;
  using EnumerableMapAddresses for EnumerableMapAddresses.AddressToAddressMap;
  using EnumerableSet for EnumerableSet.AddressSet;
  using USDPriceWith18Decimals for uint192;

  error InvalidExtraArgsTag();
  error OnlyCallableByOwnerOrFeeAdmin();
  error OnlyCallableByOwnerOrFeeAdminOrNop();
  error InvalidWithdrawalAddress(address addr);
  error InvalidFeeToken(address token);
  error NoFeesToPay();
  error NoNopsToPay();
  error InsufficientBalance();
  error TooManyNops();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error MessageGasLimitTooHigh();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error PoolDoesNotExist(address token);
  error TokenPoolMismatch();
  error SenderNotAllowed(address sender);
  error InvalidConfig();
  error InvalidAddress(bytes encodedAddress);
  error BadAFNSignal();
  error LinkBalanceNotSettled();

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event AllowListEnabledSet(bool enabled);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event NopPaid(address indexed nop, uint256 amount);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event TokenTransferFeeConfigSet(TokenTransferFeeConfigArgs[] transferFeeConfig);
  event CCIPSendRequested(Internal.EVM2EVMMessage message);
  event NopsSet(uint256 nopWeightsTotal, NopAndWeight[] nopsAndWeights);
  event PoolAdded(address token, address pool);
  event PoolRemoved(address token, address pool);

  /// @dev Struct that contains the static configuration
  struct StaticConfig {
    address linkToken; // --------┐ Link token address
    uint64 chainSelector; // -----┘ Source chainSelector
    uint64 destChainSelector; // -┐ Destination chainSelector
    uint64 defaultTxGasLimit; // -┘ Default gas limit for a tx
    address prevOnRamp; //          Address of previous-version OnRamp
  }

  /// @dev Struct to contains the dynamic configuration
  struct DynamicConfig {
    address router; //            Router address
    address priceRegistry; // --┐ Price registry address
    uint32 maxDataSize; //      | Maximum payload data size
    uint64 maxGasLimit; // -----┘ Maximum gas limit for messages targeting EVMs
    uint16 maxTokensLength; // -┐ Maximum number of distinct ERC20 tokens that can be sent per message
    address afn; // ------------┘ AFN address
  }

  /// @dev Struct to hold the execution fee configuration for a fee token
  struct FeeTokenConfig {
    uint96 networkFeeAmountUSD; // -┐ Flat network fee in USD
    uint64 multiplier; //           | Price multiplier for gas costs
    uint32 destGasOverhead; // -----┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Struct to hold the fee configuration for a fee token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeTokenConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // --------------┐ Token address
    uint64 multiplier; // ----------┘ Price multiplier for gas costs
    uint96 networkFeeAmountUSD; // -┐ Flat network fee in USD
    uint32 destGasOverhead; // -----┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Struct to hold the transfer fee configuration for token transfers
  struct TokenTransferFeeConfig {
    uint32 minFee; // ---┐ Minimum USD fee to charge, multiples of 1 US cent, or 0.01USD
    uint32 maxFee; //    | Maximum USD fee to charge, multiples of 1 US cent, or 0.01USD
    uint16 ratio; // ----┘ Ratio of token transfer value to charge as fee, multiples of 0.1bps, or 10e-5
  }

  /// @dev Same as TokenTransferFeeConfig
  /// token included so that an array of these can be passed in to setTokenTransferFeeConfig
  struct TokenTransferFeeConfigArgs {
    address token; // ---┐ Token address
    uint32 minFee; //    | Minimum USD fee to charge, multiples of 1 US cent, or 0.01USD
    uint32 maxFee; //    | Maximum USD fee to charge, multiples of 1 US cent, or 0.01USD
    uint16 ratio; // ----┘ Ratio of token transfer value to charge as fee, multiples of 0.1bps, or 10e-5
  }

  /// @dev Nop address and weight, used to set the nops and their weights
  struct NopAndWeight {
    address nop; // -----┐ Address of the node operator
    uint16 weight; // ---┘ Weight for nop rewards
  }

  struct TokenAndPool {
    address token;
    address pool;
  }

  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOnRamp 1.0.0";
  /// @dev The metadata hash for this contract
  bytes32 internal immutable i_metadataHash;
  /// @dev Default gas limit for a transactions that did not specify
  /// a gas limit in the extraArgs.
  uint64 internal immutable i_defaultTxGasLimit;
  /// @dev The link token address - known to pay nops for their work
  address internal immutable i_linkToken;
  /// @dev The chain ID of the source chain that this contract is deployed to
  uint64 internal immutable i_chainSelector;
  /// @dev The chain ID of the destination chain
  uint64 internal immutable i_destChainSelector;
  /// @dev The address of previous-version OnRamp for this lane
  address internal immutable i_prevOnRamp;
  /// @dev the maximum number of nops that can be configured at the same time.
  uint256 private constant MAX_NUMBER_OF_NOPS = 64;

  // DYNAMIC CONFIG
  /// @dev The config for the onRamp
  DynamicConfig internal s_dynamicConfig;
  /// @dev (address nop => uint256 weight)
  EnumerableMap.AddressToUintMap internal s_nops;
  /// @dev source token => token pool
  EnumerableMapAddresses.AddressToAddressMap private s_poolsBySourceToken;
  /// @dev this allowListing will be removed before public launch
  /// @dev Whether s_allowList is enabled or not.
  bool private s_allowlistEnabled;
  /// @dev A set of addresses which can make ccipSend calls.
  EnumerableSet.AddressSet private s_allowList;
  /// @dev The execution fee token config that can be set by the owner or fee admin
  mapping(address => FeeTokenConfig) internal s_feeTokenConfig;
  /// @dev The token transfer fee config that can be set by the owner or fee admin
  mapping(address => TokenTransferFeeConfig) internal s_tokenTransferFeeConfig;

  // STATE
  /// @dev The current nonce per sender
  mapping(address => uint64) internal s_senderNonce;
  /// @dev The amount of LINK available to pay NOPS
  uint96 internal s_nopFeesJuels;
  /// @dev The combined weight of all NOPs weights
  uint32 internal s_nopWeightsTotal;
  /// @dev The last used sequence number. This is zero in the case where no
  /// messages has been sent yet. 0 is not a valid sequence number for any
  /// real transaction.
  uint64 internal s_sequenceNumber;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    TokenAndPool[] memory tokensAndPools,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig,
    FeeTokenConfigArgs[] memory feeTokenConfigs,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    NopAndWeight[] memory nopsAndWeights
  ) Pausable() AggregateRateLimiter(rateLimiterConfig) {
    if (
      staticConfig.linkToken == address(0) ||
      staticConfig.chainSelector == 0 ||
      staticConfig.destChainSelector == 0 ||
      staticConfig.defaultTxGasLimit == 0
    ) revert InvalidConfig();

    i_metadataHash = keccak256(
      abi.encode(
        Internal.EVM_2_EVM_MESSAGE_HASH,
        staticConfig.chainSelector,
        staticConfig.destChainSelector,
        address(this)
      )
    );
    i_linkToken = staticConfig.linkToken;
    i_chainSelector = staticConfig.chainSelector;
    i_destChainSelector = staticConfig.destChainSelector;
    i_defaultTxGasLimit = staticConfig.defaultTxGasLimit;
    i_prevOnRamp = staticConfig.prevOnRamp;

    _setDynamicConfig(dynamicConfig);
    _setFeeTokenConfig(feeTokenConfigs);
    _setTokenTransferFeeConfig(tokenTransferFeeConfigArgs);
    _setNops(nopsAndWeights);

    // Set new tokens and pools
    for (uint256 i = 0; i < tokensAndPools.length; ++i) {
      if (tokensAndPools[i].token == address(0) || address(tokensAndPools[i].pool) == address(0))
        revert InvalidConfig();
      s_poolsBySourceToken.set(tokensAndPools[i].token, tokensAndPools[i].pool);
      emit PoolAdded(tokensAndPools[i].token, tokensAndPools[i].pool);
    }

    if (allowlist.length > 0) {
      s_allowlistEnabled = true;
      _applyAllowListUpdates(allowlist, new address[](0));
    }
  }

  // ================================================================
  // |                          Messaging                           |
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getExpectedNextSequenceNumber() external view returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getSenderNonce(address sender) external view returns (uint64) {
    return s_senderNonce[sender];
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function forwardFromRouter(
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external whenNotPaused whenHealthy returns (bytes32) {
    Client.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    // Validate the message with various checks
    _validateMessage(message.data.length, extraArgs.gasLimit, message.tokenAmounts, originalSender);
    // Rate limit on aggregated token value
    _rateLimitValue(message.tokenAmounts, IPriceRegistry(s_dynamicConfig.priceRegistry));

    if (message.receiver.length != 32) revert InvalidAddress(message.receiver);
    uint256 decodedReceiver = abi.decode(message.receiver, (uint256));
    if (decodedReceiver > type(uint160).max) revert InvalidAddress(message.receiver);

    // Convert feeToken to link if not already in link
    if (message.feeToken == i_linkToken) {
      // Since there is only 1b link this is safe
      s_nopFeesJuels += uint96(feeTokenAmount);
    } else {
      // the cast from uint256 to uint96 is considered safe, uint96 can store more than max supply of link token
      s_nopFeesJuels += uint96(
        IPriceRegistry(s_dynamicConfig.priceRegistry).convertTokenAmount(message.feeToken, feeTokenAmount, i_linkToken)
      );
    }

    if (s_senderNonce[originalSender] == 0 && i_prevOnRamp != address(0)) {
      // If this is first time send for a sender in new OnRamp, check if they have a nonce
      // from the previous OnRamp and start from there instead of zero.
      s_senderNonce[originalSender] = IEVM2AnyOnRamp(i_prevOnRamp).getSenderNonce(originalSender);
    }

    // We need the next available sequence number so we increment before we use the value
    Internal.EVM2EVMMessage memory newMessage = Internal.EVM2EVMMessage({
      sourceChainSelector: i_chainSelector,
      sequenceNumber: ++s_sequenceNumber,
      feeTokenAmount: feeTokenAmount,
      sender: originalSender,
      nonce: ++s_senderNonce[originalSender],
      gasLimit: extraArgs.gasLimit,
      strict: extraArgs.strict,
      receiver: address(uint160(decodedReceiver)),
      data: message.data,
      tokenAmounts: message.tokenAmounts,
      feeToken: message.feeToken,
      messageId: ""
    });
    newMessage.messageId = Internal._hash(newMessage, i_metadataHash);

    // Lock the tokens as last step. TokenPools may not always be trusted.
    // There should be no state changes after external call to TokenPools.
    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      Client.EVMTokenAmount memory tokenAndAmount = message.tokenAmounts[i];
      getPoolBySourceToken(IERC20(tokenAndAmount.token)).lockOrBurn(
        originalSender,
        message.receiver,
        tokenAndAmount.amount,
        i_destChainSelector,
        bytes("") // any future extraArgs component would be added here
      );
    }

    // Emit message request
    emit CCIPSendRequested(newMessage);
    return newMessage.messageId;
  }

  /// @dev Convert the extra args bytes into a struct
  /// @param extraArgs The extra args bytes
  /// @return The extra args struct
  function _fromBytes(bytes calldata extraArgs) internal view returns (Client.EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return Client.EVMExtraArgsV1({gasLimit: i_defaultTxGasLimit, strict: false});
    }
    if (bytes4(extraArgs[:4]) != Client.EVM_EXTRA_ARGS_V1_TAG) revert InvalidExtraArgsTag();
    return abi.decode(extraArgs[4:], (Client.EVMExtraArgsV1));
  }

  /// @notice Validate the forwarded message with various checks.
  /// @param dataLength The length of the data field of the message
  /// @param gasLimit The gasLimit set in message for destination execution
  /// @param tokenAmounts The token payload to be sent. They will be locked into pools by this function.
  /// @param originalSender The original sender of the message on the router.
  function _validateMessage(
    uint256 dataLength,
    uint256 gasLimit,
    Client.EVMTokenAmount[] memory tokenAmounts,
    address originalSender
  ) internal view {
    if (msg.sender != s_dynamicConfig.router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_dynamicConfig.maxDataSize))
      revert MessageTooLarge(uint256(s_dynamicConfig.maxDataSize), dataLength);
    if (gasLimit > uint256(s_dynamicConfig.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (tokenAmounts.length > uint256(s_dynamicConfig.maxTokensLength)) revert UnsupportedNumberOfTokens();
    if (s_allowlistEnabled && !s_allowList.contains(originalSender)) revert SenderNotAllowed(originalSender);
  }

  // ================================================================
  // |                           Config                             |
  // ================================================================

  /// @notice Returns the static onRamp config.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return
      StaticConfig({
        linkToken: i_linkToken,
        chainSelector: i_chainSelector,
        destChainSelector: i_destChainSelector,
        defaultTxGasLimit: i_defaultTxGasLimit,
        prevOnRamp: i_prevOnRamp
      });
  }

  /// @notice Returns the dynamic onRamp config.
  /// @return dynamicConfig the configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory dynamicConfig) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The configuration.
  function setDynamicConfig(DynamicConfig memory dynamicConfig) external onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  function _setDynamicConfig(DynamicConfig memory dynamicConfig) internal {
    if (
      dynamicConfig.router == address(0) || dynamicConfig.priceRegistry == address(0) || dynamicConfig.afn == address(0)
    ) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(
      StaticConfig({
        linkToken: i_linkToken,
        chainSelector: i_chainSelector,
        destChainSelector: i_destChainSelector,
        defaultTxGasLimit: i_defaultTxGasLimit,
        prevOnRamp: i_prevOnRamp
      }),
      dynamicConfig
    );
  }

  // ================================================================
  // |                      Tokens and pools                        |
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getSupportedTokens() public view returns (address[] memory) {
    address[] memory sourceTokens = new address[](s_poolsBySourceToken.length());
    for (uint256 i = 0; i < sourceTokens.length; ++i) {
      (sourceTokens[i], ) = s_poolsBySourceToken.at(i);
    }
    return sourceTokens;
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getPoolBySourceToken(IERC20 sourceToken) public view returns (IPool) {
    if (!s_poolsBySourceToken.contains(address(sourceToken))) revert UnsupportedToken(sourceToken);
    return IPool(s_poolsBySourceToken.get(address(sourceToken)));
  }

  /// @inheritdoc IEVM2AnyOnRamp
  /// @dev This method can only be called by the owner of the contract.
  function applyPoolUpdates(Internal.PoolUpdate[] memory removes, Internal.PoolUpdate[] memory adds) public onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      address token = removes[i].token;
      address pool = removes[i].pool;

      if (!s_poolsBySourceToken.contains(token)) revert PoolDoesNotExist(token);
      if (s_poolsBySourceToken.get(token) != pool) revert TokenPoolMismatch();

      if (s_poolsBySourceToken.remove(token)) {
        emit PoolRemoved(token, pool);
      }
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      address token = adds[i].token;
      address pool = adds[i].pool;

      if (token == address(0) || pool == address(0)) revert InvalidTokenPoolConfig();
      if (s_poolsBySourceToken.contains(token)) revert PoolAlreadyAdded();

      if (s_poolsBySourceToken.set(token, pool)) {
        emit PoolAdded(token, pool);
      }
    }
  }

  // ================================================================
  // |                             Fees                             |
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getFee(Client.EVM2AnyMessage calldata message) public view returns (uint256) {
    // ensure MessageExecutionFee is evaluated strictly before TokenTransferFee
    // if there are validation reverts, errors from MessageExecutionFee take precedence
    uint256 executionFee = _getMessageExecutionFee(message.feeToken, message.extraArgs);
    return executionFee + _getTokenTransferFee(message.feeToken, message.tokenAmounts);
  }

  /// @notice Returns the fee based on the gas usage and network fee for the
  /// given message. Fees can be different for different fee tokens.
  function _getMessageExecutionFee(address feeToken, bytes calldata extraArgs) internal view returns (uint256 fee) {
    uint256 gasLimit = _fromBytes(extraArgs).gasLimit;
    (uint192 feeTokenPrice, uint192 gasPrice) = IPriceRegistry(s_dynamicConfig.priceRegistry).getFeeTokenAndGasPrices(
      feeToken,
      i_destChainSelector
    );
    FeeTokenConfig memory feeTokenConfig = s_feeTokenConfig[feeToken];

    uint256 usdFeeAmount = gasPrice *
      (((gasLimit + feeTokenConfig.destGasOverhead) * feeTokenConfig.multiplier) / 1 ether) +
      feeTokenConfig.networkFeeAmountUSD;
    return feeTokenPrice._calcTokenAmountFromUSDValue(usdFeeAmount);
  }

  /// @notice Returns the fee based on the tokens transferred. Will always be 0 if
  /// no tokens are transferred. The token fee is calculated based on basis points.
  function _getTokenTransferFee(address feeToken, Client.EVMTokenAmount[] calldata tokenAmounts)
    internal
    view
    returns (uint256 feeTokenAmount)
  {
    uint256 numerOfTokens = tokenAmounts.length;
    // short-circuit with 0 transfer fee if no token is being transferred
    if (numerOfTokens == 0) {
      return 0;
    }

    uint192 feeTokenPrice = IPriceRegistry(s_dynamicConfig.priceRegistry).getValidatedTokenPrice(feeToken);

    for (uint256 i = 0; i < numerOfTokens; ++i) {
      Client.EVMTokenAmount memory tokenAmount = tokenAmounts[i];
      TokenTransferFeeConfig memory transferFeeConfig = s_tokenTransferFeeConfig[tokenAmount.token];

      uint256 bpsValue;
      // ratio can be 0, only calculate bps fee if ratio is greater than 0
      if (transferFeeConfig.ratio > 0) {
        uint192 tokenPrice = feeTokenPrice;
        if (tokenAmount.token != feeToken) {
          tokenPrice = IPriceRegistry(s_dynamicConfig.priceRegistry).getValidatedTokenPrice(tokenAmount.token);
        }

        uint256 tokenValue = tokenPrice._calcUSDValueFromTokenAmount(tokenAmount.amount);

        // ratio represents multiples of 0.1bps, or 10e-5
        bpsValue = (tokenValue * transferFeeConfig.ratio) / 1e5;
      }

      // convert USD values with 2 decimals to 18 decimals
      uint256 minFeeValue = uint256(transferFeeConfig.minFee) * 1e16;
      uint256 maxFeeValue = uint256(transferFeeConfig.maxFee) * 1e16;

      if (bpsValue < minFeeValue) {
        feeTokenAmount += feeTokenPrice._calcTokenAmountFromUSDValue(minFeeValue);
      } else if (bpsValue > maxFeeValue) {
        feeTokenAmount += feeTokenPrice._calcTokenAmountFromUSDValue(maxFeeValue);
      } else {
        feeTokenAmount += feeTokenPrice._calcTokenAmountFromUSDValue(bpsValue);
      }
    }

    return feeTokenAmount;
  }

  /// @notice Gets the fee configuration for a token
  /// @param token The token to get the fee configuration for
  /// @return feeTokenConfig FeeTokenConfig struct
  function getFeeTokenConfig(address token) external view returns (FeeTokenConfig memory feeTokenConfig) {
    return s_feeTokenConfig[token];
  }

  /// @notice Sets the fee configuration for a token
  /// @param feeTokenConfigs Array of FeeTokenConfigArgs structs
  function setFeeTokenConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) external onlyOwnerOrAdmin {
    _setFeeTokenConfig(feeTokenConfigs);
  }

  /// @dev Set the fee config
  /// @param feeTokenConfigs The fee token configs
  function _setFeeTokenConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) internal {
    for (uint256 i = 0; i < feeTokenConfigs.length; ++i) {
      s_feeTokenConfig[feeTokenConfigs[i].token] = FeeTokenConfig({
        networkFeeAmountUSD: feeTokenConfigs[i].networkFeeAmountUSD,
        multiplier: feeTokenConfigs[i].multiplier,
        destGasOverhead: feeTokenConfigs[i].destGasOverhead
      });
    }
    emit FeeConfigSet(feeTokenConfigs);
  }

  /// @notice Gets the transfer fee config for a given token.
  function getTokenTransferFeeConfig(address token)
    external
    view
    returns (TokenTransferFeeConfig memory tokenTransferFeeConfig)
  {
    return s_tokenTransferFeeConfig[token];
  }

  /// @notice Sets the transfer fee config.
  /// @dev only callable by the owner or admin.
  function setTokenTransferFeeConfig(TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs)
    external
    onlyOwnerOrAdmin
  {
    _setTokenTransferFeeConfig(tokenTransferFeeConfigArgs);
  }

  /// @notice internal helper to set the token transfer fee config.
  function _setTokenTransferFeeConfig(TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs) internal {
    uint256 numerOfTokens = tokenTransferFeeConfigArgs.length;

    for (uint256 i = 0; i < numerOfTokens; ++i) {
      s_tokenTransferFeeConfig[tokenTransferFeeConfigArgs[i].token] = TokenTransferFeeConfig({
        minFee: tokenTransferFeeConfigArgs[i].minFee,
        maxFee: tokenTransferFeeConfigArgs[i].maxFee,
        ratio: tokenTransferFeeConfigArgs[i].ratio
      });
    }
    emit TokenTransferFeeConfigSet(tokenTransferFeeConfigArgs);
  }

  // ================================================================
  // |                         NOP payments                         |
  // ================================================================

  /// @notice Get the total amount of fees to be paid to the Nops (in LINK)
  /// @return totalNopFees
  function getNopFeesJuels() external view returns (uint96) {
    return s_nopFeesJuels;
  }

  /// @notice Gets the Nops and their weights
  /// @return nopsAndWeights Array of NopAndWeight structs
  /// @return weightsTotal The sum weight of all Nops
  function getNops() external view returns (NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) {
    uint256 length = s_nops.length();
    nopsAndWeights = new NopAndWeight[](length);
    for (uint256 i = 0; i < length; ++i) {
      (address nopAddress, uint256 nopWeight) = s_nops.at(i);
      nopsAndWeights[i] = NopAndWeight({nop: nopAddress, weight: uint16(nopWeight)});
    }
    weightsTotal = s_nopWeightsTotal;
    return (nopsAndWeights, weightsTotal);
  }

  /// @notice Sets the Nops and their weights
  /// @param nopsAndWeights Array of NopAndWeight structs
  function setNops(NopAndWeight[] calldata nopsAndWeights) external onlyOwner {
    _setNops(nopsAndWeights);
  }

  /// @dev Set the nops and weights
  /// @param nopsAndWeights The nops and weights
  function _setNops(NopAndWeight[] memory nopsAndWeights) internal {
    // Make sure all nops have been paid before removing nops
    // We only have to pay when there are nops and there is enough
    // outstanding NOP balance to trigger a payment.
    if (s_nopWeightsTotal > 0 && s_nopFeesJuels > s_nopWeightsTotal) {
      payNops();
    }

    uint256 numberOfNops = nopsAndWeights.length;
    if (numberOfNops > MAX_NUMBER_OF_NOPS) revert TooManyNops();

    // Remove previous
    delete s_nops;

    // Add new
    uint32 nopWeightsTotal = 0;
    // nopWeightsTotal is bounded by the MAX_NUMBER_OF_NOPS and the weight of
    // a single nop being of type uint16. This ensures nopWeightsTotal will
    // always fit into the uint32 type.
    for (uint256 i = 0; i < numberOfNops; ++i) {
      s_nops.set(nopsAndWeights[i].nop, nopsAndWeights[i].weight);
      nopWeightsTotal += nopsAndWeights[i].weight;
    }
    s_nopWeightsTotal = nopWeightsTotal;
    emit NopsSet(nopWeightsTotal, nopsAndWeights);
  }

  /// @notice Pays the Node Ops their outstanding balances.
  /// @dev some balance can remain after payments are done. This is at most the sum
  /// of the weight of all nops. Since nop weights are uint16s and we can have at
  /// most MAX_NUMBER_OF_NOPS NOPs, the highest possible value is 2**22 or 0.04 gjuels.
  function payNops() public onlyOwnerOrAdminOrNop {
    uint32 weightsTotal = s_nopWeightsTotal;
    if (weightsTotal == 0) revert NoNopsToPay();

    uint96 totalFeesToPay = s_nopFeesJuels;
    if (totalFeesToPay < weightsTotal) revert NoFeesToPay();
    if (IERC20(i_linkToken).balanceOf(address(this)) < totalFeesToPay) revert InsufficientBalance();

    uint96 fundsLeft = totalFeesToPay;
    uint256 numberOfNops = s_nops.length();
    for (uint256 i = 0; i < numberOfNops; ++i) {
      (address nop, uint256 weight) = s_nops.at(i);
      // amount can never be higher than totalFeesToPay so the cast to uint96 is safe
      uint96 amount = uint96((totalFeesToPay * weight) / weightsTotal);
      fundsLeft -= amount;
      IERC20(i_linkToken).safeTransfer(nop, amount);
      emit NopPaid(nop, amount);
    }
    // Some funds can remain, since this is an incredibly small
    // amount we consider this OK.
    s_nopFeesJuels = fundsLeft;
  }

  /// @notice Allows the owner to withdraw any ERC20 token that is not the fee token
  /// @param feeToken The token to withdraw
  /// @param to The address to send the tokens to
  function withdrawNonLinkFees(address feeToken, address to) external onlyOwner {
    if (feeToken == i_linkToken) revert InvalidFeeToken(feeToken);
    if (to == address(0)) revert InvalidWithdrawalAddress(to);
    uint256 linkBalance = IERC20(i_linkToken).balanceOf(address(this));

    // We require the link balance to be settled before allowing withdrawal
    // of non-link fees.
    if (linkBalance < s_nopFeesJuels) revert LinkBalanceNotSettled();

    IERC20(feeToken).safeTransfer(to, IERC20(feeToken).balanceOf(address(this)));
  }

  // ================================================================
  // |                          Allowlist                           |
  // ================================================================

  /// @notice Gets whether the allowList functionality is enabled.
  /// @return true is enabled, false if not.
  function getAllowListEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  /// @notice Enables or disabled the allowList functionality.
  /// @param enabled Signals whether the allowlist should be enabled.
  function setAllowListEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowListEnabledSet(enabled);
  }

  /// @notice Gets the allowed addresses.
  /// @return The allowed addresses.
  /// @dev May not work if allow list gets too large. Use events in that case to compute the set.
  function getAllowList() external view returns (address[] memory) {
    address[] memory allowList = new address[](s_allowList.length());
    for (uint256 i = 0; i < s_allowList.length(); ++i) {
      allowList[i] = s_allowList.at(i);
    }
    return allowList;
  }

  /// @notice Apply updates to the allow list.
  /// @param removes The addresses to be removed.
  /// @param adds The addresses to be added.
  /// @dev allowListing will be removed before public launch
  function applyAllowListUpdates(address[] calldata removes, address[] calldata adds) external onlyOwner {
    _applyAllowListUpdates(removes, adds);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  /// @dev allowListing will be removed before public launch
  function _applyAllowListUpdates(address[] memory removes, address[] memory adds) internal {
    for (uint256 i = 0; i < removes.length; ++i) {
      address toRemove = removes[i];
      if (s_allowList.remove(toRemove)) {
        emit AllowListRemove(toRemove);
      }
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      address toAdd = adds[i];
      if (toAdd == address(0)) {
        continue;
      }
      if (s_allowList.add(toAdd)) {
        emit AllowListAdd(toAdd);
      }
    }
  }

  // ================================================================
  // |                        Access and AFN                        |
  // ================================================================

  /// @dev Require that the sender is the owner or the fee admin or a nop
  modifier onlyOwnerOrAdminOrNop() {
    if (msg.sender != owner() && msg.sender != s_admin && !s_nops.contains(msg.sender))
      revert OnlyCallableByOwnerOrFeeAdminOrNop();
    _;
  }

  /// @dev Require that the sender is the owner or the fee admin
  modifier onlyOwnerOrAdmin() {
    if (msg.sender != owner() && msg.sender != s_admin) revert OnlyCallableByOwnerOrFeeAdmin();
    _;
  }

  /// @notice Support querying whether health checker is healthy.
  function isAFNHealthy() external view returns (bool) {
    return !IAFN(s_dynamicConfig.afn).isCursed();
  }

  /// @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IAFN(s_dynamicConfig.afn).isCursed()) revert BadAFNSignal();
    _;
  }

  /// @notice Pause the contract
  /// @dev only callable by the owner
  function pause() external onlyOwner {
    _pause();
  }

  /// @notice Unpause the contract
  /// @dev only callable by the owner
  function unpause() external onlyOwner {
    _unpause();
  }
}
