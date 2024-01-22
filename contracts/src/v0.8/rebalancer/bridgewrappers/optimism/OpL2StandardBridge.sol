// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

import {L2StandardBridge} from "@eth-optimism/contracts/L2/messaging/L2StandardBridge.sol";

/// @dev This is a helper contract for generating gethwrappers
/// @dev of the optimism L2StandardBridge contract.
abstract contract OpL2StandardBridge is L2StandardBridge {
  /// @notice Emitted when an ERC20 bridge is finalized on this chain.
  /// @param localToken  Address of the ERC20 on this chain.
  /// @param remoteToken Address of the ERC20 on the remote chain.
  /// @param from        Address of the sender.
  /// @param to          Address of the receiver.
  /// @param amount      Amount of the ERC20 sent.
  /// @param extraData   Extra data sent with the transaction.
  event ERC20BridgeFinalized(
    address indexed localToken,
    address indexed remoteToken,
    address indexed from,
    address to,
    uint256 amount,
    bytes extraData
  );
}
