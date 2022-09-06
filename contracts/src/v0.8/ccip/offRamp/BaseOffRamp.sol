// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/Address.sol";
import "../health/HealthChecker.sol";
import "../pools/TokenPoolRegistry.sol";
import "../interfaces/offRamp/Any2EVMOffRampInterface.sol";
import "../rateLimiter/AggregateRateLimiter.sol";

/**
 * @notice A base OffRamp contract that every OffRamp should expand on
 */
contract BaseOffRamp is BaseOffRampInterface, HealthChecker, TokenPoolRegistry, AggregateRateLimiter {
  using Address for address;

  // Chain ID of the source chain
  uint256 public immutable i_sourceChainId;
  // Chain ID of this chain
  uint256 public immutable i_chainId;

  // The router through which all transactions will be executed
  Any2EVMOffRampRouterInterface public s_router;

  // The blob verifier contract
  BlobVerifierInterface internal s_blobVerifier;

  // The on chain offRamp configuration values
  OffRampConfig internal s_config;

  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => CCIP.MessageExecutionState) internal s_executedMessages;

  constructor(
    uint256 sourceChainId,
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    // OnrampAddress, needed for hashing in the future so already added to the interface
    address,
    AFNInterface afn,
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    RateLimiterConfig memory rateLimiterConfig,
    address tokenLimitsAdmin
  )
    HealthChecker(afn)
    TokenPoolRegistry(sourceTokens, pools)
    AggregateRateLimiter(rateLimiterConfig, tokenLimitsAdmin)
  {
    // TokenPoolRegistry does a check on tokens.length != pools.length
    i_sourceChainId = sourceChainId;
    i_chainId = chainId;
    s_config = offRampConfig;
    s_blobVerifier = blobVerifier;
  }

  /**
   * @notice Uses the pool to release or mint tokens and send them to
   *          the given `receiver` address.
   */
  function _releaseOrMintToken(
    PoolInterface pool,
    uint256 amount,
    address receiver
  ) internal {
    pool.releaseOrMint(receiver, amount);
  }

  /**
   * @notice Uses pools to release or mint a number of different tokens
   *           and send them to the given `receiver` address.
   */
  function _releaseOrMintTokens(
    PoolInterface[] memory pools,
    uint256[] memory amounts,
    address receiver
  ) internal {
    if (pools.length != amounts.length) revert TokenAndAmountMisMatch();
    for (uint256 i = 0; i < pools.length; ++i) {
      _releaseOrMintToken(pools[i], amounts[i], receiver);
    }
  }

  /**
   * @notice Verifies that the given hashed messages are valid leaves of
   *          a relayed merkle tree.
   */
  function _verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory innerProofs,
    uint256 innerProofFlagBits,
    bytes32[] memory outerProofs,
    uint256 outerProofFlagBits
  ) internal returns (uint256, uint256) {
    uint256 gasBegin = gasleft();
    uint256 timestampRelayed = s_blobVerifier.verify(
      hashedLeaves,
      innerProofs,
      innerProofFlagBits,
      outerProofs,
      outerProofFlagBits
    );
    if (timestampRelayed <= 0) revert RootNotRelayed();
    return (timestampRelayed, gasBegin - gasleft());
  }

  /**
   * @notice Try executing a message
   * @param message CCIP.Any2EVMMessageFromSender memory message
   * @return CCIP.ExecutionState
   */
  function _trialExecute(CCIP.Any2EVMMessageFromSender memory message) internal returns (CCIP.MessageExecutionState) {
    try this.executeSingleMessage(message) {} catch (bytes memory err) {
      if (BaseOffRampInterface.ReceiverError.selector == bytes4(err)) {
        return CCIP.MessageExecutionState.FAILURE;
      } else {
        revert ExecutionError();
      }
    }
    return CCIP.MessageExecutionState.SUCCESS;
  }

  /**
   * @notice Execute a single message
   * @param message The Any2EVMMessageFromSender message that will be executed
   * @dev this can only be called by the contract itself. It is part of
   * the Execute call, as we can only try/catch on external calls.
   */
  function executeSingleMessage(CCIP.Any2EVMMessageFromSender memory message) external {
    if (msg.sender != address(this)) revert CanOnlySelfCall();
    if (message.destTokens.length > 0) {
      _removeTokens(message.destTokens, message.amounts);
      _releaseOrMintTokens(message.destPools, message.amounts, message.receiver);
    }

    _callReceiver(message);
  }

  function _callReceiver(CCIP.Any2EVMMessageFromSender memory message) internal {
    if (!message.receiver.isContract()) return;
    if (!s_router.routeMessage(message)) revert ReceiverError();
  }

  /**
   * @notice Reverts as this contract should not access CCIP messages
   */
  function ccipReceive(CCIP.Any2EVMMessageFromSender calldata) external pure {
    revert();
  }

  /// @inheritdoc BaseOffRampInterface
  function execute(CCIP.ExecutionReport memory, bool) external virtual override {
    revert();
  }

  /// @inheritdoc BaseOffRampInterface
  function setRouter(Any2EVMOffRampRouterInterface router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(address(router));
  }

  /// @inheritdoc BaseOffRampInterface
  function getRouter() external view override returns (Any2EVMOffRampRouterInterface) {
    return s_router;
  }

  /// @inheritdoc BaseOffRampInterface
  function getExecutionState(uint64 sequenceNumber) public view returns (CCIP.MessageExecutionState) {
    return s_executedMessages[sequenceNumber];
  }

  /// @inheritdoc BaseOffRampInterface
  function getBlobVerifier() public view returns (BlobVerifierInterface) {
    return s_blobVerifier;
  }

  /// @inheritdoc BaseOffRampInterface
  function setBlobVerifier(BlobVerifierInterface blobVerifier) public onlyOwner {
    s_blobVerifier = blobVerifier;
  }

  /// @inheritdoc BaseOffRampInterface
  function getConfig() public view returns (OffRampConfig memory) {
    return s_config;
  }

  /// @inheritdoc BaseOffRampInterface
  function setConfig(OffRampConfig memory config) public onlyOwner {
    s_config = config;

    emit OffRampConfigSet(config);
  }

  /**
   * @notice Returns the pool for a given source chain token.
   */
  function _getPool(IERC20 token) internal view returns (PoolInterface pool) {
    pool = getPool(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
  }
}
