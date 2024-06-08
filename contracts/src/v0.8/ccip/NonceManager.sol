// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IEVM2AnyOnRamp} from "./interfaces/IEVM2AnyOnRamp.sol";

import {OwnerIsCreator} from "../shared/access/OwnerIsCreator.sol";

contract NonceManager is OwnerIsCreator {
  error OnlyCallableByOnRamp();
  error OnlyCallableByOffRamp();
  error InvalidRampUpdate();

  event PreviousOnRampUpdated(uint64 indexed destChainSelector, address prevOnRamp);
  event PreviousOffRampUpdated(uint64 indexed sourceChainSelector, address prevOffRamp);
  event OnRampUpdated(address onRamp);
  event OffRampUpdated(address offRamp);

  /// @dev Struct that contains a previous on/off ramp address with the associated source/dest chain selector
  struct PreviousRamp {
    uint64 chainSelector;
    address prevRamp;
  }

  /// @dev The current onRamp address
  address private s_onRamp;
  /// @dev The current offRamp address
  address private s_offRamp;

  /// @dev previous onRamps
  mapping(uint64 destChainSelector => address prevOnRamp) private s_prevOnRamps;
  /// @dev previous offRamps
  mapping(uint64 sourceChainSelector => address prevOffRamp) private s_prevOffRamps;
  /// @dev The current outbouncNonce per sender used on the EVM2EVMMultiOnRamp
  mapping(uint64 destChainSelector => mapping(bytes sender => uint64)) private s_outboundNonces;

  /// @notice Increments the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return The new outbound nonce
  function incrementOutboundNonce(uint64 destChainSelector, bytes calldata sender) external returns (uint64) {
    if (msg.sender != s_onRamp) revert OnlyCallableByOnRamp();

    uint64 outboundNonce = s_outboundNonces[destChainSelector][sender] + 1;

    if (outboundNonce == 0) {
      address prevOnRamp = s_prevOnRamps[destChainSelector];
      if (prevOnRamp != address(0)) {
        outboundNonce = IEVM2AnyOnRamp(prevOnRamp).getSenderNonce(abi.decode(sender, (address))) + 1;
      }
    }

    s_outboundNonces[destChainSelector][sender] = outboundNonce;

    return outboundNonce;
  }

  /// @notice Returns the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return The outbound nonce
  function getOutboundNonce(uint64 destChainSelector, bytes calldata sender) external view returns (uint64) {
    uint64 outboundNonce = s_outboundNonces[destChainSelector][sender];

    if (outboundNonce == 0) {
      address prevOnRamp = s_prevOnRamps[destChainSelector];
      if (prevOnRamp != address(0)) {
        return IEVM2AnyOnRamp(prevOnRamp).getSenderNonce(abi.decode(sender, (address)));
      }
    }

    return outboundNonce;
  }

  /// @notice Updates the ramps and previous ramps addresses
  /// @dev Only the owner can call this function
  /// @param onRamp The new onRamp address
  /// @param offRamp The new offRamp address
  /// @param prevOnRamps The previous onRamps
  /// @param prevOffRamps The previous offRamps
  function applyRampUpdates(
    address onRamp,
    address offRamp,
    PreviousRamp[] calldata prevOnRamps,
    PreviousRamp[] calldata prevOffRamps
  ) external onlyOwner {
    for (uint256 i = 0; i < prevOnRamps.length; i++) {
      PreviousRamp calldata prevOnRamp = prevOnRamps[i];

      // If the previous onRamp is address zero then it should not be included in the prevOnRamps array
      // If the previous onRamp is already set then it should not be updated
      if (prevOnRamp.prevRamp == address(0) || s_prevOnRamps[prevOnRamp.chainSelector] != address(0)) {
        revert InvalidRampUpdate();
      }

      s_prevOnRamps[prevOnRamp.chainSelector] = prevOnRamp.prevRamp;
      emit PreviousOnRampUpdated(prevOnRamp.chainSelector, prevOnRamp.prevRamp);
    }

    for (uint256 i = 0; i < prevOffRamps.length; i++) {
      PreviousRamp calldata prevOffRamp = prevOffRamps[i];

      // If the previous offRamp is address zero then it should not be included in the prevOffRamps array
      // If the previous offRamp is already set then it should not be updated
      if (prevOffRamp.prevRamp == address(0) || s_prevOffRamps[prevOffRamp.chainSelector] != address(0)) {
        revert InvalidRampUpdate();
      }

      s_prevOffRamps[prevOffRamp.chainSelector] = prevOffRamp.prevRamp;
      emit PreviousOffRampUpdated(prevOffRamp.chainSelector, prevOffRamp.prevRamp);
    }

    if (onRamp != address(0)) {
      s_onRamp = onRamp;
      emit OnRampUpdated(onRamp);
    }
    if (offRamp != address(0)) {
      s_offRamp = offRamp;
      emit OffRampUpdated(offRamp);
    }
  }

  /// @notice Gets the current onRamp address
  /// @return The onRamp address
  function getOnRamp() external view returns (address) {
    return s_onRamp;
  }

  /// @notice Gets the current offRamp address
  /// @return The offRamp address
  function getOffRamp() external view returns (address) {
    return s_offRamp;
  }

  /// @notice Gets the previous onRamp address for the given chain selector
  /// @param chainSelector The chain selector
  /// @return The previous onRamp address
  function getPrevOnRamp(uint64 chainSelector) external view returns (address) {
    return s_prevOnRamps[chainSelector];
  }

  /// @notice Gets the previous offRamp address for the given chain selector
  /// @param chainSelector The chain selector
  /// @return The previous offRamp address
  function getPrevOffRamp(uint64 chainSelector) external view returns (address) {
    return s_prevOffRamps[chainSelector];
  }
}
