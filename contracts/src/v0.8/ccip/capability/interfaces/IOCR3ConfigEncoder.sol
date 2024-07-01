// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/// @dev This is so that we can generate gethwrappers and easily encode/decode OCR3Config
/// in the offchain integration tests.
interface IOCR3ConfigEncoder {
  /// @notice PluginType indicates the type of plugin that the configuration is for.
  /// @param Commit The configuration is for the commit plugin.
  /// @param Execution The configuration is for the execution plugin.
  enum PluginType {
    Commit,
    Execution
  }

  /// @notice OCR3 configuration.
  struct OCR3Config {
    PluginType pluginType; // ────────╮ The plugin that the configuration is for.
    uint64 chainSelector; //          | The (remote) chain that the configuration is for.
    uint8 F; //                       | The "big F" parameter for the role DON.
    uint64 offchainConfigVersion; // ─╯ The version of the offchain configuration.
    bytes offrampAddress; // The remote chain offramp address.
    bytes32[] bootstrapP2PIds; // The bootstrap P2P IDs of the oracles that are part of the role DON.
    // len(p2pIds) == len(signers) == len(transmitters) == 3 * F + 1
    // NOTE: indexes matter here! The p2p ID at index i corresponds to the signer at index i and the transmitter at index i.
    // This is crucial in order to build the oracle ID <-> peer ID mapping offchain.
    bytes32[] p2pIds; // The P2P IDs of the oracles that are part of the role DON.
    bytes[] signers; // The onchain signing keys of nodes in the don.
    bytes[] transmitters; // The onchain transmitter keys of nodes in the don.
    bytes offchainConfig; // The offchain configuration for the OCR3 protocol. Protobuf encoded.
  }

  /// @notice OCR3 configuration with metadata, specifically the config count and the config digest.
  struct OCR3ConfigWithMeta {
    OCR3Config config; // The OCR3 configuration.
    uint64 configCount; // The config count used to compute the config digest.
    bytes32 configDigest; // The config digest of the OCR3 configuration.
  }

  /// @dev Encodes an array of OCR3Config into a bytes array. For test usage only.
  function exposeOCR3Config(OCR3Config[] calldata config) external view returns (bytes memory);
}
