// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IRMN} from "./interfaces/IRMN.sol";

import {EnumerableSet} from "../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/structs/EnumerableSet.sol";

import {GLOBAL_CURSE_SUBJECT, RMNBase} from "./RMNBase.sol";

/// @dev This contract is owned by RMN, if changing, please notify the RMN maintainers.
// solhint-disable chainlink-solidity/explicit-returns
contract RMN is IRMN, RMNBase {
  using EnumerableSet for EnumerableSet.AddressSet;

  string public constant override typeAndVersion = "RMN 1.5.0-dev";

  struct BlessVoteProgress {
    // This particular ordering saves us ~400 gas per voteToBless call, compared to the bool being at the bottom, even
    // though the size of the struct is the same.
    bool weightThresholdMet;
    // A BlessVoteProgress is considered invalid if weightThresholdMet is false when
    // s_versionedConfig.configVersion changes. we don't want old in-progress
    // votes to continue when we set a new config!
    // The config version at which the bless vote for a tagged root was initiated.
    uint32 configVersion;
    uint16 accumulatedWeight;
    // Care must be taken that the bitmap has at least as many bits as MAX_NUM_VOTERS.
    // uint200 is much larger than we need, but it saves us ~100 gas per voteToBless call to fill the word instead of
    // using a smaller type.
    // _bitmapGet(voterBitmap, i) = true indicates that the i-th voter has voted to bless
    uint200 voterBitmap;
  }

  mapping(bytes32 taggedRootHash => BlessVoteProgress blessVoteProgress) private s_blessVoteProgressByTaggedRootHash;

  event TaggedRootBlessed(uint32 indexed configVersion, IRMN.TaggedRoot taggedRoot, uint16 accumulatedWeight);
  event TaggedRootBlessVotesReset(uint32 indexed configVersion, IRMN.TaggedRoot taggedRoot, bool wasBlessed);
  event VotedToBless(uint32 indexed configVersion, address indexed voter, IRMN.TaggedRoot taggedRoot, uint8 weight);

  // These events make it easier for offchain logic to discover that it performs
  // the same actions multiple times.
  event AlreadyVotedToBless(uint32 indexed configVersion, address indexed voter, IRMN.TaggedRoot taggedRoot);
  event AlreadyBlessed(uint32 indexed configVersion, address indexed voter, IRMN.TaggedRoot taggedRoot);

  error VoteToBlessNoop();
  error VoteToBlessForbiddenDuringActiveGlobalCurse();

  constructor(Config memory config) RMNBase(config) {
    {
      // Ensure that the bitmap is large enough to hold MAX_NUM_VOTERS.
      // We do this in the constructor because MAX_NUM_VOTERS is constant.
      BlessVoteProgress memory vp = BlessVoteProgress({
        configVersion: 0,
        voterBitmap: type(uint200).max, // will not compile if it doesn't fit
        accumulatedWeight: 0,
        weightThresholdMet: false
      });
      assert(vp.voterBitmap >> (MAX_NUM_VOTERS - 1) >= 1);
    }
  }

  function _taggedRootHash(IRMN.TaggedRoot memory taggedRoot) internal pure returns (bytes32) {
    return keccak256(abi.encode(taggedRoot.commitStore, taggedRoot.root));
  }

  /// @param taggedRoots A tagged root is hashed as `keccak256(abi.encode(taggedRoot.commitStore
  /// /* address */, taggedRoot.root /* bytes32 */))`.
  /// @notice Tagged roots which are already (voted to be) blessed are skipped and emit corresponding events. In case
  /// the call has no effect, i.e., all passed tagged roots are skipped, the function reverts with a `VoteToBlessNoop`.
  function voteToBless(IRMN.TaggedRoot[] calldata taggedRoots) external {
    // If we have an active global curse, something is really wrong. Let's err on the
    // side of caution and not accept further blessings during this time of
    // uncertainty.
    if (isCursed(GLOBAL_CURSE_SUBJECT)) revert VoteToBlessForbiddenDuringActiveGlobalCurse();

    uint32 configVersion = s_versionedConfig.configVersion;
    BlesserRecord memory blesserRecord = s_blesserRecords[msg.sender];
    if (blesserRecord.configVersion != configVersion) revert UnauthorizedVoter(msg.sender);

    bool noop = true;
    for (uint256 i = 0; i < taggedRoots.length; ++i) {
      IRMN.TaggedRoot memory taggedRoot = taggedRoots[i];
      bytes32 taggedRootHash = _taggedRootHash(taggedRoot);
      BlessVoteProgress memory voteProgress = s_blessVoteProgressByTaggedRootHash[taggedRootHash];
      if (voteProgress.weightThresholdMet) {
        // We don't revert here because it's unreasonable to expect from the
        // voter to know exactly when to stop voting. Most likely when they
        // voted they didn't realize the threshold would be reached by the time
        // their vote was counted.
        // Additionally, there might be other tagged roots for which votes might
        // count, and we want to allow that to happen.
        emit AlreadyBlessed(configVersion, msg.sender, taggedRoot);
        continue;
      } else if (voteProgress.configVersion != configVersion) {
        // Note that voteProgress.weightThresholdMet must be false at this point

        // If votes were received while an older config was in effect,
        // invalidate them and start from scratch.
        // If votes were never received, set the current config version.
        voteProgress = BlessVoteProgress({
          configVersion: configVersion,
          voterBitmap: 0,
          accumulatedWeight: 0,
          weightThresholdMet: false
        });
      } else if (_bitmapGet(voteProgress.voterBitmap, blesserRecord.index)) {
        // We don't revert here because there might be other tagged roots for
        // which votes might count, and we want to allow that to happen.
        emit AlreadyVotedToBless(configVersion, msg.sender, taggedRoot);
        continue;
      }
      noop = false;
      voteProgress.voterBitmap = _bitmapSet(voteProgress.voterBitmap, blesserRecord.index);
      voteProgress.accumulatedWeight += blesserRecord.weight;
      emit VotedToBless(configVersion, msg.sender, taggedRoot, blesserRecord.weight);
      if (voteProgress.accumulatedWeight >= s_versionedConfig.config.blessWeightThreshold) {
        voteProgress.weightThresholdMet = true;
        emit TaggedRootBlessed(configVersion, taggedRoot, voteProgress.accumulatedWeight);
      }
      s_blessVoteProgressByTaggedRootHash[taggedRootHash] = voteProgress;
    }

    if (noop) {
      revert VoteToBlessNoop();
    }
  }

  /// @notice Can be called by the owner to remove unintentionally voted or even blessed tagged roots in a recovery
  /// scenario. The owner must ensure that there are no in-flight transactions by RMN nodes voting for any of the
  /// taggedRoots before calling this function, as such in-flight transactions could lead to the roots becoming
  /// re-blessed shortly after the call to this function, contrary to the original intention.
  function ownerResetBlessVotes(IRMN.TaggedRoot[] calldata taggedRoots) external onlyOwner {
    uint32 configVersion = s_versionedConfig.configVersion;
    for (uint256 i = 0; i < taggedRoots.length; ++i) {
      IRMN.TaggedRoot memory taggedRoot = taggedRoots[i];
      bytes32 taggedRootHash = _taggedRootHash(taggedRoot);
      BlessVoteProgress memory voteProgress = s_blessVoteProgressByTaggedRootHash[taggedRootHash];
      delete s_blessVoteProgressByTaggedRootHash[taggedRootHash];
      bool wasBlessed = voteProgress.weightThresholdMet;
      if (voteProgress.configVersion == configVersion || wasBlessed) {
        emit TaggedRootBlessVotesReset(configVersion, taggedRoot, wasBlessed);
      }
    }
  }

  /// @inheritdoc IRMN
  function isBlessed(IRMN.TaggedRoot calldata taggedRoot) external view returns (bool) {
    return s_blessVoteProgressByTaggedRootHash[_taggedRootHash(taggedRoot)].weightThresholdMet
      || s_permaBlessedCommitStores.contains(taggedRoot.commitStore);
  }

  /// @return blessVoteAddrs addresses of voters, will be empty if voting took place with an older config version
  /// @return accumulatedWeight sum of weights of voters, will be zero if voting took place with an older config version
  /// @return blessed will be accurate regardless of when voting took place
  /// @dev This is a helper method for offchain code so efficiency is not really a concern.
  function getBlessProgress(IRMN.TaggedRoot calldata taggedRoot)
    external
    view
    returns (address[] memory blessVoteAddrs, uint16 accumulatedWeight, bool blessed)
  {
    bytes32 taggedRootHash = _taggedRootHash(taggedRoot);
    BlessVoteProgress memory progress = s_blessVoteProgressByTaggedRootHash[taggedRootHash];
    blessed = progress.weightThresholdMet;
    if (progress.configVersion == s_versionedConfig.configVersion) {
      accumulatedWeight = progress.accumulatedWeight;
      uint200 bitmap = progress.voterBitmap;
      blessVoteAddrs = new address[](_bitmapCount(bitmap));
      Voter[] memory voters = s_versionedConfig.config.voters;
      uint256 j = 0;
      for (uint8 i = 0; i < voters.length; ++i) {
        if (_bitmapGet(bitmap, i)) {
          blessVoteAddrs[j] = voters[i].blessVoteAddr;
          ++j;
        }
      }
    }
  }
}
