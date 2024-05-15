// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {MultiOCR3Abstract} from "./MultiOCR3Abstract.sol";

// TODO: consider splitting configs & verification logic off to auth library (if size is prohibitive)
/// @notice Onchain verification of reports from the offchain reporting protocol
///         with multiple OCR plugin support.
abstract contract MultiOCR3Base is OwnerIsCreator, MultiOCR3Abstract {
  //   error InvalidConfig(string message);
  error WrongMessageLength(uint256 expected, uint256 actual);
  error ConfigDigestMismatch(bytes32 expected, bytes32 actual);
  error ForkedChain(uint256 expected, uint256 actual);
  error WrongNumberOfSignatures();
  error SignaturesOutOfRegistration();
  error UnauthorizedTransmitter();
  error UnauthorizedSigner();
  error NonUniqueSignatures();
  //   error OracleCannotBeZeroAddress();

  /// @dev Packing these fields used on the hot path in a ConfigInfo variable reduces the
  ///      retrieval of all of them to a minimum number of SLOADs.
  struct ConfigInfo {
    bytes32 latestConfigDigest;
    uint8 F; // ──────────────────────────────╮ maximum number of faulty/dishonest oracles
    uint8 n; //                               │ number of signers / transmitters
    bool uniqueReports; //                    │ if true, the reports should be unique
    bool isSignatureVerificationEnabled; // ──╯ if true, requires signers and verifies signatures on transmission verification
  }

  /// @notice Used for s_oracles[a].role, where a is an address, to track the purpose
  ///         of the address, or to indicate that the address is unset.
  enum Role {
    // No oracle role has been set for address a
    Unset,
    // Signing address for the s_oracles[a].index'th oracle. I.e., report
    // signatures from this oracle should ecrecover back to address a.
    Signer,
    // Transmission address for the s_oracles[a].index'th oracle. I.e., if a
    // report is received by OCR2Aggregator.transmit in which msg.sender is
    // a, it is attributed to the s_oracles[a].index'th oracle.
    Transmitter
  }

  struct Oracle {
    uint8 index; // Index of oracle in s_signers/s_transmitters
    Role role; // Role of the address which mapped to this struct
  }

  /// @notice OCR configuration for a single OCR plugin within a DON
  // TODO: make uniqueReports and skipReports static
  struct OCRConfig {
    /// @notice latest OCR config
    ConfigInfo configInfo;
    /// @notice signing address of each oracle
    address[] signers;
    /// @notice transmission address of each oracle,
    ///         i.e. the address the oracle actually sends transactions to the contract from
    address[] transmitters;
  }

  /// @notice Args to update an OCR Config
  struct OCRConfigArgs {
    uint8 ocrPluginType; // OCR plugin type to update config for
    bytes32 configDigest; // Config digest to update to
    uint8 F; // ───────────────────────────╮ maximum number of faulty/dishonest oracles
    bool uniqueReports; //                 │ if true, the reports should be unique
    bool isSignatureVerificationEnabled; // ──╯ if true, requires signers and verifies signatures on transmission verification
    address[] signers; // signing address of each oracle
    address[] transmitters; // transmission address of each oracle (i.e. the address the oracle actually sends transactions to the contract from)
  }

  /// @notice mapping of OCR plugin type -> DON config
  mapping(uint8 ocrPluginType => OCRConfig config) internal s_ocrConfigs;

  // TODO: optimization: we can use bitmaps of 2-bit width to optimise the role representation
  /// @notice OCR plugin type => signer OR transmitter address mapping
  mapping(uint8 ocrPluginType => mapping(address signerOrTransmiter => Oracle oracle)) s_oracles;

  // TODO: verify if the same TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT would apply to all plugins
  // The constant-length components of the msg.data sent to transmit.
  // See the "If we wanted to call sam" example on for example reasoning
  // https://solidity.readthedocs.io/en/v0.7.2/abi-spec.html
  uint16 private constant TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT = 4 // function selector
    + 32 * 3 // 3 words containing reportContext
    + 32 // word containing start location of abiencoded report value
    + 32 // word containing location start of abiencoded rs value
    + 32 // word containing start location of abiencoded ss value
    + 32 // rawVs value
    + 32 // word containing length of report
    + 32 // word containing length rs
    + 32; // word containing length of ss

  uint256 internal immutable i_chainID;

  // TODO: implement config sets in constructor
  constructor() {
    i_chainID = block.chainid;
  }

  /// @notice sets offchain reporting protocol configuration incl. participating oracles
  /// @param ocrConfigArgs OCR config update args
  function setOCR3Configs(OCRConfigArgs[] memory ocrConfigArgs) external onlyOwner {
    for (uint256 i; i < ocrConfigArgs.length; ++i) {
      _setOCR3Config(ocrConfigArgs[i]);
    }
  }

  /// @notice sets offchain reporting protocol configuration incl. participating oracles for a single OCR plugin type
  /// @param ocrConfigArgs OCR config update args
  function _setOCR3Config(OCRConfigArgs memory ocrConfigArgs) internal {
    uint8 ocrPluginType = ocrConfigArgs.ocrPluginType;
    OCRConfig storage ocrConfig = s_ocrConfigs[ocrPluginType];
    ConfigInfo storage configInfo = ocrConfig.configInfo;

    uint256 newTransmittersLength = ocrConfigArgs.transmitters.length;
    if (configInfo.isSignatureVerificationEnabled) {
      ocrConfig.signers = ocrConfigArgs.signers;
    }
    // TODO: validate signers / transmitters <= MAX_NUM_ORACLES

    // TODO: re-add s_oracles removal logic & validations
    //     uint256 oldSignerLength = s_signers.length;
    //     for (uint256 i = 0; i < oldSignerLength; ++i) {
    //       delete s_oracles[s_signers[i]];
    //       delete s_oracles[s_transmitters[i]];
    //     }

    //     uint256 newSignersLength = signers.length;
    //     for (uint256 i = 0; i < newSignersLength; ++i) {
    //       // add new signer/transmitter addresses
    //       address signer = signers[i];
    //       if (s_oracles[signer].role != Role.Unset) revert InvalidConfig("repeated signer address");
    //       if (signer == address(0)) revert OracleCannotBeZeroAddress();
    //       s_oracles[signer] = Oracle(uint8(i), Role.Signer);

    //       address transmitter = transmitters[i];
    //       if (s_oracles[transmitter].role != Role.Unset) revert InvalidConfig("repeated transmitter address");
    //       if (transmitter == address(0)) revert OracleCannotBeZeroAddress();
    //       s_oracles[transmitter] = Oracle(uint8(i), Role.Transmitter);
    //     }

    ocrConfig.transmitters = ocrConfigArgs.transmitters;
    configInfo.F = ocrConfigArgs.F;
    configInfo.latestConfigDigest = ocrConfigArgs.configDigest;
    configInfo.n = uint8(newTransmittersLength);

    emit ConfigSet(
      ocrPluginType, ocrConfigArgs.configDigest, ocrConfigArgs.signers, ocrConfigArgs.transmitters, ocrConfigArgs.F
    );
  }

  /// @param ocrPluginType OCR plugin type to retrieve transmitters for
  /// @return list of addresses permitted to transmit reports to this contract
  /// @dev The list will match the order used to specify the transmitter during setConfig
  function getTransmitters(uint8 ocrPluginType) external view returns (address[] memory) {
    return s_ocrConfigs[ocrPluginType].transmitters;
  }

  /// @inheritdoc MultiOCR3Abstract
  function _transmit(
    uint8 ocrPluginType,
    // NOTE: If these parameters are changed, expectedMsgDataLength and/or
    // TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT need to be changed accordingly
    bytes32[3] calldata reportContext,
    bytes calldata report,
    bytes32[] calldata rs,
    bytes32[] calldata ss,
    bytes32 rawVs // signatures
  ) internal override {
    // reportContext consists of:
    // reportContext[0]: ConfigDigest
    // reportContext[1]: 27 byte padding, 4-byte epoch and 1-byte round
    // reportContext[2]: ExtraHash
    bytes32 configDigest = reportContext[0];
    ConfigInfo memory configInfo = s_ocrConfigs[ocrPluginType].configInfo;

    if (configInfo.latestConfigDigest != configDigest) {
      revert ConfigDigestMismatch(configInfo.latestConfigDigest, configDigest);
    }
    // If the cached chainID at time of deployment doesn't match the current chainID, we reject all signed reports.
    // This avoids a (rare) scenario where chain A forks into chain A and A', A' still has configDigest
    // calculated from chain A and so OCR reports will be valid on both forks.
    if (i_chainID != block.chainid) revert ForkedChain(i_chainID, block.chainid);

    // Scoping this reduces stack pressure and gas usage
    {
      Oracle memory transmitter = s_oracles[ocrPluginType][msg.sender];
      // Check that sender is authorized to report
      if (
        !(
          transmitter.role == Role.Transmitter
            && msg.sender == s_ocrConfigs[ocrPluginType].transmitters[transmitter.index]
        )
      ) {
        revert UnauthorizedTransmitter();
      }
    }
    // Scoping this reduces stack pressure and gas usage
    {
      uint256 expectedDataLength = uint256(TRANSMIT_MSGDATA_CONSTANT_LENGTH_COMPONENT) + report.length // one byte pure entry in _report
        + rs.length * 32 // 32 bytes per entry in _rs
        + ss.length * 32; // 32 bytes per entry in _ss)
      if (msg.data.length != expectedDataLength) revert WrongMessageLength(expectedDataLength, msg.data.length);
    }

    emit Transmitted(ocrPluginType, configDigest, uint32(uint256(reportContext[1]) >> 8));

    if (configInfo.isSignatureVerificationEnabled) {
      // TODO: consider scoping this to reduce stack pressure / move to separate internal function
      uint256 expectedNumSignatures;
      if (configInfo.uniqueReports) {
        expectedNumSignatures = (configInfo.n + configInfo.F) / 2 + 1;
      } else {
        expectedNumSignatures = configInfo.F + 1;
      }
      if (rs.length != expectedNumSignatures) revert WrongNumberOfSignatures();
      if (rs.length != ss.length) revert SignaturesOutOfRegistration();

      // Verify signatures attached to report
      bytes32 h = keccak256(abi.encodePacked(keccak256(report), reportContext));
      bool[MAX_NUM_ORACLES] memory signed;

      uint256 numberOfSignatures = rs.length;
      for (uint256 i; i < numberOfSignatures; ++i) {
        // Safe from ECDSA malleability here since we check for duplicate signers.
        address signer = ecrecover(h, uint8(rawVs[i]) + 27, rs[i], ss[i]);
        // Since we disallow address(0) as a valid signer address, it can
        // never have a signer role.
        Oracle memory oracle = s_oracles[ocrPluginType][signer];
        if (oracle.role != Role.Signer) revert UnauthorizedSigner();
        if (signed[oracle.index]) revert NonUniqueSignatures();
        signed[oracle.index] = true;
      }
    }
  }

  /// @inheritdoc MultiOCR3Abstract
  function latestConfigDigestAndEpoch(uint8 ocrPluginType)
    external
    view
    virtual
    override
    returns (bool scanLogs, bytes32 configDigest, uint64 sequenceNumber)
  {
    return (true, bytes32(0), uint64(0));
  }

  /// @notice information about current offchain reporting protocol configuration
  /// @param ocrPluginType OCR plugin type to return config details for
  /// @return ocrConfig OCR config for the plugin type
  function latestConfigDetails(uint8 ocrPluginType) external view returns (OCRConfig memory ocrConfig) {
    return s_ocrConfigs[ocrPluginType];
  }
}
