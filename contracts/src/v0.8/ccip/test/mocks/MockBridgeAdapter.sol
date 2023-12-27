// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IBridgeAdapter, IL1BridgeAdapter} from "../../pools/liquidity/interfaces/IBridge.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

/// @notice Mock L1 Bridge adapter
/// @dev Sends the L1 tokens from the msg sender to the recipient on the same chain
contract MockL1BridgeAdapter is IL1BridgeAdapter {
  function sendERC20(address l1Token, address, address recipient, uint256 amount) external payable {
    IERC20(l1Token).transferFrom(msg.sender, recipient, amount);
  }

  /// @notice Mock function to finalize a withdrawal from L2
  /// @dev does nothing as the indented action cannot be inferred from the inputs
  function finalizeWithdrawERC20FromL2(
    address l2Sender,
    address l1Receiver,
    bytes calldata bridgeSpecificPayload
  ) external {}
}

/// @notice Mock L2 Bridge adapter
/// @dev Sends the L2 tokens from the msg sender to the recipient on the same chain
contract MockL2BridgeAdapter is IBridgeAdapter {
  function sendERC20(address, address l2token, address recipient, uint256 amount) external payable {
    IERC20(l2token).transferFrom(msg.sender, recipient, amount);
  }
}
