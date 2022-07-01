// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "../blobVerifier/interfaces/BlobVerifierInterface.sol";
import "../health/HealthChecker.sol";
import "../utils/CCIP.sol";
import "../pools/TokenPoolRegistry.sol";
import "./interfaces/BaseOffRampInterface.sol";

/**
 * @notice A base OffRamp contract that every OffRamp should expand on
 */
contract BaseOffRamp is BaseOffRampInterface, HealthChecker, TokenPoolRegistry {
  using Address for address;
  using SafeERC20 for IERC20;

  // Chain ID of the source chain
  uint256 public immutable SOURCE_CHAIN_ID;
  // Chain ID of this chain
  uint256 public immutable CHAIN_ID;

  // The blob verifier contract
  BlobVerifierInterface internal s_blobVerifier;

  // The on chain offRamp configuration values
  OffRampConfig internal s_config;

  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => CCIP.MessageExecutionState) internal s_executedMessages;

  constructor(
    uint256 chainId,
    OffRampConfig memory offRampConfig,
    BlobVerifierInterface blobVerifier,
    // OnrampAddress, needed for hashing in the future so already added to the interface
    address,
    AFNInterface afn,
    // TODO token limiter contract
    // https://app.shortcut.com/chainlinklabs/story/41867/contract-scaffolding-aggregatetokenlimiter-contract
    IERC20[] memory sourceTokens,
    PoolInterface[] memory pools,
    uint256 maxTimeWithoutAFNSignal
  ) HealthChecker(afn, maxTimeWithoutAFNSignal) TokenPoolRegistry(sourceTokens, pools) {
    // TokenPoolRegistry does a check on tokens.length != pools.length
    SOURCE_CHAIN_ID = offRampConfig.sourceChainId;
    CHAIN_ID = chainId;
    s_config = offRampConfig;
    s_blobVerifier = blobVerifier;
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
    if (SOURCE_CHAIN_ID != config.sourceChainId) revert InvalidSourceChain(config.sourceChainId);
    s_config = config;

    emit OffRampConfigSet(config);
  }

  /**
   * @notice Uses the pool to release or mint tokens and send them to
   *          the given `receiver` address.
   */
  function _releaseOrMintToken(
    IERC20 token,
    uint256 amount,
    address receiver
  ) internal {
    PoolInterface pool = _getPool(token);
    pool.releaseOrMint(receiver, amount);
  }

  /**
   * @notice Uses pools to release or mint a number of different tokens
   *           and send them to the given `receiver` address.
   */
  function _releaseOrMintTokens(
    IERC20[] memory tokens,
    uint256[] memory amounts,
    address receiver
  ) internal {
    if (tokens.length != amounts.length) revert TokenAndAmountMisMatch();
    for (uint256 i = 0; i < tokens.length; ++i) {
      _releaseOrMintToken(tokens[i], amounts[i], receiver);
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

  /**
   * @notice Returns the current execution state of a message based on its
   *          sequenceNumber.
   */
  function _getPool(IERC20 token) internal view returns (PoolInterface pool) {
    pool = getPool(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
  }
}
