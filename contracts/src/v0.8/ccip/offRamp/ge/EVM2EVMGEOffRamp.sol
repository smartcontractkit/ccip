// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {IBaseOffRamp} from "../../interfaces/offRamp/IBaseOffRamp.sol";
import {ICommitStore} from "../../interfaces/ICommitStore.sol";
import {IAFN} from "../../interfaces/health/IAFN.sol";
import {IPool} from "../../interfaces/pools/IPool.sol";
import {IEVM2EVMGEOffRamp} from "../../interfaces/offRamp/IEVM2EVMGEOffRamp.sol";

import {GE} from "../../models/GE.sol";
import {Common} from "../../models/Common.sol";
import {GEConsumer} from "../../models/GEConsumer.sol";
import {Internal} from "../../models/Internal.sol";
import {OCR2Base} from "../../ocr/OCR2Base.sol";
import {BaseOffRamp} from "../BaseOffRamp.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

/**
 * @notice EVM2EVMGEOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMGEOffRamp is IEVM2EVMGEOffRamp, BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMGEOffRamp 1.0.0";

  bytes32 internal immutable i_metadataHash;
  IERC20 internal immutable i_feeToken;

  mapping(address => uint256) internal s_nopBalance;
  mapping(address => uint64) internal s_senderNonce;

  GEOffRampConfig internal s_config;

  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    GEOffRampConfig memory offRampConfig,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin,
    IERC20 feeToken
  )
    OCR2Base(true)
    BaseOffRamp(
      sourceChainId,
      chainId,
      onRampAddress,
      commitStore,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig,
      tokenLimitsAdmin
    )
  {
    s_config = offRampConfig;
    i_metadataHash = _metadataHash(GE.EVM_2_EVM_GE_MESSAGE_HASH);
    i_feeToken = feeToken;
  }

  function manuallyExecute(GE.ExecutionReport memory report) external {
    _execute(report, true);
  }

  /// @inheritdoc IEVM2EVMGEOffRamp
  function getSenderNonce(address sender) public view returns (uint64 nonce) {
    return s_senderNonce[sender];
  }

  /// @inheritdoc IEVM2EVMGEOffRamp
  function getNopBalance(address nop) public view returns (uint256 balance) {
    return s_nopBalance[nop];
  }

  function _executeMessages(GE.ExecutionReport memory report, bool manualExecution) internal {
    // Report may have only price updates, so we only process messages if there are some.
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) {
      return;
    }

    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    GE.EVM2EVMGEMessage[] memory decodedMessages = new GE.EVM2EVMGEMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      GE.EVM2EVMGEMessage memory decodedMessage = abi.decode(report.encodedMessages[i], (GE.EVM2EVMGEMessage));
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      hashedLeaves[i] = GE._hash(decodedMessage, i_metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    (uint256 timestampCommitted, ) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );
    bool isOldCommitReport = (block.timestamp - timestampCommitted) > s_config.permissionLessExecutionThresholdSeconds;

    uint256 totalFeesAccrued = 0;

    // Execute messages
    for (uint256 i = 0; i < numMsgs; ++i) {
      GE.EVM2EVMGEMessage memory message = decodedMessages[i];
      Internal.MessageExecutionState originalState = getExecutionState(message.sequenceNumber);
      if (originalState == Internal.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);

      // Two valid cases here, we either have never touched this message before, or we tried to execute and failed
      if (manualExecution) {
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
        totalFeesAccrued += message.feeTokenAmount;
      }

      _isWellFormed(message);

      s_executedMessages[message.sequenceNumber] = Internal.MessageExecutionState.IN_PROGRESS;
      Internal.MessageExecutionState newState = _trialExecute(_toAny2EVMMessageFromSender(message), manualExecution);
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

    // Take the fee charged to this contract.
    _releaseOrMintToken(getPoolByDestToken(i_feeToken), totalFeesAccrued, address(this));
  }

  /**
   * @notice Execute a series of one or more messages using a merkle proof and update one or more
   * gasFeeCache prices.
   * @param report ExecutionReport
   * @param manualExecution Whether the DON auto executes or it is manually initiated
   */
  function _execute(GE.ExecutionReport memory report, bool manualExecution) internal whenNotPaused whenHealthy {
    uint256 gasStart = gasleft();

    if (address(s_router) == address(0)) revert RouterNotSet();

    // Fee updates
    if (report.feeUpdates.length != 0) {
      if (manualExecution) revert UnauthorizedGasPriceUpdate();
      s_config.gasFeeCache.updateFees(report.feeUpdates);
    }

    // Message execution
    _executeMessages(report, manualExecution);

    // Update NOP balances
    if (!manualExecution) {
      s_nopBalance[msg.sender] +=
        ((gasStart - gasleft() + s_config.gasOverhead) * tx.gasprice * report.tokenPerFeeCoin[0]) /
        1 ether;
    }
  }

  function _toAny2EVMMessageFromSender(GE.EVM2EVMGEMessage memory original)
    internal
    view
    returns (Internal.Any2EVMMessageFromSender memory)
  {
    uint256 numberOfTokens = original.tokensAndAmounts.length;
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      IPool pool = _getPool(IERC20(original.tokensAndAmounts[i].token));
      destPools[i] = address(pool);
      destTokensAndAmounts[i] = Common.EVMTokenAndAmount({
        token: address(pool.getToken()),
        amount: original.tokensAndAmounts[i].amount
      });
    }

    return
      Internal.Any2EVMMessageFromSender({
        sourceChainId: original.sourceChainId,
        sender: abi.encode(original.sender),
        receiver: original.receiver,
        data: original.data,
        destTokensAndAmounts: destTokensAndAmounts,
        destPools: destPools,
        gasLimit: original.gasLimit
      });
  }

  function _isWellFormed(GE.EVM2EVMGEMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokensAndAmounts.length > uint256(s_config.maxTokensLength))
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  /// @inheritdoc IEVM2EVMGEOffRamp
  function getGEConfig() external view returns (GEOffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc IEVM2EVMGEOffRamp
  function setGEConfig(GEOffRampConfig memory config) external onlyOwner {
    s_config = config;

    emit GEOffRampConfigChanged(config);
  }

  // ******* OCR BASE ***********
  /**
   * @notice Entry point for execution, called by the OCR network
   * @dev Expects an encoded ExecutionReport
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override {
    _execute(abi.decode(report, (GE.ExecutionReport)), false);
  }

  function _beforeSetOCR2Config(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetOCR2Config(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
