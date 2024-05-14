// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

/// @notice Abstract contract for onchain verification of reports from the offchain reporting protocol
///         with multiple OCR plugin support.
abstract contract MultiOCR3Abstract is ITypeAndVersion {
  // Maximum number of oracles the offchain reporting protocol is designed for
  // TODO: bump up to theoretical max if required
  uint256 internal constant MAX_NUM_ORACLES = 31;

  /// @notice triggers a new run of the offchain reporting protocol
  /// @param ocrPluginType OCR plugin type for which the config was set
  /// @param previousConfigBlockNumber block in which the previous config was set, to simplify historic analysis
  /// @param configDigest configDigest of this configuration
  /// @param signers ith element is address ith oracle uses to sign a report
  /// @param transmitters ith element is address ith oracle uses to transmit a report via the transmit method
  /// @param F maximum number of faulty/dishonest oracles the protocol can tolerate while still working correctly
  event ConfigSet(
    uint8 ocrPluginType,
    uint32 previousConfigBlockNumber,
    bytes32 configDigest,
    address[] signers,
    address[] transmitters,
    uint8 F
  );

  /// @notice sets offchain reporting protocol configuration incl. participating oracles
  /// @param ocrPluginType OCR plugin type to set the config for
  /// @param configDigest Config digest to assign to the DON OCR config
  /// @param signers addresses with which oracles sign the reports
  /// @param transmitters addresses oracles use to transmit the reports
  /// @param F number of faulty oracles the system can tolerate
  function setOCR3Config(
    uint8 ocrPluginType,
    bytes32 configDigest,
    address[] memory signers,
    address[] memory transmitters,
    uint8 F
  ) external virtual;

  /// @notice information about current offchain reporting protocol configuration
  /// @param ocrPluginType OCR plugin type to return config details for
  /// @return blockNumber block at which this config was set
  /// @return configDigest domain-separation tag for current config (see _configDigestFromConfigData)
  function latestConfigDetails(uint8 ocrPluginType)
    external
    view
    virtual
    returns (uint32 blockNumber, bytes32 configDigest);

  /// @notice optionally emitted to indicate the latest configDigest and sequence number
  /// for which a report was successfully transmitted. Alternatively, the contract may
  /// use latestConfigDigestAndEpoch with scanLogs set to false.
  event Transmitted(uint8 indexed ocrPluginType, bytes32 configDigest, uint64 sequenceNumber);

  /// @notice optionally returns the latest configDigest and sequence number for which
  /// a report was successfully transmitted. Alternatively, the contract may return
  /// scanLogs set to true and use Transmitted events to provide this information
  /// to offchain watchers.
  /// @param ocrPluginType OCR plugin type to fetch config digest & sequence number for
  /// @return scanLogs indicates whether to rely on the configDigest and sequence number
  /// returned or whether to scan logs for the Transmitted event instead.
  /// @return configDigest
  /// @return sequenceNumber
  function latestConfigDigestAndEpoch(uint8 ocrPluginType)
    external
    view
    virtual
    returns (bool scanLogs, bytes32 configDigest, uint64 sequenceNumber);

  /// @notice _transmit is called to post a new report to the contract.
  ///         The function should be called after the per-DON reporting logic is completed.
  /// @param ocrPluginType OCR plugin type to transmit report for
  /// @param report serialized report, which the signatures are signing.
  /// @param rs ith element is the R components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries
  /// @param ss ith element is the S components of the ith signature on report. Must have at most MAX_NUM_ORACLES entries
  /// @param rawVs ith element is the the V component of the ith signature
  function _transmit(
    uint8 ocrPluginType,
    // NOTE: If these parameters are changed, expectedMsgDataLength and/or
    // TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT need to be changed accordingly
    bytes32[3] calldata reportContext,
    bytes calldata report,
    bytes32[] calldata rs,
    bytes32[] calldata ss,
    bytes32 rawVs // signatures
  ) internal virtual;
}
