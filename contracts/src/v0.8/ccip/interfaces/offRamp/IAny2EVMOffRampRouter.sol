// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IBaseOffRamp} from "./IBaseOffRamp.sol";

import {Common} from "../../models/Common.sol";
import {Internal} from "../../models/Internal.sol";

interface IAny2EVMOffRampRouter {
  error NoOffRampsConfigured();
  error MustCallFromOffRamp(address sender);
  error SenderNotAllowed(address sender);
  error InvalidAddress();
  error OffRampNotAllowed(IBaseOffRamp offRamp);
  error AlreadyConfigured(IBaseOffRamp offRamp);

  event OffRampAdded(IBaseOffRamp indexed offRamp);
  event OffRampRemoved(IBaseOffRamp indexed offRamp);

  struct OffRampDetails {
    uint96 listIndex;
    bool allowed;
  }

  /**
   * @notice Owner can add an offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to add
   */
  function addOffRamp(IBaseOffRamp offRamp) external;

  /**
   * @notice Owner can remove a specific offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to remove
   */
  function removeOffRamp(IBaseOffRamp offRamp) external;

  /**
   * @notice Gets all configured offRamps.
   * @return offRamps The offRamp that are configured.
   */
  function getOffRamps() external view returns (IBaseOffRamp[] memory offRamps);

  /**
   * @notice Returns whether the given offRamp is set to be allowed
   * @param offRamp The offRamp to check.
   * @return allowed True if the offRamp is allowed, false if not.
   */
  function isOffRamp(IBaseOffRamp offRamp) external view returns (bool allowed);

  /**
   * @notice Route the message to its intended receiver contract
   * @param message Common.Any2EVMMessage struct
   * @param manualExecution bool to indicate manual instead of DON execution
   */
  function routeMessage(
    Common.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external returns (bool success);
}
