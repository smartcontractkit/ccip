// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {CommitStoreInterface} from "../interfaces/CommitStoreInterface.sol";
import {HealthChecker, AFNInterface} from "../health/HealthChecker.sol";
import {OCR2Base} from "../ocr/OCR2Base.sol";
import {Internal} from "../models/Internal.sol";

contract CommitStore is CommitStoreInterface, TypeAndVersionInterface, HealthChecker, OCR2Base {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "CommitStore 1.0.0";

  // Chain ID of this chain
  uint256 public immutable i_chainId;
  // Chain ID of the source chain
  uint256 public immutable i_sourceChainId;

  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;

  // The CommitStore configuration values
  CommitStoreConfig private s_config;

  // Mapping of the expected next sequence number by onRamp
  mapping(address => uint64) private s_expectedNextMinByOnRamp;

  /**
   * @dev sourceTokens are mapped to pools, and therefore should be the same length arrays.
   * @dev The AFN contract should be deployed already
   * @param chainId The ID that this contract is deployed to
   * @param afn AFN contract
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
    CommitStoreConfig memory config
  ) OCR2Base(true) HealthChecker(afn) {
    i_chainId = chainId;
    i_sourceChainId = sourceChainId;
    s_config = config;
    if (s_config.onRamps.length != s_config.minSeqNrByOnRamp.length) revert InvalidConfiguration();
    for (uint256 i = 0; i < s_config.onRamps.length; ++i) {
      s_expectedNextMinByOnRamp[s_config.onRamps[i]] = s_config.minSeqNrByOnRamp[i];
    }
  }

  /// @inheritdoc CommitStoreInterface
  function setConfig(CommitStoreConfig calldata config) external onlyOwner {
    uint256 newRampLength = config.onRamps.length;
    if (newRampLength != config.minSeqNrByOnRamp.length || newRampLength == 0) revert InvalidConfiguration();
    uint256 onRampLength = s_config.onRamps.length;
    for (uint256 i = 0; i < onRampLength; ++i) {
      delete s_expectedNextMinByOnRamp[s_config.onRamps[i]];
    }

    s_config = config;
    for (uint256 i = 0; i < newRampLength; ++i) {
      s_expectedNextMinByOnRamp[config.onRamps[i]] = config.minSeqNrByOnRamp[i];
    }

    emit CommitStoreConfigSet(config);
  }

  /// @inheritdoc CommitStoreInterface
  function getConfig() external view returns (CommitStoreConfig memory) {
    return s_config;
  }

  /// @inheritdoc CommitStoreInterface
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

  function _hashCommitStoreWithRoot(bytes32 root) internal view returns (bytes32) {
    return keccak256(abi.encode(address(this), root));
  }

  function isBlessed(bytes32 root) public view returns (bool) {
    return s_afn.isBlessed(_hashCommitStoreWithRoot(root));
  }

  /// @inheritdoc CommitStoreInterface
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
    // Only return non-zero if present and blessed.
    if (s_roots[outerRoot] == 0 || !isBlessed(outerRoot)) {
      return uint256(0);
    }
    return s_roots[outerRoot];
  }

  /// @inheritdoc CommitStoreInterface
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
      if (totalHashes > 256) revert InvalidProof();
      bytes32[] memory hashes = new bytes32[](totalHashes);
      (uint256 leafPos, uint256 hashPos, uint256 proofPos) = (0, 0, 0);

      for (uint256 i = 0; i < totalHashes; ++i) {
        hashes[i] = _hashPair(
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

  /// @inheritdoc CommitStoreInterface
  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_roots[root];
  }

  /// @inheritdoc OCR2Base
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory encodedReport
  ) internal override whenNotPaused whenHealthy {
    Internal.CommitReport memory report = abi.decode(encodedReport, (Internal.CommitReport));
    uint256 reportLength = report.onRamps.length;
    if (reportLength != report.intervals.length || reportLength != report.merkleRoots.length)
      revert InvalidCommitReport(report);
    for (uint256 i = 0; i < reportLength; ++i) {
      address onRamp = report.onRamps[i];
      uint64 expectedMinSeqNum = s_expectedNextMinByOnRamp[onRamp];
      if (expectedMinSeqNum == 0) revert UnsupportedOnRamp(onRamp);
      Internal.Interval memory repInterval = report.intervals[i];

      if (expectedMinSeqNum != repInterval.min || repInterval.min > repInterval.max)
        revert InvalidInterval(repInterval, onRamp);
      s_expectedNextMinByOnRamp[onRamp] = repInterval.max + 1;
    }
    s_roots[report.rootOfRoots] = block.timestamp;
    emit ReportAccepted(report);
  }

  /**
   * @notice Hashes two bytes32 objects. The order is taken into account,
   *          using the lower value first.
   */
  function _hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
    return a < b ? _hashInternalNode(a, b) : _hashInternalNode(b, a);
  }

  /**
   * @notice Hashes two bytes32 objects in their given order, prepended by the
   *          INTERNAL_DOMAIN_SEPARATOR.
   */
  function _hashInternalNode(bytes32 left, bytes32 right) private pure returns (bytes32 hash) {
    return keccak256(abi.encode(Internal.INTERNAL_DOMAIN_SEPARATOR, left, right));
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
