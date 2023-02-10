// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {ICommitStore} from "../interfaces/ICommitStore.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";

import {HealthChecker} from "../health/HealthChecker.sol";
import {OCR2Base} from "../ocr/OCR2Base.sol";
import {Internal} from "../models/Internal.sol";

contract CommitStore is ICommitStore, TypeAndVersionInterface, HealthChecker, OCR2Base {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "CommitStore 1.0.0";

  // Chain ID of this chain
  uint64 internal immutable i_chainId;
  // Chain ID of the source chain
  uint64 internal immutable i_sourceChainId;
  // The onRamp address on the source chain
  address internal immutable i_onRamp;

  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;

  // The min sequence number expected for future messages
  uint64 private s_minSeqNr;

  /// @dev sourceTokens are mapped to pools, and therefore should be the same length arrays.
  /// @dev The AFN contract should be deployed already
  /// @param config containing the source and dest chain Ids and the onRamp
  /// @param afn AFN contract
  /// @param minSeqNr The expected minimum sequence number
  constructor(
    CommitStoreConfig memory config,
    IAFN afn,
    uint64 minSeqNr
  ) OCR2Base() HealthChecker(afn) {
    i_chainId = config.chainId;
    i_sourceChainId = config.sourceChainId;
    i_onRamp = config.onRamp;
    s_minSeqNr = minSeqNr;
  }

  /// @inheritdoc ICommitStore
  function setMinSeqNr(uint64 minSeqNr) external onlyOwner {
    s_minSeqNr = minSeqNr;
  }

  /// @inheritdoc ICommitStore
  function getExpectedNextSequenceNumber() public view returns (uint64) {
    return s_minSeqNr;
  }

  /// @notice Used by the owner in case an invalid sequence of roots has been
  /// posted and needs to be removed. The interval in the report is trusted.
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

  /// @inheritdoc ICommitStore
  function getConfig() external view override returns (ICommitStore.CommitStoreConfig memory) {
    return ICommitStore.CommitStoreConfig({chainId: i_chainId, sourceChainId: i_sourceChainId, onRamp: i_onRamp});
  }

  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata proofs,
    uint256 proofFlagBits
  ) external view override returns (uint256 timestamp) {
    bytes32 root = merkleRoot(hashedLeaves, proofs, proofFlagBits);
    // Only return non-zero if present and blessed.
    if (s_roots[root] == 0 || !isBlessed(root)) {
      return uint256(0);
    }
    return s_roots[root];
  }

  /// @inheritdoc ICommitStore
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) public pure override returns (bytes32) {
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

  /// @inheritdoc ICommitStore
  function getMerkleRoot(bytes32 root) external view override returns (uint256) {
    return s_roots[root];
  }

  /// @inheritdoc OCR2Base
  function _report(bytes memory encodedReport) internal override whenNotPaused whenHealthy {
    ICommitStore.CommitReport memory report = abi.decode(encodedReport, (ICommitStore.CommitReport));

    if (s_minSeqNr != report.interval.min || report.interval.min > report.interval.max)
      revert InvalidInterval(report.interval);

    if (report.merkleRoot == bytes32(0)) revert InvalidRoot();

    s_minSeqNr = report.interval.max + 1;
    s_roots[report.merkleRoot] = block.timestamp;
    emit ReportAccepted(report);
  }

  /// @notice Hashes two bytes32 objects. The order is taken into account,
  /// using the lower value first.
  function _hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
    return a < b ? _hashInternalNode(a, b) : _hashInternalNode(b, a);
  }

  /// @notice Hashes two bytes32 objects in their given order, prepended by the
  /// INTERNAL_DOMAIN_SEPARATOR.
  function _hashInternalNode(bytes32 left, bytes32 right) private pure returns (bytes32 hash) {
    return keccak256(abi.encode(Internal.INTERNAL_DOMAIN_SEPARATOR, left, right));
  }

  function _beforeSetOCR2Config(uint8 f, bytes memory onchainConfig) internal override {}
}
