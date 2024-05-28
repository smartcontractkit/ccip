// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import "../../interfaces/IMessageInterceptor.sol";

contract MessageInterceptorHelper is IMessageInterceptor {
  error IncomingMessageValidationError(bytes errorReason);

  mapping(bytes32 messageId => bool isInvalid) internal s_invalidMessageIds;

  constructor() {}

  function setMessageIdValidationState(bytes32 messageId, bool isInvalid) external {
    s_invalidMessageIds[messageId] = isInvalid;
  }

  /// @inheritdoc IMessageInterceptor
  function onIncomingMessage(Client.Any2EVMMessage memory message) external view {
    if (s_invalidMessageIds[message.messageId]) {
      revert IncomingMessageValidationError(bytes("Invalid message"));
    }
  }

  /// @inheritdoc IMessageInterceptor
  function onOutgoingMessage(Client.EVM2AnyMessage memory, uint64) external pure {
    // TODO: to be implemented
    return;
  }
}
