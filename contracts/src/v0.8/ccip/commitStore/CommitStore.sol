// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {ICommitStore} from "../interfaces/ICommitStore.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IPriceRegistry} from "../interfaces/prices/IPriceRegistry.sol";

import {OCR2Base} from "../ocr/OCR2Base.sol";
import {Internal} from "../models/Internal.sol";
import {Pausable} from "../../vendor/Pausable.sol";

contract CommitStore is ICommitStore, TypeAndVersionInterface, Pausable, OCR2Base {
  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "CommitStore 1.0.0";
  // Chain ID of this chain
  uint64 internal immutable i_chainId;
  // Chain ID of the source chain
  uint64 internal immutable i_sourceChainId;
  // The onRamp address on the source chain
  address internal immutable i_onRamp;

  // DYNAMIC CONFIG
  // The dynamic commitStore config
  DynamicConfig internal s_dynamicConfig;

  // STATE
  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;
  // The min sequence number expected for future messages
  uint64 private s_minSeqNr = 1;

  /// @param staticConfig Containing the static part of the commitStore config
  /// @param dynamicConfig Containing the dynamic part of the commitStore config
  constructor(StaticConfig memory staticConfig, DynamicConfig memory dynamicConfig) OCR2Base() Pausable() {
    if (
      dynamicConfig.priceRegistry == address(0) ||
      staticConfig.onRamp == address(0) ||
      staticConfig.chainId == 0 ||
      staticConfig.sourceChainId == 0
    ) revert InvalidCommitStoreConfig();

    i_chainId = staticConfig.chainId;
    i_sourceChainId = staticConfig.sourceChainId;
    i_onRamp = staticConfig.onRamp;

    _setDynamicConfig(dynamicConfig);
  }

  /// @notice Pause the contract
  /// @dev only callable by the owner
  function pause() external onlyOwner {
    _pause();
  }

  /// @notice Unpause the contract
  /// @dev only callable by the owner
  function unpause() external onlyOwner {
    _unpause();
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

  /// @notice returns a hash of the abi encoded address of this contract and the
  /// supplied root.
  function _hashCommitStoreWithRoot(bytes32 root) internal view returns (bytes32) {
    return keccak256(abi.encode(address(this), root));
  }

  /// @inheritdoc ICommitStore
  function isBlessed(bytes32 root) public view returns (bool) {
    return IAFN(s_dynamicConfig.afn).isBlessed(_hashCommitStoreWithRoot(root));
  }

  /// @inheritdoc ICommitStore
  function getStaticConfig() external view override returns (StaticConfig memory) {
    return ICommitStore.StaticConfig({chainId: i_chainId, sourceChainId: i_sourceChainId, onRamp: i_onRamp});
  }

  /// @inheritdoc ICommitStore
  function getDynamicConfig() external view override returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @inheritdoc ICommitStore
  function setDynamicConfig(DynamicConfig memory dynamicConfig) external override onlyOwner {
    _setDynamicConfig(dynamicConfig);
  }

  /// @notice the internal version of setDynamicConfig to allow for reuse
  /// in the constructor. Emits ConfigSet on successful config set.
  function _setDynamicConfig(DynamicConfig memory dynamicConfig) internal {
    if (dynamicConfig.afn == address(0) || dynamicConfig.priceRegistry == address(0)) revert InvalidCommitStoreConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(
      ICommitStore.StaticConfig({chainId: i_chainId, sourceChainId: i_sourceChainId, onRamp: i_onRamp}),
      dynamicConfig
    );
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

    if (report.priceUpdates.tokenPriceUpdates.length > 0 || report.priceUpdates.destChainId != 0) {
      IPriceRegistry(s_dynamicConfig.priceRegistry).updatePrices(report.priceUpdates);
      // If there is no root, the report only contained fee updated and
      // we return to not revert on the empty root check below.
      if (report.merkleRoot == bytes32(0)) {
        return;
      }
    }

    // If we reached this code the report should also contain a valid root
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

  /// @notice Support querying whether health checker is healthy.
  function isAFNHealthy() external view returns (bool) {
    return !IAFN(s_dynamicConfig.afn).badSignalReceived();
  }

  /// @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IAFN(s_dynamicConfig.afn).badSignalReceived()) revert BadAFNSignal();
    _;
  }
}
