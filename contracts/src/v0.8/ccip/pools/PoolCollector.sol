// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../vendor/SafeERC20.sol";
import "../access/OwnerIsCreator.sol";
import "../onRamp/interfaces/Any2EVMTollOnRampInterface.sol";

contract PoolCollector is OwnerIsCreator {
  using SafeERC20 for IERC20;

  event FeeCharged(address from, address to, uint256 fee);
  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  error FeeTokenAmountTooLow();

  /**
   * @notice Collect the fee
   * @param onRamp OnRamp to get the fee and pools from
   * @param feeToken the feeToken to be collected
   * @param feeTokenAmount the amount of feeToken that is available
   */
  function _chargeFee(
    Any2EVMTollOnRampInterface onRamp,
    IERC20 feeToken,
    uint256 feeTokenAmount
  ) internal returns (uint256 fee) {
    if (address(feeToken) != address(0)) {
      fee = onRamp.getRequiredFee(feeToken);
      address sender = msg.sender;
      if (fee > 0) {
        if (fee > feeTokenAmount) {
          revert FeeTokenAmountTooLow();
        }
        feeToken.safeTransferFrom(sender, address(this), fee);
        feeTokenAmount -= fee;
      }
      if (feeTokenAmount > 0) {
        // Send the fee token to the pool
        PoolInterface feeTokenPool = onRamp.getTokenPool(feeToken);
        if (address(feeTokenPool) == address(0)) revert BaseOnRampInterface.UnsupportedToken(feeToken);
        feeToken.safeTransferFrom(sender, address(feeTokenPool), feeTokenAmount);
      }
      emit FeeCharged(sender, address(this), fee);
    }
  }

  /**
   * @notice Collect tokens and send them to the pools
   * @param onRamp OnRamp to get the fee and pools from
   * @param tokens the tokens to be collected
   * @param amounts the amounts of the tokens to be collected

   */
  function _collectTokens(
    BaseOnRampInterface onRamp,
    IERC20[] memory tokens,
    uint256[] memory amounts
  ) internal {
    // Send the tokens to the pools
    for (uint256 i = 0; i < tokens.length; ++i) {
      IERC20 token = tokens[i];
      PoolInterface pool = onRamp.getTokenPool(token);
      if (address(pool) == address(0)) revert BaseOnRampInterface.UnsupportedToken(token);
      token.safeTransferFrom(msg.sender, address(pool), amounts[i]);
    }
  }

  /**
   * @notice Withdraw the fee tokens accumulated in this contract
   * @dev only callable by owner
   */
  function withdrawAccumulatedFees(
    IERC20 feeToken,
    address recipient,
    uint256 amount
  ) external onlyOwner {
    feeToken.safeTransfer(recipient, amount);
    emit FeesWithdrawn(feeToken, recipient, amount);
  }
}
