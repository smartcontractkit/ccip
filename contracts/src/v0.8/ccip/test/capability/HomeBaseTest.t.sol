// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {HomeBase} from "../../capability/HomeBase.sol";
import {Internal} from "../../libraries/Internal.sol";
import {HomeBaseHelper} from "../helpers/HomeBaseHelper.sol";
import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

contract HomeBaseTest is Test {
  uint32 internal constant DON_ID = 593;
  uint8 internal constant PLUGIN_TYPE = 244;

  bytes32 internal constant ZERO_DIGEST = bytes32(uint256(0));

  HomeBaseHelper internal s_homeBase;
  address private constant CAPABILITIES_REGISTRY = address(1);

  function setUp() public virtual {
    s_homeBase = new HomeBaseHelper(CAPABILITIES_REGISTRY);
  }

  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 public constant PREFIX = 0x0c0c << (256 - 16);

  function _getConfigDigest(bytes memory staticConfig, uint32 version) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(
            keccak256(
              bytes.concat(
                abi.encode(bytes32("EVM"), block.chainid, address(s_homeBase), DON_ID, PLUGIN_TYPE, version), staticConfig
              )
            )
          ) & ~PREFIX_MASK
        )
    );
  }

  function _getStaticConfig() internal pure returns (bytes memory) {
    return abi.encode("staticConfig");
  }

  function _getDynamicConfig() internal pure returns (bytes memory) {
    return abi.encode("dynamicConfig");
  }
}

contract RMNHome_setSecondary is HomeBaseTest {
  function test_setSecondary_success() public {
    HomeBase.StoredConfig memory encodedConfig = HomeBase.StoredConfig({
      configDigest: ZERO_DIGEST,
      version: 1,
      staticConfig: _getStaticConfig(),
      dynamicConfig: _getDynamicConfig()
    });

    encodedConfig.configDigest = _getConfigDigest(encodedConfig.staticConfig, encodedConfig.version);

    vm.expectEmit();
    emit HomeBase.ConfigSet(encodedConfig);

    s_homeBase.setSecondary(DON_ID, PLUGIN_TYPE, encodedConfig.staticConfig, encodedConfig.dynamicConfig, ZERO_DIGEST);

    (HomeBase.StoredConfig memory storedConfig, bool ok) = s_homeBase.getSecondaryStoredConfig(DON_ID, PLUGIN_TYPE);
    assertTrue(ok);
    assertEq(storedConfig.version, encodedConfig.version);
    assertEq(storedConfig.configDigest, encodedConfig.configDigest);
    assertEq(storedConfig.staticConfig, encodedConfig.staticConfig);
    assertEq(storedConfig.dynamicConfig, encodedConfig.dynamicConfig);
  }

  function test_setSecondary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(HomeBase.OnlyOwnerOrSelfCallAllowed.selector);
    s_homeBase.setSecondary(DON_ID, PLUGIN_TYPE, _getStaticConfig(), _getDynamicConfig(), ZERO_DIGEST);
  }
}

contract RMNHome_setDynamicConfig is HomeBaseTest {
  function setUp() public override {
    super.setUp();
    s_homeBase.setSecondary(DON_ID, PLUGIN_TYPE, _getStaticConfig(), _getDynamicConfig(), ZERO_DIGEST);
  }

  function test_setDynamicConfig_success() public {
    (bytes32 priorPrimaryDigest, bytes32 secondaryConfigDigest) = s_homeBase.getConfigDigests(DON_ID, PLUGIN_TYPE);

    bytes memory newDynamicConfig = abi.encode("newDynamicConfig");

    vm.expectEmit();
    emit HomeBase.DynamicConfigSet(secondaryConfigDigest, newDynamicConfig);

    s_homeBase.setDynamicConfig(DON_ID, PLUGIN_TYPE, newDynamicConfig, secondaryConfigDigest);

    (HomeBase.StoredConfig memory storedConfig, bool ok) =
      s_homeBase.getStoredConfig(DON_ID, PLUGIN_TYPE, secondaryConfigDigest);
    assertTrue(ok);
    assertEq(storedConfig.dynamicConfig, newDynamicConfig);

    // Asser the digests don't change when updating the dynamic config
    (bytes32 primaryDigest, bytes32 secondaryDigest) = s_homeBase.getConfigDigests(DON_ID, PLUGIN_TYPE);
    assertEq(primaryDigest, priorPrimaryDigest);
    assertEq(secondaryDigest, secondaryConfigDigest);
  }

  function test_setDynamicConfig_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(HomeBase.OnlyOwnerOrSelfCallAllowed.selector);
    s_homeBase.setDynamicConfig(DON_ID, PLUGIN_TYPE, _getDynamicConfig(), keccak256("configDigest"));
  }
}

contract RMNHome_revokeSecondary is HomeBaseTest {
  // Sets two configs
  function setUp() public override {
    super.setUp();
    bytes32 digest = s_homeBase.setSecondary(DON_ID, PLUGIN_TYPE, _getStaticConfig(), _getDynamicConfig(), ZERO_DIGEST);
    s_homeBase.promoteSecondaryAndRevokePrimary(DON_ID, PLUGIN_TYPE, digest, ZERO_DIGEST);
    s_homeBase.setSecondary(DON_ID, PLUGIN_TYPE, _getStaticConfig(), _getDynamicConfig(), ZERO_DIGEST);
  }

  function test_revokeSecondary_success() public {
    (bytes32 priorPrimaryDigest, bytes32 priorSecondaryDigest) = s_homeBase.getConfigDigests(DON_ID, PLUGIN_TYPE);

    vm.expectEmit();
    emit HomeBase.ConfigRevoked(priorSecondaryDigest);

    s_homeBase.revokeSecondary(DON_ID, PLUGIN_TYPE, priorSecondaryDigest);

    (HomeBase.StoredConfig memory storedVersionedConfig, bool ok) =
      s_homeBase.getStoredConfig(DON_ID, PLUGIN_TYPE, priorSecondaryDigest);
    assertFalse(ok);
    // Ensure no old data is returned, even though it's still in storage
    assertEq(storedVersionedConfig.version, 0);
    assertEq(storedVersionedConfig.staticConfig.length, 0);
    assertEq(storedVersionedConfig.dynamicConfig.length, 0);

    // Asser the primary digest is unaffected but the secondary digest is set to zero
    (bytes32 primaryDigest, bytes32 secondaryDigest) = s_homeBase.getConfigDigests(DON_ID, PLUGIN_TYPE);
    assertEq(primaryDigest, priorPrimaryDigest);
    assertEq(secondaryDigest, ZERO_DIGEST);
    assertTrue(secondaryDigest != priorSecondaryDigest);
  }

  function test_revokeSecondary_ConfigDigestMismatch_reverts() public {
    (, bytes32 priorSecondaryDigest) = s_homeBase.getConfigDigests(DON_ID, PLUGIN_TYPE);

    bytes32 wrongDigest = keccak256("wrong_digest");
    vm.expectRevert(abi.encodeWithSelector(HomeBase.ConfigDigestMismatch.selector, priorSecondaryDigest, wrongDigest));
    s_homeBase.revokeSecondary(DON_ID, PLUGIN_TYPE, wrongDigest);
  }

  function test_revokeSecondary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(HomeBase.OnlyOwnerOrSelfCallAllowed.selector);
    s_homeBase.revokeSecondary(DON_ID, PLUGIN_TYPE, keccak256("configDigest"));
  }
}

contract RMNHome_promoteSecondaryAndRevokePrimary is HomeBaseTest {
  function test_promoteSecondaryAndRevokePrimary_success() public {}

  function test_promoteSecondaryAndRevokePrimary_OnlyOwner_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(HomeBase.OnlyOwnerOrSelfCallAllowed.selector);
    s_homeBase.promoteSecondaryAndRevokePrimary(DON_ID, PLUGIN_TYPE, keccak256("toPromote"), keccak256("ToRevoke"));
  }
}

contract RMNHome_beforeCapabilityConfigSet is HomeBaseTest {
  function test_beforeCapabilityConfigSet_success() public {
    vm.startPrank(address(1));

    HomeBase.StoredConfig memory encodedConfig = HomeBase.StoredConfig({
      configDigest: ZERO_DIGEST,
      version: 1,
      staticConfig: _getStaticConfig(),
      dynamicConfig: _getDynamicConfig()
    });
    encodedConfig.configDigest = _getConfigDigest(encodedConfig.staticConfig, encodedConfig.version);

    bytes memory callPayload = abi.encodeCall(
      HomeBase.setSecondary, (DON_ID, PLUGIN_TYPE, encodedConfig.staticConfig, encodedConfig.dynamicConfig, ZERO_DIGEST)
    );

    vm.expectEmit();
    emit HomeBase.ConfigSet(encodedConfig);

    s_homeBase.beforeCapabilityConfigSet(new bytes32[](0), callPayload, 0, DON_ID);
  }

  function test_beforeCapabilityConfigSet_OnlyCapabilitiesRegistryCanCall_reverts() public {
    vm.startPrank(address(0));

    vm.expectRevert(HomeBase.OnlyCapabilitiesRegistryCanCall.selector);
    s_homeBase.beforeCapabilityConfigSet(new bytes32[](0), new bytes(0), 0, DON_ID);
  }
}
