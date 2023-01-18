// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../interfaces/ICommitStore.sol";

contract MockCommitStore is ICommitStore {
  /// @inheritdoc ICommitStore
  function getCommitStoreConfig() external pure returns (CommitStoreConfig memory) {
    CommitStoreConfig memory config;
    return config;
  }

  /// @inheritdoc ICommitStore
  function setCommitStoreConfig(CommitStoreConfig calldata config) external {}

  /// @inheritdoc ICommitStore
  function getExpectedNextSequenceNumber(address) external pure returns (uint64) {
    return 1;
  }

  function getChainId() external pure returns (uint256) {
    return 1;
  }

  function getSourceChainId() external pure returns (uint256) {
    return 2;
  }

  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256,
    bytes32[] calldata,
    uint256
  ) external pure returns (uint256 timestamp) {
    return 1;
  }

  /// @inheritdoc ICommitStore
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory,
    uint256
  ) public pure returns (bytes32) {
    return leaves[0];
  }

  /// @inheritdoc ICommitStore
  function getMerkleRoot(bytes32) external pure returns (uint256) {
    return 1;
  }

  function isBlessed(bytes32) external pure returns (bool) {
    return true;
  }
}
