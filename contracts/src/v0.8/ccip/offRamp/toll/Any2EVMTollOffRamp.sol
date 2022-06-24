// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../../vendor/SafeERC20.sol";
import "../../interfaces/TollOffRampInterface.sol";
import "../../interfaces/BlobVerifierInterface.sol";
import "../../ocr/OCR2Base.sol";
import "../../health/HealthChecker.sol";
import "../../utils/CCIP.sol";
import "../../pools/TokenPoolRegistry.sol";

/**
 * @notice Any2EVMTollOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract Any2EVMTollOffRamp is
  TollOffRampInterface,
  HealthChecker,
  TokenPoolRegistry,
  TypeAndVersionInterface,
  OCR2Base
{
  using Address for address;
  using SafeERC20 for IERC20;

  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;

  // The router through which all transactions will be executed
  TollOffRampRouterInterface public s_router;
  // The blob verifier contract
  BlobVerifierInterface private s_blobVerifier;

  // The on chain offRamp configuration values
  OffRampConfig private s_config;

  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => CCIP.MessageExecutionState) public executedMessages;

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
  ) OCR2Base(true) HealthChecker(afn, maxTimeWithoutAFNSignal) TokenPoolRegistry(sourceTokens, pools) {
    SOURCE_CHAIN_ID = offRampConfig.sourceChainId;
    CHAIN_ID = chainId;
    s_config = offRampConfig;
    s_blobVerifier = blobVerifier;
  }

  /**
   * @notice setRouter sets a new router
   * @param router the new Router
   * @dev only the owner can call this function
   */
  function setRouter(TollOffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(router);
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
   * @param needFee Whether or not the executor requires a fee
   */
  function execute(CCIP.ExecutionReport memory report, bool needFee)
    external
    override
    whenNotPaused
    whenHealthy
    returns (CCIP.ExecutionResult[] memory)
  {
    if (address(s_router) == address(0)) revert RouterNotSet();
    uint256 numMsgs = report.encodedMessages.length;
    if (numMsgs == 0) revert NoMessagesToExecute();
    bytes32[] memory hashedLeaves = new bytes32[](numMsgs);
    CCIP.Any2EVMTollMessage[] memory decodedMessages = new CCIP.Any2EVMTollMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; i++) {
      decodedMessages[i] = abi.decode(report.encodedMessages[i], (CCIP.Any2EVMTollMessage));
      // TODO: hasher
      // https://app.shortcut.com/chainlinklabs/story/41625/hasher-encoder
      bytes memory data = bytes.concat(hex"00", abi.encode(decodedMessages[i]));
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

    CCIP.ExecutionResult[] memory executionResults = new CCIP.ExecutionResult[](numMsgs);

    for (uint256 i = 0; i < numMsgs; i++) {
      uint256 gasBegin = gasleft();

      CCIP.Any2EVMTollMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = _getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.Success) revert AlreadyExecuted(message.sequenceNumber);

      _isWellFormed(message);

      for (uint256 j = 0; j < message.tokens.length; j++) {
        _getPool(message.tokens[j]);
      }

      if (state != CCIP.MessageExecutionState.Failure && needFee) {
        _releaseOrMintToken(message.feeToken, message.feeTokenAmount, address(this));
      }

      executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.InProgress;
      CCIP.MessageExecutionState newState = _trialExecute(message);
      executedMessages[message.sequenceNumber] = newState;

      uint256 gasUsed = gasBegin - gasleft() + merkleGasShare;
      CCIP.ExecutionResult memory executionResult = CCIP.ExecutionResult({
        sequenceNumber: message.sequenceNumber,
        gasUsed: gasUsed,
        timestampRelayed: timestampRelayed,
        state: newState
      });
      executionResults[i] = executionResult;
      emit ExecutionCompleted(executionResult.sequenceNumber, executionResult.state);
    }

    return executionResults;
  }

  function _releaseOrMintToken(
    IERC20 token,
    uint256 amount,
    address receiver
  ) internal {
    PoolInterface pool = _getPool(token);
    pool.releaseOrMint(receiver, amount);
  }

  function _releaseOrMintTokens(
    IERC20[] memory tokens,
    uint256[] memory amounts,
    address receiver
  ) internal {
    for (uint256 i = 0; i < tokens.length; i++) {
      _releaseOrMintToken(tokens[i], amounts[i], receiver);
    }
  }

  function _callReceiver(CCIP.Any2EVMTollMessage memory message) internal {
    if (!message.receiver.isContract()) revert InvalidReceiver(message.receiver);
    CrossChainMessageReceiverInterface msgReceiver = CrossChainMessageReceiverInterface(message.receiver);
    s_router.routeMessage(msgReceiver, message);
  }

  function _trialExecute(CCIP.Any2EVMTollMessage memory message) internal returns (CCIP.MessageExecutionState) {
    // TODO(Alex) improve external execution flow
    try this.executeSingleMessage(message) {} catch (bytes memory reason) {
      return CCIP.MessageExecutionState.Failure;
      // TODO execution failure states
      // https://app.shortcut.com/chainlinklabs/story/41622/contract-scaffolding-execution-failure-states
      // revert ExecutionError(message.sequenceNumber, reason);
    }
    return CCIP.MessageExecutionState.Success;
  }

  function _verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] memory outerProofs,
    uint256 outerProofFlagBits
  ) internal returns (uint256, uint256) {
    uint256 gasBegin = gasleft();
    uint256 timestamp_relayed = s_blobVerifier.verify(
      hashedLeaves,
      innerProofs,
      innerProofFlagBits,
      outerProofs,
      outerProofFlagBits
    );
    if (timestamp_relayed <= 0) revert RootNotRelayed();
    return (timestamp_relayed, gasBegin - gasleft());
  }

  function _getExecutionState(uint64 sequenceNumber) internal view returns (CCIP.MessageExecutionState) {
    return executedMessages[sequenceNumber];
  }

  function _isWellFormed(CCIP.Any2EVMTollMessage memory message) private view {
    if (message.sourceChainId != SOURCE_CHAIN_ID) revert InvalidSourceChain(message.sourceChainId);
    if (message.tokens.length > uint256(s_config.maxTokensLength) || message.tokens.length != message.amounts.length) {
      revert UnsupportedNumberOfTokens(message.sequenceNumber);
    }
    if (message.data.length > uint256(s_config.maxDataSize))
      revert MessageTooLarge(uint256(s_config.maxDataSize), message.data.length);
  }

  function _getPool(IERC20 token) private view returns (PoolInterface pool) {
    pool = getPool(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
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
    CCIP.ExecutionResult[] memory executionResult = this.execute(executionReport, true);

    for (uint256 i = 0; i < executionReport.encodedMessages.length; i++) {
      CCIP.Any2EVMTollMessage memory message = abi.decode(
        executionReport.encodedMessages[i],
        (CCIP.Any2EVMTollMessage)
      );
      PoolInterface pool = _getPool(message.feeToken);
      uint256 tokenPerFeeCoin;
      for (uint256 j = 0; j < executionReport.tokenPerFeeCoinAddresses.length; j++) {
        if (executionReport.tokenPerFeeCoinAddresses[j] == address(message.feeToken)) {
          tokenPerFeeCoin = executionReport.tokenPerFeeCoin[j];
        }
      }
      if (tokenPerFeeCoin == uint256(0)) {
        revert MissingFeeCoinPrice(address(message.feeToken));
      }
      // Example with token being link. 1 LINK = 1e18 Juels.
      // tx.gasprice is wei / gas
      // gas * wei/gas * (juels / wei) (problem is that juels per wei could be < 1, say since 1 link < 1 eth)
      // instead we use juels per unit ETH, which > 1, assuming 1 juel < 1 ETH (safe).
      // gas * wei/gas * (juels / (ETH * 1e18 WEI/ETH))
      // gas * wei/gas * juels/ETH / (1e18 wei/ETH)
      // Example 1e6 gas * (200e9 wei / gas) * (6253149865160030 juels / ETH) / (1e18 wei/ETH) = 1.25e15 juels
      uint256 feeForGas = (executionResult[i].gasUsed * tx.gasprice * tokenPerFeeCoin) / 1e18;
      uint256 refund = message.feeTokenAmount - feeForGas;
      _releaseOrMintToken(message.feeToken, refund, message.receiver);
    }
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "Any2EVMTollOffRamp 1.0.0";
  }
}
