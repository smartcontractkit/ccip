// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {TokenPool} from "./TokenPool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.0/token/ERC20/utils/SafeERC20.sol";

/// @notice Token pool used for tokens on their native chain. This uses a lock and release mechanism.
/// Because of lock/unlock requiring liquidity, this pool contract also has function to add and remove
/// liquidity. This allows for proper bookkeeping for both user and liquidity provider balances.
/// @dev One token per LockReleaseTokenPool.
contract LockReleaseTokenPool is TokenPool {
  using SafeERC20 for IERC20;

  event LiquidityAdded(address indexed provider, uint256 indexed amount);
  event LiquidityRemoved(address indexed provider, uint256 indexed amount);

  error InsufficientLiquidity();
  error WithdrawalTooHigh();

  mapping(address => uint256) internal s_liquidityProviderBalances;

  constructor(
    IERC20 token,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig
  ) TokenPool(token, allowlist, rateLimiterConfig) {}

  /// @notice Locks the token in the pool
  /// @dev Locks are not rate limited at per-pool level. Each pool is shared across lanes,
  /// rate limiting locks does not meaningfully mitigate honeypot risk.
  /// Benefits of rate limiting here does not justify the extra gas cost.
  /// @param amount Amount to lock
  function lockOrBurn(
    address originalSender,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override whenNotPaused onlyOnRamp checkAllowList(originalSender) returns (bytes memory) {
    emit Locked(msg.sender, amount);
    return "";
  }

  /// @notice Release tokens from the pool to the recipient
  /// @param receiver Recipient address
  /// @param amount Amount to release
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64,
    bytes memory
  ) external override whenNotPaused onlyOffRamp {
    _consumeRateLimit(amount);
    getToken().safeTransfer(receiver, amount);
    emit Released(msg.sender, receiver, amount);
  }

  /// @notice Gets the amount of provided liquidity for a given address.
  /// @param provider The address for which to get the balance.
  /// @return The current provided liquidity.
  function getProvidedLiquidity(address provider) external view returns (uint256) {
    return s_liquidityProviderBalances[provider];
  }

  /// @notice Adds liquidity to the pool. The tokens should be approved first.
  /// @param amount The amount of liquidity to provide.
  function addLiquidity(uint256 amount) external {
    i_token.safeTransferFrom(msg.sender, address(this), amount);
    s_liquidityProviderBalances[msg.sender] += amount;
    emit LiquidityAdded(msg.sender, amount);
  }

  /// @notice Removed liquidity to the pool. The tokens will be sent to msg.sender.
  /// @param amount The amount of liquidity to remove.
  function removeLiquidity(uint256 amount) external {
    if (s_liquidityProviderBalances[msg.sender] < amount) revert WithdrawalTooHigh();
    if (i_token.balanceOf(address(this)) < amount) revert InsufficientLiquidity();
    s_liquidityProviderBalances[msg.sender] -= amount;
    i_token.safeTransfer(msg.sender, amount);
    emit LiquidityRemoved(msg.sender, amount);
  }
}
