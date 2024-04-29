// SPDX-License-Identifier: MIT
// Copied from https://github.com/ethereum-optimism/optimism/blob/f707883038d527cbf1e9f8ea513fe33255deadbc/packages/contracts-bedrock/src/universal/StandardBridge.sol#L88
pragma solidity ^0.8.0;

interface IOptimismStandardBridge {
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
