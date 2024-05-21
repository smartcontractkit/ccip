import {Client} from "../libraries/Client.sol";

/// @notice Interface for plug-in message validator contracts that intercept OffRamp & OnRamp messages
///         and perform validations on top of the messages. All validation functions are expected to
///         revert on validation failures.
interface IMessageValidator {
  /// @notice The error that is expected to be thrown on validation failures
  /// @param errorReason abi encoded revert reason
  error MessageValidationFailure(bytes errorReason);

  /// @notice Validates the given OffRamp message. Reverts on validation failure
  /// @param message to validate
  function validateIncomingMessage(Client.Any2EVMMessage memory message) external view;

  /// @notice Validates the given OnRamp message. Reverts on validation failure
  /// @param message to valdidate
  function validateOutgoingMessages(Client.EVM2AnyMessage memory message) external view;
}
