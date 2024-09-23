// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

abstract contract HomeBase is OwnerIsCreator, ITypeAndVersion {
  event ConfigRevoked(bytes32 configDigest);

  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);

  /// @notice The max number of configs that can be active at the same time.
  uint256 internal constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Used for encoding the config digest prefix
  uint256 internal constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  bytes32 internal constant ZERO_DIGEST = bytes32(uint256(0));

  /// @notice This array holds the digests of the configs, used for efficiency.
  /// @dev Value i in this array is valid iff it's not 0.
  bytes32[MAX_CONCURRENT_CONFIGS] internal s_configDigests;
  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configDigests[i] != 0.
  StoredConfig[MAX_CONCURRENT_CONFIGS] internal s_configs;
  /// @notice This array holds the versions of the configs.
  /// @dev Value i in this array is valid iff s_configDigests[i] != 0.
  /// @dev Since Solidity doesn't support writing complex memory structs to storage, we have to make the config calldata
  /// in setConfig and then copy it to storage. This does not allow us to modify it to add the version field, so we
  /// store the version separately.
  uint32[MAX_CONCURRENT_CONFIGS] internal s_configVersions;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 internal s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 internal s_primaryConfigIndex = 0;

  struct StoredConfig {
    bytes staticConfig;
    bytes dynamicConfig;
  }

  /// @notice Returns the current primary and secondary config digests.
  /// @dev Can be bytes32(0) if no config has been set yet or it has been revoked.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests() external view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (s_configDigests[s_primaryConfigIndex], s_configDigests[s_primaryConfigIndex ^ 1]);
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
