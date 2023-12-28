// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

/// @notice Interface for a liquidity container, this can be a CCIP token pool.
interface ILiquidityContainer {
  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed provider, uint256 indexed amount);
  
  function provideLiquidity(uint256 amount) external;

  function withdrawLiquidity(uint256 amount) external;
}
