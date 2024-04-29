// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IPool} from "../../../interfaces/IPool.sol";

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {FacadeClient} from "./FacadeClient.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract ReentrantMaliciousTokenPool is TokenPool {
  address private i_facade;

  bool private s_attacked;

  constructor(
    address facade,
    IERC20 token,
    address armProxy,
    address router
  ) TokenPool(token, new address[](0), armProxy, router) {
    i_facade = facade;
  }

  /// @dev Calls into Facade to reenter Router exactly 1 time
  function lockOrBurn(bytes calldata lockOrBurnIn) external override returns (bytes memory) {
    Pool.LockOrBurnInV1 memory lockOrBurnData = Pool._decodeLockOrBurnInV1(lockOrBurnIn);

    if (s_attacked) {
      return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), "");
    }

    s_attacked = true;

    FacadeClient(i_facade).send(lockOrBurnData.amount);
    emit Burned(msg.sender, lockOrBurnData.amount);
    return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), "");
  }

  function releaseOrMint(
    bytes memory,
    address,
    uint256 amount,
    uint64,
    IPool.SourceTokenData memory,
    bytes memory
  ) external view override returns (address, uint256) {
    return (address(i_token), amount);
  }
}
