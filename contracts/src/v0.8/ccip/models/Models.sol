// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library CCIP {
  ////////////////////////////////
  ////         COMMON         ////
  ////////////////////////////////

  // Offchain leaf domain separator
  bytes32 public constant LEAF_DOMAIN_SEPARATOR = 0x0000000000000000000000000000000000000000000000000000000000000000;
  // Internal domain separator used in proofs
  bytes32 public constant INTERNAL_DOMAIN_SEPARATOR =
    0x0000000000000000000000000000000000000000000000000000000000000001;

  struct EVMTokenAndAmount {
    address token;
    uint256 amount;
  }

  /// @notice Generalized EVM message type that is sent from EVM routers
  // to the contracts that implement the Any2EVMMessageReceiverInterface
  struct Any2EVMMessageFromSender {
    uint256 sourceChainId;
    bytes sender;
    address receiver;
    bytes data;
    // TODO consider another struct that contains pool, token and amount
    address[] destPools;
    EVMTokenAndAmount[] destTokensAndAmounts;
    uint256 gasLimit;
  }

  struct Any2EVMMessage {
    uint256 sourceChainId;
    bytes sender;
    bytes data;
    EVMTokenAndAmount[] destTokensAndAmounts;
  }

  function _toAny2EVMMessage(CCIP.Any2EVMMessageFromSender memory original)
    internal
    pure
    returns (CCIP.Any2EVMMessage memory message)
  {
    message = CCIP.Any2EVMMessage({
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

  struct FeeUpdate {
    uint256 chainId;
    uint256 gasPrice;
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

  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);

  struct EVMExtraArgsV1 {
    uint256 gasLimit;
    bool strict;
  }

  // bytes4(keccak256("CCIP EVMExtraArgsV1"));
  bytes4 public constant EVM_EXTRA_ARGS_V1_TAG = 0x97a657c9;
  uint256 public constant EVM_DEFAULT_GAS_LIMIT = 200_000;

  function _toBytes(EVMExtraArgsV1 memory extraArgs) internal pure returns (bytes memory bts) {
    return abi.encodeWithSelector(EVM_EXTRA_ARGS_V1_TAG, extraArgs);
  }

  function _fromBytes(bytes calldata extraArgs) internal pure returns (EVMExtraArgsV1 memory) {
    if (extraArgs.length == 0) {
      return CCIP.EVMExtraArgsV1({gasLimit: EVM_DEFAULT_GAS_LIMIT, strict: false});
    }
    if (bytes4(extraArgs[:4]) != EVM_EXTRA_ARGS_V1_TAG)
      revert InvalidExtraArgsTag(EVM_EXTRA_ARGS_V1_TAG, bytes4(extraArgs[:4]));
    return CCIP.EVMExtraArgsV1({gasLimit: abi.decode(extraArgs[4:36], (uint256)), strict: false});
  }

  function _addToTokensAmounts(EVMTokenAndAmount[] memory existingTokens, EVMTokenAndAmount memory newToken)
    internal
    pure
    returns (EVMTokenAndAmount[] memory)
  {
    for (uint256 i = 0; i < existingTokens.length; ++i) {
      if (existingTokens[i].token == newToken.token) {
        // already present, we need to create a new list because simply
        // incrementing the value will also mutate the original list.
        EVMTokenAndAmount[] memory copyOfTokens = new EVMTokenAndAmount[](existingTokens.length);
        for (uint256 j = 0; j < existingTokens.length; ++j) {
          copyOfTokens[j] = existingTokens[j];
        }

        copyOfTokens[i] = EVMTokenAndAmount({
          token: copyOfTokens[i].token,
          amount: copyOfTokens[i].amount + newToken.amount
        });
        return copyOfTokens;
      }
    }

    // Token is not already present, need to reallocate.
    EVMTokenAndAmount[] memory newTokens = new EVMTokenAndAmount[](existingTokens.length + 1);
    for (uint256 i = 0; i < existingTokens.length; ++i) {
      newTokens[i] = existingTokens[i];
    }
    newTokens[existingTokens.length] = newToken;
    return newTokens;
  }

  ////////////////////////////////
  ////          GE            ////
  ////////////////////////////////

  /// @notice The GE message type for EVM chains.
  struct EVM2AnyGEMessage {
    bytes receiver;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
    address feeToken;
    bytes extraArgs;
  }

  // @notice The cross chain message that gets committed to EVM GE chains
  struct EVM2EVMGEMessage {
    uint64 sequenceNumber;
    uint256 feeTokenAmount;
    address sender;
    uint64 nonce;
    uint256 gasLimit;
    bool strict;
    uint256 sourceChainId;
    // User fields
    address receiver;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
    address feeToken;
  }

  bytes32 internal constant EVM_2_EVM_GE_MESSAGE_HASH = keccak256("EVM2EVMGEMessagePlus");

  function _hash(CCIP.EVM2EVMGEMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
    return
      keccak256(
        abi.encode(
          LEAF_DOMAIN_SEPARATOR,
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

  ////////////////////////////////
  ////          TOLL          ////
  ////////////////////////////////

  /// @notice The Toll message type for EVM chains.
  struct EVM2AnyTollMessage {
    bytes receiver;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
    EVMTokenAndAmount feeTokenAndAmount;
    bytes extraArgs;
  }

  // @notice The cross chain message that gets committed to EVM toll chains
  struct EVM2EVMTollMessage {
    uint256 sourceChainId;
    uint64 sequenceNumber;
    address sender;
    address receiver;
    bytes data;
    EVMTokenAndAmount[] tokensAndAmounts;
    EVMTokenAndAmount feeTokenAndAmount;
    uint256 gasLimit;
  }

  bytes32 internal constant EVM_2_EVM_TOLL_MESSAGE_HASH = keccak256("EVM2EVMTollMessagePlus");

  function _hash(CCIP.EVM2EVMTollMessage memory original, bytes32 metadataHash) internal pure returns (bytes32) {
    return
      keccak256(
        abi.encode(
          LEAF_DOMAIN_SEPARATOR,
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
}
