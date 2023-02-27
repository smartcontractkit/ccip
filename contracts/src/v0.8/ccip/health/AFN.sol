// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

contract AFN is IAFN, OwnerIsCreator, TypeAndVersionInterface {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "AFN 1.0.0";

  // AFN participant addresses and weights
  mapping(address => uint256) private s_weightByParticipant;
  // List of AFN participant addresses
  address[] private s_participantList;
  // Quorum of good votes to reach for a blessing
  uint256 private s_weightThresholdForBlessing;
  // Quorum of bad votes to reach for a bad signal
  uint256 private s_weightThresholdForBadSignal;

  // Config version
  uint256 private s_configVersion;

  // Participants that have voted bad
  address[] private s_badVoters;
  // Has voted bad?
  mapping(address => bool) private s_hasVotedBad;
  // Total number of current bad votes. Needed to avoid iterating
  // over s_badVoters to check whether a bad signal is received.
  uint256 private s_totalBadVotes;

  // Roots stored here are blessed
  mapping(bytes32 => bool) private s_blessedRoots;
  // Votes to bless roots (root => weighted blessings from participants)
  mapping(bytes32 => uint256) private s_votesToBlessRoots;
  // A list of all the roots voted on by the current config, but have not yet been blessed
  bytes32[] private s_rootsVotedOnToBlessList;
  // Blessing record mapping(root => mapping(participant => has votes to bless))
  // We assume there will be more roots voted on than participants. Using the root as the primary
  // key enables more efficient clearing when a new config is set via `setConfig`.
  // Loop once through the roots, for each root, delete the inner mapping per participant.
  mapping(bytes32 => mapping(address => bool)) private s_rootsVotedToBlessByParticipant;

  constructor(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForBlessing,
    uint256 weightThresholdForBadSignal
  ) {
    _setAFNConfig(participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal, 1);
  }

  /// @inheritdoc IAFN
  function voteToBlessRoots(bytes32[] calldata rootsWithOrigin) external override hasNotReceivedBadSignal {
    address sender = msg.sender;
    uint256 senderWeight = s_weightByParticipant[sender];
    if (senderWeight == 0) revert InvalidVoter(sender);

    for (uint256 i = 0; i < rootsWithOrigin.length; ++i) {
      bytes32 root = rootsWithOrigin[i];
      if (isBlessed(root)) continue;
      if (s_rootsVotedToBlessByParticipant[root][sender]) continue;
      s_rootsVotedToBlessByParticipant[root][sender] = true;

      uint256 currentVotesToBless = s_votesToBlessRoots[root];
      if (currentVotesToBless == 0) {
        s_rootsVotedOnToBlessList.push(root);
      }

      uint256 newVotesToBless = currentVotesToBless + senderWeight;
      s_votesToBlessRoots[root] = newVotesToBless;
      emit VoteToBless(sender, root, senderWeight);
      if (newVotesToBless >= s_weightThresholdForBlessing) {
        s_blessedRoots[root] = true;
        emit RootBlessed(root, newVotesToBless);
      }
    }
  }

  /// @inheritdoc IAFN
  function voteBad() external override hasNotReceivedBadSignal {
    address sender = msg.sender;
    uint256 senderWeight = s_weightByParticipant[sender];
    if (senderWeight == 0) revert InvalidVoter(sender);
    if (s_hasVotedBad[sender]) revert AlreadyVoted();

    s_hasVotedBad[sender] = true;
    s_badVoters.push(sender);
    s_totalBadVotes += senderWeight;

    emit VoteBad(sender, senderWeight);

    if (badSignalReceived()) {
      emit AFNBadSignal(block.timestamp);
    }
  }

  /// @inheritdoc IAFN
  function recoverFromBadSignal() external override onlyOwner {
    if (!badSignalReceived()) revert RecoveryNotNecessary();
    _clearBadVotes();
    emit RecoveredFromBadSignal();
  }

  /// @inheritdoc IAFN
  function setAFNConfig(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForBlessing,
    uint256 weightThresholdForBadSignal
  ) external override onlyOwner {
    _setAFNConfig(participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal, s_configVersion + 1);
  }

  /// @inheritdoc IAFN
  function isBlessed(bytes32 rootWithOrigin) public view override returns (bool) {
    return s_blessedRoots[rootWithOrigin];
  }

  /// @inheritdoc IAFN
  function getVotesToBlessRoot(bytes32 root) public view override returns (uint256) {
    return s_votesToBlessRoots[root];
  }

  /// @inheritdoc IAFN
  function hasVotedToBlessRoot(address participant, bytes32 root) public view override returns (bool) {
    return s_rootsVotedToBlessByParticipant[root][participant];
  }

  /// @inheritdoc IAFN
  function badSignalReceived() public view override returns (bool) {
    return (s_totalBadVotes >= s_weightThresholdForBadSignal);
  }

  /// @inheritdoc IAFN
  function getWeightThresholds() external view override returns (uint256 blessing, uint256 badSignal) {
    return (s_weightThresholdForBlessing, s_weightThresholdForBadSignal);
  }

  /// @inheritdoc IAFN
  function getParticipants() external view override returns (address[] memory) {
    return s_participantList;
  }

  /// @inheritdoc IAFN
  function getWeightByParticipant(address participant) external view override returns (uint256) {
    return s_weightByParticipant[participant];
  }

  /// @inheritdoc IAFN
  function getConfigVersion() external view override returns (uint256) {
    return s_configVersion;
  }

  /// @inheritdoc IAFN
  function getBadVotersAndVotes() external view override returns (address[] memory voters, uint256 votes) {
    return (s_badVoters, s_totalBadVotes);
  }

  /// @inheritdoc IAFN
  function hasVotedBad(address participant) external view override returns (bool) {
    return s_hasVotedBad[participant];
  }

  /// @notice Clear all bad votes and voters
  function _clearBadVotes() private {
    address[] memory badVoters = s_badVoters;
    for (uint256 i = 0; i < badVoters.length; ++i) {
      delete s_hasVotedBad[badVoters[i]];
    }
    delete s_badVoters;
    delete s_totalBadVotes;
  }

  function _clearVotesToBlessRoots(address[] memory participants) private {
    bytes32[] memory rootsVotedOn = s_rootsVotedOnToBlessList;
    for (uint256 i = 0; i < rootsVotedOn.length; ++i) {
      bytes32 root = rootsVotedOn[i];
      delete s_votesToBlessRoots[root];
      for (uint256 j = 0; j < participants.length; ++j) {
        delete s_rootsVotedToBlessByParticipant[root][participants[j]];
      }
    }
    delete s_rootsVotedOnToBlessList;
  }

  /// @notice Set detailed config storage vars
  function _setAFNConfig(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForBlessing,
    uint256 weightThresholdForBadSignal,
    uint256 configVersion
  ) private {
    if (
      participants.length != weights.length ||
      participants.length == 0 ||
      weightThresholdForBlessing == 0 ||
      weightThresholdForBadSignal == 0 ||
      configVersion == 0
    ) revert InvalidConfig();
    // Unset existing participants
    address[] memory existingParticipants = s_participantList;
    for (uint256 i = 0; i < existingParticipants.length; ++i) {
      delete s_weightByParticipant[existingParticipants[i]];
    }

    // Update config and quorum details
    s_weightThresholdForBlessing = weightThresholdForBlessing;
    s_weightThresholdForBadSignal = weightThresholdForBadSignal;
    s_configVersion = configVersion;
    _clearBadVotes();
    _clearVotesToBlessRoots(existingParticipants);

    uint256 weightTotal = 0;
    // Set new participants
    s_participantList = participants;
    for (uint256 i = 0; i < participants.length; ++i) {
      if (participants[i] == address(0)) revert InvalidConfig();
      if (weights[i] == 0) revert InvalidWeight();
      s_weightByParticipant[participants[i]] = weights[i];
      weightTotal += weights[i];
    }
    if (weightTotal < weightThresholdForBlessing || weightTotal < weightThresholdForBadSignal) revert InvalidConfig();
    emit AFNConfigSet(participants, weights, weightThresholdForBlessing, weightThresholdForBadSignal);
  }

  modifier hasNotReceivedBadSignal() {
    if (badSignalReceived()) revert MustRecoverFromBadSignal();
    _;
  }
}
