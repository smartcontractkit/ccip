// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {IPool} from "../interfaces/pools/IPool.sol";
import {IGERouter} from "../interfaces/router/IGERouter.sol";
import {IBaseOnRampRouter} from "../interfaces/onRamp/IBaseOnRampRouter.sol";
import {IBaseOnRamp} from "../interfaces/onRamp/IBaseOnRamp.sol";
import {IEVM2AnyGEOnRamp} from "../interfaces/onRamp/IEVM2AnyGEOnRamp.sol";
import {IAny2EVMOffRampRouter} from "../interfaces/offRamp/IAny2EVMOffRampRouter.sol";
import {IAny2EVMMessageReceiver} from "../interfaces/applications/IAny2EVMMessageReceiver.sol";

import {GEConsumer} from "../models/GEConsumer.sol";
import {Internal} from "../models/Internal.sol";
import {Common} from "../models/Common.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";

import {SafeERC20} from "../../vendor/SafeERC20.sol";
import {IERC20} from "../../vendor/IERC20.sol";

contract GERouter is IGERouter, TypeAndVersionInterface, OwnerIsCreator {
  using SafeERC20 for IERC20;

  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GERouter 1.0.0";

  uint256 private constant GAS_FOR_CALL_EXACT_CHECK = 5_000;

  // Mapping from offRamp to allowed status
  mapping(address => OffRampDetails) internal s_offRamps;
  // List of all offRamps that have  OffRampDetails
  address[] internal s_offRampsList;

  // destination chain id => IOnRamp
  mapping(uint256 => IEVM2AnyGEOnRamp) private s_onRamps;

  constructor(address[] memory offRamps) {
    s_offRampsList = offRamps;
    for (uint256 i = 0; i < offRamps.length; ++i) {
      s_offRamps[offRamps[i]] = OffRampDetails({listIndex: uint96(i), allowed: true});
    }
  }

  /// @inheritdoc IGERouter
  function ccipSend(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message) external returns (bytes32) {
    IEVM2AnyGEOnRamp onRamp = s_onRamps[destinationChainId];
    // getFee checks if the onRamp is valid
    uint256 feeTokenAmount = getFee(destinationChainId, message);
    IERC20(message.feeToken).safeTransferFrom(msg.sender, address(onRamp), feeTokenAmount);

    // Transfer the tokens to the token pools.
    // TODO: Check the pool for how to take action
    for (uint256 i = 0; i < message.tokensAndAmounts.length; ++i) {
      IERC20 token = IERC20(message.tokensAndAmounts[i].token);
      IPool pool = onRamp.getPoolBySourceToken(token);
      if (address(pool) == address(0)) revert IBaseOnRamp.UnsupportedToken(token);
      token.safeTransferFrom(msg.sender, address(pool), message.tokensAndAmounts[i].amount);
    }

    return onRamp.forwardFromRouter(message, feeTokenAmount, msg.sender);
  }

  /// @inheritdoc IGERouter
  /// @dev returns 0 fee on invalid message.
  function getFee(uint64 destinationChainId, GEConsumer.EVM2AnyGEMessage memory message)
    public
    view
    returns (uint256 fee)
  {
    // Find and put the correct onRamp on the stack.
    IEVM2AnyGEOnRamp onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    return onRamp.getFee(message);
  }

  /// @inheritdoc IAny2EVMOffRampRouter
  function routeMessage(
    Common.Any2EVMMessage calldata message,
    bool manualExecution,
    uint256 gasLimit,
    address receiver
  ) external override onlyOffRamp returns (bool success) {
    bytes memory callData = abi.encodeWithSelector(IAny2EVMMessageReceiver.ccipReceive.selector, message);
    if (manualExecution) {
      // solhint-disable-next-line avoid-low-level-calls
      (success, ) = receiver.call(callData);
    } else {
      success = _callWithExactGas(gasLimit, receiver, 0, callData);
    }
  }

  /**
   * @dev calls target address with exactly gasAmount gas and data as calldata
   * @param gasAmount gas limit for this call
   * @param target target address
   * @param value call ether value
   * @param data calldata
   */
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

  /// @inheritdoc IGERouter
  function setOnRamp(uint64 chainId, IEVM2AnyGEOnRamp onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc IGERouter
  function getOnRamp(uint64 chainId) external view returns (IEVM2AnyGEOnRamp) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc IAny2EVMOffRampRouter
  function addOffRamp(address offRamp) external onlyOwner {
    if (address(offRamp) == address(0)) revert InvalidAddress();
    OffRampDetails memory details = s_offRamps[offRamp];
    // Check if the offramp is already allowed
    if (details.allowed) revert AlreadyConfigured(offRamp);

    // Set the s_offRamps with the new offRamp
    details.allowed = true;
    details.listIndex = uint96(s_offRampsList.length);
    s_offRamps[offRamp] = details;

    // Add to the s_offRampsList
    s_offRampsList.push(offRamp);

    emit OffRampAdded(offRamp);
  }

  function removeOffRamp(address offRamp) external onlyOwner {
    // Check that there are any feeds to remove
    uint256 listLength = s_offRampsList.length;
    if (listLength == 0) revert NoOffRampsConfigured();

    OffRampDetails memory oldDetails = s_offRamps[offRamp];
    // Check if it exists
    if (!oldDetails.allowed) revert OffRampNotAllowed(offRamp);

    // Swap the last item in the s_offRampsList with the item being removed,
    // update the index of the item moved from the end of the list to its new place,
    // then pop from the end of the list to remove.
    address lastItem = s_offRampsList[listLength - 1];
    // Perform swap
    s_offRampsList[listLength - 1] = s_offRampsList[oldDetails.listIndex];
    s_offRampsList[oldDetails.listIndex] = lastItem;
    // Update listIndex on moved item
    s_offRamps[lastItem].listIndex = oldDetails.listIndex;
    // Pop from list and delete from mapping
    s_offRampsList.pop();
    delete s_offRamps[offRamp];

    emit OffRampRemoved(offRamp);
  }

  /// @inheritdoc IAny2EVMOffRampRouter
  function getOffRamps() external view returns (address[] memory offRamps) {
    offRamps = s_offRampsList;
  }

  /// @inheritdoc IAny2EVMOffRampRouter
  function isOffRamp(address offRamp) external view returns (bool allowed) {
    return s_offRamps[offRamp].allowed;
  }

  /// @inheritdoc IBaseOnRampRouter
  function isChainSupported(uint64 chainId) public view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  /// @inheritdoc IGERouter
  function getSupportedTokens(uint64 destChainId) external view returns (address[] memory) {
    if (!isChainSupported(destChainId)) {
      return new address[](0);
    }
    return s_onRamps[uint256(destChainId)].getSupportedTokens();
  }

  // @notice only lets allowed offRamps execute
  modifier onlyOffRamp() {
    if (!s_offRamps[msg.sender].allowed) revert MustCallFromOffRamp(msg.sender);
    _;
  }
}
