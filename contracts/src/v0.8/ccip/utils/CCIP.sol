// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../vendor/IERC20.sol";

contract CCIP {
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

  /// @notice The event that gets emitted when an EVM to EVM cross chain request is made.
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

  /// @notice Report that is relayed by the observing DON at the relay phase
  struct RelayReport {
    bytes32 merkleRoot;
    uint64 minSequenceNumber;
    uint64 maxSequenceNumber;
  }

  // @notice The cross chain message that gets relayed to EVM chains
  struct Any2EVMTollMessage {
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

  // TODO: This is a single root for now, enable many roots in one report.
  struct ExecutionReport {
    Any2EVMTollMessage[] messages;
    bytes32[] proofs;
    uint256 proofFlagsBits;
  }
}
