// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {SafeERC20, IERC20} from "../../vendor/SafeERC20.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Common} from "../models/Common.sol";
import {IGERouter, IBaseOnRampRouter} from "../interfaces/router/IGERouter.sol";
import {IEVM2EVMGEOnRamp, IBaseOnRamp} from "../interfaces/onRamp/IEVM2EVMGEOnRamp.sol";
import {BaseOffRampRouter, IBaseOffRamp, IAny2EVMOffRampRouter} from "../offRamp/BaseOffRampRouter.sol";
import {GEConsumer} from "../models/GEConsumer.sol";
import {Internal} from "../models/Internal.sol";

contract GERouter is IGERouter, BaseOffRampRouter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GERouter 1.0.0";

  // destination chain id => IOnRamp
  mapping(uint256 => IEVM2EVMGEOnRamp) private s_onRamps;

  constructor(IBaseOffRamp[] memory offRamps) BaseOffRampRouter(offRamps) {}

  /// @inheritdoc IGERouter
  function ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message) external returns (bytes32) {
    // Find and put the correct onRamp on the stack.
    IEVM2EVMGEOnRamp onRamp = s_onRamps[destinationChainId];
    // getFee checks if the onRamp is valid
    uint256 feeTokenAmount = getFee(destinationChainId, message);

    Common.EVMTokenAndAmount[] memory combinedTokensAndAmounts = Internal._addToTokensAmounts(
      message.tokensAndAmounts,
      Common.EVMTokenAndAmount({token: message.feeToken, amount: feeTokenAmount})
    );
    // Transfer the tokensAndAmounts to the token pools.
    // TODO: Check the pool for how to take action
    for (uint256 i = 0; i < combinedTokensAndAmounts.length; ++i) {
      IERC20 token = IERC20(combinedTokensAndAmounts[i].token);
      IPool pool = onRamp.getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert IBaseOnRamp.UnsupportedToken(token);
      token.safeTransferFrom(msg.sender, address(pool), combinedTokensAndAmounts[i].amount);
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
    IEVM2EVMGEOnRamp onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    return onRamp.getFee(message);
  }

  /// @inheritdoc IGERouter
  function setOnRamp(uint64 chainId, IEVM2EVMGEOnRamp onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc IGERouter
  function getOnRamp(uint64 chainId) external view returns (IEVM2EVMGEOnRamp) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc IBaseOnRampRouter
  function isChainSupported(uint64 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
