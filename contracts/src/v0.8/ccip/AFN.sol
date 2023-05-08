// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {IAFN} from "./interfaces/IAFN.sol";

import {OwnerIsCreator} from "./OwnerIsCreator.sol";

contract AFN is IAFN, OwnerIsCreator, TypeAndVersionInterface {
  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "AFN 1.0.0";

  uint256 private constant MAX_NUM_VOTERS = 128;

  // DYNAMIC CONFIG
  struct Voter {
    // This is the address the voter should use to call voteToBless.
    address blessVoteAddr;
    // This is the address the voter should use to call voteToCurse.
    address curseVoteAddr;
    // This is the address the voter should use to call unvoteToCurse.
    address curseUnvoteAddr;
    // The weight of this voter's vote for blessing.
    uint8 blessWeight;
    // The weight of this voter's vote for cursing.
    uint8 curseWeight;
  }

  struct Config {
    Voter[] voters;
    // When the total weight of voters that have voted to bless a tagged root reaches
    // or exceeds blessWeightThreshold, the tagged root becomes blessed.
    uint16 blessWeightThreshold;
    // When the total weight of voters that have voted to curse reaches or
    // exceeds curseWeightThreshold, a curse is emitted.
    uint16 curseWeightThreshold;
  }

  struct VersionedConfig {
    Config config;
    // The version is incremented every time the config changes.
    uint32 configVersion;
    // The block number at which the config was last set. Helps the offchain
    // code check that the config was set in a stable block or double-check
    // that it has the correct config through by querying logs at that block
    // number.
    uint32 blockNumber;
  }

  VersionedConfig private s_versionedConfig;

  // STATE
  struct BlesserRecord {
    // The config version at which this BlesserRecord was last set. A blesser
    // is considered active iff this configVersion equals
    // s_versionedConfig.configVersion.
    uint32 configVersion;
    uint8 weight;
    uint8 index;
  }

  mapping(address => BlesserRecord) private s_blesserRecords;

  struct BlessVoteProgress {
    // A BlessVoteProgress is considered invalid if weightThresholdMet is false when
    // s_versionedConfig.configVersion changes. we don't want old in-progress
    // votes to continue when we set a new config!
    uint32 configVersion;
    uint16 accumulatedWeight;
    // Care must be taken that the bitmap has as many bits as MAX_NUM_VOTERS.
    uint128 voterBitmap;
    bool weightThresholdMet;
  }

  mapping(bytes32 => BlessVoteProgress) private s_blessVoteProgressByTaggedRootHash;

  struct CurserRecord {
    bool active;
    uint8 weight;
    // Stores a count of the successful voteToCurse invocations by this curser
    // since their votes were last reset (through setConfig, voteToCurse, or
    // ownerUnvoteToCurse).
    uint32 voteCount;
    address curseUnvoteAddr;
    bytes32 cursesHash;
  }

  mapping(address => CurserRecord) private s_curserRecords;

  // Maintains a per-curser set of curseIds. Entries from this mapping are
  // never cleared. Once a curseId is used it can never be reused, even after
  // an unvoteToCurse or ownerUnvoteToCurse.
  mapping(address => mapping(bytes32 => bool)) private s_curseVotes;

  struct CurseVoteProgress {
    uint16 curseWeightThreshold;
    uint16 accumulatedWeight;
    bool weightThresholdMet;
  }

  CurseVoteProgress private s_curseVoteProgress;

  // AUXILLARY STRUCTS
  struct UnvoteToCurseRecord {
    address curseVoteAddr;
    bytes32 cursesHash;
    bool forceUnvote;
  }

  // EVENTS, ERRORS
  event ConfigSet(uint32 indexed configVersion, Config config);
  error InvalidConfig();

  event TaggedRootBlessed(uint32 indexed configVersion, IAFN.TaggedRoot taggedRoot, uint16 votes);
  event VoteToBless(uint32 indexed configVersion, address indexed voter, IAFN.TaggedRoot taggedRoot, uint8 weight);

  event VoteToCurse(
    uint32 indexed configVersion,
    address indexed voter,
    uint8 weight,
    uint32 voteCount,
    bytes32 curseId,
    bytes32 cursesHash,
    uint16 accumulatedWeight
  );
  event ReusedVotesToCurse(
    uint32 indexed configVersion,
    address indexed voter,
    uint8 weight,
    uint32 voteCount,
    bytes32 cursesHash,
    uint16 accumulatedWeight
  );
  event UnvoteToCurse(
    uint32 indexed configVersion,
    address indexed voter,
    uint8 weight,
    uint32 voteCount,
    bytes32 cursesHash
  );
  event SkippedUnvoteToCurse(address indexed voter, bytes32 expectedCursesHash, bytes32 actualCursesHash);
  event Cursed(uint32 indexed configVersion, uint256 timestamp);

  // These events make it easier for offchain logic to discover that it performs
  // the same actions multiple times.
  event AlreadyVotedToBless(uint32 indexed configVersion, address indexed voter, IAFN.TaggedRoot taggedRoot);
  event AlreadyBlessed(uint32 indexed configVersion, address indexed voter, IAFN.TaggedRoot taggedRoot);

  event RecoveredFromCurse();

  error AlreadyVotedToCurse(address voter, bytes32 curseId);
  error InvalidVoter(address voter);
  error InvalidCursesHash(bytes32 expectedCursesHash, bytes32 actualCursesHash);
  error MustRecoverFromCurse();

  constructor(Config memory config) {
    {
      // Ensure that the bitmap is large enough to hold MAX_NUM_VOTERS.
      // We do this in the constructor because MAX_NUM_VOTERS is constant.
      BlessVoteProgress memory vp;
      vp.voterBitmap = uint128((1 << MAX_NUM_VOTERS) - 1);
    }
    _setConfig(config);
  }

  function _bitmapGet(uint128 bitmap, uint8 index) internal pure returns (bool) {
    return (bitmap >> index) & 1 == 1;
  }

  function _bitmapSet(uint128 bitmap, uint8 index) internal pure returns (uint128) {
    return bitmap | (uint128(1) << index);
  }

  function _bitmapCount(uint128 bitmap) internal pure returns (uint8 oneBits) {
    // https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan
    for (; bitmap != 0; ++oneBits) {
      bitmap &= bitmap - 1;
    }
  }

  function _taggedRootHash(IAFN.TaggedRoot memory taggedRoot) internal pure returns (bytes32) {
    return keccak256(abi.encode(taggedRoot.commitStore, taggedRoot.root));
  }

  /// @param taggedRoots A tagged root is hashed as `keccak256(abi.encode(taggedRoot.commitStore
  /// /* address */, taggedRoot.root /* bytes32 */))`.
  function voteToBless(IAFN.TaggedRoot[] calldata taggedRoots) external {
    // If we have an active curse, something is really wrong. Let's err on the
    // side of caution and not accept further blessings during this time of
    // uncertainty.
    if (isCursed()) revert MustRecoverFromCurse();

    uint32 configVersion = s_versionedConfig.configVersion;
    BlesserRecord memory blesserRecord = s_blesserRecords[msg.sender];
    if (blesserRecord.configVersion != configVersion) revert InvalidVoter(msg.sender);

    for (uint256 i = 0; i < taggedRoots.length; ++i) {
      IAFN.TaggedRoot memory taggedRoot = taggedRoots[i];
      bytes32 taggedRootHash = _taggedRootHash(taggedRoots[i]);
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
      }
      if (voteProgress.configVersion != configVersion) {
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
      }
      if (_bitmapGet(voteProgress.voterBitmap, blesserRecord.index)) {
        // We don't revert here because there might be other tagged roots for
        // which votes might count, and we want to allow that to happen.
        emit AlreadyVotedToBless(configVersion, msg.sender, taggedRoot);
        continue;
      }
      voteProgress.voterBitmap = _bitmapSet(voteProgress.voterBitmap, blesserRecord.index);
      voteProgress.accumulatedWeight += blesserRecord.weight;
      emit VoteToBless(configVersion, msg.sender, taggedRoot, blesserRecord.weight);
      if (voteProgress.accumulatedWeight >= s_versionedConfig.config.blessWeightThreshold) {
        voteProgress.weightThresholdMet = true;
        emit TaggedRootBlessed(configVersion, taggedRoot, voteProgress.accumulatedWeight);
      }
      s_blessVoteProgressByTaggedRootHash[taggedRootHash] = voteProgress;
    }
  }

  /// @notice Can be called by the owner to remove unintentionally blessed tagged roots
  /// in a recovery scenario.
  function ownerUnbless(IAFN.TaggedRoot[] memory taggedRoots) external onlyOwner {
    for (uint256 i = 0; i < taggedRoots.length; ++i) {
      delete s_blessVoteProgressByTaggedRootHash[_taggedRootHash(taggedRoots[i])];
    }
  }

  /// @notice Can be called by a curser to remove unintentional votes to curse.
  /// We expect this to be called very rarely, e.g. in case of a bug in the
  /// offchain code causing false voteToCurse calls.
  /// @notice Should be called from curser's corresponding curseUnvoteAddr.
  function unvoteToCurse(address curseVoteAddr, bytes32 cursesHash) external {
    _unvoteToCurse(
      false,
      UnvoteToCurseRecord({curseVoteAddr: curseVoteAddr, cursesHash: cursesHash, forceUnvote: false})
    );
  }

  function _unvoteToCurse(bool ownerCall, UnvoteToCurseRecord memory batch) internal {
    CurserRecord memory curserRecord = s_curserRecords[batch.curseVoteAddr];
    if (!ownerCall && msg.sender != curserRecord.curseUnvoteAddr) revert InvalidVoter(msg.sender);

    // If a curse is active, we want only the owner to be allowed to lift it.
    if (!ownerCall && isCursed()) revert MustRecoverFromCurse();

    if (!curserRecord.active || curserRecord.voteCount == 0) return;

    // Owner can avoid the curses hash check by setting forceUnvote to true, in case
    // a malicious curser is flooding the system with votes to curse with the
    // intention to disallow the owner to clear their curse.
    if (ownerCall && !batch.forceUnvote && curserRecord.cursesHash != batch.cursesHash) {
      emit SkippedUnvoteToCurse(batch.curseVoteAddr, curserRecord.cursesHash, batch.cursesHash);
      return;
    }
    if (msg.sender == curserRecord.curseUnvoteAddr && curserRecord.cursesHash != batch.cursesHash)
      revert InvalidCursesHash(curserRecord.cursesHash, batch.cursesHash);

    emit UnvoteToCurse(
      s_versionedConfig.configVersion,
      batch.curseVoteAddr,
      curserRecord.weight,
      curserRecord.voteCount,
      batch.cursesHash
    );
    curserRecord.voteCount = 0;
    s_curserRecords[batch.curseVoteAddr] = curserRecord;
    s_curseVoteProgress.accumulatedWeight -= curserRecord.weight;
    // If not ownerCall, no need to update weightThresholdMet as it must already have been false before.
    // If ownerCall, further logic to update weightThresholdMet follows in ownerUnvoteToCurse.
  }

  /// @notice A vote to curse is appropriate during unhealthy network conditions
  /// (eg. unexpected reorgs).
  function voteToCurse(bytes32 curseId) external {
    CurserRecord memory curserRecord = s_curserRecords[msg.sender];
    if (!curserRecord.active) revert InvalidVoter(msg.sender);
    if (s_curseVotes[msg.sender][curseId]) revert AlreadyVotedToCurse(msg.sender, curseId);
    s_curseVotes[msg.sender][curseId] = true;
    ++curserRecord.voteCount;
    curserRecord.cursesHash = keccak256(
      abi.encode(curserRecord.cursesHash, block.chainid, blockhash(block.number - 1), curseId)
    );
    s_curserRecords[msg.sender] = curserRecord;
    if (curserRecord.voteCount == 1) {
      s_curseVoteProgress.accumulatedWeight += curserRecord.weight;
      // TODO: we could add the version to configVersion to avoid the extra slot access.
      uint32 configVersion = s_versionedConfig.configVersion;
      emit VoteToCurse(
        configVersion,
        msg.sender,
        curserRecord.weight,
        curserRecord.voteCount,
        curseId,
        curserRecord.cursesHash,
        s_curseVoteProgress.accumulatedWeight
      );
      if (!s_curseVoteProgress.weightThresholdMet) {
        if (s_curseVoteProgress.accumulatedWeight >= s_curseVoteProgress.curseWeightThreshold) {
          s_curseVoteProgress.weightThresholdMet = true;
          emit Cursed(configVersion, block.timestamp);
        }
      }
    }
  }

  /// @notice Enables the owner to remove curse votes. After the curse votes are removed,
  /// this function will check whether the curse is still valid and restore the healthy state if possible.
  function ownerUnvoteToCurse(UnvoteToCurseRecord[] memory records) external onlyOwner {
    for (uint256 i = 0; i < records.length; ++i) {
      _unvoteToCurse(true, records[i]);
    }

    if (
      s_curseVoteProgress.weightThresholdMet &&
      s_curseVoteProgress.accumulatedWeight < s_curseVoteProgress.curseWeightThreshold
    ) {
      s_curseVoteProgress.weightThresholdMet = false;
      emit RecoveredFromCurse();
      // Invalidate all in-progress votes to bless by bumping the config.
      _setConfig(s_versionedConfig.config);
    }
  }

  /// @notice Will revert in case a curse is active.
  function setConfig(Config memory config) external onlyOwner {
    _setConfig(config);
  }

  /// @inheritdoc IAFN
  function isBlessed(IAFN.TaggedRoot calldata taggedRoot) public view override returns (bool) {
    return s_blessVoteProgressByTaggedRootHash[_taggedRootHash(taggedRoot)].weightThresholdMet;
  }

  /// @inheritdoc IAFN
  function isCursed() public view override returns (bool) {
    return s_curseVoteProgress.weightThresholdMet;
  }

  /// @notice Config version might be incremented for many reasons, including
  /// recovery from a curse and a regular config change.
  function getConfigDetails()
    external
    view
    returns (
      uint32 version,
      uint32 blockNumber,
      Config memory config
    )
  {
    version = s_versionedConfig.configVersion;
    blockNumber = s_versionedConfig.blockNumber;
    config = s_versionedConfig.config;
  }

  /// @notice Get addresses of those who have voted to bless tagged root, and the total weight of their votes.
  /// @return blessVoteAddrs will be empty if voting took place with an older config version
  /// @dev This is a helper method for offchain code so efficiency is not really a concern.
  function getBlessVotersAndWeight(IAFN.TaggedRoot calldata taggedRoot)
    external
    view
    returns (address[] memory blessVoteAddrs, uint16 weight)
  {
    bytes32 taggedRootHash = _taggedRootHash(taggedRoot);
    BlessVoteProgress memory progress = s_blessVoteProgressByTaggedRootHash[taggedRootHash];
    weight = progress.accumulatedWeight;
    if (progress.configVersion != s_versionedConfig.configVersion) {
      if (!progress.weightThresholdMet) {
        // If the threshold wasn't met when the vote took place with an earlier
        // config, the weight will be reset when voting restarts.
        weight = 0;
      }
      // If we're not on the same config on which the voting to bless last took
      // place we can't know for sure who voted, so return an empty array.
      return (new address[](0), weight);
    }
    uint128 bitmap = progress.voterBitmap;
    blessVoteAddrs = new address[](_bitmapCount(bitmap));
    Voter[] memory voters = s_versionedConfig.config.voters;
    uint256 j = 0;
    for (uint256 i = 0; i < voters.length; ++i) {
      if (_bitmapGet(bitmap, s_blesserRecords[voters[i].blessVoteAddr].index)) {
        blessVoteAddrs[j++] = voters[i].blessVoteAddr;
      }
    }
  }

  /// @notice Get addresses of those who have voted to curse, and the total weight of their votes.
  /// @return curseVoteAddrs will be empty if voting took place with an older config version
  /// @dev This is a helper method for offchain code so efficiency is not really a concern.
  function getCurseVotersAndWeight()
    external
    view
    returns (
      address[] memory curseVoteAddrs,
      uint16 weight,
      uint32[] memory voteCounts
    )
  {
    weight = s_curseVoteProgress.accumulatedWeight;
    uint256 numCursers;
    Voter[] memory voters = s_versionedConfig.config.voters;
    for (uint256 i = 0; i < voters.length; ++i) {
      CurserRecord memory curserRecord = s_curserRecords[voters[i].curseVoteAddr];
      if (curserRecord.active && curserRecord.voteCount > 0) {
        ++numCursers;
      }
    }
    curseVoteAddrs = new address[](numCursers);
    voteCounts = new uint32[](numCursers);
    uint256 j = 0;
    for (uint256 i = 0; i < voters.length; ++i) {
      address curseVoteAddr = voters[i].curseVoteAddr;
      CurserRecord memory curserRecord = s_curserRecords[curseVoteAddr];
      if (curserRecord.active && curserRecord.voteCount > 0) {
        curseVoteAddrs[j] = curseVoteAddr;
        voteCounts[j] = curserRecord.voteCount;
        ++j;
      }
    }
  }

  function _validateConfig(Config memory config) internal pure returns (bool) {
    if (
      config.voters.length == 0 ||
      config.voters.length > MAX_NUM_VOTERS ||
      config.blessWeightThreshold == 0 ||
      config.curseWeightThreshold == 0
    ) {
      return false;
    }

    uint256 totalBlessWeight = 0;
    uint256 totalCurseWeight = 0;
    address[] memory allAddrs = new address[](config.voters.length * 3);
    for (uint256 i = 0; i < config.voters.length; ++i) {
      Voter memory voter = config.voters[i];
      if (
        voter.blessVoteAddr == address(0) ||
        voter.curseVoteAddr == address(0) ||
        voter.curseUnvoteAddr == address(0) ||
        voter.blessWeight + voter.curseWeight == 0
      ) {
        return false;
      }
      allAddrs[3 * i + 0] = voter.blessVoteAddr;
      allAddrs[3 * i + 1] = voter.curseVoteAddr;
      allAddrs[3 * i + 2] = voter.curseUnvoteAddr;
      totalBlessWeight += voter.blessWeight;
      totalCurseWeight += voter.curseWeight;
    }
    for (uint256 i = 0; i < allAddrs.length; ++i) {
      for (uint256 j = i + 1; j < allAddrs.length; ++j) {
        if (allAddrs[i] == allAddrs[j]) {
          return false;
        }
      }
    }

    return totalBlessWeight >= config.blessWeightThreshold && totalCurseWeight >= config.curseWeightThreshold;
  }

  function _setConfig(Config memory config) private {
    if (isCursed()) revert MustRecoverFromCurse();
    if (!_validateConfig(config)) revert InvalidConfig();

    Config memory oldConfig = s_versionedConfig.config;

    // We can't directly assign s_versionedConfig.config to config
    // because copying a memory array into storage is not supported.
    {
      s_versionedConfig.config.blessWeightThreshold = config.blessWeightThreshold;
      s_versionedConfig.config.curseWeightThreshold = config.curseWeightThreshold;
      while (s_versionedConfig.config.voters.length != 0) {
        Voter memory voter = s_versionedConfig.config.voters[s_versionedConfig.config.voters.length - 1];
        delete s_blesserRecords[voter.blessVoteAddr];
        s_curserRecords[voter.curseVoteAddr].active = false;
        s_versionedConfig.config.voters.pop();
      }
      for (uint256 i = 0; i < config.voters.length; ++i) {
        s_versionedConfig.config.voters.push(config.voters[i]);
      }
    }

    uint32 configVersion = ++s_versionedConfig.configVersion;
    for (uint8 i = 0; i < config.voters.length; ++i) {
      Voter memory voter = config.voters[i];
      s_blesserRecords[voter.blessVoteAddr] = BlesserRecord({
        configVersion: configVersion,
        index: i,
        weight: voter.blessWeight
      });
      s_curserRecords[voter.curseVoteAddr].active = true;
      s_curserRecords[voter.curseVoteAddr].weight = voter.curseWeight;
      s_curserRecords[voter.curseVoteAddr].curseUnvoteAddr = voter.curseUnvoteAddr;
    }
    s_versionedConfig.blockNumber = uint32(block.number);
    emit ConfigSet(configVersion, config);

    CurseVoteProgress memory newCurseVoteProgress = CurseVoteProgress({
      curseWeightThreshold: config.curseWeightThreshold,
      accumulatedWeight: 0,
      weightThresholdMet: false
    });

    // Retain votes for the cursers who are still part of the new config.
    for (uint8 i = 0; i < oldConfig.voters.length; ++i) {
      // We could be more efficient with this but since this is only for
      // setConfig it will do for now.
      address curseVoteAddr = oldConfig.voters[i].curseVoteAddr;
      CurserRecord memory curserRecord = s_curserRecords[curseVoteAddr];
      if (curserRecord.active && curserRecord.voteCount > 0) {
        newCurseVoteProgress.accumulatedWeight += curserRecord.weight;
        emit ReusedVotesToCurse(
          configVersion,
          curseVoteAddr,
          curserRecord.weight,
          curserRecord.voteCount,
          curserRecord.cursesHash,
          newCurseVoteProgress.accumulatedWeight
        );
      }
    }
    newCurseVoteProgress.weightThresholdMet =
      newCurseVoteProgress.accumulatedWeight >= newCurseVoteProgress.curseWeightThreshold;
    if (newCurseVoteProgress.weightThresholdMet) {
      emit Cursed(configVersion, block.timestamp);
    }
    s_curseVoteProgress = newCurseVoteProgress;
  }
}
