// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilitiesRegistry} from "./interfaces/ICapabilitiesRegistry.sol";

import {Internal} from "../libraries/Internal.sol";
import {HomeBase} from "./HomeBase.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice CCIPHome stores the configuration for the CCIP capability.
/// We have two classes of configuration: chain configuration and DON (in the CapabilitiesRegistry sense) configuration.
/// Each chain will have a single configuration which includes information like the router address.
/// Each CR DON will have up to four configurations: for each of (commit, exec), one blue and one green configuration.
/// This is done in order to achieve "blue-green" deployments.
contract CCIPHome is HomeBase {
  using EnumerableSet for EnumerableSet.UintSet;

  event ChainConfigRemoved(uint64 chainSelector);
  event ChainConfigSet(uint64 chainSelector, ChainConfig chainConfig);

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
    OCR3Config staticConfig;
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

  uint256 private constant PREFIX = 0x000a << (256 - 16); // 0x000a00..00
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev 256 is the hard limit due to the bit encoding of their indexes into a uint256.
  uint256 internal constant MAX_NUM_ORACLES = 256;

  /// @dev chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => ChainConfig chainConfig) private s_chainConfigurations;

  /// @dev All chains that are configured.
  EnumerableSet.UintSet private s_remoteChainSelectors;

  constructor(address capabilitiesRegistry) HomeBase(capabilitiesRegistry) {}

  /// @notice Returns the total number of chains configured.
  /// @return The total number of chains configured.
  function getNumChainConfigurations() external view returns (uint256) {
    return s_remoteChainSelectors.length();
  }

  /// @notice The offchain code can use this to fetch an old config which might still be in use by some remotes. Use
  /// in case one of the configs is too large to be returnable by one of the other getters.
  /// @param configDigest The digest of the config to fetch.
  /// @return versionedConfig The config and its version.
  /// @return ok True if the config was found, false otherwise.
  function getConfig(
    bytes32 pluginKey,
    bytes32 configDigest
  ) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    (StoredConfig memory storedConfig, bool configOK) = _getStoredConfig(pluginKey, configDigest);
    if (configOK) {
      return (
        VersionedConfig({
          version: storedConfig.version,
          configDigest: storedConfig.configDigest,
          staticConfig: abi.decode(storedConfig.staticConfig, (OCR3Config))
        }),
        true
      );
    }

    return (versionedConfig, false);
  }

  function getAllConfigs(
    bytes32 pluginKey
  ) external view returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig) {
    (StoredConfig memory primaryStoredConfig, bool primaryOk) = _getPrimaryStoredConfig(pluginKey);

    if (primaryOk) {
      primaryConfig = VersionedConfig({
        version: primaryStoredConfig.version,
        configDigest: primaryStoredConfig.configDigest,
        staticConfig: abi.decode(primaryStoredConfig.staticConfig, (OCR3Config))
      });
    }

    (StoredConfig memory secondaryStoredConfig, bool secondaryOk) = _getSecondaryStoredConfig(pluginKey);

    if (secondaryOk) {
      secondaryConfig = VersionedConfig({
        version: secondaryStoredConfig.version,
        configDigest: secondaryStoredConfig.configDigest,
        staticConfig: abi.decode(secondaryStoredConfig.staticConfig, (OCR3Config))
      });
    }

    return (primaryConfig, secondaryConfig);
  }

  function _validateStaticAndDynamicConfig(bytes memory encodedStaticConfig, bytes memory) internal view override {
    OCR3Config memory cfg = abi.decode(encodedStaticConfig, (OCR3Config));

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

  function _validateDynamicConfig(
    bytes memory encodedStaticConfig,
    bytes memory encodedDynamicConfig
  ) internal pure override {
    // OCR doesn't use dynamic config
  }

  function _getConfigDigestPrefix() internal pure override returns (uint256) {
    return PREFIX;
  }

  // ================================================================
  // │                    Chain Configuration                       │
  // ================================================================

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
