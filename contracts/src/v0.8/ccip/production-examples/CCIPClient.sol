pragma solidity ^0.8.0;

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {CCIPReceiverWithACK} from "./CCIPReceiverWithACK.sol";

import {IRouterClient} from "../interfaces/IRouterClient.sol";
import {Client} from "../libraries/Client.sol";

contract CCIPClient is CCIPReceiverWithACK {
  using SafeERC20 for IERC20;

  error InvalidConfig();

  /// @notice You can't import CCIPReceiver and Sender due to similar parents so functionality of CCIPSender is duplicated here
  constructor(address router, IERC20 feeToken) CCIPReceiverWithACK(router, feeToken) {}

  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes memory data,
    address feeToken
  ) public payable validChain(destChainSelector) {
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: s_chains[destChainSelector],
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: s_extraArgsBytes[destChainSelector],
      feeToken: feeToken
    });

    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);

    // TODO: Decide whether workflow should assume contract is funded with tokens to send already
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      // Transfer the tokens to pay for tokens in tokenAmounts
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);

      // If they're not sending the fee token, then go ahead and approve
      if (tokenAmounts[i].token != feeToken) {
        IERC20(tokenAmounts[i].token).safeApprove(i_ccipRouter, tokenAmounts[i].amount);
      }
      // If they are sending the feeToken through, and also paying in it, then approve the router for both tokenAmount and the fee()
      else if (tokenAmounts[i].token == feeToken && feeToken != address(0)) {
        IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), fee);
        IERC20(tokenAmounts[i].token).safeApprove(i_ccipRouter, tokenAmounts[i].amount + fee);
      }
    }

    // If the user is paying in the fee token, and is NOT sending it through the bridge, then allowance() should be zero
    // and we can send just transferFrom the sender and approve the router. This is because we only approve the router
    // for the amount of tokens needed for this transaction, one at a time.
    if (feeToken != address(0) && IERC20(feeToken).allowance(address(this), i_ccipRouter) == 0) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), fee);

      // Use increaseAllowance in case the user is transfering the feeToken in tokenAmounts
      IERC20(feeToken).safeApprove(i_ccipRouter, fee);
    } else if (feeToken == address(0) && msg.value < fee) {
      revert IRouterClient.InsufficientFeeTokenAmount();
    }

    bytes32 messageId =
      IRouterClient(i_ccipRouter).ccipSend{value: feeToken == address(0) ? fee : 0}(destChainSelector, message);

    emit MessageSent(messageId);
  }
}
