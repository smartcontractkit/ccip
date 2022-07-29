// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../offRamp/mo/Any2EVMMOOffRamp.sol";

contract Any2EVMMOOffRampHelper is Any2EVMMOOffRamp {
  constructor(
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    address onRampAddress,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    uint256 maxTimeWithoutAFNSignal
  )
    Any2EVMMOOffRamp(
      chainId,
      offRampConfig,
      blobVerifier,
      onRampAddress,
      afn,
      sourceTokens,
      pools,
      maxTimeWithoutAFNSignal
    )
  {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }

  function setMessageState(uint64 sequenceNumber, CCIP.MessageExecutionState state) public {
    s_executedMessages[sequenceNumber] = state;
  }
}
