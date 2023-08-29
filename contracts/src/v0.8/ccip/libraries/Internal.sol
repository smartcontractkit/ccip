// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Client} from "./Client.sol";
import {MerkleMultiProof} from "../libraries/MerkleMultiProof.sol";

// Library for CCIP internal definitions common to multiple contracts.
library Internal {
  struct PriceUpdates {
    TokenPriceUpdate[] tokenPriceUpdates;
    uint64 destChainSelector; // --┐ Destination chain selector
    uint224 usdPerUnitGas; // -----┘ 1e18 USD per smallest unit (e.g. wei) of destination chain gas
  }

  struct TokenPriceUpdate {
    address sourceToken; // Source token
    uint224 usdPerToken; // 1e18 USD per smallest unit of token
  }

  struct TimestampedPackedUint224 {
    uint224 value; // -------┐ Value in uint224, packed.
    uint32 timestamp; // ----┘ Timestamp of the most recent update.
  }

  struct PoolUpdate {
    address token; // The IERC20 token address
    address pool; // The token pool address
  }

  struct ExecutionReport {
    EVM2EVMMessage[] messages;
    // Contains a bytes array for each message
    // each inner bytes array contains bytes per transferred token
    bytes[][] offchainTokenData;
    bytes32[] proofs;
    uint256 proofFlagBits;
  }

  // @notice The cross chain message that gets committed to EVM chains
  struct EVM2EVMMessage {
    uint64 sourceChainSelector;
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


  /// @dev EVM2EVMMessage struct has 12 fields, including 2 variable arrays.
  /// Each variable array takes 1 more slot to store its length.
  /// Hence, when abiEncoded, excluding the dynamic contents of variable arrays,
  /// EVM2EVMMessage takes up a fixed number of 14 lots, 32 bytes each.
  uint256 public constant EVM_2_EVM_MESSAGE_FIXED_BYTES = 32 * 14;

  /// @dev Each token transfer adds 1 instance of EVMTokenAmount struct to EVM2EVMMessage.
  /// EVMTokenAmount struct contains 2 fields, 32 bytes each when abiEncoded.
  uint256 public constant EVM_2_EVM_MESSAGE_BYTES_PER_TOKEN = 32 * 2;

  function _toAny2EVMMessage(
    EVM2EVMMessage memory original,
    Client.EVMTokenAmount[] memory destTokenAmounts
  ) internal pure returns (Client.Any2EVMMessage memory message) {
    message = Client.Any2EVMMessage({
      messageId: original.messageId,
      sourceChainSelector: original.sourceChainSelector,
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
          MerkleMultiProof.LEAF_DOMAIN_SEPARATOR,
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

  /// @notice Enum listing the possible message execution states within
  /// the offRamp contract.
  /// UNTOUCHED never executed
  /// IN_PROGRESS currently being executed, used a replay protection
  /// SUCCESS successfully executed. End state
  /// FAILURE unsuccessfully executed, manual execution is now enabled.
  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }
}
