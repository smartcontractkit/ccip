// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../ocr/OCR2Base.sol";
import "../interfaces/Any2EVMTollOffRampInterface.sol";
import "../../applications/interfaces/CrossChainMessageReceiverInterface.sol";
import "../BaseOffRamp.sol";

/**
 * @notice Any2EVMTollOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract Any2EVMTollOffRamp is Any2EVMTollOffRampInterface, BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "Any2EVMTollOffRamp 1.0.0";

  // The router through which all transactions will be executed
  Any2EVMTollOffRampRouterInterface private s_router;

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
  function setRouter(Any2EVMTollOffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(address(router));
  }

  /**
   * @notice ccipReceive implements the receive function to create a
   * collision if some other method happens to hash to the same signature
   */
  function ccipReceive(CCIP.Any2EVMTollMessage calldata) external pure override {
    revert();
  }

  /**
   * @notice Execute a single message
   * @param message The Any2EVMTollMessage message that will be executed
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMTollMessage memory message) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    // TODO: token limiter logic
    // https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
    _releaseOrMintTokens(message.tokens, message.amounts, message.receiver);
    _callReceiver(message);
  }

  /**
   * @notice Execute a series of one or more messages using a merkle proof
   * @param report ExecutionReport
   * @param manualExecution Whether the DON auto executes or it is manually initiated
   */
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
    CCIP.Any2EVMTollMessage[] memory decodedMessages = new CCIP.Any2EVMTollMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      decodedMessages[i] = abi.decode(report.encodedMessages[i], (CCIP.Any2EVMTollMessage));
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

    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.Any2EVMTollMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.Success) revert AlreadyExecuted(message.sequenceNumber);

      _isWellFormed(message);

      for (uint256 j = 0; j < message.tokens.length; j++) {
        _getPool(message.tokens[j]);
      }

      // If it's the first DON execution attempt, charge the fee.
      if (state == CCIP.MessageExecutionState.Untouched && !manualExecution) {
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
        for (uint256 j = 0; j < report.tokenPerFeeCoinAddresses.length; ++j) {
          if (report.tokenPerFeeCoinAddresses[j] == address(message.feeToken)) {
            tokenPerFeeCoin = report.tokenPerFeeCoin[j];
          }
        }
        if (tokenPerFeeCoin == uint256(0)) {
          revert MissingFeeCoinPrice(address(message.feeToken));
        }
        uint256 feeForGas = ((merkleGasShare + message.gasLimit) * tx.gasprice * tokenPerFeeCoin) / 1e18;
        if (feeForGas > message.feeTokenAmount) {
          revert InsufficientFeeAmount(message.sequenceNumber);
        }
        _releaseOrMintToken(message.feeToken, message.feeTokenAmount, address(this));
      }

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.InProgress;
      CCIP.MessageExecutionState newState = _trialExecute(message);
      s_executedMessages[message.sequenceNumber] = newState;

      emit ExecutionCompleted(message.sequenceNumber, newState);
    }
  }

  function _callReceiver(CCIP.Any2EVMTollMessage memory message) internal {
    if (!message.receiver.isContract()) revert InvalidReceiver(message.receiver);
    CrossChainMessageReceiverInterface msgReceiver = CrossChainMessageReceiverInterface(message.receiver);
    s_router.routeMessage(msgReceiver, message);
  }

  function _trialExecute(CCIP.Any2EVMTollMessage memory message) internal returns (CCIP.MessageExecutionState) {
    // TODO(Alex) improve external execution flow
    try this.executeSingleMessage(message) {} catch (bytes memory) {
      return CCIP.MessageExecutionState.Failure;
      // TODO execution failure states
      // https://app.shortcut.com/chainlinklabs/story/41622/contract-scaffolding-execution-failure-states
      // revert ExecutionError(message.sequenceNumber, reason);
    }
    return CCIP.MessageExecutionState.Success;
  }

  function _isWellFormed(CCIP.Any2EVMTollMessage memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length) {
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    }
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  function getRouter() external view returns (Any2EVMTollOffRampRouterInterface) {
    return s_router;
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
