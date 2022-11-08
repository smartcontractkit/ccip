// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {SafeERC20, IERC20} from "../../vendor/SafeERC20.sol";
import {EVM2AnyTollOnRampRouterInterface} from "../interfaces/onRamp/EVM2AnyTollOnRampRouterInterface.sol";
import {CCIP} from "../models/Models.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract TollSenderDapp is TypeAndVersionInterface {
  using CCIP for CCIP.EVMExtraArgsV1;
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "TollSenderDapp 1.0.0";

  // On ramp contract responsible for interacting with the DON.
  EVM2AnyTollOnRampRouterInterface public immutable i_onRampRouter;
  uint256 public immutable i_destinationChainId;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable i_destinationContract;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(
    EVM2AnyTollOnRampRouterInterface onRampRouter,
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
    address[] memory tokens,
    uint256[] memory amounts
  ) external returns (uint64 sequenceNumber) {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    for (uint256 i = 0; i < tokens.length; ++i) {
      IERC20(tokens[i]).safeTransferFrom(msg.sender, address(this), amounts[i]);
      IERC20(tokens[i]).approve(address(i_onRampRouter), amounts[i]);
    }
    // `data` format:
    //  - EOA sender address
    //  - EOA destination address
    sequenceNumber = i_onRampRouter.ccipSend(
      i_destinationChainId,
      CCIP.EVM2AnyTollMessage({
        receiver: abi.encode(i_destinationContract),
        data: abi.encode(msg.sender, destinationAddress),
        tokens: tokens,
        amounts: amounts,
        feeToken: tokens[0],
        feeTokenAmount: 0,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: 3e5})._toBytes()
      })
    );
  }
}
