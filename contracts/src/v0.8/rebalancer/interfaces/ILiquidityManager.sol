// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {IBridgeAdapter} from "./IBridge.sol";

interface ILiquidityManager {
  struct SendLiquidityParams {
    uint256 amount;
    uint64 remoteChainSelector;
  }

  struct ReceiveLiquidityParams {
    uint256 amount;
    uint64 remoteChainSelector;
    bytes bridgeData;
  }

  struct LiquidityInstructions {
    SendLiquidityParams[] sendLiquidityParams;
    ReceiveLiquidityParams[] receiveLiquidityParams;
  }

  struct CrossChainLiquidityManagerArgs {
    address remoteLiquidityManager;
    IBridgeAdapter localBridge;
    address remoteToken;
    uint64 remoteChainSelector;
    bool enabled;
  }

  /// @notice Returns the current liquidity in the liquidity container.
  function getLiquidity() external view returns (uint256 currentLiquidity);

  function getAllCrossChainLiquidityMangers() external view returns (CrossChainLiquidityManagerArgs[] memory);
}
