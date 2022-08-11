// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../subscription/Subscription.sol";
import "./Any2EVMSubscriptionOffRampRouter.sol";
import "../../ocr/OCR2Base.sol";
import "../BaseOffRamp.sol";

/**
 * @notice EVM2EVMSubscriptionOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMSubscriptionOffRamp is BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using Address for address;
  using SafeERC20 for IERC20;
  using CCIP for CCIP.EVM2EVMSubscriptionMessage;

  string public constant override typeAndVersion = "EVM2EVMSubscriptionOffRamp 1.0.0";

  mapping(address => uint64) public s_receiverToNonce;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    // OnrampAddress, needed for hashing in the future so already added to the interface
    address onRampAddress,
    AFNInterface afn,
    // TODO token limiter contract
    // https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools
  )
    OCR2Base(true)
    BaseOffRamp(sourceChainId, chainId, offRampConfig, blobVerifier, onRampAddress, afn, sourceTokens, pools)
  {}

  function execute(CCIP.ExecutionReport memory report, bool manualExecution)
    external
    override
    whenNotPaused
    whenHealthy
  {
    address routerAddress = address(s_router);
    if (routerAddress == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    CCIP.EVM2EVMSubscriptionMessage[] memory decodedMessages = new CCIP.EVM2EVMSubscriptionMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      decodedMessages[i] = abi.decode(report.encodedMessages[i], (CCIP.EVM2EVMSubscriptionMessage));
      // TODO: hasher
      // https://app.shortcut.com/chainlinklabs/story/41625/hasher-encoder
      // check abi.encodePacked usage for hash preimages, compare gas
      hashedLeaves[i] = keccak256(bytes.concat(hex"00", report.encodedMessages[i]));
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
      CCIP.EVM2EVMSubscriptionMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);
      SubscriptionInterface.OffRampSubscription memory subscription = Subscription(routerAddress).getSubscription(
        message.receiver
      );
      if (address(subscription.receiver) == address(0)) {
        revert SubscriptionInterface.SubscriptionNotFound(message.receiver);
      }

      // Any message with a nonce that is n + 1 is allowed.
      // If strict sequencing is disabled then any failed message can be re-executed out-of-order.
      bool isNextInSequence = s_receiverToNonce[message.receiver] + 1 == message.nonce;
      if (!(isNextInSequence || (!subscription.strictSequencing && state == CCIP.MessageExecutionState.FAILURE))) {
        revert IncorrectNonce(message.nonce);
      }

      _isWellFormed(message);

      for (uint256 j = 0; j < message.tokens.length; ++j) {
        _getPool(message.tokens[j]);
      }

      // Reduce stack pressure
      {
        s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
        CCIP.MessageExecutionState newState = _trialExecute(message._toAny2EVMMessage());
        s_executedMessages[message.sequenceNumber] = newState;
        emit ExecutionStateChanged(message.sequenceNumber, newState);

        // Increment the nonce of the receiver if it's the next nonce in line and it was successfully
        // executed or if the subscription doesn't require strict sequencing.
        if (isNextInSequence && (newState == CCIP.MessageExecutionState.SUCCESS || !subscription.strictSequencing)) {
          s_receiverToNonce[message.receiver]++;
        }
      }

      if (!manualExecution) {
        Any2EVMSubscriptionOffRampRouter(routerAddress).chargeSubscription(
          message.receiver,
          message.sender,
          // Gas cost in wei: gasUsed * gasPrice
          // example: 100k gas, 20 gwei = 1e5 * 20e9  = 2e15
          // Gas cost in token: costInWei * 1e18 / tokenPerFeeCoin
          // example: costInWei 2e15, tokenPerFeeCoin 2e20 = 2e15 * 2e20 / 1e18 = 4e17 tokens
          ((gasBegin - gasleft() + merkleGasShare) * tx.gasprice * tokenPerFeeCoin[i]) / 1 ether
        );
      }
    }
  }

  function _isWellFormed(CCIP.EVM2EVMSubscriptionMessage memory message) private view {
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
