// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";

import {CCIPCapabilityConfiguration} from "../../capability/CCIPCapabilityConfiguration.sol";
import {ICapabilityRegistry} from "../../capability/interfaces/ICapabilityRegistry.sol";
import {CCIPCapabilityConfigurationHelper} from "../helpers/CCIPCapabilityConfigurationHelper.sol";

contract CCIPCapabilityConfigurationSetup is Test {
  address public constant OWNER = 0x82ae2B4F57CA5C1CBF8f744ADbD3697aD1a35AFe;
  address public constant CAPABILITY_REGISTRY = 0x272aF4BF7FBFc4944Ed59F914Cd864DfD912D55e;

  CCIPCapabilityConfigurationHelper public s_ccipCC;

  function setUp() public {
    changePrank(OWNER);
    s_ccipCC = new CCIPCapabilityConfigurationHelper(CAPABILITY_REGISTRY);
  }

  function makeAssociativeArray(uint256 length, uint256 seed) internal pure returns (bytes[][] memory) {
    bytes[][] memory arr = new bytes[][](length);
    for (uint256 i = 0; i < length; i++) {
      arr[i] = new bytes[](2);
      arr[i][0] = abi.encode(keccak256(abi.encode(i, 1, seed)));
      arr[i][1] = abi.encode(address(uint160(i)));
    }
    return arr;
  }
}

contract CCIPCapabilityConfiguration_chainConfig is CCIPCapabilityConfigurationSetup {
  event ChainConfigSet(uint64 chainSelector, CCIPCapabilityConfiguration.ChainConfig chainConfig);
  event ChainConfigRemoved(uint64 chainSelector);

  function test_applyChainConfigUpdates_addChainConfigs_Success() public {
    bytes32[] memory chainReaders = new bytes32[](1);
    chainReaders[0] = keccak256(abi.encode(1));
    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory adds = new CCIPCapabilityConfiguration.ChainConfigUpdate[](2);
    adds[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: chainReaders, config: bytes("config1")})
    });
    adds[1] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 2,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: chainReaders, config: bytes("config2")})
    });

    vm.mockCall(
      CAPABILITY_REGISTRY,
      abi.encodeWithSelector(ICapabilityRegistry.getNode.selector, chainReaders[0]),
      abi.encode(
        ICapabilityRegistry.NodeParams({
          nodeOperatorId: 1,
          signer: address(1),
          p2pId: chainReaders[0],
          hashedCapabilityIds: new bytes32[](0)
        }),
        uint32(1)
      )
    );

    vm.expectEmit();
    emit ChainConfigSet(1, adds[0].chainConfig);
    vm.expectEmit();
    emit ChainConfigSet(2, adds[1].chainConfig);
    s_ccipCC.applyChainConfigUpdates(new CCIPCapabilityConfiguration.ChainConfigUpdate[](0), adds);
  }

  function test_applyChainConfigUpdates_removeChainConfigs_Success() public {
    bytes32[] memory chainReaders = new bytes32[](1);
    chainReaders[0] = keccak256(abi.encode(1));
    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory adds = new CCIPCapabilityConfiguration.ChainConfigUpdate[](2);
    adds[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: chainReaders, config: bytes("config1")})
    });
    adds[1] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 2,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: chainReaders, config: bytes("config2")})
    });

    vm.mockCall(
      CAPABILITY_REGISTRY,
      abi.encodeWithSelector(ICapabilityRegistry.getNode.selector, chainReaders[0]),
      abi.encode(
        ICapabilityRegistry.NodeParams({
          nodeOperatorId: 1,
          signer: address(1),
          p2pId: chainReaders[0],
          hashedCapabilityIds: new bytes32[](0)
        }),
        uint32(1)
      )
    );

    vm.expectEmit();
    emit ChainConfigSet(1, adds[0].chainConfig);
    vm.expectEmit();
    emit ChainConfigSet(2, adds[1].chainConfig);
    s_ccipCC.applyChainConfigUpdates(new CCIPCapabilityConfiguration.ChainConfigUpdate[](0), adds);

    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory removes =
      new CCIPCapabilityConfiguration.ChainConfigUpdate[](1);
    removes[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: new bytes32[](0), config: bytes("config1")})
    });

    vm.expectEmit();
    emit ChainConfigRemoved(1);
    s_ccipCC.applyChainConfigUpdates(removes, new CCIPCapabilityConfiguration.ChainConfigUpdate[](0));
  }

  function test_applyChainConfigUpdates_selectorNotFound_Reverts() public {
    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory removes =
      new CCIPCapabilityConfiguration.ChainConfigUpdate[](1);
    removes[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: new bytes32[](0), config: abi.encode(1, 2, 3)})
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.ChainSelectorNotFound.selector, 1));
    s_ccipCC.applyChainConfigUpdates(removes, new CCIPCapabilityConfiguration.ChainConfigUpdate[](0));
  }

  function test_applyChainConfigUpdates_nodeNotInRegistry_Reverts() public {
    bytes32[] memory chainReaders = new bytes32[](1);
    chainReaders[0] = keccak256(abi.encode(1));
    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory adds = new CCIPCapabilityConfiguration.ChainConfigUpdate[](1);
    adds[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: chainReaders, config: abi.encode(1, 2, 3)})
    });

    vm.mockCall(
      CAPABILITY_REGISTRY,
      abi.encodeWithSelector(ICapabilityRegistry.getNode.selector, chainReaders[0]),
      abi.encode(
        ICapabilityRegistry.NodeParams({
          nodeOperatorId: 0,
          signer: address(0),
          p2pId: bytes32(uint256(0)),
          hashedCapabilityIds: new bytes32[](0)
        }),
        uint32(1)
      )
    );

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.NodeNotInRegistry.selector, chainReaders[0]));
    s_ccipCC.applyChainConfigUpdates(new CCIPCapabilityConfiguration.ChainConfigUpdate[](0), adds);
  }
}

contract CCIPCapabilityConfiguration_validateConfig is CCIPCapabilityConfigurationSetup {
  event ChainConfigSet(uint64 chainSelector, CCIPCapabilityConfiguration.ChainConfig chainConfig);

  function addChainConfig(uint256 numNodes) internal returns (bytes[][] memory signers, bytes[][] memory transmitters) {
    signers = makeAssociativeArray(numNodes, 10);
    transmitters = makeAssociativeArray(numNodes, 20);
    bytes32[] memory readers = new bytes32[](numNodes);
    for (uint256 i = 0; i < numNodes; i++) {
      readers[i] = abi.decode(signers[i][0], (bytes32));
      vm.mockCall(
        CAPABILITY_REGISTRY,
        abi.encodeWithSelector(ICapabilityRegistry.getNode.selector, readers[i]),
        abi.encode(
          ICapabilityRegistry.NodeParams({
            nodeOperatorId: 1,
            signer: address(uint160(i)),
            p2pId: readers[i],
            hashedCapabilityIds: new bytes32[](0)
          }),
          uint32(1)
        )
      );
    }
    // Add chain selector for chain 1.
    CCIPCapabilityConfiguration.ChainConfigUpdate[] memory adds = new CCIPCapabilityConfiguration.ChainConfigUpdate[](1);
    adds[0] = CCIPCapabilityConfiguration.ChainConfigUpdate({
      chainSelector: 1,
      chainConfig: CCIPCapabilityConfiguration.ChainConfig({readers: readers, config: bytes("config1")})
    });

    vm.expectEmit();
    emit ChainConfigSet(1, adds[0].chainConfig);
    s_ccipCC.applyChainConfigUpdates(new CCIPCapabilityConfiguration.ChainConfigUpdate[](0), adds);

    return (signers, transmitters);
  }

  function test__validateConfig_Success() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_ChainSelectorNotSet_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 0, // invalid
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(CCIPCapabilityConfiguration.ChainSelectorNotSet.selector);
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_ChainSelectorNotFound_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 2, // not set
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.ChainSelectorNotFound.selector, 2));
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_TooManySigners_Reverts() public {
    // 32 > 31 (max num oracles)
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(32);

    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(CCIPCapabilityConfiguration.TooManySigners.selector);
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_TooManyTransmitters_Reverts() public {
    // 32 > 31 (max num oracles)
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(32);

    // truncate signers but keep transmitters > 31
    assembly {
      mstore(signers, 30)
    }

    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(CCIPCapabilityConfiguration.TooManyTransmitters.selector);
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_FMustBePositive_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 0,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(CCIPCapabilityConfiguration.FMustBePositive.selector);
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_FTooHigh_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);

    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 2,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(CCIPCapabilityConfiguration.FTooHigh.selector);
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_SignerP2PIdPairMustBeLengthTwo_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);
    signers[0] = new bytes[](1);

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.SignerP2PIdPairMustBeLengthTwo.selector, 1));
    s_ccipCC.validateConfig(config);
  }

  function test__validateConfig_NodeNotInRegistry_Reverts() public {
    (bytes[][] memory signers, bytes[][] memory transmitters) = addChainConfig(4);
    bytes32 nonExistentP2PId = keccak256("notInRegistry");
    signers[0][0] = abi.encode(nonExistentP2PId);

    vm.mockCall(
      CAPABILITY_REGISTRY,
      abi.encodeWithSelector(ICapabilityRegistry.getNode.selector, nonExistentP2PId),
      abi.encode(
        ICapabilityRegistry.NodeParams({
          nodeOperatorId: 0,
          signer: address(0),
          p2pId: bytes32(uint256(0)),
          hashedCapabilityIds: new bytes32[](0)
        }),
        uint32(1)
      )
    );

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.NodeNotInRegistry.selector, nonExistentP2PId));
    s_ccipCC.validateConfig(config);
  }
}

contract CCIPCapabilityConfiguration_ConfigStateMachine is CCIPCapabilityConfigurationSetup {
  // Successful cases.

  function test__stateFromConfigLength_Success() public {
    uint256 configLen = 0;
    CCIPCapabilityConfiguration.ConfigState state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Init));

    configLen = 1;
    state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Running));

    configLen = 2;
    state = s_ccipCC.stateFromConfigLength(configLen);
    assertEq(uint256(state), uint256(CCIPCapabilityConfiguration.ConfigState.Staging));
  }

  function test__validateConfigStateTransition_Success() public {
    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Init, CCIPCapabilityConfiguration.ConfigState.Running
    );

    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Running, CCIPCapabilityConfiguration.ConfigState.Staging
    );

    s_ccipCC.validateConfigStateTransition(
      CCIPCapabilityConfiguration.ConfigState.Staging, CCIPCapabilityConfiguration.ConfigState.Running
    );
  }

  function test__computeConfigDigest_Success() public {
    // config digest must change upon:
    // - ocr config change (e.g plugin type, chain selector, etc.)
    // - don id change
    // - config count change
    bytes[][] memory signers = makeAssociativeArray(2, 10);
    bytes[][] memory transmitters = makeAssociativeArray(2, 20);
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: transmitters,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });
    uint32 donId = 1;
    uint32 configCount = 1;

    bytes32 configDigest1 = s_ccipCC.computeConfigDigest(donId, configCount, config);

    donId = 2;
    bytes32 configDigest2 = s_ccipCC.computeConfigDigest(donId, configCount, config);

    donId = 1;
    configCount = 2;
    bytes32 configDigest3 = s_ccipCC.computeConfigDigest(donId, configCount, config);

    configCount = 1;
    config.pluginType = CCIPCapabilityConfiguration.PluginType.Execution;
    bytes32 configDigest4 = s_ccipCC.computeConfigDigest(donId, configCount, config);

    assertNotEq(configDigest1, configDigest2, "config digests 1 and 2 must not match");
    assertNotEq(configDigest1, configDigest3, "config digests 1 and 3 must not match");
    assertNotEq(configDigest1, configDigest4, "config digests 1 and 4 must not match");

    assertNotEq(configDigest2, configDigest3, "config digests 2 and 3 must not match");
    assertNotEq(configDigest2, configDigest4, "config digests 2 and 4 must not match");
  }

  function test_Fuzz__groupByPluginType_Success(uint256 numCommitCfgs, uint256 numExecCfgs) public {
    vm.assume(numCommitCfgs >= 0 && numCommitCfgs < 3);
    vm.assume(numExecCfgs >= 0 && numExecCfgs < 3);

    bytes[][] memory signers = makeAssociativeArray(4, 10);
    bytes[][] memory transmitters = makeAssociativeArray(4, 20);
    CCIPCapabilityConfiguration.OCR3Config[] memory cfgs =
      new CCIPCapabilityConfiguration.OCR3Config[](numCommitCfgs + numExecCfgs);
    for (uint256 i = 0; i < numCommitCfgs; i++) {
      cfgs[i] = CCIPCapabilityConfiguration.OCR3Config({
        pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
        chainSelector: 1,
        signers: signers,
        transmitters: transmitters,
        f: 1,
        offchainConfigVersion: 30,
        offchainConfig: abi.encode("commit", i)
      });
    }
    for (uint256 i = 0; i < numExecCfgs; i++) {
      cfgs[numCommitCfgs + i] = CCIPCapabilityConfiguration.OCR3Config({
        pluginType: CCIPCapabilityConfiguration.PluginType.Execution,
        chainSelector: 1,
        signers: signers,
        transmitters: transmitters,
        f: 1,
        offchainConfigVersion: 30,
        offchainConfig: abi.encode("exec", numCommitCfgs + i)
      });
    }
    (
      CCIPCapabilityConfiguration.OCR3Config[] memory commitCfgs,
      CCIPCapabilityConfiguration.OCR3Config[] memory execCfgs
    ) = s_ccipCC.groupByPluginType(cfgs);

    assertEq(commitCfgs.length, numCommitCfgs, "commitCfgs length must match");
    assertEq(execCfgs.length, numExecCfgs, "execCfgs length must match");
    for (uint256 i = 0; i < commitCfgs.length; i++) {
      assertEq(
        uint8(commitCfgs[i].pluginType),
        uint8(CCIPCapabilityConfiguration.PluginType.Commit),
        "plugin type must be commit"
      );
      assertEq(commitCfgs[i].offchainConfig, abi.encode("commit", i), "offchain config must match");
    }
    for (uint256 i = 0; i < execCfgs.length; i++) {
      assertEq(
        uint8(execCfgs[i].pluginType),
        uint8(CCIPCapabilityConfiguration.PluginType.Execution),
        "plugin type must be execution"
      );
      assertEq(execCfgs[i].offchainConfig, abi.encode("exec", numCommitCfgs + i), "offchain config must match");
    }
  }

  function test__computeNewConfigWithMeta_InitToRunning_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](0);
    CCIPCapabilityConfiguration.OCR3Config[] memory newConfig = new CCIPCapabilityConfiguration.OCR3Config[](1);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.ConfigState currentState = CCIPCapabilityConfiguration.ConfigState.Init;
    CCIPCapabilityConfiguration.ConfigState newState = CCIPCapabilityConfiguration.ConfigState.Running;
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfigWithMeta =
      s_ccipCC.computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, newState);
    assertEq(newConfigWithMeta.length, 1, "new config with meta length must be 1");
    assertEq(newConfigWithMeta[0].configCount, uint64(1), "config count must be 1");
    assertEq(uint8(newConfigWithMeta[0].config.pluginType), uint8(newConfig[0].pluginType), "plugin type must match");
    assertEq(newConfigWithMeta[0].config.offchainConfig, newConfig[0].offchainConfig, "offchain config must match");
    assertEq(
      newConfigWithMeta[0].configDigest,
      s_ccipCC.computeConfigDigest(donId, 1, newConfig[0]),
      "config digest must match"
    );

    // This ensures that the test case is using correct inputs.
    s_ccipCC.validateConfigTransition(currentConfig, newConfigWithMeta);
  }

  function test__computeNewConfigWithMeta_RunningToStaging_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });

    CCIPCapabilityConfiguration.OCR3Config[] memory newConfig = new CCIPCapabilityConfiguration.OCR3Config[](2);
    // existing blue config first.
    newConfig[0] = blueConfig;
    // green config next.
    newConfig[1] = greenConfig;

    CCIPCapabilityConfiguration.ConfigState currentState = CCIPCapabilityConfiguration.ConfigState.Running;
    CCIPCapabilityConfiguration.ConfigState newState = CCIPCapabilityConfiguration.ConfigState.Staging;

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfigWithMeta =
      s_ccipCC.computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, newState);
    assertEq(newConfigWithMeta.length, 2, "new config with meta length must be 2");

    assertEq(newConfigWithMeta[0].configCount, uint64(1), "config count of blue must be 1");
    assertEq(
      uint8(newConfigWithMeta[0].config.pluginType), uint8(blueConfig.pluginType), "plugin type of blue must match"
    );
    assertEq(
      newConfigWithMeta[0].config.offchainConfig, blueConfig.offchainConfig, "offchain config of blue must match"
    );
    assertEq(
      newConfigWithMeta[0].configDigest,
      s_ccipCC.computeConfigDigest(donId, 1, blueConfig),
      "config digest of blue must match"
    );

    assertEq(newConfigWithMeta[1].configCount, uint64(2), "config count of green must be 2");
    assertEq(
      uint8(newConfigWithMeta[1].config.pluginType), uint8(greenConfig.pluginType), "plugin type of green must match"
    );
    assertEq(
      newConfigWithMeta[1].config.offchainConfig, greenConfig.offchainConfig, "offchain config of green must match"
    );
    assertEq(
      newConfigWithMeta[1].configDigest,
      s_ccipCC.computeConfigDigest(donId, 2, greenConfig),
      "config digest of green must match"
    );

    // This ensures that the test case is using correct inputs.
    s_ccipCC.validateConfigTransition(currentConfig, newConfigWithMeta);
  }

  function test__computeNewConfigWithMeta_StagingToRunning_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    currentConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });
    CCIPCapabilityConfiguration.OCR3Config[] memory newConfig = new CCIPCapabilityConfiguration.OCR3Config[](1);
    newConfig[0] = greenConfig;

    CCIPCapabilityConfiguration.ConfigState currentState = CCIPCapabilityConfiguration.ConfigState.Staging;
    CCIPCapabilityConfiguration.ConfigState newState = CCIPCapabilityConfiguration.ConfigState.Running;

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfigWithMeta =
      s_ccipCC.computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, newState);

    assertEq(newConfigWithMeta.length, 1, "new config with meta length must be 1");
    assertEq(newConfigWithMeta[0].configCount, uint64(2), "config count must be 2");
    assertEq(uint8(newConfigWithMeta[0].config.pluginType), uint8(greenConfig.pluginType), "plugin type must match");
    assertEq(newConfigWithMeta[0].config.offchainConfig, greenConfig.offchainConfig, "offchain config must match");
    assertEq(
      newConfigWithMeta[0].configDigest, s_ccipCC.computeConfigDigest(donId, 2, greenConfig), "config digest must match"
    );

    // This ensures that the test case is using correct inputs.
    s_ccipCC.validateConfigTransition(currentConfig, newConfigWithMeta);
  }

  function test__validateConfigTransition_InitToRunning_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](0);

    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_RunningToStaging_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    newConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });

    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_StagingToRunning_Success() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    currentConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });

    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  // Reverts.

  function test_Fuzz__stateFromConfigLength_Reverts(uint256 configLen) public {
    vm.assume(configLen > 2);
    vm.expectRevert(CCIPCapabilityConfiguration.InvalidConfigLength.selector);
    s_ccipCC.stateFromConfigLength(configLen);
  }

  function test__groupByPluginType_threeCommitConfigs_Reverts() public {
    bytes[][] memory signers = makeAssociativeArray(4, 10);
    bytes[][] memory transmitters = makeAssociativeArray(4, 20);
    CCIPCapabilityConfiguration.OCR3Config[] memory cfgs = new CCIPCapabilityConfiguration.OCR3Config[](3);
    for (uint256 i = 0; i < 3; i++) {
      cfgs[i] = CCIPCapabilityConfiguration.OCR3Config({
        pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
        chainSelector: 1,
        signers: signers,
        transmitters: transmitters,
        f: 1,
        offchainConfigVersion: 30,
        offchainConfig: abi.encode("commit", i)
      });
    }
    vm.expectRevert();
    s_ccipCC.groupByPluginType(cfgs);
  }

  function test__groupByPluginType_threeExecutionConfigs_Reverts() public {
    bytes[][] memory signers = makeAssociativeArray(4, 10);
    bytes[][] memory transmitters = makeAssociativeArray(4, 20);
    CCIPCapabilityConfiguration.OCR3Config[] memory cfgs = new CCIPCapabilityConfiguration.OCR3Config[](3);
    for (uint256 i = 0; i < 3; i++) {
      cfgs[i] = CCIPCapabilityConfiguration.OCR3Config({
        pluginType: CCIPCapabilityConfiguration.PluginType.Execution,
        chainSelector: 1,
        signers: signers,
        transmitters: transmitters,
        f: 1,
        offchainConfigVersion: 30,
        offchainConfig: abi.encode("exec", i)
      });
    }
    vm.expectRevert();
    s_ccipCC.groupByPluginType(cfgs);
  }

  function test__groupByPluginType_TooManyOCR3Configs_Reverts() public {
    CCIPCapabilityConfiguration.OCR3Config[] memory cfgs = new CCIPCapabilityConfiguration.OCR3Config[](5);
    vm.expectRevert(CCIPCapabilityConfiguration.TooManyOCR3Configs.selector);
    s_ccipCC.groupByPluginType(cfgs);
  }

  function test__validateConfigTransition_InitToRunning_WrongConfigCount_Reverts() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 0,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](0);

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.WrongConfigCount.selector, 0, 1));
    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_RunningToStaging_WrongConfigDigestBlueGreen_Reverts() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 3, blueConfig) // wrong config digest (due to diff config count)
    });
    newConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        CCIPCapabilityConfiguration.WrongConfigDigestBlueGreen.selector,
        s_ccipCC.computeConfigDigest(donId, 3, blueConfig),
        s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
      )
    );
    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_RunningToStaging_WrongConfigCount_Reverts() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    newConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 3, // wrong config count
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 3, greenConfig)
    });

    vm.expectRevert(abi.encodeWithSelector(CCIPCapabilityConfiguration.WrongConfigCount.selector, 3, 2));
    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_StagingToRunning_WrongConfigDigest_Reverts() public {
    uint32 donId = 1;
    CCIPCapabilityConfiguration.OCR3Config memory blueConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit")
    });
    CCIPCapabilityConfiguration.OCR3Config memory greenConfig = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: makeAssociativeArray(4, 10),
      transmitters: makeAssociativeArray(4, 20),
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("commit-new")
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](2);
    currentConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 1,
      config: blueConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 1, blueConfig)
    });
    currentConfig[1] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
    });

    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    newConfig[0] = CCIPCapabilityConfiguration.OCR3ConfigWithMeta({
      configCount: 2,
      config: greenConfig,
      configDigest: s_ccipCC.computeConfigDigest(donId, 3, greenConfig) // wrong config digest
    });

    vm.expectRevert(
      abi.encodeWithSelector(
        CCIPCapabilityConfiguration.WrongConfigDigest.selector,
        s_ccipCC.computeConfigDigest(donId, 3, greenConfig),
        s_ccipCC.computeConfigDigest(donId, 2, greenConfig)
      )
    );
    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }

  function test__validateConfigTransition_NonExistentConfigTransition_Reverts() public {
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory currentConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](3);
    CCIPCapabilityConfiguration.OCR3ConfigWithMeta[] memory newConfig =
      new CCIPCapabilityConfiguration.OCR3ConfigWithMeta[](1);
    vm.expectRevert(CCIPCapabilityConfiguration.NonExistentConfigTransition.selector);
    s_ccipCC.validateConfigTransition(currentConfig, newConfig);
  }
}
