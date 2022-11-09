// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {EVM2EVMSubscriptionOnRampInterface} from "../../interfaces/onRamp/EVM2EVMSubscriptionOnRampInterface.sol";
import {EVM2AnySubscriptionOnRampRouterInterface} from "../../interfaces/onRamp/EVM2AnySubscriptionOnRampRouterInterface.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {EVM2AnySubscriptionOnRampRouter} from "./EVM2AnySubscriptionOnRampRouter.sol";
import {BaseOnRamp} from "../BaseOnRamp.sol";
import {CCIP} from "../../models/Models.sol";
import {IERC20} from "../../../vendor/IERC20.sol";
import {PoolInterface} from "../../interfaces/pools/PoolInterface.sol";

/**
 * @notice An implementation of a subscription OnRamp.
 */
contract EVM2EVMSubscriptionOnRamp is EVM2EVMSubscriptionOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using CCIP for bytes;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
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
    EVM2AnySubscriptionOnRampRouterInterface router
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
  function forwardFromRouter(CCIP.EVM2AnySubscriptionMessage calldata message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    uint256 gasLimit = message.extraArgs._fromBytes().gasLimit;
    _handleForwardFromRouter(message.data.length, gasLimit, message.tokensAndAmounts, originalSender);

    address receiver = abi.decode(message.receiver, (address));
    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    // we need the next nonce so we increment before we use the value
    CCIP.EVM2EVMSubscriptionMessage memory subscriptionMsg = CCIP.EVM2EVMSubscriptionMessage({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: i_chainId,
      sender: originalSender,
      receiver: receiver,
      nonce: ++s_receiverToNonce[receiver],
      data: message.data,
      tokensAndAmounts: message.tokensAndAmounts,
      gasLimit: gasLimit
    });
    emit CCIPSendRequested(subscriptionMsg);
    return subscriptionMsg.sequenceNumber;
  }
}
