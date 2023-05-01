// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TokenPool} from "./TokenPool.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";

import {FlexibleDecimalERC20} from "./tokens/FlexibleDecimalERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";

/// @notice This pool mints and burns its own tokens, representing a wrapped form of the native token
/// on a source chain - similar to WBTC.
contract WrappedTokenPool is TokenPool, FlexibleDecimalERC20 {
  constructor(
    string memory name,
    string memory symbol,
    uint8 decimals,
    RateLimiter.Config memory rateLimiterConfig
  ) TokenPool(IERC20(address(this)), rateLimiterConfig) FlexibleDecimalERC20(name, symbol, decimals) {}

  /// @notice Burn the token in the pool
  /// @param amount Amount to burn
  function lockOrBurn(
    address,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override whenNotPaused onlyOnRamp {
    _burn(address(this), amount);
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
    _mint(receiver, amount);
    emit Minted(msg.sender, receiver, amount);
  }
}
