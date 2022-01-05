// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

interface AFNInterface {
  struct Heartbeat {
    uint256 round;
    uint256 timestamp;
    uint256 committeeVersion;
  }

  event GoodVote(address voter, uint256 round);
  event BadVote(address voter, uint256 round);
  event AFNHeartbeat(Heartbeat heartbeat);
  event AFNBadSignal(uint256 timestamp);
  event RecoveredFromBadSignal();
  event ConfigSet(address[] parties, uint256[] weights, uint256 goodQuorum, uint256 badQuorum);

  error IncorrectRound(uint256 expected, uint256 received);
  error InvalidVoter(address voter);
  error AlreadyVoted();
  error InvalidConfig();
  error InvalidWeight();
  error MustRecoverFromBadSignal();
  error RecoveryNotNecessary();

  function hasBadSignal() external returns (bool);

  function getLastHeartbeat() external returns (Heartbeat memory);

  function voteGood(uint256 round) external;

  function voteBad() external;

  function recover() external;

  function setConfig(
    address[] memory parties,
    uint256[] memory weights,
    uint256 goodQuorum,
    uint256 badQuorum
  ) external;
}
