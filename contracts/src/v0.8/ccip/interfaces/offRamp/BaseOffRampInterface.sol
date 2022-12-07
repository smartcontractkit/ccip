// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Any2EVMOffRampRouterInterface} from "./Any2EVMOffRampRouterInterface.sol";
import {CommitStoreInterface} from "../CommitStoreInterface.sol";
import {Internal} from "../../models/Internal.sol";
import {IERC20} from "../../../vendor/IERC20.sol";

interface BaseOffRampInterface {
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
  event OffRampConfigSet(OffRampConfig config);

  struct OffRampConfig {
    // The waiting time before manual execution is enabled
    uint32 permissionLessExecutionThresholdSeconds;
    // execution delay in seconds
    uint64 executionDelaySeconds;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
  }

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner should be able to call this function
   */
  function setRouter(Any2EVMOffRampRouterInterface router) external;

  /**
   * @notice get the current router
   * @return Any2EVMOffRampRouterInterface
   */
  function getRouter() external view returns (Any2EVMOffRampRouterInterface);

  /**
   * @notice Returns the current execution state of a message based on its
   *          sequenceNumber.
   */
  function getExecutionState(uint64 sequenceNumber) external view returns (Internal.MessageExecutionState);

  /**
   * @notice Returns the current commitStore.
   */
  function getCommitStore() external view returns (CommitStoreInterface);

  /**
   * @notice Updates the commitStore.
   * @param commitStore The new commitStore
   */
  function setCommitStore(CommitStoreInterface commitStore) external;
}
