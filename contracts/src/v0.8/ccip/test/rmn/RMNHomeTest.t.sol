// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {RMNHome} from "../../rmn/RMNHome.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

contract RMNHomeTest is Test {
  bytes32 internal constant ZERO_DIGEST = bytes32(uint256(0));

  RMNHome public s_rmnHome;

  function setUp() public virtual {
    s_rmnHome = new RMNHome();
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

    return RMNHome.Config({
      staticConfig: RMNHome.StaticConfig({nodes: nodes, offchainConfig: abi.encode("static_config")}),
      dynamicConfig: RMNHome.DynamicConfig({sourceChains: sourceChains, offchainConfig: abi.encode("dynamic_config")})
    });
  }

  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 private constant PREFIX = 0x000b << (256 - 16); // 0x000b00..00

  function _getConfigDigest(RMNHome.StaticConfig memory staticConfig, uint32 version) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(keccak256(abi.encode(bytes32("EVM"), block.chainid, address(s_rmnHome), version, staticConfig)))
            & ~PREFIX_MASK
        )
    );
  }
}

contract RMNHome_setSecondary is RMNHomeTest {
  function test_setSecondary_success() public {
    RMNHome.Config memory config = _getBaseConfig();
    RMNHome.VersionedConfig memory versionedConfig = RMNHome.VersionedConfig({version: 1, config: config});
    bytes32 configDigest = _getConfigDigest(versionedConfig.config.staticConfig, versionedConfig.version);

    vm.expectEmit();
    emit RMNHome.ConfigSet(configDigest, versionedConfig);

    s_rmnHome.setSecondary(config, ZERO_DIGEST);

    (RMNHome.VersionedConfig memory storedVersionedConfig, bool ok) = s_rmnHome.getConfig(configDigest);
    assertTrue(ok);
    assertEq(storedVersionedConfig.version, versionedConfig.version);
    RMNHome.StaticConfig memory storedStaticConfig = storedVersionedConfig.config.staticConfig;
    RMNHome.DynamicConfig memory storedDynamicConfig = storedVersionedConfig.config.dynamicConfig;

    assertEq(storedStaticConfig.nodes.length, versionedConfig.config.staticConfig.nodes.length);
    for (uint256 i = 0; i < storedStaticConfig.nodes.length; i++) {
      RMNHome.Node memory storedNode = storedStaticConfig.nodes[i];
      assertEq(storedNode.peerId, versionedConfig.config.staticConfig.nodes[i].peerId);
      assertEq(storedNode.offchainPublicKey, versionedConfig.config.staticConfig.nodes[i].offchainPublicKey);
    }

    assertEq(storedDynamicConfig.sourceChains.length, versionedConfig.config.dynamicConfig.sourceChains.length);
    for (uint256 i = 0; i < storedDynamicConfig.sourceChains.length; i++) {
      RMNHome.SourceChain memory storedSourceChain = storedDynamicConfig.sourceChains[i];
      assertEq(storedSourceChain.chainSelector, versionedConfig.config.dynamicConfig.sourceChains[i].chainSelector);
      assertEq(storedSourceChain.minObservers, versionedConfig.config.dynamicConfig.sourceChains[i].minObservers);
      assertEq(
        storedSourceChain.observerNodesBitmap, versionedConfig.config.dynamicConfig.sourceChains[i].observerNodesBitmap
      );
    }
    assertEq(storedDynamicConfig.offchainConfig, versionedConfig.config.dynamicConfig.offchainConfig);
    assertEq(storedStaticConfig.offchainConfig, versionedConfig.config.staticConfig.offchainConfig);
  }

  function test_setSecondary_OutOfBoundsNodesLength_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.staticConfig.nodes = new RMNHome.Node[](257);

    vm.expectRevert(RMNHome.OutOfBoundsNodesLength.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_DuplicatePeerId_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.staticConfig.nodes[1].peerId = config.staticConfig.nodes[0].peerId;

    vm.expectRevert(RMNHome.DuplicatePeerId.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_DuplicateOffchainPublicKey_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.staticConfig.nodes[1].offchainPublicKey = config.staticConfig.nodes[0].offchainPublicKey;

    vm.expectRevert(RMNHome.DuplicateOffchainPublicKey.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_DuplicateSourceChain_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[1].chainSelector = config.dynamicConfig.sourceChains[0].chainSelector;

    vm.expectRevert(RMNHome.DuplicateSourceChain.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_OutOfBoundsObserverNodeIndex_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[0].observerNodesBitmap = 1 << config.staticConfig.nodes.length;

    vm.expectRevert(RMNHome.OutOfBoundsObserverNodeIndex.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_MinObserversTooHigh_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[0].minObservers++;

    vm.expectRevert(RMNHome.MinObserversTooHigh.selector);
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }

  function test_setSecondary_OnlyOwner_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();

    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.setSecondary(config, ZERO_DIGEST);
  }
}

contract RMNHome_setDynamicConfig is RMNHomeTest {
  function setUp() public override {
    super.setUp();
    s_rmnHome.setSecondary(_getBaseConfig(), ZERO_DIGEST);
  }

  function test_setDynamicConfig_success() public {
    (bytes32 priorPrimaryDigest,) = s_rmnHome.getConfigDigests();

    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[0].minObservers--;

    (, bytes32 secondaryConfigDigest) = s_rmnHome.getConfigDigests();

    vm.expectEmit();
    emit RMNHome.DynamicConfigSet(secondaryConfigDigest, config.dynamicConfig);

    s_rmnHome.setDynamicConfig(config.dynamicConfig, secondaryConfigDigest);

    (RMNHome.VersionedConfig memory storedVersionedConfig, bool ok) = s_rmnHome.getConfig(secondaryConfigDigest);
    assertTrue(ok);
    assertEq(
      storedVersionedConfig.config.dynamicConfig.sourceChains[0].minObservers,
      config.dynamicConfig.sourceChains[0].minObservers
    );

    // Asser the digests don't change when updating the dynamic config
    (bytes32 primaryDigest, bytes32 secondaryDigest) = s_rmnHome.getConfigDigests();
    assertEq(primaryDigest, priorPrimaryDigest);
    assertEq(secondaryDigest, secondaryConfigDigest);
  }

  // Asserts the validation function is being called
  function test_setDynamicConfig_MinObserversTooHigh_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[0].minObservers++;

    vm.expectRevert(abi.encodeWithSelector(RMNHome.DigestNotFound.selector, ZERO_DIGEST));
    s_rmnHome.setDynamicConfig(config.dynamicConfig, ZERO_DIGEST);
  }

  function test_setDynamicConfig_DigestNotFound_reverts() public {
    // Zero always reverts
    vm.expectRevert(abi.encodeWithSelector(RMNHome.DigestNotFound.selector, ZERO_DIGEST));
    s_rmnHome.setDynamicConfig(_getBaseConfig().dynamicConfig, ZERO_DIGEST);

    // Non-existent digest reverts
    bytes32 nonExistentDigest = keccak256("nonExistentDigest");
    vm.expectRevert(abi.encodeWithSelector(RMNHome.DigestNotFound.selector, nonExistentDigest));
    s_rmnHome.setDynamicConfig(_getBaseConfig().dynamicConfig, nonExistentDigest);
  }

  function test_setDynamicConfig_OnlyOwner_reverts() public {
    RMNHome.Config memory config = _getBaseConfig();

    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.setDynamicConfig(config.dynamicConfig, keccak256("configDigest"));
  }
}

contract RMNHome_revokeSecondary is RMNHomeTest {
  // Sets two configs
  function setUp() public override {
    super.setUp();
    bytes32 digest = s_rmnHome.setSecondary(_getBaseConfig(), ZERO_DIGEST);
    s_rmnHome.promoteSecondary(digest);

    RMNHome.Config memory config = _getBaseConfig();
    config.dynamicConfig.sourceChains[0].minObservers--;
    s_rmnHome.setSecondary(_getBaseConfig(), ZERO_DIGEST);
  }

  function test_revokeSecondary_success() public {
    (bytes32 priorPrimaryDigest, bytes32 priorSecondaryDigest) = s_rmnHome.getConfigDigests();

    vm.expectEmit();
    emit RMNHome.ConfigRevoked(priorSecondaryDigest);

    s_rmnHome.revokeSecondary(priorSecondaryDigest);

    (RMNHome.VersionedConfig memory storedVersionedConfig, bool ok) = s_rmnHome.getConfig(priorSecondaryDigest);
    assertFalse(ok);
    // Ensure no old data is returned, even though it's still in storage
    assertEq(storedVersionedConfig.version, 0);
    assertEq(storedVersionedConfig.config.staticConfig.nodes.length, 0);
    assertEq(storedVersionedConfig.config.dynamicConfig.sourceChains.length, 0);

    // Asser the primary digest is unaffected but the secondary digest is set to zero
    (bytes32 primaryDigest, bytes32 secondaryDigest) = s_rmnHome.getConfigDigests();
    assertEq(primaryDigest, priorPrimaryDigest);
    assertTrue(secondaryDigest != priorSecondaryDigest);
    assertEq(secondaryDigest, ZERO_DIGEST);
  }

  function test_revokeSecondary_ConfigDigestMismatch_reverts() public {
    (, bytes32 priorSecondaryDigest) = s_rmnHome.getConfigDigests();

    bytes32 wrongDigest = keccak256("wrong_digest");
    vm.expectRevert(abi.encodeWithSelector(RMNHome.ConfigDigestMismatch.selector, priorSecondaryDigest, wrongDigest));
    s_rmnHome.revokeSecondary(wrongDigest);
  }

  function test_revokeSecondary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.revokeSecondary(keccak256("configDigest"));
  }
}

contract RMNHome_promoteSecondary is RMNHomeTest {
  function test_promoteSecondary_success() public {
    (bytes32 priorPrimaryDigest, bytes32 priorSecondaryDigest) = s_rmnHome.getConfigDigests();

    vm.expectEmit();
    emit RMNHome.ConfigPromoted(priorSecondaryDigest);

    s_rmnHome.promoteSecondary(priorSecondaryDigest);

    (bytes32 primaryDigest, bytes32 secondaryDigest) = s_rmnHome.getConfigDigests();
    assertEq(primaryDigest, priorSecondaryDigest);
    assertEq(secondaryDigest, priorPrimaryDigest);
  }

  function test_promoteSecondary_ConfigDigestMismatch_reverts() public {
    (, bytes32 priorSecondaryDigest) = s_rmnHome.getConfigDigests();

    bytes32 wrongDigest = keccak256("wrong_digest");
    vm.expectRevert(abi.encodeWithSelector(RMNHome.ConfigDigestMismatch.selector, priorSecondaryDigest, wrongDigest));
    s_rmnHome.promoteSecondary(wrongDigest);
  }

  function test_promoteSecondary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.promoteSecondary(keccak256("configDigest"));
  }
}

contract RMNHome_promoteSecondaryAndRevokePrimary is RMNHomeTest {
  function test_promoteSecondaryAndRevokePrimary_success() public {}

  function test_promoteSecondaryAndRevokePrimary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert("Only callable by owner");
    s_rmnHome.promoteSecondaryAndRevokePrimary(keccak256("toPromote"), keccak256("ToRevoke"));
  }
}
