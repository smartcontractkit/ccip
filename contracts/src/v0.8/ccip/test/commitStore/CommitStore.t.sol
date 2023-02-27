// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/MerkleHelper.sol";
import "../helpers/CommitStoreHelper.sol";
import "../../health/AFN.sol";
import "../../prices/PriceRegistry.sol";
import "../prices/PriceRegistry.t.sol";

contract CommitStoreSetup is PriceRegistrySetup {
  CommitStoreHelper s_commitStore;

  function setUp() public virtual override {
    PriceRegistrySetup.setUp();

    s_commitStore = new CommitStoreHelper(
      ICommitStore.StaticConfig({chainId: DEST_CHAIN_ID, sourceChainId: SOURCE_CHAIN_ID, onRamp: ON_RAMP_ADDRESS}),
      ICommitStore.DynamicConfig({priceRegistry: address(s_priceRegistry)}),
      s_afn
    );

    address[] memory priceUpdaters = new address[](1);
    priceUpdaters[0] = address(s_commitStore);
    s_priceRegistry.addPriceUpdaters(priceUpdaters);
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
      ICommitStore.DynamicConfig({priceRegistry: address(s_priceRegistry)}),
      s_afn
    );
  }
}

/// @notice #constructor
contract CommitStore_constructor is PriceRegistrySetup {
  event StaticConfigSet(ICommitStore.StaticConfig);
  event DynamicConfigSet(ICommitStore.DynamicConfig);

  function testConstructorSuccess() public {
    ICommitStore.StaticConfig memory staticConfig = ICommitStore.StaticConfig({
      chainId: DEST_CHAIN_ID,
      sourceChainId: SOURCE_CHAIN_ID,
      onRamp: 0x2C44CDDdB6a900Fa2B585dd299E03D12Fa4293Bc
    });
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({
      priceRegistry: address(s_priceRegistry)
    });

    vm.expectEmit(false, false, false, true);
    emit StaticConfigSet(staticConfig);

    vm.expectEmit(false, false, false, true);
    emit DynamicConfigSet(dynamicConfig);

    CommitStore commitStore = new CommitStore(staticConfig, dynamicConfig, s_afn);

    ICommitStore.StaticConfig memory gotStaticConfig = commitStore.getStaticConfig();

    assertEq(staticConfig.chainId, gotStaticConfig.chainId);
    assertEq(staticConfig.sourceChainId, gotStaticConfig.sourceChainId);
    assertEq(staticConfig.onRamp, gotStaticConfig.onRamp);

    ICommitStore.DynamicConfig memory gotDynamicConfig = commitStore.getDynamicConfig();

    assertEq(dynamicConfig.priceRegistry, gotDynamicConfig.priceRegistry);

    // AFN
    assertEq(address(s_afn), address(commitStore.getAFN()));

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
  event DynamicConfigSet(ICommitStore.DynamicConfig);

  function testSetMinSeqNrSuccess() public {
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({priceRegistry: address(23784264)});

    vm.expectEmit(false, false, false, true);
    emit DynamicConfigSet(dynamicConfig);

    s_commitStore.setDynamicConfig(dynamicConfig);

    ICommitStore.DynamicConfig memory gotDynamicConfig = s_commitStore.getDynamicConfig();
    assertEq(gotDynamicConfig.priceRegistry, dynamicConfig.priceRegistry);
  }

  // Reverts
  function testOnlyOwnerReverts() public {
    ICommitStore.DynamicConfig memory dynamicConfig = ICommitStore.DynamicConfig({priceRegistry: address(23784264)});

    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_commitStore.setDynamicConfig(dynamicConfig);
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
  event ReportAccepted(ICommitStore.CommitReport report);
  event UsdPerFeeTokenUpdated(address indexed feeToken, uint256 value, uint256 timestamp);

  // Success

  function testReportOnlyRootSuccess_gas() public {
    vm.pauseGasMetering();
    uint64 max1 = 931;
    bytes32 root = "Only a single root";
    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: getEmptyPriceUpdates(),
      interval: ICommitStore.Interval(1, max1),
      merkleRoot: root
    });

    vm.expectEmit(false, false, false, true);
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
    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](1);
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceFeeToken, usdPerFeeToken: 4e18});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });

    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: priceUpdates,
      interval: ICommitStore.Interval(1, max1),
      merkleRoot: "test #2"
    });

    vm.expectEmit(false, false, false, true);
    emit ReportAccepted(report);

    s_commitStore.report(abi.encode(report));

    assertEq(max1 + 1, s_commitStore.getExpectedNextSequenceNumber());
  }

  function testOnlyPriceUpdatesSuccess() public {
    Internal.FeeTokenPriceUpdate[] memory feeTokenPriceUpdates = new Internal.FeeTokenPriceUpdate[](1);
    feeTokenPriceUpdates[0] = Internal.FeeTokenPriceUpdate({sourceFeeToken: s_sourceFeeToken, usdPerFeeToken: 4e18});
    Internal.PriceUpdates memory priceUpdates = Internal.PriceUpdates({
      feeTokenPriceUpdates: feeTokenPriceUpdates,
      destChainId: 0,
      usdPerUnitGas: 0
    });

    ICommitStore.CommitReport memory report = ICommitStore.CommitReport({
      priceUpdates: priceUpdates,
      interval: ICommitStore.Interval(0, 0),
      merkleRoot: ""
    });

    vm.expectEmit(true, true, true, true);
    emit UsdPerFeeTokenUpdated(s_sourceFeeToken, 4e18, block.timestamp);

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
    vm.expectRevert(HealthChecker.BadAFNSignal.selector);
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
