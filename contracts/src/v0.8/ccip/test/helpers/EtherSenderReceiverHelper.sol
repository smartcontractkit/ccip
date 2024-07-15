// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {CCIPSender} from "../../applications/external/CCIPSender.sol";
import {Client} from "../../libraries/Client.sol";
import {CCIPReceiverBasic} from "./receivers/CCIPReceiverBasic.sol";

import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface CCIPRouter {
  function getWrappedNative() external view returns (address);
}

contract EtherSenderReceiverHelper is CCIPSender {
  using SafeERC20 for IERC20;

  error InvalidTokenAmounts(uint256 gotAmounts);
  error InvalidToken(address gotToken, address expectedToken);
  error TokenAmountNotEqualToMsgValue(uint256 gotAmount, uint256 msgValue);
  error InsufficientMsgValue(uint256 gotAmount, uint256 msgValue);
  error InsufficientFee(uint256 gotFee, uint256 fee);
  error GasLimitTooLow(uint256 minLimit, uint256 gotLimit);

  IWrappedNative public immutable i_weth;

  constructor(address router) CCIPSender(router) {
    i_weth = IWrappedNative(CCIPRouter(router).getWrappedNative());
    IERC20(i_weth).safeApprove(router, type(uint256).max);
  }

  function validatedMessage(Client.EVM2AnyMessage calldata message) public view returns (Client.EVM2AnyMessage memory) {
    return _validatedMessage(message);
  }

  function validateFeeToken(Client.EVM2AnyMessage calldata message) public payable {
    _validateFeeToken(message);
  }

  function _validateFeeToken(Client.EVM2AnyMessage calldata message) internal view {
    uint256 tokenAmount = message.tokenAmounts[0].amount;

    if (message.feeToken != address(0)) {
      // If the fee token is NOT native, then the token amount must be equal to msg.value.
      // This is done to ensure that there is no leftover ether in this contract.
      if (msg.value != tokenAmount) {
        revert TokenAmountNotEqualToMsgValue(tokenAmount, msg.value);
      }
    }
  }

  /// @notice Validate the message content.
  /// @dev Only allows a single token to be sent. Always overwritten to be address(i_weth)
  /// and receiver is always msg.sender.
  function _validatedMessage(Client.EVM2AnyMessage calldata message)
    internal
    view
    returns (Client.EVM2AnyMessage memory)
  {
    Client.EVM2AnyMessage memory validMessage = message;

    if (validMessage.tokenAmounts.length != 1) {
      revert InvalidTokenAmounts(validMessage.tokenAmounts.length);
    }

    validMessage.data = abi.encode(msg.sender);
    validMessage.tokenAmounts[0].token = address(i_weth);

    return validMessage;
  }

  function publicCcipReceive(Client.Any2EVMMessage memory message) public {
    _ccipReceive(message);
  }

  function ccipReceive(Client.Any2EVMMessage calldata message) external virtual onlyRouter {
    _ccipReceive(message);
  }

  /// @notice Override this function in your implementation.
  /// @param message Any2EVMMessage
  function _ccipReceive(Client.Any2EVMMessage memory message) internal virtual {}
}
