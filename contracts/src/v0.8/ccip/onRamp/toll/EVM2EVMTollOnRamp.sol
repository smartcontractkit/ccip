// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IEVM2AnyTollOnRampRouter} from "../../interfaces/onRamp/IEVM2AnyTollOnRampRouter.sol";
import {IEVM2EVMTollOnRamp} from "../../interfaces/onRamp/IEVM2EVMTollOnRamp.sol";
import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IERC20, IPool} from "../../interfaces/pools/IPool.sol";
import {IAFN} from "../../interfaces/health/IAFN.sol";
import {BaseOnRamp} from "../BaseOnRamp.sol";
import {Common} from "../../models/Common.sol";
import {Toll} from "../../models/Toll.sol";
import {TollConsumer} from "../../models/TollConsumer.sol";

/**
 * @notice An implementation of a toll OnRamp.
 */
contract EVM2EVMTollOnRamp is IEVM2EVMTollOnRamp, BaseOnRamp, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMTollOnRamp 1.0.0";
  uint256 private constant EVM_DEFAULT_GAS_LIMIT = 200_000;

  // Fees per token.
  IERC20[] internal s_feeTokens;
  mapping(IERC20 => uint256) internal s_feesByToken;

  constructor(
    uint64 chainId,
    uint64 destinationChainId,
    IERC20[] memory tokens,
    IPool[] memory pools,
    address[] memory allowlist,
    IAFN afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    IEVM2AnyTollOnRampRouter router
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
  {}

  function _fromBytes(bytes calldata extraArgs) internal pure returns (TollConsumer.EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return TollConsumer.EVMExtraArgsV1({gasLimit: EVM_DEFAULT_GAS_LIMIT, strict: false});
    }
    if (bytes4(extraArgs[:4]) != TollConsumer.EVM_EXTRA_ARGS_V1_TAG)
      revert InvalidExtraArgsTag(TollConsumer.EVM_EXTRA_ARGS_V1_TAG, bytes4(extraArgs[:4]));
    return TollConsumer.EVMExtraArgsV1({gasLimit: abi.decode(extraArgs[4:36], (uint256)), strict: false});
  }

  /// @inheritdoc IEVM2EVMTollOnRamp
  function forwardFromRouter(TollConsumer.EVM2AnyTollMessage calldata message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    uint256 gasLimit = _fromBytes(message.extraArgs).gasLimit;
    _handleForwardFromRouter(message.data.length, gasLimit, message.tokensAndAmounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    Toll.EVM2EVMTollMessage memory tollMsg = Toll.EVM2EVMTollMessage({
      sourceChainId: i_chainId,
      sequenceNumber: ++s_sequenceNumber,
      sender: originalSender,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokensAndAmounts: message.tokensAndAmounts,
      feeTokenAndAmount: message.feeTokenAndAmount,
      gasLimit: gasLimit
    });
    emit CCIPSendRequested(tollMsg);
    return tollMsg.sequenceNumber;
  }

  /// @inheritdoc IEVM2EVMTollOnRamp
  // If the fee is not explicitly set, we use the solidity default of zero.
  // The set of tokens in the pool registry defines the whitelist of fee tokens.
  function setFeeConfig(FeeConfig memory feeConfig) external override onlyOwner {
    if (feeConfig.fees.length != feeConfig.feeTokens.length) revert InvalidFeeConfig();
    // Clear previously set fees.
    for (uint256 i = 0; i < s_feeTokens.length; i++) {
      delete s_feesByToken[s_feeTokens[i]];
    }
    // Set new fees
    for (uint256 i = 0; i < feeConfig.feeTokens.length; i++) {
      if (address(feeConfig.feeTokens[i]) == address(0)) revert InvalidFeeConfig();
      s_feesByToken[feeConfig.feeTokens[i]] = feeConfig.fees[i];
    }
    s_feeTokens = feeConfig.feeTokens;
  }

  /// @inheritdoc IEVM2EVMTollOnRamp
  // NOTE: Assumes fee token is valid.
  function getRequiredFee(IERC20 feeToken) external view override returns (uint256) {
    return s_feesByToken[feeToken];
  }
}
