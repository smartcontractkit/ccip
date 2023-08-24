// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TokenPool} from "../../pools/TokenPool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";

contract CustomTokenPool is TokenPool {
  event SynthBurned(uint256 amount);
  event SynthMinted(uint256 amount);

  constructor(IERC20 token, address armProxy) TokenPool(token, new address[](0), armProxy) {}

  /// @notice Locks the token in the pool
  /// @param amount Amount to lock
  function lockOrBurn(
    address,
    bytes calldata,
    uint256 amount,
    uint64,
    bytes calldata
  ) external override whenHealthy onlyOnRamp returns (bytes memory) {
    emit SynthBurned(amount);
    return "";
  }

  /// @notice Release tokens from the pool to the recipient
  /// @param amount Amount to release
  function releaseOrMint(
    bytes memory,
    address,
    uint256 amount,
    uint64,
    bytes memory
  ) external override whenHealthy onlyOffRamp {
    emit SynthMinted(amount);
  }
}
