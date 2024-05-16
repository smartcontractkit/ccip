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

  function test_applyChainConfigUpdates_removeChainConfigs_Success() public {}

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

  function test__validateConfig_Success() public {
    uint256 numNodes = 4;
    bytes[][] memory signers = makeAssociativeArray(numNodes, 10);
    bytes[][] memory transmitters = makeAssociativeArray(numNodes, 20);
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

    // Config is for 4 nodes, so f == 1.
    CCIPCapabilityConfiguration.OCR3Config memory config = CCIPCapabilityConfiguration.OCR3Config({
      pluginType: CCIPCapabilityConfiguration.PluginType.Commit,
      chainSelector: 1,
      signers: signers,
      transmitters: signers,
      f: 1,
      offchainConfigVersion: 30,
      offchainConfig: bytes("offchainConfig")
    });
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

    assertNotEq(configDigest1, configDigest2, "config digests must not match");
    assertNotEq(configDigest1, configDigest3, "config digests must not match");
    assertNotEq(configDigest1, configDigest4, "config digests must not match");

    assertNotEq(configDigest2, configDigest3, "config digests must not match");
    assertNotEq(configDigest2, configDigest4, "config digests must not match");
  }

  function test__groupByPluginType_Success() public {}

  // Reverts.

  function test_Fuzz__stateFromConfigLength_Reverts(uint256 configLen) public {
    vm.assume(configLen > 2);
    vm.expectRevert(CCIPCapabilityConfiguration.InvalidConfigLength.selector);
    s_ccipCC.stateFromConfigLength(configLen);
  }
}
