// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Pool} from "../libraries/Pool.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {IERC165} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/introspection/IERC165.sol";

// Shared public interface for multiple pool types.
// Each pool type handles a different child token model (lock/unlock, mint/burn.)
interface IPool is IERC165 {
  struct SourceTokenData {
    bytes sourcePoolAddress;
    bytes destPoolAddress;
    bytes extraData;
  }

  /// @notice Lock tokens into the pool or burn the tokens.
  /// @param lockOrBurnIn Encoded data fields for the processing of tokens on the source chain.
  /// @return lockOrBurnOut Encoded data fields for the processing of tokens on the destination chain.
  function lockOrBurn(Pool.LockOrBurnInV1 calldata lockOrBurnIn)
    external
    returns (Pool.LockOrBurnOutV1 memory lockOrBurnOut);

  /// @notice Releases or mints tokens to the receiver address.
  /// * localToken The address of the local token.
  /// * destinationAmount The amount of tokens released or minted on the destination chain,
  /// denominated in the local token's decimals.
  function releaseOrMint(Pool.ReleaseOrMintInV1 calldata releaseOrMintIn)
    external
    returns (Pool.ReleaseOrMintOutV1 memory);
}
