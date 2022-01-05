// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";

contract CCIP {
  /// @notice High level message
  struct Message {
    uint256 sequenceNumber;
    uint256 sourceChainId;
    uint256 destinationChainId;
    address sender;
    MessagePayload payload;
  }

  /// @notice Payload within the message
  struct MessagePayload {
    address receiver;
    bytes data;
    IERC20[] tokens;
    uint256[] amounts;
    address executor;
    bytes options;
  }

  /// @notice Report that is relayed by the observing DON at the relay phase
  struct RelayReport {
    bytes32 merkleRoot;
    uint256 minSequenceNumber;
    uint256 maxSequenceNumber;
  }
}
