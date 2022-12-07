// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Common} from "../../models/Common.sol";
import {GE} from "../../models/GE.sol";

interface GasFeeCacheInterface {
  error ChainNotSupported(uint64 chain);
  error FeeUpdaterNotAllowed(address feeUpdater);
  error OnlyCallableByUpdaterOrOwner();

  event FeeUpdaterSet(address feeUpdater);
  event FeeUpdaterRemoved(address feeUpdater);
  event GasFeeUpdated(uint64 destChain, uint256 linkPerUnitGas);

  function setFeeUpdater(address feeUpdater) external;

  function removeFeeUpdater(address feeUpdater) external;

  function updateFees(GE.FeeUpdate[] memory feeUpdates) external;

  function getFee(uint64 destChainId) external view returns (uint256 fee);
}
