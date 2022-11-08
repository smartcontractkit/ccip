// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {EVM2AnyTollOnRampRouterInterface} from "../../interfaces/onRamp/EVM2AnyTollOnRampRouterInterface.sol";
import {EVM2EVMTollOnRampInterface} from "../../interfaces/onRamp/EVM2EVMTollOnRampInterface.sol";
import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IERC20, PoolInterface} from "../../interfaces/pools/PoolInterface.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {BaseOnRamp} from "../BaseOnRamp.sol";
import {CCIP} from "../../models/Models.sol";

/**
 * @notice An implementation of a toll OnRamp.
 */
contract EVM2EVMTollOnRamp is EVM2EVMTollOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using CCIP for bytes;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMTollOnRamp 1.0.0";

  // Fees per token.
  IERC20[] internal s_feeTokens;
  mapping(IERC20 => uint256) internal s_feesByToken;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    address[] memory allowlist,
    AFNInterface afn,
    OnRampConfig memory config,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    EVM2AnyTollOnRampRouterInterface router
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

  /// @inheritdoc EVM2EVMTollOnRampInterface
  function forwardFromRouter(CCIP.EVM2AnyTollMessage calldata message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    uint256 gasLimit = message.extraArgs._fromBytes().gasLimit;
    _handleForwardFromRouter(message.data.length, gasLimit, message.tokens, message.amounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    CCIP.EVM2EVMTollMessage memory tollMsg = CCIP.EVM2EVMTollMessage({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: i_chainId,
      sender: originalSender,
      receiver: abi.decode(message.receiver, (address)),
      data: message.data,
      tokens: message.tokens,
      amounts: message.amounts,
      feeToken: message.feeToken,
      feeTokenAmount: message.feeTokenAmount,
      gasLimit: gasLimit
    });
    emit CCIPSendRequested(tollMsg);
    return tollMsg.sequenceNumber;
  }

  /// @inheritdoc EVM2EVMTollOnRampInterface
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

  /// @inheritdoc EVM2EVMTollOnRampInterface
  // NOTE: Assumes fee token is valid.
  function getRequiredFee(IERC20 feeToken) external view override returns (uint256) {
    return s_feesByToken[feeToken];
  }
}
