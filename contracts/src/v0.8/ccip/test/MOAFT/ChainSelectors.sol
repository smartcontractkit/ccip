// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.10;

library ChainSelectors {
  // Testnets
  string public constant SEPOLIA = "SEPOLIA";
  string public constant GNOSIS_TESTNET = "GNOSIS_TESTNET";
  string public constant BNB_TESTNET = "BNB_TESTNET";
  string public constant MODE_TESTNET = "MODE_TESTNET";
  string public constant OPT_SEPOLIA = "OPT_SEPOLIA";
  string public constant POLYGON_AMOY = "POLYGON_AMOY";
  string public constant ARB_SEPOLIA = "ARB_SEPOLIA";
  string public constant AVAX_FUJI = "AVAX_FUJI";
  string public constant BASE_SEPOLIA = "BASE_SEPOLIA";
  string public constant ZKSYNC_TESTNET = "ZKSYNC_TESTNET";

  // Mainnets
  string public constant BLAST = "BLAST";
  string public constant ETHEREUM = "ETHEREUM";
  string public constant GNOSIS = "GNOSIS";
  string public constant BNB = "BNB";
  string public constant MODE = "MODE";
  string public constant OPTIMISM = "OPTIMISM";
  string public constant POLYGON = "POLYGON";
  string public constant ARBITRUM = "ARBITRUM";
  string public constant AVAX = "AVAX";
  string public constant BASE = "BASE";
  string public constant METIS = "METIS";

  function _resolveChainSelector(
    uint64 chainSelector
  ) internal pure returns (string memory) {
    // Testnets
    if (chainSelector == 3478487238524512106) {
      return ARB_SEPOLIA;
    } else if (chainSelector == 8871595565390010547) {
      return GNOSIS_TESTNET;
    } else if (chainSelector == 16281711391670634445) {
      return POLYGON_AMOY;
    } else if (chainSelector == 13264668187771770619) {
      return BNB_TESTNET;
    } else if (chainSelector == 10344971235874465080) {
      return BASE_SEPOLIA;
    } else if (chainSelector == 829525985033418733) {
      return MODE_TESTNET;
    } else if (chainSelector == 5224473277236331295) {
      return OPT_SEPOLIA;
    } else if (chainSelector == 16015286601757825753) {
      return SEPOLIA;
    } else if (chainSelector == 6898391096552792247) {
      return ZKSYNC_TESTNET;
    } else if (chainSelector == 14767482510784806043) {
      return AVAX_FUJI;
    }
    // Mainnets
    if (chainSelector == 5009297550715157269) {
      return ETHEREUM;
    } else if (chainSelector == 4411394078118774322) {
      return BLAST;
    } else if (chainSelector == 5009297550715157269) {
      return "Ethereum mainnet";
    } else if (chainSelector == 465200170687744372) {
      return GNOSIS;
    } else if (chainSelector == 11344663589394136015) {
      return BNB;
    } else if (chainSelector == 7264351850409363825) {
      return MODE;
    } else if (chainSelector == 3734403246176062136) {
      return OPTIMISM;
    } else if (chainSelector == 4051577828743386545) {
      return POLYGON;
    } else if (chainSelector == 4949039107694359620) {
      return ARBITRUM;
    } else if (chainSelector == 6433500567565415381) {
      return AVAX;
    } else if (chainSelector == 15971525489660198786) {
      return BASE;
    } else if (chainSelector == 8805746078405598895) {
      return METIS;
    }

    return "Unknown";
  }
}
