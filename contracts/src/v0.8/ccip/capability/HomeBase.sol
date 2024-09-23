// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

abstract contract HomeBase is OwnerIsCreator, ITypeAndVersion {
  event ConfigRevoked(bytes32 configDigest);

  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);

  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  /// @notice The max number of configs that can be active at the same time.
  uint256 internal constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Helper to identify the zero config digest with less casting.
  bytes32 internal constant ZERO_DIGEST = bytes32(uint256(0));

  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configs[i].configDigest != 0.
  StoredConfig[MAX_CONCURRENT_CONFIGS] internal s_configs;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 internal s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 internal s_primaryConfigIndex = 0;

  struct StoredConfig {
    bytes32 configDigest;
    uint32 version;
    bytes staticConfig;
    bytes dynamicConfig;
  }

  /// @notice Returns the stored config for a given digest. Will always return an empty config if the digest is the zero
  /// digest. This is done to prevent exposing old config state that is invalid.
  function _getStoredConfig(bytes32 configDigest) internal view returns (StoredConfig memory storedConfig, bool ok) {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      // We never want to return true for a zero digest, even if the caller is asking for it, as this can expose old
      // config state that is invalid.
      if (s_configs[i].configDigest == configDigest && configDigest != ZERO_DIGEST) {
        return (s_configs[i], true);
      }
    }
    return (storedConfig, false);
  }

  function _getPrimaryStoredConfig() internal view returns (StoredConfig memory primaryConfig, bool ok) {
    if (s_configs[s_primaryConfigIndex].configDigest == ZERO_DIGEST) {
      return (StoredConfig(ZERO_DIGEST, 0, "", ""), false);
    }

    return (s_configs[s_primaryConfigIndex], true);
  }

  function _getSecondaryStoredConfig() internal view returns (StoredConfig memory secondaryConfig, bool ok) {
    if (s_configs[s_primaryConfigIndex ^ 1].configDigest == ZERO_DIGEST) {
      return (StoredConfig(ZERO_DIGEST, 0, "", ""), false);
    }

    return (s_configs[s_primaryConfigIndex ^ 1], true);
  }

  /// @notice Returns the current primary and secondary config digests.
  /// @dev Can be bytes32(0) if no config has been set yet or it has been revoked.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests() external view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (s_configs[s_primaryConfigIndex].configDigest, s_configs[s_primaryConfigIndex ^ 1].configDigest);
  }

  /// @notice Revokes a specific config by digest.
  /// @param configDigest The digest of the config to revoke. This is done to prevent accidental revokes.
  function revokeSecondary(bytes32 configDigest) external onlyOwner {
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
  function promoteSecondaryAndRevokePrimary(bytes32 digestToPromote, bytes32 digestToRevoke) external onlyOwner {
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
    emit ConfigRevoked(digestToRevoke);
  }

  function _calculateConfigDigest(
    bytes memory staticConfig,
    uint32 version,
    uint256 prefix
  ) internal view returns (bytes32) {
    return bytes32(
      (prefix & PREFIX_MASK)
        | (
          uint256(
            keccak256(bytes.concat(abi.encode(bytes32("EVM"), block.chainid, address(this), version), staticConfig))
          ) & ~PREFIX_MASK
        )
    );
  }
}
