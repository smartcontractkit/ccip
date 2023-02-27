// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../commitStore/CommitStore.sol";

contract CommitStoreHelper is CommitStore {
  constructor(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig,
    IAFN afn
  ) CommitStore(staticConfig, dynamicConfig, afn) {}

  /// @dev Expose _report for tests
  function report(bytes memory commitReport) external {
    _report(commitReport);
  }
}
