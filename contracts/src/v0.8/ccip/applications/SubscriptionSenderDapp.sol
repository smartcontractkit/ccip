// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/TypeAndVersionInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../interfaces/onRamp/Any2EVMSubscriptionOnRampRouterInterface.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract SubscriptionSenderDapp is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "SubscriptionSenderDapp 1.0.0";

  // On ramp contract responsible for interacting with the DON.
  Any2EVMSubscriptionOnRampRouterInterface public immutable ON_RAMP_ROUTER;
  uint256 public immutable DESTINATION_CHAIN_ID;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable DESTINATION_CONTRACT;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(
    Any2EVMSubscriptionOnRampRouterInterface onRampRouter,
    uint256 destinationChainId,
    address destinationContract
  ) {
    ON_RAMP_ROUTER = onRampRouter;
    DESTINATION_CHAIN_ID = destinationChainId;
    DESTINATION_CONTRACT = destinationContract;
  }

  /**
   * @notice Send tokens to the destination chain.
   * @dev msg.sender must first call TOKEN.approve for this contract to spend the tokens.
   */
  function sendTokens(
    address destinationAddress,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) external returns (uint64 sequenceNumber) {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    for (uint256 i = 0; i < tokens.length; ++i) {
      tokens[i].safeTransferFrom(msg.sender, address(this), amounts[i]);
      tokens[i].approve(address(ON_RAMP_ROUTER), amounts[i]);
    }
    // `data` format:
    //  - EOA sender address
    //  - EOA destination address
    sequenceNumber = ON_RAMP_ROUTER.ccipSend(
      DESTINATION_CHAIN_ID,
      CCIP.EVM2AnySubscriptionMessage({
        receiver: DESTINATION_CONTRACT,
        data: abi.encode(msg.sender, destinationAddress),
        tokens: tokens,
        amounts: amounts,
        gasLimit: 0
      })
    );
  }
}
