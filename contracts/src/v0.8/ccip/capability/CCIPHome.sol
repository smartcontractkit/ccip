// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilityConfiguration} from "../../keystone/interfaces/ICapabilityConfiguration.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {ICapabilitiesRegistry} from "./interfaces/ICapabilitiesRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Internal} from "../libraries/Internal.sol";

import {IERC165} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/interfaces/IERC165.sol";
import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice CCIPHome stores the configuration for the CCIP capability.
/// We have two classes of configuration: chain configuration and DON (in the CapabilitiesRegistry sense) configuration.
/// Each chain will have a single configuration which includes information like the router address.
/// Each CR DON will have up to four configurations: for each of (commit, exec), one blue and one green configuration.
/// This is done in order to achieve "blue-green" deployments.
contract CCIPHome is OwnerIsCreator, ITypeAndVersion, ICapabilityConfiguration, IERC165 {
  using EnumerableSet for EnumerableSet.UintSet;

  event ChainConfigRemoved(uint64 chainSelector);
  event ChainConfigSet(uint64 chainSelector, ChainConfig chainConfig);
  event ConfigSet(bytes32 indexed configDigest, VersionedConfig versionedConfig);
  event ConfigRevoked(bytes32 indexed configDigest);
  event DynamicConfigSet(bytes32 indexed configDigest, bytes dynamicConfig);
  event ConfigPromoted(bytes32 indexed configDigest);

  error OutOfBoundsNodesLength();
  error DuplicatePeerId();
  error DuplicateOffchainPublicKey();
  error DuplicateSourceChain();
  error OutOfBoundsObserverNodeIndex();
  error MinObserversTooHigh();
  error NodeNotInRegistry(bytes32 p2pId);
  error ChainSelectorNotFound(uint64 chainSelector);
  error FChainMustBePositive();
  error ChainSelectorNotSet();
  error InvalidPluginType();
  error OfframpAddressCannotBeZero();
  error FChainTooHigh(uint256 fChain, uint256 FRoleDON);
  error TooManySigners();
  error FTooHigh();
  error InvalidNode(OCR3Node node);
  error NotEnoughTransmitters(uint256 got, uint256 minimum);
  error OnlyCapabilitiesRegistryCanCall();
  error ZeroAddressNotAllowed();
  error ConfigDigestMismatch(bytes32 expectedConfigDigest, bytes32 gotConfigDigest);
  error DigestNotFound(bytes32 configDigest);

  error InvalidStateTransition(
    bytes32 currentPrimaryDigest, bytes32 currentSecondaryDigest, bytes32 blue, bytes32 green
  );

  /// @notice Represents an oracle node in OCR3 configs part of the role DON.
  /// Every configured node should be a signer, but does not have to be a transmitter.
  struct OCR3Node {
    bytes32 p2pId; // Peer2Peer connection ID of the oracle
    bytes signerKey; // On-chain signer public key
    bytes transmitterKey; // On-chain transmitter public key. Can be set to empty bytes to represent that the node is a signer but not a transmitter.
  }

  /// @notice OCR3 configuration.
  /// Note that FRoleDON >= fChain, since FRoleDON represents the role DON, and fChain represents sub-committees.
  /// FRoleDON values are typically identical across multiple OCR3 configs since the chains pertain to one role DON,
  /// but FRoleDON values can change across OCR3 configs to indicate role DON splits.
  struct OCR3Config {
    Internal.OCRPluginType pluginType; // ─╮ The plugin that the configuration is for.
    uint64 chainSelector; //               | The (remote) chain that the configuration is for.
    uint8 FRoleDON; //                     | The "big F" parameter for the role DON.
    uint64 offchainConfigVersion; // ──────╯ The version of the offchain configuration.
    bytes offrampAddress; // The remote chain offramp address.
    OCR3Node[] nodes; // Keys & IDs of nodes part of the role DON
    bytes offchainConfig; // The offchain configuration for the OCR3 protocol. Protobuf encoded.
  }

  struct VersionedConfig {
    uint32 version;
    bytes32 configDigest;
    OCR3Config config;
  }

  /// @notice Chain configuration.
  /// Changes to chain configuration are detected out-of-band in plugins and decoded offchain.
  struct ChainConfig {
    bytes32[] readers; // The P2P IDs of the readers for the chain. These IDs must be registered in the capabilities registry.
    uint8 fChain; // The fault tolerance parameter of the chain.
    bytes config; // The chain configuration. This is kept intentionally opaque so as to add fields in the future if needed.
  }

  /// @notice Chain configuration information struct used in applyChainConfigUpdates and getAllChainConfigs.
  struct ChainConfigArgs {
    uint64 chainSelector;
    ChainConfig chainConfig;
  }

  string public constant override typeAndVersion = "CCIPHome 1.6.0-dev";

  /// @dev A prefix added to all config digests that is unique to the implementation
  uint256 private constant PREFIX = 0x000a << (256 - 16); // 0x000a00..00
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev 256 is the hard limit due to the bit encoding of their indexes into a uint256.
  uint256 internal constant MAX_NUM_ORACLES = 256;

  /// @notice Used for encoding the config digest prefix
  uint256 private constant PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..00
  /// @notice The max number of configs that can be active at the same time.
  uint256 private constant MAX_CONCURRENT_CONFIGS = 2;
  /// @notice Helper to identify the zero config digest with less casting.
  bytes32 private constant ZERO_DIGEST = bytes32(uint256(0));

  /// @dev The canonical capabilities registry address.
  address internal immutable i_capabilitiesRegistry;

  /// @dev chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => ChainConfig chainConfig) private s_chainConfigurations;

  /// @dev All chains that are configured.
  EnumerableSet.UintSet private s_remoteChainSelectors;

  /// @notice This array holds the configs.
  /// @dev Value i in this array is valid iff s_configs[i].configDigest != 0.
  mapping(bytes32 pluginKey => VersionedConfig[MAX_CONCURRENT_CONFIGS]) private s_configs;

  /// @notice The total number of configs ever set, used for generating the version of the configs.
  uint32 private s_configCount = 0;
  /// @notice The index of the primary config.
  uint32 private s_primaryConfigIndex = 0;

  /// @notice Constructor for the CCIPHome contract takes in the address of the capabilities registry. This address
  /// is the only allowed caller to mutate the configuration through beforeCapabilityConfigSet.
  constructor(address capabilitiesRegistry) {
    if (capabilitiesRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_capabilitiesRegistry = capabilitiesRegistry;
  }

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
    // Config count is unused because we don't want to invalidate a config on blue/green promotions so we keep track of
    // the actual newly submitted configs instead of the number of config mutations.
    uint64, // config count
    uint32 donId
  ) external override {
    if (msg.sender != i_capabilitiesRegistry) {
      revert OnlyCapabilitiesRegistryCanCall();
    }

    (OCR3Config memory blue, OCR3Config memory green) = abi.decode(update, (OCR3Config, OCR3Config));
    bytes32 pluginKey = bytes32(uint256(donId));
    uint32 newConfigVersion = s_configCount + 1;

    (bytes32 currentBlueDigest, bytes32 currentGreenDigest) = getConfigDigests(pluginKey);
    bytes32 newBlueDigest = _calculateConfigDigest(pluginKey, abi.encode(blue), newConfigVersion);
    bytes32 newGreenDigest = _calculateConfigDigest(pluginKey, abi.encode(green), newConfigVersion);

    // Check the possible steps of the state machine
    // 1. promoteSecondaryAndRevokePrimary requires
    //   - blue digest to be the current secondary digest
    //   - green digest to be the zero digest
    if (currentGreenDigest == newBlueDigest && newGreenDigest == ZERO_DIGEST) {
      _promoteSecondaryAndRevokePrimary(pluginKey, newBlueDigest);
      return;
    }
    // setSecondary and revokeSecondary require no changes to the blue config
    if (currentBlueDigest == newBlueDigest) {
      // 2. If the green config is non-zero, we call setSecondary
      if (newGreenDigest != ZERO_DIGEST) {
        _setSecondary(pluginKey, blue, newBlueDigest);
        return;
      } else {
        // 3. If the green config is zero, we call revokeSecondary
        _revokeSecondary(pluginKey, newGreenDigest);
        return;
      }
    }

    // There are no other valid state transitions so we revert if we have not returned by now.
    revert InvalidStateTransition(currentBlueDigest, currentGreenDigest, newBlueDigest, newGreenDigest);
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
  /// @param pluginKey The key of the plugin to get the config digests for.
  /// @return primaryConfigDigest The digest of the primary config.
  /// @return secondaryConfigDigest The digest of the secondary config.
  function getConfigDigests(
    bytes32 pluginKey
  ) public view returns (bytes32 primaryConfigDigest, bytes32 secondaryConfigDigest) {
    return (
      s_configs[pluginKey][s_primaryConfigIndex].configDigest,
      s_configs[pluginKey][s_primaryConfigIndex ^ 1].configDigest
    );
  }

  /// @notice Returns the primary config digest for for a given key.
  /// @param pluginKey The key of the plugin to get the config digests for.
  function getPrimaryDigest(bytes32 pluginKey) public view returns (bytes32) {
    return s_configs[pluginKey][s_primaryConfigIndex].configDigest;
  }

  /// @notice Returns the secondary config digest for for a given key.
  /// @param pluginKey The key of the plugin to get the config digests for.
  function getSecondaryDigest(bytes32 pluginKey) public view returns (bytes32) {
    return s_configs[pluginKey][s_primaryConfigIndex ^ 1].configDigest;
  }

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  /// @param pluginKey The unique key for the DON that the configuration applies to.
  /// @param configDigest The digest of the config to fetch.
  /// @return versionedConfig The config and its version.
  /// @return ok True if the config was found, false otherwise.
  function getConfig(
    bytes32 pluginKey,
    bytes32 configDigest
  ) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    for (uint256 i = 0; i < MAX_CONCURRENT_CONFIGS; ++i) {
      // We never want to return true for a zero digest, even if the caller is asking for it, as this can expose old
      // config state that is invalid.
      if (s_configs[pluginKey][i].configDigest == configDigest && configDigest != ZERO_DIGEST) {
        return (s_configs[pluginKey][i], true);
      }
    }
    // versionConfig is uninitialized so it contains default values.
    return (versionedConfig, false);
  }

  /// @notice Returns the primary and secondary configuration for a given plugin key.
  /// @param pluginKey The unique key for the DON that the configuration applies to.
  /// @return primaryConfig The primary configuration.
  /// @return secondaryConfig The secondary configuration.
  function getAllConfigs(
    bytes32 pluginKey
  ) external view returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig) {
    VersionedConfig memory storedPrimaryConfig = s_configs[pluginKey][s_primaryConfigIndex];
    if (storedPrimaryConfig.configDigest != ZERO_DIGEST) {
      primaryConfig = storedPrimaryConfig;
    }

    VersionedConfig memory storedSecondaryConfig = s_configs[pluginKey][s_primaryConfigIndex ^ 1];
    if (storedSecondaryConfig.configDigest != ZERO_DIGEST) {
      secondaryConfig = storedSecondaryConfig;
    }

    return (primaryConfig, secondaryConfig);
  }

  // ================================================================
  // │                     State transitions                        │
  // ================================================================

  /// @notice Sets a new config as the secondary config. Does not influence the primary config.
  /// @param pluginKey The key of the plugin to set the config for.
  /// @return newConfigDigest The digest of the new config.
  function _setSecondary(
    bytes32 pluginKey,
    OCR3Config memory config,
    bytes32 newDigest
  ) internal returns (bytes32 newConfigDigest) {
    _validateConfig(config);

    bytes32 existingDigest = getSecondaryDigest(pluginKey);
    // are we going to overwrite a config? If so, emit an event.
    if (getSecondaryDigest(pluginKey) != ZERO_DIGEST) {
      emit ConfigRevoked(existingDigest);
    }

    uint32 newVersion = ++s_configCount;

    VersionedConfig memory newConfig = VersionedConfig({configDigest: newDigest, version: newVersion, config: config});

    VersionedConfig storage existingConfig = s_configs[pluginKey][s_primaryConfigIndex ^ 1];
    // TODO existingConfig.config = config;
    existingConfig.version = newVersion;
    existingConfig.configDigest = newDigest;

    emit ConfigSet(newConfig.configDigest, newConfig);

    return newDigest;
  }

  /// @notice Revokes a specific config by digest.
  /// @param pluginKey The key of the plugin to revoke the config for.
  /// @param configDigest The digest of the config to revoke. This is done to prevent accidental revokes.
  function _revokeSecondary(bytes32 pluginKey, bytes32 configDigest) internal {
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
  /// @param pluginKey The key of the plugin to promote the config for.
  /// @param digestToPromote The digest of the config to promote.
  function _promoteSecondaryAndRevokePrimary(bytes32 pluginKey, bytes32 digestToPromote) internal {
    uint256 secondaryConfigIndex = s_primaryConfigIndex ^ 1;
    if (s_configs[pluginKey][secondaryConfigIndex].configDigest != digestToPromote) {
      revert ConfigDigestMismatch(s_configs[pluginKey][secondaryConfigIndex].configDigest, digestToPromote);
    }

    uint256 primaryConfigIndex = s_primaryConfigIndex;

    delete s_configs[pluginKey][primaryConfigIndex].configDigest;

    bytes32 digestToRevoke = s_configs[pluginKey][primaryConfigIndex].configDigest;
    if (digestToRevoke != ZERO_DIGEST) {
      emit ConfigRevoked(digestToRevoke);
    }

    s_primaryConfigIndex ^= 1;

    emit ConfigPromoted(digestToPromote);
  }

  /// @notice Calculates the config digest for a given plugin key, static config, and version.
  /// @param pluginKey The key of the plugin to calculate the digest for.
  /// @param staticConfig The static part of the config.
  /// @param version The version of the config.
  /// @return The calculated config digest.
  function _calculateConfigDigest(
    bytes32 pluginKey,
    bytes memory staticConfig,
    uint32 version
  ) internal view returns (bytes32) {
    return bytes32(
      (PREFIX & PREFIX_MASK)
        | (
          uint256(
            keccak256(
              bytes.concat(abi.encode(bytes32("EVM"), block.chainid, address(this), pluginKey, version), staticConfig)
            )
          ) & ~PREFIX_MASK
        )
    );
  }

  // ================================================================
  // │                         Validation                           │
  // ================================================================

  function _validateConfig(OCR3Config memory cfg) internal view {
    if (cfg.chainSelector == 0) revert ChainSelectorNotSet();
    if (cfg.pluginType != Internal.OCRPluginType.Commit && cfg.pluginType != Internal.OCRPluginType.Execution) {
      revert InvalidPluginType();
    }
    if (cfg.offrampAddress.length == 0 || keccak256(cfg.offrampAddress) == EMPTY_ENCODED_ADDRESS_HASH) {
      revert OfframpAddressCannotBeZero();
    }
    if (!s_remoteChainSelectors.contains(cfg.chainSelector)) revert ChainSelectorNotFound(cfg.chainSelector);

    // fChain cannot exceed FRoleDON, since it is a subcommittee in the larger DON
    uint256 FRoleDON = cfg.FRoleDON;
    uint256 fChain = s_chainConfigurations[cfg.chainSelector].fChain;
    // fChain > 0 is enforced in applyChainConfigUpdates, and the presence of a chain config is checked above
    // FRoleDON != 0 because FRoleDON >= fChain is enforced here
    if (fChain > FRoleDON) {
      revert FChainTooHigh(fChain, FRoleDON);
    }

    // len(nodes) >= 3 * FRoleDON + 1
    // len(nodes) == numberOfSigners
    uint256 numberOfNodes = cfg.nodes.length;
    if (numberOfNodes > MAX_NUM_ORACLES) revert TooManySigners();
    if (numberOfNodes <= 3 * FRoleDON) revert FTooHigh();

    uint256 nonZeroTransmitters = 0;
    bytes32[] memory p2pIds = new bytes32[](numberOfNodes);
    for (uint256 i = 0; i < numberOfNodes; ++i) {
      OCR3Node memory node = cfg.nodes[i];

      // 3 * fChain + 1 <= nonZeroTransmitters <= 3 * FRoleDON + 1
      // Transmitters can be set to 0 since there can be more signers than transmitters,
      if (node.transmitterKey.length != 0) {
        nonZeroTransmitters++;
      }

      // Signer key and p2pIds must always be present
      if (node.signerKey.length == 0 || node.p2pId == bytes32(0)) {
        revert InvalidNode(node);
      }

      p2pIds[i] = node.p2pId;
    }

    // We check for chain config presence above, so fChain here must be non-zero. fChain <= FRoleDON due to the checks above.
    // There can be less transmitters than signers - so they can be set to zero (which indicates that a node is a signer, but not a transmitter).
    uint256 minTransmittersLength = 3 * fChain + 1;
    if (nonZeroTransmitters < minTransmittersLength) {
      revert NotEnoughTransmitters(nonZeroTransmitters, minTransmittersLength);
    }

    // Check that the readers are in the capabilities registry.
    _ensureInRegistry(p2pIds);
  }

  // ================================================================
  // │                    Chain Configuration                       │
  // ================================================================

  /// @notice Returns the total number of chains configured.
  /// @return The total number of chains configured.
  function getNumChainConfigurations() external view returns (uint256) {
    return s_remoteChainSelectors.length();
  }

  /// @notice Returns all the chain configurations.
  /// @param pageIndex The page index.
  /// @param pageSize The page size.
  /// @return paginatedChainConfigs chain configurations.
  function getAllChainConfigs(uint256 pageIndex, uint256 pageSize) external view returns (ChainConfigArgs[] memory) {
    uint256 numberOfChains = s_remoteChainSelectors.length();
    uint256 startIndex = pageIndex * pageSize;

    if (pageSize == 0 || startIndex >= numberOfChains) {
      return new ChainConfigArgs[](0); // Return an empty array if pageSize is 0 or pageIndex is out of bounds
    }

    uint256 endIndex = startIndex + pageSize;
    if (endIndex > numberOfChains) {
      endIndex = numberOfChains;
    }

    ChainConfigArgs[] memory paginatedChainConfigs = new ChainConfigArgs[](endIndex - startIndex);

    uint256[] memory chainSelectors = s_remoteChainSelectors.values();
    for (uint256 i = startIndex; i < endIndex; ++i) {
      uint64 chainSelector = uint64(chainSelectors[i]);
      paginatedChainConfigs[i - startIndex] =
        ChainConfigArgs({chainSelector: chainSelector, chainConfig: s_chainConfigurations[chainSelector]});
    }

    return paginatedChainConfigs;
  }

  /// @notice Sets and/or removes chain configurations.
  /// Does not validate that fChain <= FRoleDON and relies on OCR3Configs to be changed in case fChain becomes larger than the FRoleDON value.
  /// @param chainSelectorRemoves The chain configurations to remove.
  /// @param chainConfigAdds The chain configurations to add.
  function applyChainConfigUpdates(
    uint64[] calldata chainSelectorRemoves,
    ChainConfigArgs[] calldata chainConfigAdds
  ) external onlyOwner {
    // Process removals first.
    for (uint256 i = 0; i < chainSelectorRemoves.length; ++i) {
      // check if the chain selector is in s_remoteChainSelectors first.
      if (!s_remoteChainSelectors.contains(chainSelectorRemoves[i])) {
        revert ChainSelectorNotFound(chainSelectorRemoves[i]);
      }

      delete s_chainConfigurations[chainSelectorRemoves[i]];
      s_remoteChainSelectors.remove(chainSelectorRemoves[i]);

      emit ChainConfigRemoved(chainSelectorRemoves[i]);
    }

    // Process additions next.
    for (uint256 i = 0; i < chainConfigAdds.length; ++i) {
      ChainConfig memory chainConfig = chainConfigAdds[i].chainConfig;
      uint64 chainSelector = chainConfigAdds[i].chainSelector;

      // Verify that the provided readers are present in the capabilities registry.
      _ensureInRegistry(chainConfig.readers);

      // Verify that fChain is positive.
      if (chainConfig.fChain == 0) {
        revert FChainMustBePositive();
      }

      s_chainConfigurations[chainSelector] = chainConfig;
      s_remoteChainSelectors.add(chainSelector);

      emit ChainConfigSet(chainSelector, chainConfig);
    }
  }

  /// @notice Helper function to ensure that a node is in the capabilities registry.
  /// @param p2pIds The P2P IDs of the node to check.
  function _ensureInRegistry(bytes32[] memory p2pIds) internal view {
    for (uint256 i = 0; i < p2pIds.length; ++i) {
      // TODO add a method that does the validation in the ICapabilitiesRegistry contract
      if (ICapabilitiesRegistry(i_capabilitiesRegistry).getNode(p2pIds[i]).p2pId == bytes32("")) {
        revert NodeNotInRegistry(p2pIds[i]);
      }
    }
  }
}
