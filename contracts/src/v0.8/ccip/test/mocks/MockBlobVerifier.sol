// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/BlobVerifierInterface.sol";

contract MockBlobVerifier is BlobVerifierInterface {
  /// @inheritdoc BlobVerifierInterface
  function getConfig() external pure returns (BlobVerifierConfig memory) {
    BlobVerifierConfig memory config;
    return config;
  }

  /// @inheritdoc BlobVerifierInterface
  function setConfig(BlobVerifierConfig calldata config) external {}

  /// @inheritdoc BlobVerifierInterface
  function getExpectedNextSequenceNumber(address) external pure returns (uint64) {
    return 1;
  }

  /// @inheritdoc BlobVerifierInterface
  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256,
    bytes32[] calldata,
    uint256
  ) external pure returns (uint256 timestamp) {
    return 1;
  }

  /// @inheritdoc BlobVerifierInterface
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory,
    uint256
  ) public pure returns (bytes32) {
    return leaves[0];
  }

  /// @inheritdoc BlobVerifierInterface
  function getMerkleRoot(bytes32) external pure returns (uint256) {
    return 1;
  }

  function isBlessed(bytes32) external pure returns (bool) {
    return true;
  }
}
