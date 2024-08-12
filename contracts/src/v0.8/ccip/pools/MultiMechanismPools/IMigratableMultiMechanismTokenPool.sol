pragma solidity ^0.8.0;

interface IMigratableMultiMechanismTokenPool {
  function withdrawLiquidity(uint64 remoteChainSelector, uint256 amount) external;
}
