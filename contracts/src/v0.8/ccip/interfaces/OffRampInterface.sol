// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/PoolInterface.sol";
import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface OffRampInterface {
  error RelayReportError();
  error SequenceError(uint256 lastMaxSequenceNumber, uint256 newMinSequenceNumber);
  error MerkleProofError(CCIP.MerkleProof proof, CCIP.Message message);
  error TokenMismatch();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error AlreadyExecuted(uint256 sequenceNumber);
  error InvalidExecutor(uint256 sequenceNumber);
  error ExecutionError(uint256 sequenceNumber, bytes reason);
  error FeeError();
  error ExecutionDelayError();
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);

  event ReportAccepted(CCIP.RelayReport report);
  event CrossChainMessageExecuted(uint256 indexed sequenceNumber);
  event ExecutionDelaySecondsSet(uint256 delay);
  event ExecutionFeeLinkSet(uint256 executionFee);
  event MaxDataSizeSet(uint256 size);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  /**
   * @notice Execute the delivery of a message by using its merkle proof
   * @param proof Merkle proof
   * @param message Original message object
   * @param needFee Whether or not the executor requires a fee
   */
  function executeTransaction(
    CCIP.Message memory message,
    CCIP.MerkleProof memory proof,
    bool needFee
  ) external;
}
