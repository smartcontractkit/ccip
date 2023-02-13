// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IEVM2AnyOnRamp} from "./IEVM2AnyOnRamp.sol";
import {Internal} from "../../models/Internal.sol";
import {IAllowList} from "../access/IAllowList.sol";
import {IERC20} from "../../../vendor/IERC20.sol";
import {IPool} from "../pools/IPool.sol";

interface IEVM2EVMOnRamp is IEVM2AnyOnRamp, IAllowList {
  error InvalidExtraArgsTag(bytes4 expected, bytes4 got);
  error OnlyCallableByOwnerOrFeeAdmin();
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error MessageGasLimitTooHigh();
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();
  error RouterNotSet();
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error PoolDoesNotExist(IERC20 token);
  error TokenPoolMismatch();

  event FeeAdminSet(address feeAdmin);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event CCIPSendRequested(Internal.EVM2EVMMessage message);
  event RouterSet(address router);
  event OnRampConfigSet(OnRampConfig config);
  event PoolAdded(IERC20 token, IPool pool);
  event PoolRemoved(IERC20 token, IPool pool);

  function setFeeAdmin(address feeAdmin) external;

  /// @dev Struct to hold the fee configuration for a token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // -------┐ Token address
    uint64 multiplier; // ---┘ Price multiplier for gas costs
    uint96 feeAmount; // ------┐ Flat fee in feeToken
    uint32 destGasOverhead; //-┘ Extra gas charged on top of the gasLimit
  }

  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external;

  /**
   * @notice Get the pool for a specific token
   * @param sourceToken The source chain token to get the pool for
   * @return pool IPool
   */
  function getPoolBySourceToken(IERC20 sourceToken) external view returns (IPool);

  /**
   * @notice Get all supported source tokens
   * @return Array of supported source tokens
   */
  function getSupportedTokens() external view returns (address[] memory);

  /**
   * @notice Gets the next sequence number to be used in the onRamp
   * @return the next sequence number to be used
   */
  function getExpectedNextSequenceNumber() external view returns (uint64);

  /**
   * @notice Get the next nonce for a given sender
   * @param sender The sender to get the nonce for
   * @return nonce The next nonce for the sender
   */
  function getSenderNonce(address sender) external view returns (uint64 nonce);

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
   * @notice Gets the chain ID of the chain this contract is deployed on
   * @return chainID
   */
  function getChainId() external view returns (uint64);

  /**
   * @notice Get the destination chain ID
   * @return destinationChainID
   */
  function getDestinationChainId() external view returns (uint64);

  struct OnRampConfig {
    // maximum payload data size
    uint32 maxDataSize;
    // Maximum number of distinct ERC20 tokens that can be sent in a message
    uint16 maxTokensLength;
    // Maximum gasLimit for messages targeting EVMs
    uint64 maxGasLimit;
  }

  /**
   * @notice Sets the onRamp config to the given OnRampConfig object
   * @param config The new OnRampConfig
   */
  function setOnRampConfig(OnRampConfig calldata config) external;

  /**
   * @notice Gets the current onRamp configuration
   * @return config The current configuration
   */
  function getOnRampConfig() external view returns (OnRampConfig memory config);
}
