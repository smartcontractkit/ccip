// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../blobVerifier/BlobVerifier.sol";

contract BlobVerifierHelper is BlobVerifier {
  constructor(
    uint256 chainId,
    uint256 sourceChainId,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    BlobVerifierConfig memory config
  ) BlobVerifier(chainId, sourceChainId, afn, maxTimeWithoutAFNSignal, config) {}

  /**
   * @dev Expose _report for tests
   */
  function report(bytes memory rp) external {
    _report(bytes32(0), 0, rp);
  }
}
