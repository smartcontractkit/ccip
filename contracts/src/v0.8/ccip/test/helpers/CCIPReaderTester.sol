// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

contract CCIPReaderTester {
  struct SourceChainConfig {
    bool isEnabled;
    uint64 minSeqNr;
    bytes onRamp;
  }

  mapping(uint64 sourceChainSelector => SourceChainConfig sourceChainConfig) internal s_sourceChainConfigs;

  function getSourceChainConfig(uint64 sourceChainSelector) external view returns (SourceChainConfig memory) {
    return s_sourceChainConfigs[sourceChainSelector];
  }

  function setSourceChainConfig(uint64 sourceChainSelector, SourceChainConfig memory sourceChainConfig) external {
    s_sourceChainConfigs[sourceChainSelector] = sourceChainConfig;
  }
}
