// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../CommitStore.sol";

contract CommitStoreHelper is CommitStore {
  constructor(StaticConfig memory staticConfig) CommitStore(staticConfig) {}

  /// @dev Expose _report for tests
  function report(bytes memory commitReport) external {
    _report(commitReport);
  }
}
