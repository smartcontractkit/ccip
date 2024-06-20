// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IMessageInterceptor} from "../../interfaces/IMessageInterceptor.sol";
import {Client} from "../../libraries/Client.sol";

contract MessageInterceptorHelper is IMessageInterceptor {
  error IncomingMessageValidationError(bytes errorReason);

  mapping(bytes32 messageId => bool isInvalid) internal s_invalidMessageIds;

  constructor() {}

  function setMessageIdValidationState(bytes32 messageId, bool isInvalid) external {
    s_invalidMessageIds[messageId] = isInvalid;
  }

  /// @inheritdoc IMessageInterceptor
  function onIncomingMessage(Client.Any2EVMMessage calldata message) external view {
    if (s_invalidMessageIds[message.messageId]) {
      revert IncomingMessageValidationError(bytes("Invalid message"));
    }
  }

  /// @inheritdoc IMessageInterceptor
  function onOutgoingMessage(uint64, Client.EVM2AnyMessage calldata message) external view {
    if (s_invalidMessageIds[keccak256(abi.encode(message))]) {
      revert IncomingMessageValidationError(bytes("Invalid message"));
    }
    return;
  }
}
