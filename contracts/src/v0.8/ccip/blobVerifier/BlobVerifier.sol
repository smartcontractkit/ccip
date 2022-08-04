// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/Address.sol";
import "../../vendor/SafeERC20.sol";
import "../health/HealthChecker.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "./interfaces/BlobVerifierInterface.sol";

contract BlobVerifier is BlobVerifierInterface, TypeAndVersionInterface, HealthChecker, OCR2Base {
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "BlobVerifier 1.0.0";

  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;
  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Offchain leaf domain separator
  bytes1 private constant LEAF_DOMAIN_SEPARATOR = 0x00;
  // Internal domain separator used in proofs
  bytes1 private constant INTERNAL_DOMAIN_SEPARATOR = 0x01;

  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;

  // The BlobVerifier configuration values
  BlobVerifierConfig private s_config;

  // Mapping of the expected next sequence number by onRamp
  mapping(address => uint64) private s_expectedNextMinByOnRamp;

  /**
   * @dev sourceTokens are mapped to pools, and therefore should be the same length arrays.
   * @dev The AFN contract should be deployed already
   * @param chainId The ID that this contract is deployed to
   * @param afn AFN contract
   * @param maxTimeWithoutAFNSignal An AFN config setting
   * @param config containing:
   * sourceChainId: the source chain ID
   * onRamps: the addresses of the connected onRamps on the source chain for when overwriting
   *    the s_expectedNextMinByOnRamp mapping
   * minSeqNrByOnRamp: the new values when overwriting the s_expectedNextMinByOnRamp mapping
   */
  constructor(
    uint256 chainId,
    uint256 sourceChainId,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    BlobVerifierConfig memory config
  ) OCR2Base(true) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    CHAIN_ID = chainId;
    SOURCE_CHAIN_ID = sourceChainId;
    s_config = config;
    if (s_config.onRamps.length != s_config.minSeqNrByOnRamp.length) {
      revert InvalidConfiguration();
    }
    for (uint256 i = 0; i < s_config.onRamps.length; ++i) {
      s_expectedNextMinByOnRamp[s_config.onRamps[i]] = s_config.minSeqNrByOnRamp[i];
    }
  }

  /// @inheritdoc BlobVerifierInterface
  function setConfig(BlobVerifierConfig calldata config) external onlyOwner {
    uint256 newRampLength = config.onRamps.length;
    if (newRampLength != config.minSeqNrByOnRamp.length || newRampLength == 0) {
      revert InvalidConfiguration();
    }
    uint256 onRampLength = s_config.onRamps.length;
    for (uint256 i = 0; i < onRampLength; ++i) {
      delete s_expectedNextMinByOnRamp[s_config.onRamps[i]];
    }

    s_config = config;
    for (uint256 i = 0; i < newRampLength; ++i) {
      s_expectedNextMinByOnRamp[config.onRamps[i]] = config.minSeqNrByOnRamp[i];
    }

    emit BlobVerifierConfigSet(config);
  }

  /// @inheritdoc BlobVerifierInterface
  function getConfig() external view returns (BlobVerifierConfig memory) {
    return s_config;
  }

  /// @inheritdoc BlobVerifierInterface
  function getExpectedNextSequenceNumber(address onRamp) public view returns (uint64) {
    return s_expectedNextMinByOnRamp[onRamp];
  }

  /**
   * @notice Used by the owner in case an invalid sequence of roots has been
   * posted and needs to be removed. The interval in the report is trusted.
   */
  function resetUnblessedRoots(bytes32[] calldata rootToReset) external onlyOwner {
    for (uint256 i = 0; i < rootToReset.length; ++i) {
      // TODO: AFN check ( assert not self.afn.is_blessed(root))
      delete s_roots[rootToReset[i]];
    }
  }

  /// @inheritdoc BlobVerifierInterface
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] calldata outerProofs,
    uint256 outerProofFlagBits
  ) external view returns (uint256 timestamp) {
    bytes32[] memory outerLeaves = new bytes32[](1);
    // Use the result of the inner merkle proof as the single leaf of the outer merkle tree.
    outerLeaves[0] = merkleRoot(hashedLeaves, innerProofs, innerProofFlagBits);
    bytes32 outerRoot = merkleRoot(outerLeaves, outerProofs, outerProofFlagBits);
    return s_roots[outerRoot];
  }

  /// @inheritdoc BlobVerifierInterface
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) public pure returns (bytes32) {
    unchecked {
      uint256 leavesLen = leaves.length;
      // As of Solidity 0.6.5, overflow is not possible here because in-memory arrays are limited to
      // a max length of 2**64-1. Two uint64 values will not overflow a uint256.
      // See: https://blog.soliditylang.org/2020/04/06/memory-creation-overflow-bug/
      // Underflow is possible if leaves and proofs are empty, resulting in totalHashes = 2**256-1
      // This will be caught in the `require(totalHashes <= 256)` statement.
      uint256 totalHashes = leavesLen + proofs.length - 1;
      if (totalHashes == 0) {
        return leaves[0];
      }
      require(totalHashes <= 256);
      bytes32[] memory hashes = new bytes32[](totalHashes);
      (uint256 leafPos, uint256 hashPos, uint256 proofPos) = (0, 0, 0);

      for (uint256 i = 0; i < totalHashes; ++i) {
        hashes[i] = hashPair(
          // Checks if the bit flag signals the use of a supplied proof or a leaf/previous hash.
          ((proofFlagBits >> i) & uint256(1)) == 1
            ? (leafPos < leavesLen ? leaves[leafPos++] : hashes[hashPos++]) // Use a leaf or a previously computed hash
            : proofs[proofPos++], // Use a supplied proof.
          // The second part of the hashed pair is never a proof as hashing two proofs would result in a
          // hash that can already be computed offchain.
          leafPos < leavesLen ? leaves[leafPos++] : hashes[hashPos++]
        );
      }
      // Return the last hash.
      return hashes[totalHashes - 1];
    }
  }

  /// @inheritdoc BlobVerifierInterface
  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_roots[root];
  }

  /// @inheritdoc OCR2Base
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory encodedReport
  ) internal override whenNotPaused whenHealthy {
    CCIP.RelayReport memory report = abi.decode(encodedReport, (CCIP.RelayReport));
    uint256 reportLength = report.onRamps.length;
    if (reportLength != report.intervals.length || reportLength != report.merkleRoots.length) {
      revert InvalidRelayReport(report);
    }
    for (uint256 i = 0; i < reportLength; ++i) {
      address onRamp = report.onRamps[i];
      uint64 expectedMinSeqNum = s_expectedNextMinByOnRamp[onRamp];
      if (expectedMinSeqNum == 0) {
        revert UnsupportedOnRamp(onRamp);
      }
      CCIP.Interval memory repInterval = report.intervals[i];

      if (expectedMinSeqNum != repInterval.min || repInterval.min > repInterval.max) {
        revert InvalidInterval(repInterval, onRamp);
      }
      s_expectedNextMinByOnRamp[onRamp] = repInterval.max + 1;
    }
    s_roots[report.rootOfRoots] = block.timestamp;
    emit ReportAccepted(report);
  }

  /**
   * @notice Hashes two bytes32 objects. The order is taken into account,
   *          using the lower value first.
   */
  function hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
    return a < b ? _hashInternalNode(a, b) : _hashInternalNode(b, a);
  }

  /**
   * @notice Hashes two bytes32 objects in their given order, prepended by the
   *          INTERNAL_DOMAIN_SEPARATOR.
   */
  function _hashInternalNode(bytes32 left, bytes32 right) private pure returns (bytes32 hash) {
    return keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, left, right));
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
