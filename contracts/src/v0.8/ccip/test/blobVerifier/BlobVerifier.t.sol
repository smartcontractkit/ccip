// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./BlobVerifierSetup.t.sol";
import "../helpers/MerkleHelper.sol";

contract BlobVerifierTest is BlobVerifierSetup {
  function setUp() public virtual override {
    BlobVerifierSetup.setUp();
  }

  function testReportSuccess() public {
    uint64 max1 = 931;
    uint64 max2 = 2;
    uint64 max3 = 15;
    CCIP.Interval[] memory intervals = new CCIP.Interval[](3);
    intervals[0] = CCIP.Interval(1, max1);
    intervals[1] = CCIP.Interval(1, max2);
    intervals[2] = CCIP.Interval(1, max3);
    bytes32[] memory merkleRoots = new bytes32[](3);
    merkleRoots[0] = "test #1";
    merkleRoots[1] = "test #2";
    merkleRoots[2] = "test #3";
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: s_config.onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });
    s_blobVerifier.report(abi.encode(report));

    assertEq(max1 + 1, s_blobVerifier.s_expectedNextMinByOnRamp(s_config.onRamps[0]));
    assertEq(max2 + 1, s_blobVerifier.s_expectedNextMinByOnRamp(s_config.onRamps[1]));
    assertEq(max3 + 1, s_blobVerifier.s_expectedNextMinByOnRamp(s_config.onRamps[2]));
  }

  function testReportAndVerifySingleRamp() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(1, 2);
    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = "rootAndAlsoRootOfRoots";
    address[] memory onRamps = new address[](1);
    onRamps[0] = s_config.onRamps[0];
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: merkleRoots[0]
    });
    s_blobVerifier.report(abi.encode(report));
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_blobVerifier.verify(merkleRoots, proofs, 2**1, proofs, 2**1);
    assertEq(s_blockTime, timestamp);
  }
}
