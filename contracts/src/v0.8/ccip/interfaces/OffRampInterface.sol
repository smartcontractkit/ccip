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
  error UnexpectedPayloadData(uint256 sequenceNumber);

  event ReportAccepted(CCIP.RelayReport report);
  event CrossChainMessageExecuted(uint256 indexed sequenceNumber);
  event ExecutionDelaySecondsSet(uint64 delay);
  event ExecutionFeeLinkSet(uint64 executionFee);
  event MaxDataSizeSet(uint64 size);
  event MaxTokensLengthSet(uint64 length);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  struct OffRampConfig{
    // Execution fee in Juels (smallest denomination of LINK)
    uint64 executionFeeJuels;
    // execution delay in seconds
    uint64 executionDelaySeconds;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
  }

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
