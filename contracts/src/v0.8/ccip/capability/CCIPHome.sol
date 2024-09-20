// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilityConfiguration} from "../../keystone/interfaces/ICapabilityConfiguration.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {ICapabilitiesRegistry} from "./interfaces/ICapabilitiesRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Internal} from "../libraries/Internal.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice CCIPHome stores the configuration for the CCIP capability.
/// We have two classes of configuration: chain configuration and DON (in the CapabilitiesRegistry sense) configuration.
/// Each chain will have a single configuration which includes information like the router address.
/// Each CR DON will have up to four configurations: for each of (commit, exec), one blue and one green configuration.
/// This is done in order to achieve "blue-green" deployments.
contract CCIPHome is ITypeAndVersion, ICapabilityConfiguration, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.UintSet;

  /// @notice Emitted when a chain's configuration is set.
  /// @param chainSelector The chain selector.
  /// @param chainConfig The chain configuration.
  event ChainConfigSet(uint64 chainSelector, ChainConfig chainConfig);
  /// @notice Emitted when a chain's configuration is removed.
  /// @param chainSelector The chain selector.
  event ChainConfigRemoved(uint64 chainSelector);
  event ConfigSet(uint32 indexed donId, uint8 indexed pluginType, OCR3ConfigWithMeta[] config);

  error NodeNotInRegistry(bytes32 p2pId);
  error OnlyCapabilitiesRegistryCanCall();
  error ChainSelectorNotFound(uint64 chainSelector);
  error ChainSelectorNotSet();
  error TooManyOCR3Configs();
  error TooManySigners();
  error InvalidNode(OCR3Node node);
  error NotEnoughTransmitters(uint256 got, uint256 minimum);
  error FChainMustBePositive();
  error FTooHigh();
  error FChainTooHigh(uint256 fChain, uint256 FRoleDON);
  error InvalidPluginType();
  error PluginTypeMismatch(
    bytes32 configDigest, Internal.OCRPluginType expectedPluginType, Internal.OCRPluginType gotPluginType
  );
  error OfframpAddressCannotBeZero();
  error InvalidConfigLength(uint256 length);
  error InvalidConfigStateTransition(ConfigState currentState, ConfigState proposedState);
  error NonExistentConfigTransition();
  error WrongConfigCount(uint64 got, uint64 expected);
  error WrongConfigDigest(bytes32 got, bytes32 expected);
  error WrongConfigDigestBlueGreen(bytes32 got, bytes32 expected);
  error ZeroAddressNotAllowed();

  /// @notice ConfigState indicates the state of the configuration.
  /// A DON's configuration always starts out in the "Init" state - this is the starting state.
  /// The only valid transition from "Init" is to the "Running" state - this is the first ever configuration.
  /// The only valid transition from "Running" is to the "Staging" state - this is a blue/green proposal.
  /// The only valid transition from "Staging" is back to the "Running" state - this is a promotion.
  /// In order to rollback a configuration, we must therefore do the following:
  /// - Suppose that we have a correct configuration in the "Running" state (V1).
  /// - We propose a new configuration and transition to the "Staging" state (V2).
  /// - V2 turns out to be buggy
  /// - In the same transaction, we must:
  ///   - Promote V2
  ///   - Re-propose V1
  ///   - Promote V1
  enum ConfigState {
    Init,
    Running,
    Staging
  }

  /// @notice Chain configuration.
  /// Changes to chain configuration are detected out-of-band in plugins and decoded offchain.
  struct ChainConfig {
    bytes32[] readers; // The P2P IDs of the readers for the chain. These IDs must be registered in the capabilities registry.
    uint8 fChain; // The fault tolerance parameter of the chain.
    bytes config; // The chain configuration. This is kept intentionally opaque so as to add fields in the future if needed.
  }

  /// @notice Chain configuration information struct used in applyChainConfigUpdates and getAllChainConfigs.
  struct ChainConfigInfo {
    uint64 chainSelector;
    ChainConfig chainConfig;
  }

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

  /// @notice OCR3 configuration with metadata, specifically the config count and the config digest.
  struct OCR3ConfigWithMeta {
    OCR3Config config; // The OCR3 configuration.
    uint64 version; // The version used to compute the config digest, starts at 1 and increments by 1 for each new config.
    bytes32 configDigest; // The config digest of the OCR3 configuration.
  }

  /// @dev Type and version override.
  string public constant override typeAndVersion = "CCIPHome 1.6.0-dev";

  /// @dev The canonical capabilities registry address.
  address internal immutable i_capabilitiesRegistry;

  uint8 internal constant MAX_OCR3_CONFIGS_PER_PLUGIN = 2;
  /// @dev We have two plugins: commit and execution, for each we can store a blue and a green config.
  uint8 internal constant MAX_OCR3_CONFIGS_PER_DON = 2 * MAX_OCR3_CONFIGS_PER_PLUGIN;
  uint256 internal constant CONFIG_DIGEST_PREFIX_MASK = type(uint256).max << (256 - 16); // 0xFFFF00..0
  /// @dev must be equal to libocr multi role: https://github.com/smartcontractkit/libocr/blob/ae747ca5b81236ffdbf1714318c652e923a5ff4d/offchainreporting2plus/types/config_digest.go#L28
  uint256 internal constant CONFIG_DIGEST_PREFIX = 0x000a << (256 - 16); // 0x000a00..00
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev 256 is the hard limit due to the bit encoding of their indexes into a uint256.
  uint256 internal constant MAX_NUM_ORACLES = 256;

  /// @dev chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => ChainConfig chainConfig) private s_chainConfigurations;

  /// @dev All chains that are configured.
  EnumerableSet.UintSet private s_remoteChainSelectors;

  /// @dev OCR3 configurations for each DON.
  /// Each CR DON will have a commit and execution configuration.
  /// This means that a DON can have up to 4 configurations, since we are implementing blue/green deployments.
  mapping(uint32 donId => mapping(Internal.OCRPluginType pluginType => OCR3ConfigWithMeta[] ocr3Configs)) private
    s_ocr3Configs;

  /// @param capabilitiesRegistry the canonical capabilities registry address.
  constructor(address capabilitiesRegistry) {
    if (capabilitiesRegistry == address(0)) {
      revert ZeroAddressNotAllowed();
    }
    i_capabilitiesRegistry = capabilitiesRegistry;
  }

  // ================================================================
  // │                    Config Getters                            │
  // ================================================================

  /// @notice Returns the capabilities registry address.
  /// @return The capabilities registry address.
  function getCapabilityRegistry() external view returns (address) {
    return i_capabilitiesRegistry;
  }

  /// @notice Returns the total number of chains configured.
  /// @return The total number of chains configured.
  function getNumChainConfigurations() external view returns (uint256) {
    return s_remoteChainSelectors.length();
  }

  /// @notice Returns all the chain configurations.
  /// @param pageIndex The page index.
  /// @param pageSize The page size.
  /// @return paginatedChainConfigs chain configurations.
  function getAllChainConfigs(uint256 pageIndex, uint256 pageSize) external view returns (ChainConfigInfo[] memory) {
    uint256 totalItems = s_remoteChainSelectors.length(); // Total number of chain selectors
    uint256 startIndex = pageIndex * pageSize;

    if (pageSize == 0 || startIndex >= totalItems) {
      return new ChainConfigInfo[](0); // Return an empty array if pageSize is 0 or pageIndex is out of bounds
    }

    uint256 endIndex = startIndex + pageSize;
    if (endIndex > totalItems) {
      endIndex = totalItems;
    }

    ChainConfigInfo[] memory paginatedChainConfigs = new ChainConfigInfo[](endIndex - startIndex);

    uint256[] memory chainSelectors = s_remoteChainSelectors.values();
    for (uint256 i = startIndex; i < endIndex; ++i) {
      uint64 chainSelector = uint64(chainSelectors[i]);
      paginatedChainConfigs[i - startIndex] =
        ChainConfigInfo({chainSelector: chainSelector, chainConfig: s_chainConfigurations[chainSelector]});
    }

    return paginatedChainConfigs;
  }

  /// @notice Returns the OCR configuration for the given don ID and plugin type.
  /// @param donId The DON ID.
  /// @param pluginType The plugin type.
  /// @return The OCR3 configurations, up to 2 (blue and green).
  function getOCRConfig(
    uint32 donId,
    Internal.OCRPluginType pluginType
  ) external view returns (OCR3ConfigWithMeta[] memory) {
    return s_ocr3Configs[donId][pluginType];
  }

  // ================================================================
  // │                    Capability Configuration                  │
  // ================================================================

  /// @inheritdoc ICapabilityConfiguration
  /// @dev The CCIP capability will fetch the configuration needed directly from this contract.
  /// The offchain syncer will call this function, however, so its important that it doesn't revert.
  function getCapabilityConfiguration(uint32 /* donId */ ) external pure override returns (bytes memory configuration) {
    return bytes("");
  }

  /// @notice Called by the registry prior to the config being set for a particular DON.
  /// @param config The configuration, abi encoded as (OCR3Config[] commitConfigs, OCR3Config[] execConfigs).
  /// @dev precondition Requires destination chain config to be set
  function beforeCapabilityConfigSet(
    bytes32[] calldata, // nodes
    bytes calldata config,
    uint64, // configCount
    uint32 donId
  ) external override {
    if (msg.sender != i_capabilitiesRegistry) {
      revert OnlyCapabilitiesRegistryCanCall();
    }

    (OCR3Config[] memory commitConfigs, OCR3Config[] memory execConfigs) =
      abi.decode(config, (OCR3Config[], OCR3Config[]));

    if (commitConfigs.length > 0) {
      _updatePluginConfig(donId, Internal.OCRPluginType.Commit, commitConfigs);
    }
    if (execConfigs.length > 0) {
      _updatePluginConfig(donId, Internal.OCRPluginType.Execution, execConfigs);
    }
  }

  /// @notice Sets a new OCR3 config for a specific plugin type for a DON.
  /// @param donId The DON ID.
  /// @param pluginType The plugin type.
  /// @param newConfig The new configuration.
  function _updatePluginConfig(uint32 donId, Internal.OCRPluginType pluginType, OCR3Config[] memory newConfig) internal {
    OCR3ConfigWithMeta[] memory currentConfig = s_ocr3Configs[donId][pluginType];

    // Validate the state transition being proposed, which is implicitly defined by the combination
    // of lengths of the current and new configurations.
    ConfigState currentState = _stateFromConfigLength(currentConfig.length);
    ConfigState proposedState = _stateFromConfigLength(newConfig.length);
    _validateConfigStateTransition(currentState, proposedState);

    // Build the new configuration with metadata and validate that the transition is valid.
    OCR3ConfigWithMeta[] memory newConfigWithMeta =
      _computeNewConfigWithMeta(donId, currentConfig, newConfig, currentState, proposedState);

    _validateConfigTransition(currentConfig, newConfigWithMeta);

    // Update contract state with new configuration if its valid.
    // We won't run out of gas from this delete since the array is at most 2 elements long.
    delete s_ocr3Configs[donId][pluginType];
    for (uint256 i = 0; i < newConfigWithMeta.length; ++i) {
      OCR3Config memory newOcr3Config = newConfigWithMeta[i].config;

      if (newOcr3Config.pluginType != pluginType) {
        revert PluginTypeMismatch(newConfigWithMeta[i].configDigest, pluginType, newOcr3Config.pluginType);
      }

      // Struct has to be manually copied since there is a nested OCR3Node array. Direct assignment
      // will result in Unimplemented Feature issue.
      OCR3ConfigWithMeta storage ocr3ConfigWithMeta = s_ocr3Configs[donId][pluginType].push();
      ocr3ConfigWithMeta.configDigest = newConfigWithMeta[i].configDigest;
      ocr3ConfigWithMeta.version = newConfigWithMeta[i].version;

      OCR3Config storage ocr3Config = ocr3ConfigWithMeta.config;
      ocr3Config.pluginType = newOcr3Config.pluginType;
      ocr3Config.chainSelector = newOcr3Config.chainSelector;
      ocr3Config.FRoleDON = newOcr3Config.FRoleDON;
      ocr3Config.offchainConfigVersion = newOcr3Config.offchainConfigVersion;
      ocr3Config.offrampAddress = newOcr3Config.offrampAddress;
      ocr3Config.offchainConfig = newOcr3Config.offchainConfig;

      for (uint256 j = 0; j < newOcr3Config.nodes.length; ++j) {
        ocr3Config.nodes.push(newOcr3Config.nodes[j]);
      }
    }

    emit ConfigSet(donId, uint8(pluginType), newConfigWithMeta);
  }

  // ================================================================
  // │                    Config State Machine                      │
  // ================================================================

  /// @notice Determine the config state of the configuration from the length of the config.
  /// @param configLen The length of the configuration.
  /// @return The config state.
  function _stateFromConfigLength(uint256 configLen) internal pure returns (ConfigState) {
    if (configLen > 2) {
      revert InvalidConfigLength(configLen);
    }
    return ConfigState(configLen);
  }

  /// @notice Validates the state transition between two config states.
  /// The only valid state transitions are the following:
  /// Init    -> Running (first ever config)
  /// Running -> Staging (blue/green proposal)
  /// Staging -> Running (promotion)
  /// Everything else is invalid and should revert.
  /// @param currentState The current state.
  /// @param newState The new state.
  function _validateConfigStateTransition(ConfigState currentState, ConfigState newState) internal pure {
    // Calculate the difference between the new state and the current state
    int256 stateDiff = int256(uint256(newState)) - int256(uint256(currentState));

    // Check if the state transition is valid:
    // Valid transitions:
    // 1. currentState -> newState (where stateDiff == 1)
    //    e.g., init -> running or running -> staging
    // 2. staging -> running (where stateDiff == -1)
    if (stateDiff == 1 || (stateDiff == -1 && currentState == ConfigState.Staging)) {
      return;
    }
    revert InvalidConfigStateTransition(currentState, newState);
  }

  /// @notice Validates the transition between two OCR3 configurations.
  /// @param currentConfig The current configuration with metadata.
  /// @param newConfigWithMeta The new configuration with metadata.
  function _validateConfigTransition(
    OCR3ConfigWithMeta[] memory currentConfig,
    OCR3ConfigWithMeta[] memory newConfigWithMeta
  ) internal pure {
    uint256 currentConfigLen = currentConfig.length;
    uint256 newConfigLen = newConfigWithMeta.length;
    if (currentConfigLen == 0 && newConfigLen == 1) {
      // Config counts always must start at 1 for the first ever config.
      if (newConfigWithMeta[0].version != 1) {
        revert WrongConfigCount(newConfigWithMeta[0].version, 1);
      }
      return;
    }

    if (currentConfigLen == 1 && newConfigLen == 2) {
      // On a blue/green proposal:
      // * the config digest of the blue config must remain unchanged.
      // * the green config count must be the blue config count + 1.
      if (newConfigWithMeta[0].configDigest != currentConfig[0].configDigest) {
        revert WrongConfigDigestBlueGreen(newConfigWithMeta[0].configDigest, currentConfig[0].configDigest);
      }
      if (newConfigWithMeta[1].version != currentConfig[0].version + 1) {
        revert WrongConfigCount(newConfigWithMeta[1].version, currentConfig[0].version + 1);
      }
      return;
    }

    if (currentConfigLen == 2 && newConfigLen == 1) {
      // On a promotion, the green config digest must become the blue config digest.
      if (newConfigWithMeta[0].configDigest != currentConfig[1].configDigest) {
        revert WrongConfigDigest(newConfigWithMeta[0].configDigest, currentConfig[1].configDigest);
      }
      return;
    }

    revert NonExistentConfigTransition();
  }

  /// @notice Computes a new configuration with metadata based on the current configuration and the new configuration.
  /// @param donId The DON ID.
  /// @param currentConfig The current configuration, including metadata.
  /// @param newConfig The new configuration, without metadata.
  /// @param currentState The current state of the configuration.
  /// @param newState The new state of the configuration.
  /// @return The new configuration with metadata.
  function _computeNewConfigWithMeta(
    uint32 donId,
    OCR3ConfigWithMeta[] memory currentConfig,
    OCR3Config[] memory newConfig,
    ConfigState currentState,
    ConfigState newState
  ) internal view returns (OCR3ConfigWithMeta[] memory) {
    uint64[] memory configCounts = new uint64[](newConfig.length);

    // Set config counts based on the only valid state transitions.
    // Init    -> Running (first ever config)
    // Running -> Staging (blue/green proposal)
    // Staging -> Running (promotion)
    if (currentState == ConfigState.Init && newState == ConfigState.Running) {
      // First ever config starts with config count == 1.
      configCounts[0] = 1;
    } else if (currentState == ConfigState.Running && newState == ConfigState.Staging) {
      // On a blue/green proposal, the config count of the green config is the blue config count + 1.
      configCounts[0] = currentConfig[0].version;
      configCounts[1] = currentConfig[0].version + 1;
    } else if (currentState == ConfigState.Staging && newState == ConfigState.Running) {
      // On a promotion, the config count of the green config becomes the blue config count.
      configCounts[0] = currentConfig[1].version;
    } else {
      revert InvalidConfigStateTransition(currentState, newState);
    }

    OCR3ConfigWithMeta[] memory newConfigWithMeta = new OCR3ConfigWithMeta[](newConfig.length);
    for (uint256 i = 0; i < configCounts.length; ++i) {
      _validateConfig(newConfig[i]);
      newConfigWithMeta[i] = OCR3ConfigWithMeta({
        config: newConfig[i],
        version: configCounts[i],
        configDigest: _computeConfigDigest(donId, configCounts[i], newConfig[i])
      });
    }

    return newConfigWithMeta;
  }

  /// @notice Validates an OCR3 configuration.
  /// @param cfg The OCR3 configuration.
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

  /// @notice Computes the digest of the provided configuration.
  /// @dev In traditional OCR config digest computation, block.chainid and address(this) are used
  /// in order to further domain separate the digest. We can't do that here since the digest will
  /// be used on remote chains; so we use the chain selector instead of block.chainid. The don ID
  /// replaces the address(this) in the traditional computation.
  /// @param donId The DON ID.
  /// @param version The version number of the config.
  /// @param ocr3Config The OCR3 configuration.
  /// @return The computed digest.
  function _computeConfigDigest(
    uint32 donId,
    uint64 version,
    OCR3Config memory ocr3Config
  ) internal pure returns (bytes32) {
    uint256 h = uint256(
      keccak256(
        abi.encode(
          ocr3Config.chainSelector,
          donId,
          ocr3Config.pluginType,
          ocr3Config.offrampAddress,
          version,
          ocr3Config.nodes,
          ocr3Config.FRoleDON,
          ocr3Config.offchainConfigVersion,
          ocr3Config.offchainConfig
        )
      )
    );

    return bytes32((CONFIG_DIGEST_PREFIX & CONFIG_DIGEST_PREFIX_MASK) | (h & ~CONFIG_DIGEST_PREFIX_MASK));
  }

  // ================================================================
  // │                    Chain Configuration                       │
  // ================================================================

  /// @notice Sets and/or removes chain configurations.
  /// Does not validate that fChain <= FRoleDON and relies on OCR3Configs to be changed in case fChain becomes larger than the FRoleDON value.
  /// @param chainSelectorRemoves The chain configurations to remove.
  /// @param chainConfigAdds The chain configurations to add.
  function applyChainConfigUpdates(
    uint64[] calldata chainSelectorRemoves,
    ChainConfigInfo[] calldata chainConfigAdds
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
