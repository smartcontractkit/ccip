// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/MerkleHelper.sol";
import "../helpers/CommitStoreHelper.sol";
import "../../AFN.sol";
import "../../PriceRegistry.sol";
import "../priceRegistry/PriceRegistry.t.sol";

contract CommitStoreSetup is PriceRegistrySetup {
  CommitStoreHelper s_commitStore;

  function setUp() public virtual override {
    PriceRegistrySetup.setUp();

    s_commitStore = new CommitStoreHelper(
      ICommitStore.StaticConfig({chainId: DEST_CHAIN_ID, sourceChainId: SOURCE_CHAIN_ID, onRamp: ON_RAMP_ADDRESS}),
      ICommitStore.DynamicConfig({priceRegistry: address(s_priceRegistry), afn: address(s_afn)})
    );

    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = address(s_commitStore);
    s_priceRegistry.applyPriceUpdatersUpdates(priceUpdaters, new address[](0));
  }
}

contract CommitStoreRealAFNSetup is PriceRegistrySetup {
  CommitStoreHelper s_commitStore;

  function setUp() public virtual override {
    PriceRegistrySetup.setUp();
    address[] memory participants = new address[](1);
    participants[0] = OWNER;
    uint256[] memory weights = new uint256[](1);
    weights[0] = 1;
    s_afn = new AFN(participants, weights, 1, 1); // Overwrite base mock afn with real.
    s_commitStore = new CommitStoreHelper(
      ICommitStore.StaticConfig({chainId: DEST_CHAIN_ID, sourceChainId: SOURCE_CHAIN_ID, onRamp: ON_RAMP_ADDRESS}),
      ICommitStore.DynamicConfig({priceRegistry: address(s_priceRegistry), afn: address(s_afn)})
    );
  }
}

/// @notice #constructor
contract CommitStore_constructor is PriceRegistrySetup {
  event ConfigSet(ICommitStore.StaticConfig, ICommitStore.DynamicConfig);

  function testConstructorSuccess() public {
    ICommitStore.StaticConfig memory staticConfig = ICommitStore.StaticConfig({
      chainId: DEST_CHAIN_ID,
      sourceChainId: SOURCE_CHAIN_ID,
      onRamp: 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc
    });
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({
      priceRegistry: address(s_priceRegistry),
      afn: address(s_afn)
    });

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    CommitStore commitStore = new CommitStore(staticConfig, dynamicConfig);

    ICommitStore.StaticConfig memory gotStaticConfig = commitStore.getStaticConfig();

    assertEq(staticConfig.chainId, gotStaticConfig.chainId);
    assertEq(staticConfig.sourceChainId, gotStaticConfig.sourceChainId);
    assertEq(staticConfig.onRamp, gotStaticConfig.onRamp);

    ICommitStore.DynamicConfig memory gotDynamicConfig = commitStore.getDynamicConfig();

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
  event ConfigSet(ICommitStore.StaticConfig, ICommitStore.DynamicConfig);

  function testSetMinSeqNrSuccess() public {
    ICommitStore.StaticConfig memory staticConfig = s_commitStore.getStaticConfig();
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({
      priceRegistry: address(23784264),
      afn: address(s_afn)
    });

    vm.expectEmit();
    emit ConfigSet(staticConfig, dynamicConfig);

    s_commitStore.setDynamicConfig(dynamicConfig);

    ICommitStore.DynamicConfig memory gotDynamicConfig = s_commitStore.getDynamicConfig();
    assertEq(gotDynamicConfig.priceRegistry, dynamicConfig.priceRegistry);
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({
      priceRegistry: address(23784264),
      afn: address(s_afn)
    });

    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_commitStore.setDynamicConfig(dynamicConfig);
  }
}

/// @notice #resetUnblessedRoots
contract CommitStore_resetUnblessedRoots is CommitStoreSetup {
  event RootRemoved(bytes32 root);

  function setUp() public virtual override {
    CommitStoreSetup.setUp();

    // Setup a real AFN instead of the mock one that always
    // returns "true" for any root.
    address[] memory participants = new address[](1);
    participants[0] = OWNER;
    uint256[] memory weights = new uint256[](1);
    weights[0] = 2000;

    s_afn = new AFN(participants, weights, weights[0], weights[0]);

    s_commitStore.setDynamicConfig(
      ICommitStore.DynamicConfig({priceRegistry: address(s_priceRegistry), afn: address(s_afn)})
    );
  }

  function testResetUnblessedRootsSuccess() public {
    bytes32[] memory rootsToReset = new bytes32[](3);
    rootsToReset[0] = "1";
    rootsToReset[1] = "2";
    rootsToReset[2] = "3";

    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(1, 2),
      merkleRoot: rootsToReset[0]
    });

    s_commitStore.report(abi.encode(report));

    report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(3, 4),
      merkleRoot: rootsToReset[1]
    });

    s_commitStore.report(abi.encode(report));

    report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(5, 5),
      merkleRoot: rootsToReset[2]
    });

    s_commitStore.report(abi.encode(report));

    bytes32[] memory blessedRoots = new bytes32[](1);
    blessedRoots[0] = keccak256(abi.encode(address(s_commitStore), rootsToReset[1]));

    s_afn.voteToBlessRoots(blessedRoots);

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
  event ReportAccepted(ICommitStore.CommitReport report);
  event UsdPerTokenUpdated(address indexed feeToken, uint256 value, uint256 timestamp);

  function testReportOnlyRootSuccess_gas() public {
    vm.pauseGasMetering();
    uint64 max1 = 931;
    bytes32 root = "Only a single root";
    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(1, max1),
      merkleRoot: root
    });

    vm.expectEmit();
    emit ReportAccepted(report);

    vm.resumeGasMetering();
    s_commitStore.report(abi.encode(report));
    vm.pauseGasMetering();

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber());
    assertEq(block.timestamp, s_commitStore.getMerkleRoot(root));
    vm.resumeGasMetering();
  }

  function testReportAndPriceUpdateSuccess() public {
    uint64 max1 = 12;
    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](1);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceFeeToken, usdPerToken: 4e18});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: tokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });

    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: priceUpdates,
      interval: ICommitStore.Interval(1, max1),
      merkleRoot: "test #2"
    });

    vm.expectEmit();
    emit ReportAccepted(report);

    s_commitStore.report(abi.encode(report));

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber());
  }

  function testOnlyPriceUpdatesSuccess() public {
    Internal.TokenPriceUpdate[] memory tokenPriceUpdates = new Internal.TokenPriceUpdate[](1);
    tokenPriceUpdates[0] = Internal.TokenPriceUpdate({sourceToken: s_sourceFeeToken, usdPerToken: 4e18});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      tokenPriceUpdates: tokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });

    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: priceUpdates,
      interval: ICommitStore.Interval(0, 0),
      merkleRoot: ""
    });

    vm.expectEmit();
    emit UsdPerTokenUpdated(s_sourceFeeToken, 4e18, block.timestamp);

    s_commitStore.report(abi.encode(report));
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
    vm.expectRevert(ICommitStore.BadAFNSignal.selector);
    bytes memory report;
    s_commitStore.report(report);
  }

  function testInvalidRootRevert() public {
    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(1, 4),
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(ICommitStore.InvalidRoot.selector);
    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalReverts() public {
    ICommitStore.Interval memory interval = ICommitStore.Interval(2, 2);
    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: interval,
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidInterval.selector, interval));

    s_commitStore.report(abi.encode(report));
  }

  function testInvalidIntervalMinLargerThanMaxReverts() public {
    ICommitStore.Interval memory interval = ICommitStore.Interval(1, 0);
    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: interval,
      merkleRoot: bytes32(0)
    });

    vm.expectRevert(abi.encodeWithSelector(ICommitStore.InvalidInterval.selector, interval));

    s_commitStore.report(abi.encode(report));
  }
}

/// @notice #verify
contract CommitStore_verify is CommitStoreRealAFNSetup {
  function testNotBlessedSuccess() public {
    bytes32[] memory leaves = new bytes32[](1);
    leaves[0] = "rootAndAlsoRootOfRoots";
    s_commitStore.report(
      abi.encode(
        ICommitStore.CommitReport({
          priceUpdates: getEmptyPriceUpdates(),
          interval: ICommitStore.Interval(1, 2),
          merkleRoot: leaves[0]
        })
      )
    );
    bytes32[] memory proofs = new bytes32[](0);
    // We have not blessed this root, should return 0.
    uint256 timestamp = s_commitStore.verify(leaves, proofs, 2**1);
    assertEq(uint256(0), timestamp);
  }

  function testBlessedSuccess() public {
    bytes32[] memory leaves = new bytes32[](1);
    leaves[0] = "rootAndAlsoRootOfRoots";
    s_commitStore.report(
      abi.encode(
        ICommitStore.CommitReport({
          priceUpdates: getEmptyPriceUpdates(),
          interval: ICommitStore.Interval(1, 2),
          merkleRoot: leaves[0]
        })
      )
    );
    // Bless that root.
    bytes32[] memory rootsWithOrigin = new bytes32[](1);
    rootsWithOrigin[0] = keccak256(abi.encode(address(s_commitStore), leaves[0]));
    s_afn.voteToBlessRoots(rootsWithOrigin);
    bytes32[] memory proofs = new bytes32[](0);
    uint256 timestamp = s_commitStore.verify(leaves, proofs, 2**1);
    assertEq(BLOCK_TIME, timestamp);
  }

  // Reverts

  function testTooManyLeavesReverts() public {
    bytes32[] memory leaves = new bytes32[](258);
    bytes32[] memory proofs = new bytes32[](0);

    vm.expectRevert(ICommitStore.InvalidProof.selector);

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
    s_afn.voteBad();
    assertEq(s_commitStore.isAFNHealthy(), false);
    s_afn.recoverFromBadSignal();
    assertEq(s_commitStore.isAFNHealthy(), true);
  }
}
