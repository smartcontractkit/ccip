// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";

interface GasFeeCacheInterface {
  error ChainNotSupported(uint256 chain);
  error FeeUpdaterNotAllowed(address feeUpdater);
  error OnlyCallableByUpdaterOrOwner();

  event FeeUpdaterSet(address feeUpdater);
  event FeeUpdaterRemoved(address feeUpdater);
  event GasFeeUpdated(uint256 destChain, uint256 linkPerUnitGas);

  function setFeeUpdater(address feeUpdater) external;

  function removeFeeUpdater(address feeUpdater) external;

  function updateFees(CCIP.FeeUpdate[] memory feeUpdates) external;

  function getFee(uint256 destChainId) external returns (uint256 fee);
}
