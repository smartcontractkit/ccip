// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/CommitStoreHelper.sol";
import {AFN} from "../../AFN.sol";
import {IAFN} from "../../interfaces/IAFN.sol";
import "../../PriceRegistry.sol";
import "../priceRegistry/PriceRegistry.t.sol";
import "../ocr/OCR2Base.t.sol";
import "../../ocr/OCR2Abstract.sol";

contract CommitStoreSetup is PriceRegistrySetup, OCR2BaseSetup {
  event ConfigSet(CommitStore.StaticConfig, CommitStore.DynamicConfig);

  CommitStoreHelper s_commitStore;

  function setUp() public virtual override(PriceRegistrySetup, OCR2BaseSetup) {
    PriceRegistrySetup.setUp();
    OCR2BaseSetup.setUp();

    s_commitStore = new CommitStoreHelper(
      CommitStore.StaticConfig({
        chainSelector: DEST_CHAIN_ID,
        sourceChainSelector: SOURCE_CHAIN_ID,
        onRamp: ON_RAMP_ADDRESS
      })
    );
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: address(s_priceRegistry),
      afn: address(s_mockAFN)
    });
    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );

    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = address(s_commitStore);
    s_priceRegistry.applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
  }
}

contract CommitStoreRealAFNSetup is PriceRegistrySetup, OCR2BaseSetup {
  CommitStoreHelper s_commitStore;

  AFN internal s_afn;

  function setUp() public virtual override(PriceRegistrySetup, OCR2BaseSetup) {
    PriceRegistrySetup.setUp();
    OCR2BaseSetup.setUp();

    AFN.Voter[] memory voters = new AFN.Voter[](1);
    voters[0] = AFN.Voter({
      blessVoteAddr: OWNER,
      curseVoteAddr: address(9999),
      curseUnvoteAddr: address(19999),
      blessWeight: 1,
      curseWeight: 1
    });
    // Overwrite base mock afn with real.
    s_afn = new AFN(AFN.Config({voters: voters, blessWeightThreshold: 1, curseWeightThreshold: 1}));
    s_commitStore = new CommitStoreHelper(
      CommitStore.StaticConfig({
        chainSelector: DEST_CHAIN_ID,
        sourceChainSelector: SOURCE_CHAIN_ID,
        onRamp: ON_RAMP_ADDRESS
      })
    );
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: address(s_priceRegistry),
      afn: address(s_afn)
    });
    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );
  }
}

/// @notice #constructor
contract CommitStore_constructor is PriceRegistrySetup, OCR2BaseSetup {
  event ConfigSet(CommitStore.StaticConfig, CommitStore.DynamicConfig);

  function setUp() public virtual override(PriceRegistrySetup, OCR2BaseSetup) {
    PriceRegistrySetup.setUp();
    OCR2BaseSetup.setUp();
  }

  function testConstructorSuccess() public {
    CommitStore.StaticConfig memory staticConfig = CommitStore.StaticConfig({
      chainSelector: DEST_CHAIN_ID,
      sourceChainSelector: SOURCE_CHAIN_ID,
      onRamp: 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc
    });
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: address(s_priceRegistry),
      afn: address(s_mockAFN)
    });

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    CommitStore commitStore = new CommitStore(staticConfig);
    commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );

    CommitStore.StaticConfig memory gotStaticConfig = commitStore.getStaticConfig();

    assertEq(staticConfig.chainSelector, gotStaticConfig.chainSelector);
    assertEq(staticConfig.sourceChainSelector, gotStaticConfig.sourceChainSelector);
    assertEq(staticConfig.onRamp, gotStaticConfig.onRamp);

    CommitStore.DynamicConfig memory gotDynamicConfig = commitStore.getDynamicConfig();

    assertEq(dynamicConfig.priceRegistry, gotDynamicConfig.priceRegistry);
    assertEq(dynamicConfig.afn, gotDynamicConfig.afn);

    // CommitStore initial values
    assertEq(1, commitStore.getExpectedNextSequenceNumber());
    assertEq(commitStore.typeAndVersion(), "CommitStore 1.0.0");
    assertEq(OWNER, commitStore.owner());
  }
}

/// @notice #setMinSeqNr
contract CommitStore_setMinSeqNr is CommitStoreSetup {
  function testSetMinSeqNrSuccess(uint64 minSeqNr) public {
    s_commitStore.setMinSeqNr(minSeqNr);

    assertEq(s_commitStore.getExpectedNextSequenceNumber(), minSeqNr);
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_commitStore.setMinSeqNr(6723);
  }
}

/// @notice #setDynamicConfig
contract CommitStore_setDynamicConfig is CommitStoreSetup {
  function testSetMinSeqNrSuccess(address priceRegistry, address afn) public {
    vm.assume(priceRegistry != address(0) && afn != address(0));
    CommitStore.StaticConfig memory staticConfig = s_commitStore.getStaticConfig();
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: priceRegistry,
      afn: afn
    });
    bytes memory onchainConfig = abi.encode(dynamicConfig);

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    uint32 configCount = 1;

    vm.expectEmit();
    emit ConfigSet(
      uint32(block.number),
      getBasicConfigDigest(address(s_commitStore), s_f, configCount, onchainConfig),
      configCount + 1,
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      onchainConfig,
      s_offchainConfigVersion,
      abi.encode("")
    );

    CommitStore.DynamicConfig memory gotDynamicConfig = s_commitStore.getDynamicConfig();
    assertEq(gotDynamicConfig.priceRegistry, dynamicConfig.priceRegistry);
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: address(23784264),
      afn: address(s_mockAFN)
    });

    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );
  }

  function testInvalidCommitStoreConfigReverts() public {
    CommitStore.DynamicConfig memory dynamicConfig = CommitStore.DynamicConfig({
      priceRegistry: address(0),
      afn: address(1)
    });

    vm.expectRevert(CommitStore.InvalidCommitStoreConfig.selector);
    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );

    dynamicConfig.priceRegistry = address(1);
    dynamicConfig.afn = address(0);

    vm.expectRevert(CommitStore.InvalidCommitStoreConfig.selector);
    s_commitStore.setOCR2Config(
      s_valid_signers,
      s_valid_transmitters,
      s_f,
      abi.encode(dynamicConfig),
      s_offchainConfigVersion,
      abi.encode("")
    );
  }
}

/// @notice #resetUnblessedRoots
contract CommitStore_resetUnblessedRoots is CommitStoreRealAFNSetup {
  event RootRemoved(bytes32 root);

  function testResetUnblessedRootsSuccess() public {
    bytes32[] memory rootsToReset = new bytes32[](3);
    rootsToReset[0] = "1";
    rootsToReset[1] = "2";
    rootsToReset[2] = "3";

    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: CommitStore.Interval(1, 2),
      merkleRoot: rootsToReset[0]
    });

    s_commitStore.report(abi.encode(report));

    report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: CommitStore.Interval(3, 4),
      merkleRoot: rootsToReset[1]
    });

    s_commitStore.report(abi.encode(report));

    report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: CommitStore.Interval(5, 5),
      merkleRoot: rootsToReset[2]
    });

    s_commitStore.report(abi.encode(report));

    IAFN.TaggedRoot[] memory blessedTaggedRoots = new IAFN.TaggedRoot[](1);
    blessedTaggedRoots[0] = IAFN.TaggedRoot({commitStore: address(s_commitStore), root: rootsToReset[1]});

    s_afn.voteToBless(blessedTaggedRoots);

    vm.expectEmit(false, false, false, true);
    emit RootRemoved(rootsToReset[0]);

    vm.expectEmit(false, false, false, true);
    emit RootRemoved(rootsToReset[2]);

    s_commitStore.resetUnblessedRoots(rootsToReset);

    assertEq(0, s_commitStore.getMerkleRoot(rootsToReset[0]));
    assertEq(BLOCK_TIME, s_commitStore.getMerkleRoot(rootsToReset[1]));
    assertEq(0, s_commitStore.getMerkleRoot(rootsToReset[2]));
  }

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
  event ReportAccepted(CommitStore.CommitReport report);
  event UsdPerTokenUpdated(address indexed feeToken, uint256 value, uint256 timestamp);

  function testReportOnlyRootSuccess_gas() public {
    vm.pauseGasMetering();
    uint64 max1 = 931;
    bytes32 root = "Only a single root";
    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: CommitStore.Interval(1, max1),
      merkleRoot: root
    });

    vm.expectEmit();
    emit ReportAccepted(report);

    bytes memory encodedReport = abi.encode(report);

    vm.resumeGasMetering();
    s_commitStore.report(encodedReport);
    vm.pauseGasMetering();

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber());
    assertEq(block.timestamp, s_commitStore.getMerkleRoot(root));
    vm.resumeGasMetering();
  }

  function testReportAndPriceUpdateSuccess() public {
    uint64 max1 = 12;

    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getSinglePriceUpdateStruct(s_sourceFeeToken, 4e18),
      interval: CommitStore.Interval(1, max1),
      merkleRoot: "test #2"
    });

    vm.expectEmit();
    emit ReportAccepted(report);

    s_commitStore.report(abi.encode(report));

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber());
  }

  function testOnlyPriceUpdatesSuccess() public {
    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getSinglePriceUpdateStruct(s_sourceFeeToken, 4e18),
      interval: CommitStore.Interval(0, 0),
      merkleRoot: ""
    });

    vm.expectEmit();
    emit UsdPerTokenUpdated(s_sourceFeeToken, 4e18, block.timestamp);

    s_commitStore.report(abi.encode(report));
  }

  // Reverts

  function testPausedReverts() public {
    s_commitStore.pause();
    bytes memory report;
    vm.expectRevert("Pausable: paused");
    s_commitStore.report(report);
  }

  function testUnhealthyReverts() public {
    s_mockAFN.voteToCurse(0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
    vm.expectRevert(CommitStore.BadAFNSignal.selector);
    bytes memory report;
    s_commitStore.report(report);
  }

  function testInvalidRootRevert() public {
    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: CommitStore.Interval(1, 4),
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(CommitStore.InvalidRoot.selector);
    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalReverts() public {
    CommitStore.Interval memory interval = CommitStore.Interval(2, 2);
    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: interval,
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(abi.encodeWithSelector(CommitStore.InvalidInterval.selector, interval));

    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalMinLargerThanMaxReverts() public {
    CommitStore.Interval memory interval = CommitStore.Interval(1, 0);
    CommitStore.CommitReport memory report = CommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: interval,
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(abi.encodeWithSelector(CommitStore.InvalidInterval.selector, interval));

    s_commitStore.report(abi.encode(report));
  }
}

/// @notice #verify
contract CommitStore_verify is CommitStoreRealAFNSetup {
  function testNotBlessedSuccess() public {
    bytes32[] memory leaves = new bytes32[](1);
    leaves[0] = "root";
    s_commitStore.report(
      abi.encode(
        CommitStore.CommitReport({
          priceUpdates: getEmptyPriceUpdates(),
          interval: CommitStore.Interval(1, 2),
          merkleRoot: leaves[0]
        })
      )
    );
    bytes32[] memory proofs = new bytes32[](0);
    // We have not blessed this root, should return 0.
    uint256 timestamp = s_commitStore.verify(leaves, proofs, 0);
    assertEq(uint256(0), timestamp);
  }

  function testBlessedSuccess() public {
    bytes32[] memory leaves = new bytes32[](1);
    leaves[0] = "root";
    s_commitStore.report(
      abi.encode(
        CommitStore.CommitReport({
          priceUpdates: getEmptyPriceUpdates(),
          interval: CommitStore.Interval(1, 2),
          merkleRoot: leaves[0]
        })
      )
    );
    // Bless that root.
    IAFN.TaggedRoot[] memory taggedRoots = new IAFN.TaggedRoot[](1);
    taggedRoots[0] = IAFN.TaggedRoot({commitStore: address(s_commitStore), root: leaves[0]});
    s_afn.voteToBless(taggedRoots);
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_commitStore.verify(leaves, proofs, 0);
    assertEq(BLOCK_TIME, timestamp);
  }

  // Reverts

  function testPausedReverts() public {
    s_commitStore.pause();

    bytes32[] memory hashedLeaves = new bytes32[](0);
    bytes32[] memory proofs = new bytes32[](0);
    uint256 proofFlagBits = 0;

    vm.expectRevert("Pausable: paused");
    s_commitStore.verify(hashedLeaves, proofs, proofFlagBits);
  }

  function testTooManyLeavesReverts() public {
    bytes32[] memory leaves = new bytes32[](258);
    bytes32[] memory proofs = new bytes32[](0);

    vm.expectRevert(MerkleMultiProof.InvalidProof.selector);

    s_commitStore.verify(leaves, proofs, 0);
  }
}

contract CommitStore_afn is CommitStoreSetup {
  function testAFN() public {
    // Test pausing
    assertEq(s_commitStore.paused(), false);
    s_commitStore.pause();
    assertEq(s_commitStore.paused(), true);
    s_commitStore.unpause();
    assertEq(s_commitStore.paused(), false);

    // Test afn
    assertEq(s_commitStore.isAFNHealthy(), true);
    s_mockAFN.voteToCurse(0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
    assertEq(s_commitStore.isAFNHealthy(), false);
    AFN.UnvoteToCurseRecord[] memory records = new AFN.UnvoteToCurseRecord[](1);
    records[0] = AFN.UnvoteToCurseRecord({curseVoteAddr: OWNER, cursesHash: bytes32(uint256(0)), forceUnvote: true});
    s_mockAFN.ownerUnvoteToCurse(records);
    assertEq(s_commitStore.isAFNHealthy(), true);
  }
}
