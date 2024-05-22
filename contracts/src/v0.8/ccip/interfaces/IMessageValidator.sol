// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../libraries/Client.sol";

/// @notice Interface for plug-in message validator contracts that intercept OffRamp & OnRamp messages
///         and perform validations on top of the messages. All validation functions are expected to
///         revert on validation failures.
interface IMessageValidator {
  /// @notice Common error that can be thrown on validation failures and used by consumers
  /// @param errorReason abi encoded revert reason
  error MessageValidationError(bytes errorReason);

  /// @notice Validates the given OffRamp message. Reverts on validation failure
  /// @param message to validate
  function validateIncomingMessage(Client.Any2EVMMessage memory message) external;

  /// @notice Validates the given OnRamp message. Reverts on validation failure
  /// @param message to valdidate
  /// @param destChainSelector dest chain selector where the message is being sent to
  function validateOutgoingMessage(Client.EVM2AnyMessage memory message, uint64 destChainSelector) external;
}
