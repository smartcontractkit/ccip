// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ICapabilitiesRegistry} from "./interfaces/ICapabilitiesRegistry.sol";

import {Internal} from "../libraries/Internal.sol";
import {HomeBase} from "./HomeBase.sol";
import {CCIPConfigTypes} from "./libraries/CCIPConfigTypes.sol";

import {EnumerableSet} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/structs/EnumerableSet.sol";

/// @notice Stores the home configuration for RMN, that is referenced by CCIP oracles, RMN nodes, and the RMNRemote
/// contracts.
contract RMNHome is HomeBase {
  using EnumerableSet for EnumerableSet.UintSet;

  event ChainConfigRemoved(uint64 chainSelector);
  /// @notice Emitted when a chain's configuration is set.
  /// @param chainSelector The chain selector.
  /// @param chainConfig The chain configuration.
  event ChainConfigSet(uint64 chainSelector, CCIPConfigTypes.ChainConfig chainConfig);

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
  error InvalidNode(CCIPConfigTypes.OCR3Node node);
  error NotEnoughTransmitters(uint256 got, uint256 minimum);

  struct Node {
    bytes32 peerId; //            Used for p2p communication.
    bytes32 offchainPublicKey; // Observations are signed with this public key, and are only verified offchain.
  }

  struct SourceChain {
    uint64 chainSelector; // ─────╮ The Source chain selector.
    uint64 minObservers; // ──────╯ Required number of observers to agree on an observation for this source chain.
    uint256 observerNodesBitmap; // ObserverNodesBitmap & (1<<i) == (1<<i) iff nodes[i] is an observer for this source chain.
  }

  struct StaticConfig {
    // No sorting requirement for nodes, but ensure that SourceChain.observerNodeIndices in the home chain config &
    // Signer.nodeIndex in the remote chain configs are appropriately updated when changing this field.
    Node[] nodes;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct DynamicConfig {
    // No sorting requirement for source chains, it is most gas efficient to append new source chains to the right.
    SourceChain[] sourceChains;
    bytes offchainConfig; // Offchain configuration for RMN nodes.
  }

  struct VersionedConfig {
    uint32 version;
    bytes32 configDigest;
    StaticConfig staticConfig;
    DynamicConfig dynamicConfig;
  }

  string public constant override typeAndVersion = "CCIPHome 1.6.0-dev";

  uint256 private constant PREFIX = 0x000a << (256 - 16); // 0x000a00..00
  bytes32 internal constant EMPTY_ENCODED_ADDRESS_HASH = keccak256(abi.encode(address(0)));
  /// @dev 256 is the hard limit due to the bit encoding of their indexes into a uint256.
  uint256 internal constant MAX_NUM_ORACLES = 256;

  /// @dev chain configuration for each chain that CCIP is deployed on.
  mapping(uint64 chainSelector => CCIPConfigTypes.ChainConfig chainConfig) private s_chainConfigurations;

  /// @dev All chains that are configured.
  EnumerableSet.UintSet private s_remoteChainSelectors;

  constructor() HomeBase(address(1)) {}

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
  function getConfig(bytes32 configDigest) external view returns (VersionedConfig memory versionedConfig, bool ok) {
    (StoredConfig memory storedConfig, bool configOK) = _getStoredConfig(0, 0, configDigest);
    if (configOK) {
      return (
        VersionedConfig({
          version: storedConfig.version,
          configDigest: storedConfig.configDigest,
          staticConfig: abi.decode(storedConfig.staticConfig, (StaticConfig)),
          dynamicConfig: abi.decode(storedConfig.dynamicConfig, (DynamicConfig))
        }),
        true
      );
    }

    return (versionedConfig, false);
  }

  function getAllConfigs()
    external
    view
    returns (VersionedConfig memory primaryConfig, VersionedConfig memory secondaryConfig)
  {
    (StoredConfig memory primaryStoredConfig, bool primaryOk) = _getPrimaryStoredConfig(0, 0);

    if (primaryOk) {
      primaryConfig = VersionedConfig({
        version: primaryStoredConfig.version,
        configDigest: primaryStoredConfig.configDigest,
        staticConfig: abi.decode(primaryStoredConfig.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(primaryStoredConfig.dynamicConfig, (DynamicConfig))
      });
    }

    (StoredConfig memory secondaryStoredConfig, bool secondaryOk) = _getSecondaryStoredConfig(0, 0);

    if (secondaryOk) {
      secondaryConfig = VersionedConfig({
        version: secondaryStoredConfig.version,
        configDigest: secondaryStoredConfig.configDigest,
        staticConfig: abi.decode(secondaryStoredConfig.staticConfig, (StaticConfig)),
        dynamicConfig: abi.decode(secondaryStoredConfig.dynamicConfig, (DynamicConfig))
      });
    }

    return (primaryConfig, secondaryConfig);
  }

  function _validateStaticAndDynamicConfig(bytes memory encodedStaticConfig, bytes memory) internal view override {
    CCIPConfigTypes.OCR3Config memory cfg = abi.decode(encodedStaticConfig, (CCIPConfigTypes.OCR3Config));

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
      CCIPConfigTypes.OCR3Node memory node = cfg.nodes[i];

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

  /// @notice Sets and/or removes chain configurations.
  /// Does not validate that fChain <= FRoleDON and relies on OCR3Configs to be changed in case fChain becomes larger than the FRoleDON value.
  /// @param chainSelectorRemoves The chain configurations to remove.
  /// @param chainConfigAdds The chain configurations to add.
  function applyChainConfigUpdates(
    uint64[] calldata chainSelectorRemoves,
    CCIPConfigTypes.ChainConfigInfo[] calldata chainConfigAdds
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
      CCIPConfigTypes.ChainConfig memory chainConfig = chainConfigAdds[i].chainConfig;
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
