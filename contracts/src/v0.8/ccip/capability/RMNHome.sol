// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {HomeBase} from "./HomeBase.sol";

/// @notice Stores the home configuration for RMN, that is referenced by CCIP oracles, RMN nodes, and the RMNRemote
/// contracts.
contract RMNHome is HomeBase {
  error OutOfBoundsNodesLength();
  error DuplicatePeerId();
  error DuplicateOffchainPublicKey();
  error DuplicateSourceChain();
  error OutOfBoundsObserverNodeIndex();
  error MinObserversTooHigh();

  struct Node {
    bytes32 peerId; //            Used for p2p communication.
    bytes32 offchainPublicKey; // Observations are signed with this public key, and are only verified offchain.
  }

  struct SourceChain {
    uint64 chainSelector; // ─────╮ The Source chain selector.
    uint64 minObservers; // ──────╯ Required number of observers to agree on an observation for this source chain.
    uint256 observerNodesBitmap; // ObserverNodesBitmap & (1<<i) == (1<<i) iff nodes[i] is an observer for this source chain.
  }

  struct StaticConfig {
    // No sorting requirement for nodes, but ensure that SourceChain.observerNodeIndices in the home chain config &
    // Signer.nodeIndex in the remote chain configs are appropriately updated when changing this field.
    Node[] nodes;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct DynamicConfig {
    // No sorting requirement for source chains, it is most gas efficient to append new source chains to the right.
    SourceChain[] sourceChains;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct VersionedConfig {
    uint32 version;
    bytes32 configDigest;
    StaticConfig staticConfig;
    DynamicConfig dynamicConfig;
  }

  string public constant override typeAndVersion = "RMNHome 1.6.0-dev";

  uint256 private constant PREFIX = 0x000b << (256 - 16); // 0x000b00..00

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  /// @param configDigest The digest of the config to fetch.
  /// @return versionedConfig The config and its version.
  /// @return ok True if the config was found, false otherwise.
  function getConfig(bytes32 configDigest) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    (StoredConfig memory storedConfig, bool configOK) = _getStoredConfig(configDigest);
    if (configOK) {
      return (
        VersionedConfig({
          version: storedConfig.version,
          configDigest: storedConfig.configDigest,
          staticConfig: abi.decode(storedConfig.staticConfig, (StaticConfig)),
          dynamicConfig: abi.decode(storedConfig.dynamicConfig, (DynamicConfig))
        }),
        true
      );
    }

    return (versionedConfig, false);
  }

  function getAllConfigs()
    external
    view
    returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig)
  {
    (StoredConfig memory primaryStoredConfig, bool primaryOk) = _getPrimaryStoredConfig();

    if (primaryOk) {
      primaryConfig = VersionedConfig({
        version: primaryStoredConfig.version,
        configDigest: primaryStoredConfig.configDigest,
        staticConfig: abi.decode(primaryStoredConfig.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(primaryStoredConfig.dynamicConfig, (DynamicConfig))
      });
    }

    (StoredConfig memory secondaryStoredConfig, bool secondaryOk) = _getSecondaryStoredConfig();

    if (secondaryOk) {
      secondaryConfig = VersionedConfig({
        version: secondaryStoredConfig.version,
        configDigest: secondaryStoredConfig.configDigest,
        staticConfig: abi.decode(secondaryStoredConfig.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(secondaryStoredConfig.dynamicConfig, (DynamicConfig))
      });
    }

    return (primaryConfig, secondaryConfig);
  }

  function _validateStaticAndDynamicConfig(
    bytes memory encodedStaticConfig,
    bytes memory encodedDynamicConfig
  ) internal pure override {
    StaticConfig memory staticConfig = abi.decode(encodedStaticConfig, (StaticConfig));
    // Ensure that observerNodesBitmap can be bit-encoded into a uint256.
    if (staticConfig.nodes.length > 256) {
      revert OutOfBoundsNodesLength();
    }

    // Ensure no peerId or offchainPublicKey is duplicated.
    for (uint256 i = 0; i < staticConfig.nodes.length; ++i) {
      for (uint256 j = i + 1; j < staticConfig.nodes.length; ++j) {
        if (staticConfig.nodes[i].peerId == staticConfig.nodes[j].peerId) {
          revert DuplicatePeerId();
        }
        if (staticConfig.nodes[i].offchainPublicKey == staticConfig.nodes[j].offchainPublicKey) {
          revert DuplicateOffchainPublicKey();
        }
      }
    }

    _validateDynamicConfigParsed(abi.decode(encodedDynamicConfig, (DynamicConfig)), staticConfig.nodes.length);
  }

  function _validateDynamicConfigParsed(DynamicConfig memory dynamicConfig, uint256 numberOfNodes) internal pure {
    uint256 numberOfSourceChains = dynamicConfig.sourceChains.length;
    for (uint256 i = 0; i < numberOfSourceChains; ++i) {
      SourceChain memory currentSourceChain = dynamicConfig.sourceChains[i];
      // Ensure the source chain is unique.
      for (uint256 j = i + 1; j < numberOfSourceChains; ++j) {
        if (currentSourceChain.chainSelector == dynamicConfig.sourceChains[j].chainSelector) {
          revert DuplicateSourceChain();
        }
      }

      // all observer node indices are valid
      uint256 bitmap = currentSourceChain.observerNodesBitmap;
      // Check if there are any bits set for indexes outside of the expected range.
      if (bitmap & (type(uint256).max >> (256 - numberOfNodes)) != bitmap) {
        revert OutOfBoundsObserverNodeIndex();
      }

      uint256 observersCount = 0;
      for (; bitmap != 0; ++observersCount) {
        bitmap &= bitmap - 1;
      }

      // minObservers are tenable
      if (currentSourceChain.minObservers > observersCount) {
        revert MinObserversTooHigh();
      }
    }
  }

  function _validateDynamicConfig(
    bytes memory encodedStaticConfig,
    bytes memory encodedDynamicConfig
  ) internal pure override {
    uint256 numberOfNodes = abi.decode(encodedStaticConfig, (StaticConfig)).nodes.length;
    DynamicConfig memory dynamicConfig = abi.decode(encodedDynamicConfig, (DynamicConfig));

    _validateDynamicConfigParsed(dynamicConfig, numberOfNodes);
  }

  function _getConfigDigestPrefix() internal pure override returns (uint256) {
    return PREFIX;
  }
}
