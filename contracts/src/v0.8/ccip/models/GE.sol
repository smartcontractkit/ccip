// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {GEConsumer} from "./GEConsumer.sol";
import {Common} from "./Common.sol";
import {Internal} from "./Internal.sol";

// GE message type specific structs.
library GE {
  struct FeeUpdate {
    address token;
    uint64 chainId;
    uint128 linkPerUnitGas;
  }

  struct ExecutionReport {
    uint64[] sequenceNumbers;
    address[] tokenPerFeeCoinAddresses;
    uint256[] tokenPerFeeCoin;
    FeeUpdate[] feeUpdates;
    bytes[] encodedMessages;
    bytes32[] innerProofs;
    uint256 innerProofFlagBits;
    bytes32[] outerProofs;
    uint256 outerProofFlagBits;
  }

  // @notice The cross chain message that gets committed to EVM GE chains
  struct EVM2EVMGEMessage {
    uint64 sourceChainId;
    uint64 sequenceNumber;
    uint256 feeTokenAmount;
    address sender;
    uint64 nonce;
    uint256 gasLimit;
    bool strict;
    // User fields
    address receiver;
    bytes data;
    Common.EVMTokenAndAmount[] tokensAndAmounts;
    address feeToken;
    bytes32 messageId;
  }

  function _toAny2EVMMessage(EVM2EVMGEMessage memory original, Common.EVMTokenAndAmount[] memory destTokensAndAmounts)
    internal
    pure
    returns (Common.Any2EVMMessage memory message)
  {
    message = Common.Any2EVMMessage({
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      data: original.data,
      destTokensAndAmounts: destTokensAndAmounts
    });
  }

  bytes32 internal constant EVM_2_EVM_GE_MESSAGE_HASH = keccak256("EVM2EVMGEMessageEvent");

  function _hash(GE.EVM2EVMGEMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
    return
      keccak256(
        abi.encode(
          Internal.LEAF_DOMAIN_SEPARATOR,
          metadataHash,
          original.sequenceNumber,
          original.nonce,
          original.sender,
          original.receiver,
          keccak256(original.data),
          keccak256(abi.encode(original.tokensAndAmounts)),
          original.gasLimit,
          original.strict,
          original.feeToken,
          original.feeTokenAmount
        )
      );
  }
}
