// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface ILiquidityManager {
  struct SendLiquidityParams {
    uint256 amount;
    uint64 destChainSelector;
  }

  struct ReceiveLiquidityParams {
    uint256 amount;
    uint64 sourceChainSelector;
    bytes bridgeData;
  }

  struct LiquidityInstructions {
    SendLiquidityParams[] sendLiquidityParams;
    ReceiveLiquidityParams[] receiveLiquidityParams;
  }

  /// @notice Returns the current liquidity in the liquidity container.
  function getLiquidity() external view returns (uint256 currentLiquidity);
}
