// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

contract CCIPReaderTester {
  struct SourceChainConfig {
    bool isEnabled;
    uint64 minSeqNr;
    bytes onRamp;
  }

  struct EVM2AnyRampMessage {
    RampMessageHeader header;
    address sender;
  }

  struct RampMessageHeader {
    bytes32 messageId;
    uint64 sourceChainSelector;
    uint64 destChainSelector;
    uint64 sequenceNumber;
    uint64 nonce;
  }

  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) internal s_sourceChainConfigs;

  function getSourceChainConfig(uint64 sourceChainSelector) external view returns (SourceChainConfig memory) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  function setSourceChainConfig(uint64 sourceChainSelector, SourceChainConfig memory sourceChainConfig) external {
    s_sourceChainConfigs[sourceChainSelector] = sourceChainConfig;
  }

  event CCIPSendRequested(uint64 indexed destChainSelector, EVM2AnyRampMessage message);

  function EmitCCIPSendRequested(uint64 destChainSelector, EVM2AnyRampMessage memory message) external {
    emit CCIPSendRequested(destChainSelector, message);
  }

  enum MessageExecutionState {
    UNTOUCHED,
    IN_PROGRESS,
    SUCCESS,
    FAILURE
  }

  event ExecutionStateChanged(
    uint64 indexed sourceChainSelector,
    uint64 indexed sequenceNumber,
    bytes32 indexed messageId,
    MessageExecutionState state,
    bytes returnData
  );

  function EmitExecutionStateChanged(
    uint64 sourceChainSelector,
    uint64 sequenceNumber,
    bytes32 messageId,
    MessageExecutionState state,
    bytes memory returnData
  ) external {
    emit ExecutionStateChanged(sourceChainSelector, sequenceNumber, messageId, state, returnData);
  }

  struct Interval {
    uint64 min;
    uint64 max;
  }

  struct MerkleRoot {
    uint64 sourceChainSelector;
    Interval interval;
    bytes32 merkleRoot;
  }

  struct GasPriceUpdate {
    uint64 destChainSelector;
    uint224 usdPerUnitGas;
  }

  struct TokenPriceUpdate {
    address sourceToken;
    uint224 usdPerToken;
  }

  struct PriceUpdates {
    TokenPriceUpdate[] tokenPriceUpdates;
    GasPriceUpdate[] gasPriceUpdates;
  }

  struct CommitReport {
    PriceUpdates priceUpdates;
    MerkleRoot[] merkleRoots;
  }

  event CommitReportAccepted(CommitReport report);

  function EmitCommitReportAccepted(CommitReport memory report) external {
    emit CommitReportAccepted(report);
  }
}
