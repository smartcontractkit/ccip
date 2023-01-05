// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IAFN} from "../interfaces/health/IAFN.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

import {Pausable} from "../../vendor/Pausable.sol";

contract HealthChecker is Pausable, OwnerIsCreator {
  // AFN contract to check health of the system
  IAFN internal s_afn;

  error BadAFNSignal();
  error BadHealthConfig();

  event AFNSet(IAFN oldAFN, IAFN newAFN);

  /**
   * @param afn The AFN contract to check health
   */
  constructor(IAFN afn) {
    if (address(afn) == address(0)) revert BadHealthConfig();
    s_afn = afn;
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
  function setAFN(IAFN afn) external onlyOwner {
    if (address(afn) == address(0)) revert BadHealthConfig();
    IAFN old = s_afn;
    s_afn = afn;
    emit AFNSet(old, afn);
  }

  /**
   * @notice Get the current AFN contract
   * @return Current AFN
   */
  function getAFN() external view returns (IAFN) {
    return s_afn;
  }

  /**
   * @notice Support querying whether health checker is healthy.
   */
  function isAFNHealthy() external view returns (bool) {
    return !s_afn.badSignalReceived();
  }

  /**
   * @notice Ensure that the AFN has not emitted a bad signal, and that the latest heartbeat is not stale.
   */
  modifier whenHealthy() {
    if (s_afn.badSignalReceived()) revert BadAFNSignal();
    _;
  }
}
