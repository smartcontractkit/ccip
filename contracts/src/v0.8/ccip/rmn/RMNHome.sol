// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

/// @notice Stores the home configuration for RMN, that is referenced by CCIP oracles, RMN nodes, and the RMNRemote
/// contracts.
contract RMNHome is OwnerIsCreator, ITypeAndVersion {
  error OutOfBoundsNodesLength();
  error DuplicatePeerId();
  error DuplicateOffchainPublicKey();
  error DuplicateSourceChain();
  error OutOfBoundsObserverNodeIndex();
  error MinObserversTooHigh();
  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);
  error DigestNotFound(bytes32 configDigest);

  event ConfigSet(bytes32 configDigest, VersionedConfig versionedConfig);
  event ConfigRevoked(bytes32 configDigest);
  event DynamicConfigSet(bytes32 indexed configDigest, DynamicConfig dynamicConfig);
  event ConfigPromoted(bytes32 configDigest);

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

  struct Config {
    StaticConfig staticConfig;
    DynamicConfig dynamicConfig;
  }

  struct VersionedConfig {
    uint32 version; // The version of this config, starting from 1 it increments with each new config.
    Config config;
  }

  struct VersionedConfigWithDigest {
    bytes32 configDigest;
    VersionedConfig versionedConfig;
  }

  string public constant override typeAndVersion = "RMNHome 1.6.0-dev";

  /// @notice The max number of configs that can be active at the same time.
  uint256 private constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 private constant PREFIX = 0x000b << (256 - 16); // 0x000b00..00
  bytes32 private constant ZERO_DIGEST = bytes32(uint256(0));

  /// @notice This array holds the digests of the configs, used for efficiency.
  /// @dev Value i in this array is valid iff it's not 0.
  bytes32[MAX_CONCURRENT_CONFIGS] private s_configDigests;
  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configDigests[i] != 0.
  Config[MAX_CONCURRENT_CONFIGS] private s_configs;
  /// @notice This array holds the versions of the configs.
  /// @dev Value i in this array is valid iff s_configDigests[i] != 0.
  /// @dev Since Solidity doesn't support writing complex memory structs to storage, we have to make the config calldata
  /// in setConfig and then copy it to storage. This does not allow us to modify it to add the version field, so we
  /// store the version separately.
  uint32[MAX_CONCURRENT_CONFIGS] private s_configVersions;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 private s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 private s_primaryConfigIndex = 0;

  /// @notice Returns the current primary and secondary config digests.
  /// @dev Can be bytes32(0) if no config has been set yet or it has been revoked.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests() external view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (s_configDigests[s_primaryConfigIndex], s_configDigests[s_primaryConfigIndex ^ 1]);
  }

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
        return (VersionedConfig({config: s_configs[i], version: s_configVersions[i]}), true);
      }
    }
    return (versionedConfig, false);
  }

  function getAllConfigs()
    external
    view
    returns (VersionedConfigWithDigest memory primaryConfig, VersionedConfigWithDigest memory secondaryConfig)
  {
    // We need to explicitly check if the digest exists, because we don't clear out revoked config state. Not doing this
    // check would result in potentially returning previous configs.
    uint256 primaryConfigIndex = s_primaryConfigIndex;
    bytes32 primaryConfigDigest = s_configDigests[primaryConfigIndex];
    if (primaryConfigDigest != ZERO_DIGEST) {
      primaryConfig = VersionedConfigWithDigest({
        configDigest: primaryConfigDigest,
        versionedConfig: (
          VersionedConfig({config: s_configs[primaryConfigIndex], version: s_configVersions[primaryConfigIndex]})
        )
      });
    }

    uint256 secondaryConfigIndex = primaryConfigIndex ^ 1;
    bytes32 secondaryConfigDigest = s_configDigests[secondaryConfigIndex];
    if (secondaryConfigDigest != ZERO_DIGEST) {
      secondaryConfig = VersionedConfigWithDigest({
        configDigest: secondaryConfigDigest,
        versionedConfig: (
          VersionedConfig({config: s_configs[secondaryConfigIndex], version: s_configVersions[secondaryConfigIndex]})
        )
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
    s_configs[secondaryConfigIndex] = newConfig;
    s_configVersions[secondaryConfigIndex] = newVersion;
    s_configDigests[secondaryConfigIndex] = newConfigDigest;

    emit ConfigSet(newConfigDigest, VersionedConfig({version: newVersion, config: newConfig}));

    return newConfigDigest;
  }

  function setDynamicConfig(DynamicConfig calldata newDynamicConfig, bytes32 currentDigest) external onlyOwner {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      if (s_configDigests[i] == currentDigest && currentDigest != ZERO_DIGEST) {
        Config memory currentConfig = s_configs[i];
        _validateDynamicConfig(newDynamicConfig, currentConfig.staticConfig.nodes.length);

        // Since the dynamic config doesn't change we don't have to update the digest or version.
        s_configs[i].dynamicConfig = newDynamicConfig;

        emit DynamicConfigSet(currentDigest, newDynamicConfig);
        return;
      }
    }

    revert DigestNotFound(currentDigest);
  }

  /// @notice Revokes a specific config by digest.
  /// @param configDigest The digest of the config to revoke. This is done to prevent accidental revokes.
  function revokeSecondary(bytes32 configDigest) external onlyOwner {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configDigests[secondaryConfigIndex] != configDigest) {
      revert ConfigDigestMismatch(s_configDigests[secondaryConfigIndex], configDigest);
    }

    emit ConfigRevoked(configDigest);
    // Delete only the digest, as that's what's used to determine if a config is active. This means the actual
    // config stays in storage which should significantly reduce the gas cost of overwriting that storage space in
    // the future.
    delete s_configDigests[secondaryConfigIndex];
  }

  /// @notice Promotes the secondary config to the primary config. This demotes the primary to be the secondary but does
  /// not revoke it. To revoke the primary, use `promoteSecondaryAndRevokePrimary` instead.
  function promoteSecondary(bytes32 digestToPromote) external onlyOwner {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configDigests[secondaryConfigIndex] != digestToPromote) {
      revert ConfigDigestMismatch(s_configDigests[secondaryConfigIndex], digestToPromote);
    }

    s_primaryConfigIndex ^= 1;

    emit ConfigPromoted(digestToPromote);
  }

  /// @notice Promotes the secondary config to the primary config and revokes the primary config.
  function promoteSecondaryAndRevokePrimary(bytes32 digestToPromote, bytes32 digestToRevoke) external onlyOwner {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configDigests[secondaryConfigIndex] != digestToPromote) {
      revert ConfigDigestMismatch(s_configDigests[secondaryConfigIndex], digestToPromote);
    }

    uint256 primaryConfigIndex = s_primaryConfigIndex;
    if (s_configDigests[primaryConfigIndex] != digestToRevoke) {
      revert ConfigDigestMismatch(s_configDigests[primaryConfigIndex], digestToRevoke);
    }

    delete s_configDigests[primaryConfigIndex];

    s_primaryConfigIndex ^= 1;
    emit ConfigRevoked(digestToRevoke);
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
