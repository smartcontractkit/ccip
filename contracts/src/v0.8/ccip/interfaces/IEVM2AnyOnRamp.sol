// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IPool} from "./pools/IPool.sol";

import {Client} from "../libraries/Client.sol";
import {Internal} from "../libraries/Internal.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";

import {IEVM2AnyOnRampRouter} from "./IEVM2AnyOnRampRouter.sol";

interface IEVM2AnyOnRamp is IEVM2AnyOnRampRouter {
  /// @notice Gets the next sequence number to be used in the onRamp
  /// @return the next sequence number to be used
  function getExpectedNextSequenceNumber() external view returns (uint64);

  /// @notice Get the next nonce for a given sender
  /// @param sender The sender to get the nonce for
  /// @return nonce The next nonce for the sender
  function getSenderNonce(address sender) external view returns (uint64 nonce);

  /// @notice Adds and removed token pools.
  /// @param removes The tokens and pools to be removed
  /// @param adds The tokens and pools to be added.
  function applyPoolUpdates(Internal.PoolUpdate[] memory removes, Internal.PoolUpdate[] memory adds) external;
}
