// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMNRemote} from "./rmn/RMNRemote.sol";

/// @dev this file exposes structs that are otherwise internal to the CCIP codebase
/// doing this allows those structs to be encoded and decoded with type safety in offchain code
/// and tests because generated wrappers are available
contract CCIPEncodingUtils {
  constructor() {
    revert("do not deploy");
  }

  /// @dev the RMN Report struct is used in integration / E2E tests
  function _rmnReport(bytes32 rmnVersionString, RMNRemote.Report memory rmnReport) external {}
}
