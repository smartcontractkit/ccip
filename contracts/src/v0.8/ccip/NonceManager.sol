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
  // TODO: add prevOffRamp
  struct PreviousRamps {
    address prevOnRamp;
  }

  /// @dev Struct that contains the chain selector and the previous on/off ramps, same as PreviousRamps but with the chain selector
  /// so that an array of these can be passed to the applyRampUpdates function
  struct PreviousRampsArgs {
    uint64 chainSelector;
    PreviousRamps prevRamps;
  }

  /// @dev Struct that contains a sender's outbound and inbound nonces
  struct Nonce {
    uint64 outbound;
    uint64 inbound;
  }

  /// @dev The current onRamp address
  address private s_onRamp;
  /// TODO: add s_offRamp;

  /// @dev previous ramps
  mapping(uint64 chainSelector => PreviousRamps prevRamps) private s_prevRamps;
  /// @dev The current nonces per sender used on the on/off ramps
  mapping(uint64 chainSelector => mapping(bytes sender => Nonce nonce)) private s_nonces;

  /// @notice Increments the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return The new outbound nonce
  function incrementOutboundNonce(uint64 destChainSelector, bytes calldata sender) external returns (uint64) {
    if (msg.sender != s_onRamp) revert OnlyCallableByOnRamp();

    Nonce storage nonce = s_nonces[destChainSelector][sender];

    uint64 outboundNonce = nonce.outbound + 1;

    if (outboundNonce == 0) {
      address prevOnRamp = s_prevRamps[destChainSelector].prevOnRamp;
      if (prevOnRamp != address(0)) {
        outboundNonce = IEVM2AnyOnRamp(prevOnRamp).getSenderNonce(abi.decode(sender, (address))) + 1;
      }
    }

    nonce.outbound = outboundNonce;

    return outboundNonce;
  }

  /// TODO: add incrementInboundNonce

  /// @notice Returns the outbound nonce for the given sender on the given destination chain
  /// @param destChainSelector The destination chain selector
  /// @param sender The encoded sender address
  /// @return The outbound nonce
  function getOutboundNonce(uint64 destChainSelector, bytes calldata sender) external view returns (uint64) {
    uint64 outboundNonce = s_nonces[destChainSelector][sender].outbound;

    if (outboundNonce == 0) {
      address prevOnRamp = s_prevRamps[destChainSelector].prevOnRamp;
      if (prevOnRamp != address(0)) {
        return IEVM2AnyOnRamp(prevOnRamp).getSenderNonce(abi.decode(sender, (address)));
      }
    }

    return outboundNonce;
  }

  /// TODO: add getInboundNonce

  /// @notice Updates the ramps and previous ramps addresses
  /// @param onRamp The new onRamp address
  /// @param prevRampsArgs The previous onRamps
  /// TODO: add offRamp
  function applyRampUpdates(address onRamp, PreviousRampsArgs[] calldata prevRampsArgs) external onlyOwner {
    if (onRamp != address(0)) {
      s_onRamp = onRamp;
      emit OnRampUpdated(onRamp);
    }

    for (uint256 i = 0; i < prevRampsArgs.length; i++) {
      PreviousRampsArgs calldata prevRampsArg = prevRampsArgs[i];

      PreviousRamps storage prevRamps = s_prevRamps[prevRampsArg.chainSelector];

      // If the previous onRamp is already set then it should not be updated
      if (prevRamps.prevOnRamp != address(0)) {
        revert InvalidRampUpdate();
      }

      prevRamps.prevOnRamp = prevRampsArg.prevRamps.prevOnRamp;
      emit PreviousOnRampUpdated(prevRampsArg.chainSelector, prevRamps.prevOnRamp);
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
  function getPrevRamps(uint64 chainSelector) external view returns (PreviousRamps memory) {
    return s_prevRamps[chainSelector];
  }
}
