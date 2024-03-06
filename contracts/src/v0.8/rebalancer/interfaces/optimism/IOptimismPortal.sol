// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {Types} from "./Types.sol";

interface IOptimismPortal {
  /// @notice Semantic version.
  function version() external view returns (string memory);

  /// @notice Proves a withdrawal transaction.
  /// @param _tx              Withdrawal transaction to finalize.
  /// @param _l2OutputIndex   L2 output index to prove against.
  /// @param _outputRootProof Inclusion proof of the L2ToL1MessagePasser contract's storage root.
  /// @param _withdrawalProof Inclusion proof of the withdrawal in L2ToL1MessagePasser contract.
  function proveWithdrawalTransaction(
    Types.WithdrawalTransaction memory _tx,
    uint256 _l2OutputIndex,
    Types.OutputRootProof calldata _outputRootProof,
    bytes[] calldata _withdrawalProof
  ) external;
}
