// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "../../interfaces/IEVM2AnyOnRamp.sol";
import {EVM2EVMOnRamp} from "../../onRamp/EVM2EVMOnRamp.sol";

interface IEVM2AnyOnRampExtended is IEVM2AnyOnRamp {
  /// @notice Pays the Node Ops their outstanding balances.
  function payNops() external;

  /// @notice Gets the Nops and their weights
  /// @return nopsAndWeights Array of NopAndWeight structs
  /// @return weightsTotal The sum weight of all Nops
  function getNops() external view returns (EVM2EVMOnRamp.NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal);

  /// @notice Returns the link that would remain after paying the Nops.
  function linkAvailableForPayment() external view returns (int256);

  /// @notice Get the total amount of fees to be paid to the Nops (in LINK)
  /// @return totalNopFees
  function getNopFeesJuels() external view returns (uint96);
}
