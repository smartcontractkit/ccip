// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../../models/Common.sol";
import {GE} from "../../models/GE.sol";

interface IFeeManager {
  error ChainNotSupported(uint64 chain);
  error FeeUpdaterNotAllowed(address feeUpdater);
  error OnlyCallableByUpdaterOrOwner();
  error StaleFee(uint256 threshold, uint256 timePassed);

  event FeeUpdaterSet(address feeUpdater);
  event FeeUpdaterRemoved(address feeUpdater);
  event GasFeeUpdated(uint64 destChain, uint128 linkPerUnitGas, uint128 timestamp);

  struct TimestampedFeeUpdate {
    uint128 linkPerUnitGas;
    uint128 timestamp;
  }

  function setFeeUpdater(address feeUpdater) external;

  function removeFeeUpdater(address feeUpdater) external;

  function updateFees(GE.FeeUpdate[] memory feeUpdates) external;

  function getFee(uint64 destChainId) external view returns (uint128 fee);
}
