// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../pools/PoolCollector.sol";

contract PoolCollectorHelper is PoolCollector {
  function collectTokens(IEVM2EVMTollOnRamp onRamp, Common.EVMTokenAndAmount[] memory tokensAndAmounts) external {
    _collectTokens(onRamp, tokensAndAmounts);
  }

  function chargeFee(
    IEVM2EVMTollOnRamp onRamp,
    IERC20 feeToken,
    uint256 feeTokenAmount
  ) external {
    _chargeFee(onRamp, feeToken, feeTokenAmount);
  }
}
