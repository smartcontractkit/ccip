// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "./Common.sol";

// Library for CCIP internal definitions common to multiple contracts.
library Internal {
  function _addToTokensAmounts(
    Common.EVMTokenAndAmount[] memory existingTokens,
    Common.EVMTokenAndAmount memory newToken
  ) internal pure returns (Common.EVMTokenAndAmount[] memory) {
    for (uint256 i = 0; i < existingTokens.length; ++i) {
      if (existingTokens[i].token == newToken.token) {
        // already present, we need to create a new list because simply
        // incrementing the value will also mutate the original list.
        Common.EVMTokenAndAmount[] memory copyOfTokens = new Common.EVMTokenAndAmount[](existingTokens.length);
        for (uint256 j = 0; j < existingTokens.length; ++j) {
          copyOfTokens[j] = existingTokens[j];
        }

        copyOfTokens[i] = Common.EVMTokenAndAmount({
          token: copyOfTokens[i].token,
          amount: copyOfTokens[i].amount + newToken.amount
        });
        return copyOfTokens;
      }
    }

    // Token is not already present, need to reallocate.
    Common.EVMTokenAndAmount[] memory newTokens = new Common.EVMTokenAndAmount[](existingTokens.length + 1);
    for (uint256 i = 0; i < existingTokens.length; ++i) {
      newTokens[i] = existingTokens[i];
    }
    newTokens[existingTokens.length] = newToken;
    return newTokens;
  }

  // Offchain leaf domain separator
  bytes32 public constant LEAF_DOMAIN_SEPARATOR = 0x0000000000000000000000000000000000000000000000000000000000000000;
  // Internal domain separator used in proofs
  bytes32 public constant INTERNAL_DOMAIN_SEPARATOR =
    0x0000000000000000000000000000000000000000000000000000000000000001;

  /// @notice Bridge struct between message type specific struct and Any2EVMMessage for the end user.
  // TODO: Maybe not even needed? Just have separate arguments + an Any2EVMMessage to simplify?
  struct Any2EVMMessageFromSender {
    uint64 sourceChainId;
    bytes sender;
    address receiver;
    bytes data;
    // TODO consider another struct that contains pool, token and amount
    address[] destPools;
    Common.EVMTokenAndAmount[] destTokensAndAmounts;
    uint256 gasLimit;
  }

  function _toAny2EVMMessage(Any2EVMMessageFromSender memory original)
    internal
    pure
    returns (Common.Any2EVMMessage memory message)
  {
    message = Common.Any2EVMMessage({
      sourceChainId: original.sourceChainId,
      sender: original.sender,
      data: original.data,
      destTokensAndAmounts: original.destTokensAndAmounts
    });
  }

  /// @notice a sequenceNumber interval
  struct Interval {
    uint64 min;
    uint64 max;
  }

  /// @notice Report that is committed by the observing DON at the committing phase
  struct CommitReport {
    address[] onRamps;
    Interval[] intervals;
    bytes32[] merkleRoots;
    bytes32 rootOfRoots;
  }

  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }

  struct ExecutionResult {
    uint64 sequenceNumber;
    MessageExecutionState state;
  }
}
