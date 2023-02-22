// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "./Client.sol";

// Library for CCIP internal definitions common to multiple contracts.
library Internal {
  struct PriceUpdates {
    FeeTokenPriceUpdate[] feeTokenPriceUpdates;
    uint64 destChainId; // ------┐ Destination chain Id
    uint128 usdPerUnitGas; // ---┘ USD per unit of destination chain gas
  }

  struct FeeTokenPriceUpdate {
    address sourceFeeToken;
    uint128 usdPerFeeToken;
  }

  struct ExecutionReport {
    uint64[] sequenceNumbers;
    bytes[] encodedMessages;
    bytes32[] proofs;
    uint256 proofFlagBits;
  }

  // @notice The cross chain message that gets committed to EVM chains
  struct EVM2EVMMessage {
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
    Client.EVMTokenAmount[] tokenAmounts;
    address feeToken;
    bytes32 messageId;
  }

  function _toAny2EVMMessage(EVM2EVMMessage memory original, Client.EVMTokenAmount[] memory destTokenAmounts)
    internal
    pure
    returns (Client.Any2EVMMessage memory message)
  {
    message = Client.Any2EVMMessage({
      messageId: original.messageId,
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      data: original.data,
      destTokenAmounts: destTokenAmounts
    });
  }

  bytes32 internal constant EVM_2_EVM_MESSAGE_HASH = keccak256("EVM2EVMMessageEvent");

  function _hash(EVM2EVMMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
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
          keccak256(abi.encode(original.tokenAmounts)),
          original.gasLimit,
          original.strict,
          original.feeToken,
          original.feeTokenAmount
        )
      );
  }

  function _addToTokensAmounts(Client.EVMTokenAmount[] memory existingTokens, Client.EVMTokenAmount memory newToken)
    internal
    pure
    returns (Client.EVMTokenAmount[] memory)
  {
    for (uint256 i = 0; i < existingTokens.length; ++i) {
      if (existingTokens[i].token == newToken.token) {
        // already present, we need to create a new list because simply
        // incrementing the value will also mutate the original list.
        Client.EVMTokenAmount[] memory copyOfTokens = new Client.EVMTokenAmount[](existingTokens.length);
        for (uint256 j = 0; j < existingTokens.length; ++j) {
          copyOfTokens[j] = existingTokens[j];
        }

        copyOfTokens[i] = Client.EVMTokenAmount({
          token: copyOfTokens[i].token,
          amount: copyOfTokens[i].amount + newToken.amount
        });
        return copyOfTokens;
      }
    }

    // Token is not already present, need to reallocate.
    Client.EVMTokenAmount[] memory newTokens = new Client.EVMTokenAmount[](existingTokens.length + 1);
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

  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }
}
