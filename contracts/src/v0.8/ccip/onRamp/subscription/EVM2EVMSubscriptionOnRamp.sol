// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "./EVM2AnySubscriptionOnRampRouter.sol";
import "../BaseOnRamp.sol";

/**
 * @notice An implementation of a subscription OnRamp.
 */
contract EVM2EVMSubscriptionOnRamp is Any2EVMSubscriptionOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2EVMSubscriptionOnRamp 1.0.0";

  // The last used sequence number per receiver address. This is zero in the case
  // where no messages has been sent yet. 0 is not a valid sequence number for any
  // real transaction.
  mapping(address => uint64) s_receiverToNonce;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    OnRampConfig memory config,
    Any2EVMSubscriptionOnRampRouterInterface router
  )
    BaseOnRamp(
      chainId,
      destinationChainId,
      tokens,
      pools,
      feeds,
      allowlist,
      afn,
      maxTimeWithoutAFNSignal,
      config,
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
    CCIP.EVM2EVMSubscriptionEvent memory subscriptionEvent = CCIP.EVM2EVMSubscriptionEvent({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: CHAIN_ID,
      sender: originalSender,
      receiver: message.receiver,
      nonce: ++s_receiverToNonce[message.receiver],
      data: message.data,
      tokens: message.tokens,
      amounts: message.amounts,
      gasLimit: message.gasLimit
    });
    emit CCIPSendRequested(subscriptionEvent);
    return subscriptionEvent.sequenceNumber;
  }
}
