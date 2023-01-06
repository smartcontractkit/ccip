// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IEVM2AnyTollOnRampRouter} from "../../interfaces/onRamp/IEVM2AnyTollOnRampRouter.sol";
import {IEVM2EVMTollOnRamp} from "../../interfaces/onRamp/IEVM2EVMTollOnRamp.sol";
import {IBaseOnRampRouter} from "../../interfaces/onRamp/IBaseOnRampRouter.sol";
import {IBaseOnRamp} from "../../interfaces/onRamp/IBaseOnRamp.sol";

import {PoolCollector} from "../../pools/PoolCollector.sol";
import {OwnerIsCreator} from "../../access/OwnerIsCreator.sol";
import {TollConsumer} from "../../models/TollConsumer.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

contract EVM2AnyTollOnRampRouter is IEVM2AnyTollOnRampRouter, TypeAndVersionInterface, OwnerIsCreator, PoolCollector {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2AnyTollOnRampRouter 1.0.0";

  // destination chain id => IOnRamp
  mapping(uint256 => IEVM2EVMTollOnRamp) private s_onRamps;

  /// @inheritdoc IEVM2AnyTollOnRampRouter
  function ccipSend(uint64 destinationChainId, TollConsumer.EVM2AnyTollMessage memory message)
    external
    returns (uint64)
  {
    // Find and put the correct onRamp on the stack.
    IEVM2EVMTollOnRamp onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);

    // Charge the fee and subtract that amount from the feeTokenAmount. This will revert if
    // the given feeTokenAmount is too low for the needed fee.
    message.feeTokenAndAmount.amount -= _chargeFee(
      onRamp,
      IERC20(message.feeTokenAndAmount.token),
      message.feeTokenAndAmount.amount
    );
    // Transfer the tokensAndAmounts to the token pools.
    _collectTokens(onRamp, message.tokensAndAmounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc IEVM2AnyTollOnRampRouter
  function setOnRamp(uint64 chainId, IEVM2EVMTollOnRamp onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc IEVM2AnyTollOnRampRouter
  function getOnRamp(uint64 chainId) external view returns (IEVM2EVMTollOnRamp) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc IBaseOnRampRouter
  function isChainSupported(uint64 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
