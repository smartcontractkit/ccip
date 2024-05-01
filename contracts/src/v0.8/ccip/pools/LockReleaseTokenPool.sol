// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {ILiquidityContainer} from "../../liquiditymanager/interfaces/ILiquidityContainer.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";

import {Pool} from "../libraries/Pool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
/// Because of lock/unlock requiring liquidity, this pool contract also has function to add and remove
/// liquidity. This allows for proper bookkeeping for both user and liquidity provider balances.
/// @dev One token per LockReleaseTokenPool.
contract LockReleaseTokenPool is TokenPool, ILiquidityContainer, ITypeAndVersion {
  using SafeERC20 for IERC20;

  error InsufficientLiquidity();
  error LiquidityNotAccepted();
  error Unauthorized(address caller);

  string public constant override typeAndVersion = "LockReleaseTokenPool 1.5.0-dev";

  /// @dev The unique lock release pool flag to signal through EIP 165.
  bytes4 private constant LOCK_RELEASE_INTERFACE_ID = bytes4(keccak256("LockReleaseTokenPool"));

  /// @dev Whether or not the pool accepts liquidity.
  /// External liquidity is not required when there is one canonical token deployed to a chain,
  /// and CCIP is facilitating mint/burn on all the other chains, in which case the invariant
  /// balanceOf(pool) on home chain == sum(totalSupply(mint/burn "wrapped" token) on all remote chains) should always hold
  bool internal immutable i_acceptLiquidity;
  /// @notice The address of the rebalancer.
  address internal s_rebalancer;
  /// @notice The address of the rate limiter admin.
  /// @dev Can be address(0) if none is configured.
  address internal s_rateLimitAdmin;

  constructor(
    IERC20 token,
    address[] memory allowlist,
    address armProxy,
    bool acceptLiquidity,
    address router
  ) TokenPool(token, allowlist, armProxy, router) {
    i_acceptLiquidity = acceptLiquidity;
  }

  /// @notice Locks the token in the pool
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function lockOrBurn(Pool.LockOrBurnInV1 calldata lockOrBurnIn)
    external
    virtual
    override
    whenHealthy
    returns (Pool.LockOrBurnOutV1 memory)
  {
    _checkAllowList(lockOrBurnIn.originalSender);
    _onlyOnRamp(lockOrBurnIn.remoteChainSelector);
    _consumeOutboundRateLimit(lockOrBurnIn.remoteChainSelector, lockOrBurnIn.amount);

    emit Locked(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({destPoolAddress: getRemotePool(lockOrBurnIn.remoteChainSelector), destPoolData: ""});
  }

  /// @notice Release tokens from the pool to the recipient
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function releaseOrMint(Pool.ReleaseOrMintInV1 calldata releaseOrMintIn)
    external
    virtual
    override
    whenHealthy
    returns (Pool.ReleaseOrMintOutV1 memory)
  {
    _onlyOffRamp(releaseOrMintIn.remoteChainSelector);
    _validateSourceCaller(releaseOrMintIn.remoteChainSelector, releaseOrMintIn.sourcePoolAddress);
    _consumeInboundRateLimit(releaseOrMintIn.remoteChainSelector, releaseOrMintIn.amount);

    getToken().safeTransfer(releaseOrMintIn.receiver, releaseOrMintIn.amount);

    emit Released(msg.sender, releaseOrMintIn.receiver, releaseOrMintIn.amount);

    return Pool.ReleaseOrMintOutV1({localToken: address(i_token), destinationAmount: releaseOrMintIn.amount});
  }

  /// @notice returns the lock release interface flag used for EIP165 identification.
  function getLockReleaseInterfaceId() public pure returns (bytes4) {
    return LOCK_RELEASE_INTERFACE_ID;
  }

  // @inheritdoc IERC165
  function supportsInterface(bytes4 interfaceId) public pure virtual override returns (bool) {
    return interfaceId == LOCK_RELEASE_INTERFACE_ID || interfaceId == type(ILiquidityContainer).interfaceId
      || super.supportsInterface(interfaceId);
  }

  /// @notice Gets LiquidityManager, can be address(0) if none is configured.
  /// @return The current liquidity manager.
  function getRebalancer() external view returns (address) {
    return s_rebalancer;
  }

  /// @notice Sets the LiquidityManager address.
  /// @dev Only callable by the owner.
  function setRebalancer(address rebalancer) external onlyOwner {
    s_rebalancer = rebalancer;
  }

  /// @notice Sets the rate limiter admin address.
  /// @dev Only callable by the owner.
  /// @param rateLimitAdmin The new rate limiter admin address.
  function setRateLimitAdmin(address rateLimitAdmin) external onlyOwner {
    s_rateLimitAdmin = rateLimitAdmin;
  }

  /// @notice Gets the rate limiter admin address.
  function getRateLimitAdmin() external view returns (address) {
    return s_rateLimitAdmin;
  }

  /// @notice Checks if the pool can accept liquidity.
  /// @return true if the pool can accept liquidity, false otherwise.
  function canAcceptLiquidity() external view returns (bool) {
    return i_acceptLiquidity;
  }

  /// @notice Adds liquidity to the pool. The tokens should be approved first.
  /// @param amount The amount of liquidity to provide.
  function provideLiquidity(uint256 amount) external {
    if (!i_acceptLiquidity) revert LiquidityNotAccepted();
    if (s_rebalancer != msg.sender) revert Unauthorized(msg.sender);

    i_token.safeTransferFrom(msg.sender, address(this), amount);
    emit LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param amount The amount of liquidity to remove.
  function withdrawLiquidity(uint256 amount) external {
    if (s_rebalancer != msg.sender) revert Unauthorized(msg.sender);

    if (i_token.balanceOf(address(this)) < amount) revert InsufficientLiquidity();
    i_token.safeTransfer(msg.sender, amount);
    emit LiquidityRemoved(msg.sender, amount);
  }

  /// @notice Sets the rate limiter admin address.
  /// @dev Only callable by the owner or the rate limiter admin. NOTE: overwrites the normal
  /// onlyAdmin check in the base implementation to also allow the rate limiter admin.
  /// @param remoteChainSelector The remote chain selector for which the rate limits apply.
  /// @param outboundConfig The new outbound rate limiter config.
  /// @param inboundConfig The new inbound rate limiter config.
  function setChainRateLimiterConfig(
    uint64 remoteChainSelector,
    RateLimiter.Config memory outboundConfig,
    RateLimiter.Config memory inboundConfig
  ) external override {
    if (msg.sender != s_rateLimitAdmin && msg.sender != owner()) revert Unauthorized(msg.sender);

    _setRateLimitConfig(remoteChainSelector, outboundConfig, inboundConfig);
  }
}
