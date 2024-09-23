pragma solidity ^0.8.0;

import {IBurnMintERC20} from "../../../shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../../libraries/Pool.sol";
import {BurnMintTokenPool} from "../BurnMintTokenPool.sol";

contract BurnMintWithLockReleaseFlagTokenPool is BurnMintTokenPool {
  /// bytes4(keccak256("NO_CCTP_USE_LOCK_RELEASE"))
  bytes4 public constant LOCK_RELEASE_FLAG = 0xfa7c07de;

  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) BurnMintTokenPool(token, allowlist, rmnProxy, router) {}

  /// @notice Burn the token in the pool
  /// @dev The _validateLockOrBurn check is an essential security check
  function lockOrBurn(
    Pool.LockOrBurnInV1 calldata lockOrBurnIn
  ) external virtual override returns (Pool.LockOrBurnOutV1 memory) {
    _validateLockOrBurn(lockOrBurnIn);

    _burn(lockOrBurnIn.amount);

    emit Burned(msg.sender, lockOrBurnIn.amount);

    return Pool.LockOrBurnOutV1({
      destTokenAddress: getRemoteToken(lockOrBurnIn.remoteChainSelector),
      destPoolData: abi.encode(LOCK_RELEASE_FLAG)
    });
  }
}
