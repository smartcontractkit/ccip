pragma solidity ^0.8.24;

import {ICapabilityConfiguration} from "./interfaces/ICapabilityConfiguration.sol";
import {ICapabilityRegistry} from "./interfaces/ICapabilityRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

/// @notice CCIPCapabilityConfiguration stores the configuration for the CCIP capability.
/// We have two classes of configuration: chain configuration and DON (in the CapabilityRegistry sense) configuration.
/// Each chain will have a single configuration which includes information like the router address.
/// Each CR DON will have up to four configurations: for each of (commit, exec), one blue and one green configuration.
/// This is done in order to achieve "blue-green" deployments.
contract CCIPCapabilityConfiguration is ICapabilityConfiguration, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.UintSet;

  /// @notice Emitted when a chain's configuration is set.
  /// @param chainSelector The chain selector.
  /// @param chainConfig The chain configuration.
  event ChainConfigSet(uint64 chainSelector, ChainConfig chainConfig);

  /// @notice Emitted when a chain's configuration is removed.
  /// @param chainSelector The chain selector.
  event ChainConfigRemoved(uint64 chainSelector);

  error InvalidConfigOperation();
  error NoCapabilityConfigurationSet(uint32 donId);
  error NodeNotInRegistry(bytes32 p2pId);
  error OnlyCapabilityRegistryCanCall();
  error ChainSelectorNotFound(uint64 chainSelector);
  error ChainSelectorNotSet();
  error SignerP2PIdPairMustBeLengthTwo(uint256 gotLength);

  /// @notice PluginId indicates the type of plugin that the configuration is for.
  /// @param Commit The configuration is for the commit plugin.
  /// @param Execution The configuration is for the execution plugin.
  enum PluginId {
    Commit,
    Execution
  }

  /// @notice Chain configuration.
  /// @param readers The P2P IDs of the readers for the chain. These IDs must be registered in the capability registry.
  /// @param config The chain configuration. This is kept intentionally opaque so as to add fields in the future if needed.
  /// Changes to chain configs are detected out-of-band offchain and decoded offchain.
  struct ChainConfig {
    bytes32[] readers;
    bytes config;
  }

  /// @notice Chain configuration update struct used in applyChainConfigUpdates.
  struct ChainConfigUpdate {
    uint64 chainSelector;
    ChainConfig chainConfig;
  }

  /// @notice OCR3 configuration.
  /// @param pluginId The plugin that the configuration is for.
  /// @param chainSelector The (remote) chain that the configuration is for.
  /// @param signers An associative array that contains (onchain signer public key, p2p id) pairs.
  /// @param transmitters An associative array that contains (transmitter, p2p id) pairs.
  /// @param f The "big F" parameter for the role DON.
  /// @param onchainConfig The onchain configuration for the OCR3 protocol. This will likely go unused.
  /// @param offchainConfigVersion The version of the offchain configuration.
  /// @param offchainConfig The offchain configuration for the OCR3 protocol. Protobuf encoded.
  struct OCR3Config {
    PluginId pluginId;
    uint64 chainSelector;
    bytes[][] signers;
    bytes[][] transmitters;
    uint8 f;
    bytes onchainConfig;
    uint64 offchainConfigVersion;
    bytes offchainConfig;
  }

  /// @notice The canonical capability registry address.
  address internal immutable i_capabilityRegistry;

  /// @notice chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => ChainConfig chainConfig) internal s_chainConfigurations;

  /// @notice All chains that are configured.
  EnumerableSet.UintSet internal s_chainSelectors;

  /// @notice OCR3 configurations for each DON.
  /// Each CR DON will have a commit and execution configuration.
  /// This means that a DON can have up to 4 configurations, since we are implementing blue/green deployments.
  mapping(uint32 donId => mapping(PluginId pluginId => OCR3Config[] ocr3Configs)) internal s_ocr3Configs;

  /// @notice The DONs that have been configured.
  EnumerableSet.UintSet internal s_donIds;

  /// @notice Configuration count domain separation parameter.
  /// Incremented for each configuration change across all DONs.
  uint64 internal s_configCount;

  /// @param capabilityRegistry the canonical capability registry address.
  constructor(address capabilityRegistry) {
    i_capabilityRegistry = capabilityRegistry;
  }

  // ================================================================
  // │                    Config Getters                            │
  // ================================================================

  /// @notice Returns the latest OCR3 configurations for all DONs.
  /// The offchain code will call this function to get the latest OCR3 configurations
  /// and spin up the appropriate OCR3 instances.
  /// This is expected to be called off-chain only and can have prohibitively high gas costs.
  /// @return The latest OCR3 configurations for all DONs.
  // TODO: will this eventually hit the RPC max response size limit?
  function getAllOCRConfigs() external view returns (OCR3Config[][] memory) {
    OCR3Config[][] memory ocr3Configs = new OCR3Config[][](s_donIds.length());
    for (uint256 i = 0; i < s_donIds.length(); i++) {
      // This is a safe cast, don IDs are uint32 in the CR.
      uint32 donId = uint32(s_donIds.at(i));
      ocr3Configs[i] = new OCR3Config[](s_ocr3Configs[donId][PluginId.Commit].length + s_ocr3Configs[donId][PluginId.Execution].length);
      for (uint256 j = 0; j < s_ocr3Configs[donId][PluginId.Commit].length; j++) {
        ocr3Configs[i][j] = s_ocr3Configs[donId][PluginId.Commit][j];
      }
      for (uint256 j = 0; j < s_ocr3Configs[donId][PluginId.Execution].length; j++) {
        ocr3Configs[i][j + s_ocr3Configs[donId][PluginId.Commit].length] = s_ocr3Configs[donId][PluginId.Execution][j];
      }
    }
    return ocr3Configs;
  }

  /// @notice Returns all the chain configurations.
  /// @return The chain configurations.
  // TODO: will this eventually hit the RPC max response size limit?
  function getAllChainConfigs() external view returns (ChainConfig[] memory) {
    ChainConfig[] memory chainConfigs = new ChainConfig[](s_chainSelectors.length());
    for (uint256 i = 0; i < s_chainSelectors.length(); i++) {
      chainConfigs[i] = s_chainConfigurations[uint64(s_chainSelectors.at(i))];
    }
    return chainConfigs;
  }

  // ================================================================
  // │                    Capability Configuration                  │
  // ================================================================

  /// @inheritdoc ICapabilityConfiguration
  function getCapabilityConfiguration(uint256 /* donId */) external pure override returns (bytes memory configuration) {
    // The CCIP capability will fetch the configuration needed directly from this contract.
    // The offchain syncer will call this function, however, so its important that it doesn't revert.
    return bytes("");
  }

  /// @notice Called by the registry prior to the config being set for a particular DON.
  function beforeCapabilityConfigSet(bytes32[] calldata nodes, bytes calldata config, uint64 configCount, uint32 donId) external override {
    if (msg.sender != i_capabilityRegistry) {
      revert OnlyCapabilityRegistryCanCall();
    }

    revert("unimplemented");
  }

  function _validateConfig(OCR3Config memory ocr3Config) internal view {
    if (ocr3Config.chainSelector == 0) {
      revert ChainSelectorNotSet();
    }

    // Check that the chain configuration is set.
    if (!s_chainSelectors.contains(ocr3Config.chainSelector)) {
      revert ChainSelectorNotFound(ocr3Config.chainSelector);
    }

    // Check that the readers are in the capability registry.
    for (uint256 j = 0; j < ocr3Config.signers.length; j++) {
      // We expect a pair of (p2pId, signer) for each element in the signers array.
      bytes[] memory signerP2PIdPair = ocr3Config.signers[j];
      if (signerP2PIdPair.length != 2) {
        revert SignerP2PIdPairMustBeLengthTwo(signerP2PIdPair.length);
      }

      // The provided p2pId must be in the capability registry.
      _ensureInRegistry(abi.decode(signerP2PIdPair[0], (bytes32)));
    }
  }

  // ================================================================
  // │                    Chain Configuration                       │
  // ================================================================

  /// @notice Sets and/or removes chain configurations.
  /// @param removes The chain configurations to remove.
  /// @param adds The chain configurations to add.
  function applyChainConfigUpdates(ChainConfigUpdate[] calldata removes, ChainConfigUpdate[] calldata adds) external onlyOwner {
    // Process removals first.
    for (uint256 i = 0; i < removes.length; i++) {
      // check if the chain selector is in s_chainSelectors first.
      if (!s_chainSelectors.contains(removes[i].chainSelector)) {
        revert ChainSelectorNotFound(removes[i].chainSelector);
      }

      delete s_chainConfigurations[removes[i].chainSelector];
      s_chainSelectors.remove(removes[i].chainSelector);

      emit ChainConfigRemoved(removes[i].chainSelector);
    }

    // Process additions next.
    for (uint256 i = 0; i < adds.length; i++) {
      bytes32[] memory readers = adds[i].chainConfig.readers;
      uint64 chainSelector = adds[i].chainSelector;

      // Verify that the provided readers are present in the capability registry.
      for (uint256 j = 0; i < readers.length; j++) {
        _ensureInRegistry(readers[j]);
      }

      s_chainConfigurations[chainSelector] = adds[i].chainConfig;
      s_chainSelectors.add(chainSelector);

      emit ChainConfigSet(chainSelector, adds[i].chainConfig);
    }
  }

  function _ensureInRegistry(bytes32 p2pId) internal view {
    (ICapabilityRegistry.NodeParams memory node, ) = ICapabilityRegistry(i_capabilityRegistry).getNode(p2pId);
    if (node.p2pId != p2pId) {
      revert NodeNotInRegistry(p2pId);
    }
  }
}
