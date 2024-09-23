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
  error DigestNotFound(bytes32 configDigest);

  event ConfigSet(VersionedConfig versionedConfig);
  event DynamicConfigSet(bytes32 indexed configDigest, DynamicConfig dynamicConfig);
  event ConfigPromoted(bytes32 configDigest);

  struct Node {
    bytes32 peerId; //            Used for p2p communication.
    bytes32 offchainPublicKey; // Observations are signed with this public key, and are only verified offchain.
  }

  struct StaticConfig {
    // No sorting requirement for nodes, but ensure that SourceChain.observerNodeIndices in the home chain config &
    // Signer.nodeIndex in the remote chain configs are appropriately updated when changing this field.
    Node[] nodes;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct SourceChain {
    uint64 chainSelector; // ─────╮ The Source chain selector.
    uint64 minObservers; // ──────╯ Required number of observers to agree on an observation for this source chain.
    uint256 observerNodesBitmap; // ObserverNodesBitmap & (1<<i) == (1<<i) iff nodes[i] is an observer for this source chain.
  }

  struct DynamicConfig {
    // No sorting requirement for source chains, it is most gas efficient to append new source chains to the right.
    SourceChain[] sourceChains;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct Config {
    StaticConfig staticConfig;
    DynamicConfig dynamicConfig;
  }

  struct VersionedConfig {
    uint32 version; // The version of this config, starting from 1 it increments with each new config.
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
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      // We never want to return true for a zero digest, even if the caller is asking for it, as this can expose old
      // config state that is invalid.
      if (s_configDigests[i] == configDigest && configDigest != ZERO_DIGEST) {
        StoredConfig memory config = s_configs[i];
        return (
          VersionedConfig({
            version: s_configVersions[i],
            configDigest: configDigest,
            staticConfig: abi.decode(config.staticConfig, (StaticConfig)),
            dynamicConfig: abi.decode(config.dynamicConfig, (DynamicConfig))
          }),
          true
        );
      }
    }
    return (versionedConfig, false);
  }

  function getAllConfigs()
    external
    view
    returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig)
  {
    // We need to explicitly check if the digest exists, because we don't clear out revoked config state. Not doing this
    // check would result in potentially returning previous configs.
    uint256 primaryConfigIndex = s_primaryConfigIndex;
    bytes32 primaryConfigDigest = s_configDigests[primaryConfigIndex];
    if (primaryConfigDigest != ZERO_DIGEST) {
      StoredConfig memory config = s_configs[primaryConfigIndex];

      primaryConfig = VersionedConfig({
        version: s_configVersions[primaryConfigIndex],
        configDigest: primaryConfigDigest,
        staticConfig: abi.decode(config.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(config.dynamicConfig, (DynamicConfig))
      });
    }

    uint256 secondaryConfigIndex = primaryConfigIndex ^ 1;
    bytes32 secondaryConfigDigest = s_configDigests[secondaryConfigIndex];
    if (secondaryConfigDigest != ZERO_DIGEST) {
      StoredConfig memory config = s_configs[secondaryConfigIndex];

      secondaryConfig = VersionedConfig({
        version: s_configVersions[secondaryConfigIndex],
        configDigest: secondaryConfigDigest,
        staticConfig: abi.decode(config.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(config.dynamicConfig, (DynamicConfig))
      });
    }

    return (primaryConfig, secondaryConfig);
  }

  /// @notice Sets a new config as the secondary config. Does not influence the primary config.
  /// @param newConfig The new config to set.
  /// @param digestToOverwrite The digest of the config to overwrite, or ZERO_DIGEST if no config is to be overwritten.
  /// This is done to prevent accidental overwrites.
  function setSecondary(
    Config calldata newConfig,
    bytes32 digestToOverwrite
  ) external onlyOwner returns (bytes32 newConfigDigest) {
    _validateStaticConfig(newConfig.staticConfig);
    _validateDynamicConfig(newConfig.dynamicConfig, newConfig.staticConfig.nodes.length);

    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;

    if (s_configDigests[secondaryConfigIndex] != digestToOverwrite) {
      revert ConfigDigestMismatch(s_configDigests[secondaryConfigIndex], digestToOverwrite);
    }

    // are we going to overwrite a config? If so, emit an event.
    if (digestToOverwrite != ZERO_DIGEST) {
      emit ConfigRevoked(digestToOverwrite);
    }

    uint32 newVersion = ++s_configCount;
    newConfigDigest = _getConfigDigest(newConfig.staticConfig, newVersion);
    s_configs[secondaryConfigIndex] =
      StoredConfig(abi.encode(newConfig.staticConfig), abi.encode(newConfig.dynamicConfig));
    s_configVersions[secondaryConfigIndex] = newVersion;
    s_configDigests[secondaryConfigIndex] = newConfigDigest;

    emit ConfigSet(
      VersionedConfig({
        version: newVersion,
        configDigest: newConfigDigest,
        staticConfig: newConfig.staticConfig,
        dynamicConfig: newConfig.dynamicConfig
      })
    );

    return newConfigDigest;
  }

  function setDynamicConfig(DynamicConfig calldata newDynamicConfig, bytes32 currentDigest) external onlyOwner {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      if (s_configDigests[i] == currentDigest && currentDigest != ZERO_DIGEST) {
        StaticConfig memory staticConfig = abi.decode(s_configs[i].staticConfig, (StaticConfig));
        _validateDynamicConfig(newDynamicConfig, staticConfig.nodes.length);

        // Since the static config doesn't change we don't have to update the digest or version.
        s_configs[i].dynamicConfig = abi.encode(newDynamicConfig);

        emit DynamicConfigSet(currentDigest, newDynamicConfig);
        return;
      }
    }

    revert DigestNotFound(currentDigest);
  }

  function _validateStaticConfig(StaticConfig calldata newStaticConfig) internal pure {
    // Ensure that observerNodesBitmap can be bit-encoded into a uint256.
    if (newStaticConfig.nodes.length > 256) {
      revert OutOfBoundsNodesLength();
    }

    // Ensure no peerId or offchainPublicKey is duplicated.
    for (uint256 i = 0; i < newStaticConfig.nodes.length; ++i) {
      for (uint256 j = i + 1; j < newStaticConfig.nodes.length; ++j) {
        if (newStaticConfig.nodes[i].peerId == newStaticConfig.nodes[j].peerId) {
          revert DuplicatePeerId();
        }
        if (newStaticConfig.nodes[i].offchainPublicKey == newStaticConfig.nodes[j].offchainPublicKey) {
          revert DuplicateOffchainPublicKey();
        }
      }
    }
  }

  function _validateDynamicConfig(DynamicConfig calldata dynamicConfig, uint256 numberOfNodes) internal pure {
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

  function _getConfigDigest(StaticConfig memory staticConfig, uint32 version) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(keccak256(abi.encode(bytes32("EVM"), block.chainid, address(this), version, staticConfig)))
            & ~PREFIX_MASK
        )
    );
  }
}
