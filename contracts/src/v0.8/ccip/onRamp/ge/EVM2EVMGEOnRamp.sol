// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IERC20, PoolInterface} from "../../interfaces/pools/PoolInterface.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {BaseOnRamp} from "../BaseOnRamp.sol";
import {Common} from "../../models/Common.sol";
import {GEConsumer} from "../../models/GEConsumer.sol";
import {GE} from "../../models/GE.sol";
import {EVM2EVMGEOnRampInterface} from "../../interfaces/onRamp/EVM2EVMGEOnRampInterface.sol";
import {GERouterInterface} from "../../interfaces/router/GERouterInterface.sol";
import {GasFeeCacheInterface} from "../../interfaces/gasFeeCache/GasFeeCacheInterface.sol";

contract EVM2EVMGEOnRamp is EVM2EVMGEOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMGEOnRamp 1.0.0";
  uint256 private constant EVM_DEFAULT_GAS_LIMIT = 200_000;

  bytes32 internal immutable i_metadataHash;

  mapping(address => uint64) internal s_nonceBySender;
  DynamicFeeConfig internal s_feeConfig;
  address internal s_feeAdmin;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    address[] memory allowlist,
    AFNInterface afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    GERouterInterface router,
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

  /// @inheritdoc EVM2EVMGEOnRampInterface
  function forwardFromRouter(
    GEConsumer.EVM2AnyGEMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external override whenNotPaused whenHealthy returns (bytes32) {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();

    GEConsumer.EVMExtraArgsV1 memory extraArgs = _fromBytes(message.extraArgs);
    _handleForwardFromRouter(message.data.length, extraArgs.gasLimit, message.tokensAndAmounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    GE.EVM2EVMGEMessage memory GEMsg = GE.EVM2EVMGEMessage({
      sourceChainId: i_chainId,
      sequenceNumber: ++s_sequenceNumber,
      feeTokenAmount: feeTokenAmount,
      sender: originalSender,
      nonce: ++s_nonceBySender[originalSender],
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

  function getFee(GEConsumer.EVM2AnyGEMessage calldata message) public view returns (uint256 fee) {
    if (s_feeConfig.feeToken != message.feeToken) revert MismatchedFeeToken(s_feeConfig.feeToken, message.feeToken);
    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    uint256 linkPerUnitGas = GasFeeCacheInterface(s_feeConfig.gasFeeCache).getFee(s_feeConfig.destChainId);

    return
      s_feeConfig.feeAmount + // Flat fee
      ((gasLimit + s_feeConfig.destGasOverhead) * linkPerUnitGas * s_feeConfig.multiplier) / // Total gas reserved for tx
      1 ether; // latest gas reported gas fee with a safety margin
  }

  function setFeeAdmin(address feeAdmin) external onlyOwner {
    s_feeAdmin = feeAdmin;
    emit FeeAdminSet(feeAdmin);
  }

  function setFeeConfig(DynamicFeeConfig calldata feeConfig) external onlyOwner {
    s_feeConfig = feeConfig;
    emit FeeConfigSet(feeConfig);
  }
}
