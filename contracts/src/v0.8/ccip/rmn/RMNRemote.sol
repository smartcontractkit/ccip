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

  struct Signer {
    address onchainPublicKey; // ────╮ for signing reports
    uint64 nodeIndex; // ────────────╯ maps to nodes in home chain config, should be strictly increasing
  }

  struct Config {
    bytes32 rmnHomeContractConfigDigest; // digest of the RMNHome contract config
    Signer[] signers; // list of signers
    uint64 minSigners; // threshold for the number of signers required to verify a report
  }

  struct VersionedConfig {
    uint32 version; // for tracking the version of the config
    Config config; // the config
  }

  struct Report {
    uint256 destChainId; // to guard against chain selector misconfiguration
    uint64 destChainSelector; // the chain selector of the destination chain
    address rmnRemoteContractAddress; // the address of this contract
    address offrampAddress; // the address of the offramp on the same chain as this contract
    bytes32 rmnHomeContractConfigDigest; // the digest of the RMNHome contract config
    Internal.MerkleRoot[] destLaneUpdates; // the dest lane updates
  }

  // ================================================================
  // │                           Storage                            │
  // ================================================================

  Config s_config;
  uint32 s_configCount;

  string public constant override typeAndVersion = "RMNRemote 1.6.0-dev";
  uint64 internal immutable i_chainSelector;

  bytes16[] private s_cursedSubjectsSequence;
  mapping(bytes16 subject => uint256 indexPlusOne) private s_cursedSubjectsIndexPlusOne;
  mapping(address signer => bool exists) s_signers; // for more gas efficient verify

  // ================================================================
  // │                         Constructor                          │
  // ================================================================

  /// @param chainSelector the chain selector of the chain this contract is deployed to
  constructor(uint64 chainSelector) {
    i_chainSelector = chainSelector;
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
          destChainSelector: i_chainSelector,
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

  function getVersionedConfig() external view returns (VersionedConfig memory) {
    return VersionedConfig({version: s_configCount, config: s_config});
  }

  /// @notice Returns the chain selector configured at deployment time
  /// @return chainSelector the chain selector (not the chain ID)
  function getChainSelector() external view returns (uint64 chainSelector) {
    return i_chainSelector;
  }

  // ================================================================
  // │                           Cursing                            │
  // ================================================================

  function curse(bytes16 subject) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    curse(subjects);
  }

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

  function uncurse(bytes16 subject) external {
    bytes16[] memory subjects = new bytes16[](1);
    subjects[0] = subject;
    uncurse(subjects);
  }

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

  function getCursedSubjects() external view returns (bytes16[] memory) {
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
