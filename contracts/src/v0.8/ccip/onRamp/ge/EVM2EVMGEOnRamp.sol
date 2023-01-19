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

  mapping(address => uint64) internal s_senderNonce;
  DynamicFeeConfig internal s_feeConfig;
  address internal s_feeAdmin;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    address[] memory tokens,
    IPool[] memory pools,
    address[] memory allowlist,
    IAFN afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    IGERouter router,
    DynamicFeeConfig memory feeConfig
  )
    BaseOnRamp(
      chainId,
      destinationChainId,
      tokens,
      pools,
      allowlist,
      afn,
      config,
      rateLimiterConfig,
      tokenLimitsAdmin,
      address(router)
    )
  {
    s_feeConfig = feeConfig;
    i_metadataHash = keccak256(abi.encode(GE.EVM_2_EVM_GE_MESSAGE_HASH, chainId, destinationChainId, address(this)));
    emit FeeConfigSet(feeConfig);
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
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();

    // If link is used as fee token send it to the link pool
    // If a non-link token is used send it to the feeManager contract to
    // convert it to link.
    if (message.feeToken == s_feeConfig.linkToken) {
      address pool = address(getPoolBySourceToken(IERC20(message.feeToken)));
      if (pool == address(0)) revert UnsupportedToken(IERC20(message.feeToken));
      IERC20(message.feeToken).safeTransfer(pool, feeTokenAmount);
    } else {
      IERC20(message.feeToken).safeTransfer(address(s_feeConfig.feeManager), feeTokenAmount);
    }

    GEConsumer.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    _handleForwardFromRouter(message.data.length, extraArgs.gasLimit, message.tokensAndAmounts, originalSender);

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
    uint256 linkPerUnitGas = IFeeManager(s_feeConfig.feeManager).getFee(message.feeToken, s_feeConfig.destChainId);
    if (linkPerUnitGas == 0) revert IFeeManager.TokenOrChainNotSupported(message.feeToken, s_feeConfig.destChainId);

    return
      s_feeConfig.feeAmount + // Flat fee
      ((gasLimit + s_feeConfig.destGasOverhead) * linkPerUnitGas * s_feeConfig.multiplier) / // Total gas reserved for tx
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
  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external override onlyOwner {
    s_feeConfig = feeConfig;
    emit FeeConfigSet(feeConfig);
  }
}
