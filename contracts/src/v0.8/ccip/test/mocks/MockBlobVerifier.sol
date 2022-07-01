// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../blobVerifier/interfaces/BlobVerifierInterface.sol";

contract MockBlobVerifier is BlobVerifierInterface {
  function setConfig(BlobVerifierConfig calldata config) external {}

  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256,
    bytes32[] calldata,
    uint256
  ) external pure returns (uint256 timestamp) {
    return 1;
  }

  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory,
    uint256
  ) public pure returns (bytes32) {
    return leaves[0];
  }
}
