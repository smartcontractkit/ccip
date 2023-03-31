// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {Pausable} from "../../../vendor/Pausable.sol";

contract MockCommitStore is ICommitStore, Pausable {
  /// @inheritdoc ICommitStore
  function verify(
    bytes32[] calldata,
    bytes32[] calldata,
    uint256
  ) external view whenNotPaused returns (uint256 timestamp) {
    return 1;
  }

  function pause() external {
    _pause();
  }
}
