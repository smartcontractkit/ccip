// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {IRMNV2} from "../interfaces/IRMNV2.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {Internal} from "../libraries/Internal.sol";

/// @dev this is included in the preimage of the digest that RMN nodes sign
bytes32 constant RMN_V1_6_ANY2EVM_REPORT = keccak256("RMN_V1_6_ANY2EVM_REPORT");

/// @dev An active curse on this subject will cause isCursed() to return true. Use this subject if there is an issue with a
/// remote chain, for which there exists a legacy lane contract deployed on the same chain as this RMN contract is
/// deployed, relying on isCursed().
bytes16 constant LEGACY_CURSE_SUBJECT = 0x01000000000000000000000000000000;

/// @dev An active curse on this subject will cause isCursed() and isCursed(bytes16) to return true. Use this subject for
/// issues affecting all of CCIP chains, or pertaining to the chain that this contract is deployed on, instead of using
/// the local chain selector as a subject.
bytes16 constant GLOBAL_CURSE_SUBJECT = 0x01000000000000000000000000000001;

/// @notice This contract supports verification of RMN reports for any Any2EVM OffRamp.
contract RMNRemote is OwnerIsCreator, ITypeAndVersion, IRMNV2 {
  error AlreadyCursed(bytes16 subject);
  error ConfigNotSet();
  error DuplicateOnchainPublicKey();
  error InvalidSignature();
  error InvalidSignerOrder();
  error MinSignersTooHigh();
  error NotCursed(bytes16 subject);
  error OutOfOrderSignatures();
  error ThresholdNotMet();
  error UnexpectedSigner();

  event ConfigSet(VersionedConfig versionedConfig);
  event Cursed(bytes16[] subjects);
  event Uncursed(bytes16[] subjects);

  /// @dev the configuration of an RMN signer
  struct Signer {
    address onchainPublicKey; // ────╮ For signing reports
    uint64 nodeIndex; // ────────────╯ Maps to nodes in home chain config, should be strictly increasing
  }

  /// @dev the contract config
  struct Config {
    bytes32 rmnHomeContractConfigDigest; // Digest of the RMNHome contract config
    Signer[] signers; // List of signers
    uint64 minSigners; // Threshold for the number of signers required to verify a report
  }

  /// @dev the contract config + a version number
  struct VersionedConfig {
    uint32 version; // For tracking the version of the config
    Config config; // The config
  }

  /// @dev part of the payload that RMN nodes sign: keccak256(abi.encode(RMN_V1_6_ANY2EVM_REPORT, report))
  struct Report {
    uint256 destChainId; // To guard against chain selector misconfiguration
    uint64 destChainSelector; // The chain selector of the destination chain
    address rmnRemoteContractAddress; // The address of this contract
    address offrampAddress; // The address of the offramp on the same chain as this contract
    bytes32 rmnHomeContractConfigDigest; // The digest of the RMNHome contract config
    Internal.MerkleRoot[] destLaneUpdates; // The dest lane updates
  }

  Config s_config;
  uint32 s_configCount;

  string public constant override typeAndVersion = "RMNRemote 1.6.0-dev";
  uint64 internal immutable i_localChainSelector;

  bytes16[] private s_cursedSubjectsSequence;
  /// @dev the index+1 is stored to easily distinguish b/t noncursed and cursed at the 0 index
  mapping(bytes16 subject => uint256 indexPlusOne) private s_cursedSubjectsIndexPlusOne;
  mapping(address signer => bool exists) s_signers; // for more gas efficient verify

  /// @param localChainSelector the chain selector of the chain this contract is deployed to
  constructor(uint64 localChainSelector) {
    i_localChainSelector = localChainSelector;
  }

  // ================================================================
  // │                         Verification                         │
  // ================================================================

  /// @inheritdoc IRMNV2
  function verify(
    address offrampAddress,
    Internal.MerkleRoot[] memory destLaneUpdates,
    Signature[] memory signatures
  ) external view {
    if (s_configCount == 0) {
      revert ConfigNotSet();
    }

    bytes32 signedHash = keccak256(
      abi.encode(
        RMN_V1_6_ANY2EVM_REPORT,
        Report({
          destChainId: block.chainid,
          destChainSelector: i_localChainSelector,
          rmnRemoteContractAddress: address(this),
          offrampAddress: offrampAddress,
          rmnHomeContractConfigDigest: s_config.rmnHomeContractConfigDigest,
          destLaneUpdates: destLaneUpdates
        })
      )
    );

    uint256 numSigners = 0;
    address prevAddress = address(0);
    for (uint256 i = 0; i < signatures.length; ++i) {
      Signature memory sig = signatures[i];
      address signerAddress = ecrecover(signedHash, 27, sig.r, sig.s);
      if (signerAddress == address(0)) revert InvalidSignature();
      if (!(prevAddress < signerAddress)) revert OutOfOrderSignatures();
      if (!s_signers[signerAddress]) revert UnexpectedSigner();
      prevAddress = signerAddress;
      ++numSigners;
    }
    if (numSigners < s_config.minSigners) revert ThresholdNotMet();
  }

  // ================================================================
  // │                            Config                            │
  // ================================================================

  /// @notice Sets the configuration of the contract
  /// @param newConfig the new configuration
  /// @dev setting congig is atomic; we delete all pre-existing config and set everything from scratch
  function setConfig(Config calldata newConfig) external onlyOwner {
    // sanity checks
    {
      // signers are in ascending order of nodeIndex
      for (uint256 i = 1; i < newConfig.signers.length; ++i) {
        if (!(newConfig.signers[i - 1].nodeIndex < newConfig.signers[i].nodeIndex)) {
          revert InvalidSignerOrder();
        }
      }

      // minSigners is tenable
      if (!(newConfig.minSigners <= newConfig.signers.length)) {
        revert MinSignersTooHigh();
      }
    }

    // clear the old signers
    {
      Config storage oldConfig = s_config;
      while (oldConfig.signers.length > 0) {
        delete s_signers[oldConfig.signers[oldConfig.signers.length - 1].onchainPublicKey];
        oldConfig.signers.pop();
      }
    }

    // set the new signers
    {
      for (uint256 i = 0; i < newConfig.signers.length; ++i) {
        if (s_signers[newConfig.signers[i].onchainPublicKey]) {
          revert DuplicateOnchainPublicKey();
        }
        s_signers[newConfig.signers[i].onchainPublicKey] = true;
      }
    }

    s_config = newConfig;
    uint32 newConfigCount = ++s_configCount;
    emit ConfigSet(VersionedConfig({version: newConfigCount, config: newConfig}));
  }

  /// @notice Returns the current configuration of the contract + a version number
  /// @return versionedConfig the current configuration + version
  function getVersionedConfig() external view returns (VersionedConfig memory) {
    return VersionedConfig({version: s_configCount, config: s_config});
  }

  /// @notice Returns the chain selector configured at deployment time
  /// @return localChainSelector the chain selector (not the chain ID)
  function getLocalChainSelector() external view returns (uint64 localChainSelector) {
    return i_localChainSelector;
  }

  // ================================================================
  // │                           Cursing                            │
  // ================================================================

  /// @notice Curse a single subject
  /// @param subject the subject to curse
  function curse(bytes16 subject) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    curse(subjects);
  }

  /// @notice Curse an array of subjects
  /// @param subjects the subjects to curse
  /// @dev reverts if any of the subjects are already cursed or if there is a duplicate
  function curse(bytes16[] memory subjects) public onlyOwner {
    for (uint256 i = 0; i < subjects.length; ++i) {
      bytes16 toCurseSubject = subjects[i];
      if (s_cursedSubjectsIndexPlusOne[toCurseSubject] != 0) {
        revert AlreadyCursed(toCurseSubject);
      }
      s_cursedSubjectsSequence.push(toCurseSubject);
      s_cursedSubjectsIndexPlusOne[toCurseSubject] = s_cursedSubjectsSequence.length;
    }
    emit Cursed(subjects);
  }

  /// @notice Uncurse a single subject
  /// @param subject the subject to uncurse
  function uncurse(bytes16 subject) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    uncurse(subjects);
  }

  /// @notice Uncurse an array of subjects
  /// @param subjects the subjects to uncurse
  /// @dev reverts if any of the subjects are not cursed or if there is a duplicate
  function uncurse(bytes16[] memory subjects) public onlyOwner {
    for (uint256 i = 0; i < subjects.length; ++i) {
      bytes16 toUncurseSubject = subjects[i];
      uint256 toUncurseSubjectIndexPlusOne = s_cursedSubjectsIndexPlusOne[toUncurseSubject];
      if (toUncurseSubjectIndexPlusOne == 0) {
        revert NotCursed(toUncurseSubject);
      }
      uint256 toUncurseSubjectIndex = toUncurseSubjectIndexPlusOne - 1;
      // copy the last subject to the position of the subject to uncurse
      bytes16 lastSubject = s_cursedSubjectsSequence[s_cursedSubjectsSequence.length - 1];
      s_cursedSubjectsSequence[toUncurseSubjectIndex] = lastSubject;
      s_cursedSubjectsIndexPlusOne[lastSubject] = toUncurseSubjectIndexPlusOne;
      // then pop, since we have the last subject also in toUncurseSubjectIndex
      s_cursedSubjectsSequence.pop();
      delete s_cursedSubjectsIndexPlusOne[toUncurseSubject];
    }
    emit Uncursed(subjects);
  }

  /// @inheritdoc IRMNV2
  function getCursedSubjects() external view returns (bytes16[] memory subjects) {
    return s_cursedSubjectsSequence;
  }

  /// @inheritdoc IRMNV2
  function isCursed() external view returns (bool) {
    if (s_cursedSubjectsSequence.length == 0) {
      return false;
    }
    return
      s_cursedSubjectsIndexPlusOne[LEGACY_CURSE_SUBJECT] > 0 || s_cursedSubjectsIndexPlusOne[GLOBAL_CURSE_SUBJECT] > 0;
  }

  /// @inheritdoc IRMNV2
  function isCursed(bytes16 subject) external view returns (bool) {
    if (s_cursedSubjectsSequence.length == 0) {
      return false;
    }
    return s_cursedSubjectsIndexPlusOne[subject] > 0 || s_cursedSubjectsIndexPlusOne[GLOBAL_CURSE_SUBJECT] > 0;
  }
}
