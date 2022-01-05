// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../utils/TokenLimits.sol";
import "../health/HealthChecker.sol";
import "../../vendor/Address.sol";

contract SingleTokenOffRamp is OffRampInterface, TypeAndVersionInterface, HealthChecker, OCR2Base {
  using Address for address;
  using TokenLimits for TokenLimits.TokenBucket;

  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;
  // Token pool contract
  PoolInterface public immutable POOL;
  // Token contract
  IERC20 public immutable TOKEN;
  // Offchain leaf domain separator
  bytes1 private constant LEAF_DOMAIN_SEPARATOR = 0x00;
  // Internal domain separator used in proofs
  bytes1 private constant INTERNAL_DOMAIN_SEPARATOR = 0x01;
  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_merkleRoots;
  // sequenceNumber => executed
  mapping(uint256 => bool) private s_executed;
  // execution delay in seconds
  uint256 private s_executionDelaySeconds;
  // Last relay report
  CCIP.RelayReport private s_lastReport;
  // Token bucket for token rate limiting
  TokenLimits.TokenBucket private s_tokenBucket;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    IERC20 token,
    PoolInterface pool,
    uint256 tokenBucketRate,
    uint256 tokenBucketCapacity,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    uint256 executionDelaySeconds
  ) OCR2Base(true) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    if (pool.getToken() != token) revert TokenMismatch();
    SOURCE_CHAIN_ID = sourceChainId;
    CHAIN_ID = chainId;
    TOKEN = token;
    POOL = pool;
    s_tokenBucket = TokenLimits.constructTokenBucket(tokenBucketRate, tokenBucketCapacity, true);
    s_executionDelaySeconds = executionDelaySeconds;
  }

  /**
   * @notice Extending OCR2Base._report
   * @dev assumes the report is a bytes encoded bytes32 merkle root
   * @dev will be called by Chainlink nodes on transmit()
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override whenNotPaused whenHealthy {
    CCIP.RelayReport memory newRelayReport = abi.decode(report, (CCIP.RelayReport));
    // check that the sequence numbers make sense
    if (newRelayReport.minSequenceNumber > newRelayReport.maxSequenceNumber) revert RelayReportError();
    CCIP.RelayReport memory lastRelayReport = s_lastReport;
    // if this is not the first relay report, make sure the sequence numbers
    // are greater than the previous report.
    if (lastRelayReport.merkleRoot != bytes32(0)) {
      if (newRelayReport.minSequenceNumber != lastRelayReport.maxSequenceNumber + 1) {
        revert SequenceError(lastRelayReport.maxSequenceNumber, newRelayReport.minSequenceNumber);
      }
    }

    s_merkleRoots[newRelayReport.merkleRoot] = block.timestamp;
    s_lastReport = newRelayReport;
    emit ReportAccepted(newRelayReport);
  }

  /**
   * @notice Execute a specific payload
   * @param proof Merkle proof in the order bottom to top of the tree
   * @param message Message that is to be sent
   * @param index Index of the leaf
   * @dev Can be called by anyone
   */
  function executeTransaction(
    bytes32[] memory proof,
    CCIP.Message memory message,
    uint256 index
  ) external override whenNotPaused whenHealthy {
    // Verify merkle proof
    // The leaf offchain is keccak256(LEAF_DOMAIN_SEPARATOR || CrossChainSendRequested event data),
    // where the CrossChainSendRequested event data is abi.encode(CCIP.Message).
    bytes32 leaf = keccak256(abi.encodePacked(LEAF_DOMAIN_SEPARATOR, abi.encode(message)));

    // Get root from proof
    bytes32 root = generateMerkleRoot(proof, leaf, index);

    // Check that root has been relayed
    uint256 reportTimestamp = s_merkleRoots[root];
    if (reportTimestamp == 0) revert MerkleProofError(proof, message, index);

    // Execution delay
    if (reportTimestamp + s_executionDelaySeconds >= block.timestamp) revert ExecutionDelayError();

    // Disallow double-execution.
    if (s_executed[message.sequenceNumber]) revert AlreadyExecuted(message.sequenceNumber);

    // The transaction can only be executed by the designated executor, if one exists.
    if (message.payload.executor != address(0) && message.payload.executor != msg.sender)
      revert InvalidExecutor(message.sequenceNumber);

    // Validity checks for the message.
    _isWellFormed(message);

    // Avoid shooting ourselves in the foot by disallowing calls to some
    // privileged OffRamp function as OffRamp.
    // In the wild: https://rekt.news/polynetwork-rekt/
    _validateReceiver(message);

    // Mark as executed before external calls
    s_executed[message.sequenceNumber] = true;

    // Remove the tokens from the rate limiting bucket
    uint256 numberOfTokens = message.payload.amounts[0];
    if (!s_tokenBucket.remove(numberOfTokens)) revert ExceedsTokenLimit(s_tokenBucket.tokens, numberOfTokens);

    // Release tokens to receiver
    POOL.releaseOrMint(message.payload.receiver, message.payload.amounts[0]);

    // Try send the message, emit fulfillment error if fails
    try CrossChainMessageReceiverInterface(message.payload.receiver).receiveMessage(message) {
      emit CrossChainMessageExecuted(message.sequenceNumber);
    } catch (bytes memory reason) {
      revert ExecutionError(message.sequenceNumber, reason);
    }
  }

  /**
   * @notice Generate a Merkle Root from Proof.
   * @param proof Merkle proof in the order bottom to top of the tree
   * @param leaf bytes32 leaf hash
   * @param index Index of the leaf
   * @return bytes32 root generated by proof
   */
  function generateMerkleRoot(
    bytes32[] memory proof,
    bytes32 leaf,
    uint256 index
  ) public pure returns (bytes32) {
    bytes32 hash = leaf;

    for (uint256 i = 0; i < proof.length; i++) {
      bytes32 proofElement = proof[i];

      if (index % 2 == 0) {
        hash = keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, hash, proofElement));
      } else {
        hash = keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, proofElement, hash));
      }
      index = index / 2;
    }
    return hash;
  }

  /**
   * @notice Message receiver checks
   */
  function _validateReceiver(CCIP.Message memory message) private view {
    if (
      address(message.payload.receiver) == address(this) ||
      address(message.payload.receiver) == address(POOL) ||
      address(message.payload.receiver) == address(TOKEN) ||
      !address(message.payload.receiver).isContract()
    ) revert InvalidReceiver(message.payload.receiver);
  }

  function _isWellFormed(CCIP.Message memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.payload.tokens.length != 1 || message.payload.amounts.length != 1) revert UnsupportedNumberOfTokens();
    if (message.payload.tokens[0] != TOKEN) revert UnsupportedToken(message.payload.tokens[0]);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {
    // TODO
  }

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function configureTokenBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external onlyOwner {
    s_tokenBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewTokenBucketConstructed(rate, capacity, full);
  }

  function getTokenBucket() external view returns (TokenLimits.TokenBucket memory) {
    return s_tokenBucket;
  }

  function setExecutionDelaySeconds(uint256 executionDelaySeconds) external onlyOwner {
    s_executionDelaySeconds = executionDelaySeconds;
    emit ExecutionDelaySecondsSet(executionDelaySeconds);
  }

  function getExecutionDelaySeconds() external view returns (uint256) {
    return s_executionDelaySeconds;
  }

  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_merkleRoots[root];
  }

  function getExecuted(uint256 sequenceNumber) external view returns (bool) {
    return s_executed[sequenceNumber];
  }

  function getLastReport() external view returns (CCIP.RelayReport memory) {
    return s_lastReport;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "SingleTokenOffRamp 1.1.0";
  }
}
