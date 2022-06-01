// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../pools/PoolCollector.sol";

contract PoolCollectorHelper is PoolCollector {
  function collectTokens(
    OnRampInterface onRamp,
    IERC20[] calldata tokens,
    uint256[] calldata amounts,
    IERC20 feeToken,
    uint256 feeTokenAmount
  ) external {
    _collectTokens(onRamp, tokens, amounts, feeToken, feeTokenAmount);
  }
}
