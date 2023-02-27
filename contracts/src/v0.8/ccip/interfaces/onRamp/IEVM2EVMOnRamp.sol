// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IAllowList} from "../access/IAllowList.sol";
import {IPool} from "../pools/IPool.sol";

import {Internal} from "../../models/Internal.sol";
import {IEVM2AnyOnRamp} from "./IEVM2AnyOnRamp.sol";

import {IERC20} from "../../../vendor/IERC20.sol";

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
  error InvalidTokenPoolConfig();
  error PoolAlreadyAdded();
  error PoolDoesNotExist(IERC20 token);
  error TokenPoolMismatch();
  error TokenOrChainNotSupported(address token, uint64 chain);
  error InvalidConfig();

  event StaticConfigSet(StaticConfig);
  event DynamicConfigSet(DynamicConfig);
  event NopPaid(address indexed nop, uint256 amount);
  event FeeConfigSet(FeeTokenConfigArgs[] feeConfig);
  event CCIPSendRequested(Internal.EVM2EVMMessage message);
  event NopsSet(uint256 nopWeightsTotal, NopAndWeight[] nopsAndWeights);
  event PoolAdded(IERC20 token, IPool pool);
  event PoolRemoved(IERC20 token, IPool pool);

  /// @dev Struct that contains the static configuration
  struct StaticConfig {
    address linkToken; // --------┐ Link token address
    uint64 chainId; // -----------┘ Source chain Id
    uint64 destChainId; // -------┐ Destination chain Id
    uint64 defaultTxGasLimit; // -┘ Default gas limit for a tx
  }

  /// @dev Struct to contains the dynamic configuration
  struct DynamicConfig {
    address router; //            Router address
    address priceRegistry; // --┐ Price registry address
    uint32 maxDataSize; //      | Maximum payload data size
    uint64 maxGasLimit; // -----┘ Maximum gas limit for messages targeting EVMs
    uint16 maxTokensLength; // -┐ Maximum number of distinct ERC20 tokens that can be sent per message
    address feeAdmin; // -------┘ Fee admin address
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

  /// @notice Gets the Nops and their weights
  /// @return nopsAndWeights Array of NopAndWeight structs
  /// @return weightsTotal The sum weight of all Nops
  function getNops() external view returns (NopAndWeight[] memory nopsAndWeights, uint256 weightsTotal);

  /// @notice Sets the Nops and their weights
  /// @param nopsAndWeights Array of NopAndWeight structs
  function setNops(NopAndWeight[] calldata nopsAndWeights) external;

  /// @notice Allows the owner to withdraw any ERC20 token that is not the fee token
  /// @param feeToken The token to withdraw
  /// @param to The address to send the tokens to
  function withdrawNonLinkFees(address feeToken, address to) external;

  /// @notice Pays the Nops
  function payNops() external;

  /// @notice Get the total amount of fees to be paid to the Nops (in LINK)
  /// @return totalNopFees
  function getNopFeesJuels() external view returns (uint256);

  /// @notice Returns the static onRamp config.
  /// @return the configuration.
  function getStaticConfig() external view returns (StaticConfig memory);

  /// @notice Returns the dynamic onRamp config.
  /// @return the configuration.
  function getDynamicConfig() external view returns (DynamicConfig memory);

  /// @notice Sets the dynamic configuration.
  /// @param dynamicConfig The configuration.
  function setDynamicConfig(DynamicConfig memory dynamicConfig) external;

  /// @notice Gets the fee configuration for a token
  /// @param token The token to get the fee configuration for
  /// @return feeTokenConfig FeeTokenConfig struct
  function getFeeConfig(address token) external returns (FeeTokenConfig memory);

  /// @notice Sets the fee configuration for a token
  /// @param feeTokenConfigs Array of FeeTokenConfigArgs structs
  function setFeeConfig(FeeTokenConfigArgs[] calldata feeTokenConfigs) external;
}
