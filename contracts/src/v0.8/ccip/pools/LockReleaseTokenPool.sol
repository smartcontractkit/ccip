// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPool} from "./TokenPool.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {SafeERC20} from "../../vendor/SafeERC20.sol";

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

  constructor(IERC20 token) TokenPool(token) {}

  /// @notice Locks the token in the pool
  /// @param amount Amount to lock
  function lockOrBurn(uint256 amount) external override whenNotPaused validateOwnerOrOnRamp {
    emit Locked(msg.sender, amount);
  }

  /// @notice Release tokens from the pool to the recipient
  /// @param recipient Recipient address
  /// @param amount Amount to release
  function releaseOrMint(address recipient, uint256 amount) external override whenNotPaused validateOwnerOrOffRamp {
    getToken().safeTransfer(recipient, amount);
    emit Released(msg.sender, recipient, amount);
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
    i_token.safeTransfer(msg.sender, amount);
    s_liquidityProviderBalances[msg.sender] -= amount;
    emit LiquidityRemoved(msg.sender, amount);
  }
}
