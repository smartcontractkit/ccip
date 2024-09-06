// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMNV2} from "../interfaces/IRMNV2.sol";
import {Internal} from "../libraries/Internal.sol";

/// @dev XXX DO NOT USE THIS CONTRACT, FOR TESTING ONLY XXX
contract DummyRMN is IRMNV2 {
  /// @inheritdoc IRMNV2
  function verify(Internal.MerkleRoot[] memory destLaneUpdates, Signature[] memory signatures) external view {
    return;
  }

  /// @inheritdoc IRMNV2
  function isCursed() external view returns (bool) {
    return false;
  }

  /// @inheritdoc IRMNV2
  function isCursed(bytes16 subject) external view returns (bool) {
    return false;
  }
}
