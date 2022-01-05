// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/AFNInterface.sol";
import "../access/OwnerIsCreator.sol";

contract AFN is AFNInterface, OwnerIsCreator {
  // AFN party addresses and weights
  mapping(address => uint256) private s_parties;
  // List of AFN party addresses
  address[] private s_partyList;
  // Quorum of good votes to reach
  uint256 private s_goodQuorum;
  // Quorum of bad votes to reach
  uint256 private s_badQuorum;
  // The current round ID
  uint256 private s_round;
  // Version of the set of parties
  uint256 private s_committeeVersion;

  // Last heartbeat
  Heartbeat private s_lastHeartbeat;
  // The last round that a party voted good
  mapping(address => uint256) private s_lastGoodVote;
  // round => total good votes
  mapping(uint256 => uint256) private s_goodVotes;

  // Has a party voted bad
  mapping(address => bool) private s_hasVotedBad;
  // Parties that have voted bad
  address[] private s_badVoters;
  // Total bad votes
  uint256 private s_badVotes;
  // Whether or not there is a bad signal
  bool private s_badSignal;

  constructor(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) {
    _setConfig(parties, weights, goodQuorum, badQuorum, 1, 1);
  }

  ////////  VOTING  ////////

  /**
   * @notice Submit a good vote
   * @dev msg.sender must be a registered party
   * @param round the current round
   */
  function voteGood(uint256 round) external override {
    uint256 currentRound = s_round;
    if (round != currentRound) revert IncorrectRound(currentRound, round);
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    if (s_parties[sender] == 0) revert InvalidVoter(sender);
    if (s_lastGoodVote[sender] == currentRound) revert AlreadyVoted();

    s_lastGoodVote[sender] = currentRound;
    s_goodVotes[currentRound] += s_parties[sender];
    emit GoodVote(sender, currentRound);

    if (s_goodVotes[currentRound] >= s_goodQuorum) {
      Heartbeat memory heartbeat = Heartbeat({
        round: currentRound,
        timestamp: uint64(block.timestamp),
        committeeVersion: s_committeeVersion
      });
      s_lastHeartbeat = heartbeat;
      s_round++;
      emit AFNHeartbeat(heartbeat);
    }
  }

  /**
   * @notice Submit a bad vote
   * @dev msg.sender must be a registered party
   */
  function voteBad() external override {
    if (s_badSignal) revert MustRecoverFromBadSignal();
    address sender = msg.sender;
    uint256 senderWeight = s_parties[sender];
    if (senderWeight == 0) revert InvalidVoter(sender);
    if (s_hasVotedBad[sender]) revert AlreadyVoted();

    s_hasVotedBad[sender] = true;
    s_badVoters.push(sender);
    s_badVotes += senderWeight;

    if (s_badVotes >= s_badQuorum) {
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
    address[] memory badVoters = s_badVoters;
    for (uint256 i = 0; i < badVoters.length; i++) {
      s_hasVotedBad[badVoters[i]] = false;
    }
    s_badVotes = 0;
    delete s_badVoters;
    s_badSignal = false;
    emit RecoveredFromBadSignal();
  }

  /**
   * @notice Set config storage vars
   * @dev only callable by the owner
   * @param parties parties allowed to vote
   * @param weights weights of each party's vote
   * @param goodQuorum threshold to emit a heartbeat
   * @param badQuorum threashold to emit a bad signal
   */
  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external override onlyOwner {
    _setConfig(parties, weights, goodQuorum, badQuorum, s_round + 1, s_committeeVersion + 1);
  }

  ////////  Views ////////

  function hasBadSignal() external view override returns (bool) {
    return s_badSignal;
  }

  function getLastHeartbeat() external view override returns (Heartbeat memory) {
    return s_lastHeartbeat;
  }

  function getQuorums() external view returns (uint256 good, uint256 bad) {
    return (s_goodQuorum, s_badQuorum);
  }

  function getParties() external view returns (address[] memory) {
    return s_partyList;
  }

  function getWeight(address party) external view returns (uint256) {
    return s_parties[party];
  }

  function getRound() external view returns (uint256) {
    return s_round;
  }

  function getCommitteeVersion() external view returns (uint256) {
    return s_committeeVersion;
  }

  function getLastGoodVote(address party) external view returns (uint256) {
    return s_lastGoodVote[party];
  }

  function getGoodVotes(uint256 round) external view returns (uint256) {
    return s_goodVotes[round];
  }

  function getBadVotersAndVotes() external view returns (address[] memory voters, uint256 votes) {
    return (s_badVoters, s_badVotes);
  }

  function hasVotedBad(address party) external view returns (bool) {
    return s_hasVotedBad[party];
  }

  ////////  Private ////////

  /**
   * @notice Set detailed config storage vars
   */
  function _setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum,
    uint256 round,
    uint256 committeeVersion
  ) private {
    if (
      parties.length != weights.length ||
      parties.length == 0 ||
      goodQuorum == 0 ||
      badQuorum == 0 ||
      round == 0 ||
      committeeVersion == 0
    ) {
      revert InvalidConfig();
    }
    // Unset existing parties
    address[] memory existingParties = s_partyList;
    for (uint256 i = 0; i < existingParties.length; i++) {
      s_parties[existingParties[i]] = 0;
    }

    // Update round, committee and quorum details
    s_goodQuorum = goodQuorum;
    s_badQuorum = badQuorum;
    s_round = round;
    s_committeeVersion = committeeVersion;

    // Set new parties
    s_partyList = parties;
    for (uint256 i = 0; i < parties.length; i++) {
      if (weights[i] == 0) revert InvalidWeight();
      s_parties[parties[i]] = weights[i];
    }
    emit ConfigSet(parties, weights, goodQuorum, badQuorum);
  }
}
