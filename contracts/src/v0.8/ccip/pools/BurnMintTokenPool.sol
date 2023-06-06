// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../shared/token/ERC20/IBurnMintERC20.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "./TokenPool.sol";

/// @notice This pool mints and burns a 3rd-party token.
/// @dev Pool whitelisting mode is set in the constructor and cannot be modified later.
/// It either accepts any address as originalSender, or only accepts whitelisted originalSender.
/// The only way to change whitelisting mode is to deploy a new pool.
/// If that is expected, please make sure the token's burner/minter roles are adjustable.
contract BurnMintTokenPool is TokenPool {
  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    RateLimiter.Config memory rateLimiterConfig
  ) TokenPool(token, allowlist, rateLimiterConfig) {}

  /// @notice Burn the token in the pool
  /// @dev Burn is not rate limited at per-pool level. Burn does not contribute to honey pot risk.
  /// Benefits of rate limiting here does not justify the extra gas cost.
  /// @param amount Amount to burn
  function lockOrBurn(
    address originalSender,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override whenNotPaused onlyOnRamp checkAllowList(originalSender) {
    IBurnMintERC20(address(i_token)).burn(amount);
    emit Burned(msg.sender, amount);
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @param receiver Recipient address
  /// @param amount Amount to mint
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64,
    bytes memory
  ) external override whenNotPaused onlyOffRamp {
    _consumeRateLimit(amount);
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
    emit Minted(msg.sender, receiver, amount);
  }
}
