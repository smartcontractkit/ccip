// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {BaseOffRampInterface} from "../../interfaces/offRamp/BaseOffRampInterface.sol";
import {CommitStoreInterface} from "../../interfaces/CommitStoreInterface.sol";
import {OCR2Base} from "../../ocr/OCR2Base.sol";
import {BaseOffRamp} from "../BaseOffRamp.sol";
import {CCIP} from "../../models/Models.sol";
import {IERC20} from "../../../vendor/IERC20.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {PoolInterface} from "../../interfaces/pools/PoolInterface.sol";
import {EVM2EVMGEOffRampInterface} from "../../interfaces/offRamp/EVM2EVMGEOffRampInterface.sol";

/**
 * @notice EVM2EVMGEOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMGEOffRamp is EVM2EVMGEOffRampInterface, BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using CCIP for CCIP.EVM2EVMGEMessage;
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMGEOffRamp 1.0.0";

  bytes32 internal immutable i_metadataHash;

  mapping(address => uint256) internal s_nopBalance;
  mapping(address => uint64) internal s_senderNonce;

  GEOffRampConfig internal s_config;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    GEOffRampConfig memory offRampConfig,
    address onRampAddress,
    CommitStoreInterface commitStore,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
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
    i_metadataHash = _metadataHash(CCIP.EVM_2_EVM_GE_MESSAGE_HASH);
  }

  /// @inheritdoc BaseOffRamp
  function manuallyExecute(CCIP.ExecutionReport memory report) external override {
    _execute(report, true);
  }

  /// @inheritdoc EVM2EVMGEOffRampInterface
  function getSenderNonce(address sender) public view returns (uint256 nonce) {
    return s_senderNonce[sender];
  }

  /// @inheritdoc EVM2EVMGEOffRampInterface
  function getNopBalance(address nop) public view returns (uint256 balance) {
    return s_nopBalance[nop];
  }

  function _executeMessages(CCIP.ExecutionReport memory report, bool manualExecution) internal {
    // Report may have only price updates, so we only process messages if there are some.
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) {
      return;
    }

    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    CCIP.EVM2EVMGEMessage[] memory decodedMessages = new CCIP.EVM2EVMGEMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMGEMessage memory decodedMessage = abi.decode(report.encodedMessages[i], (CCIP.EVM2EVMGEMessage));
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      hashedLeaves[i] = decodedMessage._hash(i_metadataHash);
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

    // Execute messages
    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMGEMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState originalState = getExecutionState(message.sequenceNumber);
      if (originalState == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);

      // Two valid cases here, we either have never touched this message before, or we tried to execute and failed
      if (manualExecution) {
        // Manually execution is fine if we previously failed or if the commit report is just too old
        if (!(isOldCommitReport || originalState == CCIP.MessageExecutionState.FAILURE))
          revert ManualExecutionNotYetEnabled();
      } else {
        // DON can only execute a message once
        if (originalState != CCIP.MessageExecutionState.UNTOUCHED) revert AlreadyAttempted(message.sequenceNumber);
      }

      _isWellFormed(message);

      // If this is the first time executing this message we take the fee
      if (originalState == CCIP.MessageExecutionState.UNTOUCHED) {
        if (s_senderNonce[message.sender] + 1 != message.nonce) revert IncorrectNonce(message.nonce);

        // Take the fee charged to this contract.
        // _releaseOrMintToken converts the message.feeToken to the proper destination token
        _releaseOrMintToken(_getPool(IERC20(message.feeToken)), message.feeTokenAmount, address(this));
      }

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
      CCIP.MessageExecutionState newState = _trialExecute(_toAny2EVMMessageFromSender(message));
      s_executedMessages[message.sequenceNumber] = newState;

      if (!(message.strict && newState == CCIP.MessageExecutionState.FAILURE)) {
        s_senderNonce[message.sender]++;
      }

      emit ExecutionStateChanged(message.sequenceNumber, newState);
    }
  }

  /**
   * @notice Execute a series of one or more messages using a merkle proof and update one or more
   * gasFeeCache prices.
   * @param report ExecutionReport
   * @param manualExecution Whether the DON auto executes or it is manually initiated
   */
  function _execute(CCIP.ExecutionReport memory report, bool manualExecution) internal whenNotPaused whenHealthy {
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

  function _toAny2EVMMessageFromSender(CCIP.EVM2EVMGEMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory)
  {
    uint256 numberOfTokens = original.tokensAndAmounts.length;
    CCIP.EVMTokenAndAmount[] memory destTokensAndAmounts = new CCIP.EVMTokenAndAmount[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = _getPool(IERC20(original.tokensAndAmounts[i].token));
      destPools[i] = address(pool);
      destTokensAndAmounts[i] = CCIP.EVMTokenAndAmount({
        token: address(pool.getToken()),
        amount: original.tokensAndAmounts[i].amount
      });
    }

    return
      CCIP.Any2EVMMessageFromSender({
        sourceChainId: original.sourceChainId,
        sender: abi.encode(original.sender),
        receiver: original.receiver,
        data: original.data,
        destTokensAndAmounts: destTokensAndAmounts,
        destPools: destPools,
        gasLimit: original.gasLimit
      });
  }

  function _isWellFormed(CCIP.EVM2EVMGEMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokensAndAmounts.length > uint256(s_config.maxTokensLength))
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  /// @inheritdoc EVM2EVMGEOffRampInterface
  function getGEConfig() external view returns (GEOffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc EVM2EVMGEOffRampInterface
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
    _execute(abi.decode(report, (CCIP.ExecutionReport)), false);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
