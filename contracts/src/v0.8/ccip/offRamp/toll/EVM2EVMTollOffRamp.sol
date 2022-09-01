// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../interfaces/offRamp/Any2EVMOffRampInterface.sol";
import "../../interfaces/BlobVerifierInterface.sol";
import "../../ocr/OCR2Base.sol";
import "../BaseOffRamp.sol";

/**
 * @notice EVM2EVMTollOffRamp enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract EVM2EVMTollOffRamp is BaseOffRamp, TypeAndVersionInterface, OCR2Base {
  using CCIP for CCIP.EVM2EVMTollMessage;

  string public constant override typeAndVersion = "EVM2EVMTollOffRamp 1.0.0";

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
    CCIP.EVM2EVMTollMessage[] memory decodedMessages = new CCIP.EVM2EVMTollMessage[](numMsgs);

    for (uint256 i = 0; i < numMsgs; ++i) {
      decodedMessages[i] = abi.decode(report.encodedMessages[i], (CCIP.EVM2EVMTollMessage));
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

    for (uint256 i = 0; i < numMsgs; ++i) {
      CCIP.EVM2EVMTollMessage memory message = decodedMessages[i];
      CCIP.MessageExecutionState state = getExecutionState(message.sequenceNumber);
      if (state == CCIP.MessageExecutionState.SUCCESS) revert AlreadyExecuted(message.sequenceNumber);

      _isWellFormed(message);

      for (uint256 j = 0; j < message.tokens.length; ++j) {
        _getPool(message.tokens[j]);
      }

      // If it's the first DON execution attempt, charge the fee.
      if (state == CCIP.MessageExecutionState.UNTOUCHED && !manualExecution) {
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
        address destinationFeeTokenAddress = address(_getPool(message.feeToken).getToken());
        for (uint256 j = 0; j < report.tokenPerFeeCoinAddresses.length; ++j) {
          if (report.tokenPerFeeCoinAddresses[j] == destinationFeeTokenAddress) {
            tokenPerFeeCoin = report.tokenPerFeeCoin[j];
          }
        }
        if (tokenPerFeeCoin == uint256(0)) {
          revert MissingFeeCoinPrice(destinationFeeTokenAddress);
        }
        // Gas cost in wei: gasUsed * gasPrice
        // example: 100k gas, 20 gwei = 1e5 * 20e9  = 2e15
        // Gas cost in token: costInWei * 1e18 / tokenPerFeeCoin
        // example: costInWei 2e15, tokenPerFeeCoin 2e20 = 2e15 * 2e20 / 1e18 = 4e17 tokens
        uint256 feeForGas = ((merkleGasShare + message.gasLimit) * tx.gasprice * tokenPerFeeCoin) / 1 ether;
        if (feeForGas > message.feeTokenAmount) {
          revert InsufficientFeeAmount(message.sequenceNumber, feeForGas, message.feeTokenAmount);
        }

        // _releaseOrMintToken converts the message.feeToken to the proper destination token
        _releaseOrMintToken(message.feeToken, message.feeTokenAmount, address(this));
      }

      s_executedMessages[message.sequenceNumber] = CCIP.MessageExecutionState.IN_PROGRESS;
      CCIP.MessageExecutionState newState = _trialExecute(message._toAny2EVMMessage());
      s_executedMessages[message.sequenceNumber] = newState;

      emit ExecutionStateChanged(message.sequenceNumber, newState);
    }
  }

  function _isWellFormed(CCIP.EVM2EVMTollMessage memory message) private view {
    if (message.sourceChainId != i_sourceChainId) revert InvalidSourceChain(message.sourceChainId);
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
    this.execute(abi.decode(report, (CCIP.ExecutionReport)), false);
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {}

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {}

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {}
}
