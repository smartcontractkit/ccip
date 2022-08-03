// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./BlobVerifierSetup.t.sol";
import "../helpers/MerkleHelper.sol";

/// @notice #constructor
contract BlobVerifier_constructor is BaseTest {
  function testSuccess() public {
    address[] memory onRamps = new address[](3);
    onRamps[0] = ON_RAMP_ADDRESS;
    onRamps[1] = 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc;
    onRamps[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint64[] memory minSequenceNumbers = new uint64[](3);
    minSequenceNumbers[0] = 1;
    minSequenceNumbers[1] = 2;
    minSequenceNumbers[2] = 4;
    BlobVerifierInterface.BlobVerifierConfig memory config = BlobVerifierInterface.BlobVerifierConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSequenceNumbers
    });
    BlobVerifier blobVerifier = new BlobVerifier(DEST_CHAIN_ID, SOURCE_CHAIN_ID, s_afn, HEARTBEAT, config);

    // BlobVerifier config
    assertEq(minSequenceNumbers[0], blobVerifier.getExpectedNextSequenceNumber(onRamps[0]));
    assertEq(minSequenceNumbers[1], blobVerifier.getExpectedNextSequenceNumber(onRamps[1]));
    assertEq(minSequenceNumbers[2], blobVerifier.getExpectedNextSequenceNumber(onRamps[2]));

    BlobVerifierInterface.BlobVerifierConfig memory contractConfig = blobVerifier.getConfig();
    assertEq(keccak256(abi.encode(config.minSeqNrByOnRamp)), keccak256(abi.encode(contractConfig.minSeqNrByOnRamp)));
    assertEq(config.onRamps, contractConfig.onRamps);

    // typeAndVersion
    assertEq("BlobVerifier 1.0.0", blobVerifier.typeAndVersion());

    // owner
    assertEq(OWNER, blobVerifier.owner());

    // HealthChecker
    assertEq(HEARTBEAT, blobVerifier.getMaxSecondsWithoutAFNHeartbeat());
    assertEq(address(s_afn), address(blobVerifier.getAFN()));
  }

  function testInvalidConfigurationReverts() public {
    address[] memory onRamps = new address[](3);
    uint64[] memory minSequenceNumbers = new uint64[](2);

    vm.expectRevert(BlobVerifierInterface.InvalidConfiguration.selector);

    new BlobVerifier(
      DEST_CHAIN_ID,
      SOURCE_CHAIN_ID,
      s_afn,
      HEARTBEAT,
      BlobVerifierInterface.BlobVerifierConfig({onRamps: onRamps, minSeqNrByOnRamp: minSequenceNumbers})
    );
  }
}

/// @notice #setConfig
contract BlobVerifier_setConfig is BlobVerifierSetup {
  // Success

  function testSuccess() public {
    address[] memory onRamps = new address[](1);
    onRamps[0] = address(1);
    uint64[] memory minSeqNrByOnRamp = new uint64[](1);
    minSeqNrByOnRamp[0] = 200;
    BlobVerifierInterface.BlobVerifierConfig memory newConfig = BlobVerifierInterface.BlobVerifierConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    // Assert the current value for ON_RAMP_ADDRESS is set
    assertEq(1, s_blobVerifier.getExpectedNextSequenceNumber(ON_RAMP_ADDRESS));

    vm.expectEmit(false, false, false, false);
    emit BlobVerifierConfigSet(newConfig);

    s_blobVerifier.setConfig(newConfig);

    // Checks whether the new onramp is properly set to the given value
    assertEq(minSeqNrByOnRamp[0], s_blobVerifier.getExpectedNextSequenceNumber(onRamps[0]));
    // Assert the previously checked value is now 0, indicating successful removal
    // from the supported onRamps list.
    assertEq(0, s_blobVerifier.getExpectedNextSequenceNumber(ON_RAMP_ADDRESS));
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    BlobVerifierInterface.BlobVerifierConfig memory newConfig;
    s_blobVerifier.setConfig(newConfig);
  }

  function testInvalidConfigurationLengthMismatchReverts() public {
    address[] memory onRamps = new address[](2);
    uint64[] memory minSeqNrByOnRamp = new uint64[](1);
    BlobVerifierInterface.BlobVerifierConfig memory newConfig = BlobVerifierInterface.BlobVerifierConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    vm.expectRevert(BlobVerifierInterface.InvalidConfiguration.selector);

    s_blobVerifier.setConfig(newConfig);
  }

  function testInvalidConfigurationZeroRampsReverts() public {
    address[] memory onRamps = new address[](0);
    uint64[] memory minSeqNrByOnRamp = new uint64[](0);
    BlobVerifierInterface.BlobVerifierConfig memory newConfig = BlobVerifierInterface.BlobVerifierConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    vm.expectRevert(BlobVerifierInterface.InvalidConfiguration.selector);

    s_blobVerifier.setConfig(newConfig);
  }
}

/// @notice #resetUnblessedRoots
contract BlobVerifier_resetUnblessedRoots is BlobVerifierSetup {
  // TODO proper AFN blessing handling

  // Reverts
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    bytes32[] memory rootToReset;
    s_blobVerifier.resetUnblessedRoots(rootToReset);
  }
}

/// @notice #report
contract BlobVerifier_report is BlobVerifierSetup {
  event ReportAccepted(CCIP.RelayReport report);

  // Success

  function testSuccess() public {
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
    BlobVerifierInterface.BlobVerifierConfig memory config = blobVerifierConfig();
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: config.onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectEmit(true, false, false, true);

    s_blobVerifier.report(abi.encode(report));
    emit ReportAccepted(report);

    assertEq(max1 + 1, s_blobVerifier.getExpectedNextSequenceNumber(config.onRamps[0]));
    assertEq(max2 + 1, s_blobVerifier.getExpectedNextSequenceNumber(config.onRamps[1]));
    assertEq(max3 + 1, s_blobVerifier.getExpectedNextSequenceNumber(config.onRamps[2]));
  }

  // Reverts

  function testPausedReverts() public {
    s_blobVerifier.pause();
    vm.expectRevert("Pausable: paused");
    bytes memory report;
    s_blobVerifier.report(report);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    bytes memory report;
    s_blobVerifier.report(report);
  }

  function testInvalidRelayReportRootLengthReverts() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](3);
    bytes32[] memory merkleRoots = new bytes32[](2);
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: blobVerifierConfig().onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(BlobVerifierInterface.InvalidRelayReport.selector, report));

    s_blobVerifier.report(abi.encode(report));
  }

  function testInvalidRelayReportIntervalLengthReverts() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](2);
    bytes32[] memory merkleRoots = new bytes32[](3);
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: blobVerifierConfig().onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(BlobVerifierInterface.InvalidRelayReport.selector, report));

    s_blobVerifier.report(abi.encode(report));
  }

  function testUnsupportedOnRampReverts() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    address[] memory onRamps = new address[](1);
    bytes32[] memory merkleRoots = new bytes32[](1);
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(BlobVerifierInterface.UnsupportedOnRamp.selector, onRamps[0]));

    s_blobVerifier.report(abi.encode(report));
  }

  function testInvalidIntervalReverts() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(2, 2);
    address[] memory onRamps = new address[](1);
    onRamps[0] = blobVerifierConfig().onRamps[0];
    bytes32[] memory merkleRoots = new bytes32[](1);
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(BlobVerifierInterface.InvalidInterval.selector, intervals[0], onRamps[0]));

    s_blobVerifier.report(abi.encode(report));
  }

  function testInvalidIntervalMinLargerThanMaxReverts() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(1, 0);
    address[] memory onRamps = new address[](1);
    onRamps[0] = blobVerifierConfig().onRamps[0];
    bytes32[] memory merkleRoots = new bytes32[](1);
    CCIP.RelayReport memory report = CCIP.RelayReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(BlobVerifierInterface.InvalidInterval.selector, intervals[0], onRamps[0]));

    s_blobVerifier.report(abi.encode(report));
  }
}

/// @notice #verify
contract BlobVerifier_verify is BlobVerifierSetup {
  function testSuccess() public {
    CCIP.Interval[] memory intervals = new CCIP.Interval[](1);
    intervals[0] = CCIP.Interval(1, 2);
    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = "rootAndAlsoRootOfRoots";
    address[] memory onRamps = new address[](1);
    onRamps[0] = blobVerifierConfig().onRamps[0];
    s_blobVerifier.report(
      abi.encode(
        CCIP.RelayReport({
          onRamps: onRamps,
          intervals: intervals,
          merkleRoots: merkleRoots,
          rootOfRoots: merkleRoots[0]
        })
      )
    );
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_blobVerifier.verify(merkleRoots, proofs, 2**1, proofs, 2**1);
    assertEq(BLOCK_TIME, timestamp);
  }

  // Reverts

  function testTooManyLeavesReverts() public {
    bytes32[] memory merkleRoots = new bytes32[](258);
    bytes32[] memory proofs = new bytes32[](0);

    vm.expectRevert();

    s_blobVerifier.verify(merkleRoots, proofs, 0, proofs, 0);
  }
}
