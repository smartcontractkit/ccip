// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";
import "../pools/TokenPool.sol";

library CCIP {
  ////////////////////////////////
  ////         COMMON         ////
  ////////////////////////////////

  // Offchain leaf domain separator
  bytes32 public constant LEAF_DOMAIN_SEPARATOR = 0x0000000000000000000000000000000000000000000000000000000000000000;
  // Internal domain separator used in proofs
  bytes32 public constant INTERNAL_DOMAIN_SEPARATOR =
    0x0000000000000000000000000000000000000000000000000000000000000001;

  /// @notice Generalized EVM message type that is sent from EVM routers
  // to the contracts that implement the Any2EVMMessageReceiverInterface
  struct Any2EVMMessageFromSender {
    uint256 sourceChainId;
    bytes sender;
    address receiver;
    bytes data;
    IERC20[] destTokens;
    PoolInterface[] destPools;
    uint256[] amounts;
    uint256 gasLimit;
  }

  struct Any2EVMMessage {
    uint256 sourceChainId;
    bytes sender;
    bytes data;
    IERC20[] destTokens;
    uint256[] amounts;
  }

  function _toAny2EVMMessage(CCIP.Any2EVMMessageFromSender memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory message)
  {
    message = CCIP.Any2EVMMessage({
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      data: original.data,
      destTokens: original.destTokens,
      amounts: original.amounts
    });
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
    // Only used by toll as subscriptions use a known fee token.
    address[] tokenPerFeeCoinAddresses;
    // For subscriptions only the first value is used as
    // all subscriptions use the same fee token.
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

  ////////////////////////////////
  ////          TOLL          ////
  ////////////////////////////////

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

  bytes32 internal constant EVM_2_EVM_TOLL_MESSAGE_HASH = keccak256("EVM2EVMTollMessagePlus");

  function _hash(CCIP.EVM2EVMTollMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
    return
      keccak256(
        abi.encode(
          LEAF_DOMAIN_SEPARATOR,
          metadataHash,
          original.sequenceNumber,
          original.sender,
          original.receiver,
          keccak256(original.data),
          keccak256(abi.encode(original.tokens)),
          keccak256(abi.encode(original.amounts)),
          original.gasLimit,
          original.feeToken,
          original.feeTokenAmount
        )
      );
  }

  ////////////////////////////////
  ////      SUBSCRIPTION      ////
  ////////////////////////////////

  struct EVM2AnySubscriptionMessage {
    address receiver;
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

  bytes32 internal constant EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH = keccak256("EVM2EVMSubscriptionMessagePlus");

  function _hash(CCIP.EVM2EVMSubscriptionMessage memory original, bytes32 metadataHash)
    internal
    pure
    returns (bytes32)
  {
    return
      keccak256(
        abi.encode(
          LEAF_DOMAIN_SEPARATOR,
          metadataHash,
          original.sequenceNumber,
          original.sender,
          original.receiver,
          keccak256(original.data),
          keccak256(abi.encode(original.tokens)),
          keccak256(abi.encode(original.amounts)),
          original.gasLimit,
          original.nonce
        )
      );
  }
}
