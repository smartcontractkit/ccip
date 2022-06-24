// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../interfaces/TollOnRampInterface.sol";
import "../../utils/CCIP.sol";
import "../BaseOnRamp.sol";

/**
 * @notice An implementation of a toll OnRamp.
 */
contract EVM2EVMTollOnRamp is TollOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  // OnRamp config
  OnRampConfig private s_config;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    OnRampConfig memory config
  ) BaseOnRamp(chainId, destinationChainId, tokens, pools, feeds, allowlist, afn, maxTimeWithoutAFNSignal) {
    s_config = config;
  }

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    address sender = msg.sender;
    if (sender != s_config.router) revert MustBeCalledByRouter();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    // Check that payload is formed correctly
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length)
      revert UnsupportedNumberOfTokens();

    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    for (uint256 i = 0; i < message.tokens.length; i++) {
      IERC20 token = message.tokens[i];
      PoolInterface pool = getPool(token);
      if (address(pool) == address(0)) revert UnsupportedToken(token);
      uint256 amount = message.amounts[i];
      pool.lockOrBurn(amount);
    }

    // Emit message request
    CCIP.EVM2EVMTollEvent memory tollEvent = CCIP.EVM2EVMTollEvent({
      sequenceNumber: s_sequenceNumber++,
      sourceChainId: CHAIN_ID,
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

  /**
   * @notice Get the required fee for a specific fee token
   * @param feeToken token to get the fee for
   * @return fee uint256
   */
  function getRequiredFee(IERC20 feeToken) public view override returns (uint256) {
    AggregatorV2V3Interface feed = getFeed(feeToken);
    if (address(feed) == address(0)) revert UnsupportedFeeToken(feeToken);
    return s_config.relayingFeeJuels * uint256(feed.latestAnswer());
  }

  function setConfig(OnRampConfig calldata config) external onlyOwner {
    s_config = config;
    emit OnRampConfigSet(config);
  }

  function getConfig() external view returns (OnRampConfig memory config) {
    return s_config;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EVM2EVMTollOnRamp 1.0.0";
  }
}
