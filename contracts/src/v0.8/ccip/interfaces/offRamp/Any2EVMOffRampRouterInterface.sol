// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {BaseOffRampInterface} from "./BaseOffRampInterface.sol";
import {Common} from "../../models/Common.sol";
import {Internal} from "../../models/Internal.sol";

interface Any2EVMOffRampRouterInterface {
  error NoOffRampsConfigured();
  error MustCallFromOffRamp(address sender);
  error SenderNotAllowed(address sender);
  error InvalidAddress();
  error OffRampNotAllowed(BaseOffRampInterface offRamp);
  error AlreadyConfigured(BaseOffRampInterface offRamp);

  event OffRampAdded(BaseOffRampInterface indexed offRamp);
  event OffRampRemoved(BaseOffRampInterface indexed offRamp);

  struct OffRampDetails {
    uint96 listIndex;
    bool allowed;
  }

  /**
   * @notice Owner can add an offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to add
   */
  function addOffRamp(BaseOffRampInterface offRamp) external;

  /**
   * @notice Owner can remove a specific offRamp from the allowlist
   * @dev Only callable by the owner
   * @param offRamp The offRamp to remove
   */
  function removeOffRamp(BaseOffRampInterface offRamp) external;

  /**
   * @notice Gets all configured offRamps.
   * @return offRamps The offRamp that are configured.
   */
  function getOffRamps() external view returns (BaseOffRampInterface[] memory offRamps);

  /**
   * @notice Returns whether the given offRamp is set to be allowed
   * @param offRamp The offRamp to check.
   * @return allowed True if the offRamp is allowed, false if not.
   */
  function isOffRamp(BaseOffRampInterface offRamp) external view returns (bool allowed);

  /**
   * @notice Route the message to its intended receiver contract
   * @param message Internal.Any2EVMMessageFromSender struct
   * @param manualExecution bool to indicate manual instead of DON execution
   */
  function routeMessage(Internal.Any2EVMMessageFromSender calldata message, bool manualExecution)
    external
    returns (bool success);
}
