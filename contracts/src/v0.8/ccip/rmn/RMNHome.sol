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

  event ConfigSet(bytes32 configDigest, VersionedConfig versionedConfig);
  event ConfigRevoked(bytes32 configDigest);

  struct Node {
    bytes32 peerId; //            Used for p2p communication.
    bytes32 offchainPublicKey; // Observations are signed with this public key, and are only verified offchain.
  }

  struct SourceChain {
    uint64 chainSelector; // ─────╮ The Source chain selector.
    uint64 minObservers; // ──────╯ Required number of observers to agree on an observation for this source chain.
    uint256 observerNodesBitmap; // ObserverNodesBitmap & (1<<i) == (1<<i) iff nodes[i] is an observer for this source chain.
  }

  struct Config {
    // No sorting requirement for nodes, but ensure that SourceChain.observerNodeIndices in the home chain config &
    // Signer.nodeIndex in the remote chain configs are appropriately updated when changing this field.
    Node[] nodes;
    // No sorting requirement for source chains, it is most gas efficient to append new source chains to the right.
    SourceChain[] sourceChains;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
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
  uint256 private constant CONFIG_RING_BUFFER_SIZE = 2;
  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  uint256 private constant PREFIX = 0x000b << (256 - 16); // 0x000b00..00

  /// @notice This array holds just the versions of the configs, the actual configs are stored in s_configs.
  /// s_configVersions[i] == 0 iff s_configs[i] is unusable, either never set or revoked.
  uint32[CONFIG_RING_BUFFER_SIZE] private s_configVersions;
  /// @notice This array holds the digests of the configs, used for efficiency.
  /// @dev Value i in this array is valid iff s_configVersions[i] != 0.
  bytes32[CONFIG_RING_BUFFER_SIZE] private s_configDigests;
  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configVersions[i] != 0.
  Config[CONFIG_RING_BUFFER_SIZE] private s_configs;
  /// @notice The index of the latest config in the ring buffer. Given a ring buffer of 2 this values will be 0 or 1.
  /// @dev Since this value is packed with the config count, it won't flip the slot from 0 to 1 or vice versa, meaning
  /// there's no gas impact in using 0 as a value.
  uint32 private s_latestConfigIndex;
  /// @notice The total number of configs set, used for generating the version of the configs.
  uint32 private s_configCount;

  /// @notice Returns the current ring buffer size.
  function getRingBufferSize() external pure returns (uint256) {
    return CONFIG_RING_BUFFER_SIZE;
  }

  /// @notice Sets a new config.
  /// Setting a new config while the ring buffer is full will revoke the oldest config
  /// @param newConfig The new config to set.
  function setConfig(Config calldata newConfig) external onlyOwner {
    // sanity checks
    {
      // Ensure that observerNodesBitmap can be bit-encoded into a uint256.
      if (newConfig.nodes.length > 256) {
        revert OutOfBoundsNodesLength();
      }

      // Ensure no peerId or offchainPublicKey is duplicated.
      for (uint256 i = 0; i < newConfig.nodes.length; ++i) {
        for (uint256 j = i + 1; j < newConfig.nodes.length; ++j) {
          if (newConfig.nodes[i].peerId == newConfig.nodes[j].peerId) {
            revert DuplicatePeerId();
          }
          if (newConfig.nodes[i].offchainPublicKey == newConfig.nodes[j].offchainPublicKey) {
            revert DuplicateOffchainPublicKey();
          }
        }
      }

      uint256 numberOfSourceChains = newConfig.sourceChains.length;
      for (uint256 i = 0; i < numberOfSourceChains; ++i) {
        // Ensure the source chain is unique.
        for (uint256 j = i + 1; j < numberOfSourceChains; ++j) {
          if (newConfig.sourceChains[i].chainSelector == newConfig.sourceChains[j].chainSelector) {
            revert DuplicateSourceChain();
          }
        }

        // all observer node indices are valid
        uint256 bitmap = newConfig.sourceChains[i].observerNodesBitmap;
        // Check if there are any bits set for indexes outside of the expected range.
        if (bitmap & (type(uint256).max >> (256 - newConfig.nodes.length)) != bitmap) {
          revert OutOfBoundsObserverNodeIndex();
        }

        uint256 observersCount = 0;
        for (; bitmap != 0; ++observersCount) {
          bitmap &= bitmap - 1;
        }

        // minObservers are tenable
        if (newConfig.sourceChains[i].minObservers > observersCount) {
          revert MinObserversTooHigh();
        }
      }
    }

    uint256 newConfigIndex = (s_latestConfigIndex + 1) % CONFIG_RING_BUFFER_SIZE;

    // are we going to overwrite a config?
    if (s_configVersions[newConfigIndex] > 0) {
      emit ConfigRevoked(s_configDigests[newConfigIndex]);
    }

    uint32 newConfigCount = ++s_configCount;
    VersionedConfig memory newVersionedConfig = VersionedConfig({version: newConfigCount, config: newConfig});
    bytes32 newConfigDigest = _getConfigDigest(newVersionedConfig);
    s_configs[newConfigIndex] = newConfig;
    s_configVersions[newConfigIndex] = newConfigCount;
    s_configDigests[newConfigIndex] = newConfigDigest;
    s_latestConfigIndex = uint32(newConfigIndex);

    emit ConfigSet(newConfigDigest, newVersionedConfig);
  }

  /// @notice Revokes past configs, so that only the latest config remains. Call to promote staging to production. If
  /// the latest config that was set through setConfig, was subsequently revoked through revokeConfig, this function
  /// will revoke _all_ configs.
  function revokeAllConfigsButLatest() external onlyOwner {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      // Find all configs that are not the latest.
      if (s_latestConfigIndex != i && s_configVersions[i] > 0) {
        emit ConfigRevoked(_getConfigDigest(VersionedConfig({version: s_configVersions[i], config: s_configs[i]})));
        // Delete only the version, as that's what's used to determine if a config is active. This means the actual
        // config stays in storage which should significantly reduce the gas cost of overwriting that storage space in
        // the future.
        delete s_configVersions[i];
      }
    }
  }

  /// @notice Revokes a specific config by digest.
  /// @param configDigest The digest of the config to revoke.
  function revokeConfig(bytes32 configDigest) external onlyOwner {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      if (s_configDigests[i] == configDigest && s_configVersions[i] > 0) {
        emit ConfigRevoked(configDigest);
        // Delete only the version, as that's what's used to determine if a config is active. This means the actual
        // config stays in storage which should significantly reduce the gas cost of overwriting that storage space in
        // the future.
        delete s_configVersions[i];
        break;
      }
    }
  }

  // ================================================================
  // │                       Offchain getters                       |
  // │            Gas is not a concern for these functions          |
  // ================================================================

  /// @return configDigests ordered from oldest to latest set
  function getConfigDigests() external view returns (bytes32[] memory configDigests) {
    uint256 len = 0;
    for (uint256 act = 0; act <= 1; ++act) {
      if (act == 1) {
        configDigests = new bytes32[](len);
      }

      uint256 i = s_latestConfigIndex;
      do {
        if (s_configVersions[i] > 0) {
          if (act == 0) {
            ++len;
          } else if (act == 1) {
            configDigests[--len] = s_configDigests[i];
          }
        }
        if (i == 0) {
          i = CONFIG_RING_BUFFER_SIZE - 1;
        } else {
          --i;
        }
      } while (i != s_latestConfigIndex);
    }
    return configDigests;
  }

  /// @param offset setting to 0 will put the newest config in the last position of the returned array, setting to 1
  /// will put the second newest config in the last position, and so on
  /// @param limit len(versionedConfigsWithDigests) <= limit, set to 1 to get just the latest config,
  /// set to CONFIG_RING_BUFFER_SIZE to ensure that all configs are returned
  /// @return versionedConfigsWithDigests ordered from oldest to latest set
  function getVersionedConfigsWithDigests(
    uint256 offset,
    uint256 limit
  ) external view returns (VersionedConfigWithDigest[] memory versionedConfigsWithDigests) {
    uint256[2] memory ignored;
    uint256[2] memory accounted;
    uint256 len; // clobbered by the end of the loop
    for (uint256 act = 0; act <= 1; ++act) {
      if (act == 1) {
        len = accounted[0];
        versionedConfigsWithDigests = new VersionedConfigWithDigest[](len);
      }

      uint256 i = s_latestConfigIndex;
      do {
        if (accounted[act] >= limit) {
          break;
        }

        if (s_configVersions[i] > 0) {
          if (ignored[act] < offset) {
            ++ignored[act];
          } else {
            ++accounted[act];
            if (act == 1) {
              versionedConfigsWithDigests[--len] = VersionedConfigWithDigest({
                configDigest: s_configDigests[i],
                versionedConfig: VersionedConfig({version: s_configVersions[i], config: s_configs[i]})
              });
            }
          }
        }
        i = i == 0 ? CONFIG_RING_BUFFER_SIZE - 1 : i - 1;
      } while (i != s_latestConfigIndex);
    }
    return versionedConfigsWithDigests;
  }

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  function getVersionedConfig(
    bytes32 configDigest
  ) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      if (s_configVersions[i] > 0 && s_configDigests[i] == configDigest) {
        return (VersionedConfig({version: s_configVersions[i], config: s_configs[i]}), true);
      }
    }
    return (versionedConfig, false);
  }

  function _getConfigDigest(VersionedConfig memory versionedConfig) internal pure returns (bytes32) {
    return bytes32((PREFIX & PREFIX_MASK) | (uint256(keccak256(abi.encode(versionedConfig))) & ~PREFIX_MASK));
  }
}
