// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../interfaces/AFNInterface.sol";
import "../access/OwnerIsCreator.sol";
import "../../interfaces/TypeAndVersionInterface.sol";

contract AFN is AFNInterface, OwnerIsCreator, TypeAndVersionInterface {
  // AFN participant addresses and weights
  mapping(address => uint256) private s_weightByParticipant;
  // List of AFN participant addresses
  address[] private s_participantList;
  // Quorum of good votes to reach for a heartbeat
  uint256 private s_weightThresholdForHeartbeat;
  // Quorum of bad votes to reach for a bad signal
  uint256 private s_weightThresholdForBadSignal;
  // The current round ID
  uint256 private s_round;
  // Version of the set of participants
  uint256 private s_committeeVersion;

  // Last heartbeat
  Heartbeat private s_lastHeartbeat;
  // The last round that a participant voted good
  mapping(address => uint256) private s_lastGoodVote;
  // Total good votes for this round
  uint256 private s_goodVotes;

  // Has a participant voted bad
  mapping(address => bool) private s_hasVotedBad;
  // participants that have voted bad
  address[] private s_badVoters;
  // Total bad votes
  uint256 private s_badVotes;
  // Whether or not there is a bad signal
  bool private s_badSignal;

  constructor(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForHeartbeat,
    uint256 weightThresholdForBadSignal
  ) {
    _setConfig(participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal, 1, 1);
  }

  ////////  VOTING  ////////

  /**
   * @notice Submit a good vote
   * @dev msg.sender must be a registered participant
   * @param round the current round
   */
  function voteGood(uint256 round) external override {
    uint256 currentRound = s_round;
    if (round != currentRound) revert IncorrectRound(currentRound, round);
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    if (s_weightByParticipant[sender] == 0) revert InvalidVoter(sender);
    if (s_lastGoodVote[sender] == currentRound) revert AlreadyVoted();

    s_lastGoodVote[sender] = currentRound;
    s_goodVotes += s_weightByParticipant[sender];
    emit GoodVote(sender, currentRound);

    if (s_goodVotes >= s_weightThresholdForHeartbeat) {
      Heartbeat memory heartbeat = Heartbeat({
        round: currentRound,
        timestamp: uint64(block.timestamp),
        committeeVersion: s_committeeVersion
      });
      s_lastHeartbeat = heartbeat;
      _newRound();
      emit AFNHeartbeat(heartbeat);
    }
  }

  /**
   * @notice Submit a bad vote
   * @dev msg.sender must be a registered participant
   */
  function voteBad() external override {
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    uint256 senderWeight = s_weightByParticipant[sender];
    if (senderWeight == 0) revert InvalidVoter(sender);
    if (s_hasVotedBad[sender]) revert AlreadyVoted();

    s_hasVotedBad[sender] = true;
    s_badVoters.push(sender);
    s_badVotes += senderWeight;

    if (s_badVotes >= s_weightThresholdForBadSignal) {
      s_badSignal = true;
      emit AFNBadSignal(block.timestamp);
    }
  }

  ////////  OnlyOwner ////////

  /**
   * @notice Recover from a bad signal
   * @dev only callable by the owner
   */
  function recover() external override onlyOwner {
    if (!s_badSignal) revert RecoveryNotNecessary();
    _clearBadVotes();
    s_badSignal = false;
    emit RecoveredFromBadSignal();
  }

  /**
   * @notice Set config storage vars
   * @dev only callable by the owner
   * @param participants participants allowed to vote
   * @param weights weights of each participant's vote
   * @param weightThresholdForHeartbeat threshold to emit a heartbeat
   * @param weightThresholdForBadSignal threashold to emit a bad signal
   */
  function setConfig(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForHeartbeat,
    uint256 weightThresholdForBadSignal
  ) external override onlyOwner {
    _setConfig(
      participants,
      weights,
      weightThresholdForHeartbeat,
      weightThresholdForBadSignal,
      s_round + 1,
      s_committeeVersion + 1
    );
  }

  ////////  Views ////////

  function hasBadSignal() external view override returns (bool) {
    return s_badSignal;
  }

  function getLastHeartbeat() external view override returns (Heartbeat memory) {
    return s_lastHeartbeat;
  }

  function getWeightThresholds() external view returns (uint256 good, uint256 bad) {
    return (s_weightThresholdForHeartbeat, s_weightThresholdForBadSignal);
  }

  function getParticipants() external view returns (address[] memory) {
    return s_participantList;
  }

  function getWeightByParticipant(address participant) external view returns (uint256) {
    return s_weightByParticipant[participant];
  }

  function getRound() external view returns (uint256) {
    return s_round;
  }

  function getCommitteeVersion() external view returns (uint256) {
    return s_committeeVersion;
  }

  function getLastGoodVoteByParticipant(address participant) external view returns (uint256) {
    return s_lastGoodVote[participant];
  }

  function getGoodVotes() external view returns (uint256) {
    return s_goodVotes;
  }

  function getBadVotersAndVotes() external view returns (address[] memory voters, uint256 votes) {
    return (s_badVoters, s_badVotes);
  }

  function hasVotedBad(address participant) external view returns (bool) {
    return s_hasVotedBad[participant];
  }

  ////////  Private ////////

  /**
   * @notice Increment the round and reset good votes
   */
  function _newRound() private {
    s_round++;
    s_goodVotes = 0;
  }

  /**
   * @notice Clear all bad votes and voters
   */
  function _clearBadVotes() private {
    address[] memory badVoters = s_badVoters;
    for (uint256 i = 0; i < badVoters.length; i++) {
      s_hasVotedBad[badVoters[i]] = false;
    }
    s_badVotes = 0;
    delete s_badVoters;
  }

  /**
   * @notice Set detailed config storage vars
   */
  function _setConfig(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForHeartbeat,
    uint256 weightThresholdForBadSignal,
    uint256 round,
    uint256 committeeVersion
  ) private {
    if (
      participants.length != weights.length ||
      participants.length == 0 ||
      weightThresholdForHeartbeat == 0 ||
      weightThresholdForBadSignal == 0 ||
      round == 0 ||
      committeeVersion == 0
    ) {
      revert InvalidConfig();
    }
    // Unset existing participants
    address[] memory existingParticipants = s_participantList;
    for (uint256 i = 0; i < existingParticipants.length; i++) {
      s_weightByParticipant[existingParticipants[i]] = 0;
    }

    // Update round, committee and quorum details
    s_weightThresholdForHeartbeat = weightThresholdForHeartbeat;
    s_weightThresholdForBadSignal = weightThresholdForBadSignal;
    s_committeeVersion = committeeVersion;
    _newRound();
    _clearBadVotes();

    uint256 weightTotal = 0;
    // Set new participants
    s_participantList = participants;
    for (uint256 i = 0; i < participants.length; i++) {
      if (participants[i] == address(0)) revert InvalidConfig();
      if (weights[i] == 0) revert InvalidWeight();
      s_weightByParticipant[participants[i]] = weights[i];
      weightTotal += weights[i];
    }
    if (weightTotal < weightThresholdForHeartbeat || weightTotal < weightThresholdForBadSignal) {
      revert InvalidConfig();
    }
    emit ConfigSet(participants, weights, weightThresholdForHeartbeat, weightThresholdForBadSignal);
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "AFN 0.0.1";
  }
}
