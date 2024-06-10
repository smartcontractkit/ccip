// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IEVM2AnyOnRamp} from "./interfaces/IEVM2AnyOnRamp.sol";

import {OwnerIsCreator} from "../shared/access/OwnerIsCreator.sol";

contract NonceManager is OwnerIsCreator {
  error OnlyCallableByOnRamp();
  error InvalidRampUpdate();

  event PreviousOnRampUpdated(uint64 indexed destChainSelector, address prevOnRamp);
  event OnRampUpdated(address onRamp);

  /// @dev Struct that contains a previous on/off ramp address with the associated source/dest chain selector
  struct PreviousRamp {
    uint64 chainSelector;
    address prevRamp;
  }

  /// @dev The current onRamp address
  address private s_onRamp;
  /// TODO: add s_offRamp;

  /// @dev previous onRamps
  mapping(uint64 destChainSelector => address prevOnRamp) private s_prevOnRamps;
  /// TODO: add previous offRamps
  /// @dev The current outbouncNonce per sender used on the EVM2EVMMultiOnRamp
  mapping(uint64 destChainSelector => mapping(bytes sender => uint64)) private s_outboundNonces;
  /// TODO: add inboundNonces

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

  /// TODO: add incrementInboundNonce

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

  /// TODO: add getInboundNonce

  /// @notice Updates the ramps and previous ramps addresses
  /// @param onRamp The new onRamp address
  /// @param prevOnRamps The previous onRamps
  function applyRampUpdates(address onRamp, PreviousRamp[] calldata prevOnRamps) external onlyOwner {
    if (onRamp != address(0)) {
      s_onRamp = onRamp;
      emit OnRampUpdated(onRamp);
    }

    for (uint256 i = 0; i < prevOnRamps.length; i++) {
      PreviousRamp calldata prevOnRamp = prevOnRamps[i];

      // If the previous onRamp is already set then it should not be updated
      if (s_prevOnRamps[prevOnRamp.chainSelector] != address(0)) {
        revert InvalidRampUpdate();
      }

      s_prevOnRamps[prevOnRamp.chainSelector] = prevOnRamp.prevRamp;
      emit PreviousOnRampUpdated(prevOnRamp.chainSelector, prevOnRamp.prevRamp);
    }

    // TODO: add offRamp logic
  }

  /// @notice Gets the current onRamp address
  /// @return The onRamp address
  function getOnRamp() external view returns (address) {
    return s_onRamp;
  }

  /// @notice Gets the previous onRamp address for the given chain selector
  /// @param chainSelector The chain selector
  /// @return The previous onRamp address
  function getPrevOnRamp(uint64 chainSelector) external view returns (address) {
    return s_prevOnRamps[chainSelector];
  }

  // TODO: add offRamp functions
}
