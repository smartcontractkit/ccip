// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "./EVM2AnySubscriptionOnRampRouter.sol";
import "../BaseOnRamp.sol";

/**
 * @notice An implementation of a subscription OnRamp.
 */
contract EVM2EVMSubscriptionOnRamp is Any2EVMSubscriptionOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  string public constant override typeAndVersion = "EVM2EVMSubscriptionOnRamp 1.0.0";

  // The last used sequence number per receiver address. This is zero in the case
  // where no messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  mapping(address => uint64) internal s_receiverToNonce;

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
    Any2EVMSubscriptionOnRampRouterInterface router
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

  /**
   * @notice Send a message to the remote chain
   * @dev approve() must have already been called on the token using the this ramp address as the spender.
   * @dev if the contract is paused, this function will revert.
   * @param message Message struct to send
   * @param originalSender The original initiator of the CCIP request
   */
  function forwardFromRouter(CCIP.EVM2AnySubscriptionMessage memory message, address originalSender)
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
    // we need the next nonce so we increment before we use the value
    CCIP.EVM2EVMSubscriptionMessage memory subscriptionMsg = CCIP.EVM2EVMSubscriptionMessage({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: i_chainId,
      sender: originalSender,
      receiver: message.receiver,
      nonce: ++s_receiverToNonce[message.receiver],
      data: message.data,
      tokens: message.tokens,
      amounts: message.amounts,
      gasLimit: message.gasLimit
    });
    emit CCIPSendRequested(subscriptionMsg);
    return subscriptionMsg.sequenceNumber;
  }
}
