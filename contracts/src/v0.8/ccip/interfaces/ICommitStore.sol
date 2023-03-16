// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Internal} from "../models/Internal.sol";

interface ICommitStore {
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

  /// @notice Sets the minimum sequence number.
  /// @param minSeqNr The new minimum sequence number
  function setMinSeqNr(uint64 minSeqNr) external;

  /// @notice Returns the next expected sequence number.
  /// @return the next expected sequenceNumber.
  function getExpectedNextSequenceNumber() external view returns (uint64);

  /// @notice Returns timestamp of when root was accepted or -1 if verification fails.
  /// @dev This method uses a merkle tree within a merkle tree, with the hashedLeaves,
  /// proofs and proofFlagBits being used to get the root of the inner tree.
  /// This root is then used as the singular leaf of the outer tree.
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata proofs,
    uint256 proofFlagBits
  ) external returns (uint256 timestamp);

  /// @notice Returns the timestamp of a potentially previously committed merkle root.
  /// If the root was never committed 0 will be returned.
  /// @param root The merkle root to check the commit status for.
  /// @return the timestamp of the committed root or zero in the case that it was never
  /// committed.
  function getMerkleRoot(bytes32 root) external view returns (uint256);

  /// @notice Returns if a root is blessed or not.
  /// @param root The merkle root to check the blessing status for.
  /// @return whether the root is blessed or not.
  function isBlessed(bytes32 root) external view returns (bool);

  /// @notice Returns the static commit store config.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory);

  /// @notice Returns the dynamic commit store config.
  /// @return the configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory);

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The configuration.
  function setDynamicConfig(DynamicConfig memory dynamicConfig) external;
}
