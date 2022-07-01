// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/PoolCollector.sol";

contract PoolCollectorHelper is PoolCollector {
  function collectTokens(
    Any2EVMTollOnRampInterface onRamp,
    IERC20[] calldata tokens,
    uint256[] calldata amounts
  ) external {
    _collectTokens(onRamp, tokens, amounts);
  }

  function chargeFee(
    Any2EVMTollOnRampInterface onRamp,
    IERC20 feeToken,
    uint256 feeTokenAmount
  ) external {
    _chargeFee(onRamp, feeToken, feeTokenAmount);
  }
}
