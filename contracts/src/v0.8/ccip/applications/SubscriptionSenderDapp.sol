// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/TypeAndVersionInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../interfaces/onRamp/EVM2AnySubscriptionOnRampRouterInterface.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract SubscriptionSenderDapp is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "SubscriptionSenderDapp 1.0.0";

  // On ramp contract responsible for interacting with the DON.
  EVM2AnySubscriptionOnRampRouterInterface public immutable i_onRampRouter;
  uint256 public immutable i_destinationChainId;

  constructor(EVM2AnySubscriptionOnRampRouterInterface onRampRouter, uint256 destinationChainId) {
    i_onRampRouter = onRampRouter;
    i_destinationChainId = destinationChainId;
  }

  /**
   * @notice Send messages to the destination chain.
   * @dev msg.sender must first call TOKEN.approve for this contract to spend the tokens.
   */
  function sendMessage(CCIP.EVM2AnySubscriptionMessage memory message) external returns (uint64 sequenceNumber) {
    IERC20[] memory tokens = message.tokens;
    uint256[] memory amounts = message.amounts;
    for (uint256 i = 0; i < tokens.length; ++i) {
      tokens[i].safeTransferFrom(msg.sender, address(this), amounts[i]);
      tokens[i].approve(address(i_onRampRouter), amounts[i]);
    }
    sequenceNumber = i_onRampRouter.ccipSend(i_destinationChainId, message);
  }

  /**
   * @notice Funds the subscription for On ramp.
   * @dev msg.sender must first call TOKEN.approve and transfer for this contract to spend the tokens.
   */
  function fundSubscription(IERC20 feeToken, uint256 amount) external {
    feeToken.approve(address(i_onRampRouter), amount);
    i_onRampRouter.fundSubscription(amount);
  }

  function unfundSubscription(uint256 amount) external {
    i_onRampRouter.unfundSubscription(amount);
  }
}
