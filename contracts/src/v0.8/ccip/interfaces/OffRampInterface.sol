// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/PoolInterface.sol";
import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface OffRampInterface {
  error RelayReportError();
  error SequenceError(uint256 lastMaxSequenceNumber, uint256 newMinSequenceNumber);
  error MerkleProofError(bytes32[] proof, CCIP.Message message, uint256 index);
  error TokenMismatch();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error ExceedsTokenLimit(uint256 currentLimit, uint256 requested);
  error AlreadyExecuted(uint256 sequenceNumber);
  error InvalidExecutor(uint256 sequenceNumber);
  error ExecutionError(uint256 sequenceNumber, bytes reason);
  error ExecutionDelayError();
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);

  event ReportAccepted(CCIP.RelayReport report);
  event CrossChainMessageExecuted(uint256 indexed sequenceNumber);
  event ExecutionDelaySecondsSet(uint256 delay);
  event NewTokenBucketConstructed(uint256 rate, uint256 capacity, bool full);

  /**
   * @notice Execute the delivery of a message by using its merkle proof
   * @param proof Merkle proof
   * @param message Original message object
   * @param index Index of the message in the original tree
   */
  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external;
}
