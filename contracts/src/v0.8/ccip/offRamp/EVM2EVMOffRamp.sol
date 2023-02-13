// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {ICommitStore} from "../interfaces/ICommitStore.sol";
import {IFeeManager} from "../interfaces/fees/IFeeManager.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IEVM2EVMOffRamp} from "../interfaces/offRamp/IEVM2EVMOffRamp.sol";
import {IRouter} from "../interfaces/router/IRouter.sol";
import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";

import {Internal} from "../models/Internal.sol";
import {Common} from "../models/Common.sol";
import {Consumer} from "../models/Consumer.sol";
import {Internal} from "../models/Internal.sol";
import {OCR2Base} from "../ocr/OCR2Base.sol";
import {HealthChecker} from "../health/HealthChecker.sol";
import {OffRampTokenPoolRegistry} from "../pools/OffRampTokenPoolRegistry.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {Address} from "../../vendor/Address.sol";
import {ERC165Checker} from "../../vendor/ERC165Checker.sol";

/// @notice EVM2EVMOffRamp enables OCR networks to execute multiple messages
/// in an OffRamp in a single transaction.
contract EVM2EVMOffRamp is
  IEVM2EVMOffRamp,
  HealthChecker,
  OffRampTokenPoolRegistry,
  AggregateRateLimiter,
  TypeAndVersionInterface,
  OCR2Base
{
  using Address for address;
  using ERC165Checker for address;

  // STATIC CONFIG
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMOffRamp 1.0.0";
  // Chain ID of the source chain
  uint64 internal immutable i_sourceChainId;
  // Chain ID of this chain
  uint64 internal immutable i_chainId;
  // OnRamp address on the source chain
  address internal immutable i_onRampAddress;
  // metadataHash is a prefix for a message hash preimage to ensure uniqueness.
  bytes32 internal immutable i_metadataHash;

  // DYNAMIC CONFIG
  OffRampConfig internal s_config;

  // STATE
  mapping(address => uint64) internal s_senderNonce;
  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => Internal.MessageExecutionState) internal s_executedMessages;

  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    address onRampAddress,
    OffRampConfig memory offRampConfig,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  )
    OCR2Base()
    HealthChecker(afn)
    OffRampTokenPoolRegistry(sourceTokens, pools)
    AggregateRateLimiter(rateLimiterConfig)
  {
    if (onRampAddress == address(0)) revert ZeroAddressNotAllowed();
    _setOffRampConfig(offRampConfig);

    i_sourceChainId = sourceChainId;
    i_chainId = chainId;
    i_onRampAddress = onRampAddress;
    i_metadataHash = _metadataHash(Internal.EVM_2_EVM_MESSAGE_HASH);
  }

  function _metadataHash(bytes32 prefix) internal view returns (bytes32) {
    return keccak256(abi.encode(prefix, i_sourceChainId, i_chainId, i_onRampAddress));
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function getOffRampConfig() external view override returns (OffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function setOffRampConfig(OffRampConfig memory config) external override onlyOwner {
    _setOffRampConfig(config);
  }

  function _setOffRampConfig(OffRampConfig memory config) private {
    if (config.router == address(0) || config.commitStore == address(0) || config.feeManager == address(0))
      revert InvalidOffRampConfig(config);

    s_config = config;
    emit OffRampConfigChanged(config);
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function getExecutionState(uint64 sequenceNumber) public view returns (Internal.MessageExecutionState) {
    return s_executedMessages[sequenceNumber];
  }

  function getChainIDs() external view returns (uint64 sourceChainId, uint64 chainId) {
    sourceChainId = i_sourceChainId;
    chainId = i_chainId;
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function getSenderNonce(address sender) public view override returns (uint64 nonce) {
    return s_senderNonce[sender];
  }

  /// @notice Uses the pool to release or mint tokens and send them to
  ///         the given `receiver` address.
  function _releaseOrMintToken(
    IPool pool,
    uint256 amount,
    address receiver
  ) internal {
    pool.releaseOrMint(receiver, amount);
  }

  /// @notice Uses pools to release or mint a number of different tokens
  ///           and send them to the given `receiver` address.
  function _releaseOrMintTokens(Common.EVMTokenAndAmount[] memory sourceTokensAndAmounts, address receiver)
    internal
    returns (Common.EVMTokenAndAmount[] memory)
  {
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](
      sourceTokensAndAmounts.length
    );
    for (uint256 i = 0; i < sourceTokensAndAmounts.length; ++i) {
      IPool pool = getPoolBySourceToken(IERC20(sourceTokensAndAmounts[i].token));
      if (address(pool) == address(0)) revert UnsupportedToken(IERC20(sourceTokensAndAmounts[i].token));
      _releaseOrMintToken(pool, sourceTokensAndAmounts[i].amount, receiver);
      destTokensAndAmounts[i].token = address(pool.getToken());
      destTokensAndAmounts[i].amount = sourceTokensAndAmounts[i].amount;
    }
    _removeTokens(destTokensAndAmounts);
    return destTokensAndAmounts;
  }

  /// @notice Execute a single message
  /// @param message The Any2EVMMessageFromSender message that will be executed
  /// @param manualExecution bool to indicate manual instead of DON execution
  /// @dev this can only be called by the contract itself. It is part of
  /// the Execute call, as we can only try/catch on external calls.
  function executeSingleMessage(Internal.EVM2EVMMessage memory message, bool manualExecution) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](0);
    if (message.tokensAndAmounts.length > 0) {
      destTokensAndAmounts = _releaseOrMintTokens(message.tokensAndAmounts, message.receiver);
    }
    if (
      !message.receiver.isContract() || !message.receiver.supportsInterface(type(IAny2EVMMessageReceiver).interfaceId)
    ) return;
    if (
      !IRouter(s_config.router).routeMessage(
        Internal._toAny2EVMMessage(message, destTokensAndAmounts),
        manualExecution,
        message.gasLimit,
        message.receiver
      )
    ) revert ReceiverError();
  }

  /// @notice Try executing a message
  /// @param message Common.Any2EVMMessage memory message
  /// @param manualExecution bool to indicate manual instead of DON execution
  /// @return Internal.ExecutionState
  function _trialExecute(Internal.EVM2EVMMessage memory message, bool manualExecution)
    internal
    returns (Internal.MessageExecutionState)
  {
    try this.executeSingleMessage(message, manualExecution) {} catch (bytes memory err) {
      if (IEVM2EVMOffRamp.ReceiverError.selector == bytes4(err)) {
        return Internal.MessageExecutionState.FAILURE;
      } else {
        revert ExecutionError(err);
      }
    }
    return Internal.MessageExecutionState.SUCCESS;
  }

  function _isWellFormed(Internal.EVM2EVMMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokensAndAmounts.length > uint256(s_config.maxTokensLength))
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  function _executeMessages(Internal.ExecutionReport memory report, bool manualExecution) internal {
    // Report may have only price updates, so we only process messages if there are some.
    uint256 numMsgs = report.encodedMessages.length;
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    Internal.EVM2EVMMessage[] memory decodedMessages = new Internal.EVM2EVMMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      Internal.EVM2EVMMessage memory decodedMessage = abi.decode(report.encodedMessages[i], (Internal.EVM2EVMMessage));
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      hashedLeaves[i] = Internal._hash(decodedMessage, i_metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    // SECURITY CRITICAL CHECK
    uint256 timestampCommitted = ICommitStore(s_config.commitStore).verify(
      hashedLeaves,
      report.proofs,
      report.proofFlagBits
    );
    if (timestampCommitted <= 0) revert RootNotCommitted();

    // Execute messages
    for (uint256 i = 0; i < numMsgs; ++i) {
      Internal.EVM2EVMMessage memory message = decodedMessages[i];
      Internal.MessageExecutionState originalState = getExecutionState(message.sequenceNumber);
      // Two valid cases here, we either have never touched this message before, or we tried to execute
      // and failed. This check protects against reentry and re-execution because the other states are
      // IN_PROGRESS and SUCCESS, both should not be allowed to execute.
      if (
        !(originalState == Internal.MessageExecutionState.UNTOUCHED ||
          originalState == Internal.MessageExecutionState.FAILURE)
      ) revert AlreadyExecuted(message.sequenceNumber);

      if (manualExecution) {
        bool isOldCommitReport = (block.timestamp - timestampCommitted) >
          s_config.permissionLessExecutionThresholdSeconds;
        // Manually execution is fine if we previously failed or if the commit report is just too old
        // Acceptable state transitions: FAILURE->SUCCESS, UNTOUCHED->SUCCESS, FAILURE->FAILURE
        if (!(isOldCommitReport || originalState == Internal.MessageExecutionState.FAILURE))
          revert ManualExecutionNotYetEnabled();
      } else {
        // DON can only execute a message once
        // Acceptable state transitions: UNTOUCHED->SUCCESS, UNTOUCHED->FAILURE
        if (originalState != Internal.MessageExecutionState.UNTOUCHED) revert AlreadyAttempted(message.sequenceNumber);
      }

      // If this is the first time executing this message we take the fee
      if (originalState == Internal.MessageExecutionState.UNTOUCHED) {
        // UNTOUCHED messages MUST be executed in order always.
        if (s_senderNonce[message.sender] + 1 != message.nonce) {
          // We skip the message if the nonce is incorrect
          emit SkippedIncorrectNonce(message.nonce, message.sender);
          continue;
        }
      }

      _isWellFormed(message);

      s_executedMessages[message.sequenceNumber] = Internal.MessageExecutionState.IN_PROGRESS;
      Internal.MessageExecutionState newState = _trialExecute(message, manualExecution);
      s_executedMessages[message.sequenceNumber] = newState;

      if (manualExecution) {
        // Nonce changes per state transition:
        // FAILURE->SUCCESS: no nonce bump unless strict
        // UNTOUCHED->SUCCESS: nonce bump
        // FAILURE->FAILURE: no nonce bump
        if (
          (message.strict &&
            originalState == Internal.MessageExecutionState.FAILURE &&
            newState == Internal.MessageExecutionState.SUCCESS) ||
          (originalState == Internal.MessageExecutionState.UNTOUCHED &&
            newState == Internal.MessageExecutionState.SUCCESS)
        ) {
          s_senderNonce[message.sender]++;
        }
      } else {
        // Nonce changes per state transition:
        // UNTOUCHED->SUCCESS: nonce bump
        // UNTOUCHED->FAILURE: nonce bump unless strict
        if (!(message.strict && newState == Internal.MessageExecutionState.FAILURE)) {
          s_senderNonce[message.sender]++;
        }
      }

      emit ExecutionStateChanged(message.sequenceNumber, message.messageId, newState);
    }
  }

  /// @notice Execute a series of one or more messages using a merkle proof and update one or more
  /// feeManager prices.
  /// @param report ExecutionReport
  /// @param manualExecution Whether the DON auto executes or it is manually initiated
  function _execute(Internal.ExecutionReport memory report, bool manualExecution) internal whenNotPaused whenHealthy {
    // Fee updates
    if (report.feeUpdates.length != 0) {
      if (manualExecution) revert UnauthorizedGasPriceUpdate();
      IFeeManager(s_config.feeManager).updateFees(report.feeUpdates);
    }

    // Messages execution
    if (report.encodedMessages.length != 0) {
      _executeMessages(report, manualExecution);
    }
  }

  /// @inheritdoc IEVM2EVMOffRamp
  function manuallyExecute(Internal.ExecutionReport memory report) external override {
    _execute(report, true);
  }

  /// @notice Reverts as this contract should not access CCIP messages
  function ccipReceive(Common.Any2EVMMessage calldata) external pure {
    // solhint-disable-next-line reason-string
    revert();
  }

  // ******* OCR BASE ***********
  ///
  /// @notice Entry point for execution, called by the OCR network
  /// @dev Expects an encoded ExecutionReport
  ///
  function _report(bytes memory report) internal override {
    _execute(abi.decode(report, (Internal.ExecutionReport)), false);
  }

  function _beforeSetOCR2Config(uint8 f, bytes memory onchainConfig) internal override {}
}
