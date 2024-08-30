// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/Ownable2Step.sol";

import {ITypeAndVersion} from "../shared/interfaces/ITypeAndVersion.sol";

/// @notice Stores the home configuration for RMN, that is referenced by CCIP oracles, RMN nodes, and the RMNRemote
/// contracts.
contract RMNHome is Ownable2Step, ITypeAndVersion {
  /// @dev temp placeholder to exclude this contract from coverage
  function test() public {}

  string public constant override typeAndVersion = "RMNHome 1.6.0-dev";

  struct Node {
    bytes32 peerId; // used for p2p communication
    bytes32 offchainPublicKey; // observations are signed with this public key, and are only verified offchain
  }

  struct SourceChain {
    uint64 chainSelector;
    uint64 minObservers; // required to agree on an observation for this source chain
    uint256 observerNodesBitmap; // observerNodesBitmap & (1<<i) == (1<<i) iff nodes[i] is an observer for this source chain
  }

  struct Config {
    // No sorting requirement for nodes, but ensure that SourceChain.observerNodeIndices in the home chain config &
    // Signer.nodeIndex in the remote chain configs are appropriately updated when changing this field
    Node[] nodes;
    // No sorting requirement for source chains, it is most gas efficient to append new source chains to the right.
    SourceChain[] sourceChains;
    // Offchain configuration
    bytes offchainConfig;
  }

  struct VersionedConfig {
    uint32 version;
    Config config;
  }

  function _configDigest(
    VersionedConfig memory versionedConfig
  ) internal view returns (bytes32) {
    uint256 h = uint256(keccak256(abi.encode(bytes32("EVM"), block.chainid, address(this), versionedConfig)));
    uint256 prefixMask = type(uint256).max << (256 - 16); // 0xFFFF00..00
    uint256 prefix = 0x000b << (256 - 16); // 0x000b00..00
    return bytes32((prefix & prefixMask) | (h & ~prefixMask));
  }

  // if we were to have VersionedConfig instead of Config in the ring buffer, we couldn't assign directly to it in
  // setConfig without via-ir
  uint256 public constant CONFIG_RING_BUFFER_SIZE = 2;
  uint32 s_configCount;
  uint32[CONFIG_RING_BUFFER_SIZE] s_configVersions; // s_configVersions[i] == 0 iff s_configs[i] is unusable
  bytes32[CONFIG_RING_BUFFER_SIZE] s_configDigests;
  Config[CONFIG_RING_BUFFER_SIZE] s_configs;
  uint256 s_latestConfigIndex;

  function setConfig(
    Config calldata newConfig
  ) external onlyOwner {
    // sanity checks
    {
      if (newConfig.nodes.length > 256) {
        // so that observerNodesBitmap can be uint256
        revert OutOfBoundsNodesLength();
      }

      // no peerId or offchainPublicKey is duplicated
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

      for (uint256 i = 0; i < newConfig.sourceChains.length; ++i) {
        // the source chain is unique
        for (uint256 j = i + 1; j < newConfig.sourceChains.length; ++j) {
          if (newConfig.sourceChains[i].chainSelector == newConfig.sourceChains[j].chainSelector) {
            revert DuplicateSourceChain();
          }
        }

        // all observer node indices are valid
        uint256 bitmap = newConfig.sourceChains[i].observerNodesBitmap;
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

    uint256 oldConfigIndex = s_latestConfigIndex;
    uint256 newConfigIndex = (oldConfigIndex + 1) % CONFIG_RING_BUFFER_SIZE;

    // are we going to overwrite a config?
    if (s_configVersions[newConfigIndex] > 0) {
      emit ConfigRevoked(s_configDigests[newConfigIndex]);
    }

    uint32 newConfigCount = ++s_configCount;
    VersionedConfig memory newVersionedConfig = VersionedConfig({version: newConfigCount, config: newConfig});
    bytes32 newConfigDigest = _configDigest(newVersionedConfig);
    s_configs[newConfigIndex] = newConfig;
    s_configVersions[newConfigIndex] = newConfigCount;
    s_configDigests[newConfigIndex] = newConfigDigest;
    s_latestConfigIndex = newConfigIndex;
    emit ConfigSet(newConfigDigest, newVersionedConfig);
  }

  /// @notice Revokes past configs, so that only the latest config remains. Call to promote staging to production. If
  /// the latest config that was set through setConfig, was subsequently revoked through revokeConfig, this function
  /// will revoke _all_ configs.
  function revokeAllConfigsButLatest() external onlyOwner {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      if (s_latestConfigIndex != i && s_configVersions[i] > 0) {
        emit ConfigRevoked(_configDigest(VersionedConfig({version: s_configVersions[i], config: s_configs[i]})));
        delete s_configVersions[i];
      }
    }
  }

  function revokeConfig(
    bytes32 configDigest
  ) external onlyOwner {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      if (s_configVersions[i] > 0 && s_configDigests[i] == configDigest) {
        emit ConfigRevoked(configDigest);
        delete s_configVersions[i];
        break;
      }
    }
  }

  ///
  /// Offchain getters
  /// Only to be called by offchain code, efficiency is not a concern
  ///

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
        i = i == 0 ? CONFIG_RING_BUFFER_SIZE - 1 : i - 1;
      } while (i != s_latestConfigIndex);
    }
  }

  struct VersionedConfigWithDigest {
    bytes32 configDigest;
    VersionedConfig versionedConfig;
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
  }

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  function getVersionedConfig(
    bytes32 configDigest
  ) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    for (uint256 i = 0; i < CONFIG_RING_BUFFER_SIZE; ++i) {
      if (s_configVersions[i] > 0 && s_configDigests[i] == configDigest) {
        versionedConfig = VersionedConfig({version: s_configVersions[i], config: s_configs[i]});
        ok = true;
        break;
      }
    }
  }

  ///
  /// Events
  ///

  event ConfigSet(bytes32 configDigest, VersionedConfig versionedConfig);
  event ConfigRevoked(bytes32 configDigest);

  ///
  /// Errors
  ///

  error OutOfBoundsNodesLength();
  error DuplicatePeerId();
  error DuplicateOffchainPublicKey();
  error DuplicateSourceChain();
  error OutOfBoundsObserverNodeIndex();
  error MinObserversTooHigh();
}
