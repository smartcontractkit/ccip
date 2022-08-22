// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "../../vendor/Pausable.sol";
import "../access/OwnerIsCreator.sol";
import "../onRamp/BaseOnRamp.sol";
import "../interfaces/onRamp/BaseOnRampInterface.sol";
import "../interfaces/offRamp/Any2EVMOffRampInterface.sol";

/**
 * @notice Base abstract class with common functions for all token pools
 */
abstract contract TokenPool is PoolInterface, OwnerIsCreator, Pausable {
  using TokenLimits for TokenLimits.TokenBucket;

  IERC20 internal immutable i_token;
  mapping(BaseOnRampInterface => bool) internal s_onRamps;
  mapping(BaseOffRampInterface => bool) internal s_offRamps;
  TokenLimits.TokenBucket internal s_lockOrBurnBucket;
  TokenLimits.TokenBucket internal s_releaseOrMintBucket;

  error PermissionsError();

  constructor(
    IERC20 token,
    BucketConfig memory lockOrBurnConfig,
    BucketConfig memory releaseOrMintConfig
  ) {
    i_token = token;
    s_lockOrBurnBucket = TokenLimits.constructTokenBucket(lockOrBurnConfig.rate, lockOrBurnConfig.capacity, true);
    s_releaseOrMintBucket = TokenLimits.constructTokenBucket(
      releaseOrMintConfig.rate,
      releaseOrMintConfig.capacity,
      true
    );
  }

  /**
   * @notice Pause the pool
   * @dev Only callable by the owner
   */
  function pause() external override onlyOwner {
    _pause();
  }

  /**
   * @notice Unpause the pool
   * @dev Only callable by the owner
   */
  function unpause() external override onlyOwner {
    _unpause();
  }

  /**
   * @notice Set an onRamp's permissions
   * @dev Only callable by the owner
   * @param onRamp The onRamp
   * @param permission Whether or not the onRamp has onRamp permissions on this contract
   */
  function setOnRamp(BaseOnRampInterface onRamp, bool permission) public onlyOwner {
    s_onRamps[onRamp] = permission;
  }

  /**
   * @notice Set an offRamp's permissions
   * @dev Only callable by the owner
   * @param offRamp The offRamp
   * @param permission Whether or not the offRamp has offRamp permissions on this contract
   */
  function setOffRamp(BaseOffRampInterface offRamp, bool permission) public onlyOwner {
    s_offRamps[offRamp] = permission;
  }

  /**
   * @notice Set the lock or burn bucket by constructing a new bucket with new parameters
   * @dev Only callable by the owner
   * @param rate The rate of refill
   * @param capacity Full capacity of the bucket
   * @param full Start full?
   */
  function setLockOrBurnBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external override onlyOwner {
    s_lockOrBurnBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewLockBurnBucketConstructed(rate, capacity, full);
  }

  /**
   * @notice Set the release or mint bucket by constructing a new bucket with new parameters
   * @dev Only callable by the owner
   * @param rate The rate of refill
   * @param capacity Full capacity of the bucket
   * @param full Start full?
   */
  function setReleaseOrMintBucket(
    uint256 rate,
    uint256 capacity,
    bool full
  ) external override onlyOwner {
    s_releaseOrMintBucket = TokenLimits.constructTokenBucket(rate, capacity, full);
    emit NewReleaseMintBucketConstructed(rate, capacity, full);
  }

  /**
   * @notice Get the lock or burn bucket
   * @return bucket
   */
  function getLockOrBurnBucket() external view override returns (TokenLimits.TokenBucket memory) {
    return s_lockOrBurnBucket;
  }

  /**
   * @notice Get the mint or release bucket
   * @return bucket
   */
  function getReleaseOrMintBucket() external view override returns (TokenLimits.TokenBucket memory) {
    return s_releaseOrMintBucket;
  }

  /**
   * @notice Checks whether something is a permissioned onRamp on this contract
   * @return boolean
   */
  function isOnRamp(BaseOnRampInterface onRamp) public view returns (bool) {
    return s_onRamps[onRamp];
  }

  /**
   * @notice Checks whether something is a permissioned offRamp on this contract
   * @return boolean
   */
  function isOffRamp(BaseOffRampInterface offRamp) public view returns (bool) {
    return s_offRamps[offRamp];
  }

  /**
   * @notice Gets the underlying token
   * @return token
   */
  function getToken() public view override returns (IERC20 token) {
    return i_token;
  }

  /**
   * @notice Checks whether the msg.sender is either the owner, or a permissioned onRamp on this contract
   * @dev Reverts with a PermissionsError if check fails
   */
  function _validateOwnerOrOnRamp() internal view {
    if (msg.sender != owner() && !isOnRamp(Any2EVMTollOnRampInterface(msg.sender))) revert PermissionsError();
  }

  /**
   * @notice Checks whether the msg.sender is either the owner, or a permissioned offRamp on this contract
   * @dev Reverts with a PermissionsError if check fails
   */
  function _validateOwnerOrOffRamp() internal view {
    if (msg.sender != owner() && !isOffRamp(BaseOffRampInterface(msg.sender))) revert PermissionsError();
  }

  /**
   * @notice Checks whether the amount requested to lock or burn is within the lock or burn token bucket limit
   * @dev NOTE: this mutates s_lockOrBurnBucket in storage
   * @dev Reverts with a ExceedsTokenLimit is check fails
   */
  function _verifyLockOrBurnWithinLimits(uint256 amount) internal {
    // Alters the s_lockOrBurnBucket in storage
    if (!s_lockOrBurnBucket.remove(amount)) revert ExceedsTokenLimit(s_lockOrBurnBucket.tokens, amount);
  }

  /**
   * @notice Checks whether the amount requested to mint or release is within the mint or release token bucket limit
   * @dev NOTE: this mutates s_releaseOrMintBucket in storage
   * @dev Reverts with a ExceedsTokenLimit is check fails
   */
  function _verifyMintOrReleaseWithinLimits(uint256 amount) internal {
    // Alters the s_releaseOrMintBucket in storage
    if (!s_releaseOrMintBucket.remove(amount)) revert ExceedsTokenLimit(s_releaseOrMintBucket.tokens, amount);
  }

  /**
   * @notice Check permissions and limits of a lock or burn
   */
  modifier assertLockOrBurn(uint256 amount) {
    _validateOwnerOrOnRamp();
    _verifyLockOrBurnWithinLimits(amount);
    _;
  }

  /**
   * @notice Check permissions and limits of a lock or burn
   */
  modifier assertMintOrRelease(uint256 amount) {
    _validateOwnerOrOffRamp();
    _verifyMintOrReleaseWithinLimits(amount);
    _;
  }
}
