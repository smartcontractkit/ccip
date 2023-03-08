// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IEVM2EVMOnRamp} from "../interfaces/onRamp/IEVM2EVMOnRamp.sol";
import {IPriceRegistry} from "../interfaces/prices/IPriceRegistry.sol";
import {IEVM2AnyOnRamp} from "../interfaces/onRamp/IEVM2AnyOnRamp.sol";
import {IAggregateRateLimiter} from "../interfaces/rateLimiter/IAggregateRateLimiter.sol";

import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {Client} from "../models/Client.sol";
import {Internal} from "../models/Internal.sol";
import {EnumerableMapAddresses} from "../../libraries/internal/EnumerableMapAddresses.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableSet.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableMap.sol";
import {Pausable} from "../../vendor/Pausable.sol";

contract EVM2EVMOnRamp is IEVM2EVMOnRamp, Pausable, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.AddressToUintMap;
  using EnumerableMapAddresses for EnumerableMapAddresses.AddressToAddressMap;
  using EnumerableSet for EnumerableSet.AddressSet;

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
  /// @dev The amount of LINK available to pay NOPS
  uint256 internal s_nopFeesJuels;
  /// @dev The total weight of all NOPs weights
  uint256 internal s_nopWeightsTotal;
  /// @dev The current nonce per sender
  mapping(address => uint64) internal s_senderNonce;
  /// @dev The last used sequence number. This is zero in the case where no
  /// messages has been sent yet. 0 is not a valid sequence number for any
  /// real transaction.
  uint64 internal s_sequenceNumber;

  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    TokenAndPool[] memory tokensAndPools,
    address[] memory allowlist,
    IAggregateRateLimiter.RateLimiterConfig memory rateLimiterConfig,
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

  /// @inheritdoc IEVM2AnyOnRamp
  function getPoolBySourceToken(IERC20 sourceToken) external view override returns (IPool) {
    return _getPoolBySourceToken(sourceToken);
  }

  /// @dev Get the pool for the given source token
  /// @param sourceToken The source token to get the pool for
  /// @return The pool for the given source token
  function _getPoolBySourceToken(IERC20 sourceToken) private view returns (IPool) {
    if (!s_poolsBySourceToken.contains(address(sourceToken))) revert UnsupportedToken(sourceToken);
    return IPool(s_poolsBySourceToken.get(address(sourceToken)));
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

    _removeTokens(tokenAmounts);
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function forwardFromRouter(
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external override whenNotPaused whenHealthy returns (bytes32) {
    Client.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    // Validate the message with various checks
    _validateMessage(message.data.length, extraArgs.gasLimit, message.tokenAmounts, originalSender);
    if (message.receiver.length != 32) revert InvalidAddress(message.receiver);
    uint256 decodedReceiver = abi.decode(message.receiver, (uint256));
    if (decodedReceiver > type(uint160).max) revert InvalidAddress(message.receiver);

    // Convert feeToken to link if not already in link
    if (message.feeToken == i_linkToken) {
      s_nopFeesJuels += feeTokenAmount;
    } else {
      s_nopFeesJuels += IPriceRegistry(s_dynamicConfig.priceRegistry).convertFeeTokenAmountToLinkAmount(
        i_linkToken,
        message.feeToken,
        feeTokenAmount
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

  /// @inheritdoc IEVM2AnyOnRamp
  function getFee(Client.EVM2AnyMessage calldata message) public view override returns (uint256 fee) {
    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    uint256 feeTokenBaseUnitsPerUnitGas = IPriceRegistry(s_dynamicConfig.priceRegistry).getFeeTokenBaseUnitsPerUnitGas(
      message.feeToken,
      i_destChainId
    );
    if (feeTokenBaseUnitsPerUnitGas == 0) revert TokenOrChainNotSupported(message.feeToken, i_destChainId);

    // NOTE: if a fee token is not configured, formula below will intentionally
    // return zero, i.e. zeroing the fees for that feeToken.
    FeeTokenConfig memory feeTokenConfig = s_feeTokenConfig[message.feeToken];
    return
      feeTokenConfig.feeAmount + // Flat fee
      ((gasLimit + feeTokenConfig.destGasOverhead) * feeTokenBaseUnitsPerUnitGas * feeTokenConfig.multiplier) / // Total gas reserved for tx
      1 ether; // latest gas reported gas fee with a safety margin
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getExpectedNextSequenceNumber() external view override returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getStaticConfig() external view override returns (IEVM2EVMOnRamp.StaticConfig memory) {
    return
      IEVM2EVMOnRamp.StaticConfig({
        linkToken: i_linkToken,
        chainId: i_chainId,
        destChainId: i_destChainId,
        defaultTxGasLimit: i_defaultTxGasLimit
      });
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getDynamicConfig() external view override returns (DynamicConfig memory config) {
    return s_dynamicConfig;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setDynamicConfig(DynamicConfig memory dynamicConfig) external override onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Internal version of setDynamicConfig to allow for reuse in the constructor.
  function _setDynamicConfig(DynamicConfig memory dynamicConfig) internal {
    if (
      dynamicConfig.router == address(0) || dynamicConfig.priceRegistry == address(0) || dynamicConfig.afn == address(0)
    ) revert InvalidConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(
      IEVM2EVMOnRamp.StaticConfig({
        linkToken: i_linkToken,
        chainId: i_chainId,
        destChainId: i_destChainId,
        defaultTxGasLimit: i_defaultTxGasLimit
      }),
      dynamicConfig
    );
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getNopFeesJuels() external view override returns (uint256) {
    return s_nopFeesJuels;
  }

  /// #@inheritdoc IEVM2AnyOnRamp
  /// @dev This method can only be called by the owner of the contract.
  function applyPoolUpdates(Internal.PoolUpdate[] memory adds, Internal.PoolUpdate[] memory removes) public onlyOwner {
    for (uint256 i = 0; i < adds.length; ++i) {
      address token = adds[i].token;
      address pool = adds[i].pool;

      if (token == address(0) || pool == address(0)) revert InvalidTokenPoolConfig();
      if (s_poolsBySourceToken.contains(token)) revert PoolAlreadyAdded();
      s_poolsBySourceToken.set(token, pool);

      emit PoolAdded(token, pool);
    }

    for (uint256 i = 0; i < removes.length; ++i) {
      address token = removes[i].token;
      address pool = removes[i].pool;

      if (!s_poolsBySourceToken.contains(token)) revert PoolDoesNotExist(token);
      if (s_poolsBySourceToken.get(token) != pool) revert TokenPoolMismatch();
      s_poolsBySourceToken.remove(token);

      emit PoolRemoved(token, pool);
    }
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getSupportedTokens() public view returns (address[] memory) {
    address[] memory sourceTokens = new address[](s_poolsBySourceToken.length());
    for (uint256 i = 0; i < sourceTokens.length; ++i) {
      (sourceTokens[i], ) = s_poolsBySourceToken.at(i);
    }
    return sourceTokens;
  }

  /// @inheritdoc IEVM2AnyOnRamp
  function getSenderNonce(address sender) external view override returns (uint64) {
    return s_senderNonce[sender];
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function withdrawNonLinkFees(address feeToken, address to) external override onlyOwner {
    if (feeToken == i_linkToken) revert InvalidFeeToken(feeToken);
    if (to == address(0)) revert InvalidWithdrawalAddress(to);
    IERC20(feeToken).safeTransfer(to, IERC20(feeToken).balanceOf(address(this)));
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setNops(NopAndWeight[] calldata nopsAndWeights) external onlyOwner {
    _setNops(nopsAndWeights);
  }

  /// @dev Set the nops and weights
  /// @param nopsAndWeights The nops and weights
  function _setNops(NopAndWeight[] memory nopsAndWeights) internal {
    // Remove previous
    delete s_nops;

    // Add new
    uint256 nopWeightsTotal = 0;
    for (uint256 i = 0; i < nopsAndWeights.length; ++i) {
      s_nops.set(nopsAndWeights[i].nop, nopsAndWeights[i].weight);
      nopWeightsTotal += nopsAndWeights[i].weight;
    }
    s_nopWeightsTotal = nopWeightsTotal;
    emit NopsSet(nopWeightsTotal, nopsAndWeights);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getNops() external view override returns (NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal) {
    uint256 length = s_nops.length();
    nopsAndWeights = new NopAndWeight[](length);
    for (uint256 i = 0; i < length; ++i) {
      (address nopAddress, uint256 nopWeight) = s_nops.at(i);
      nopsAndWeights[i] = NopAndWeight({nop: nopAddress, weight: nopWeight});
    }
    weightsTotal = s_nopWeightsTotal;
    return (nopsAndWeights, weightsTotal);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getFeeConfig(address token) external view override returns (FeeTokenConfig memory config) {
    return s_feeTokenConfig[token];
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setFeeConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) external override onlyOwnerOrFeeAdmin {
    _setFeeConfig(feeTokenConfigs);
  }

  /// @inheritdoc IEVM2EVMOnRamp
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

  /// @inheritdoc IEVM2EVMOnRamp
  function setAllowListEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowListEnabledSet(enabled);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getAllowListEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getAllowList() external view returns (address[] memory) {
    address[] memory allowList = new address[](s_allowList.length());
    for (uint256 i = 0; i < s_allowList.length(); ++i) {
      allowList[i] = s_allowList.at(i);
    }
    return allowList;
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

  /// @inheritdoc IEVM2EVMOnRamp
  function payNops() external onlyOwnerOrFeeAdminOrNop {
    uint256 weightsTotal = s_nopWeightsTotal;
    if (weightsTotal == 0) revert NoNopsToPay();

    uint256 totalFeesToPay = s_nopFeesJuels;
    if (totalFeesToPay == 0 || totalFeesToPay < weightsTotal) revert NoFeesToPay();

    uint256 contractBalance = IERC20(i_linkToken).balanceOf(address(this));
    if (contractBalance < totalFeesToPay) revert InsufficientBalance();

    uint256 nopFee = (totalFeesToPay * 1e18) / weightsTotal;
    for (uint256 i = 0; i < s_nops.length(); ++i) {
      (address nop, uint256 weight) = s_nops.at(i);
      uint256 amount = (nopFee * weight) / 1e18;
      IERC20(i_linkToken).safeTransfer(nop, amount);
      totalFeesToPay -= amount;
      emit NopPaid(nop, amount);
    }
    s_nopFeesJuels = totalFeesToPay;
  }

  /// @dev Require that the sender is the owner or the fee admin or a nop
  modifier onlyOwnerOrFeeAdminOrNop() {
    if (msg.sender != owner() && msg.sender != s_dynamicConfig.feeAdmin && !s_nops.contains(msg.sender))
      revert OnlyCallableByOwnerOrFeeAdminOrNop();
    _;
  }

  /// @dev Require that the sender is the owner or the fee admin
  modifier onlyOwnerOrFeeAdmin() {
    if (msg.sender != owner() && msg.sender != s_dynamicConfig.feeAdmin) revert OnlyCallableByOwnerOrFeeAdmin();
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
}
