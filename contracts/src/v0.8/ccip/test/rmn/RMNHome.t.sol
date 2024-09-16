// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {RMNHome} from "../../rmn/RMNHome.sol";
import {BaseTest} from "../BaseTest.t.sol";
import {Vm} from "forge-std/Vm.sol";

contract RMNHomeTest is BaseTest {
  RMNHome public s_rmnHome;

  function setUp() public virtual override {
    super.setUp();
    s_rmnHome = new RMNHome();

    s_rmnHome.revokeAllConfigsButLatest();
  }

  function _getBaseConfig() internal pure returns (RMNHome.Config memory) {
    RMNHome.Node[] memory nodes = new RMNHome.Node[](3);
    nodes[0] = RMNHome.Node({peerId: keccak256("peerId_0"), offchainPublicKey: keccak256("offchainPublicKey_0")});
    nodes[1] = RMNHome.Node({peerId: keccak256("peerId_1"), offchainPublicKey: keccak256("offchainPublicKey_1")});
    nodes[2] = RMNHome.Node({peerId: keccak256("peerId_2"), offchainPublicKey: keccak256("offchainPublicKey_2")});

    RMNHome.SourceChain[] memory sourceChains = new RMNHome.SourceChain[](2);
    // Observer 0 for source chain 9000
    sourceChains[0] = RMNHome.SourceChain({chainSelector: 9000, minObservers: 1, observerNodesBitmap: 1 << 0});
    // Observers 1 and 2 for source chain 9001
    sourceChains[1] = RMNHome.SourceChain({chainSelector: 9001, minObservers: 2, observerNodesBitmap: 1 << 1 | 1 << 2});

    return RMNHome.Config({nodes: nodes, sourceChains: sourceChains, offchainConfig: abi.encode("offchainConfig")});
  }

  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 private constant PREFIX = 0x000b << (256 - 16); // 0x000b00..00

  function _getConfigDigest(RMNHome.VersionedConfig memory versionedConfig) internal pure returns (bytes32) {
    return bytes32((PREFIX & PREFIX_MASK) | (uint256(keccak256(abi.encode(versionedConfig))) & ~PREFIX_MASK));
  }
}

contract RMNHome_setConfig is RMNHomeTest {
  function test_setConfig_success() public {
    RMNHome.Config memory config = _getBaseConfig();
    RMNHome.VersionedConfig memory versionedConfig = RMNHome.VersionedConfig({version: 1, config: config});
    bytes32 configDigest = _getConfigDigest(versionedConfig);

    vm.expectEmit();
    emit RMNHome.ConfigSet(configDigest, versionedConfig);

    s_rmnHome.setConfig(config);

    (RMNHome.VersionedConfig memory storedVersionedConfig, bool ok) = s_rmnHome.getVersionedConfig(configDigest);
    assertTrue(ok);
    assertEq(storedVersionedConfig.version, versionedConfig.version);
    assertEq(storedVersionedConfig.config.nodes.length, versionedConfig.config.nodes.length);
    for (uint256 i = 0; i < storedVersionedConfig.config.nodes.length; i++) {
      RMNHome.Node memory storedNode = storedVersionedConfig.config.nodes[i];
      assertEq(storedNode.peerId, versionedConfig.config.nodes[i].peerId);
      assertEq(storedNode.offchainPublicKey, versionedConfig.config.nodes[i].offchainPublicKey);
    }

    assertEq(storedVersionedConfig.config.sourceChains.length, versionedConfig.config.sourceChains.length);
    for (uint256 i = 0; i < storedVersionedConfig.config.sourceChains.length; i++) {
      RMNHome.SourceChain memory storedSourceChain = storedVersionedConfig.config.sourceChains[i];
      assertEq(storedSourceChain.chainSelector, versionedConfig.config.sourceChains[i].chainSelector);
      assertEq(storedSourceChain.minObservers, versionedConfig.config.sourceChains[i].minObservers);
      assertEq(storedSourceChain.observerNodesBitmap, versionedConfig.config.sourceChains[i].observerNodesBitmap);
    }
    assertEq(storedVersionedConfig.config.offchainConfig, versionedConfig.config.offchainConfig);
  }

  function test_setConfig_moreThanRingBuffer_success() public {
    RMNHome.Config memory config = _getBaseConfig();
    uint256 ringBufferSize = s_rmnHome.getRingBufferSize();

    vm.recordLogs();

    // Insert ring buffer size + 1 configs
    for (uint256 i = 0; i < ringBufferSize + 1; i++) {
      s_rmnHome.setConfig(config);
    }

    Vm.Log[] memory entries = vm.getRecordedLogs();
    vm.assertEq(entries.length, ringBufferSize + 2);

    for (uint256 i = 0; i < ringBufferSize; i++) {
      Vm.Log memory entry = entries[i];
      vm.assertEq(entry.topics[0], bytes32(uint256(RMNHome.ConfigSet.selector)));
    }

    // The last config should revoke the first config, which means it should emit a ConfigRevoked event
    vm.assertEq(entries[ringBufferSize].topics[0], bytes32(uint256(RMNHome.ConfigRevoked.selector)));
    vm.assertEq(entries[ringBufferSize + 1].topics[0], bytes32(uint256(RMNHome.ConfigSet.selector)));
  }

  function test_setConfig_OutOfBoundsNodesLength_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.nodes = new RMNHome.Node[](257);

    vm.expectRevert(RMNHome.OutOfBoundsNodesLength.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_DuplicatePeerId_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.nodes[1].peerId = config.nodes[0].peerId;

    vm.expectRevert(RMNHome.DuplicatePeerId.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_DuplicateOffchainPublicKey_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.nodes[1].offchainPublicKey = config.nodes[0].offchainPublicKey;

    vm.expectRevert(RMNHome.DuplicateOffchainPublicKey.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_DuplicateSourceChain_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.sourceChains[1].chainSelector = config.sourceChains[0].chainSelector;

    vm.expectRevert(RMNHome.DuplicateSourceChain.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_OutOfBoundsObserverNodeIndex_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.sourceChains[0].observerNodesBitmap = 1 << config.nodes.length;

    vm.expectRevert(RMNHome.OutOfBoundsObserverNodeIndex.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_MinObserversTooHigh_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.sourceChains[0].minObservers++;

    vm.expectRevert(RMNHome.MinObserversTooHigh.selector);
    s_rmnHome.setConfig(config);
  }

  function test_setConfig_OnlyOwner_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();

    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.setConfig(config);
  }
}

contract RMNHome_revokeAllConfigsButLatest is RMNHomeTest {
  function test_revokeAllConfigsButLatest_success() public {}

  function test_setConfig_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.revokeAllConfigsButLatest();
  }
}

contract RMNHome_revokeConfig is RMNHomeTest {
  function test_revokeConfig_success() public {}

  function test_setConfig_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.revokeConfig(keccak256("configDigest"));
  }
}
