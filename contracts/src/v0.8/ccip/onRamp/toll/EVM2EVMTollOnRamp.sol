// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/onRamp/Any2EVMTollOnRampRouterInterface.sol";
import "../../../interfaces/TypeAndVersionInterface.sol";
import "../BaseOnRamp.sol";

/**
 * @notice An implementation of a toll OnRamp.
 */
contract EVM2EVMTollOnRamp is Any2EVMTollOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using Address for address;
  using SafeERC20 for IERC20;

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
    Any2EVMTollOnRampRouterInterface router
  ) BaseOnRamp(chainId, destinationChainId, tokens, pools, allowlist, afn, config, address(router)) {}

  /// @inheritdoc Any2EVMTollOnRampInterface
  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    handleForwardFromRouter(message.data.length, message.tokens, message.amounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    CCIP.EVM2EVMTollEvent memory tollEvent = CCIP.EVM2EVMTollEvent({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: i_chainId,
      sender: originalSender,
      receiver: message.receiver,
      data: message.data,
      tokens: message.tokens,
      amounts: message.amounts,
      feeToken: message.feeToken,
      feeTokenAmount: message.feeTokenAmount,
      gasLimit: message.gasLimit
    });
    emit CCIPSendRequested(tollEvent);
    return tollEvent.sequenceNumber;
  }

  /// @inheritdoc Any2EVMTollOnRampInterface
  // If the fee is not explicitly set, we use the solidity default of zero.
  // The set of tokens in the pool registry defines the whitelist of fee tokens.
  function setFeeConfig(FeeConfig memory feeConfig) external override onlyOwner {
    if (feeConfig.fees.length != feeConfig.feeTokens.length) {
      revert InvalidFeeConfig();
    }
    // Clear previously set fees.
    for (uint256 i = 0; i < s_feeTokens.length; i++) {
      delete s_feesByToken[s_feeTokens[i]];
    }
    // Set new fees
    for (uint256 i = 0; i < feeConfig.feeTokens.length; i++) {
      if (address(feeConfig.feeTokens[i]) == address(0)) {
        revert InvalidFeeConfig();
      }
      s_feesByToken[feeConfig.feeTokens[i]] = feeConfig.fees[i];
    }
    s_feeTokens = feeConfig.feeTokens;
  }

  /// @inheritdoc Any2EVMTollOnRampInterface
  // NOTE: Assumes fee token is valid.
  function getRequiredFee(IERC20 feeToken) external view override returns (uint256) {
    return s_feesByToken[feeToken];
  }
}
