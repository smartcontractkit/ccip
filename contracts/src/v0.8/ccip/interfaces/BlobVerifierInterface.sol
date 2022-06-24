// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "../interfaces/TollOffRampRouterInterface.sol";

interface BlobVerifierInterface {
  error UnSupportedOnRamp(address onRamp);
  error InvalidInterval(CCIP.Interval interval, address onRamp);
  error InvalidRelayReport(CCIP.RelayReport report);

  event ReportAccepted(CCIP.RelayReport report);
  event BlobVerifierConfigSet(BlobVerifierConfig config);

  struct BlobVerifierConfig {
    uint256 sourceChainId;
    address[] onRamps;
    uint64[] minSeqNrByOnRamp;
  }

  function setConfig(BlobVerifierConfig calldata config) external;

  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] calldata outerProofs,
    uint256 outerProofFlagBits
  ) external returns (uint256 timestamp);

  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) external pure returns (bytes32);
}
