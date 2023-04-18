// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../interfaces/TypeAndVersionInterface.sol";
import {IRouterClient} from "./interfaces/IRouterClient.sol";
import {IRouter} from "./interfaces/IRouter.sol";
import {IEVM2AnyOnRamp} from "./interfaces/IEVM2AnyOnRamp.sol";
import {IWrappedNative} from "./interfaces/IWrappedNative.sol";
import {IAny2EVMMessageReceiver} from "./interfaces/IAny2EVMMessageReceiver.sol";

import {Client} from "./libraries/Client.sol";
import {Internal} from "./libraries/Internal.sol";
import {OwnerIsCreator} from "./OwnerIsCreator.sol";

import {SafeERC20} from "../vendor/SafeERC20.sol";
import {IERC20} from "../vendor/IERC20.sol";

/// @title Guaranteed Execution Router
/// @notice This is the entry point for the end user wishing to send a cross
/// chain message.
/// @dev This contract is used as a router for both on-ramps and off-ramps
contract Router is IRouter, IRouterClient, TypeAndVersionInterface, OwnerIsCreator {
  using SafeERC20 for IERC20;

  event OnRampSet(uint64 indexed destChainId, address onRamp);
  event OffRampAdded(uint64 indexed sourceChainId, address offRamp);
  event OffRampRemoved(uint64 indexed sourceChainId, address offRamp);

  struct OnRampUpdate {
    uint64 destChainId; // --┐  Destination chain Id.
    address onRamp; // ------┘  OnRamp address that is allowed to use this router.
  }
  struct OffRampUpdate {
    uint64 sourceChainId; //    Source chain Id.
    address[] offRamps; //      List of offRamps that are allowed to use this router.
  }

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "Router 1.0.0";
  // The minimum amount of gas to perform the call with exact gas
  uint64 private constant GAS_FOR_CALL_EXACT_CHECK = 5_000;

  // DYNAMIC CONFIG
  address private s_wrappedNative;
  // destChainId => onRamp address
  // Only ever one onRamp enabled at a time for a given destChainId.
  mapping(uint256 => address) private s_onRamps;
  // Can be multiple offRamps enabled at a time for a given sourceChainId.
  mapping(uint256 => address[]) private s_offRamps;
  // Mapping of offRamps to source chain ids
  mapping(address => uint256) private s_offRampSourceChainIds;

  constructor(address wrappedNative) {
    // Zero address indicates unsupported auto-wrapping.
    s_wrappedNative = wrappedNative;
  }

  // ================================================================
  // |                       Message sending                        |
  // ================================================================

  /// @inheritdoc IRouterClient
  /// @dev returns 0 fee on invalid message.
  function getFee(uint64 destinationChainId, Client.EVM2AnyMessage memory message) external view returns (uint256 fee) {
    if (message.feeToken == address(0)) {
      // For empty feeToken return native quote.
      message.feeToken = address(s_wrappedNative);
    }
    address onRamp = s_onRamps[destinationChainId];
    if (onRamp == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    return IEVM2AnyOnRamp(onRamp).getFee(message);
  }

  /// @inheritdoc IRouterClient
  function getSupportedTokens(uint64 chainId) external view returns (address[] memory) {
    if (!isChainSupported(chainId)) {
      return new address[](0);
    }
    return IEVM2AnyOnRamp(s_onRamps[uint256(chainId)]).getSupportedTokens();
  }

  /// @inheritdoc IRouterClient
  function isChainSupported(uint64 chainId) public view returns (bool) {
    return s_onRamps[chainId] != address(0);
  }

  /// @inheritdoc IRouterClient
  function ccipSend(uint64 destinationChainId, Client.EVM2AnyMessage memory message)
    external
    payable
    returns (bytes32)
  {
    address onRamp = s_onRamps[destinationChainId];
    if (onRamp == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    uint256 feeTokenAmount;
    // address(0) signals payment in true native
    if (message.feeToken == address(0)) {
      // for fee calculation we check the wrapped native price as we wrap
      // as part of the native fee coin payment.
      message.feeToken = s_wrappedNative;
      feeTokenAmount = IEVM2AnyOnRamp(onRamp).getFee(message);
      // Ensure sufficient native.
      if (msg.value < feeTokenAmount) revert InsufficientFeeTokenAmount();
      // Wrap and send native payment.
      // Note we take the whole msg.value regardless if its larger.
      feeTokenAmount = msg.value;
      IWrappedNative(message.feeToken).deposit{value: feeTokenAmount}();
      IERC20(message.feeToken).safeTransferFrom(address(this), onRamp, feeTokenAmount);
    } else {
      if (msg.value > 0) revert InvalidMsgValue();
      feeTokenAmount = IEVM2AnyOnRamp(onRamp).getFee(message);
      IERC20(message.feeToken).safeTransferFrom(msg.sender, onRamp, feeTokenAmount);
    }

    // Transfer the tokens to the token pools.
    for (uint256 i = 0; i < message.tokenAmounts.length; ++i) {
      IERC20 token = IERC20(message.tokenAmounts[i].token);
      token.safeTransferFrom(
        msg.sender,
        address(IEVM2AnyOnRamp(onRamp).getPoolBySourceToken(token)),
        message.tokenAmounts[i].amount
      );
    }

    return IEVM2AnyOnRamp(onRamp).forwardFromRouter(message, feeTokenAmount, msg.sender);
  }

  // ================================================================
  // |                      Message execution                       |
  // ================================================================

  /// @inheritdoc IRouter
  function routeMessage(
    Client.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external override onlyOffRamp(message.sourceChainId) returns (bool success) {
    bytes memory callData = abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message);
    if (manualExecution) {
      // solhint-disable-next-line avoid-low-level-calls
      (success, ) = receiver.call(callData);
    } else {
      success = _callWithExactGas(gasLimit, receiver, 0, callData);
    }
  }

  /// @dev calls target address with exactly gasAmount gas and data as calldata
  /// @param gasAmount gas limit for this call
  /// @param target target address
  /// @param value call ether value
  /// @param data calldata
  function _callWithExactGas(
    uint256 gasAmount,
    address target,
    uint256 value,
    bytes memory data
  ) internal returns (bool success) {
    // solhint-disable-next-line no-inline-assembly
    assembly {
      let g := gas()
      // Compute g -= GAS_FOR_CALL_EXACT_CHECK and check for underflow
      // The gas actually passed to the callee is _min(gasAmount, 63//64*gas available).
      // We want to ensure that we revert if gasAmount >  63//64*gas available
      // as we do not want to provide them with less, however that check itself costs
      // gas.  GAS_FOR_CALL_EXACT_CHECK ensures we have at least enough gas to be able
      // to revert if gasAmount >  63//64*gas available.
      if lt(g, GAS_FOR_CALL_EXACT_CHECK) {
        revert(0, 0)
      }
      g := sub(g, GAS_FOR_CALL_EXACT_CHECK)
      // if g - g//64 <= gasAmount, revert
      // (we subtract g//64 because of EIP-150)
      if iszero(gt(sub(g, div(g, 64)), gasAmount)) {
        revert(0, 0)
      }
      // solidity calls check that a contract actually exists at the destination, so we do the same
      if iszero(extcodesize(target)) {
        revert(0, 0)
      }
      // call and return whether we succeeded. ignore return data
      // call(gas,addr,value,argsOffset,argsLength,retOffset,retLength)
      success := call(gasAmount, target, value, add(data, 0x20), mload(data), 0, 0)
    }
    return (success);
  }

  // ================================================================
  // |                           Config                             |
  // ================================================================

  /// @notice Gets the wrapped representation of the native fee coin.
  /// @return The address of the ERC20 wrapped native.
  function getWrappedNative() external view returns (address) {
    return s_wrappedNative;
  }

  /// @notice Sets a new wrapped native token.
  /// @param wrappedNative The address of the new wrapped native ERC20 token.
  function setWrappedNative(address wrappedNative) external onlyOwner {
    s_wrappedNative = wrappedNative;
  }

  /// @notice Get the onramp for a destination chain.
  /// @param destChainId The destination chain Id to get the onRamp for.
  /// @return The address of the onRamp.
  function getOnRamp(uint64 destChainId) external view returns (address) {
    return s_onRamps[destChainId];
  }

  /// @notice Get a list of offRamps for a source chain.
  function getOffRamps(uint64 sourceChainId) external view returns (address[] memory) {
    return s_offRamps[sourceChainId];
  }

  /// @notice Set applies a set of ingress and egress config updates.
  /// @dev only callable by owner.
  function applyRampUpdates(OnRampUpdate[] memory onRampUpdates, OffRampUpdate[] memory offRampUpdates)
    external
    onlyOwner
  {
    // Apply egress updates.
    // We permit zero address as way to disable egress.
    for (uint256 i = 0; i < onRampUpdates.length; ++i) {
      s_onRamps[onRampUpdates[i].destChainId] = onRampUpdates[i].onRamp;
      emit OnRampSet(onRampUpdates[i].destChainId, onRampUpdates[i].onRamp);
    }
    // Apply ingress updates.
    // We permit an empty list as a way to disable ingress.
    for (uint256 i = 0; i < offRampUpdates.length; ++i) {
      uint64 sourceChainId = offRampUpdates[i].sourceChainId;
      // For this source chain, clear all the existing offRamps.
      for (uint256 j = 0; j < s_offRamps[sourceChainId].length; ++j) {
        delete s_offRampSourceChainIds[s_offRamps[sourceChainId][j]];
        emit OffRampRemoved(sourceChainId, s_offRamps[sourceChainId][j]);
      }
      delete s_offRamps[sourceChainId];
      // For this source chain, add all the new offRamps (there may be zero)
      for (uint256 j = 0; j < offRampUpdates[i].offRamps.length; ++j) {
        s_offRampSourceChainIds[offRampUpdates[i].offRamps[j]] = sourceChainId;
        emit OffRampAdded(sourceChainId, offRampUpdates[i].offRamps[j]);
      }
      s_offRamps[sourceChainId] = offRampUpdates[i].offRamps;
    }
  }

  // ================================================================
  // |                           Access                             |
  // ================================================================

  // @notice only lets allowed offRamps execute
  modifier onlyOffRamp(uint64 sourceChainId) {
    if (s_offRampSourceChainIds[msg.sender] == 0 || s_offRampSourceChainIds[msg.sender] != sourceChainId)
      revert OnlyOffRamp();
    _;
  }
}
