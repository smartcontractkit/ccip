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
  error MessageGasLimitTooHigh();
  error UnsupportedToken(address token);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  // TODO: rename to InvalidStaticConfig
  error InvalidConfig();
  error CursedByRMN(uint64 sourceChainSelector);
  error GetSupportedTokensFunctionalityRemovedCheckAdminRegistry();
  error InvalidDestChainConfig(uint64 destChainSelector);
  error DestChainAlreadyConfigured(uint64 destChainSelector);
  error DestinationChainNotEnabled(uint64 destChainSelector);

  event AdminSet(address newAdmin);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event FeePaid(address indexed feeToken, uint256 feeValueJuels);
  event FeeTokenWithdrawn(address indexed feeAggregator, address indexed feeToken, uint256 amount);
  /// RMN depends on this event, if changing, please notify the RMN maintainers.
  event CCIPSendRequested(uint64 indexed destChainSelector, Internal.EVM2AnyRampMessage message);
  event DestChainAdded(uint64 indexed destChainSelector, DestChainConfig destChainConfig);

  /// @dev Struct that contains the static configuration
  /// RMN depends on this struct, if changing, please notify the RMN maintainers.
  // solhint-disable-next-line gas-struct-packing
  struct StaticConfig {
    uint64 chainSelector; // ─────╮ Source chainSelector
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

  /// @dev Struct to hold the configs for a destination chain
  struct DestChainConfig {
    address prevOnRamp; // ────────────────────╮ Address of previous-version OnRamp
    uint64 sequenceNumber; // ─────────────────╯ The last used sequence number. This is zero in the case where no messages has been sent yet.
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
    uint64 destChainSelector; // ─────────╮ Destination chain selector
    address prevOnRamp; // ───────────────╯ Address of previous-version OnRamp.
  }

  // STATIC CONFIG
  string public constant override typeAndVersion = "EVM2EVMMultiOnRamp 1.6.0-dev";
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
      staticConfig.chainSelector == 0 || staticConfig.rmnProxy == address(0) || staticConfig.nonceManager == address(0)
        || staticConfig.tokenAdminRegistry == address(0)
    ) {
      revert InvalidConfig();
    }

    i_chainSelector = staticConfig.chainSelector;
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
    // TODO: inline generateNewMessage
    Internal.EVM2AnyRampMessage memory newMessage =
      _generateNewMessage(destChainConfig, destChainSelector, message, feeTokenAmount, originalSender);

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
    // NOTE: assumes the message has already been validated through the getFee call
    // Validate message sender is set and allowed. Not validated in `getFee` since it is not user-driven.
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Router address may be zero intentionally to pause.
    if (msg.sender != s_dynamicConfig.router) revert MustBeCalledByRouter();

    // Validate the message with various checks
    uint256 numberOfTokens = message.tokenAmounts.length;

    Internal.EVM2AnyRampMessage memory rampMessage = Internal.EVM2AnyRampMessage({
      header: Internal.RampMessageHeader({
        // Should be generated after the message is complete
        messageId: "",
        sourceChainSelector: i_chainSelector,
        destChainSelector: destChainSelector,
        // We need the next available sequence number so we increment before we use the value
        sequenceNumber: ++destChainConfig.sequenceNumber,
        nonce: 0
      }),
      sender: originalSender,
      data: message.data,
      extraArgs: message.extraArgs,
      receiver: message.receiver,
      feeToken: message.feeToken,
      feeTokenAmount: feeTokenAmount,
      tokenAmounts: new Internal.RampTokenAmount[](numberOfTokens) // should be populated after generation
    });

    // Lock the tokens as last step. TokenPools may not always be trusted.
    // There should be no state changes after external call to TokenPools.
    for (uint256 i = 0; i < numberOfTokens; ++i) {
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

      // NOTE: pool validations are outsourced to the PriceRegistry to handle family-specific logic handling

      rampMessage.tokenAmounts[i] = Internal.RampTokenAmount({
        sourcePoolAddress: abi.encode(sourcePool),
        destTokenAddress: poolReturnData.destTokenAddress,
        extraData: poolReturnData.destPoolData,
        amount: tokenAndAmount.amount
      });
    }

    (uint256 msgFeeJuels, bool isOutOfOrderExecution) =
      IPriceRegistry(s_dynamicConfig.priceRegistry).getValidatedRampMessageParams(rampMessage, message.tokenAmounts);
    emit FeePaid(message.feeToken, msgFeeJuels);

    if (!isOutOfOrderExecution) {
      // Only bump nonce for messages that specify allowOutOfOrderExecution == false. Otherwise, we
      // may block ordered message nonces, which is not what we want.
      rampMessage.header.nonce =
        INonceManager(i_nonceManager).getIncrementedOutboundNonce(destChainSelector, originalSender);
    }

    address messageValidator = s_dynamicConfig.messageValidator;
    if (messageValidator != address(0)) {
      try IMessageInterceptor(messageValidator).onOutboundMessage(destChainSelector, message) {}
      catch (bytes memory err) {
        revert IMessageInterceptor.MessageValidationError(err);
      }
    }

    // Hash only after all fields have been set
    rampMessage.header.messageId = Internal._hash(rampMessage, destChainConfig.metadataHash);

    return rampMessage;
  }

  // ================================================================
  // │                           Config                             │
  // ================================================================

  /// @notice Returns the static onRamp config.
  /// @dev RMN depends on this function, if changing, please notify the RMN maintainers.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({
      chainSelector: i_chainSelector,
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
        chainSelector: i_chainSelector,
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
    if (IRMN(i_rmnProxy).isCursed(bytes16(uint128(destChainSelector)))) revert CursedByRMN(destChainSelector);
    if (s_destChainConfig[destChainSelector].metadataHash == bytes32("")) {
      revert DestinationChainNotEnabled(destChainSelector);
    }

    return IPriceRegistry(s_dynamicConfig.priceRegistry).getValidatedFee(destChainSelector, message);
  }

  // TODO: revisit removing dest chain configs (fetch through PriceRegistry)
  /// @notice Updates the destination chain specific config.
  /// @param destChainConfigArgs Array of source chain specific configs.
  function applyDestChainConfigUpdates(DestChainConfigArgs[] memory destChainConfigArgs) external onlyOwner {
    _applyDestChainConfigUpdates(destChainConfigArgs);
  }

  /// @notice Internal version of applyDestChainConfigUpdates.
  function _applyDestChainConfigUpdates(DestChainConfigArgs[] memory destChainConfigArgs) internal {
    for (uint256 i = 0; i < destChainConfigArgs.length; ++i) {
      DestChainConfigArgs memory destChainConfigArg = destChainConfigArgs[i];
      uint64 destChainSelector = destChainConfigArg.destChainSelector;

      if (destChainSelector == 0) {
        revert InvalidDestChainConfig(destChainSelector);
      }

      DestChainConfig storage destChainConfig = s_destChainConfig[destChainSelector];

      if (destChainConfig.metadataHash == 0) {
        DestChainConfig memory newDestChainConfig = DestChainConfig({
          prevOnRamp: destChainConfigArg.prevOnRamp,
          // Sequence numbers start at 0 for newly configured chains
          sequenceNumber: 0,
          metadataHash: keccak256(
            abi.encode(Internal.EVM_2_ANY_MESSAGE_HASH, i_chainSelector, destChainSelector, address(this))
            )
        });

        s_destChainConfig[destChainSelector] = newDestChainConfig;

        emit DestChainAdded(destChainSelector, newDestChainConfig);
      } else {
        revert DestChainAlreadyConfigured(destChainSelector);
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
