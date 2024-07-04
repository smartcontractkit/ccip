// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "../../libraries/Client.sol";
import {CCIPClientBase} from "./CCIPClientBase.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableMap} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableMap.sol";

contract CCIPReceiver is CCIPClientBase {
  using SafeERC20 for IERC20;
  using EnumerableMap for EnumerableMap.Bytes32ToUintMap;

  error OnlySelf();
  error MessageNotFailed(bytes32 messageId);

  event MessageFailed(bytes32 indexed messageId, bytes reason);
  event MessageSucceeded(bytes32 indexed messageId);
  event MessageRecovered(bytes32 indexed messageId);
  event MessageAbandoned(bytes32 indexed messageId, address tokenReceiver);

  // Example error code, could have many different error codes.
  enum ErrorCode {
    // RESOLVED is first so that the default value is resolved.
    RESOLVED,
    FAILED,
    // TODO comment
    ABANDONED
  }

  // The message contents of failed messages are stored here.
  mapping(bytes32 messageId => Client.Any2EVMMessage contents) internal s_messageContents;

  // Contains failed messages and their state.
  EnumerableMap.Bytes32ToUintMap internal s_failedMessages;

  constructor(address router) CCIPClientBase(router) {}

  function typeAndVersion() external pure virtual returns (string memory) {
    return "CCIPReceiver 1.0.0-dev";
  }

  // ================================================================
  // │                  Incoming Message Processing                 |
  // ================================================================

  /// @notice The entrypoint for the CCIP router to call. This function should
  /// never revert, all errors should be handled internally in this contract.
  /// @param message The message to process.
  /// @dev Extremely important to ensure only router calls this.
  function ccipReceive(Client.Any2EVMMessage calldata message)
    external
    virtual
    onlyRouter
    isValidChain(message.sourceChainSelector)
  {
    try this.processMessage(message) {}
    catch (bytes memory err) {
      // TODO kinda dangerous comments here. 
      //  The original suggestion was:
      //  if you want custom retry logic, plus owner extracting tokens as a last resort for recovery, use this try-catch pattern in ccipReceiver
      //  this means the message will appear as success to CCIP, and you can track the actual message state within the dapp
      //  if you do not need custom retry logic, and you don't need owner token recovery function, then you don't need the try-catch here
      //  because you can use ccip manualExecution as a retry function
      //
      // we should not phrase it such that "Any failures should be tracked by individual Dapps", it's only the case if they want custom retry pattern
      // as opposed to default manual exec, or they want token recovery feature.

      // Mark the message as having failed. Any failures should be tracked by individual Dapps, since CCIP
      // will mark the message as having been successfully delivered. CCIP makes no assurances about message delivery
      // other than invocation with proper gas limit. Any logic/execution errors should be tracked by separately.
      s_failedMessages.set(message.messageId, uint256(ErrorCode.FAILED));

      // Store the message contents in case it needs to be retried or abandoned
      s_messageContents[message.messageId] = message;

      // Don't revert because CCIPRouter doesn't revert. Emit event instead.
      // The message can be retried or abandoned later without having to do manual execution of CCIP, which should
      // be reserved for retrying with a higher gas limit.
      emit MessageFailed(message.messageId, err);
      return;
    }

    emit MessageSucceeded(message.messageId);
  }

  /// @notice Contains arbitrary application-logic for incoming CCIP messages.
  /// @dev It has to be external because of the try/catch of ccipReceive() which invokes it
  function processMessage(Client.Any2EVMMessage calldata message)
    external
    virtual
    onlySelf
    isValidSender(message.sourceChainSelector, message.sender)
  {}

  // ================================================================
  // │                  Failed Message Processing                   |
  // ================== ==============================================

  // TODO we don't do new @dev on every new line, same for @notice, when in doubt check other contracts.
  /// @notice Execute a message that failed initial delivery, but with different logic specifically for re-execution.
  /// @dev Since the function invoked _retryFailedMessage(), which is marked onlyOwner, this may only be called by the Owner as well.
  /// @dev function will revert if the messageId was not already stored as having failed its initial execution
  /// @param messageId the unique ID of the CCIP-message which failed initial-execution.
  function retryFailedMessage(bytes32 messageId) external {
    if (s_failedMessages.get(messageId) != uint256(ErrorCode.FAILED)) revert MessageNotFailed(messageId);

    // Set the error code to 0 to disallow reentry and retry the same failed message
    // multiple times.
    s_failedMessages.set(messageId, uint256(ErrorCode.RESOLVED));

    // Allow developer to implement arbitrary functionality on retried messages, such as just releasing the associated tokens
    Client.Any2EVMMessage memory message = s_messageContents[messageId];

    // Allow the user override the implementation, since different workflow may be desired for retrying a merssage
    _retryFailedMessage(message);

    emit MessageRecovered(messageId);
  }

  /// @notice A function that should contain any special logic needed to "retry" processing of a previously failed message.
  /// TODO these funtions do not exist
  /// @dev If the owner wants to retrieve tokens without special logic, then abandonMessage() or recoverTokens() should be used instead
  /// @dev function is marked onlyOwner, but is virtual. Allowing permissionless execution is not recommended but may be allowed if function is overridden
  function _retryFailedMessage(Client.Any2EVMMessage memory message) internal virtual onlyOwner {
    // TODO how about we add a default implementation that calls `processMessage`, and comments for overrides
  }

  // TODO double notices
  /// @notice Should be used to recover tokens from a failed message, while ensuring the message cannot be retried
  /// @notice function will send tokens to destination, but will NOT invoke any arbitrary logic afterwards.
  /// @dev this function is only callable as the owner, and
  function abandonFailedMessage(bytes32 messageId, address receiver) external onlyOwner {
    if (s_failedMessages.get(messageId) != uint256(ErrorCode.FAILED)) revert MessageNotFailed(messageId);

    s_failedMessages.set(messageId, uint256(ErrorCode.ABANDONED));
    Client.Any2EVMMessage memory message = s_messageContents[messageId];

    for (uint256 i = 0; i < message.destTokenAmounts.length; ++i) {
      IERC20(message.destTokenAmounts[i].token).safeTransfer(receiver, message.destTokenAmounts[i].amount);
    }

    emit MessageAbandoned(messageId, receiver);
  }

  // ================================================================
  // │                  Message Tracking                            │
  // ================================================================

  /// @param messageId the ID of the message delivered by the CCIP Router
  /// @return Any2EVMMessage a standard CCIP message for EVM-compatible networks
  function getMessageContents(bytes32 messageId) public view returns (Client.Any2EVMMessage memory) {
    return s_messageContents[messageId];
  }

  /// @param messageId the ID of the message delivered by the CCIP Router
  function getMessageStatus(bytes32 messageId) public view returns (uint256) {
    return s_failedMessages.get(messageId);
  }

  modifier onlySelf() {
    if (msg.sender != address(this)) revert OnlySelf();
    _;
  }
}
