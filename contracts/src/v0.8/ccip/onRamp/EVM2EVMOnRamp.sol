// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IARM} from "../interfaces/IARM.sol";
import {IPriceRegistry} from "../interfaces/IPriceRegistry.sol";
import {IEVM2AnyOnRamp} from "../interfaces/IEVM2AnyOnRamp.sol";
import {ILinkAvailable} from "../interfaces/automation/ILinkAvailable.sol";

import {AggregateRateLimiter} from "../AggregateRateLimiter.sol";
import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";
import {EnumerableMapAddresses} from "../../shared/enumerable/EnumerableMapAddresses.sol";

import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/utils/structs/EnumerableMap.sol";

/// @notice The onRamp is a contract that handles lane-specific fee logic, NOP payments and
/// bridgeable token support.
/// @dev The EVM2EVMOnRamp, CommitStore and EVM2EVMOffRamp form an xchain upgradeable unit. Any change to one of them
/// results an onchain upgrade of all 3.
contract EVM2EVMOnRamp is IEVM2AnyOnRamp, ILinkAvailable, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.AddressToUintMap;
  using EnumerableMapAddresses for EnumerableMapAddresses.AddressToAddressMap;
  using USDPriceWith18Decimals for uint224;

  error InvalidExtraArgsTag();
  error OnlyCallableByOwnerOrAdmin();
  error OnlyCallableByOwnerOrAdminOrNop();
  error InvalidWithdrawParams();
  error NoFeesToPay();
  error NoNopsToPay();
  error InsufficientBalance();
  error TooManyNops();
  error MaxFeeBalanceReached();
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
  error InvalidConfig();
  error InvalidAddress(bytes encodedAddress);
  error BadARMSignal();
  error LinkBalanceNotSettled();
  error InvalidNopAddress(address nop);
  error NotAFeeToken(address token);
  error CannotSendZeroTokens();

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
    address linkToken; // ────────╮ Link token address
    uint64 chainSelector; // ─────╯ Source chainSelector
    uint64 destChainSelector; // ─╮ Destination chainSelector
    uint64 defaultTxGasLimit; //  │ Default gas limit for a tx
    uint96 maxNopFeesJuels; // ───╯ Max nop fee balance onramp can have
    address prevOnRamp; //          Address of previous-version OnRamp
    address armProxy; //            Address of ARM proxy
  }

  /// @dev Struct to contains the dynamic configuration
  struct DynamicConfig {
    address router; // ─────────────────────────╮ Router address
    uint16 maxTokensLength; //                  │ Maximum number of ERC20 token transfers per message
    uint32 destGasOverhead; //                  │ Extra gas charged on top of the gasLimit
    uint16 destGasPerPayloadByte; //            │ Destination chain gas charged per byte of `data` payload
    uint32 destDataAvailabilityOverheadGas; //  │ Extra data availability gas charged on top of message data
    uint16 destGasPerDataAvailabilityByte; // ──╯ Amount of gas to charge per byte of data that needs availability
    uint16 destDataAvailabilityMultiplier; // ──╮ Multiplier for data availability gas, multiples of 1e-4, or 0.0001
    address priceRegistry; //                   │ Price registry address
    uint32 maxDataSize; //                      │ Maximum payload data size, max 4GB
    uint32 maxGasLimit; // ─────────────────────╯ Maximum gas limit for messages targeting EVMs, max 4 Billion gas
  }

  /// @dev Struct to hold the execution fee configuration for a fee token
  struct FeeTokenConfig {
    uint32 networkFeeUSD; // ───────────╮ Flat network fee to charge for messages,  multiples of 0.01 USD
    uint32 minTokenTransferFeeUSD; //   │ Minimum fee to charge for token transfers, multiples of 0.01 USD
    uint32 maxTokenTransferFeeUSD; //   │ Maximum fee to charge for token transfers, multiples of 0.01 USD
    uint64 gasMultiplier; //            │ Price multiplier for gas costs, 1e18 based so 11e17 = 10% extra cost.
    uint64 premiumMultiplier; //        │ Multiplier for fee-token-specific premiums
    bool enabled; // ───────────────────╯ Whether this fee token is enabled
  }

  /// @dev Struct to hold the fee configuration for a fee token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeTokenConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // ──────────────────╮ Token address
    uint32 networkFeeUSD; //            │ Flat network fee to charge for messages,  multiples of 0.01 USD
    uint32 minTokenTransferFeeUSD; //   │ Minimum fee to charge for token transfers, multiples of 0.01 USD
    uint32 maxTokenTransferFeeUSD; //───╯ Maximum fee to charge for token transfers, multiples of 0.01 USD
    uint64 gasMultiplier; //         ───╮ Price multiplier for gas costs, 1e18 based so 11e17 = 10% extra cost
    uint64 premiumMultiplier; //        │ Multiplier for fee-token-specific premiums
    bool enabled; // ───────────────────╯ Whether this fee token is enabled
  }

  /// @dev Struct to hold the transfer fee configuration for token transfers
  struct TokenTransferFeeConfig {
    uint16 ratio; // ───────────────────╮ Ratio of token transfer value to charge as fee, multiples of 0.1bps, or 1e-5
    uint32 destGasOverhead; //          │ Gas charged to execute the token transfer on the destination chain
    uint32 destBytesOverhead; // ───────╯ Extra data availability bytes on top of transfer data, e.g. USDC offchain data
  }

  /// @dev Same as TokenTransferFeeConfig
  /// token included so that an array of these can be passed in to setTokenTransferFeeConfig
  struct TokenTransferFeeConfigArgs {
    address token; // ──────────────────╮ Token address
    uint16 ratio; //                    │ Ratio of token transfer value to charge as fee, multiples of 0.1bps, or 1e-5
    uint32 destGasOverhead; //          │ Gas charged to execute the token transfer on the destination chain
    uint32 destBytesOverhead; // ───────╯ Extra data availability bytes on top of transfer data, e.g. USDC offchain data
  }

  /// @dev Nop address and weight, used to set the nops and their weights
  struct NopAndWeight {
    address nop; // ────╮ Address of the node operator
    uint16 weight; // ──╯ Weight for nop rewards
  }

  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOnRamp 1.2.0";
  /// @dev metadataHash is a lane-specific prefix for a message hash preimage which ensures global uniqueness
  /// Ensures that 2 identical messages sent to 2 different lanes will have a distinct hash.
  /// Must match the metadataHash used in computing leaf hashes offchain for the root committed in
  /// the commitStore and i_metadataHash in the offRamp.
  bytes32 internal immutable i_metadataHash;
  /// @dev Default gas limit for a transactions that did not specify
  /// a gas limit in the extraArgs.
  uint64 internal immutable i_defaultTxGasLimit;
  /// @dev Maximum nop fee that can accumulate in this onramp
  uint96 internal immutable i_maxNopFeesJuels;
  /// @dev The link token address - known to pay nops for their work
  address internal immutable i_linkToken;
  /// @dev The chain ID of the source chain that this contract is deployed to
  uint64 internal immutable i_chainSelector;
  /// @dev The chain ID of the destination chain
  uint64 internal immutable i_destChainSelector;
  /// @dev The address of previous-version OnRamp for this lane
  /// Used to be able to provide sequencing continuity during a zero downtime upgrade.
  address internal immutable i_prevOnRamp;
  /// @dev The address of the arm proxy
  address internal immutable i_armProxy;
  /// @dev the maximum number of nops that can be configured at the same time.
  /// Used to bound gas for loops over nops.
  uint256 private constant MAX_NUMBER_OF_NOPS = 64;

  // DYNAMIC CONFIG
  /// @dev The config for the onRamp
  DynamicConfig internal s_dynamicConfig;
  /// @dev (address nop => uint256 weight)
  EnumerableMap.AddressToUintMap internal s_nops;
  /// @dev source token => token pool
  EnumerableMapAddresses.AddressToAddressMap private s_poolsBySourceToken;

  /// @dev The execution fee token config that can be set by the owner or fee admin
  mapping(address token => FeeTokenConfig feeTokenConfig) internal s_feeTokenConfig;
  /// @dev The token transfer fee config that can be set by the owner or fee admin
  mapping(address token => TokenTransferFeeConfig tranferFeeConfig) internal s_tokenTransferFeeConfig;

  // STATE
  /// @dev The current nonce per sender.
  /// The offramp has a corresponding s_senderNonce mapping to ensure messages
  /// are executed in the same order they are sent.
  mapping(address sender => uint64 nonce) internal s_senderNonce;
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
    Internal.PoolUpdate[] memory tokensAndPools,
    RateLimiter.Config memory rateLimiterConfig,
    FeeTokenConfigArgs[] memory feeTokenConfigs,
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs,
    NopAndWeight[] memory nopsAndWeights
  ) AggregateRateLimiter(rateLimiterConfig) {
    if (
      staticConfig.linkToken == address(0) ||
      staticConfig.chainSelector == 0 ||
      staticConfig.destChainSelector == 0 ||
      staticConfig.defaultTxGasLimit == 0 ||
      staticConfig.armProxy == address(0)
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
    i_maxNopFeesJuels = staticConfig.maxNopFeesJuels;
    i_prevOnRamp = staticConfig.prevOnRamp;
    i_armProxy = staticConfig.armProxy;

    _setDynamicConfig(dynamicConfig);
    _setFeeTokenConfig(feeTokenConfigs);
    _setTokenTransferFeeConfig(tokenTransferFeeConfigArgs);
    _setNops(nopsAndWeights);

    // Set new tokens and pools
    _applyPoolUpdates(new Internal.PoolUpdate[](0), tokensAndPools);
  }

  // ================================================================
  // │                          Messaging                           │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getExpectedNextSequenceNumber() external view returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getSenderNonce(address sender) external view returns (uint64) {
    uint256 senderNonce = s_senderNonce[sender];

    if (senderNonce == 0 && i_prevOnRamp != address(0)) {
      // If OnRamp was upgraded, check if sender has a nonce from the previous OnRamp.
      return IEVM2AnyOnRamp(i_prevOnRamp).getSenderNonce(sender);
    }
    return uint64(senderNonce);
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function forwardFromRouter(
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external whenHealthy returns (bytes32) {
    // Validate message sender is set and allowed. Not validated in `getFee` since it is not user-driven.
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Router address may be zero intentionally to pause.
    if (msg.sender != s_dynamicConfig.router) revert MustBeCalledByRouter();

    // EVM destination addresses should be abi encoded and therefore always 32 bytes long
    // Not duplicately validated in `getFee`. Invalid address is uncommon, gas cost outweighs UX gain.
    if (message.receiver.length != 32) revert InvalidAddress(message.receiver);
    uint256 decodedReceiver = abi.decode(message.receiver, (uint256));
    // We want to disallow sending to address(0) and to precompiles, which exist on address(1) through address(9).
    if (decodedReceiver > type(uint160).max || decodedReceiver < 10) revert InvalidAddress(message.receiver);

    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    // Validate the message with various checks

    uint256 numberOfTokens = message.tokenAmounts.length;
    _validateMessage(message.data.length, gasLimit, numberOfTokens);

    // Only check token value if there are tokens
    if (numberOfTokens > 0) {
      for (uint256 i = 0; i < numberOfTokens; ++i) {
        if (message.tokenAmounts[i].amount == 0) revert CannotSendZeroTokens();
      }
      // Rate limit on aggregated token value
      _rateLimitValue(message.tokenAmounts, IPriceRegistry(s_dynamicConfig.priceRegistry));
    }

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
    if (s_nopFeesJuels > i_maxNopFeesJuels) revert MaxFeeBalanceReached();

    if (s_senderNonce[originalSender] == 0 && i_prevOnRamp != address(0)) {
      // If this is first time send for a sender in new OnRamp, check if they have a nonce
      // from the previous OnRamp and start from there instead of zero.
      s_senderNonce[originalSender] = IEVM2AnyOnRamp(i_prevOnRamp).getSenderNonce(originalSender);
    }

    // We need the next available sequence number so we increment before we use the value
    Internal.EVM2EVMMessage memory newMessage = Internal.EVM2EVMMessage({
      sourceChainSelector: i_chainSelector,
      sender: originalSender,
      receiver: address(uint160(decodedReceiver)),
      sequenceNumber: ++s_sequenceNumber,
      gasLimit: gasLimit,
      strict: false,
      nonce: ++s_senderNonce[originalSender],
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      data: message.data,
      tokenAmounts: message.tokenAmounts,
      sourceTokenData: new bytes[](numberOfTokens), // will be filled in later
      messageId: ""
    });

    // Lock the tokens as last step. TokenPools may not always be trusted.
    // There should be no state changes after external call to TokenPools.
    for (uint256 i = 0; i < numberOfTokens; ++i) {
      Client.EVMTokenAmount memory tokenAndAmount = message.tokenAmounts[i];
      newMessage.sourceTokenData[i] = getPoolBySourceToken(IERC20(tokenAndAmount.token)).lockOrBurn(
        originalSender,
        message.receiver,
        tokenAndAmount.amount,
        i_destChainSelector,
        bytes("") // any future extraArgs component would be added here
      );
    }

    // Hash only after the sourceTokenData has been set
    newMessage.messageId = Internal._hash(newMessage, i_metadataHash);

    // Emit message request
    emit CCIPSendRequested(newMessage);
    return newMessage.messageId;
  }

  /// @dev Convert the extra args bytes into a struct
  /// @param extraArgs The extra args bytes
  /// @return The extra args struct
  function _fromBytes(bytes calldata extraArgs) internal view returns (Client.EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return Client.EVMExtraArgsV1({gasLimit: i_defaultTxGasLimit});
    }
    if (bytes4(extraArgs) != Client.EVM_EXTRA_ARGS_V1_TAG) revert InvalidExtraArgsTag();
    // EVMExtraArgsV1 originally included a second boolean (strict) field which we have deprecated entirely.
    // Clients may still send that version but it will be ignored.
    return abi.decode(extraArgs[4:], (Client.EVMExtraArgsV1));
  }

  /// @notice Validate the forwarded message with various checks.
  /// @dev This function can be called multiple times during a CCIPSend,
  /// only common user-driven mistakes are validated here to minimize duplicate validation cost.
  /// @param dataLength The length of the data field of the message.
  /// @param gasLimit The gasLimit set in message for destination execution.
  /// @param numberOfTokens The number of tokens to be sent.
  function _validateMessage(uint256 dataLength, uint256 gasLimit, uint256 numberOfTokens) internal view {
    // Check that payload is formed correctly
    uint256 maxDataSize = uint256(s_dynamicConfig.maxDataSize);
    if (dataLength > maxDataSize) revert MessageTooLarge(maxDataSize, dataLength);
    if (gasLimit > uint256(s_dynamicConfig.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (numberOfTokens > uint256(s_dynamicConfig.maxTokensLength)) revert UnsupportedNumberOfTokens();
  }

  // ================================================================
  // │                           Config                             │
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
        maxNopFeesJuels: i_maxNopFeesJuels,
        prevOnRamp: i_prevOnRamp,
        armProxy: i_armProxy
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
    // We permit router to be set to zero as a way to pause the contract.
    if (dynamicConfig.priceRegistry == address(0)) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(
      StaticConfig({
        linkToken: i_linkToken,
        chainSelector: i_chainSelector,
        destChainSelector: i_destChainSelector,
        defaultTxGasLimit: i_defaultTxGasLimit,
        maxNopFeesJuels: i_maxNopFeesJuels,
        prevOnRamp: i_prevOnRamp,
        armProxy: i_armProxy
      }),
      dynamicConfig
    );
  }

  // ================================================================
  // │                      Tokens and pools                        │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getSupportedTokens() external view returns (address[] memory) {
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
  function applyPoolUpdates(
    Internal.PoolUpdate[] memory removes,
    Internal.PoolUpdate[] memory adds
  ) external onlyOwner {
    _applyPoolUpdates(removes, adds);
  }

  function _applyPoolUpdates(Internal.PoolUpdate[] memory removes, Internal.PoolUpdate[] memory adds) internal {
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
      if (token != address(IPool(pool).getToken())) revert TokenPoolMismatch();

      if (s_poolsBySourceToken.set(token, pool)) {
        emit PoolAdded(token, pool);
      } else {
        revert PoolAlreadyAdded();
      }
    }
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  /// @dev getFee MUST revert if the feeToken is not listed in the fee token config.
  /// as the router assumes it does.
  function getFee(Client.EVM2AnyMessage calldata message) external view returns (uint256) {
    Client.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    // Validate the message with various checks
    _validateMessage(message.data.length, extraArgs.gasLimit, message.tokenAmounts.length);

    FeeTokenConfig memory feeTokenConfig = s_feeTokenConfig[message.feeToken];
    if (!feeTokenConfig.enabled) revert NotAFeeToken(message.feeToken);

    (uint224 feeTokenPrice, uint224 packedGasPrice) = IPriceRegistry(s_dynamicConfig.priceRegistry)
      .getTokenAndGasPrices(message.feeToken, i_destChainSelector);
    uint112 executionGasPrice = uint112(packedGasPrice);

    // Calculate premiumFee in USD with 18 decimals precision.
    // If there are token transfers, premiumFee is calculated from token transfer fees.
    // If there are no token transfers, we charge a flat network fee.
    uint256 premiumFeeUSD = 0;
    uint32 tokenTransferGas = 0;
    uint32 tokenTransferBytesOverhead = 0;
    if (message.tokenAmounts.length > 0) {
      (premiumFeeUSD, tokenTransferGas, tokenTransferBytesOverhead) = _getTokenTransferCost(
        message.feeToken,
        feeTokenPrice,
        message.tokenAmounts,
        feeTokenConfig
      );
    } else {
      // Convert USD values with 2 decimals to 18 decimals.
      premiumFeeUSD = uint256(feeTokenConfig.networkFeeUSD) * 1e16;
    }

    // Apply a feeToken-specific multiplier with 18 decimals, arrive at 36 decimals
    premiumFeeUSD = premiumFeeUSD * feeTokenConfig.premiumMultiplier;

    // Calculate execution gas fee on destination chain in USD with 36 decimals.
    // We add the message gas limit, the overhead gas, and the data availability gas together.
    // We then multiple this destination gas total with the gas multiplier and convert it into USD.
    uint256 executionCostUSD = executionGasPrice *
      ((extraArgs.gasLimit +
        s_dynamicConfig.destGasOverhead +
        (message.data.length * s_dynamicConfig.destGasPerPayloadByte) +
        tokenTransferGas) * feeTokenConfig.gasMultiplier);

    uint256 dataAvailabilityCostUSD = 0;
    // Only calculate data availability cost if multiplier is non-zero.
    // The multiplier should be set to 0 if destination chain does not charge data availability cost.
    if (s_dynamicConfig.destDataAvailabilityMultiplier > 0) {
      uint112 dataAvailabilityGasPrice = uint112(packedGasPrice >> Internal.GAS_PRICE_BITS);

      dataAvailabilityCostUSD = _getDataAvailabilityCostUSD(
        dataAvailabilityGasPrice,
        message.data.length,
        message.tokenAmounts.length,
        tokenTransferBytesOverhead
      );
    }

    // Calculate number of fee tokens to charge.
    // Total USD fee is in 36 decimals, feeTokenPrice is in 18 decimals USD for 1e18 smallest token denominations.
    // Result of the division is the number of smallest token denominations.
    return (premiumFeeUSD + executionCostUSD + dataAvailabilityCostUSD) / feeTokenPrice;
  }

  /// @notice Returns the estimated data availability cost of the message.
  /// @dev To save on gas, we use a single destGasPerDataAvailabilityByte value for both zero and non-zero bytes.
  /// @param dataAvailabilityGasPrice USD per data availability gas in 18 decimals.
  /// @param messageDataLength length of the data field in the message.
  /// @param numberOfTokens number of distinct token transfers in the message.
  /// @param tokenTransferBytesOverhead additional token transfer data passed to destination, e.g. USDC attestation.
  /// @return dataAvailabilityCostUSD total data availability cost in USD with 36 decimals.
  function _getDataAvailabilityCostUSD(
    uint112 dataAvailabilityGasPrice,
    uint256 messageDataLength,
    uint256 numberOfTokens,
    uint32 tokenTransferBytesOverhead
  ) internal view returns (uint256 dataAvailabilityCostUSD) {
    uint256 dataAvailabilityLengthBytes = Internal.MESSAGE_FIXED_BYTES +
      messageDataLength +
      (numberOfTokens * Internal.MESSAGE_BYTES_PER_TOKEN) +
      tokenTransferBytesOverhead;

    uint256 dataAvailabilityGas = (dataAvailabilityLengthBytes * s_dynamicConfig.destGasPerDataAvailabilityByte) +
      s_dynamicConfig.destDataAvailabilityOverheadGas;

    // dataAvailabilityGasPrice is in 18 decimals, destDataAvailabilityMultiplier is in 4 decimals
    // we pad 14 decimals to bring the result to 36 decimals, in line with token bps and execution fee.
    return ((dataAvailabilityGas * dataAvailabilityGasPrice) * s_dynamicConfig.destDataAvailabilityMultiplier) * 1e14;
  }

  /// @notice Returns the token transfer fee.
  /// A basis point fee is calculated from the USD value of each token transfer.
  /// Sum of basis point fees is confined within range [minTokenTransferFeeUSD, maxTokenTransferFeeUSD].
  /// @dev Assumes that tokenAmounts are validated to be listed tokens elsewhere.
  /// @dev Splitting one token transfer into multiple transfers is discouraged,
  /// as it will result in a transferFee equal or greater than the same amount aggregated/de-duped.
  /// @param feeToken address of the feeToken.
  /// @param feeTokenPrice price of feeToken in USD with 18 decimals.
  /// @param tokenAmounts token transfers in the message.
  /// @param feeTokenConfig configuration struct of fee token.
  /// @return tokenTransferFeeUSD total token transfer bps fee in USD with 36 decimals.
  /// @return tokenTransferGas total execution gas of the token transfers.
  /// @return tokenTransferBytesOverhead additional token transfer data passed to destination, e.g. USDC attestation.
  function _getTokenTransferCost(
    address feeToken,
    uint224 feeTokenPrice,
    Client.EVMTokenAmount[] calldata tokenAmounts,
    FeeTokenConfig memory feeTokenConfig
  ) internal view returns (uint256 tokenTransferFeeUSD, uint32 tokenTransferGas, uint32 tokenTransferBytesOverhead) {
    uint256 numberOfTokens = tokenAmounts.length;

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      Client.EVMTokenAmount memory tokenAmount = tokenAmounts[i];
      TokenTransferFeeConfig memory transferFeeConfig = s_tokenTransferFeeConfig[tokenAmount.token];

      // Validate if the token is supported, do not calculate fee for unsupported tokens.
      if (!s_poolsBySourceToken.contains(tokenAmount.token)) revert UnsupportedToken(IERC20(tokenAmount.token));

      uint256 bpsFeeUSD = 0;
      // Only calculate bps fee if ratio is greater than 0. Ratio of 0 means no bps fee for a token.
      // Useful for when the PriceRegistry cannot return a valid price for the token.
      if (transferFeeConfig.ratio > 0) {
        uint224 tokenPrice = 0;
        if (tokenAmount.token != feeToken) {
          tokenPrice = IPriceRegistry(s_dynamicConfig.priceRegistry).getValidatedTokenPrice(tokenAmount.token);
        } else {
          tokenPrice = feeTokenPrice;
        }

        // Calculate token transfer value, then apply fee ratio
        // ratio represents multiples of 0.1bps, or 1e-5
        bpsFeeUSD = (tokenPrice._calcUSDValueFromTokenAmount(tokenAmount.amount) * transferFeeConfig.ratio) / 1e5;
      }

      tokenTransferFeeUSD += bpsFeeUSD;
      tokenTransferGas += transferFeeConfig.destGasOverhead;
      tokenTransferBytesOverhead += transferFeeConfig.destBytesOverhead;
    }

    // Convert USD values with 2 decimals to 18 decimals.
    // Sum of bps fees should be kept within range of [minTokenTransferFeeUSD, maxTokenTransferFeeUSD].
    uint256 minTransferFeeUSD = uint256(feeTokenConfig.minTokenTransferFeeUSD) * 1e16;
    if (tokenTransferFeeUSD < minTransferFeeUSD) {
      return (minTransferFeeUSD, tokenTransferGas, tokenTransferBytesOverhead);
    }

    uint256 maxTransferFeeUSD = uint256(feeTokenConfig.maxTokenTransferFeeUSD) * 1e16;
    if (tokenTransferFeeUSD > maxTransferFeeUSD) {
      return (maxTransferFeeUSD, tokenTransferGas, tokenTransferBytesOverhead);
    }

    return (tokenTransferFeeUSD, tokenTransferGas, tokenTransferBytesOverhead);
  }

  /// @notice Gets the fee configuration for a token
  /// @param token The token to get the fee configuration for
  /// @return feeTokenConfig FeeTokenConfig struct
  function getFeeTokenConfig(address token) external view returns (FeeTokenConfig memory feeTokenConfig) {
    return s_feeTokenConfig[token];
  }

  /// @notice Sets the fee configuration for a token
  /// @param feeTokenConfigArgs Array of FeeTokenConfigArgs structs.
  function setFeeTokenConfig(FeeTokenConfigArgs[] memory feeTokenConfigArgs) external onlyOwnerOrAdmin {
    _setFeeTokenConfig(feeTokenConfigArgs);
  }

  /// @dev Set the fee config
  /// @param feeTokenConfigArgs The fee token configs.
  function _setFeeTokenConfig(FeeTokenConfigArgs[] memory feeTokenConfigArgs) internal {
    for (uint256 i = 0; i < feeTokenConfigArgs.length; ++i) {
      FeeTokenConfigArgs memory configArg = feeTokenConfigArgs[i];

      s_feeTokenConfig[configArg.token] = FeeTokenConfig({
        networkFeeUSD: configArg.networkFeeUSD,
        minTokenTransferFeeUSD: configArg.minTokenTransferFeeUSD,
        maxTokenTransferFeeUSD: configArg.maxTokenTransferFeeUSD,
        gasMultiplier: configArg.gasMultiplier,
        premiumMultiplier: configArg.premiumMultiplier,
        enabled: configArg.enabled
      });
    }
    emit FeeConfigSet(feeTokenConfigArgs);
  }

  /// @notice Gets the transfer fee config for a given token.
  function getTokenTransferFeeConfig(
    address token
  ) external view returns (TokenTransferFeeConfig memory tokenTransferFeeConfig) {
    return s_tokenTransferFeeConfig[token];
  }

  /// @notice Sets the transfer fee config.
  /// @dev only callable by the owner or admin.
  function setTokenTransferFeeConfig(
    TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs
  ) external onlyOwnerOrAdmin {
    _setTokenTransferFeeConfig(tokenTransferFeeConfigArgs);
  }

  /// @notice internal helper to set the token transfer fee config.
  function _setTokenTransferFeeConfig(TokenTransferFeeConfigArgs[] memory tokenTransferFeeConfigArgs) internal {
    for (uint256 i = 0; i < tokenTransferFeeConfigArgs.length; ++i) {
      TokenTransferFeeConfigArgs memory configArg = tokenTransferFeeConfigArgs[i];

      s_tokenTransferFeeConfig[configArg.token] = TokenTransferFeeConfig({
        ratio: configArg.ratio,
        destGasOverhead: configArg.destGasOverhead,
        destBytesOverhead: configArg.destBytesOverhead
      });
    }
    emit TokenTransferFeeConfigSet(tokenTransferFeeConfigArgs);
  }

  // ================================================================
  // │                         NOP payments                         │
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
  function setNops(NopAndWeight[] calldata nopsAndWeights) external onlyOwnerOrAdmin {
    _setNops(nopsAndWeights);
  }

  /// @param nopsAndWeights New set of nops and weights
  /// @dev Clears existing nops, sets new nops and weights
  /// @dev We permit fees to accrue before nops are configured, in which case
  /// they will go to the first set of configured nops.
  function _setNops(NopAndWeight[] memory nopsAndWeights) internal {
    uint256 numberOfNops = nopsAndWeights.length;
    if (numberOfNops > MAX_NUMBER_OF_NOPS) revert TooManyNops();

    // Make sure all nops have been paid before removing nops
    // We only have to pay when there are nops and there is enough
    // outstanding NOP balance to trigger a payment.
    if (s_nopWeightsTotal > 0 && s_nopFeesJuels >= s_nopWeightsTotal) {
      payNops();
    }

    // Remove all previous nops, move from end to start to avoid shifting
    for (uint256 i = s_nops.length(); i > 0; --i) {
      (address nop, ) = s_nops.at(i - 1);
      s_nops.remove(nop);
    }

    // Add new
    uint32 nopWeightsTotal = 0;
    // nopWeightsTotal is bounded by the MAX_NUMBER_OF_NOPS and the weight of
    // a single nop being of type uint16. This ensures nopWeightsTotal will
    // always fit into the uint32 type.
    for (uint256 i = 0; i < numberOfNops; ++i) {
      // Make sure the LINK token is not a nop because the link token doesn't allow
      // self transfers. If set as nop, payNops would always revert. Since setNops
      // calls payNops, we can never remove the LINK token as a nop.
      address nop = nopsAndWeights[i].nop;
      uint16 weight = nopsAndWeights[i].weight;
      if (nop == i_linkToken || nop == address(0)) revert InvalidNopAddress(nop);
      s_nops.set(nop, weight);
      nopWeightsTotal += weight;
    }
    s_nopWeightsTotal = nopWeightsTotal;
    emit NopsSet(nopWeightsTotal, nopsAndWeights);
  }

  /// @notice Pays the Node Ops their outstanding balances.
  /// @dev some balance can remain after payments are done. This is at most the sum
  /// of the weight of all nops. Since nop weights are uint16s and we can have at
  /// most MAX_NUMBER_OF_NOPS NOPs, the highest possible value is 2**22 or 0.04 gjuels.
  function payNops() public onlyOwnerOrAdminOrNop {
    uint256 weightsTotal = s_nopWeightsTotal;
    if (weightsTotal == 0) revert NoNopsToPay();

    uint96 totalFeesToPay = s_nopFeesJuels;
    if (totalFeesToPay < weightsTotal) revert NoFeesToPay();
    if (_linkLeftAfterNopFees() < 0) revert InsufficientBalance();

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
  function withdrawNonLinkFees(address feeToken, address to) external onlyOwnerOrAdmin {
    if (feeToken == i_linkToken || to == address(0)) revert InvalidWithdrawParams();

    // We require the link balance to be settled before allowing withdrawal
    // of non-link fees.
    if (_linkLeftAfterNopFees() < 0) revert LinkBalanceNotSettled();

    IERC20(feeToken).safeTransfer(to, IERC20(feeToken).balanceOf(address(this)));
  }

  // ================================================================
  // │                        Link monitoring                       │
  // ================================================================

  /// @notice Calculate remaining LINK balance after paying nops
  /// @return balance if nops were to be paid
  function _linkLeftAfterNopFees() private view returns (int256) {
    // Since LINK caps at uint96, casting to int256 is safe
    return int256(IERC20(i_linkToken).balanceOf(address(this))) - int256(uint256(s_nopFeesJuels));
  }

  /// @notice Allow keeper to monitor funds available for paying nops
  function linkAvailableForPayment() external view returns (int256) {
    return _linkLeftAfterNopFees();
  }

  // ================================================================
  // │                        Access and ARM                        │
  // ================================================================

  /// @dev Require that the sender is the owner or the fee admin or a nop
  modifier onlyOwnerOrAdminOrNop() {
    if (msg.sender != owner() && msg.sender != s_admin && !s_nops.contains(msg.sender))
      revert OnlyCallableByOwnerOrAdminOrNop();
    _;
  }

  /// @dev Require that the sender is the owner or the fee admin
  modifier onlyOwnerOrAdmin() {
    if (msg.sender != owner() && msg.sender != s_admin) revert OnlyCallableByOwnerOrAdmin();
    _;
  }

  /// @notice Ensure that the ARM has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IARM(i_armProxy).isCursed()) revert BadARMSignal();
    _;
  }
}
