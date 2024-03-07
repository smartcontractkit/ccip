// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {GameType} from "./DisputeTypes.sol";

interface IOptimismPortal2 {
  /// @notice The game type that the OptimismPortal consults for output proposals.
  function respectedGameType() external view returns (GameType);
}
