// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../utils/CCIP.sol";
import "./EVM2AnyMOOnRampRouter.sol";
import "../BaseOnRamp.sol";
import "../interfaces/Any2EVMMOOnRampInterface.sol";

/**
 * @notice An implementation of a subscription OnRamp.
 */
contract EVM2EVMMOOnRamp is Any2EVMMOOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2EVMMOOnRamp 1.0.0";

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
    Any2EVMMOOnRampRouterInterface router
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
  function forwardFromRouter(CCIP.EVM2AnyMOMessage memory message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    if (s_router == address(0)) revert RouterNotSet();
    if (originalSender == address(0)) revert RouterMustSetOriginalSender();
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
    if (s_allowlistEnabled && !s_allowed[originalSender]) revert SenderNotAllowed(originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    // we need the next nonce so we increment before we use the value
    CCIP.EVM2EVMMOEvent memory subscriptionEvent = CCIP.EVM2EVMMOEvent({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: CHAIN_ID,
      sender: originalSender,
      receiver: message.receiver,
      nonce: ++s_receiverToNonce[message.receiver],
      data: message.data,
      gasLimit: message.gasLimit
    });
    emit CCIPSendRequested(subscriptionEvent);
    return subscriptionEvent.sequenceNumber;
  }
}
