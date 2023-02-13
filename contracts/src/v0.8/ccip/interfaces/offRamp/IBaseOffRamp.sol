// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouter} from "../router/IRouter.sol";
import {ICommitStore} from "../ICommitStore.sol";

import {Internal} from "../../models/Internal.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

interface IBaseOffRamp {
  error ZeroAddressNotAllowed();
  error AlreadyExecuted(uint64 sequenceNumber);
  error ExecutionError(bytes error);
  error InvalidSourceChain(uint64 sourceChainId);
  error NoMessagesToExecute();
  error ManualExecutionNotYetEnabled();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error RouterNotSet();
  error RootNotCommitted();
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error TokenAndAmountMisMatch();
  error UnsupportedToken(IERC20 token);
  error CanOnlySelfCall();
  error ReceiverError();
  error MissingFeeCoinPrice(address feeCoin);
  error InsufficientFeeAmount(uint64 sequenceNumber, uint256 expectedFeeTokens, uint256 feeTokenAmount);

  event OffRampRouterSet(address indexed router, uint64 sourceChainId, address onRampAddress);

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner should be able to call this function
   */
  function setRouter(IRouter router) external;

  /**
   * @notice get the current router
   * @return IRouter    */
  function getRouter() external view returns (IRouter);

  /**
   * @notice Returns the current execution state of a message based on its
   *          sequenceNumber.
   */
  function getExecutionState(uint64 sequenceNumber) external view returns (Internal.MessageExecutionState);

  /**
   * @notice Returns the current commitStore.
   */
  function getCommitStore() external view returns (ICommitStore);

  /**
   * @notice Updates the commitStore.
   * @param commitStore The new commitStore
   */
  function setCommitStore(ICommitStore commitStore) external;
}
