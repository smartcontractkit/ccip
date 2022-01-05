// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../ramps/SingleTokenOnRamp.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../../vendor/SafeERC20.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract EOASingleTokenSender is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // On ramp contract responsible for interacting with the DON.
  SingleTokenOnRamp public immutable ON_RAMP;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable DESTINATION_CONTRACT;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(SingleTokenOnRamp onRamp, address destinationContract) {
    ON_RAMP = onRamp;
    DESTINATION_CONTRACT = destinationContract;
  }

  /**
   * @notice Send tokens to the destination chain.
   * @dev msg.sender must first call TOKEN.approve for this contract to spend the tokens.
   */
  function sendTokens(
    address destinationAddress,
    uint256 amount,
    address executor
  ) external returns (uint256 sequenceNumber) {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    bytes memory options;
    // Set tokens using the ramp token
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = ON_RAMP.TOKEN();
    // Set the amounts using the amount parameter
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = amount;
    address originalSender = msg.sender;
    // Init the MessagePayload struct
    // `payload.data` format:
    //  - EOA sender address
    //  - EOA destination address
    CCIP.MessagePayload memory payload = CCIP.MessagePayload({
      receiver: DESTINATION_CONTRACT,
      data: abi.encode(originalSender, destinationAddress),
      tokens: tokens,
      amounts: amounts,
      executor: executor,
      options: options
    });
    tokens[0].safeTransferFrom(originalSender, address(this), amount);
    tokens[0].approve(address(ON_RAMP.POOL()), amount);
    sequenceNumber = ON_RAMP.requestCrossChainSend(payload);
  }

  /**
   * @notice Get the details of the ramp. This includes the destination chain details
   */
  function rampDetails()
    external
    view
    returns (
      IERC20 token,
      uint256 destinationChainId,
      IERC20 destinationChainToken
    )
  {
    token = ON_RAMP.TOKEN();
    destinationChainId = ON_RAMP.DESTINATION_CHAIN_ID();
    destinationChainToken = ON_RAMP.DESTINATION_TOKEN();
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "EOASingleTokenSender 1.0.0";
  }
}
