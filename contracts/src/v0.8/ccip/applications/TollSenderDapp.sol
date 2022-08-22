// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/TypeAndVersionInterface.sol";
import "../../vendor/SafeERC20.sol";
import "../interfaces/onRamp/Any2EVMTollOnRampRouterInterface.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract TollSenderDapp is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "TollSenderDapp 1.0.0";

  // On ramp contract responsible for interacting with the DON.
  Any2EVMTollOnRampRouterInterface public immutable i_onRampRouter;
  uint256 public immutable i_destinationChainId;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable i_destinationContract;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(
    Any2EVMTollOnRampRouterInterface onRampRouter,
    uint256 destinationChainId,
    address destinationContract
  ) {
    i_onRampRouter = onRampRouter;
    i_destinationChainId = destinationChainId;
    i_destinationContract = destinationContract;
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
      tokens[i].approve(address(i_onRampRouter), amounts[i]);
    }
    // `data` format:
    //  - EOA sender address
    //  - EOA destination address
    sequenceNumber = i_onRampRouter.ccipSend(
      i_destinationChainId,
      CCIP.EVM2AnyTollMessage({
        receiver: i_destinationContract,
        data: abi.encode(msg.sender, destinationAddress),
        tokens: tokens,
        amounts: amounts,
        feeToken: tokens[0],
        feeTokenAmount: 0,
        gasLimit: 0
      })
    );
  }
}
