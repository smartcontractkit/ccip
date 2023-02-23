// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IEVM2EVMOnRamp} from "../interfaces/onRamp/IEVM2EVMOnRamp.sol";
import {IPriceRegistry} from "../interfaces/prices/IPriceRegistry.sol";
import {IEVM2AnyOnRamp} from "../interfaces/onRamp/IEVM2AnyOnRamp.sol";
import {IAggregateRateLimiter} from "../interfaces/rateLimiter/IAggregateRateLimiter.sol";

import {HealthChecker} from "../health/HealthChecker.sol";
import {AllowList} from "../access/AllowList.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {Client} from "../models/Client.sol";
import {Internal} from "../models/Internal.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";
import {EnumerableMap} from "../../vendor/openzeppelin-solidity/v4.7.3/contracts/utils/structs/EnumerableMap.sol";

contract EVM2EVMOnRamp is IEVM2EVMOnRamp, HealthChecker, AllowList, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.AddressToUintMap;

  struct TokenAndPool {
    address token;
    IPool pool;
  }

  struct PoolConfig {
    IPool pool;
    bool enabled;
  }

  /// @dev Default gas limit for EVM transactions
  uint256 private constant EVM_DEFAULT_GAS_LIMIT = 200_000;
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOnRamp 1.0.0";
  /// @dev The metadata hash for this contract
  bytes32 internal immutable i_metadataHash;
  /// @dev The chain ID of the source chain that this contract is deployed to
  uint64 internal immutable i_chainId;
  /// @dev The chain ID of the destination chain
  uint64 internal immutable i_destChainId;

  /// @dev The config for the onRamp
  OnRampConfig internal s_config;
  /// @dev The address of the router contract
  address internal s_router;

  /// @dev Special permissioned address that can set the fee for each token
  address internal s_feeAdmin;
  /// @dev The fee token config that can be set by the owner or fee admin
  mapping(address => FeeTokenConfig) internal s_feeTokenConfig;

  /// @dev The link token address - known to pay nops for their work
  address internal immutable i_linkToken;
  /// @dev The price registry address - used to get destination gas and token prices for conversions.
  address internal s_priceRegistry;

  /// @dev The amount of LINK available to pay NOPS
  uint256 internal s_nopFeesJuels;
  /// @dev The total weight of all NOPs weights
  uint256 internal s_nopWeightsTotal;
  /// @dev (address nop => uint256 weight)
  EnumerableMap.AddressToUintMap internal s_nops;

  /// @dev The current nonce per sender
  mapping(address => uint64) internal s_senderNonce;

  /// @dev The last used sequence number. This is zero in the case where no
  /// messages has been sent yet. 0 is not a valid sequence number for any
  /// real transaction.
  uint64 internal s_sequenceNumber;

  /// @dev source token => token pool
  mapping(IERC20 => PoolConfig) private s_poolsBySourceToken;
  /// @dev The list of source tokens that are supported
  address[] private s_sourceTokenList;

  constructor(
    Chains memory chainIds,
    TokenAndPool[] memory tokensAndPools,
    address[] memory allowlist,
    IAFN afn,
    OnRampConfig memory config,
    IAggregateRateLimiter.RateLimiterConfig memory rateLimiterConfig,
    address router,
    address priceRegistry,
    FeeTokenConfigArgs[] memory feeTokenConfigs,
    address linkToken,
    NopAndWeight[] memory nopsAndWeights
  ) HealthChecker(afn) AllowList(allowlist) AggregateRateLimiter(rateLimiterConfig) {
    if (router == address(0) || priceRegistry == address(0) || linkToken == address(0)) revert NullAddressNotAllowed();

    i_metadataHash = keccak256(
      abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, chainIds.chainId, chainIds.destChainId, address(this))
    );
    i_linkToken = linkToken;
    i_chainId = chainIds.chainId;
    i_destChainId = chainIds.destChainId;
    s_config = config;
    s_router = router;
    s_priceRegistry = priceRegistry;
    s_sequenceNumber = 0;
    _setFeeConfig(feeTokenConfigs);
    _setNops(nopsAndWeights);

    address[] memory newTokens = new address[](tokensAndPools.length);
    // Set new tokens and pools
    for (uint256 i = 0; i < tokensAndPools.length; ++i) {
      s_poolsBySourceToken[IERC20(tokensAndPools[i].token)] = PoolConfig({pool: tokensAndPools[i].pool, enabled: true});
      newTokens[i] = address(tokensAndPools[i].token);
    }
    s_sourceTokenList = newTokens;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getPoolBySourceToken(IERC20 sourceToken) external view override returns (IPool) {
    return _getPoolBySourceToken(sourceToken);
  }

  /// @dev Get the pool for the given source token
  /// @param sourceToken The source token to get the pool for
  /// @return The pool for the given source token
  function _getPoolBySourceToken(IERC20 sourceToken) private view returns (IPool) {
    PoolConfig memory poolConfig = s_poolsBySourceToken[sourceToken];
    if (poolConfig.enabled) {
      return s_poolsBySourceToken[sourceToken].pool;
    }
    revert UnsupportedToken(sourceToken);
  }

  /// @dev Convert the extra args bytes into a struct
  /// @param extraArgs The extra args bytes
  /// @return The extra args struct
  function _fromBytes(bytes calldata extraArgs) internal pure returns (Client.EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return Client.EVMExtraArgsV1({gasLimit: EVM_DEFAULT_GAS_LIMIT, strict: false});
    }
    if (bytes4(extraArgs[:4]) != Client.EVM_EXTRA_ARGS_V1_TAG)
      revert InvalidExtraArgsTag(Client.EVM_EXTRA_ARGS_V1_TAG, bytes4(extraArgs[:4]));
    return Client.EVMExtraArgsV1({gasLimit: abi.decode(extraArgs[4:36], (uint256)), strict: false});
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
    if (msg.sender != s_router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_config.maxDataSize)) revert MessageTooLarge(uint256(s_config.maxDataSize), dataLength);
    if (gasLimit > uint256(s_config.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (tokenAmounts.length > uint256(s_config.maxTokensLength)) revert UnsupportedNumberOfTokens();
    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

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

    // Convert feeToken to link if not already in link
    if (message.feeToken == i_linkToken) {
      s_nopFeesJuels += feeTokenAmount;
    } else {
      s_nopFeesJuels += IPriceRegistry(s_priceRegistry).convertFeeTokenAmountToLinkAmount(
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
      receiver: abi.decode(message.receiver, (address)),
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
    uint256 feeTokenBaseUnitsPerUnitGas = IPriceRegistry(s_priceRegistry).getFeeTokenBaseUnitsPerUnitGas(
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

  /// @inheritdoc IEVM2EVMOnRamp
  function getExpectedNextSequenceNumber() external view override returns (uint64) {
    return s_sequenceNumber + 1;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setRouter(address router) public override onlyOwner {
    s_router = router;
    emit RouterSet(router);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getRouter() external view override returns (address router) {
    return s_router;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setOnRampConfig(OnRampConfig calldata config) external override onlyOwner {
    s_config = config;
    emit OnRampConfigSet(config);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getOnRampConfig() external view override returns (OnRampConfig memory config) {
    return s_config;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getChainId() external view override returns (uint64) {
    return i_chainId;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getDestChainId() external view override returns (uint64) {
    return i_destChainId;
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getNopFeesJuels() external view override returns (uint256) {
    return s_nopFeesJuels;
  }

  /// @notice Add a new token pool
  /// @param token The source token
  /// @param pool The pool that will be used
  /// @dev This method can only be called by the owner of the contract.
  function addPool(IERC20 token, IPool pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    if (s_poolsBySourceToken[token].enabled) revert PoolAlreadyAdded();

    s_poolsBySourceToken[token] = PoolConfig({pool: pool, enabled: true});
    s_sourceTokenList.push(address(token));

    emit PoolAdded(token, pool);
  }

  /// @notice Remove a token pool
  /// @param token The source token
  /// @param pool The pool that will be removed
  /// @dev This method can only be called by the owner of the contract.
  function removePool(IERC20 token, IPool pool) public onlyOwner {
    PoolConfig memory oldConfig = s_poolsBySourceToken[token];
    // Check if the pool exists
    if (address(oldConfig.pool) == address(0)) revert PoolDoesNotExist(token);
    // Sanity check
    if (address(oldConfig.pool) != address(pool)) revert TokenPoolMismatch();

    s_poolsBySourceToken[token].enabled = false;

    emit PoolRemoved(token, pool);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getSupportedTokens() public view returns (address[] memory) {
    // TODO: We should just fully remove the pool in remove pool in the
    // same way we do for OffRamps? Seems better than keeping the full list
    // of once configured tokens then filter out disabled ones on every query
    uint256 numberOfSupportedTokens = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[IERC20(s_sourceTokenList[i])].enabled) {
        numberOfSupportedTokens++;
      }
    }

    address[] memory sourceTokens = new address[](numberOfSupportedTokens);
    uint256 j = 0;
    for (uint256 i = 0; i < s_sourceTokenList.length; ++i) {
      if (s_poolsBySourceToken[IERC20(s_sourceTokenList[i])].enabled) {
        sourceTokens[j++] = s_sourceTokenList[i];
      }
    }
    return sourceTokens;
  }

  /// @inheritdoc IEVM2EVMOnRamp
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
  function setFeeAdmin(address feeAdmin) external override onlyOwner {
    s_feeAdmin = feeAdmin;
    emit FeeAdminSet(feeAdmin);
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
  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external override onlyOwnerOrFeeAdmin {
    _setFeeConfig(feeTokenConfigs);
  }

  /// @dev Set the fee config
  /// @param feeTokenConfigs The fee token configs
  function _setFeeConfig(FeeTokenConfigArgs[] memory feeTokenConfigs) internal {
    for (uint256 i = 0; i < feeTokenConfigs.length; i++) {
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
    for (uint256 i = 0; i < s_nops.length(); i++) {
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
    if (msg.sender != owner() && msg.sender != s_feeAdmin && !s_nops.contains(msg.sender))
      revert OnlyCallableByOwnerOrFeeAdminOrNop();
    _;
  }

  /// @dev Require that the sender is the owner or the fee admin
  modifier onlyOwnerOrFeeAdmin() {
    if (msg.sender != owner() && msg.sender != s_feeAdmin) revert OnlyCallableByOwnerOrFeeAdmin();
    _;
  }
}
