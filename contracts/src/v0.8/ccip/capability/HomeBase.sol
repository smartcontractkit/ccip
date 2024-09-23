// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilityConfiguration} from "../../keystone/interfaces/ICapabilityConfiguration.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {IERC165} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/interfaces/IERC165.sol";

abstract contract HomeBase is OwnerIsCreator, ITypeAndVersion, ICapabilityConfiguration, IERC165 {
  event ConfigSet(StoredConfig versionedConfig);
  event ConfigRevoked(bytes32 indexed configDigest);
  event DynamicConfigSet(bytes32 indexed configDigest, bytes dynamicConfig);
  event ConfigPromoted(bytes32 indexed configDigest);

  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);
  error DigestNotFound(bytes32 configDigest);
  error ZeroAddressNotAllowed();
  error OnlyCapabilitiesRegistryCanCall();
  error OnlyOwnerOrSelfCallAllowed();

  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  /// @notice The max number of configs that can be active at the same time.
  uint256 private constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Helper to identify the zero config digest with less casting.
  bytes32 private constant ZERO_DIGEST = bytes32(uint256(0));

  /// @dev The canonical capabilities registry address.
  address internal immutable i_capabilitiesRegistry;

  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configs[i].configDigest != 0.
  mapping(bytes32 pluginKey => StoredConfig[MAX_CONCURRENT_CONFIGS]) private s_configs;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 private s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 private s_primaryConfigIndex = 0;

  struct StoredConfig {
    bytes32 configDigest;
    uint32 version;
    bytes staticConfig;
    bytes dynamicConfig;
  }

  /// @param capabilitiesRegistry the canonical capabilities registry address.
  constructor(address capabilitiesRegistry) {
    if (capabilitiesRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_capabilitiesRegistry = capabilitiesRegistry;
  }

  function _validateStaticAndDynamicConfig(bytes memory staticConfig, bytes memory dynamicConfig) internal view virtual;

  function _validateDynamicConfig(bytes memory staticConfig, bytes memory dynamicConfig) internal view virtual;

  function _getConfigDigestPrefix() internal pure virtual returns (uint256);

  // ================================================================
  // │                    Capability Registry                       │
  // ================================================================

  /// @notice Returns the capabilities registry address.
  /// @return The capabilities registry address.
  function getCapabilityRegistry() external view returns (address) {
    return i_capabilitiesRegistry;
  }

  /// @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) external pure override returns (bool) {
    return interfaceId == type(ICapabilityConfiguration).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  /// @notice Called by the registry prior to the config being set for a particular DON.
  /// @dev precondition Requires destination chain config to be set
  function beforeCapabilityConfigSet(
    bytes32[] calldata, // nodes
    bytes calldata update,
    uint64, // configCount
    uint32 // donId
  ) external override {
    if (msg.sender != i_capabilitiesRegistry) {
      revert OnlyCapabilitiesRegistryCanCall();
    }
    (bool success, bytes memory errorData) = address(this).call(update);
    if (!success) {
      revert(string(errorData));
    }
  }

  /// @inheritdoc ICapabilityConfiguration
  /// @dev The CCIP capability will fetch the configuration needed directly from this contract.
  /// The offchain syncer will call this function, so its important that it doesn't revert.
  function getCapabilityConfiguration(uint32 /* donId */ ) external pure override returns (bytes memory configuration) {
    return bytes("");
  }

  // ================================================================
  // │                          Getters                             │
  // ================================================================

  /// @notice Returns the current primary and secondary config digests.
  /// @dev Can be bytes32(0) if no config has been set yet or it has been revoked.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests(
    bytes32 pluginKey
  ) external view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (
      s_configs[pluginKey][s_primaryConfigIndex].configDigest,
      s_configs[pluginKey][s_primaryConfigIndex ^ 1].configDigest
    );
  }

  function getPrimaryDigest(bytes32 pluginKey) public view returns (bytes32) {
    return s_configs[pluginKey][s_primaryConfigIndex].configDigest;
  }

  function getSecondaryDigest(bytes32 pluginKey) public view returns (bytes32) {
    return s_configs[pluginKey][s_primaryConfigIndex ^ 1].configDigest;
  }

  /// @notice Returns the stored config for a given digest. Will always return an empty config if the digest is the zero
  /// digest. This is done to prevent exposing old config state that is invalid.
  function _getStoredConfig(
    bytes32 pluginKey,
    bytes32 configDigest
  ) internal view returns (StoredConfig memory storedConfig, bool ok) {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      // We never want to return true for a zero digest, even if the caller is asking for it, as this can expose old
      // config state that is invalid.
      if (s_configs[pluginKey][i].configDigest == configDigest && configDigest != ZERO_DIGEST) {
        return (s_configs[pluginKey][i], true);
      }
    }
    return (storedConfig, false);
  }

  function _getPrimaryStoredConfig(
    bytes32 pluginKey
  ) internal view returns (StoredConfig memory primaryConfig, bool ok) {
    if (s_configs[pluginKey][s_primaryConfigIndex].configDigest == ZERO_DIGEST) {
      return (StoredConfig(ZERO_DIGEST, 0, "", ""), false);
    }

    return (s_configs[pluginKey][s_primaryConfigIndex], true);
  }

  function _getSecondaryStoredConfig(
    bytes32 pluginKey
  ) internal view returns (StoredConfig memory secondaryConfig, bool ok) {
    if (s_configs[pluginKey][s_primaryConfigIndex ^ 1].configDigest == ZERO_DIGEST) {
      return (StoredConfig(ZERO_DIGEST, 0, "", ""), false);
    }

    return (s_configs[pluginKey][s_primaryConfigIndex ^ 1], true);
  }

  // ================================================================
  // │                     State transitions                        │
  // ================================================================

  /// @notice Sets a new config as the secondary config. Does not influence the primary config.
  /// @param digestToOverwrite The digest of the config to overwrite, or ZERO_DIGEST if no config is to be overwritten.
  /// This is done to prevent accidental overwrites.
  function setSecondary(
    bytes32 pluginKey,
    bytes calldata encodedStaticConfig,
    bytes calldata encodedDynamicConfig,
    bytes32 digestToOverwrite
  ) external OnlyOwnerOrSelfCall returns (bytes32 newConfigDigest) {
    _validateStaticAndDynamicConfig(encodedStaticConfig, encodedDynamicConfig);

    bytes32 existingDigest = getSecondaryDigest(pluginKey);

    if (existingDigest != digestToOverwrite) {
      revert ConfigDigestMismatch(existingDigest, digestToOverwrite);
    }

    // are we going to overwrite a config? If so, emit an event.
    if (existingDigest != ZERO_DIGEST) {
      emit ConfigRevoked(digestToOverwrite);
    }

    uint32 newVersion = ++s_configCount;
    newConfigDigest = _calculateConfigDigest(pluginKey, encodedStaticConfig, newVersion);

    StoredConfig memory newConfig = StoredConfig({
      configDigest: newConfigDigest,
      version: newVersion,
      staticConfig: encodedStaticConfig,
      dynamicConfig: encodedDynamicConfig
    });

    s_configs[pluginKey][s_primaryConfigIndex ^ 1] = newConfig;

    emit ConfigSet(newConfig);

    return newConfigDigest;
  }

  /// @notice Revokes a specific config by digest.
  /// @param configDigest The digest of the config to revoke. This is done to prevent accidental revokes.
  function revokeSecondary(bytes32 pluginKey, bytes32 configDigest) external OnlyOwnerOrSelfCall {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configs[pluginKey][secondaryConfigIndex].configDigest != configDigest) {
      revert ConfigDigestMismatch(s_configs[pluginKey][secondaryConfigIndex].configDigest, configDigest);
    }

    emit ConfigRevoked(configDigest);
    // Delete only the digest, as that's what's used to determine if a config is active. This means the actual
    // config stays in storage which should significantly reduce the gas cost of overwriting that storage space in
    // the future.
    delete s_configs[pluginKey][secondaryConfigIndex].configDigest;
  }

  /// @notice Promotes the secondary config to the primary config and revokes the primary config.
  function promoteSecondaryAndRevokePrimary(
    bytes32 pluginKey,
    bytes32 digestToPromote,
    bytes32 digestToRevoke
  ) external OnlyOwnerOrSelfCall {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configs[pluginKey][secondaryConfigIndex].configDigest != digestToPromote) {
      revert ConfigDigestMismatch(s_configs[pluginKey][secondaryConfigIndex].configDigest, digestToPromote);
    }

    uint256 primaryConfigIndex = s_primaryConfigIndex;
    if (s_configs[pluginKey][primaryConfigIndex].configDigest != digestToRevoke) {
      revert ConfigDigestMismatch(s_configs[pluginKey][primaryConfigIndex].configDigest, digestToRevoke);
    }

    delete s_configs[pluginKey][primaryConfigIndex].configDigest;

    s_primaryConfigIndex ^= 1;
    if (digestToRevoke != ZERO_DIGEST) {
      emit ConfigRevoked(digestToRevoke);
    }
    emit ConfigPromoted(digestToPromote);
  }

  function setDynamicConfig(
    bytes32 pluginKey,
    bytes calldata newDynamicConfig,
    bytes32 currentDigest
  ) external OnlyOwnerOrSelfCall {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      if (s_configs[pluginKey][i].configDigest == currentDigest && currentDigest != ZERO_DIGEST) {
        _validateDynamicConfig(s_configs[pluginKey][i].staticConfig, newDynamicConfig);

        // Since the static config doesn't change we don't have to update the digest or version.
        s_configs[pluginKey][i].dynamicConfig = newDynamicConfig;

        emit DynamicConfigSet(currentDigest, newDynamicConfig);
        return;
      }
    }

    revert DigestNotFound(currentDigest);
  }

  function _calculateConfigDigest(
    bytes32 pluginKey,
    bytes memory staticConfig,
    uint32 version
  ) internal view returns (bytes32) {
    return bytes32(
      (_getConfigDigestPrefix() & PREFIX_MASK)
        | (
          uint256(
            keccak256(
              bytes.concat(abi.encode(bytes32("EVM"), block.chainid, address(this), pluginKey, version), staticConfig)
            )
          ) & ~PREFIX_MASK
        )
    );
  }

  modifier OnlyOwnerOrSelfCall() {
    if (msg.sender != owner() && msg.sender != address(this)) {
      revert OnlyOwnerOrSelfCallAllowed();
    }
    _;
  }
}
