// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "./Common.sol";
import {Internal} from "./Internal.sol";

library Toll {
  // @notice The cross chain message that gets committed to EVM toll chains
  struct EVM2EVMTollMessage {
    uint64 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    bytes data;
    Common.EVMTokenAndAmount[] tokensAndAmounts;
    Common.EVMTokenAndAmount feeTokenAndAmount;
    uint256 gasLimit;
  }

  function _toAny2EVMMessage(EVM2EVMTollMessage memory original, Common.EVMTokenAndAmount[] memory destTokensAndAmounts)
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

  bytes32 internal constant EVM_2_EVM_TOLL_MESSAGE_HASH = keccak256("EVM2EVMTollMessageEvent");

  function _hash(Toll.EVM2EVMTollMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
    return
      keccak256(
        abi.encode(
          Internal.LEAF_DOMAIN_SEPARATOR,
          metadataHash,
          original.sequenceNumber,
          original.sender,
          original.receiver,
          keccak256(original.data),
          keccak256(abi.encode(original.tokensAndAmounts)),
          original.gasLimit,
          original.feeTokenAndAmount
        )
      );
  }

  struct ExecutionReport {
    uint64[] sequenceNumbers;
    address[] tokenPerFeeCoinAddresses;
    uint256[] tokenPerFeeCoin;
    bytes[] encodedMessages;
    bytes32[] innerProofs;
    uint256 innerProofFlagBits;
    bytes32[] outerProofs;
    uint256 outerProofFlagBits;
  }
}
