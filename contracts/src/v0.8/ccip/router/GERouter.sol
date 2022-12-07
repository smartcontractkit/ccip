// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {SafeERC20, IERC20} from "../../vendor/SafeERC20.sol";
import {PoolInterface} from "../interfaces/pools/PoolInterface.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Common} from "../models/Common.sol";
import {GERouterInterface, BaseOnRampRouterInterface} from "../interfaces/router/GERouterInterface.sol";
import {EVM2EVMGEOnRampInterface, BaseOnRampInterface} from "../interfaces/onRamp/EVM2EVMGEOnRampInterface.sol";
import {BaseOffRampRouter, BaseOffRampInterface, Any2EVMOffRampRouterInterface} from "../offRamp/BaseOffRampRouter.sol";
import {GEConsumer} from "../models/GEConsumer.sol";
import {Internal} from "../models/Internal.sol";

contract GERouter is GERouterInterface, BaseOffRampRouter, TypeAndVersionInterface {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GERouter 1.0.0";

  // destination chain id => OnRampInterface
  mapping(uint256 => EVM2EVMGEOnRampInterface) private s_onRamps;

  constructor(BaseOffRampInterface[] memory offRamps) BaseOffRampRouter(offRamps) {}

  /// @inheritdoc GERouterInterface
  function ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message) external returns (bytes32) {
    // Find and put the correct onRamp on the stack.
    EVM2EVMGEOnRampInterface onRamp = s_onRamps[destinationChainId];
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
      PoolInterface pool = onRamp.getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert BaseOnRampInterface.UnsupportedToken(token);
      token.safeTransferFrom(msg.sender, address(pool), combinedTokensAndAmounts[i].amount);
    }

    return onRamp.forwardFromRouter(message, feeTokenAmount, msg.sender);
  }

  /// @inheritdoc GERouterInterface
  // @dev returns 0 fee on invalid message.
  function getFee(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message)
    public
    view
    returns (uint256 fee)
  {
    // Find and put the correct onRamp on the stack.
    EVM2EVMGEOnRampInterface onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    return onRamp.getFee(message);
  }

  /// @inheritdoc GERouterInterface
  function setOnRamp(uint64 chainId, EVM2EVMGEOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc GERouterInterface
  function getOnRamp(uint64 chainId) external view returns (EVM2EVMGEOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint64 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }
}
