// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {Pausable} from "../../../vendor/Pausable.sol";

contract MockCommitStore is ICommitStore, Pausable {
  uint64 private s_expectedNextSequenceNumber = 1;

  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256
  ) external view whenNotPaused returns (uint256 timestamp) {
    return 1;
  }

  function getExpectedNextSequenceNumber() external view returns (uint64) {
    return s_expectedNextSequenceNumber;
  }

  function setExpectedNextSequenceNumber(uint64 nextSeqNum) external {
    s_expectedNextSequenceNumber = nextSeqNum;
  }

  function pause() external {
    _pause();
  }
}
