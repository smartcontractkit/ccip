// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IERC20} from "../../../vendor/IERC20.sol";

// Shared public interface for multiple pool types.
// Each pool type handles a different child token model (lock/unlock, mint/burn.)
interface IPool {
  error ExceedsTokenLimit(uint256 currentLimit, uint256 requested);
  error PermissionsError();
  error NullAddressNotAllowed();

  event Locked(address indexed sender, uint256 amount);
  event Burned(address indexed sender, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  /// @notice Lock or burn the token in the pool.
  /// @param amount Amount to lock or burn.
  function lockOrBurn(uint256 amount, address originalSender) external;

  /// @notice Release or mint tokens from the pool to the recipient.
  /// @param recipient Recipient address.
  /// @param amount Amount to release or mint.
  function releaseOrMint(address recipient, uint256 amount) external;

  /// @notice Gets the IERC20 token that this pool can lock or burn.
  /// @return token The IERC20 token representation.
  function getToken() external view returns (IERC20 token);

  /// @notice Pauses the token pool.
  function pause() external;

  /// @notice Unpauses the token pool.
  function unpause() external;
}
