// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilityConfiguration} from "../../../keystone/interfaces/ICapabilityConfiguration.sol";
import {ICapabilitiesRegistry} from "../../capability/interfaces//ICapabilitiesRegistry.sol";

import {CCIPHome} from "../../capability/CCIPHome.sol";
import {Internal} from "../../libraries/Internal.sol";
import {CCIPHomeHelper} from "../helpers/CCIPHomeHelper.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

import {IERC165} from "../../../vendor/openzeppelin-solidity/v5.0.2/contracts/interfaces/IERC165.sol";

contract CCIPHomeTest is Test {
  bytes32 internal constant ZERO_DIGEST = bytes32(uint256(0));
  address internal constant CAPABILITIES_REGISTRY = address(0x0000000123123123123);
  Internal.OCRPluginType internal constant DEFAULT_PLUGIN_TYPE = Internal.OCRPluginType.Commit;
  uint32 internal constant DEFAULT_DON_ID = 78978987;

  CCIPHomeHelper public s_ccipHome = new CCIPHomeHelper(CAPABILITIES_REGISTRY);

  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 private constant PREFIX = 0x000a << (256 - 16); // 0x000b00..00

  uint64 private constant DEFAULT_CHAIN_SELECTOR = 9381579735;

  function setUp() public virtual {
    s_ccipHome.applyChainConfigUpdates(new uint64[](0), _getBaseChainConfigs());

    ICapabilitiesRegistry.NodeInfo memory nodeInfo = ICapabilitiesRegistry.NodeInfo({
      p2pId: keccak256("p2pId"),
      signer: keccak256("signer"),
      nodeOperatorId: 1,
      configCount: 1,
      workflowDONId: 1,
      hashedCapabilityIds: new bytes32[](0),
      capabilitiesDONIds: new uint256[](0)
    });

    vm.mockCall(
      CAPABILITIES_REGISTRY, abi.encodeWithSelector(ICapabilitiesRegistry.getNode.selector), abi.encode(nodeInfo)
    );

    vm.startPrank(address(s_ccipHome));
  }

  function _getBaseChainConfigs() internal pure returns (CCIPHome.ChainConfigArgs[] memory) {
    CCIPHome.ChainConfigArgs[] memory configs = new CCIPHome.ChainConfigArgs[](1);
    CCIPHome.ChainConfig memory chainConfig =
      CCIPHome.ChainConfig({readers: new bytes32[](0), fChain: 1, config: abi.encode("chainConfig")});
    configs[0] = CCIPHome.ChainConfigArgs({chainSelector: DEFAULT_CHAIN_SELECTOR, chainConfig: chainConfig});

    return configs;
  }

  function _getConfigDigest(
    uint32 donId,
    Internal.OCRPluginType pluginType,
    bytes memory config,
    uint32 version
  ) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(
            keccak256(
              bytes.concat(
                abi.encode(bytes32("EVM"), block.chainid, address(s_ccipHome), donId, pluginType, version), config
              )
            )
          ) & ~PREFIX_MASK
        )
    );
  }

  function _getBaseConfig() internal pure returns (CCIPHome.OCR3Config memory) {
    CCIPHome.OCR3Node[] memory nodes = new CCIPHome.OCR3Node[](4);
    for (uint256 i = 0; i < nodes.length; i++) {
      nodes[i] = CCIPHome.OCR3Node({
        p2pId: keccak256(abi.encode("p2pId", i)),
        signerKey: abi.encode("signerKey"),
        transmitterKey: abi.encode("transmitterKey")
      });
    }

    return CCIPHome.OCR3Config({
      pluginType: Internal.OCRPluginType.Commit,
      chainSelector: DEFAULT_CHAIN_SELECTOR,
      FRoleDON: 1,
      offchainConfigVersion: 98765,
      offrampAddress: abi.encode("offrampAddress"),
      nodes: nodes,
      offchainConfig: abi.encode("offchainConfig")
    });
  }
}

contract CCIPHome_constructor is CCIPHomeTest {
  function test_constructor_success() public {
    CCIPHome ccipHome = new CCIPHome(CAPABILITIES_REGISTRY);

    assertEq(address(ccipHome.getCapabilityRegistry()), CAPABILITIES_REGISTRY);
  }

  function test_supportsInterface_success() public view {
    assertTrue(s_ccipHome.supportsInterface(type(IERC165).interfaceId));
    assertTrue(s_ccipHome.supportsInterface(type(ICapabilityConfiguration).interfaceId));
  }

  function test_getCapabilityConfiguration_success() public view {
    bytes memory config = s_ccipHome.getCapabilityConfiguration(DEFAULT_DON_ID);
    assertEq(config.length, 0);
  }

  function test_constructor_CapabilitiesRegistryAddressZero_reverts() public {
    vm.expectRevert(CCIPHome.ZeroAddressNotAllowed.selector);
    new CCIPHome(address(0));
  }
}

contract CCIPHome_beforeCapabilityConfigSet is CCIPHomeTest {
  function setUp() public virtual override {
    super.setUp();
    vm.stopPrank();
    vm.startPrank(address(CAPABILITIES_REGISTRY));
  }

  function test_beforeCapabilityConfigSet_success() public {
    // first set a config
    bytes memory callData =
      abi.encodeCall(CCIPHome.setCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, _getBaseConfig(), ZERO_DIGEST));

    vm.expectCall(address(s_ccipHome), callData);

    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);

    // Then revoke the config
    bytes32 candidateDigest = s_ccipHome.getCandidateDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertNotEq(candidateDigest, ZERO_DIGEST);

    callData = abi.encodeCall(CCIPHome.revokeCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, candidateDigest));

    vm.expectCall(address(s_ccipHome), callData);

    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);

    // Then set a new config
    callData =
      abi.encodeCall(CCIPHome.setCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, _getBaseConfig(), ZERO_DIGEST));

    vm.expectCall(address(s_ccipHome), callData);

    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);

    // Then promote the new config

    bytes32 newCandidateDigest = s_ccipHome.getCandidateDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertNotEq(newCandidateDigest, ZERO_DIGEST);

    callData = abi.encodeCall(
      CCIPHome.promoteCandidateAndRevokeActive, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, newCandidateDigest, ZERO_DIGEST)
    );

    vm.expectCall(address(s_ccipHome), callData);

    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);

    bytes32 activeDigest = s_ccipHome.getActiveDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, newCandidateDigest);
  }

  function test_beforeCapabilityConfigSet_OnlyCapabilitiesRegistryCanCall_reverts() public {
    bytes memory callData =
      abi.encodeCall(CCIPHome.setCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, _getBaseConfig(), ZERO_DIGEST));

    vm.stopPrank();

    vm.expectRevert(CCIPHome.OnlyCapabilitiesRegistryCanCall.selector);

    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);
  }

  function test_beforeCapabilityConfigSet_InvalidSelector_reverts() public {
    bytes memory callData = abi.encodeCall(CCIPHome.getConfigDigests, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE));

    vm.expectRevert(abi.encodeWithSelector(CCIPHome.InvalidSelector.selector, CCIPHome.getConfigDigests.selector));
    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);
  }

  function test_beforeCapabilityConfigSet_DONIdMismatch_reverts() public {
    uint32 wrongDonId = DEFAULT_DON_ID + 1;

    bytes memory callData =
      abi.encodeCall(CCIPHome.setCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, _getBaseConfig(), ZERO_DIGEST));

    vm.expectRevert(abi.encodeWithSelector(CCIPHome.DONIdMismatch.selector, DEFAULT_DON_ID, wrongDonId));
    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, wrongDonId);
  }

  function test_beforeCapabilityConfigSet_InnerCallReverts_reverts() public {
    bytes memory callData = abi.encodeCall(CCIPHome.revokeCandidate, (DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, ZERO_DIGEST));

    vm.expectRevert(CCIPHome.RevokingZeroDigestNotAllowed.selector);
    s_ccipHome.beforeCapabilityConfigSet(new bytes32[](0), callData, 0, DEFAULT_DON_ID);
  }
}

contract CCIPHome_getConfigDigests is CCIPHomeTest {
  function test_getConfigDigests_success() public {
    (bytes32 activeDigest, bytes32 candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, ZERO_DIGEST);
    assertEq(candidateDigest, ZERO_DIGEST);

    CCIPHome.OCR3Config memory config = _getBaseConfig();
    bytes32 firstDigest = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    (activeDigest, candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, ZERO_DIGEST);
    assertEq(candidateDigest, firstDigest);

    s_ccipHome.promoteCandidateAndRevokeActive(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, firstDigest, ZERO_DIGEST);

    (activeDigest, candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, firstDigest);
    assertEq(candidateDigest, ZERO_DIGEST);

    bytes32 secondDigest = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    (activeDigest, candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, firstDigest);
    assertEq(candidateDigest, secondDigest);

    assertEq(activeDigest, s_ccipHome.getActiveDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE));
    assertEq(candidateDigest, s_ccipHome.getCandidateDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE));
  }
}

contract CCIPHome_setCandidate is CCIPHomeTest {
  function test_setCandidate_success() public {
    CCIPHome.OCR3Config memory config = _getBaseConfig();
    CCIPHome.VersionedConfig memory versionedConfig =
      CCIPHome.VersionedConfig({version: 1, config: config, configDigest: ZERO_DIGEST});

    versionedConfig.configDigest =
      _getConfigDigest(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, abi.encode(versionedConfig.config), versionedConfig.version);

    vm.expectEmit();
    emit CCIPHome.ConfigSet(versionedConfig.configDigest, versionedConfig.version, versionedConfig.config);

    s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, versionedConfig.config, ZERO_DIGEST);

    (CCIPHome.VersionedConfig memory storedVersionedConfig, bool ok) =
      s_ccipHome.getConfig(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, versionedConfig.configDigest);
    assertTrue(ok);
    assertEq(storedVersionedConfig.version, versionedConfig.version);
    assertEq(storedVersionedConfig.configDigest, versionedConfig.configDigest);
    assertEq(keccak256(abi.encode(storedVersionedConfig.config)), keccak256(abi.encode(versionedConfig.config)));
  }

  function test_setCandidate_ConfigDigestMismatch_reverts() public {
    CCIPHome.OCR3Config memory config = _getBaseConfig();

    bytes32 digest = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    vm.expectRevert(abi.encodeWithSelector(CCIPHome.ConfigDigestMismatch.selector, digest, ZERO_DIGEST));
    s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    vm.expectEmit();
    emit CCIPHome.CandidateConfigRevoked(digest);

    s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, digest);
  }

  function test_setCandidate_CanOnlySelfCall_reverts() public {
    vm.stopPrank();

    vm.expectRevert(CCIPHome.CanOnlySelfCall.selector);
    s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, _getBaseConfig(), ZERO_DIGEST);
  }
}

contract CCIPHome_revokeCandidate is CCIPHomeTest {
  // Sets two configs
  function setUp() public virtual override {
    super.setUp();
    CCIPHome.OCR3Config memory config = _getBaseConfig();
    bytes32 digest = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);
    s_ccipHome.promoteCandidateAndRevokeActive(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, digest, ZERO_DIGEST);

    config.offrampAddress = abi.encode("new_offrampAddress");
    s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);
  }

  function test_revokeCandidate_success() public {
    (bytes32 priorActiveDigest, bytes32 priorCandidateDigest) =
      s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);

    vm.expectEmit();
    emit CCIPHome.CandidateConfigRevoked(priorCandidateDigest);

    s_ccipHome.revokeCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, priorCandidateDigest);

    (CCIPHome.VersionedConfig memory storedVersionedConfig, bool ok) =
      s_ccipHome.getConfig(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, priorCandidateDigest);
    assertFalse(ok);
    // Ensure no old data is returned, even though it's still in storage
    assertEq(storedVersionedConfig.version, 0);
    assertEq(storedVersionedConfig.config.chainSelector, 0);
    assertEq(storedVersionedConfig.config.FRoleDON, 0);

    // Asser the active digest is unaffected but the candidate digest is set to zero
    (bytes32 activeDigest, bytes32 candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, priorActiveDigest);
    assertEq(candidateDigest, ZERO_DIGEST);
    assertTrue(candidateDigest != priorCandidateDigest);
  }

  function test_revokeCandidate_ConfigDigestMismatch_reverts() public {
    (, bytes32 priorCandidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);

    bytes32 wrongDigest = keccak256("wrong_digest");
    vm.expectRevert(abi.encodeWithSelector(CCIPHome.ConfigDigestMismatch.selector, priorCandidateDigest, wrongDigest));
    s_ccipHome.revokeCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, wrongDigest);
  }

  function test_revokeCandidate_RevokingZeroDigestNotAllowed_reverts() public {
    vm.expectRevert(CCIPHome.RevokingZeroDigestNotAllowed.selector);
    s_ccipHome.revokeCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, ZERO_DIGEST);
  }

  function test_revokeCandidate_CanOnlySelfCall_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(CCIPHome.CanOnlySelfCall.selector);
    s_ccipHome.revokeCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, keccak256("configDigest"));
  }
}

contract CCIPHome_promoteCandidateAndRevokeActive is CCIPHomeTest {
  function test_promoteCandidateAndRevokeActive_success() public {
    CCIPHome.OCR3Config memory config = _getBaseConfig();
    bytes32 firstConfigToPromote = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    vm.expectEmit();
    emit CCIPHome.ConfigPromoted(firstConfigToPromote);

    s_ccipHome.promoteCandidateAndRevokeActive(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, firstConfigToPromote, ZERO_DIGEST);

    // Assert the active digest is updated and the candidate digest is set to zero
    (bytes32 activeDigest, bytes32 candidateDigest) = s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeDigest, firstConfigToPromote);
    assertEq(candidateDigest, ZERO_DIGEST);

    // Set a new candidate to promote over a non-zero active config.
    config.offchainConfig = abi.encode("new_offchainConfig_config");
    bytes32 secondConfigToPromote = s_ccipHome.setCandidate(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, config, ZERO_DIGEST);

    vm.expectEmit();
    emit CCIPHome.ActiveConfigRevoked(firstConfigToPromote);

    vm.expectEmit();
    emit CCIPHome.ConfigPromoted(secondConfigToPromote);

    s_ccipHome.promoteCandidateAndRevokeActive(
      DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, secondConfigToPromote, firstConfigToPromote
    );

    (CCIPHome.VersionedConfig memory activeConfig, CCIPHome.VersionedConfig memory candidateConfig) =
      s_ccipHome.getAllConfigs(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    assertEq(activeConfig.configDigest, secondConfigToPromote);
    assertEq(candidateConfig.configDigest, ZERO_DIGEST);
    assertEq(keccak256(abi.encode(activeConfig.config)), keccak256(abi.encode(config)));
  }

  function test_promoteCandidateAndRevokeActive_NoOpStateTransitionNotAllowed_reverts() public {
    vm.expectRevert(CCIPHome.NoOpStateTransitionNotAllowed.selector);
    s_ccipHome.promoteCandidateAndRevokeActive(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, ZERO_DIGEST, ZERO_DIGEST);
  }

  function test_promoteCandidateAndRevokeActive_ConfigDigestMismatch_reverts() public {
    (bytes32 priorActiveDigest, bytes32 priorCandidateDigest) =
      s_ccipHome.getConfigDigests(DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE);
    bytes32 wrongActiveDigest = keccak256("wrongActiveDigest");
    bytes32 wrongCandidateDigest = keccak256("wrongCandidateDigest");

    vm.expectRevert(
      abi.encodeWithSelector(CCIPHome.ConfigDigestMismatch.selector, priorActiveDigest, wrongCandidateDigest)
    );
    s_ccipHome.promoteCandidateAndRevokeActive(
      DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, wrongCandidateDigest, wrongActiveDigest
    );

    vm.expectRevert(
      abi.encodeWithSelector(CCIPHome.ConfigDigestMismatch.selector, priorActiveDigest, wrongActiveDigest)
    );

    s_ccipHome.promoteCandidateAndRevokeActive(
      DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, priorCandidateDigest, wrongActiveDigest
    );
  }

  function test_promoteCandidateAndRevokeActive_CanOnlySelfCall_reverts() public {
    vm.stopPrank();

    vm.expectRevert(CCIPHome.CanOnlySelfCall.selector);
    s_ccipHome.promoteCandidateAndRevokeActive(
      DEFAULT_DON_ID, DEFAULT_PLUGIN_TYPE, keccak256("toPromote"), keccak256("ToRevoke")
    );
  }
}

contract CCIPHome__validateConfig is CCIPHomeTest {
//  function test_validateStaticAndDynamicConfig_OutOfBoundsNodesLength_reverts() public {
//    CCIPHome.OCR3Config memory config = _getBaseConfig();
//    config.staticConfig.nodes = new CCIPHome.Node[](257);
//
//    vm.expectRevert(CCIPHome.OutOfBoundsNodesLength.selector);
//    s_ccipHome.setCandidate(config.staticConfig, config.dynamicConfig, ZERO_DIGEST);
//  }
}
