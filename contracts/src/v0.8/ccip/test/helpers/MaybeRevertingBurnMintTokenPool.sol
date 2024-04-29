// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";
import {IPool} from "../../interfaces/IPool.sol";

import {Pool} from "../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";

contract MaybeRevertingBurnMintTokenPool is BurnMintTokenPool {
  bytes public s_revertReason = "";
  bytes public s_sourceTokenData = "";

  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    address armProxy,
    address router
  ) BurnMintTokenPool(token, allowlist, armProxy, router) {}

  function setShouldRevert(bytes calldata revertReason) external {
    s_revertReason = revertReason;
  }

  function setSourceTokenData(bytes calldata sourceTokenData) external {
    s_sourceTokenData = sourceTokenData;
  }

  function lockOrBurn(bytes calldata lockOrBurnIn) external virtual override whenHealthy returns (bytes memory) {
    Pool.LockOrBurnInV1 memory lockOrBurnData = Pool._decodeLockOrBurnInV1(lockOrBurnIn);
    _checkAllowList(lockOrBurnData.originalSender);
    _onlyOnRamp(lockOrBurnData.remoteChainSelector);
    _consumeOutboundRateLimit(lockOrBurnData.remoteChainSelector, lockOrBurnData.amount);

    bytes memory revertReason = s_revertReason;
    if (revertReason.length != 0) {
      assembly {
        revert(add(32, revertReason), mload(revertReason))
      }
    }

    IBurnMintERC20(address(i_token)).burn(lockOrBurnData.amount);
    emit Burned(msg.sender, lockOrBurnData.amount);
    return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), s_sourceTokenData);
  }

  /// @notice Reverts depending on the value of `s_revertReason`
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector,
    IPool.SourceTokenData memory sourceTokenData,
    bytes memory
  ) external virtual override whenHealthy onlyOffRamp(remoteChainSelector) returns (address, uint256) {
    _validateSourceCaller(remoteChainSelector, sourceTokenData.sourcePoolAddress);
    bytes memory revertReason = s_revertReason;
    if (revertReason.length != 0) {
      assembly {
        revert(add(32, revertReason), mload(revertReason))
      }
    }
    _consumeInboundRateLimit(remoteChainSelector, amount);
    IBurnMintERC20(address(i_token)).mint(receiver, amount);
    emit Minted(msg.sender, receiver, amount);
    return (address(i_token), amount);
  }
}
