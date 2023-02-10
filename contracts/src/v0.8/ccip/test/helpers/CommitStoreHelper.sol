// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../commitStore/CommitStore.sol";

contract CommitStoreHelper is CommitStore {
  constructor(
    CommitStoreConfig memory config,
    IAFN afn,
    uint64 minSeqNr
  ) CommitStore(config, afn, minSeqNr) {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory commitReport) external {
    _report(commitReport);
  }
}
