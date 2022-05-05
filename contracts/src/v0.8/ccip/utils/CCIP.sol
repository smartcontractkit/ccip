// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../vendor/IERC20.sol";

contract CCIP {
  /// @notice High level message
  struct Message {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    MessagePayload payload;
  }

  /// @notice Payload within the message
  struct MessagePayload {
    IERC20[] tokens;
    uint256[] amounts;
    uint256 destinationChainId;
    address receiver;
    address executor;
    bytes data;
  }

  /// @notice Report that is relayed by the observing DON at the relay phase
  struct RelayReport {
    bytes32 merkleRoot;
    uint64 minSequenceNumber;
    uint64 maxSequenceNumber;
  }

  // TODO: This is a single root for now, enable many roots in one report.
  struct ExecutionReport {
    Message[] messages;
    bytes32[] proofs;
    uint256 proofFlagsBits;
  }
}
