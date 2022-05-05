// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../interfaces/PoolInterface.sol";
import "../interfaces/OffRampRouterInterface.sol";
import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface OffRampInterface {
  error RelayReportError();
  error SequenceError(uint64 lastMaxSequenceNumber, uint64 newMinSequenceNumber);
  error MerkleProofError(bytes32 root);
  error TokenMismatch();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error AlreadyExecuted(uint64 sequenceNumber);
  error InvalidExecutor(uint64 sequenceNumber);
  error ExecutionError(uint64 sequenceNumber, bytes reason);
  error FeeError();
  error ExecutionDelayError();
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnexpectedPayloadData(uint64 sequenceNumber);
  error RouterNotSet();

  event ReportAccepted(CCIP.RelayReport report);
  event CrossChainMessageExecuted(uint64 indexed sequenceNumber);
  event OffRampConfigSet(OffRampConfig config);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);
  event OffRampRouterSet(OffRampRouterInterface router);

  struct OffRampConfig {
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
   * @param report ExecutionReport
   * @param needFee Whether or not the executor requires a fee
   */
  function executeTransaction(CCIP.ExecutionReport memory report, bool needFee) external;
}
