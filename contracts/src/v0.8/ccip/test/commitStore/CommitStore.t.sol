// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/MerkleHelper.sol";
import "../helpers/CommitStoreHelper.sol";
import "../BaseTest.t.sol";
import "../../health/AFN.sol";

contract CommitStoreSetup is BaseTest {
  event CommitStoreConfigSet(ICommitStore.CommitStoreConfig config);

  CommitStoreHelper s_commitStore;

  function setUp() public virtual override {
    BaseTest.setUp();

    s_commitStore = new CommitStoreHelper(DEST_CHAIN_ID, SOURCE_CHAIN_ID, s_afn, commitStoreConfig());
  }
}

contract CommitStoreRealAFNSetup is BaseTest {
  CommitStoreHelper s_commitStore;

  function setUp() public virtual override {
    BaseTest.setUp();
    address[] memory participants = new address[](1);
    participants[0] = OWNER;
    uint256[] memory weights = new uint256[](1);
    weights[0] = 1;
    s_afn = new AFN(participants, weights, 1, 1); // Overwrite base mock afn with real.
    s_commitStore = new CommitStoreHelper(SOURCE_CHAIN_ID, DEST_CHAIN_ID, s_afn, commitStoreConfig());
  }
}

/// @notice #constructor
contract CommitStore_constructor is BaseTest {
  function testSuccess() public {
    address[] memory onRamps = new address[](3);
    onRamps[0] = ON_RAMP_ADDRESS;
    onRamps[1] = 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc;
    onRamps[2] = 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC;
    uint64[] memory minSequenceNumbers = new uint64[](3);
    minSequenceNumbers[0] = 1;
    minSequenceNumbers[1] = 2;
    minSequenceNumbers[2] = 4;
    ICommitStore.CommitStoreConfig memory config = ICommitStore.CommitStoreConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSequenceNumbers
    });
    CommitStore commitStore = new CommitStore(DEST_CHAIN_ID, SOURCE_CHAIN_ID, s_afn, config);

    // CommitStore config
    assertEq(minSequenceNumbers[0], commitStore.getExpectedNextSequenceNumber(onRamps[0]));
    assertEq(minSequenceNumbers[1], commitStore.getExpectedNextSequenceNumber(onRamps[1]));
    assertEq(minSequenceNumbers[2], commitStore.getExpectedNextSequenceNumber(onRamps[2]));

    ICommitStore.CommitStoreConfig memory contractConfig = commitStore.getConfig();
    assertEq(keccak256(abi.encode(config.minSeqNrByOnRamp)), keccak256(abi.encode(contractConfig.minSeqNrByOnRamp)));
    assertEq(config.onRamps, contractConfig.onRamps);

    // typeAndVersion
    assertEq("CommitStore 1.0.0", commitStore.typeAndVersion());

    // owner
    assertEq(OWNER, commitStore.owner());

    // HealthChecker
    assertEq(address(s_afn), address(commitStore.getAFN()));
  }

  function testInvalidConfigurationReverts() public {
    address[] memory onRamps = new address[](3);
    uint64[] memory minSequenceNumbers = new uint64[](2);

    vm.expectRevert(ICommitStore.InvalidConfiguration.selector);

    new CommitStore(
      DEST_CHAIN_ID,
      SOURCE_CHAIN_ID,
      s_afn,
      ICommitStore.CommitStoreConfig({onRamps: onRamps, minSeqNrByOnRamp: minSequenceNumbers})
    );
  }
}

/// @notice #setConfig
contract CommitStore_setConfig is CommitStoreSetup {
  // Success

  function testSuccess() public {
    address[] memory onRamps = new address[](1);
    onRamps[0] = address(1);
    uint64[] memory minSeqNrByOnRamp = new uint64[](1);
    minSeqNrByOnRamp[0] = 200;
    ICommitStore.CommitStoreConfig memory newConfig = ICommitStore.CommitStoreConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    // Assert the current value for ON_RAMP_ADDRESS is set
    assertEq(1, s_commitStore.getExpectedNextSequenceNumber(ON_RAMP_ADDRESS));

    vm.expectEmit(false, false, false, false);
    emit CommitStoreConfigSet(newConfig);

    s_commitStore.setConfig(newConfig);

    // Checks whether the new onramp is properly set to the given value
    assertEq(minSeqNrByOnRamp[0], s_commitStore.getExpectedNextSequenceNumber(onRamps[0]));
    // Assert the previously checked value is now 0, indicating successful removal
    // from the supported onRamps list.
    assertEq(0, s_commitStore.getExpectedNextSequenceNumber(ON_RAMP_ADDRESS));
  }

  // Reverts

  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    ICommitStore.CommitStoreConfig memory newConfig;
    s_commitStore.setConfig(newConfig);
  }

  function testInvalidConfigurationLengthMismatchReverts() public {
    address[] memory onRamps = new address[](2);
    uint64[] memory minSeqNrByOnRamp = new uint64[](1);
    ICommitStore.CommitStoreConfig memory newConfig = ICommitStore.CommitStoreConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    vm.expectRevert(ICommitStore.InvalidConfiguration.selector);

    s_commitStore.setConfig(newConfig);
  }

  function testInvalidConfigurationZeroRampsReverts() public {
    address[] memory onRamps = new address[](0);
    uint64[] memory minSeqNrByOnRamp = new uint64[](0);
    ICommitStore.CommitStoreConfig memory newConfig = ICommitStore.CommitStoreConfig({
      onRamps: onRamps,
      minSeqNrByOnRamp: minSeqNrByOnRamp
    });
    vm.expectRevert(ICommitStore.InvalidConfiguration.selector);

    s_commitStore.setConfig(newConfig);
  }
}

/// @notice #resetUnblessedRoots
contract CommitStore_resetUnblessedRoots is CommitStoreSetup {
  // TODO proper AFN blessing handling

  // Reverts
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    bytes32[] memory rootToReset;
    s_commitStore.resetUnblessedRoots(rootToReset);
  }
}

/// @notice #report
contract CommitStore_report is CommitStoreSetup {
  event ReportAccepted(Internal.CommitReport report);

  // Success

  function testSuccess() public {
    uint64 max1 = 931;
    uint64 max2 = 2;
    uint64 max3 = 15;
    Internal.Interval[] memory intervals = new Internal.Interval[](3);
    intervals[0] = Internal.Interval(1, max1);
    intervals[1] = Internal.Interval(1, max2);
    intervals[2] = Internal.Interval(1, max3);
    bytes32[] memory merkleRoots = new bytes32[](3);
    merkleRoots[0] = "test #1";
    merkleRoots[1] = "test #2";
    merkleRoots[2] = "test #3";
    ICommitStore.CommitStoreConfig memory config = commitStoreConfig();
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: config.onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectEmit(true, false, false, true);

    s_commitStore.report(abi.encode(report));
    emit ReportAccepted(report);

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber(config.onRamps[0]));
    assertEq(max2 + 1, s_commitStore.getExpectedNextSequenceNumber(config.onRamps[1]));
    assertEq(max3 + 1, s_commitStore.getExpectedNextSequenceNumber(config.onRamps[2]));
  }

  // Reverts

  function testPausedReverts() public {
    s_commitStore.pause();
    vm.expectRevert("Pausable: paused");
    bytes memory report;
    s_commitStore.report(report);
  }

  function testUnhealthyReverts() public {
    s_afn.voteBad();
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
    bytes memory report;
    s_commitStore.report(report);
  }

  function testInvalidCommitReportRootLengthReverts() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](3);
    bytes32[] memory merkleRoots = new bytes32[](2);
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: commitStoreConfig().onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidCommitReport.selector, report));

    s_commitStore.report(abi.encode(report));
  }

  function testInvalidCommitReportIntervalLengthReverts() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](2);
    bytes32[] memory merkleRoots = new bytes32[](3);
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: commitStoreConfig().onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidCommitReport.selector, report));

    s_commitStore.report(abi.encode(report));
  }

  function testUnsupportedOnRampReverts() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](1);
    address[] memory onRamps = new address[](1);
    bytes32[] memory merkleRoots = new bytes32[](1);
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.UnsupportedOnRamp.selector, onRamps[0]));

    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalReverts() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](1);
    intervals[0] = Internal.Interval(2, 2);
    address[] memory onRamps = new address[](1);
    onRamps[0] = commitStoreConfig().onRamps[0];
    bytes32[] memory merkleRoots = new bytes32[](1);
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidInterval.selector, intervals[0], onRamps[0]));

    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalMinLargerThanMaxReverts() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](1);
    intervals[0] = Internal.Interval(1, 0);
    address[] memory onRamps = new address[](1);
    onRamps[0] = commitStoreConfig().onRamps[0];
    bytes32[] memory merkleRoots = new bytes32[](1);
    Internal.CommitReport memory report = Internal.CommitReport({
      onRamps: onRamps,
      intervals: intervals,
      merkleRoots: merkleRoots,
      rootOfRoots: "root"
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidInterval.selector, intervals[0], onRamps[0]));

    s_commitStore.report(abi.encode(report));
  }
}

/// @notice #verify
contract CommitStore_verify is CommitStoreRealAFNSetup {
  function testNotBlessedSuccess() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](1);
    intervals[0] = Internal.Interval(1, 2);
    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = "rootAndAlsoRootOfRoots";
    address[] memory onRamps = new address[](1);
    onRamps[0] = commitStoreConfig().onRamps[0];
    s_commitStore.report(
      abi.encode(
        Internal.CommitReport({
          onRamps: onRamps,
          intervals: intervals,
          merkleRoots: merkleRoots,
          rootOfRoots: merkleRoots[0]
        })
      )
    );
    bytes32[] memory proofs = new bytes32[](0);
    // We have not blessed this root, should return 0.
    uint256 timestamp = s_commitStore.verify(merkleRoots, proofs, 2**1, proofs, 2**1);
    assertEq(uint256(0), timestamp);
  }

  function testBlessedSuccess() public {
    Internal.Interval[] memory intervals = new Internal.Interval[](1);
    intervals[0] = Internal.Interval(1, 2);
    bytes32[] memory merkleRoots = new bytes32[](1);
    merkleRoots[0] = "rootAndAlsoRootOfRoots";
    address[] memory onRamps = new address[](1);
    onRamps[0] = commitStoreConfig().onRamps[0];
    s_commitStore.report(
      abi.encode(
        Internal.CommitReport({
          onRamps: onRamps,
          intervals: intervals,
          merkleRoots: merkleRoots,
          rootOfRoots: merkleRoots[0]
        })
      )
    );
    // Bless that root.
    bytes32[] memory rootsWithOrigin = new bytes32[](1);
    rootsWithOrigin[0] = keccak256(abi.encode(address(s_commitStore), merkleRoots[0]));
    s_afn.voteToBlessRoots(rootsWithOrigin);
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_commitStore.verify(merkleRoots, proofs, 2**1, proofs, 2**1);
    assertEq(BLOCK_TIME, timestamp);
  }

  // Reverts

  function testTooManyLeavesReverts() public {
    bytes32[] memory merkleRoots = new bytes32[](258);
    bytes32[] memory proofs = new bytes32[](0);

    vm.expectRevert(ICommitStore.InvalidProof.selector);

    s_commitStore.verify(merkleRoots, proofs, 0, proofs, 0);
  }
}
