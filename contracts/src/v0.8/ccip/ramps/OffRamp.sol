// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../health/HealthChecker.sol";
import "../pools/TokenPoolRegistry.sol";
import "../../vendor/Address.sol";
import "./PriceFeedRegistry.sol";
import "../../vendor/SafeERC20.sol";

contract OffRamp is
  OffRampInterface,
  TypeAndVersionInterface,
  HealthChecker,
  TokenPoolRegistry,
  PriceFeedRegistry,
  OCR2Base
{
  using Address for address;
  using SafeERC20 for IERC20;

  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;
  // Offchain leaf domain separator
  bytes1 private constant LEAF_DOMAIN_SEPARATOR = 0x00;
  // Internal domain separator used in proofs
  bytes1 private constant INTERNAL_DOMAIN_SEPARATOR = 0x01;
  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_merkleRoots;
  // sequenceNumber => executed
  mapping(uint64 => bool) private s_executed;
  // Last relay report
  CCIP.RelayReport private s_lastReport;

  // Configuration values
  OffRampConfig private s_config;
  // Router
  OffRampRouterInterface private s_router;

  /**
   * @dev sourceTokens are mapped to pools, and therefore should be the same length arrays.
   * @dev The AFN contract should be deployed already
   * @param sourceChainId The ID of the source chain
   * @param chainId The ID that this contract is deployed to
   * @param sourceTokens Array of source chain tokens that this contract supports
   * @param pools Array token token pools on this chain (Must map 1:1 with sourceTokens)
   * @param afn AFN contract
   * @param config containing:
   * - maxTimeWithoutAFNSignal Maximum number of seconds allows between AFN singals
   * - executionDelaySeconds Delay, in seconds, between the relay and execution of a message
   * - maxTokensLength The maximum number of different tokens allowed to be sent in a single message
   * - executionFeeJuels The execution fee, denominated in JUELS
   */
  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    OffRampConfig memory config
  )
    OCR2Base(true)
    HealthChecker(afn, maxTimeWithoutAFNSignal)
    TokenPoolRegistry(sourceTokens, pools)
    PriceFeedRegistry(sourceTokens, feeds)
  {
    SOURCE_CHAIN_ID = sourceChainId;
    CHAIN_ID = chainId;
    s_config = config;
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
   * @notice Execute the delivery of a message by using its merkle proof
   * @param report Execution report containing multi proofs
   * @param needFee Whether or not the executor requires a fee
   * @dev Can be called by anyone
   * @dev If the caller wishes to collect fees from the execution, needFee should be true.
   * This will send fee tokens directly to the executor address (msg.sender)
   */
  function executeTransaction(CCIP.ExecutionReport memory report, bool needFee)
    external
    override
    whenNotPaused
    whenHealthy
  {
    if (address(s_router) == address(0)) revert RouterNotSet();
    // Get root from path
    bytes32 root = merkleRoot(report);

    // Check that root has been relayed
    uint256 reportTimestamp = s_merkleRoots[root];
    if (reportTimestamp == 0) revert MerkleProofError(root);

    // Execution delay
    if (reportTimestamp + uint256(s_config.executionDelaySeconds) >= block.timestamp) revert ExecutionDelayError();

    for (uint256 i = 0; i < report.messages.length; i++) {
      CCIP.AnyToEVMTollMessage memory message = report.messages[i];

      // Disallow double-execution.
      if (s_executed[message.sequenceNumber]) revert AlreadyExecuted(message.sequenceNumber);

      // Validity checks for the message.
      _isWellFormed(message);

      // Avoid shooting ourselves in the foot by disallowing calls to some
      // privileged OffRamp function as OffRamp.
      // In the wild: https://rekt.news/polynetwork-rekt/
      _validateReceiver(message);

      // Mark as executed before external calls
      s_executed[message.sequenceNumber] = true;

      if (needFee) {
        uint256 fee = 0;
        IERC20 feeToken = message.tokens[0];
        AggregatorV2V3Interface feed = getFeed(feeToken);
        if (address(feed) == address(0)) revert FeeError();
        fee = uint256(s_config.executionFeeJuels) * uint256(feed.latestAnswer());
        if (fee > 0) {
          message.amounts[0] -= fee;
          _getPool(feeToken).releaseOrMint(msg.sender, fee);
        }
      }

      for (uint256 j = 0; j < message.tokens.length; j++) {
        // Release tokens to receiver
        _getPool(message.tokens[j]).releaseOrMint(message.receiver, message.amounts[j]);
      }

      // Try send the message, revert if fails
      if (message.receiver.isContract()) {
        try s_router.routeMessage(CrossChainMessageReceiverInterface(message.receiver), message) {} catch (
          bytes memory reason
        ) {
          // TODO: Figure out a better way to handle failed executions
          revert ExecutionError(message.sequenceNumber, reason);
        }
      } else {
        if (message.data.length > 0) {
          revert UnexpectedPayloadData(message.sequenceNumber);
        }
      }
      emit CrossChainMessageExecuted(message.sequenceNumber);
      // TODO: gas based fee calculation
    }
  }

  /**
   * @notice Generate a Merkle Root from an ExecutionReport
   * @param report ExecutionReport
   */
  function merkleRoot(CCIP.ExecutionReport memory report) public pure returns (bytes32) {
    uint256 leavesLen = report.messages.length;
    uint256 totalHashes = leavesLen + report.proofs.length - 1;
    require(totalHashes <= 256);
    unchecked {
      bytes32[] memory hashes = new bytes32[](totalHashes);
      uint256 leafPos = 0;
      uint256 hashPos = 0;
      uint256 proofPos = 0;
      for (uint256 i = 0; i < totalHashes; ++i) {
        bool proofFlag = ((report.proofFlagsBits >> i) & uint256(1)) == 1;
        hashes[i] = hashPair(
          proofFlag
            ? (leafPos < leavesLen ? _hashLeafNode(report.messages[leafPos++]) : hashes[hashPos++])
            : report.proofs[proofPos++],
          leafPos < leavesLen ? _hashLeafNode(report.messages[leafPos++]) : hashes[hashPos++]
        );
      }

      if (totalHashes > 0) {
        return hashes[totalHashes - 1];
      }
      return _hashLeafNode(report.messages[0]);
    }
  }

  function hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
    return a < b ? _hashInternalNode(a, b) : _hashInternalNode(b, a);
  }

  function _hashLeafNode(CCIP.AnyToEVMTollMessage memory message) private pure returns (bytes32) {
    // The hash offchain is keccak256(LEAF_DOMAIN_SEPARATOR || CrossChainSendRequested event data),
    // where the CrossChainSendRequested event data is abi.encode(CCIP.AnyToEVMTollMessage).
    return keccak256(abi.encodePacked(LEAF_DOMAIN_SEPARATOR, abi.encode(message)));
  }

  function _hashInternalNode(bytes32 left, bytes32 right) private pure returns (bytes32 hash) {
    return keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, left, right));
  }

  function _getPool(IERC20 token) private view returns (PoolInterface pool) {
    pool = getPool(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
  }

  /**
   * @notice Message receiver checks
   */
  function _validateReceiver(CCIP.AnyToEVMTollMessage memory message) private view {
    if (address(message.receiver) == address(this) || isPool(address(message.receiver)))
      revert InvalidReceiver(message.receiver);
  }

  function _isWellFormed(CCIP.AnyToEVMTollMessage memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length) {
      revert UnsupportedNumberOfTokens();
    }
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  /**
   * @notice TODO Withraw function that will be removed once transmitter renumeration is implemented
   */
  function withdrawAccumulatedFees(
    IERC20 feeToken,
    address recipient,
    uint256 amount
  ) external onlyOwner {
    feeToken.safeTransfer(recipient, amount);
    emit FeesWithdrawn(feeToken, recipient, amount);
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

  function setRouter(OffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(router);
  }

  function getRouter() external view returns (OffRampRouterInterface) {
    return s_router;
  }

  function setOffRampConfig(OffRampConfig calldata config) external onlyOwner {
    s_config = config;
    emit OffRampConfigSet(config);
  }

  function getOffRampConfig() external view returns (OffRampConfig memory) {
    return s_config;
  }

  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_merkleRoots[root];
  }

  function getExecuted(uint64 sequenceNumber) external view returns (bool) {
    return s_executed[sequenceNumber];
  }

  function getLastReport() external view returns (CCIP.RelayReport memory) {
    return s_lastReport;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "OffRamp 0.0.1";
  }
}
