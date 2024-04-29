// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBurnMintERC20} from "../../shared/token/ERC20/IBurnMintERC20.sol";
import {IPool} from "../interfaces/IPool.sol";

import {Pool} from "../libraries/Pool.sol";
import {TokenPool} from "./TokenPool.sol";

abstract contract BurnMintTokenPoolAbstract is TokenPool {
  /// @notice Contains the specific burn call for a pool.
  /// @dev overriding this method allows us to create pools with different burn signatures
  /// without duplicating the underlying logic.
  function _burn(uint256 amount) internal virtual;

  /// @notice Burn the token in the pool
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function lockOrBurn(bytes calldata lockOrBurnIn)
    external
    virtual
    override
    whenHealthy
    returns (Pool.LockOrBurnOutV1 memory)
  {
    Pool.LockOrBurnInV1 memory lockOrBurnData = Pool._decodeLockOrBurnInV1(lockOrBurnIn);
    _checkAllowList(lockOrBurnData.originalSender);
    _onlyOnRamp(lockOrBurnData.remoteChainSelector);
    _consumeOutboundRateLimit(lockOrBurnData.remoteChainSelector, lockOrBurnData.amount);

    _burn(lockOrBurnData.amount);

    emit Burned(msg.sender, lockOrBurnData.amount);

    return Pool._encodeLockOrBurnOutV1(getRemotePool(lockOrBurnData.remoteChainSelector), "");
  }

  /// @notice Mint tokens from the pool to the recipient
  /// @param receiver Recipient address
  /// @param amount Amount to mint
  /// @dev The whenHealthy check is important to ensure that even if a ramp is compromised
  /// we're able to stop token movement via ARM.
  function releaseOrMint(
    bytes memory,
    address receiver,
    uint256 amount,
    uint64 remoteChainSelector,
    IPool.SourceTokenData memory sourceTokenData,
    bytes memory
  ) external virtual override whenHealthy returns (address, uint256) {
    _onlyOffRamp(remoteChainSelector);
    _validateSourceCaller(remoteChainSelector, sourceTokenData.sourcePoolAddress);
    _consumeInboundRateLimit(remoteChainSelector, amount);

    IBurnMintERC20(address(i_token)).mint(receiver, amount);

    emit Minted(msg.sender, receiver, amount);

    return (address(i_token), amount);
  }
}
