// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../models/Models.sol";

interface CommitStoreInterface {
  error UnsupportedOnRamp(address onRamp);
  error InvalidInterval(CCIP.Interval interval, address onRamp);
  error InvalidRelayReport(CCIP.RelayReport report);
  error InvalidConfiguration();
  error InvalidProof();

  event ReportAccepted(CCIP.RelayReport report);
  event CommitStoreConfigSet(CommitStoreConfig config);

  struct CommitStoreConfig {
    address[] onRamps;
    uint64[] minSeqNrByOnRamp;
  }

  /**
   * @notice Gets the current configuration.
   * @return the currently configured CommitStoreConfig.
   */
  function getConfig() external view returns (CommitStoreConfig memory);

  /**
   * @notice Sets the new CommitStoreConfig and updates the s_expectedNextMinByOnRamp
   *      mapping. It will first blank the entire mapping and then input the new values.
   *      This means that any onRamp previously set but not included in the new config
   *      will be unsupported afterwards.
   * @param config The new configuration.
   */
  function setConfig(CommitStoreConfig calldata config) external;

  /**
   * @notice Returns the next expected sequence number for a given onRamp.
   * @param onRamp The onRamp for which to get the next sequence number.
   * @return the next expected sequenceNumber for the given onRamp.
   */
  function getExpectedNextSequenceNumber(address onRamp) external view returns (uint64);

  /**
   * @notice Returns timestamp of when root was accepted or -1 if verification fails.
   * @dev This method uses a merkle tree within a merkle tree, with the hashedLeaves,
   *        innerProofs and innerProofFlagBits being used to get the root of the inner
   *        tree. This root is then used as the singular leaf of the outer tree.
   */
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] calldata outerProofs,
    uint256 outerProofFlagBits
  ) external returns (uint256 timestamp);

  /**
   * @notice Generates a Merkle Root based on the given leaves, proofs and proofFlagBits.
   *          This method can proof multiple leaves at the same time.
   * @param leaves The leaf hashes of the merkle tree.
   * @param proofs The hashes to be used instead of a leaf hash when the proofFlagBits
   *          indicates a proof should be used.
   * @param proofFlagBits A single uint256 of which each bit indicates whether a leaf or
   *          a proof needs to be used in a hash operation.
   * @dev the maximum number of hash operations it set to 256. Any input that would require
   *          more than 256 hashes to get to a root will revert.
   * @dev For given input `leaves` = [a,b,c] `proofs` = [D] and `proofFlagBits` = 5
   *     totalHashes = 3 + 1 - 1 = 3
   *  ** round 1 **
   *     proofFlagBits = (5 >> 0) & 1 = true
   *     hashes[0] = hashPair(a, b)
   *     (leafPos, hashPos, proofPos) = (2, 0, 0);
   *
   *  ** round 2 **
   *     proofFlagBits = (5 >> 1) & 1 = false
   *     hashes[1] = hashPair(D, c)
   *     (leafPos, hashPos, proofPos) = (3, 0, 1);
   *
   *  ** round 3 **
   *     proofFlagBits = (5 >> 2) & 1 = true
   *     hashes[2] = hashPair(hashed[0], hashes[1])
   *     (leafPos, hashPos, proofPos) = (3, 2, 1);
   *
   *     i = 3 and no longer < totalHashes. The algorithm is done
   *     return hashes[totalHashes - 1] = hashes[2]; the last hash we computed.
   */
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) external pure returns (bytes32);

  /**
   * @notice Returns the timestamp of a potentially previously relayed merkle root. If
   *          the root was never relayed 0 will be returned.
   * @param root The merkle root to check the relay status for.
   * @return the timestamp of the relayed root or zero in the case that it was never
   *          relayed.
   */
  function getMerkleRoot(bytes32 root) external view returns (uint256);

  /**
   * @notice Returns if a root is blessed or not.
   * @param root The merkle root to check the blessing status for.
   * @return whether the root is blessed or not.
   */
  function isBlessed(bytes32 root) external view returns (bool);
}
