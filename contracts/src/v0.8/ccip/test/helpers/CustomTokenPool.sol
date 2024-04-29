// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Pool} from "../../libraries/Pool.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract CustomTokenPool is TokenPool {
  event SynthBurned(uint256 amount);
  event SynthMinted(uint256 amount);

  constructor(IERC20 token, address armProxy, address router) TokenPool(token, new address[](0), armProxy, router) {}

  /// @notice Locks the token in the pool
  function lockOrBurn(bytes calldata lockOrBurnIn) external virtual override whenHealthy returns (bytes memory) {
    Pool.LockOrBurnInV1 memory lockOrBurnData = Pool._decodeLockOrBurnInV1(lockOrBurnIn);
    _onlyOnRamp(lockOrBurnData.remoteChainSelector);
    emit SynthBurned(lockOrBurnData.amount);
    return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), "");
  }

  /// @notice Release tokens from the pool to the recipient
  /// @param amount Amount to release
  function releaseOrMint(
    bytes memory,
    address,
    uint256 amount,
    uint64 remoteChainSelector,
    SourceTokenData memory,
    bytes memory
  ) external override whenHealthy onlyOffRamp(remoteChainSelector) returns (address, uint256) {
    emit SynthMinted(amount);
    return (address(i_token), amount);
  }
}
