// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

interface ICapabilityConfiguration {
  /// @notice Returns the capability configuration for a particular DON instance.
  /// @dev donId is required to get DON-specific configuration. It avoids a
  /// situation where configuration size grows too large.
  /// @param donId The DON instance ID. These are stored in the CapabilityRegistry.
  /// @return configuration DON's configuration for the capability.
  function getCapabilityConfiguration(uint32 donId) external view returns (bytes memory configuration);

  /// @notice Emitted when a capability configuration is set.
  event CapabilityConfigurationSet();
}

interface ICapabilityRegistry {
  struct Node {
    /// @notice The id of the node operator that manages this node
    uint256 nodeOperatorId;
    /// @notice This is an Ed25519 public key that is used to identify a node.
    /// This key is guaranteed to be unique in the CapabilityRegistry. It is
    /// used to identify a node in the the P2P network.
    bytes32 p2pId;
    /// @notice The signer address for application-layer message verification.
    address signer;
    /// @notice The list of hashed capability IDs this node supports. This list is
    /// never empty and all capabilities are guaranteed to exist in the
    /// CapabilityRegistry.
    bytes32[] supportedHashedCapabilityIds;
  }

  /// @notice Gets a node's data
  /// @param p2pId The P2P ID of the node to query for
  /// @return Node The node data
  function getNode(bytes32 p2pId) external view returns (Node memory);
}

/// @notice CCIPCapabilityConfiguration stores the configuration for the CCIP capability.
/// We have two classes of configuration: chain configuration and DON (in the CapabilityRegistry sense) configuration.
/// Each chain will have a single configuration which includes information like the router address.
/// Each CR DON will have up to two configurations: one for the staging protocol instance and one for the production protocol instance.
/// This is done in order to achieve "blue-green" deployments.
contract CCIPCapabilityConfiguration is ICapabilityConfiguration, OwnerIsCreator {
  using EnumerableSet for EnumerableSet.UintSet;
  using EnumerableSet for EnumerableSet.Bytes32Set;

  /// @notice Various operations that can be performed when setting a configuration.
  /// @param SetProductionConfig Set the configuration of the production OCR3 instance.
  /// @param SetStagingConfig Set the configuration of the staging OCR3 instance.
  /// @param PromoteStagingConfig Atomically promote the staging configuration to production - this effectively "moves" the staging
  /// configuration to be the production configuration and wipes out the staging configuration.
  enum ConfigOperation {
    SetProductionConfig,
    SetStagingConfig,
    PromoteStagingConfig
  }

  /// @notice Chain configuration.
  /// @param readers The P2P IDs of the readers for the chain. These IDs must be registered in the capability registry.
  /// @param config The chain configuration. This is kept intentionally opaque so as to add fields in the future if needed.
  /// It is decoded offchain by the CCIP capability.
  struct ChainConfig {
    bytes32[] readers;
    bytes config;
  }

  struct ChainConfigUpdate {
    uint64 chainSelector;
    ChainConfig chainConfig;
  }

  /// @notice OCR3 configuration.
  /// @param signers The onchain signer public keys for the OCR3 protocol (e.g, addresses if the destination is EVM compatible).
  /// @param transmitters The onchain transmitter public keys for the OCR3 protocol (e.g addresses if the destination is EVM compatible).
  /// @param f The "big F" parameter for the role DON.
  /// @param onchainConfig The onchain configuration for the OCR3 protocol. This will likely go unused.
  /// @param offchainConfigVersion The version of the offchain configuration.
  /// @param offchainConfig The offchain configuration for the OCR3 protocol. Protobuf encoded.
  struct OCR3Config {
    bytes[] signers;
    bytes[] transmitters;
    uint8 f;
    bytes onchainConfig;
    uint64 offchainConfigVersion;
    bytes offchainConfig;
  }

  /// @notice Configuration of a particular OCR instance.
  /// @param configDigest The config digest of the configuration. Used as a domain separator when signing reports.
  /// @param ocr3Config The OCR3 configuration.
  struct InstanceConfig {
    bytes32 configDigest;
    OCR3Config ocr3Config;
  }

  /// @notice Configuration for a blue-green deployment.
  /// Each DON will have two configurations: one for the production instance and one for the staging instance.
  /// The production instance is the "blue" instance and the staging instance is the "green" instance.
  /// @param productionConfig The configuration for the production instance.
  /// @param stagingConfig The configuration for the staging instance.
  struct BlueGreenConfig {
    InstanceConfig productionConfig;
    InstanceConfig stagingConfig;
  }

  /// @notice A configuration update.
  /// @param op The operation to perform. See ConfigOperation.
  /// @param ocr3Config The OCR3 configuration. If the operation is PromoteStagingConfig, it isn't needed.
  struct ConfigUpdate {
    ConfigOperation op;
    OCR3Config ocr3Config;
  }

  /// @notice Full capability configuration, returned from getCapabilityConfiguration.
  /// @param chainConfigs The chain configurations for all chains that are configured.
  /// @param blueGreenConfig The instance configuration for the DON.
  struct FullCapabilityConfiguration {
    ChainConfig[] chainConfigs;
    BlueGreenConfig blueGreenConfig;
  }

  /// @notice chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => ChainConfig chainConfig) internal s_chainConfigurations;

  /// @notice All chains that are configured.
  EnumerableSet.UintSet internal s_chainSelectors;

  /// @notice The capability registry address.
  ICapabilityRegistry immutable internal i_capabilityRegistry;

  mapping(uint32 donId => BlueGreenConfig blueGreenConfig) s_blueGreenConfigs;

  error InvalidConfigOperation();
  error NoCapabilityConfigurationSet(uint32 donId);
  error NodeNotInRegistry(bytes32 p2pId);
  error OnlyCapabilityRegistryCanCall();
  error ChainSelectorNotFound(uint64 chainSelector);

  /// @notice Emitted when a chain's configuration is set.
  /// @param chainSelector The chain selector.
  /// @param chainConfig The chain configuration.
  event ChainConfigSet(uint64 chainSelector, ChainConfig chainConfig);

  /// @notice Emitted when a chain's configuration is removed.
  /// @param chainSelector The chain selector.
  event ChainConfigRemoved(uint64 chainSelector);

  /// @notice Emitted when a production or staging configuration is set.
  /// @param donId The DON ID.
  /// @param config The configuration that got set.
  event ProductionConfigSet(uint32 donId, InstanceConfig config);

  /// @notice Emitted when a staging configuration is set.
  /// @param donId The DON ID.
  /// @param config The configuration that got set.
  event StagingConfigSet(uint32 donId, InstanceConfig config);

  /// @notice Emitted when a staging configuration is promoted to production.
  /// @param donId The DON ID.
  /// @param stagingConfig The staging configuration that got promoted.
  /// @param oldProductionConfig The old production configuration.
  event StagingConfigPromoted(uint32 donId, InstanceConfig stagingConfig, InstanceConfig oldProductionConfig);

  constructor(address capabilityRegistry) {
    i_capabilityRegistry = ICapabilityRegistry(capabilityRegistry);
  }

  /// @inheritdoc ICapabilityConfiguration
  function getCapabilityConfiguration(uint32 donId) external view returns (bytes memory configuration) {
    BlueGreenConfig memory blueGreenConfig = s_blueGreenConfigs[donId];
    if (blueGreenConfig.productionConfig.configDigest == bytes32(uint256(0)) &&
      blueGreenConfig.stagingConfig.configDigest == bytes32(uint256(0))) {
      revert NoCapabilityConfigurationSet(donId);
    }

    // get all the chain configs for all the chains that are configured.
    ChainConfig[] memory chainConfigs = new ChainConfig[](s_chainSelectors.length());
    for (uint256 i = 0; i < s_chainSelectors.length(); i++) {
      chainConfigs[i] = s_chainConfigurations[uint64(s_chainSelectors.at(i))];
    }

    FullCapabilityConfiguration memory fullConfig = FullCapabilityConfiguration(chainConfigs, blueGreenConfig);
    return abi.encode(fullConfig);
  }

  /// @notice Called by the registry prior to the config being set for a particular DON.
  function beforeCapabilityConfigSet(bytes32[] calldata nodes, bytes calldata config, uint64 configCount, uint32 donId) external {
    if (msg.sender != address(i_capabilityRegistry)) {
      revert OnlyCapabilityRegistryCanCall();
    }

    // TODO: check that configCount is increasing?

    ConfigUpdate memory update = abi.decode(config, (ConfigUpdate));
    bytes32 configDigest = _validateConfig(update.ocr3Config, configCount);
    InstanceConfig memory instanceConfig = InstanceConfig(configDigest, update.ocr3Config);
    if (update.op == ConfigOperation.SetProductionConfig) {
      s_blueGreenConfigs[donId].productionConfig = instanceConfig;

      emit ProductionConfigSet(donId, instanceConfig);
    } else if (update.op == ConfigOperation.SetStagingConfig) {
      s_blueGreenConfigs[donId].stagingConfig = instanceConfig;

      emit StagingConfigSet(donId, instanceConfig);
    } else if (update.op == ConfigOperation.PromoteStagingConfig) {
      // Promote the staging instanceConfig to production.
      InstanceConfig memory stagingConfig = s_blueGreenConfigs[donId].stagingConfig;
      InstanceConfig memory oldProductionConfig = s_blueGreenConfigs[donId].productionConfig;
      s_blueGreenConfigs[donId].productionConfig = stagingConfig;

      // Clear the staging instanceConfig.
      InstanceConfig memory emptyConfig;
      s_blueGreenConfigs[donId].stagingConfig = emptyConfig;

      emit StagingConfigPromoted(donId, stagingConfig, oldProductionConfig);
    } else {
      revert InvalidConfigOperation();
    }

    emit CapabilityConfigurationSet();
  }

  function _validateConfig(OCR3Config memory config, uint64 configCount) internal pure returns (bytes32 configDigest) {
    // check that signers/transmitters are not repeated.
    // check that signers/transmitters are not zero.
    revert("unimplemented");
  }

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
        ICapabilityRegistry.Node memory node = i_capabilityRegistry.getNode(readers[j]);
        if (node.p2pId != readers[j]) {
          revert NodeNotInRegistry(readers[j]);
        }
      }

      s_chainConfigurations[chainSelector] = adds[i].chainConfig;
      s_chainSelectors.add(chainSelector);

      emit ChainConfigSet(chainSelector, adds[i].chainConfig);
    }
  }
}
