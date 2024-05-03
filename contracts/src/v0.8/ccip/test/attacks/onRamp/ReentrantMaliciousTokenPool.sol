// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Pool} from "../../../libraries/Pool.sol";
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
  function lockOrBurn(Pool.LockOrBurnInV1 calldata lockOrBurnIn)
    external
    override
    returns (Pool.LockOrBurnOutV1 memory)
  {
    if (s_attacked) {
      return Pool.LockOrBurnOutV1({destPoolAddress: getRemotePool(lockOrBurnIn.remoteChainSelector), destPoolData: ""});
    }

    s_attacked = true;

    FacadeClient(i_facade).send(lockOrBurnIn.amount);
    emit Burned(msg.sender, lockOrBurnIn.amount);
    return Pool.LockOrBurnOutV1({destPoolAddress: getRemotePool(lockOrBurnIn.remoteChainSelector), destPoolData: ""});
  }

  function releaseOrMint(Pool.ReleaseOrMintInV1 calldata releaseOrMintIn)
    external
    view
    override
    returns (Pool.ReleaseOrMintOutV1 memory)
  {
    return Pool.ReleaseOrMintOutV1({localToken: address(i_token), destinationAmount: releaseOrMintIn.amount});
  }
}
