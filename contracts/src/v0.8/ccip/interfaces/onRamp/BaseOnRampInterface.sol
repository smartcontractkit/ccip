// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {AllowListInterface} from "../access/AllowListInterface.sol";
import {PoolInterface, IERC20} from "../pools/PoolInterface.sol";

interface BaseOnRampInterface is AllowListInterface {
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error MessageGasLimitTooHigh();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error RouterNotSet();
  // Token pools
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error PoolDoesNotExist(IERC20 token);
  error TokenPoolMismatch();

  event RouterSet(address router);
  event OnRampConfigSet(OnRampConfig config);

  // Token pools
  event PoolAdded(IERC20 token, PoolInterface pool);
  event PoolRemoved(IERC20 token, PoolInterface pool);

  struct OnRampConfig {
    // Fee for sending message taken in this contract
    uint64 commitFeeJuels;
    // maximum payload data size
    uint64 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint64 maxTokensLength;
    // Maximum gasLimit for messages targeting EVMs
    uint64 maxGasLimit;
  }

  struct PoolConfig {
    PoolInterface pool;
    bool enabled;
  }

  /**
   * @notice Get the pool for a specific token
   * @param sourceToken The source chain token to get the pool for
   * @return pool PoolInterface
   */
  function getPoolBySourceToken(IERC20 sourceToken) external returns (PoolInterface);

  /**
   * @notice Get all configured source tokens
   * @return Array of configured source tokens
   */
  function getPoolTokens() external view returns (IERC20[] memory);

  /**
   * @notice Gets the next sequence number to be used in the onRamp
   * @return the next sequence number to be used
   */
  function getExpectedNextSequenceNumber() external view returns (uint64);

  /**
   * @notice Sets the router to the given router
   * @param router The new router
   */
  function setRouter(address router) external;

  /**
   * @notice Gets the configured router
   * @return The set router
   */
  function getRouter() external view returns (address);

  /**
   * @notice Sets the onRamp config to the given OnRampConfig object
   * @param config The new OnRampConfig
   */
  function setConfig(OnRampConfig calldata config) external;

  /**
   * @notice Gets the current onRamp configuration
   * @return config The current configuration
   */
  function getConfig() external view returns (OnRampConfig memory config);
}
