// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IERC20, PoolInterface} from "../../interfaces/pools/PoolInterface.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {BaseOnRamp} from "../BaseOnRamp.sol";
import {CCIP} from "../../models/Models.sol";
import {EVM2EVMGEOnRampInterface} from "../../interfaces/onRamp/EVM2EVMGEOnRampInterface.sol";
import {GERouterInterface} from "../../interfaces/router/GERouterInterface.sol";
import {DynamicFeeCalculator} from "../../dynamicFeeCalculator/DynamicFeeCalculator.sol";

contract EVM2EVMGEOnRamp is EVM2EVMGEOnRampInterface, BaseOnRamp, DynamicFeeCalculator, TypeAndVersionInterface {
  using CCIP for bytes;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMGEOnRamp 1.0.0";

  // Fees per token.
  IERC20[] internal s_feeTokens;
  mapping(IERC20 => uint256) internal s_feesByToken;
  mapping(address => uint64) internal s_nonceBySender;

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
    DynamicFeeCalculator(chainId, feeConfig)
  {}

  /// @inheritdoc EVM2EVMGEOnRampInterface
  function forwardFromRouter(
    CCIP.EVM2AnyGEMessage calldata message,
    uint256 feeTokenAmount,
    address originalSender
  ) external override whenNotPaused whenHealthy returns (uint64) {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();

    CCIP.EVMExtraArgsV1 memory extraArgs = message.extraArgs._fromBytes();
    _handleForwardFromRouter(message.data.length, extraArgs.gasLimit, message.tokensAndAmounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    CCIP.EVM2EVMGEMessage memory GEMsg = CCIP.EVM2EVMGEMessage({
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
      feeToken: message.feeToken
    });
    emit CCIPSendRequested(GEMsg);
    return GEMsg.sequenceNumber;
  }
}
