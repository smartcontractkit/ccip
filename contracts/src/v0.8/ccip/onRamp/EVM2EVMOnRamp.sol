// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IEVM2EVMOnRamp} from "../interfaces/onRamp/IEVM2EVMOnRamp.sol";
import {IEVM2AnyOnRamp} from "../interfaces/onRamp/IEVM2AnyOnRamp.sol";
import {IAggregateRateLimiter} from "../interfaces/rateLimiter/IAggregateRateLimiter.sol";
import {IFeeManager} from "../interfaces/fees/IFeeManager.sol";

import {HealthChecker} from "../health/HealthChecker.sol";
import {AllowList} from "../access/AllowList.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {Common} from "../models/Common.sol";
import {Client} from "../models/Client.sol";
import {Internal} from "../models/Internal.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";

contract EVM2EVMOnRamp is IEVM2EVMOnRamp, HealthChecker, AllowList, AggregateRateLimiter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOnRamp 1.0.0";

  // Static configuration
  uint256 private constant EVM_DEFAULT_GAS_LIMIT = 200_000;
  bytes32 internal immutable i_metadataHash;
  // Chain ID of the source chain (where this contract is deployed)
  uint64 internal immutable i_chainId;
  // Chain ID of the destination chain (where this contract sends messages)
  uint64 internal immutable i_destinationChainId;

  // Dynamic configuration
  OnRampConfig internal s_config;
  // TODO: these three should be part of config?
  address internal s_feeAdmin;
  address internal s_feeManager;
  address internal s_router;

  // State
  mapping(address => uint64) internal s_senderNonce;
  /// @dev Struct to hold the fee configuration for a token
  struct FeeTokenConfig {
    uint96 feeAmount; // ---------┐ Flat fee
    uint64 multiplier; //         | Price multiplier for gas costs
    uint32 destGasOverhead; // ---┘ Extra gas charged on top of the gasLimit
  }
  mapping(address => FeeTokenConfig) internal s_feeTokenConfig;
  // The last used sequence number. This is zero in the case where no
  // messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  uint64 internal s_sequenceNumber;
  struct PoolConfig {
    IPool pool;
    bool enabled;
  }
  // source token => token pool
  mapping(IERC20 => PoolConfig) private s_poolsBySourceToken;
  address[] private s_sourceTokenList;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    address[] memory tokens,
    IPool[] memory pools,
    address[] memory allowlist,
    IAFN afn,
    OnRampConfig memory config,
    IAggregateRateLimiter.RateLimiterConfig memory rateLimiterConfig,
    address router,
    address feeManager,
    FeeTokenConfigArgs[] memory feeTokenConfigs
  ) HealthChecker(afn) AllowList(allowlist) AggregateRateLimiter(rateLimiterConfig) {
    i_metadataHash = keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, chainId, destinationChainId, address(this)));
    s_feeManager = feeManager;
    _setFeeConfig(feeTokenConfigs);
    i_chainId = chainId;
    i_destinationChainId = destinationChainId;
    s_config = config;
    s_router = router;
    s_sequenceNumber = 0;

    if (tokens.length != pools.length) revert InvalidTokenPoolConfig();
    s_sourceTokenList = tokens;
    // Set new tokens and pools
    for (uint256 i = 0; i < tokens.length; ++i) {
      s_poolsBySourceToken[IERC20(tokens[i])] = PoolConfig({pool: pools[i], enabled: true});
    }
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function getPoolBySourceToken(IERC20 sourceToken) external view override returns (IPool) {
    return _getPoolBySourceToken(sourceToken);
  }

  function _getPoolBySourceToken(IERC20 sourceToken) private view returns (IPool) {
    PoolConfig memory poolConfig = s_poolsBySourceToken[sourceToken];
    if (poolConfig.enabled) {
      return s_poolsBySourceToken[sourceToken].pool;
    }
    revert UnsupportedToken(sourceToken);
  }

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
  /// @param tokensAndAmounts The token payload to be sent. They will be locked into pools by this function.
  /// @param originalSender The original sender of the message on the router.
  function _validateMessage(
    uint256 dataLength,
    uint256 gasLimit,
    Common.EVMTokenAndAmount[] memory tokensAndAmounts,
    address originalSender
  ) internal {
    if (msg.sender != s_router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (dataLength > uint256(s_config.maxDataSize)) revert MessageTooLarge(uint256(s_config.maxDataSize), dataLength);
    if (gasLimit > uint256(s_config.maxGasLimit)) revert MessageGasLimitTooHigh();
    if (tokensAndAmounts.length > uint256(s_config.maxTokensLength)) revert UnsupportedNumberOfTokens();
    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    _removeTokens(tokensAndAmounts);
  }

  /// @inheritdoc  IEVM2AnyOnRamp
  function forwardFromRouter(
    Client.EVM2AnyMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external override whenNotPaused whenHealthy returns (bytes32) {
    Client.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    // Validate the message with various checks
    _validateMessage(message.data.length, extraArgs.gasLimit, message.tokensAndAmounts, originalSender);

    // Send feeToken directly to the Fee Manager
    IERC20(message.feeToken).safeTransfer(address(s_feeManager), feeTokenAmount);

    for (uint256 i = 0; i < message.tokensAndAmounts.length; ++i) {
      Common.EVMTokenAndAmount memory tokenAndAmount = message.tokensAndAmounts[i];
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
      tokensAndAmounts: message.tokensAndAmounts,
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
    uint256 feeTokenBaseUnitsPerUnitGas = IFeeManager(s_feeManager).getFeeTokenBaseUnitsPerUnitGas(
      message.feeToken,
      i_destinationChainId
    );
    if (feeTokenBaseUnitsPerUnitGas == 0)
      revert IFeeManager.TokenOrChainNotSupported(message.feeToken, i_destinationChainId);

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
  function getDestinationChainId() external view override returns (uint64) {
    return i_destinationChainId;
  }

  /**
   * @notice Add a new token pool
   * @param token The source token
   * @param pool The pool that will be used
   * @dev This method can only be called by the owner of the contract.
   */
  function addPool(IERC20 token, IPool pool) public onlyOwner {
    if (address(token) == address(0) || address(pool) == address(0)) revert InvalidTokenPoolConfig();
    if (s_poolsBySourceToken[token].enabled) revert PoolAlreadyAdded();

    s_poolsBySourceToken[token] = PoolConfig({pool: pool, enabled: true});
    s_sourceTokenList.push(address(token));

    emit PoolAdded(token, pool);
  }

  /**
   * @notice Remove a token pool
   * @param token The source token
   * @param pool The pool that will be removed
   * @dev This method can only be called by the owner of the contract.
   */
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
  function setFeeAdmin(address feeAdmin) external override onlyOwner {
    s_feeAdmin = feeAdmin;
    emit FeeAdminSet(feeAdmin);
  }

  /// @inheritdoc IEVM2EVMOnRamp
  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external override onlyOwnerOrFeeAdmin {
    _setFeeConfig(feeTokenConfigs);
  }

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

  modifier onlyOwnerOrFeeAdmin() {
    if (msg.sender != owner() && msg.sender != s_feeAdmin) revert OnlyCallableByOwnerOrFeeAdmin();
    _;
  }
}
