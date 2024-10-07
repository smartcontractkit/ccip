// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IBurnMintERC20} from "../../shared/token/ERC20/IBurnMintERC20.sol";

import {Pool} from "../libraries/Pool.sol";
import {BurnWithFromMintTokenPool} from "./BurnWithFromMintTokenPool.sol";

import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {TokenPool} from "./TokenPool.sol";

/// @notice This pool mints and burns a 3rd-party token.
/// @dev Pool whitelisting mode is set in the constructor and cannot be modified later.
/// It either accepts any address as originalSender, or only accepts whitelisted originalSender.
/// The only way to change whitelisting mode is to deploy a new pool.
/// If that is expected, please make sure the token's burner/minter roles are adjustable.
/// @dev This contract is a variant of BurnMintTokenPool that uses `burn(from, amount)`.
contract BurnWithFromMintRebasingTokenPool is BurnWithFromMintTokenPool {
  using SafeERC20 for IBurnMintERC20;

  error NegativeMintAmount(uint256 amountBurned);

  string public constant override typeAndVersion = "BurnWithFromMintRebasingTokenPool 1.5.0";

  constructor(
    IBurnMintERC20 token,
    address[] memory allowlist,
    address rmnProxy,
    address router
  ) BurnWithFromMintTokenPool(token, allowlist, rmnProxy, router) {}

  /// @notice Mint tokens from the pool to the recipient
  /// @dev The _validateReleaseOrMint check is an essential security check
  function releaseOrMint(
    Pool.ReleaseOrMintInV1 calldata releaseOrMintIn
  ) external virtual override returns (Pool.ReleaseOrMintOutV1 memory) {
    _validateReleaseOrMint(releaseOrMintIn);

    uint256 balancePre = IBurnMintERC20(address(i_token)).balanceOf(releaseOrMintIn.receiver);

    // Mint to the receiver
    IBurnMintERC20(address(i_token)).mint(releaseOrMintIn.receiver, releaseOrMintIn.amount);

    uint256 balancePost = IBurnMintERC20(address(i_token)).balanceOf(releaseOrMintIn.receiver);

    // Mint should not reduce the number of tokens in the receiver, if it does it will revert the call.
    if (balancePost < balancePre) {
      revert NegativeMintAmount(balancePre - balancePost);
    }

    emit Minted(msg.sender, releaseOrMintIn.receiver, balancePost - balancePre);

    return Pool.ReleaseOrMintOutV1({destinationAmount: balancePost - balancePre});
  }
}
