// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../../interfaces/pools/IPool.sol";
import {IAFN} from "../../interfaces/health/IAFN.sol";
import {IEVM2AnyGEOnRamp} from "../../interfaces/onRamp/IEVM2AnyGEOnRamp.sol";
import {IEVM2EVMGEOnRamp} from "../../interfaces/onRamp/IEVM2EVMGEOnRamp.sol";
import {IGERouter} from "../../interfaces/router/IGERouter.sol";
import {IFeeManager} from "../../interfaces/fees/IFeeManager.sol";

import {BaseOnRamp} from "../BaseOnRamp.sol";
import {Common} from "../../models/Common.sol";
import {GEConsumer} from "../../models/GEConsumer.sol";
import {GE} from "../../models/GE.sol";

import {SafeERC20} from "../../../vendor/SafeERC20.sol";
import {IERC20} from "../../../vendor/IERC20.sol";

contract EVM2EVMGEOnRamp is IEVM2EVMGEOnRamp, BaseOnRamp, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMGEOnRamp 1.0.0";
  uint256 private constant EVM_DEFAULT_GAS_LIMIT = 200_000;

  bytes32 internal immutable i_metadataHash;
  address internal s_feeAdmin;

  mapping(address => uint64) internal s_senderNonce;

  address internal s_feeManager;

  mapping(address => FeeTokenConfig) internal s_feeTokenConfig;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    address[] memory tokens,
    IPool[] memory pools,
    address[] memory allowlist,
    IAFN afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    IGERouter router,
    address feeManager,
    FeeTokenConfigArgs[] memory feeTokenConfigs
  ) BaseOnRamp(chainId, destinationChainId, tokens, pools, allowlist, afn, config, rateLimiterConfig, address(router)) {
    i_metadataHash = keccak256(abi.encode(GE.EVM_2_EVM_GE_MESSAGE_HASH, chainId, destinationChainId, address(this)));
    s_feeManager = feeManager;
    _setFeeConfig(feeTokenConfigs);
  }

  function _fromBytes(bytes calldata extraArgs) internal pure returns (GEConsumer.EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return GEConsumer.EVMExtraArgsV1({gasLimit: EVM_DEFAULT_GAS_LIMIT, strict: false});
    }
    if (bytes4(extraArgs[:4]) != GEConsumer.EVM_EXTRA_ARGS_V1_TAG)
      revert InvalidExtraArgsTag(GEConsumer.EVM_EXTRA_ARGS_V1_TAG, bytes4(extraArgs[:4]));
    return GEConsumer.EVMExtraArgsV1({gasLimit: abi.decode(extraArgs[4:36], (uint256)), strict: false});
  }

  /// @inheritdoc IEVM2AnyGEOnRamp
  function getPoolBySourceToken(IERC20 sourceToken)
    public
    view
    virtual
    override(BaseOnRamp, IEVM2AnyGEOnRamp)
    returns (IPool)
  {
    return BaseOnRamp.getPoolBySourceToken(sourceToken);
  }

  /// @inheritdoc IEVM2AnyGEOnRamp
  function getSupportedTokens()
    public
    view
    virtual
    override(BaseOnRamp, IEVM2AnyGEOnRamp)
    returns (address[] memory tokens)
  {
    return BaseOnRamp.getSupportedTokens();
  }

  /// @inheritdoc IEVM2AnyGEOnRamp
  function forwardFromRouter(
    GEConsumer.EVM2AnyGEMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external override whenNotPaused whenHealthy returns (bytes32) {
    GEConsumer.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    // Validate the message with various checks
    _validateMessage(message.data.length, extraArgs.gasLimit, message.tokensAndAmounts, originalSender);

    // Send feeToken directly to the Fee Manager
    IERC20(message.feeToken).safeTransfer(address(s_feeManager), feeTokenAmount);

    for (uint256 i = 0; i < message.tokensAndAmounts.length; ++i) {
      Common.EVMTokenAndAmount memory tokenAndAmount = message.tokensAndAmounts[i];
      IERC20 token = IERC20(tokenAndAmount.token);
      IPool pool = getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      pool.lockOrBurn(tokenAndAmount.amount);
    }

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    GE.EVM2EVMGEMessage memory GEMsg = GE.EVM2EVMGEMessage({
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
    GEMsg.messageId = GE._hash(GEMsg, i_metadataHash);
    emit CCIPSendRequested(GEMsg);
    return GEMsg.messageId;
  }

  /// @inheritdoc IEVM2AnyGEOnRamp
  function getFee(GEConsumer.EVM2AnyGEMessage calldata message) public view override returns (uint256 fee) {
    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    uint256 feeTokenBaseUnitsPerUnitGas = IFeeManager(s_feeManager).getFeeTokenBaseUnitsPerUnitGas(
      message.feeToken,
      i_destinationChainId
    );
    if (feeTokenBaseUnitsPerUnitGas == 0)
      revert IFeeManager.TokenOrChainNotSupported(message.feeToken, i_destinationChainId);

    FeeTokenConfig memory feeTokenConfig = s_feeTokenConfig[message.feeToken];
    return
      feeTokenConfig.feeAmount + // Flat fee
      ((gasLimit + feeTokenConfig.destGasOverhead) * feeTokenBaseUnitsPerUnitGas * feeTokenConfig.multiplier) / // Total gas reserved for tx
      1 ether; // latest gas reported gas fee with a safety margin
  }

  /// @inheritdoc IEVM2AnyGEOnRamp
  function getSenderNonce(address sender) external view override returns (uint64) {
    return s_senderNonce[sender];
  }

  /// @inheritdoc IEVM2EVMGEOnRamp
  function setFeeAdmin(address feeAdmin) external override onlyOwner {
    s_feeAdmin = feeAdmin;
    emit FeeAdminSet(feeAdmin);
  }

  /// @inheritdoc IEVM2EVMGEOnRamp
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
