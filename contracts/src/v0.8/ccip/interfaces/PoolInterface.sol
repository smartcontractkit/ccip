// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../../vendor/IERC20.sol";
import "../utils/TokenLimits.sol";

// Shared public interface for multiple pool types.
// Each pool type handles a different child token model (lock/unlock, mint/burn.)
interface PoolInterface {
  error ExceedsTokenLimit(uint256 currentLimit, uint256 requested);

  event Locked(address indexed sender, address indexed depositor, uint256 amount);
  event Burned(address indexed sender, address indexed depositor, uint256 amount);
  event Released(address indexed sender, address indexed recipient, uint256 amount);
  event Minted(address indexed sender, address indexed recipient, uint256 amount);
  event NewLockBurnBucketConstructed(uint256 rate, uint256 capacity, bool full);
  event NewReleaseMintBucketConstructed(uint256 rate, uint256 capacity, bool full);

  /**
   * @notice Lock or burn the token in the pool
   * @param depositor Token holder address
   * @param amount Amount to lock or burn
   */
  function lockOrBurn(address depositor, uint256 amount) external;

  /**
   * @notice Release or mint tokens fromm the pool to the recipient
   * @param recipient Recipient address
   * @param amount Amount to release or mint
   */
  function releaseOrMint(address recipient, uint256 amount) external;

  function getToken() external view returns (IERC20 pool);

  function setLockOrBurnBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external;

  function setReleaseOrMintBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external;

  function getLockOrBurnBucket() external view returns (TokenLimits.TokenBucket memory);

  function getReleaseOrMintBucket() external view returns (TokenLimits.TokenBucket memory);

  function pause() external;

  function unpause() external;
}
