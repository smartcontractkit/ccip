// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {ILiquidityContainer} from "../../../liquidity-manager/interfaces/ILiquidityContainer.sol";
import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {LiquidityManager} from "../../pools/liquidity/LiquidityManager.sol";
import {ILiquidityManager} from "../../pools/liquidity/interfaces/ILiquidityManager.sol";

contract TestOnlyLiquidityManager is LiquidityManager {
  constructor(IERC20 token, uint64 localChainSelector, ILiquidityContainer localLiquidityContainer) LiquidityManager(token, localChainSelector, localLiquidityContainer) {}

  function publicReport(bytes calldata report, uint64 ocrSeqNum) external {
    _report(report, ocrSeqNum);
  }

  /// @dev exposed so that we can encode the report offchain
  function publicEncodeReport(ILiquidityManager.LiquidityInstructions memory liquidityInstructions) external pure returns (bytes memory) {
    return abi.encode(liquidityInstructions);
  }
}
