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
  error OnlyCallableByOwnerOrFeeAdminOrNop();
  error InvalidWithdrawalAddress(address addr);
  error InvalidFeeToken(address token);
  error NoFeesToPay();
  error NoNopsToPay();
  error InsufficientBalance();
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
  error TokenOrChainNotSupported(address token, uint64 chain);
  error NullAddressNotAllowed();

  event FeeAdminSet(address feeAdmin);
  event NopPaid(address indexed nop, uint256 amount);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event CCIPSendRequested(Internal.EVM2EVMMessage message);
  event NopsSet(uint256 nopWeightsTotal, NopAndWeight[] nopsAndWeights);
  event RouterSet(address router);
  event OnRampConfigSet(OnRampConfig config);
  event PoolAdded(IERC20 token, IPool pool);
  event PoolRemoved(IERC20 token, IPool pool);

  struct Chains {
    uint64 chainId; // -------┐ Source chain Id
    uint64 destChainId; // ---┘ Destination chain Id
  }

  /// @dev Struct to hold the fee configuration for a token
  struct FeeTokenConfig {
    uint96 feeAmount; // --------┐ Flat fee
    uint64 multiplier; //        | Price multiplier for gas costs
    uint32 destGasOverhead; // --┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Struct to hold the fee configuration for a token, same as the FeeTokenConfig but with
  /// token included so that an array of these can be passed in to setFeeConfig to set the mapping
  struct FeeTokenConfigArgs {
    address token; // ---------┐ Token address
    uint64 multiplier; // -----┘ Price multiplier for gas costs
    uint96 feeAmount; // ------┐ Flat fee in feeToken
    uint32 destGasOverhead; //-┘ Extra gas charged on top of the gasLimit
  }

  /// @dev Nop address and weight, used to set the nops and their weights
  struct NopAndWeight {
    address nop;
    uint256 weight;
  }

  /// @dev Struct to hold the onRamp configuration
  struct OnRampConfig {
    uint32 maxDataSize; // ----┐ Maximum payload data size
    uint16 maxTokensLength; // | Maximum number of distinct ERC20 tokens that can be sent per message
    uint64 maxGasLimit; // ----┘ Maximum gas limit for messages targeting EVMs
  }

  /// @notice Sets the Nops and their weights
  /// @param nopsAndWeights Array of NopAndWeight structs
  function setNops(NopAndWeight[] calldata nopsAndWeights) external;

  /// @notice Gets the Nops and their weights
  /// @return nopsAndWeights Array of NopAndWeight structs
  /// @return weightsTotal The sum weight of all Nops
  function getNops() external view returns (NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal);

  /// @notice Allows the owner to withdraw any ERC20 token that is not the fee token
  /// @param feeToken The token to withdraw
  /// @param to The address to send the tokens to
  function withdrawNonLinkFees(address feeToken, address to) external;

  /// @notice Pays the Nops
  function payNops() external;

  /// @notice Sets the fee admin
  /// @param feeAdmin The address of the fee admin
  function setFeeAdmin(address feeAdmin) external;

  /// @notice Sets the fee configuration for a token
  /// @param feeTokenConfigs Array of FeeTokenConfigArgs structs
  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external;

  /// @notice Gets the fee configuration for a token
  /// @param token The token to get the fee configuration for
  /// @return feeTokenConfig FeeTokenConfig struct
  function getFeeConfig(address token) external returns (FeeTokenConfig memory);

  /// @notice Get the pool for a specific token
  /// @param sourceToken The source chain token to get the pool for
  /// @return pool IPool
  function getPoolBySourceToken(IERC20 sourceToken) external view returns (IPool);

  /// @notice Get all supported source tokens
  /// @return Array of supported source tokens
  function getSupportedTokens() external view returns (address[] memory);

  /// @notice Gets the next sequence number to be used in the onRamp
  /// @return the next sequence number to be used
  function getExpectedNextSequenceNumber() external view returns (uint64);

  /// @notice Get the next nonce for a given sender
  /// @param sender The sender to get the nonce for
  /// @return nonce The next nonce for the sender
  function getSenderNonce(address sender) external view returns (uint64 nonce);

  /// @notice Sets the router to the given router
  /// @param router The new router
  function setRouter(address router) external;

  /// @notice Gets the configured router
  /// @return The set router
  function getRouter() external view returns (address);

  /// @notice Gets the chain ID of the chain this contract is deployed on
  /// @return chainID
  function getChainId() external view returns (uint64);

  /// @notice Get the destination chain ID
  /// @return destinationChainID
  function getDestChainId() external view returns (uint64);

  /// @notice Get the total amount of fees to be paid to the Nops (in LINK)
  /// @return totalNopFees
  function getNopFeesJuels() external view returns (uint256);

  /// @notice Sets the onRamp config to the given OnRampConfig object
  /// @param config The new OnRampConfig
  function setOnRampConfig(OnRampConfig calldata config) external;

  /// @notice Gets the current onRamp configuration
  /// @return config The current configuration
  function getOnRampConfig() external view returns (OnRampConfig memory config);
}
