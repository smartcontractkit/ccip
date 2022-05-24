// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../interfaces/OnRampInterface.sol";
import "../../vendor/SafeERC20.sol";

contract PoolCollector {
  using SafeERC20 for IERC20;

  event FeeCharged(address from, address to, uint256 fee);

  /**
   * @notice Collect tokens from the payload and send them to the pools
   * @param onRamp OnRamp to get the fee and pools from
   * @param payload Message payload
   */
  function _collectTokens(OnRampInterface onRamp, CCIP.MessagePayload memory payload) internal {
    address sender = msg.sender;
    IERC20 feeToken = payload.tokens[0];
    uint256 fee = onRamp.getRequiredFee(feeToken);
    feeToken.safeTransferFrom(sender, address(this), fee);
    payload.amounts[0] -= fee;
    emit FeeCharged(sender, address(this), fee);
    for (uint256 i = 0; i < payload.tokens.length; i++) {
      IERC20 token = payload.tokens[i];
      PoolInterface pool = onRamp.getTokenPool(token);
      if (address(pool) == address(0)) revert OnRampInterface.UnsupportedToken(token);
      token.safeTransferFrom(sender, address(pool), payload.amounts[i]);
    }
  }
}
