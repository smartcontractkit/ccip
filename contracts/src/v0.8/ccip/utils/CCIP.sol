// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";

library CCIP {
  ////----**** TOLL ****----////

  /// @notice The Toll message type for EVM chains.
  struct EVM2AnyTollMessage {
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    IERC20 feeToken;
    uint256 feeTokenAmount;
    uint256 gasLimit;
  }

  /// @notice The event that gets emitted when an EVM to EVM cross chain toll request is made.
  struct EVM2EVMTollEvent {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    IERC20 feeToken;
    uint256 feeTokenAmount;
    uint256 gasLimit;
  }

  // @notice The cross chain message that gets relayed to EVM toll chains
  struct EVM2EVMTollMessage {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    IERC20 feeToken;
    uint256 feeTokenAmount;
    uint256 gasLimit;
  }

  function _toAny2EVMMessage(CCIP.EVM2EVMTollMessage memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory message)
  {
    message = CCIP.Any2EVMMessage({
      sourceChainId: original.sourceChainId,
      sequenceNumber: original.sequenceNumber,
      sender: abi.encode(original.sender),
      receiver: original.receiver,
      data: original.data,
      tokens: original.tokens,
      amounts: original.amounts,
      gasLimit: original.gasLimit
    });
  }

  ////----**** SUSBCRIPTION ****----////

  struct EVM2AnySubscriptionMessage {
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    uint256 gasLimit;
  }

  /// @notice The event that gets emitted when an EVM to EVM cross chain subscription request is made.
  struct EVM2EVMSubscriptionEvent {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    uint64 nonce;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    uint256 gasLimit;
  }

  // @notice The cross chain message that gets relayed to EVM subscription chains
  struct EVM2EVMSubscriptionMessage {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    uint64 nonce;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    uint256 gasLimit;
  }

  function _toAny2EVMMessage(CCIP.EVM2EVMSubscriptionMessage memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory message)
  {
    message = CCIP.Any2EVMMessage({
      sourceChainId: original.sourceChainId,
      sequenceNumber: original.sequenceNumber,
      sender: abi.encode(original.sender),
      receiver: original.receiver,
      data: original.data,
      tokens: original.tokens,
      amounts: original.amounts,
      gasLimit: original.gasLimit
    });
  }

  ////----**** COMMON ****----////

  /// @notice Generalized received EVM message
  struct Any2EVMMessage {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    bytes sender;
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    uint256 gasLimit;
  }

  /// @notice a sequenceNumber interval
  struct Interval {
    uint64 min;
    uint64 max;
  }

  /// @notice Report that is relayed by the observing DON at the relay phase
  struct RelayReport {
    address[] onRamps;
    Interval[] intervals;
    bytes32[] merkleRoots;
    bytes32 rootOfRoots;
  }

  struct ExecutionReport {
    uint64[] sequenceNumbers;
    address[] tokenPerFeeCoinAddresses;
    uint256[] tokenPerFeeCoin;
    bytes[] encodedMessages;
    bytes32[] innerProofs;
    uint256 innerProofFlagBits;
    bytes32[] outerProofs;
    uint256 outerProofFlagBits;
  }

  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }

  struct ExecutionResult {
    uint64 sequenceNumber;
    MessageExecutionState state;
  }
}
