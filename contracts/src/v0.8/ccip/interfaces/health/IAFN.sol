// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IAFN {
  event RootBlessed(bytes32 indexed root, uint256 votes);
  event VoteToBless(address indexed voter, bytes32 indexed root, uint256 weight);
  event VoteBad(address indexed voter, uint256 weight);
  event AFNBadSignal(uint256 timestamp);
  event RecoveredFromBadSignal();
  event AFNConfigSet(address[] parties, uint256[] weights, uint256 goodQuorum, uint256 badQuorum);

  error InvalidVoter(address voter);
  error AlreadyVoted();
  error InvalidConfig();
  error InvalidWeight();
  error MustRecoverFromBadSignal();
  error RecoveryNotNecessary();

  /**
   * @notice Check if a bad signal has been received
   * @return bool badSignal
   */
  function badSignalReceived() external view returns (bool);

  /**
   * @notice Vote to bless a set of roots with Origin
   * @param rootsWithOrigin - array of roots
   */
  function voteToBlessRoots(bytes32[] calldata rootsWithOrigin) external;

  /**
   * @notice Get thresholds for blessing and bad signal
   * @return blessing threshold for blessing
   * @return badSignal threshold for bad signal
   */
  function getWeightThresholds() external returns (uint256 blessing, uint256 badSignal);

  /**
   * @notice Check if a participant has voted to bless a root
   * @param participant address
   * @param root bytes32
   * @return bool has voted to bless
   */
  function hasVotedToBlessRoot(address participant, bytes32 root) external view returns (bool);

  /**
   * @notice Get all configured participants
   * @return participants address array
   */
  function getParticipants() external returns (address[] memory);

  /**
   * @notice Get the weight of a participant
   * @param participant address
   * @return weight uint256
   */
  function getWeightByParticipant(address participant) external view returns (uint256);

  /**
   * @notice Get the config version
   * @return version uint256
   */
  function getConfigVersion() external view returns (uint256);

  /**
   * @notice Get participants who have voted bad, and the total number of bad votes
   * @return voters address array
   * @return votes total number of bad votes
   */
  function getBadVotersAndVotes() external view returns (address[] memory voters, uint256 votes);

  /**
   * @notice Get the number of votes to bless a particular root
   * @param root bytes32
   * @return votes number of votes
   */
  function getVotesToBlessRoot(bytes32 root) external view returns (uint256);

  /**
   * @notice Check if a participant has voted bad
   * @param participant address
   * @return hasVotedBad bool
   */
  function hasVotedBad(address participant) external view returns (bool);

  /**
   * @notice Vote bad
   */
  function voteBad() external;

  /**
   * @notice Check if a root is blessed
   * @param rootWithOrigin bytes32
   * @return isBlessed bool
   */
  function isBlessed(bytes32 rootWithOrigin) external view returns (bool);

  /**
   * @notice Recover from a bad signal
   */
  function recoverFromBadSignal() external;

  /**
   * @notice Set config storage vars
   * @dev only callable by the owner
   * @param participants participants allowed to vote
   * @param weights weights of each participant's vote
   * @param weightThresholdForBlessing threshold to emit a blessing
   * @param weightThresholdForBadSignal threshold to emit a bad signal
   */
  function setAFNConfig(
    address[] memory participants,
    uint256[] memory weights,
    uint256 weightThresholdForBlessing,
    uint256 weightThresholdForBadSignal
  ) external;
}
