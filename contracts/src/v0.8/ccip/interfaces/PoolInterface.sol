// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";

// Shared public interface for multiple pool types.
// Each pool type handles a different child token model (lock/unlock, mint/burn.)
interface PoolInterface {
  event Locked(address indexed sender, address indexed depositor, uint256 amount);
  event Burnt(address indexed sender, address indexed depositor, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);

  function lockOrBurn(address depositor, uint256 amount) external;

  function releaseOrMint(address recipient, uint256 amount) external;

  function getToken() external view returns (IERC20 pool);

  function pause() external;

  function unpause() external;
}
