// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IEVM2AnyTollOnRampRouter} from "../interfaces/onRamp/IEVM2AnyTollOnRampRouter.sol";

import {TollConsumer} from "../models/TollConsumer.sol";
import {Common} from "../models/Common.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {SafeERC20} from "../../vendor/SafeERC20.sol";

/**
 * @notice This contract enables EOAs to send a single asset across to the chain
 * represented by the On Ramp. Consider this an "Application Layer" contract that utilise the
 * underlying protocol.
 */
contract TollSenderDapp is TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "TollSenderDapp 1.0.0";

  // On ramp contract responsible for interacting with the DON.
  IEVM2AnyTollOnRampRouter public immutable i_onRampRouter;
  uint64 public immutable i_destinationChainId;
  // Corresponding contract on the destination chain responsible for receiving the message
  // and enabling the EOA on the destination chain to access the tokens that are sent.
  // For this scenario, it would be the address of a deployed EOASingleTokenReceiver.
  address public immutable i_destinationContract;

  error InvalidDestinationAddress(address invalidAddress);

  constructor(
    IEVM2AnyTollOnRampRouter onRampRouter,
    uint64 destinationChainId,
    address destinationContract
  ) {
    i_onRampRouter = onRampRouter;
    i_destinationChainId = destinationChainId;
    i_destinationContract = destinationContract;
  }

  /**
   * @notice Send tokensAndAmounts to the destination chain.
   * @dev msg.sender must first call TOKEN.approve for this contract to spend the tokensAndAmounts.
   */
  function sendTokens(address destinationAddress, Common.EVMTokenAndAmount[] memory tokensAndAmounts)
    external
    returns (uint64 sequenceNumber)
  {
    if (destinationAddress == address(0)) revert InvalidDestinationAddress(destinationAddress);
    for (uint256 i = 0; i < tokensAndAmounts.length; ++i) {
      IERC20(tokensAndAmounts[i].token).safeTransferFrom(msg.sender, address(this), tokensAndAmounts[i].amount);
      IERC20(tokensAndAmounts[i].token).approve(address(i_onRampRouter), tokensAndAmounts[i].amount);
    }
    // `data` format:
    //  - EOA sender address
    //  - EOA destination address
    sequenceNumber = i_onRampRouter.ccipSend(
      i_destinationChainId,
      TollConsumer.EVM2AnyTollMessage({
        receiver: abi.encode(i_destinationContract),
        data: abi.encode(msg.sender, destinationAddress),
        tokensAndAmounts: tokensAndAmounts,
        feeTokenAndAmount: tokensAndAmounts[0],
        extraArgs: TollConsumer._argsToBytes(TollConsumer.EVMExtraArgsV1({gasLimit: 3e5, strict: false}))
      })
    );
  }
}
