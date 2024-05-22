// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import "../../interfaces/IMessageValidator.sol";

contract MessageValidatorHelper is IMessageValidator {
  error IncomingMessageValidationError(bytes errorReason);

  mapping(bytes32 messageId => bool isInvalid) internal s_invalidMessageIds;

  constructor() {}

  /// @inheritdoc IMessageValidator
  function validateIncomingMessage(Client.Any2EVMMessage memory message) external view {
    if (s_invalidMessageIds[message.messageId]) {
      revert IncomingMessageValidationError(bytes("Invalid message"));
    }
  }

  /// @inheritdoc IMessageValidator
  function validateOutgoingMessage(Client.EVM2AnyMessage memory, uint64) external pure {
    // TODO: to be implemented
    return;
  }
}
