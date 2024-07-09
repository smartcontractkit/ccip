// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {IEVM2AnyOnRampClient} from "../interfaces/IEVM2AnyOnRampClient.sol";
import {IMessageInterceptor} from "../interfaces/IMessageInterceptor.sol";
import {INonceManager} from "../interfaces/INonceManager.sol";
import {IPoolV1} from "../interfaces/IPool.sol";
import {IPriceRegistry} from "../interfaces/IPriceRegistry.sol";
import {IRMN} from "../interfaces/IRMN.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";
import {Pool} from "../libraries/Pool.sol";
import {USDPriceWith18Decimals} from "../libraries/USDPriceWith18Decimals.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice The EVM2EVMMultiOnRamp is a contract that handles lane-specific fee logic
/// @dev The EVM2EVMMultiOnRamp, MultiCommitStore and EVM2EVMMultiOffRamp form an xchain upgradeable unit. Any change to one of them
/// results an onchain upgrade of all 3.
contract EVM2EVMMultiOnRamp is IEVM2AnyOnRampClient, ITypeAndVersion, OwnerIsCreator {
  using SafeERC20 for IERC20;
  using USDPriceWith18Decimals for uint224;

  error CannotSendZeroTokens();
  error InvalidExtraArgsTag();
  error ExtraArgOutOfOrderExecutionMustBeTrue();
  error OnlyCallableByOwnerOrAdmin();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error MessageGasLimitTooHigh();
  error MessageFeeTooHigh(uint256 msgFeeJuels, uint256 maxFeeJuelsPerMsg);
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(address token);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error InvalidConfig();
  error CursedByRMN(uint64 sourceChainSelector);
  error SourceTokenDataTooLarge(address token);
  error GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  error InvalidDestChainConfig(uint64 destChainSelector);
  error DestinationChainNotEnabled(uint64 destChainSelector);
  error InvalidDestBytesOverhead(address token, uint32 destBytesOverhead);

  event AdminSet(address newAdmin);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event FeePaid(address indexed feeToken, uint256 feeValueJuels);
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  /// RMN depends on this event, if changing, please notify the RMN maintainers.
  event CCIPSendRequested(uint64 indexed destChainSelector, Internal.EVM2AnyRampMessage message);
  event DestChainAdded(uint64 indexed destChainSelector, DestChainConfig destChainConfig);
  event DestChainDynamicConfigUpdated(uint64 indexed destChainSelector, DestChainDynamicConfig dynamicConfig);

  /// @dev Struct that contains the static configuration
  /// RMN depends on this struct, if changing, please notify the RMN maintainers.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    address linkToken; // ────────╮ Link token address
    uint64 chainSelector; // ─────╯ Source chainSelector
    uint96 maxFeeJuelsPerMsg; // ─╮ Maximum fee that can be charged for a message
    address rmnProxy; // ─────────╯ Address of RMN proxy
    address nonceManager; // Address of the nonce manager
    address tokenAdminRegistry; // Token admin registry address
  }

  /// @dev Struct to contains the dynamic configuration
  // solhint-disable-next-line gas-struct-packing
  struct DynamicConfig {
    address router; // Router address
    address priceRegistry; // Price registry address
    address messageValidator; // Optional message validator to validate outbound messages (zero address = no validator)
    address feeAggregator; // Fee aggregator address
  }

  /// @dev Struct to hold the dynamic configs for a destination chain
  struct DestChainDynamicConfig {
    bool isEnabled; // ──────────────────────────╮ Whether this destination chain is enabled
    uint16 maxNumberOfTokensPerMsg; //           │ Maximum number of distinct ERC20 token transferred per message
    uint32 maxDataBytes; //                      │ Maximum payload data size in bytes
    uint32 maxPerMsgGasLimit; //                 │ Maximum gas limit for messages targeting EVMs
    // TODO
    // The following three properties are defaults, they can be overridden by setting the TokenTransferFeeConfig for a token
    uint64 defaultTxGasLimit; //                 │ Default gas limit for a tx
    bool enforceOutOfOrder; //                   │ Whether to enforce the allowOutOfOrderExecution extraArg value to be true.
    bytes4 chainFamilySelector; // ──────────────╯ Selector that identifies the destination chain's family. Used to determine the correct validations to perform for the dest chain.
  }

  /// @dev Struct to hold the configs for a destination chain
  struct DestChainConfig {
    DestChainDynamicConfig dynamicConfig; // ──╮ Dynamic configs for a destination chain
    address prevOnRamp; // ────────────────────╯ Address of previous-version OnRamp
    uint64 sequenceNumber; // The last used sequence number. This is zero in the case where no messages has been sent yet.
    // 0 is not a valid sequence number for any real transaction.
    /// @dev metadataHash is a lane-specific prefix for a message hash preimage which ensures global uniqueness
    /// Ensures that 2 identical messages sent to 2 different lanes will have a distinct hash.
    /// Must match the metadataHash used in computing leaf hashes offchain for the root committed in
    /// the commitStore and i_metadataHash in the offRamp.
    bytes32 metadataHash;
  }

  /// @dev Struct to hold the dynamic configs, its destination chain selector and previous onRamp.
  /// Same as DestChainConfig but with the destChainSelector and the prevOnRamp so that an array of these
  /// can be passed in the constructor and the applyDestChainConfigUpdates function
  //solhint-disable gas-struct-packing
  struct DestChainConfigArgs {
    uint64 destChainSelector; // Destination chain selector
    DestChainDynamicConfig dynamicConfig; // Struct to hold the configs for a destination chain
    address prevOnRamp; // Address of previous-version OnRamp.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "EVM2EVMMultiOnRamp 1.6.0-dev";
  /// @dev Maximum fee that can be charged for a message. This is a guard to prevent massively overcharging due to misconfiguation.
  uint96 internal immutable i_maxFeeJuelsPerMsg;
  /// @dev The link token address
  address internal immutable i_linkToken;
  /// @dev The chain ID of the source chain that this contract is deployed to
  uint64 internal immutable i_chainSelector;
  /// @dev The address of the rmn proxy
  address internal immutable i_rmnProxy;
  /// @dev The address of the nonce manager
  address internal immutable i_nonceManager;
  /// @dev The address of the token admin registry
  address internal immutable i_tokenAdminRegistry;
  /// @dev the maximum number of nops that can be configured at the same time.
  /// Used to bound gas for loops over nops.
  uint256 private constant MAX_NUMBER_OF_NOPS = 64;

  // DYNAMIC CONFIG
  /// @dev The config for the onRamp
  DynamicConfig internal s_dynamicConfig;

  /// @dev The destination chain specific configs
  mapping(uint64 destChainSelector => DestChainConfig destChainConfig) internal s_destChainConfig;

  // STATE
  /// @dev The amount of LINK available to pay NOPS
  uint96 internal s_nopFeesJuels;
  /// @dev The combined weight of all NOPs weights
  uint32 internal s_nopWeightsTotal;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    DestChainConfigArgs[] memory destChainConfigArgs
  ) {
    if (
      staticConfig.linkToken == address(0) || staticConfig.chainSelector == 0 || staticConfig.rmnProxy == address(0)
        || staticConfig.nonceManager == address(0) || staticConfig.tokenAdminRegistry == address(0)
    ) {
      revert InvalidConfig();
    }

    i_linkToken = staticConfig.linkToken;
    i_chainSelector = staticConfig.chainSelector;
    i_maxFeeJuelsPerMsg = staticConfig.maxFeeJuelsPerMsg;
    i_rmnProxy = staticConfig.rmnProxy;
    i_nonceManager = staticConfig.nonceManager;
    i_tokenAdminRegistry = staticConfig.tokenAdminRegistry;

    _setDynamicConfig(dynamicConfig);
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  // ================================================================
  // │                          Messaging                           │
  // ================================================================

  /// @notice Gets the next sequence number to be used in the onRamp
  /// @param destChainSelector The destination chain selector
  /// @return the next sequence number to be used
  function getExpectedNextSequenceNumber(uint64 destChainSelector) external view returns (uint64) {
    return s_destChainConfig[destChainSelector].sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2AnyOnRampClient
  function forwardFromRouter(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external returns (bytes32) {
    DestChainConfig storage destChainConfig = s_destChainConfig[destChainSelector];
    Internal.EVM2AnyRampMessage memory newMessage =
      _generateNewMessage(destChainConfig, destChainSelector, message, feeTokenAmount, originalSender);

    // Lock the tokens as last step. TokenPools may not always be trusted.
    // There should be no state changes after external call to TokenPools.
    for (uint256 i = 0; i < newMessage.tokenAmounts.length; ++i) {
      Client.EVMTokenAmount memory tokenAndAmount = message.tokenAmounts[i];

      if (tokenAndAmount.amount == 0) revert CannotSendZeroTokens();

      IPoolV1 sourcePool = getPoolBySourceToken(destChainSelector, IERC20(tokenAndAmount.token));
      // We don't have to check if it supports the pool version in a non-reverting way here because
      // if we revert here, there is no effect on CCIP. Therefore we directly call the supportsInterface
      // function and not through the ERC165Checker.
      if (address(sourcePool) == address(0) || !sourcePool.supportsInterface(Pool.CCIP_POOL_V1)) {
        revert UnsupportedToken(tokenAndAmount.token);
      }

      Pool.LockOrBurnOutV1 memory poolReturnData = sourcePool.lockOrBurn(
        Pool.LockOrBurnInV1({
          receiver: message.receiver,
          remoteChainSelector: destChainSelector,
          originalSender: originalSender,
          amount: tokenAndAmount.amount,
          localToken: tokenAndAmount.token
        })
      );

      // Since the DON has to pay for the extraData to be included on the destination chain, we cap the length of the
      // extraData. This prevents gas bomb attacks on the NOPs. As destBytesOverhead accounts for both
      // extraData and offchainData, this caps the worst case abuse to the number of bytes reserved for offchainData.
      if (poolReturnData.destPoolData.length > Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
        // TODO: re-add validation
        // && poolReturnData.destPoolData.length
        // > s_tokenTransferFeeConfig[destChainSelector][tokenAndAmount.token].destBytesOverhead

        revert SourceTokenDataTooLarge(tokenAndAmount.token);
      }

      _validateDestFamilyAddress(destChainConfig.dynamicConfig.chainFamilySelector, poolReturnData.destTokenAddress);

      newMessage.tokenAmounts[i] = Internal.RampTokenAmount({
        sourcePoolAddress: abi.encode(sourcePool),
        destTokenAddress: poolReturnData.destTokenAddress,
        extraData: poolReturnData.destPoolData,
        amount: tokenAndAmount.amount
      });
    }

    // Hash only after the sourceTokenData has been set
    newMessage.header.messageId = Internal._hash(newMessage, destChainConfig.metadataHash);

    // Emit message request
    // This must happen after any pool events as some tokens (e.g. USDC) emit events that we expect to precede this
    // event in the offchain code.
    emit CCIPSendRequested(destChainSelector, newMessage);
    return newMessage.header.messageId;
  }

  /// @notice Helper function to relieve stack pressure from `forwardFromRouter`
  /// @param destChainConfig The destination chain config storage pointer
  /// @param destChainSelector The destination chain selector
  /// @param message Message struct to send
  /// @param feeTokenAmount Amount of fee tokens for payment
  /// @param originalSender The original initiator of the CCIP request
  function _generateNewMessage(
    DestChainConfig storage destChainConfig,
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) internal returns (Internal.EVM2AnyRampMessage memory) {
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destChainSelector)))) revert CursedByRMN(destChainSelector);
    // Validate message sender is set and allowed. Not validated in `getFee` since it is not user-driven.
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Router address may be zero intentionally to pause.
    if (msg.sender != s_dynamicConfig.router) revert MustBeCalledByRouter();
    if (!destChainConfig.dynamicConfig.isEnabled) revert DestinationChainNotEnabled(destChainSelector);

    // Validate the message with various checks
    uint256 numberOfTokens = message.tokenAmounts.length;
    // TODO: optimization - consider removing this call, since the Router calls into getFee -> _validateMessage
    // _validateMessage(destChainSelector, message.data.length, numberOfTokens, message.receiver);

    // Only check token value if there are tokens
    if (numberOfTokens > 0) {
      address messageValidator = s_dynamicConfig.messageValidator;
      if (messageValidator != address(0)) {
        try IMessageInterceptor(messageValidator).onOutboundMessage(destChainSelector, message) {}
        catch (bytes memory err) {
          revert IMessageInterceptor.MessageValidationError(err);
        }
      }
    }

    uint256 msgFeeJuels;
    // Convert feeToken to link if not already in link
    if (message.feeToken == i_linkToken) {
      msgFeeJuels = feeTokenAmount;
    } else {
      msgFeeJuels =
        IPriceRegistry(s_dynamicConfig.priceRegistry).convertTokenAmount(message.feeToken, feeTokenAmount, i_linkToken);
    }

    emit FeePaid(message.feeToken, msgFeeJuels);

    if (msgFeeJuels > i_maxFeeJuelsPerMsg) revert MessageFeeTooHigh(msgFeeJuels, i_maxFeeJuelsPerMsg);

    // NOTE: when supporting non-EVM chains, revisit this and parse non-EVM args
    // Assumes strict ordering, unless the chain is of the EVM family and the extra args indicate out of order execution
    uint64 nonce = 0;
    if (!_parseEVMExtraArgsFromBytes(message.extraArgs, destChainConfig.dynamicConfig).allowOutOfOrderExecution) {
      // Only bump nonce for messages that specify allowOutOfOrderExecution == false. Otherwise, we
      // may block ordered message nonces, which is not what we want.
      nonce = INonceManager(i_nonceManager).getIncrementedOutboundNonce(destChainSelector, originalSender);
    }

    Internal.EVM2AnyRampMessage memory rampMessage = Internal.EVM2AnyRampMessage({
      header: Internal.RampMessageHeader({
        // Should be generated after the message is complete
        messageId: "",
        sourceChainSelector: i_chainSelector,
        destChainSelector: destChainSelector,
        // We need the next available sequence number so we increment before we use the value
        sequenceNumber: ++destChainConfig.sequenceNumber,
        nonce: nonce
      }),
      sender: originalSender,
      data: message.data,
      extraArgs: _convertParsedExtraArgs(message.extraArgs, destChainConfig.dynamicConfig),
      receiver: message.receiver,
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      tokenAmounts: new Internal.RampTokenAmount[](numberOfTokens) // should be populated after generation
    });

    return rampMessage;
  }

  // TODO: rm extra args parsing

  /// @dev Parses extraArgs with dest chain config family tag validation, and re-encodes the args to the latest arguments version.
  /// Used to generate an EVM2AnyRampMessage with the accurate representation of the parsed extraArgs.
  /// @param extraArgs The extra args bytes
  /// @param destChainDynamicConfig Dest chain config to validate against
  /// @return encodedExtraArgs the parsed & encoded extra args
  function _convertParsedExtraArgs(
    bytes calldata extraArgs,
    DestChainDynamicConfig memory destChainDynamicConfig
  ) internal pure returns (bytes memory encodedExtraArgs) {
    bytes4 chainFamilySelector = destChainDynamicConfig.chainFamilySelector;
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_EVM) {
      return abi.encode(_parseEVMExtraArgsFromBytes(extraArgs, destChainDynamicConfig));
    }
    // Invalid chain family selectors cannot be configured - ignore invalid cases
    return encodedExtraArgs;
  }

  /// @dev Convert the extra args bytes into a struct with validations against the dest chain config
  /// @param extraArgs The extra args bytes
  /// @param destChainDynamicConfig Dest chain config to validate against
  /// @return EVMExtraArgs the extra args struct (latest version)
  function _parseEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    DestChainDynamicConfig memory destChainDynamicConfig
  ) internal pure returns (Client.EVMExtraArgsV2 memory) {
    Client.EVMExtraArgsV2 memory evmExtraArgs =
      _parseUnvalidatedEVMExtraArgsFromBytes(extraArgs, destChainDynamicConfig.defaultTxGasLimit);

    if (evmExtraArgs.gasLimit > uint256(destChainDynamicConfig.maxPerMsgGasLimit)) revert MessageGasLimitTooHigh();
    if (destChainDynamicConfig.enforceOutOfOrder && !evmExtraArgs.allowOutOfOrderExecution) {
      revert ExtraArgOutOfOrderExecutionMustBeTrue();
    }

    return evmExtraArgs;
  }

  /// @dev Convert the extra args bytes into a struct
  /// @param extraArgs The extra args bytes
  /// @param defaultTxGasLimit default tx gas limit to use in the absence of extra args
  /// @return EVMExtraArgs the extra args struct (latest version)
  function _parseUnvalidatedEVMExtraArgsFromBytes(
    bytes calldata extraArgs,
    uint64 defaultTxGasLimit
  ) private pure returns (Client.EVMExtraArgsV2 memory) {
    if (extraArgs.length == 0) {
      // If extra args are empty, generate default values
      return Client.EVMExtraArgsV2({gasLimit: defaultTxGasLimit, allowOutOfOrderExecution: false});
    }

    bytes4 extraArgsTag = bytes4(extraArgs);
    bytes memory argsData = extraArgs[4:];

    if (extraArgsTag == Client.EVM_EXTRA_ARGS_V2_TAG) {
      return abi.decode(argsData, (Client.EVMExtraArgsV2));
    } else if (extraArgsTag == Client.EVM_EXTRA_ARGS_V1_TAG) {
      // EVMExtraArgsV1 originally included a second boolean (strict) field which has been deprecated.
      // Clients may still include it but it will be ignored.
      return Client.EVMExtraArgsV2({gasLimit: abi.decode(argsData, (uint256)), allowOutOfOrderExecution: false});
    }

    revert InvalidExtraArgsTag();
  }

  // TODO: revisit moving this to PriceRegistry
  /// @notice Validate the forwarded message with various checks.
  /// @dev This function can be called multiple times during a CCIPSend,
  /// only common user-driven mistakes are validated here to minimize duplicate validation cost.
  /// @param destChainSelector The destination chain selector.
  /// @param dataLength The length of the data field of the message.
  /// @param numberOfTokens The number of tokens to be sent.
  function _validateMessage(uint64 destChainSelector, uint256 dataLength, uint256 numberOfTokens) internal view {
    // Check that payload is formed correctly
    DestChainDynamicConfig storage destChainDynamicConfig = s_destChainConfig[destChainSelector].dynamicConfig;
    if (dataLength > uint256(destChainDynamicConfig.maxDataBytes)) {
      revert MessageTooLarge(uint256(destChainDynamicConfig.maxDataBytes), dataLength);
    }
    if (numberOfTokens > uint256(destChainDynamicConfig.maxNumberOfTokensPerMsg)) revert UnsupportedNumberOfTokens();
  }

  /// @notice Validates that the destAddress matches the expected format of the family.
  /// @param chainFamilySelector Tag to identify the target family
  /// @param destAddress Dest address to validate
  /// @dev precondition - assumes the family tag is correct and validated
  function _validateDestFamilyAddress(bytes4 chainFamilySelector, bytes memory destAddress) internal pure {
    if (chainFamilySelector == Internal.CHAIN_FAMILY_SELECTOR_EVM) {
      Internal._validateEVMAddress(destAddress);
    }
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static onRamp config.
  /// @dev RMN depends on this function, if changing, please notify the RMN maintainers.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({
      linkToken: i_linkToken,
      chainSelector: i_chainSelector,
      maxFeeJuelsPerMsg: i_maxFeeJuelsPerMsg,
      rmnProxy: i_rmnProxy,
      nonceManager: i_nonceManager,
      tokenAdminRegistry: i_tokenAdminRegistry
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
    if (dynamicConfig.priceRegistry == address(0) || dynamicConfig.feeAggregator == address(0)) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(
      StaticConfig({
        linkToken: i_linkToken,
        chainSelector: i_chainSelector,
        maxFeeJuelsPerMsg: i_maxFeeJuelsPerMsg,
        rmnProxy: i_rmnProxy,
        nonceManager: i_nonceManager,
        tokenAdminRegistry: i_tokenAdminRegistry
      }),
      dynamicConfig
    );
  }

  // ================================================================
  // │                      Tokens and pools                        │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRampClient
  function getPoolBySourceToken(uint64, /*destChainSelector*/ IERC20 sourceToken) public view returns (IPoolV1) {
    return IPoolV1(ITokenAdminRegistry(i_tokenAdminRegistry).getPool(address(sourceToken)));
  }

  /// @inheritdoc IEVM2AnyOnRampClient
  function getSupportedTokens(uint64 /*destChainSelector*/ ) external pure returns (address[] memory) {
    revert GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  }

  // ================================================================
  // │                             Fees                             │
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRampClient
  /// @dev getFee MUST revert if the feeToken is not listed in the fee token config, as the router assumes it does.
  /// @param destChainSelector The destination chain selector.
  /// @param message The message to get quote for.
  /// @return feeTokenAmount The amount of fee token needed for the fee, in smallest denomination of the fee token.
  function getFee(
    uint64 destChainSelector,
    Client.EVM2AnyMessage calldata message
  ) external view returns (uint256 feeTokenAmount) {
    DestChainDynamicConfig storage destChainDynamicConfig = s_destChainConfig[destChainSelector].dynamicConfig;

    // TODO: move check to PriceRegistry?
    if (!destChainDynamicConfig.isEnabled) revert DestinationChainNotEnabled(destChainSelector);

    return IPriceRegistry(s_dynamicConfig.priceRegistry).getFee(destChainSelector, message);
  }

  /// @notice Updates the destination chain specific config.
  /// @param destChainConfigArgs Array of source chain specific configs.
  function applyDestChainConfigUpdates(DestChainConfigArgs[] memory destChainConfigArgs) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  /// @notice Internal version of applyDestChainConfigUpdates.
  function _applyDestChainConfigUpdates(DestChainConfigArgs[] memory destChainConfigArgs) internal {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArgs[i].destChainSelector;

      // NOTE: when supporting non-EVM chains, update chainFamilySelector validations
      if (
        destChainSelector == 0 || destChainConfigArg.dynamicConfig.defaultTxGasLimit == 0
          || destChainConfigArg.dynamicConfig.chainFamilySelector != Internal.CHAIN_FAMILY_SELECTOR_EVM
      ) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      DestChainConfig storage destChainConfig = s_destChainConfig[destChainSelector];
      address prevOnRamp = destChainConfigArg.prevOnRamp;

      DestChainConfig memory newDestChainConfig = DestChainConfig({
        dynamicConfig: destChainConfigArg.dynamicConfig,
        prevOnRamp: prevOnRamp,
        sequenceNumber: destChainConfig.sequenceNumber,
        metadataHash: destChainConfig.metadataHash
      });

      destChainConfig.dynamicConfig = newDestChainConfig.dynamicConfig;

      if (destChainConfig.metadataHash == 0) {
        newDestChainConfig.metadataHash =
          keccak256(abi.encode(Internal.EVM_2_ANY_MESSAGE_HASH, i_chainSelector, destChainSelector, address(this)));
        destChainConfig.metadataHash = newDestChainConfig.metadataHash;
        if (prevOnRamp != address(0)) destChainConfig.prevOnRamp = prevOnRamp;

        emit DestChainAdded(destChainSelector, destChainConfig);
      } else {
        if (destChainConfig.prevOnRamp != prevOnRamp) revert InvalidDestChainConfig(destChainSelector);
        // TODO: move validation to PriceRegistry
        // if (destChainConfigArg.dynamicConfig.defaultTokenDestBytesOverhead < Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES) {
        //   revert InvalidDestBytesOverhead(address(0), destChainConfigArg.dynamicConfig.defaultTokenDestBytesOverhead);
        // }

        emit DestChainDynamicConfigUpdated(destChainSelector, destChainConfigArg.dynamicConfig);
      }
    }
  }

  /// @notice Returns the destination chain config for given destination chain selector.
  /// @param destChainSelector The destination chain selector.
  /// @return The destination chain config.
  function getDestChainConfig(uint64 destChainSelector) external view returns (DestChainConfig memory) {
    return s_destChainConfig[destChainSelector];
  }

  /// @notice Withdraws the outstanding fee token balances to the fee aggregator.
  /// @dev This function can be permissionless as it only transfers accepted fee tokens to the fee aggregator which is a trusted address.
  function withdrawFeeTokens() external {
    address[] memory feeTokens = IPriceRegistry(s_dynamicConfig.priceRegistry).getFeeTokens();
    address feeAggregator = s_dynamicConfig.feeAggregator;

    for (uint256 i = 0; i < feeTokens.length; ++i) {
      IERC20 feeToken = IERC20(feeTokens[i]);
      uint256 feeTokenBalance = feeToken.balanceOf(address(this));

      if (feeTokenBalance > 0) {
        feeToken.safeTransfer(feeAggregator, feeTokenBalance);

        emit FeeTokenWithdrawn(feeAggregator, address(feeToken), feeTokenBalance);
      }
    }
  }
}
