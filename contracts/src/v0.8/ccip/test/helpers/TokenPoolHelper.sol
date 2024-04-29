// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Pool} from "../../libraries/Pool.sol";
import "../../pools/TokenPool.sol";

contract TokenPoolHelper is TokenPool {
  event LockOrBurn(uint256 amount);
  event ReleaseOrMint(address indexed recipient, uint256 amount);
  event AssertionPassed();

  constructor(
    IERC20 token,
    address[] memory allowlist,
    address armProxy,
    address router
  ) TokenPool(token, allowlist, armProxy, router) {}

  function lockOrBurn(bytes calldata lockOrBurnIn) external override returns (bytes memory) {
    Pool.LockOrBurnInV1 memory lockOrBurnData = Pool._decodeLockOrBurnInV1(lockOrBurnIn);
    emit LockOrBurn(lockOrBurnData.amount);
    return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), "");
  }

  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64,
    IPool.SourceTokenData memory,
    bytes memory
  ) external override returns (address, uint256) {
    emit ReleaseOrMint(receiver, amount);
    return (address(i_token), amount);
  }

  function onlyOnRampModifier(uint64 remoteChainSelector) external view {
    _onlyOnRamp(remoteChainSelector);
  }

  function onlyOffRampModifier(uint64 remoteChainSelector) external onlyOffRamp(remoteChainSelector) {}
}
