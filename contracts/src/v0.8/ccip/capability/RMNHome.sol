// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

/// @notice Stores the home configuration for RMN, that is referenced by CCIP oracles, RMN nodes, and the RMNRemote
/// contracts.
contract RMNHome is OwnerIsCreator, ITypeAndVersion {
  event ConfigSet(bytes32 indexed configDigest, uint32 version, StaticConfig staticConfig, DynamicConfig dynamicConfig);
  event ConfigRevoked(bytes32 indexed configDigest);
  event DynamicConfigSet(bytes32 indexed configDigest, DynamicConfig dynamicConfig);
  event ConfigPromoted(bytes32 indexed configDigest);

  error OutOfBoundsNodesLength();
  error DuplicatePeerId();
  error DuplicateOffchainPublicKey();
  error DuplicateSourceChain();
  error OutOfBoundsObserverNodeIndex();
  error MinObserversTooHigh();
  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);
  error DigestNotFound(bytes32 configDigest);

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
  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  /// @notice The max number of configs that can be active at the same time.
  uint256 private constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Helper to identify the zero config digest with less casting.
  bytes32 private constant ZERO_DIGEST = bytes32(uint256(0));

  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configs[i].configDigest != 0.
  VersionedConfig[MAX_CONCURRENT_CONFIGS] private s_configs;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 private s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 private s_primaryConfigIndex = 0;

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  /// @param configDigest The digest of the config to fetch.
  /// @return versionedConfig The config and its version.
  /// @return ok True if the config was found, false otherwise.
  function getConfig(bytes32 configDigest) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    return _getStoredConfig(configDigest);
  }

  function getAllConfigs()
    external
    view
    returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig)
  {
    VersionedConfig memory storedPrimaryConfig = s_configs[s_primaryConfigIndex];
    if (storedPrimaryConfig.configDigest != ZERO_DIGEST) {
      primaryConfig = storedPrimaryConfig;
    }

    VersionedConfig memory storedSecondaryConfig = s_configs[s_primaryConfigIndex ^ 1];
    if (storedSecondaryConfig.configDigest != ZERO_DIGEST) {
      secondaryConfig = storedSecondaryConfig;
    }

    return (primaryConfig, secondaryConfig);
  }

  function _validateStaticAndDynamicConfig(
    StaticConfig memory staticConfig,
    DynamicConfig memory dynamicConfig
  ) internal pure {
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

    _validateDynamicConfigParsed(dynamicConfig, staticConfig.nodes.length);
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

  function _validateDynamicConfig(StaticConfig memory staticConfig, DynamicConfig memory dynamicConfig) internal pure {
    _validateDynamicConfigParsed(dynamicConfig, staticConfig.nodes.length);
  }

  function _validateCaller() internal view {
    _validateOwnership();
  }

  // ================================================================
  // │                          Getters                             │
  // ================================================================

  /// @notice Returns the current primary and secondary config digests.
  /// @dev Can be bytes32(0) if no config has been set yet or it has been revoked.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests() external view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (s_configs[s_primaryConfigIndex].configDigest, s_configs[s_primaryConfigIndex ^ 1].configDigest);
  }

  /// @notice Returns the primary config digest for for a given key.
  function getPrimaryDigest() public view returns (bytes32) {
    return s_configs[s_primaryConfigIndex].configDigest;
  }

  /// @notice Returns the secondary config digest for for a given key.
  function getSecondaryDigest() public view returns (bytes32) {
    return s_configs[s_primaryConfigIndex ^ 1].configDigest;
  }

  /// @notice Returns the stored config for a given digest. Will always return an empty config if the digest is the zero
  /// digest. This is done to prevent exposing old config state that is invalid.
  /// @param configDigest The digest of the config to fetch.
  function _getStoredConfig(bytes32 configDigest) internal view returns (VersionedConfig memory storedConfig, bool ok) {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      // We never want to return true for a zero digest, even if the caller is asking for it, as this can expose old
      // config state that is invalid.
      if (s_configs[i].configDigest == configDigest && configDigest != ZERO_DIGEST) {
        return (s_configs[i], true);
      }
    }
    return (storedConfig, false);
  }

  /// @notice Returns the primary stored config for a given key.
  /// @return primaryConfig The primary stored config.
  /// @return ok True if the config was found, false otherwise.
  function _getPrimaryStoredConfig() internal view returns (VersionedConfig memory primaryConfig, bool ok) {
    if (s_configs[s_primaryConfigIndex].configDigest == ZERO_DIGEST) {
      return (primaryConfig, false);
    }

    return (s_configs[s_primaryConfigIndex], true);
  }

  /// @notice Returns the secondary stored config for a given key.
  /// @return secondaryConfig The secondary stored config.
  /// @return ok True if the config was found, false otherwise.
  function _getSecondaryStoredConfig() internal view returns (VersionedConfig memory secondaryConfig, bool ok) {
    if (s_configs[s_primaryConfigIndex ^ 1].configDigest == ZERO_DIGEST) {
      return (secondaryConfig, false);
    }

    return (s_configs[s_primaryConfigIndex ^ 1], true);
  }

  // ================================================================
  // │                     State transitions                        │
  // ================================================================

  /// @notice Sets a new config as the secondary config. Does not influence the primary config.
  /// @param staticConfig The static part of the config.
  /// @param dynamicConfig The dynamic part of the config.
  /// @param digestToOverwrite The digest of the config to overwrite, or ZERO_DIGEST if no config is to be overwritten.
  /// This is done to prevent accidental overwrites.
  /// @return newConfigDigest The digest of the new config.
  function setSecondary(
    StaticConfig calldata staticConfig,
    DynamicConfig calldata dynamicConfig,
    bytes32 digestToOverwrite
  ) external returns (bytes32 newConfigDigest) {
    _validateCaller();
    _validateStaticAndDynamicConfig(staticConfig, dynamicConfig);

    bytes32 existingDigest = getSecondaryDigest();

    if (existingDigest != digestToOverwrite) {
      revert ConfigDigestMismatch(existingDigest, digestToOverwrite);
    }

    // are we going to overwrite a config? If so, emit an event.
    if (existingDigest != ZERO_DIGEST) {
      emit ConfigRevoked(digestToOverwrite);
    }

    uint32 newVersion = ++s_configCount;
    newConfigDigest = _calculateConfigDigest(abi.encode(staticConfig), newVersion);

    VersionedConfig storage existingConfig = s_configs[s_primaryConfigIndex ^ 1];
    existingConfig.configDigest = newConfigDigest;
    existingConfig.version = newVersion;
    existingConfig.staticConfig = staticConfig;
    existingConfig.dynamicConfig = dynamicConfig;

    emit ConfigSet(newConfigDigest, newVersion, staticConfig, dynamicConfig);

    return newConfigDigest;
  }

  /// @notice Revokes a specific config by digest.
  /// @param configDigest The digest of the config to revoke. This is done to prevent accidental revokes.
  function revokeSecondary(bytes32 configDigest) external {
    _validateCaller();

    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configs[secondaryConfigIndex].configDigest != configDigest) {
      revert ConfigDigestMismatch(s_configs[secondaryConfigIndex].configDigest, configDigest);
    }

    emit ConfigRevoked(configDigest);
    // Delete only the digest, as that's what's used to determine if a config is active. This means the actual
    // config stays in storage which should significantly reduce the gas cost of overwriting that storage space in
    // the future.
    delete s_configs[secondaryConfigIndex].configDigest;
  }

  /// @notice Promotes the secondary config to the primary config and revokes the primary config.
  /// @param digestToPromote The digest of the config to promote.
  /// @param digestToRevoke The digest of the config to revoke.
  function promoteSecondaryAndRevokePrimary(bytes32 digestToPromote, bytes32 digestToRevoke) external {
    _validateCaller();

    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configs[secondaryConfigIndex].configDigest != digestToPromote) {
      revert ConfigDigestMismatch(s_configs[secondaryConfigIndex].configDigest, digestToPromote);
    }

    uint256 primaryConfigIndex = s_primaryConfigIndex;
    if (s_configs[primaryConfigIndex].configDigest != digestToRevoke) {
      revert ConfigDigestMismatch(s_configs[primaryConfigIndex].configDigest, digestToRevoke);
    }

    delete s_configs[primaryConfigIndex].configDigest;

    s_primaryConfigIndex ^= 1;
    if (digestToRevoke != ZERO_DIGEST) {
      emit ConfigRevoked(digestToRevoke);
    }
    emit ConfigPromoted(digestToPromote);
  }

  /// @notice Sets the dynamic config for a specific config.
  /// @param newDynamicConfig The new dynamic config.
  /// @param currentDigest The digest of the config to update.
  /// @dev This does not update the config digest as only the static config is part of the digest.
  function setDynamicConfig(DynamicConfig calldata newDynamicConfig, bytes32 currentDigest) external {
    _validateCaller();

    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      if (s_configs[i].configDigest == currentDigest && currentDigest != ZERO_DIGEST) {
        _validateDynamicConfig(s_configs[i].staticConfig, newDynamicConfig);

        // Since the static config doesn't change we don't have to update the digest or version.
        s_configs[i].dynamicConfig = newDynamicConfig;

        emit DynamicConfigSet(currentDigest, newDynamicConfig);
        return;
      }
    }

    revert DigestNotFound(currentDigest);
  }

  /// @notice Calculates the config digest for a given plugin key, static config, and version.
  /// @param staticConfig The static part of the config.
  /// @param version The version of the config.
  /// @return The calculated config digest.
  function _calculateConfigDigest(bytes memory staticConfig, uint32 version) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(
            keccak256(bytes.concat(abi.encode(bytes32("EVM"), block.chainid, address(this), version), staticConfig))
          ) & ~PREFIX_MASK
        )
    );
  }
}
