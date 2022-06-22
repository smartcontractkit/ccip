// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../interfaces/PoolInterface.sol";
import "../interfaces/TollOffRampRouterInterface.sol";
import "../interfaces/CrossChainMessageReceiverInterface.sol";
import "../utils/CCIP.sol";

interface TollOffRampInterface {
  error AlreadyExecuted(uint64 sequenceNumber);
  error CanOnlySelfCall();
  error ExecutionError(uint64 sequenceNumber, bytes reason);
  error InvalidReceiver(address receiver);
  error InvalidSourceChain(uint256 sourceChainId);
  error NoMessagesToExecute();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error RouterNotSet();
  error RootNotRelayed();
  error UnsupportedNumberOfTokens(uint64 sequenceNumber);
  error UnsupportedToken(IERC20 token);
  error MissingFeeCoinPrice(address feeCoin);

  event ExecutionCompleted(uint64 sequenceNumber, CCIP.MessageExecutionState state);
  event OffRampConfigSet(OffRampConfig config);
  event OffRampRouterSet(TollOffRampRouterInterface router);

  struct OffRampConfig {
    // The ID of the source chain
    uint256 sourceChainId;
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
  function setRouter(TollOffRampRouterInterface router) external;

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature/
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata message) external;

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param needFee Whether or not the executor requires a fee
   */
  function execute(CCIP.ExecutionReport memory report, bool needFee) external returns (CCIP.ExecutionResult[] memory);

  /**
   * @notice Execute a single message
   * @param message The Any2EVMTollMessage message that will be executed
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMTollMessage memory message) external;
}
