// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../ocr/OCR2Base.sol";
import "../interfaces/Any2EVMSubscriptionOffRampInterface.sol";
import "../BaseOffRamp.sol";

/**
 * @notice Any2EVMSubscriptionOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract Any2EVMSubscriptionOffRamp is
  Any2EVMSubscriptionOffRampInterface,
  BaseOffRamp,
  TypeAndVersionInterface,
  OCR2Base
{
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "Any2EVMSubscriptionOffRamp 1.0.0";

  // The router through which all transactions will be executed
  Any2EVMSubscriptionOffRampRouterInterface public s_router;

  mapping(address => uint64) public s_receiverToNonce;

  constructor(
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    // OnrampAddress, needed for hashing in the future so already added to the interface
    address onRampAddress,
    AFNInterface afn,
    // TODO token limiter contract
    // https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    uint256 maxTimeWithoutAFNSignal
  )
    OCR2Base(true)
    BaseOffRamp(chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools, maxTimeWithoutAFNSignal)
  {}

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner can call this function
   */
  function setRouter(Any2EVMSubscriptionOffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(address(router));
  }

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature
   */
  function ccipReceive(CCIP.Any2EVMSubscriptionMessage calldata) external pure override {
    revert();
  }

  /**
   * @notice Execute a single message
   * @param message The Any2EVMSubscriptionMessage message that will be executed
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMSubscriptionMessage memory message) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    // TODO: token limiter logic
    // https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
    _releaseOrMintTokens(message.tokens, message.amounts, message.receiver);
    _callReceiver(message);
  }

  /// @inheritdoc Any2EVMSubscriptionOffRampInterface
  function execute(CCIP.ExecutionReport memory report, bool manualExecution)
    external
    override
    whenNotPaused
    whenHealthy
  {
    if (address(s_router) == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    CCIP.Any2EVMSubscriptionMessage[] memory decodedMessages = new CCIP.Any2EVMSubscriptionMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      decodedMessages[i] = abi.decode(report.encodedMessages[i], (CCIP.Any2EVMSubscriptionMessage));
      // TODO: hasher
      // https://app.shortcut.com/chainlinklabs/story/41625/hasher-encoder
      bytes memory data = bytes.concat(hex"00", report.encodedMessages[i]);
      hashedLeaves[i] = keccak256(data);
    }

    (uint256 timestampRelayed, uint256 gasUsedByMerkle) = _verifyMessages(
      hashedLeaves,
      report.innerProofs,
      report.innerProofFlagBits,
      report.outerProofs,
      report.outerProofFlagBits
    );
    uint256 merkleGasShare = gasUsedByMerkle / decodedMessages.length;

    // only allow manual execution if the report is old enough
    if (manualExecution && (block.timestamp - timestampRelayed) < s_config.permissionLessExecutionThresholdSeconds) {
      revert ManualExecutionNotYetEnabled();
    }

    uint256[] memory tokenPerFeeCoin = report.tokenPerFeeCoin;
    for (uint256 i = 0; i < numMsgs; ++i) {
      uint256 gasBegin = gasleft();
      CCIP.Any2EVMSubscriptionMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.Success) revert AlreadyExecuted(message.sequenceNumber);
      SubscriptionInterface.OffRampSubscription memory subscription = s_router.getSubscription(message.receiver);
      if (address(subscription.receiver) == address(0)) {
        revert SubscriptionInterface.SubscriptionNotFound(message.receiver);
      }

      // Any message with a nonce that is n + 1 is allowed.
      // If strict sequencing is disabled then any failed message can be re-executed out-of-order.
      bool isNextInSequence = s_receiverToNonce[message.receiver] + 1 == message.nonce;
      if (!(isNextInSequence || (!subscription.strictSequencing && state == CCIP.MessageExecutionState.Failure))) {
        revert IncorrectNonce(message.nonce);
      }

      _isWellFormed(message);

      for (uint256 j = 0; j < message.tokens.length; j++) {
        _getPool(message.tokens[j]);
      }

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.InProgress;
      CCIP.MessageExecutionState newState = _trialExecute(message);
      s_executedMessages[message.sequenceNumber] = newState;

      // Increment the nonce of the receiver if it's the next nonce in line and it was successfully
      // executed or if the subscription doesn't require strict sequencing.
      if (isNextInSequence && (newState == CCIP.MessageExecutionState.Success || !subscription.strictSequencing)) {
        s_receiverToNonce[message.receiver]++;
      }

      if (!manualExecution) {
        s_router.chargeSubscription(
          message.receiver,
          message.sender,
          ((gasBegin - gasleft() + merkleGasShare) * tx.gasprice * tokenPerFeeCoin[i]) / 1e18
        );
      }
      emit ExecutionCompleted(message.sequenceNumber, newState);
    }
  }

  function _callReceiver(CCIP.Any2EVMSubscriptionMessage memory message) internal {
    if (!message.receiver.isContract()) revert InvalidReceiver(message.receiver);
    CrossChainMessageReceiverInterface msgReceiver = CrossChainMessageReceiverInterface(message.receiver);
    s_router.routeMessage(msgReceiver, message);
  }

  function _trialExecute(CCIP.Any2EVMSubscriptionMessage memory message) internal returns (CCIP.MessageExecutionState) {
    // TODO(Alex) improve external execution flow
    try this.executeSingleMessage(message) {} catch (bytes memory) {
      return CCIP.MessageExecutionState.Failure;
      // TODO execution failure states
      // https://app.shortcut.com/chainlinklabs/story/41622/contract-scaffolding-execution-failure-states
      // revert ExecutionError(message.sequenceNumber, reason);
    }
    return CCIP.MessageExecutionState.Success;
  }

  function _isWellFormed(CCIP.Any2EVMSubscriptionMessage memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length) {
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    }
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
    CCIP.ExecutionReport memory executionReport = abi.decode(report, (CCIP.ExecutionReport));
    this.execute(executionReport, true);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
