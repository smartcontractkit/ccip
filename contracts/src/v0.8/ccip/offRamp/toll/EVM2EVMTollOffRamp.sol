// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {BaseOffRampInterface} from "../../interfaces/offRamp/BaseOffRampInterface.sol";
import {BlobVerifierInterface} from "../../interfaces/BlobVerifierInterface.sol";
import {OCR2Base} from "../../ocr/OCR2Base.sol";
import {BaseOffRamp} from "../BaseOffRamp.sol";
import {CCIP} from "../../models/Models.sol";
import {IERC20} from "../../../vendor/IERC20.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";
import {PoolInterface} from "../../interfaces/pools/PoolInterface.sol";

/**
 * @notice EVM2EVMTollOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMTollOffRamp is BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using CCIP for CCIP.EVM2EVMTollMessage;
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMTollOffRamp 1.0.0";
  uint256 private constant TOLL_CONSTANT_MESSAGE_PART_BYTES = (20 + // receiver
    20 + // sender
    2 + // chain id
    8 + // sequence number
    32 + // gas limit
    20 + // fee token address
    32); // fee token amount
  uint256 private constant TOLL_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS = (2_100 + // COLD_SLOAD_COST for first reading the state
    20_000 + // SSTORE_SET_GAS for writing from 0 (untouched) to non-zero (in-progress)
    100); // SLOAD_GAS = WARM_STORAGE_READ_COST for rewriting from non-zero (in-progress) to non-zero (success/failure)

  mapping(uint256 => uint256) public feeTaken;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
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
      offRampConfig,
      blobVerifier,
      afn,
      sourceTokens,
      pools,
      rateLimiterConfig,
      tokenLimitsAdmin
    )
  {}

  /**
   * @notice Compute the overhead gas for a given message given its share of the merkle root verification costs.
   * We need to compute this to bill the user upfront so we can let them know how much of a refund they get.
   */
  function overheadGasToll(uint256 merkleGasShare, CCIP.EVM2EVMTollMessage memory message)
    public
    pure
    returns (uint256)
  {
    uint256 messageBytes = (TOLL_CONSTANT_MESSAGE_PART_BYTES +
      (EVM_ADDRESS_LENGTH_BYTES + EVM_WORD_BYTES) *
      message.tokensAndAmounts.length +
      message.data.length);
    uint256 messageCalldataGas = messageBytes * CALLDATA_GAS_PER_BYTE;
    return (messageCalldataGas +
      merkleGasShare +
      TOLL_EXECUTION_STATE_PROCESSING_OVERHEAD_GAS +
      PER_TOKEN_OVERHEAD_GAS *
      (message.tokensAndAmounts.length + 1) +
      RATE_LIMITER_OVERHEAD_GAS +
      EXTERNAL_CALL_OVERHEAD_GAS);
  }

  /**
   * @notice Compute the fee for a given message using token prices in report.
   * @dev Reduces stack pressure to have an explicit function for this.
   */
  function _computeFee(
    uint256 merkleGasShare,
    CCIP.ExecutionReport memory report,
    CCIP.EVM2EVMTollMessage memory message
  ) internal view returns (uint256) {
    // Charge the gas share & gas limit of the message multiplied by the token per fee coin for
    // the given message.
    // Example with token being link. 1 LINK = 1e18 Juels.
    // tx.gasprice is wei / gas
    // gas * wei/gas * (juels / wei) (problem is that juels per wei could be < 1, say since 1 link < 1 eth)
    // instead we use juels per unit ETH, which > 1, assuming 1 juel < 1 ETH (safe).
    // gas * wei/gas * (juels / (ETH * 1e18 WEI/ETH))
    // gas * wei/gas * juels/ETH / (1e18 wei/ETH)
    // Example 1e6 gas * (200e9 wei / gas) * (6253149865160030 juels / ETH) / (1e18 wei/ETH) = 1.25e15 juels
    uint256 tokenPerFeeCoin;
    // tokenPerFeeCoinAddresses is keyed in destination chain tokens so we need to convert the feeToken
    // before we do the lookup
    address destinationFeeTokenAddress = address(_getPool(IERC20(message.feeTokenAndAmount.token)).getToken());
    for (uint256 j = 0; j < report.tokenPerFeeCoinAddresses.length; ++j) {
      if (report.tokenPerFeeCoinAddresses[j] == destinationFeeTokenAddress) {
        tokenPerFeeCoin = report.tokenPerFeeCoin[j];
      }
    }
    if (tokenPerFeeCoin == uint256(0)) revert MissingFeeCoinPrice(destinationFeeTokenAddress);
    // Gas cost in wei: gasUsed * gasPrice
    // example: 100k gas, 20 gwei = 1e5 * 20e9  = 2e15
    // Gas cost in token: costInWei * 1e18 / tokenPerFeeCoin
    // example: costInWei 2e15, tokenPerFeeCoin 2e20 = 2e15 * 2e20 / 1e18 = 4e17 tokens
    uint256 feeTokenCharged = ((overheadGasToll(merkleGasShare, message) + message.gasLimit) *
      tx.gasprice *
      tokenPerFeeCoin) / 1 ether;
    if (feeTokenCharged > message.feeTokenAndAmount.amount)
      revert InsufficientFeeAmount(message.sequenceNumber, feeTokenCharged, message.feeTokenAndAmount.amount);
    return feeTokenCharged;
  }

  /// @inheritdoc BaseOffRamp
  function manuallyExecute(CCIP.ExecutionReport memory report) external override {
    _execute(report, true);
  }

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param manualExecution Whether the DON auto executes or it is manually initiated
   */
  function _execute(CCIP.ExecutionReport memory report, bool manualExecution) internal whenNotPaused whenHealthy {
    if (address(s_router) == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();

    CCIP.EVM2EVMTollMessage[] memory decodedMessages = new CCIP.EVM2EVMTollMessage[](numMsgs);
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    bytes32 metadataHash = _metadataHash(CCIP.EVM_2_EVM_TOLL_MESSAGE_HASH);
    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMTollMessage memory decodedMessage = abi.decode(report.encodedMessages[i], (CCIP.EVM2EVMTollMessage));
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      hashedLeaves[i] = decodedMessage._hash(metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    (uint256 timestampRelayed, uint256 gasUsedByMerkle) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );
    bool isOldRelayReport = (block.timestamp - timestampRelayed) > s_config.permissionLessExecutionThresholdSeconds;

    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMTollMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState originalState = getExecutionState(message.sequenceNumber);
      if (originalState == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);

      // Manually execution is fine if we previously failed or if the relay report is just too old
      if (!(!manualExecution || isOldRelayReport || originalState == CCIP.MessageExecutionState.FAILURE))
        revert ManualExecutionNotYetEnabled();

      _isWellFormed(message);

      uint256 feeTokenCharged;
      // If it's the first DON execution attempt, charge the fee.
      if (originalState == CCIP.MessageExecutionState.UNTOUCHED && !manualExecution) {
        feeTokenCharged = _computeFee(gasUsedByMerkle / decodedMessages.length, report, message);
        // Take the fee charged to this contract.
        // _releaseOrMintToken converts the message.feeTokenAndAmount to the proper destination token
        PoolInterface feeTokenPool = _getPool(IERC20(message.feeTokenAndAmount.token));
        _releaseOrMintToken(feeTokenPool, feeTokenCharged, address(this));
        // Forward the refund amount to the user so they know how much they were refunded.
        message.feeTokenAndAmount.amount -= feeTokenCharged;
      }

      if (originalState != CCIP.MessageExecutionState.UNTOUCHED) {
        // We have taken a fee already, remove from message to avoid
        // double-minting.
        message.feeTokenAndAmount.amount -= feeTaken[message.sequenceNumber];
      }

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
      // NOTE: toAny2EVMMessageFromSender merges the fee token into the token set.
      CCIP.MessageExecutionState newState = _trialExecute(_toAny2EVMMessageFromSender(message));
      s_executedMessages[message.sequenceNumber] = newState;

      if (originalState == CCIP.MessageExecutionState.UNTOUCHED && newState == CCIP.MessageExecutionState.FAILURE) {
        feeTaken[message.sequenceNumber] = feeTokenCharged;
      }

      emit ExecutionStateChanged(message.sequenceNumber, newState);
    }
  }

  // @notice IMPORTANT: Merges the fee token into the set of (tokens, amounts)
  function _toAny2EVMMessageFromSender(CCIP.EVM2EVMTollMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory message)
  {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = CCIP._addToTokensAmounts(
      original.tokensAndAmounts,
      original.feeTokenAndAmount
    );
    uint256 numberOfTokens = tokensAndAmounts.length;
    CCIP.EVMTokenAndAmount[] memory destTokensAndAmounts = new CCIP.EVMTokenAndAmount[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = _getPool(IERC20(tokensAndAmounts[i].token));
      destPools[i] = address(pool);
      destTokensAndAmounts[i] = CCIP.EVMTokenAndAmount({
        token: address(pool.getToken()),
        amount: tokensAndAmounts[i].amount
      });
    }

    message = CCIP.Any2EVMMessageFromSender({
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      receiver: original.receiver,
      data: original.data,
      destTokensAndAmounts: destTokensAndAmounts,
      destPools: destPools,
      gasLimit: original.gasLimit
    });
  }

  function _isWellFormed(CCIP.EVM2EVMTollMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokensAndAmounts.length > uint256(s_config.maxTokensLength))
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
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
