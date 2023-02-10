// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {IBaseOffRamp} from "../interfaces/offRamp/IBaseOffRamp.sol";
import {IAny2EVMOffRampRouter} from "../interfaces/offRamp/IAny2EVMOffRampRouter.sol";
import {ICommitStore} from "../interfaces/ICommitStore.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IAFN} from "../interfaces/health/IAFN.sol";

import {Internal} from "../models/Internal.sol";
import {Common} from "../models/Common.sol";
import {AggregateRateLimiter} from "../rateLimiter/AggregateRateLimiter.sol";
import {HealthChecker} from "../health/HealthChecker.sol";
import {OffRampTokenPoolRegistry} from "../pools/OffRampTokenPoolRegistry.sol";

import {Address} from "../../vendor/Address.sol";
import {IERC20} from "../../vendor/IERC20.sol";

/**
 * @notice An Any2EVM base OffRamp contract that every OffRamp should expand on
 */
contract Any2EVMBaseOffRamp is IBaseOffRamp, HealthChecker, OffRampTokenPoolRegistry, AggregateRateLimiter {
  using Address for address;

  // Chain ID of the source chain
  uint64 internal immutable i_sourceChainId;
  // Chain ID of this chain
  uint64 internal immutable i_chainId;
  // OnRamp address on the source chain
  address internal immutable i_onRampAddress;

  // The router through which all transactions will be executed
  IAny2EVMOffRampRouter internal s_router;

  // The commitStore contract
  ICommitStore internal s_commitStore;

  // A mapping of sequence numbers to execution state.
  // This makes sure we never execute a message twice.
  mapping(uint64 => Internal.MessageExecutionState) internal s_executedMessages;

  constructor(
    uint64 sourceChainId,
    uint64 chainId,
    address onRampAddress,
    ICommitStore commitStore,
    IAFN afn,
    IERC20[] memory sourceTokens,
    IPool[] memory pools,
    RateLimiterConfig memory rateLimiterConfig
  ) HealthChecker(afn) OffRampTokenPoolRegistry(sourceTokens, pools) AggregateRateLimiter(rateLimiterConfig) {
    if (onRampAddress == address(0)) revert ZeroAddressNotAllowed();
    // OffRampTokenPoolRegistry does a check on tokensAndAmounts.length != pools.length
    i_sourceChainId = sourceChainId;
    i_chainId = chainId;
    i_onRampAddress = onRampAddress;
    s_commitStore = commitStore;
  }

  /**
   * @notice Uses the pool to release or mint tokens and send them to
   *          the given `receiver` address.
   */
  function _releaseOrMintToken(
    IPool pool,
    uint256 amount,
    address receiver
  ) internal {
    pool.releaseOrMint(receiver, amount);
  }

  /**
   * @notice Uses pools to release or mint a number of different tokens
   *           and send them to the given `receiver` address.
   */
  function _releaseOrMintTokens(Common.EVMTokenAndAmount[] memory sourceTokensAndAmounts, address receiver)
    internal
    returns (Common.EVMTokenAndAmount[] memory)
  {
    Common.EVMTokenAndAmount[] memory destTokensAndAmounts = new Common.EVMTokenAndAmount[](
      sourceTokensAndAmounts.length
    );
    for (uint256 i = 0; i < sourceTokensAndAmounts.length; ++i) {
      IPool pool = _getPool(IERC20(sourceTokensAndAmounts[i].token));
      _releaseOrMintToken(pool, sourceTokensAndAmounts[i].amount, receiver);
      destTokensAndAmounts[i].token = address(pool.getToken());
      destTokensAndAmounts[i].amount = sourceTokensAndAmounts[i].amount;
    }
    _removeTokens(destTokensAndAmounts);
    return destTokensAndAmounts;
  }

  /**
   * @notice Verifies that the given hashed messages are valid leaves of
   *          a committed merkle tree.
   */
  function _verifyMessages(
    bytes32[] memory hashedLeaves,
    bytes32[] memory proofs,
    uint256 proofFlagBits
  ) internal returns (uint256, uint256) {
    uint256 gasBegin = gasleft();
    uint256 timestampCommitted = s_commitStore.verify(hashedLeaves, proofs, proofFlagBits);
    if (timestampCommitted <= 0) revert RootNotCommitted();
    return (timestampCommitted, gasBegin - gasleft());
  }

  /**
   * @notice Reverts as this contract should not access CCIP messages
   */
  function ccipReceive(Common.Any2EVMMessage calldata) external pure {
    // solhint-disable-next-line reason-string
    revert();
  }

  /// @inheritdoc IBaseOffRamp
  function setRouter(IAny2EVMOffRampRouter router) external onlyOwner {
    s_router = router;
    emit OffRampRouterSet(address(router), i_sourceChainId, i_onRampAddress);
  }

  /// @inheritdoc IBaseOffRamp
  function getRouter() external view override returns (IAny2EVMOffRampRouter) {
    return s_router;
  }

  /// @inheritdoc IBaseOffRamp
  function getExecutionState(uint64 sequenceNumber) public view returns (Internal.MessageExecutionState) {
    return s_executedMessages[sequenceNumber];
  }

  /// @inheritdoc IBaseOffRamp
  function getCommitStore() external view returns (ICommitStore) {
    return s_commitStore;
  }

  /// @inheritdoc IBaseOffRamp
  function setCommitStore(ICommitStore commitStore) external onlyOwner {
    s_commitStore = commitStore;
  }

  function getChainIDs() external view returns (uint64 sourceChainId, uint64 chainId) {
    sourceChainId = i_sourceChainId;
    chainId = i_chainId;
  }

  /**
   * @notice Returns the pool for a given source chain token.
   */
  function _getPool(IERC20 token) internal view returns (IPool pool) {
    pool = getPoolBySourceToken(token);
    if (address(pool) == address(0)) revert UnsupportedToken(token);
  }

  function _metadataHash(bytes32 prefix) internal view returns (bytes32) {
    return keccak256(abi.encode(prefix, i_sourceChainId, i_chainId, i_onRampAddress));
  }
}
