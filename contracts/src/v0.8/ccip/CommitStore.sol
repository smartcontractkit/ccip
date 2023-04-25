// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {ICommitStore} from "./interfaces/ICommitStore.sol";
import {IAFN} from "./interfaces/IAFN.sol";
import {IPriceRegistry} from "./interfaces/IPriceRegistry.sol";

import {OCR2Base} from "./ocr/OCR2Base.sol";
import {Internal} from "./libraries/Internal.sol";
import {Pausable} from "../vendor/Pausable.sol";
import {MerkleMultiProof} from "./libraries/MerkleMultiProof.sol";

contract CommitStore is ICommitStore, TypeAndVersionInterface, Pausable, OCR2Base {
  error InvalidInterval(Interval interval);
  error InvalidRoot();
  error InvalidCommitStoreConfig();
  error BadAFNSignal();

  event ReportAccepted(CommitReport report);
  event ConfigSet(StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event RootRemoved(bytes32 root);

  /// @notice Static commit store config
  struct StaticConfig {
    uint64 chainId; // -------┐  Destination chain Id
    uint64 sourceChainId; // -┘  Source chain Id
    address onRamp; //           OnRamp address on the source chain
  }

  /// @notice Dynamic commit store config
  struct DynamicConfig {
    address priceRegistry; // Price registry address on the destination chain
    address afn; // AFN
  }

  /// @notice a sequenceNumber interval
  struct Interval {
    uint64 min; // ---┐ Minimum sequence number, inclusive
    uint64 max; // ---┘ Maximum sequence number, inclusive
  }

  /// @notice Report that is committed by the observing DON at the committing phase
  struct CommitReport {
    Internal.PriceUpdates priceUpdates;
    Interval interval;
    bytes32 merkleRoot;
  }

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
  // The min sequence number expected for future messages
  uint64 private s_minSeqNr = 1;
  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;

  /// @param staticConfig Containing the static part of the commitStore config
  constructor(StaticConfig memory staticConfig) OCR2Base() Pausable() {
    if (staticConfig.onRamp == address(0) || staticConfig.chainId == 0 || staticConfig.sourceChainId == 0)
      revert InvalidCommitStoreConfig();

    i_chainId = staticConfig.chainId;
    i_sourceChainId = staticConfig.sourceChainId;
    i_onRamp = staticConfig.onRamp;
  }

  // ================================================================
  // |                        Verification                          |
  // ================================================================

  /// @notice Returns the next expected sequence number.
  /// @return the next expected sequenceNumber.
  function getExpectedNextSequenceNumber() public view returns (uint64) {
    return s_minSeqNr;
  }

  /// @notice Sets the minimum sequence number.
  /// @param minSeqNr The new minimum sequence number
  function setMinSeqNr(uint64 minSeqNr) external onlyOwner {
    s_minSeqNr = minSeqNr;
  }

  /// @notice Returns the timestamp of a potentially previously committed merkle root.
  /// If the root was never committed 0 will be returned.
  /// @param root The merkle root to check the commit status for.
  /// @return the timestamp of the committed root or zero in the case that it was never
  /// committed.
  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_roots[root];
  }

  /// @notice Returns if a root is blessed or not.
  /// @param root The merkle root to check the blessing status for.
  /// @return whether the root is blessed or not.
  function isBlessed(bytes32 root) public view returns (bool) {
    return IAFN(s_dynamicConfig.afn).isBlessed(_hashCommitStoreWithRoot(root));
  }

  /// @notice Used by the owner in case an invalid sequence of roots has been
  /// posted and needs to be removed. The interval in the report is trusted.
  function resetUnblessedRoots(bytes32[] calldata rootToReset) external onlyOwner {
    for (uint256 i = 0; i < rootToReset.length; ++i) {
      bytes32 root = rootToReset[i];
      if (!isBlessed(root)) {
        delete s_roots[root];
        emit RootRemoved(root);
      }
    }
  }

  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata proofs,
    uint256 proofFlagBits
  ) external view override whenNotPaused returns (uint256 timestamp) {
    bytes32 root = MerkleMultiProof.merkleRoot(hashedLeaves, proofs, proofFlagBits);
    // Only return non-zero if present and blessed.
    if (s_roots[root] == 0 || !isBlessed(root)) {
      return uint256(0);
    }
    return s_roots[root];
  }

  /// @inheritdoc OCR2Base
  function _report(bytes memory encodedReport) internal override whenNotPaused whenHealthy {
    CommitReport memory report = abi.decode(encodedReport, (CommitReport));

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

  /// @notice returns a hash of the abi encoded address of this contract and the
  /// supplied root.
  function _hashCommitStoreWithRoot(bytes32 root) internal view returns (bytes32) {
    return keccak256(abi.encode(address(this), root));
  }

  // ================================================================
  // |                           Config                             |
  // ================================================================

  /// @notice Returns the static commit store config.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory) {
    return StaticConfig({chainId: i_chainId, sourceChainId: i_sourceChainId, onRamp: i_onRamp});
  }

  /// @notice Returns the dynamic commit store config.
  /// @return the configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory) {
    return s_dynamicConfig;
  }

  /// @notice Sets the dynamic config. This function is called during `setOCR2Config` flow
  function _beforeSetConfig(bytes memory onchainConfig) internal override {
    DynamicConfig memory dynamicConfig = abi.decode(onchainConfig, (DynamicConfig));

    if (dynamicConfig.afn == address(0) || dynamicConfig.priceRegistry == address(0)) revert InvalidCommitStoreConfig();

    s_dynamicConfig = dynamicConfig;

    emit ConfigSet(StaticConfig({chainId: i_chainId, sourceChainId: i_sourceChainId, onRamp: i_onRamp}), dynamicConfig);
  }

  // ================================================================
  // |                        Access and AFN                        |
  // ================================================================

  /// @notice Support querying whether health checker is healthy.
  function isAFNHealthy() external view returns (bool) {
    return !IAFN(s_dynamicConfig.afn).badSignalReceived();
  }

  /// @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
  modifier whenHealthy() {
    if (IAFN(s_dynamicConfig.afn).badSignalReceived()) revert BadAFNSignal();
    _;
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
}
