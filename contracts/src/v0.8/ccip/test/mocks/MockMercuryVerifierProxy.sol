// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

contract MockMercuryVerifierProxy {
  /**
   * @notice Bulk verifies that the data encoded has been signed
   * correctly by routing to the correct verifier, and bills the user if applicable.
   * @param payloads The encoded payloads to be verified, including the signed
   * report.
   * @param parameterPayload fee metadata for billing
   * @return verifiedReports The encoded reports from the verifier.
   */
  function verifyBulk(
    bytes[] calldata payloads,
    bytes calldata parameterPayload
  ) external payable returns (bytes[] memory verifiedReports) {
    // always successfully verify for tests to pass
    return new bytes[](0);
  }
}
