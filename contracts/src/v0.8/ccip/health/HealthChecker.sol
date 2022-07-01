// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/Pausable.sol";
import "../health/interfaces/AFNInterface.sol";
import "../access/OwnerIsCreator.sol";

contract HealthChecker is Pausable, OwnerIsCreator {
  // AFN contract to check health of the system
  AFNInterface private s_afn;
  // The maximum time since the last AFN heartbeat before it is considered unhealthy
  uint256 private s_maxSecondsWithoutAFNHeartbeat;

  error BadAFNSignal();
  error StaleAFNHeartbeat();
  error BadHealthConfig();

  event AFNSet(AFNInterface oldAFN, AFNInterface newAFN);
  event AFNMaxHeartbeatTimeSet(uint256 oldTime, uint256 newTime);

  /**
   * @param afn The AFN contract to check health
   * @param maxSecondsWithoutAFNHeartbeat maximum seconds allowed between heartbeats to consider
   * the network "healthy".
   */
  constructor(AFNInterface afn, uint256 maxSecondsWithoutAFNHeartbeat) {
    if (address(afn) == address(0) || maxSecondsWithoutAFNHeartbeat == 0) revert BadHealthConfig();
    s_afn = afn;
    s_maxSecondsWithoutAFNHeartbeat = maxSecondsWithoutAFNHeartbeat;
  }

  /**
   * @notice Pause the contract
   * @dev only callable by the owner
   */
  function pause() external onlyOwner {
    _pause();
  }

  /**
   * @notice Unpause the contract
   * @dev only callable by the owner
   */
  function unpause() external onlyOwner {
    _unpause();
  }

  /**
   * @notice Change the afn contract to track
   * @dev only callable by the owner
   * @param afn new AFN contract
   */
  function setAFN(AFNInterface afn) external onlyOwner {
    if (address(afn) == address(0)) revert BadHealthConfig();
    AFNInterface old = s_afn;
    s_afn = afn;
    emit AFNSet(old, afn);
  }

  /**
   * @notice Get the current AFN contract
   * @return Current AFN
   */
  function getAFN() external view returns (AFNInterface) {
    return s_afn;
  }

  /**
   * @notice Change the maximum time allowed without a heartbeat
   * @dev only callable by the owner
   * @param newTime the new max time
   */
  function setMaxSecondsWithoutAFNHeartbeat(uint256 newTime) external onlyOwner {
    if (newTime == 0) revert BadHealthConfig();
    uint256 oldTime = s_maxSecondsWithoutAFNHeartbeat;
    s_maxSecondsWithoutAFNHeartbeat = newTime;
    emit AFNMaxHeartbeatTimeSet(oldTime, newTime);
  }

  /**
   * @notice Get the current max time without heartbeat
   * @return current max time
   */
  function getMaxSecondsWithoutAFNHeartbeat() external view returns (uint256) {
    return s_maxSecondsWithoutAFNHeartbeat;
  }

  /**
   * @notice Support querying whether health checker is healthy.
   */
  function isHealthy(uint256 timeNow) external view returns (bool) {
    return !s_afn.hasBadSignal() && ((timeNow - s_afn.getLastHeartbeat().timestamp) <= s_maxSecondsWithoutAFNHeartbeat);
  }

  /**
   * @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
   */
  modifier whenHealthy() {
    if (s_afn.hasBadSignal()) revert BadAFNSignal();
    AFNInterface.Heartbeat memory lastHeartbeat = s_afn.getLastHeartbeat();
    if ((block.timestamp - uint256(lastHeartbeat.timestamp)) > s_maxSecondsWithoutAFNHeartbeat)
      revert StaleAFNHeartbeat();
    _;
  }
}
