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
import {EnumerableMapAddresses} from "../../libraries/internal/EnumerableMapAddresses.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {Pausable} from "../../vendor/Pausable.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableMap.sol";

contract EVM2EVMOnRamp is IEVM2AnyOnRamp, Pausable, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.AddressToUintMap;
  using EnumerableMapAddresses for EnumerableMapAddresses.AddressToAddressMap;
  using EnumerableSet for EnumerableSet.AddressSet;

  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);
  error OnlyCallableByOwnerOrFeeAdmin();
  error OnlyCallableByOwnerOrFeeAdminOrNop();
  error InvalidWithdrawalAddress(address addr);
  error InvalidFeeToken(address token);
  error NoFeesToPay();
  error NoNopsToPay();
  error InsufficientBalance();
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

  event AllowListAdd(address sender);
  event AllowListRemove(address sender);
  event AllowListEnabledSet(bool enabled);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event NopPaid(address indexed nop, uint256 amount);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event CCIPSendRequested(Internal.EVM2EVMMessage message);
  event NopsSet(uint256 nopWeightsTotal, NopAndWeight[] nopsAndWeights);
  event PoolAdded(address token, address pool);
  event PoolRemoved(address token, address pool);

  /// @dev Struct that contains the static configuration
  struct StaticConfig {
    address linkToken; // --------┐ Link token address
    uint64 chainId; // -----------┘ Source chain Id
    uint64 destChainId; // -------┐ Destination chain Id
    uint64 defaultTxGasLimit; // -┘ Default gas limit for a tx
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

  /// @dev Struct to hold the fee configuration for a token
  struct FeeTokenConfig {
    uint96 feeAmount; // --------┐ Flat fee
    uint64 multiplier; //        | Price multiplier for gas costs
    uint32 destGasOverhead; // --┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Struct to hold the fee configuration for a token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // ---------┐ Token address
    uint64 multiplier; // -----┘ Price multiplier for gas costs
    uint96 feeAmount; // ------┐ Flat fee in feeToken
    uint32 destGasOverhead; //-┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Nop address and weight, used to set the nops and their weights
  struct NopAndWeight {
    address nop; // ----┐ Address of the node operator
    uint16 weight; // --┘ Weight for nop rewards
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
  uint64 internal immutable i_chainId;
  /// @dev The chain ID of the destination chain
  uint64 internal immutable i_destChainId;

  // DYNAMIC CONFIG
  /// @dev The config for the onRamp
  DynamicConfig internal s_dynamicConfig;
  /// @dev (address nop => uint256 weight)
  EnumerableMap.AddressToUintMap internal s_nops;
  /// @dev source token => token pool
  EnumerableMapAddresses.AddressToAddressMap private s_poolsBySourceToken;
  /// @dev Whether s_allowList is enabled or not.
  bool private s_allowlistEnabled;
  /// @dev A set of addresses which can make ccipSend calls.
  EnumerableSet.AddressSet private s_allowList;
  /// @dev The fee token config that can be set by the owner or fee admin
  mapping(address => FeeTokenConfig) internal s_feeTokenConfig;

  // STATE
  /// @dev The current nonce per sender
  mapping(address => uint64) internal s_senderNonce;
  /// @dev The amount of LINK available to pay NOPS
  uint96 internal s_nopFeesJuels;
  /// @dev The total weight of all NOPs weights
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
    NopAndWeight[] memory nopsAndWeights
  ) Pausable() AggregateRateLimiter(rateLimiterConfig) {
    if (
      staticConfig.linkToken == address(0) ||
      staticConfig.chainId == 0 ||
      staticConfig.destChainId == 0 ||
      staticConfig.defaultTxGasLimit == 0
    ) revert InvalidConfig();

    i_metadataHash = keccak256(
      abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, staticConfig.chainId, staticConfig.destChainId, address(this))
    );
    i_linkToken = staticConfig.linkToken;
    i_chainId = staticConfig.chainId;
    i_destChainId = staticConfig.destChainId;
    i_defaultTxGasLimit = staticConfig.defaultTxGasLimit;

    _setDynamicConfig(dynamicConfig);
    _setFeeConfig(feeTokenConfigs);
    _setNops(nopsAndWeights);

    // Set new tokens and pools
    for (uint256 i = 0; i < tokensAndPools.length; ++i) {
      if (tokensAndPools[i].token == address(0) || address(tokensAndPools[i].pool) == address(0))
        revert InvalidConfig();
      s_poolsBySourceToken.set(tokensAndPools[i].token, tokensAndPools[i].pool);
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

    // Lock the tokens
    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      Client.EVMTokenAmount memory tokenAndAmount = message.tokenAmounts[i];
      IPool pool = _getPoolBySourceToken(IERC20(tokenAndAmount.token));
      pool.lockOrBurn(tokenAndAmount.amount, originalSender);
    }

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    Internal.EVM2EVMMessage memory newMessage = Internal.EVM2EVMMessage({
      sourceChainId: i_chainId,
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
    if (bytes4(extraArgs[:4]) != Client.EVM_EXTRA_ARGS_V1_TAG)
      revert InvalidExtraArgsTag(Client.EVM_EXTRA_ARGS_V1_TAG, bytes4(extraArgs[:4]));
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
  ) internal {
    if (msg.sender != s_dynamicConfig.router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_dynamicConfig.maxDataSize))
      revert MessageTooLarge(uint256(s_dynamicConfig.maxDataSize), dataLength);
    if (gasLimit > uint256(s_dynamicConfig.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (tokenAmounts.length > uint256(s_dynamicConfig.maxTokensLength)) revert UnsupportedNumberOfTokens();
    if (s_allowlistEnabled && !s_allowList.contains(originalSender)) revert SenderNotAllowed(originalSender);

    _rateLimitValue(tokenAmounts);
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
        chainId: i_chainId,
        destChainId: i_destChainId,
        defaultTxGasLimit: i_defaultTxGasLimit
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
        chainId: i_chainId,
        destChainId: i_destChainId,
        defaultTxGasLimit: i_defaultTxGasLimit
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
  function getPoolBySourceToken(IERC20 sourceToken) external view returns (IPool) {
    return _getPoolBySourceToken(sourceToken);
  }

  /// @dev Get the pool for the given source token
  /// @param sourceToken The source token to get the pool for
  /// @return The pool for the given source token
  function _getPoolBySourceToken(IERC20 sourceToken) private view returns (IPool) {
    if (!s_poolsBySourceToken.contains(address(sourceToken))) revert UnsupportedToken(sourceToken);
    return IPool(s_poolsBySourceToken.get(address(sourceToken)));
  }

  /// #@inheritdoc IEVM2AnyOnRamp
  /// @dev This method can only be called by the owner of the contract.
  function applyPoolUpdates(Internal.PoolUpdate[] memory removes, Internal.PoolUpdate[] memory adds) public onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      address token = removes[i].token;
      address pool = removes[i].pool;

      if (!s_poolsBySourceToken.contains(token)) revert PoolDoesNotExist(token);
      if (s_poolsBySourceToken.get(token) != pool) revert TokenPoolMismatch();
      s_poolsBySourceToken.remove(token);

      emit PoolRemoved(token, pool);
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      address token = adds[i].token;
      address pool = adds[i].pool;

      if (token == address(0) || pool == address(0)) revert InvalidTokenPoolConfig();
      if (s_poolsBySourceToken.contains(token)) revert PoolAlreadyAdded();
      s_poolsBySourceToken.set(token, pool);

      emit PoolAdded(token, pool);
    }
  }

  // ================================================================
  // |                             Fees                             |
  // ================================================================

  /// @inheritdoc IEVM2AnyOnRamp
  function getFee(Client.EVM2AnyMessage calldata message) public view returns (uint256 fee) {
    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    uint256 feeTokenBaseUnitsPerUnitGas = IPriceRegistry(s_dynamicConfig.priceRegistry).getFeeTokenBaseUnitsPerUnitGas(
      message.feeToken,
      i_destChainId
    );

    // NOTE: if a fee token is not configured, formula below will intentionally
    // return zero, i.e. zeroing the fees for that feeToken.
    FeeTokenConfig memory feeTokenConfig = s_feeTokenConfig[message.feeToken];
    return
      feeTokenConfig.feeAmount + // Flat fee
      ((gasLimit + feeTokenConfig.destGasOverhead) * feeTokenBaseUnitsPerUnitGas * feeTokenConfig.multiplier) / // Total gas reserved for tx
      1 ether; // latest gas reported gas fee with a safety margin
  }

  /// @notice Gets the fee configuration for a token
  /// @param token The token to get the fee configuration for
  /// @return feeTokenConfig FeeTokenConfig struct
  function getFeeConfig(address token) external view returns (FeeTokenConfig memory feeTokenConfig) {
    return s_feeTokenConfig[token];
  }

  /// @notice Sets the fee configuration for a token
  /// @param feeTokenConfigs Array of FeeTokenConfigArgs structs
  function setFeeConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) external onlyOwnerOrAdmin {
    _setFeeConfig(feeTokenConfigs);
  }

  /// @dev Set the fee config
  /// @param feeTokenConfigs The fee token configs
  function _setFeeConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) internal {
    for (uint256 i = 0; i < feeTokenConfigs.length; ++i) {
      s_feeTokenConfig[feeTokenConfigs[i].token] = FeeTokenConfig({
        feeAmount: feeTokenConfigs[i].feeAmount,
        multiplier: feeTokenConfigs[i].multiplier,
        destGasOverhead: feeTokenConfigs[i].destGasOverhead
      });
    }
    emit FeeConfigSet(feeTokenConfigs);
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
    // Remove previous
    delete s_nops;

    // Add new
    uint32 nopWeightsTotal = 0;
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      s_nops.set(nopsAndWeights[i].nop, nopsAndWeights[i].weight);
      nopWeightsTotal += nopsAndWeights[i].weight;
    }
    s_nopWeightsTotal = nopWeightsTotal;
    emit NopsSet(nopWeightsTotal, nopsAndWeights);
  }

  /// @notice Pays the Node Ops their outstanding balances.
  function payNops() external onlyOwnerOrAdminOrNop {
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
    s_nopFeesJuels = fundsLeft;
  }

  /// @notice Allows the owner to withdraw any ERC20 token that is not the fee token
  /// @param feeToken The token to withdraw
  /// @param to The address to send the tokens to
  function withdrawNonLinkFees(address feeToken, address to) external onlyOwner {
    if (feeToken == i_linkToken) revert InvalidFeeToken(feeToken);
    if (to == address(0)) revert InvalidWithdrawalAddress(to);
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
  /// @param adds The added addresses.
  /// @param adds The removed addresses.
  function applyAllowListUpdates(address[] calldata adds, address[] calldata removes) external onlyOwner {
    _applyAllowListUpdates(adds, removes);
  }

  /// @notice Internal version of applyAllowListUpdates to allow for reuse in the constructor.
  function _applyAllowListUpdates(address[] memory adds, address[] memory removes) internal {
    for (uint256 i = 0; i < removes.length; ++i) {
      s_allowList.remove(removes[i]);
      emit AllowListRemove(removes[i]);
    }
    for (uint256 i = 0; i < adds.length; ++i) {
      s_allowList.add(adds[i]);
      emit AllowListAdd(adds[i]);
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
    return !IAFN(s_dynamicConfig.afn).badSignalReceived();
  }

  /// @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IAFN(s_dynamicConfig.afn).badSignalReceived()) revert BadAFNSignal();
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
