// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../offRamp/interfaces/Any2EVMTollOffRampRouterInterface.sol";
import "../../utils/CCIP.sol";

interface BlobVerifierInterface {
  error UnsupportedOnRamp(address onRamp);
  error InvalidInterval(CCIP.Interval interval, address onRamp);
  error InvalidRelayReport(CCIP.RelayReport report);
  error InvalidConfiguration();

  event ReportAccepted(CCIP.RelayReport report);
  event BlobVerifierConfigSet(BlobVerifierConfig config);

  struct BlobVerifierConfig {
    address[] onRamps;
    uint64[] minSeqNrByOnRamp;
  }

  /**
   * @notice Sets the new BlobVerifierConfig and updates the s_expectedNextMinByOnRamp
   *      mapping. It will first blank the entire mapping and then input the new values.
   *      This means that any onRamp previously set but not included in the new config
   *      will be unsupported afterwards.
   */
  function setConfig(BlobVerifierConfig calldata config) external;

  /**
   * @notice Verifies the given merkle tree input in a two-tiered merkle
   *          tree setup.
   */
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] calldata outerProofs,
    uint256 outerProofFlagBits
  ) external returns (uint256 timestamp);

  /**
   * @notice Computes the merkle root of a given set of leaves/proofs.
   * @dev This method can proof multiple leaves at the same time.
   */
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) external pure returns (bytes32);
}
