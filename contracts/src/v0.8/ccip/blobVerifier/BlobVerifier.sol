// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../vendor/Address.sol";
import "../../vendor/SafeERC20.sol";
import "../health/HealthChecker.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../interfaces/BlobVerifierInterface.sol";

contract BlobVerifier is BlobVerifierInterface, TypeAndVersionInterface, HealthChecker, OCR2Base {
  using Address for address;
  using SafeERC20 for IERC20;

  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;
  // Offchain leaf domain separator
  bytes1 private constant LEAF_DOMAIN_SEPARATOR = 0x00;
  // Internal domain separator used in proofs
  bytes1 private constant INTERNAL_DOMAIN_SEPARATOR = 0x01;

  // merkleRoot => timestamp when received
  mapping(bytes32 => uint256) private s_roots;

  // The BlobVerifier configuration values
  BlobVerifierConfig private s_config;

  // Mapping of the expected next sequence number by onRamp
  mapping(address => uint64) public s_expectedNextMinByOnRamp;

  /**
   * @dev sourceTokens are mapped to pools, and therefore should be the same length arrays.
   * @dev The AFN contract should be deployed already
   * @param chainId The ID that this contract is deployed to
   * @param afn AFN contract
   * @param maxTimeWithoutAFNSignal An AFN config setting
   * @param config containing:
   * sourceChainId: the source chain ID
   * onRamps: the addresses of the connected onRamps on the source chain for when overwriting
   *    the s_expectedNextMinByOnRamp mapping
   * minSeqNrByOnRamp: the new values when overwriting the s_expectedNextMinByOnRamp mapping
   */
  constructor(
    uint256 chainId,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    BlobVerifierConfig memory config
  ) OCR2Base(true) HealthChecker(afn, maxTimeWithoutAFNSignal) {
    CHAIN_ID = chainId;
    s_config = config;
    for (uint256 i = 0; i < s_config.onRamps.length; ++i) {
      s_expectedNextMinByOnRamp[s_config.onRamps[i]] = s_config.minSeqNrByOnRamp[i];
    }
  }

  /**
   * @notice Sets the new BlobVerifierConfig and updates the s_expectedNextMinByOnRamp
   * mapping. This will overwrite the existing expected minimum sequence numbers
   * for the supplied onRamps but will not change them for onRamps that are not
   * included in the BlobVerifierConfig.onRamps property.
   */
  function setConfig(BlobVerifierConfig calldata config) external onlyOwner {
    s_config = config;
    // TODO overwriting mappings is hard. Old values can persist with this implementation
    for (uint256 i = 0; i < s_config.onRamps.length; ++i) {
      s_expectedNextMinByOnRamp[s_config.onRamps[i]] = s_config.minSeqNrByOnRamp[i];
    }

    emit BlobVerifierConfigSet(config);
  }

  /**
   * @notice Returns the current config.
   */
  function getConfig() external view returns (BlobVerifierConfig memory) {
    return s_config;
  }

  /**
   * @notice Used by the owner in case an invalid sequence of roots has been
   * posted and needs to be removed. The interval in the report is trusted.
   */
  function resetUnblessedRoots(bytes32[] calldata rootToReset) external onlyOwner {
    for (uint256 i = 0; i < rootToReset.length; ++i) {
      // TODO: AFN check ( assert not self.afn.is_blessed(root))
      delete s_roots[rootToReset[i]];
    }
  }

  /**
   * @notice Extending OCR2Base._report
   * @dev assumes the report is a bytes encoded CCIP.RelayReport
   * @dev report is called by the Relaying DON through the ReportingPlugin
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory encodedReport
  ) internal override whenNotPaused whenHealthy {
    CCIP.RelayReport memory report = abi.decode(encodedReport, (CCIP.RelayReport));
    uint256 reportLength = report.onRamps.length;
    if (reportLength != report.intervals.length || reportLength != report.merkleRoots.length) {
      revert InvalidRelayReport(report);
    }
    for (uint256 i = 0; i < reportLength; ++i) {
      address onRamp = report.onRamps[i];
      uint64 expectedMinSeqNum = s_expectedNextMinByOnRamp[onRamp];
      CCIP.Interval memory repInterval = report.intervals[i];

      if (expectedMinSeqNum == 0) {
        revert UnSupportedOnRamp(onRamp);
      }
      if (expectedMinSeqNum != repInterval.min || repInterval.min > repInterval.max) {
        revert InvalidInterval(repInterval, onRamp);
      }
      s_expectedNextMinByOnRamp[onRamp] = repInterval.max + 1;
    }
    s_roots[report.rootOfRoots] = block.timestamp;
    emit ReportAccepted(report);
  }

  /**
   * @notice Returns timestamp of when root was accepted or -1 if verification fails
   */
  function verify(
    bytes32[] calldata hashedLeaves,
    bytes32[] calldata innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] calldata outerProofs,
    uint256 outerProofFlagBits
  ) external view returns (uint256 timestamp) {
    bytes32 innerRoot = merkleRoot(hashedLeaves, innerProofs, innerProofFlagBits);
    bytes32[] memory outerLeaves = new bytes32[](1);
    outerLeaves[0] = innerRoot;
    bytes32 outerRoot = merkleRoot(outerLeaves, outerProofs, outerProofFlagBits);
    return s_roots[outerRoot];
  }

  /**
   * @notice Generate a Merkle Root
   */
  function merkleRoot(
    bytes32[] memory leaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) public pure returns (bytes32) {
    uint256 leavesLen = leaves.length;
    uint256 totalHashes = leavesLen + proofs.length - 1;
    require(totalHashes <= 256);
    unchecked {
      bytes32[] memory hashes = new bytes32[](totalHashes);
      uint256 leafPos = 0;
      uint256 hashPos = 0;
      uint256 proofPos = 0;
      for (uint256 i = 0; i < totalHashes; ++i) {
        bool proofFlag = ((proofFlagBits >> i) & uint256(1)) == 1;
        hashes[i] = hashPair(
          proofFlag ? (leafPos < leavesLen ? leaves[leafPos++] : hashes[hashPos++]) : proofs[proofPos++],
          leafPos < leavesLen ? leaves[leafPos++] : hashes[hashPos++]
        );
      }

      if (totalHashes > 0) {
        return hashes[totalHashes - 1];
      }
      return leaves[0];
    }
  }

  function getMerkleRoot(bytes32 root) external view returns (uint256) {
    return s_roots[root];
  }

  function hashPair(bytes32 a, bytes32 b) private pure returns (bytes32) {
    return a < b ? _hashInternalNode(a, b) : _hashInternalNode(b, a);
  }

  function _hashInternalNode(bytes32 left, bytes32 right) private pure returns (bytes32 hash) {
    return keccak256(abi.encodePacked(INTERNAL_DOMAIN_SEPARATOR, left, right));
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {
    // TODO
  }

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "BlobVerifier 1.0.0";
  }
}
