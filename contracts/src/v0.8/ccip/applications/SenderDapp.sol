// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../interfaces/TypeAndVersionInterface.sol";
import "../utils/CCIP.sol";
import "../../vendor/SafeERC20.sol";
import "../onRamp/interfaces/TollOnRampRouterInterface.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract SenderDapp is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // On ramp contract responsible for interacting with the DON.
  TollOnRampRouterInterface public immutable ON_RAMP_ROUTER;
  uint256 public immutable DESTINATION_CHAIN_ID;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable DESTINATION_CONTRACT;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(
    TollOnRampRouterInterface onRampRouter,
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
    uint256[] memory amounts,
    address executor
  ) external returns (uint64 sequenceNumber) {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    address originalSender = msg.sender;
    for (uint256 i = 0; i < tokens.length; i++) {
      tokens[i].safeTransferFrom(originalSender, address(this), amounts[i]);
      tokens[i].approve(address(ON_RAMP_ROUTER), amounts[i]);
    }
    // `data` format:
    //  - EOA sender address
    //  - EOA destination address
    sequenceNumber = ON_RAMP_ROUTER.ccipSend(
      DESTINATION_CHAIN_ID,
      CCIP.EVM2AnyTollMessage({
        receiver: destinationAddress,
        data: abi.encode(originalSender, destinationAddress),
        tokens: tokens,
        amounts: amounts,
        feeToken: tokens[0],
        feeTokenAmount: 0,
        gasLimit: 0
      })
    );
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "SenderDapp 1.0.0";
  }
}
