// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IGERouter} from "../interfaces/router/IGERouter.sol";
import {IBaseOnRampRouter} from "../interfaces/onRamp/IBaseOnRampRouter.sol";
import {IBaseOnRamp} from "../interfaces/onRamp/IBaseOnRamp.sol";
import {IEVM2AnyGEOnRamp} from "../interfaces/onRamp/IEVM2AnyGEOnRamp.sol";
import {IBaseOffRamp} from "../interfaces/offRamp/IBaseOffRamp.sol";

import {BaseOffRampRouter} from "../offRamp/BaseOffRampRouter.sol";
import {GEConsumer} from "../models/GEConsumer.sol";
import {Internal} from "../models/Internal.sol";
import {Common} from "../models/Common.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";

contract GERouter is IGERouter, BaseOffRampRouter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GERouter 1.0.0";

  // destination chain id => IOnRamp
  mapping(uint256 => IEVM2AnyGEOnRamp) private s_onRamps;

  constructor(IBaseOffRamp[] memory offRamps) BaseOffRampRouter(offRamps) {}

  /// @inheritdoc IGERouter
  function ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message) external returns (bytes32) {
    IEVM2AnyGEOnRamp onRamp = s_onRamps[destinationChainId];
    // getFee checks if the onRamp is valid
    uint256 feeTokenAmount = getFee(destinationChainId, message);
    IERC20(message.feeToken).safeTransferFrom(msg.sender, address(onRamp), feeTokenAmount);

    // Transfer the tokens to the token pools.
    // TODO: Check the pool for how to take action
    for (uint256 i = 0; i < message.tokensAndAmounts.length; ++i) {
      IERC20 token = IERC20(message.tokensAndAmounts[i].token);
      IPool pool = onRamp.getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert IBaseOnRamp.UnsupportedToken(token);
      token.safeTransferFrom(msg.sender, address(pool), message.tokensAndAmounts[i].amount);
    }

    return onRamp.forwardFromRouter(message, feeTokenAmount, msg.sender);
  }

  /// @inheritdoc IGERouter
  // @dev returns 0 fee on invalid message.
  function getFee(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message)
    public
    view
    returns (uint256 fee)
  {
    // Find and put the correct onRamp on the stack.
    IEVM2AnyGEOnRamp onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    return onRamp.getFee(message);
  }

  /// @inheritdoc IGERouter
  function setOnRamp(uint64 chainId, IEVM2AnyGEOnRamp onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc IGERouter
  function getOnRamp(uint64 chainId) external view returns (IEVM2AnyGEOnRamp) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc IBaseOnRampRouter
  function isChainSupported(uint64 chainId) public view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  /// @inheritdoc IGERouter
  function getSupportedTokens(uint64 destChainId) external view returns (address[] memory) {
    if (!isChainSupported(destChainId)) {
      return new address[](0);
    }
    return s_onRamps[uint256(destChainId)].getSupportedTokens();
  }
}
