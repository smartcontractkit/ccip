// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {Any2EVMSubscriptionOffRampRouter} from "./Any2EVMSubscriptionOffRampRouter.sol";
import {OCR2Base} from "../../ocr/OCR2Base.sol";
import {BaseOffRamp} from "../BaseOffRamp.sol";
import {CCIP} from "../../models/Models.sol";
import {IERC20} from "../../../vendor/IERC20.sol";
import {PoolInterface} from "../../interfaces/pools/PoolInterface.sol";
import {BlobVerifierInterface} from "../../interfaces/BlobVerifierInterface.sol";
import {AFNInterface} from "../../interfaces/health/AFNInterface.sol";

/**
 * @notice EVM2EVMSubscriptionOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMFreeOffRamp is BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using CCIP for CCIP.EVM2EVMSubscriptionMessage;
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2EVMSubscriptionOffRamp 1.0.0";

  mapping(address => uint64) internal s_receiverToNonce;

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

  function execute(CCIP.ExecutionReport memory report, bool manualExecution)
    external
    override
    whenNotPaused
    whenHealthy
  {
    if (address(s_router) == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();

    CCIP.EVM2EVMSubscriptionMessage[] memory decodedMessages = new CCIP.EVM2EVMSubscriptionMessage[](numMsgs);
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    // TODO optimise gas cost of hashing/caching hash
    bytes32 metadataHash = _metadataHash(CCIP.EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH);
    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMSubscriptionMessage memory decodedMessage = abi.decode(
        report.encodedMessages[i],
        (CCIP.EVM2EVMSubscriptionMessage)
      );
      // We do this hash here instead of in _verifyMessages to avoid two separate loops
      // over the same data, which increases gas cost
      // TODO: golf check
      hashedLeaves[i] = decodedMessage._hash(metadataHash);
      decodedMessages[i] = decodedMessage;
    }

    // TODO: Spec difference measuring gas used by verification vs calculating it?
    // imo billing calculated values > billing measured to help with cost predictability
    // solhint-disable-next-line no-unused-vars
    (uint256 timestampRelayed,) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );

    // only allow manual execution if the report is old enough
    if (manualExecution && (block.timestamp - timestampRelayed) < s_config.permissionLessExecutionThresholdSeconds)
      revert ManualExecutionNotYetEnabled();

    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMSubscriptionMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);

      // Any message with a nonce that is n + 1 is allowed.
      bool isNextInSequence = s_receiverToNonce[message.receiver] + 1 == message.nonce;
      if (!(isNextInSequence || state == CCIP.MessageExecutionState.FAILURE)) revert IncorrectNonce(message.nonce);

      _isWellFormed(message);

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
      CCIP.MessageExecutionState newState = _trialExecute(_toAny2EVMMessageFromSender(message));
      s_executedMessages[message.sequenceNumber] = newState;
      emit ExecutionStateChanged(message.sequenceNumber, newState);

      // Increment the nonce of the receiver if it's the next nonce in line and it was successfully executed .
      if (isNextInSequence) {
        s_receiverToNonce[message.receiver]++;
      }
    }
  }

  function _toAny2EVMMessageFromSender(CCIP.EVM2EVMSubscriptionMessage memory original)
    internal
    view
    returns (CCIP.Any2EVMMessageFromSender memory message)
  {
    uint256 numberOfTokens = original.tokens.length;
    address[] memory destTokens = new address[](numberOfTokens);
    address[] memory destPools = new address[](numberOfTokens);

    for (uint256 i = 0; i < numberOfTokens; ++i) {
      PoolInterface pool = _getPool(IERC20(original.tokens[i]));
      destPools[i] = address(pool);
      destTokens[i] = address(pool.getToken());
    }

    message = CCIP.Any2EVMMessageFromSender({
      sourceChainId: original.sourceChainId,
      sender: abi.encode(original.sender),
      receiver: original.receiver,
      data: original.data,
      destTokens: destTokens,
      destPools: destPools,
      amounts: original.amounts,
      gasLimit: original.gasLimit
    });
  }

  function getNonce(address receiver) external view returns (uint64) {
    return s_receiverToNonce[receiver];
  }

  function _isWellFormed(CCIP.EVM2EVMSubscriptionMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length)
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
    this.execute(abi.decode(report, (CCIP.ExecutionReport)), false);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
