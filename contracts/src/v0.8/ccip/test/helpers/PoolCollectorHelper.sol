// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/PoolCollector.sol";

contract PoolCollectorHelper is PoolCollector {
  function collectTokens(
    EVM2EVMTollOnRampInterface onRamp,
    address[] calldata tokens,
    uint256[] calldata amounts
  ) external {
    _collectTokens(onRamp, tokens, amounts);
  }

  function chargeFee(
    EVM2EVMTollOnRampInterface onRamp,
    IERC20 feeToken,
    uint256 feeTokenAmount
  ) external {
    _chargeFee(onRamp, feeToken, feeTokenAmount);
  }
}
