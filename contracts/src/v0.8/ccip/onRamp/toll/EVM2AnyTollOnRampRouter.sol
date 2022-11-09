// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {EVM2AnyTollOnRampRouterInterface, EVM2EVMTollOnRampInterface, BaseOnRampRouterInterface} from "../../interfaces/onRamp/EVM2AnyTollOnRampRouterInterface.sol";
import {BaseOnRampInterface} from "../../interfaces/onRamp/BaseOnRampInterface.sol";
import {PoolCollector} from "../../pools/PoolCollector.sol";
import {OwnerIsCreator} from "../../access/OwnerIsCreator.sol";
import {CCIP} from "../../models/Models.sol";
import {IERC20} from "../../../vendor/IERC20.sol";

contract EVM2AnyTollOnRampRouter is
  EVM2AnyTollOnRampRouterInterface,
  TypeAndVersionInterface,
  OwnerIsCreator,
  PoolCollector
{
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2AnyTollOnRampRouter 1.0.0";

  // destination chain id => OnRampInterface
  mapping(uint256 => EVM2EVMTollOnRampInterface) private s_onRamps;

  /// @inheritdoc EVM2AnyTollOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyTollMessage memory message) external returns (uint64) {
    // Find and put the correct onRamp on the stack.
    EVM2EVMTollOnRampInterface onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);

    // Charge the fee and subtract that amount from the feeTokenAmount. This will revert if
    // the given feeTokenAmount is too low for the needed fee.
    message.feeTokenAndAmount.amount -= _chargeFee(onRamp, IERC20(message.feeTokenAndAmount.token), message.feeTokenAndAmount.amount);
    // Transfer the tokensAndAmounts to the token pools.
    _collectTokens(onRamp, message.tokensAndAmounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc EVM2AnyTollOnRampRouterInterface
  function setOnRamp(uint256 chainId, EVM2EVMTollOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc EVM2AnyTollOnRampRouterInterface
  function getOnRamp(uint256 chainId) external view returns (EVM2EVMTollOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
