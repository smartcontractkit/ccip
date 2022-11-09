// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/CommitStoreInterface.sol";

contract MockCommitStore is CommitStoreInterface {
  /// @inheritdoc CommitStoreInterface
  function getConfig() external pure returns (CommitStoreConfig memory) {
    CommitStoreConfig memory config;
    return config;
  }

  /// @inheritdoc CommitStoreInterface
  function setConfig(CommitStoreConfig calldata config) external {}

  /// @inheritdoc CommitStoreInterface
  function getExpectedNextSequenceNumber(address) external pure returns (uint64) {
    return 1;
  }

  /// @inheritdoc CommitStoreInterface
  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256,
    bytes32[] calldata,
    uint256
  ) external pure returns (uint256 timestamp) {
    return 1;
  }

  /// @inheritdoc CommitStoreInterface
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory,
    uint256
  ) public pure returns (bytes32) {
    return leaves[0];
  }

  /// @inheritdoc CommitStoreInterface
  function getMerkleRoot(bytes32) external pure returns (uint256) {
    return 1;
  }

  function isBlessed(bytes32) external pure returns (bool) {
    return true;
  }
}
