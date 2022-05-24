// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../pools/PoolCollector.sol";

contract PoolCollectorHelper is PoolCollector {
  function collectTokens(OnRampInterface onRamp, CCIP.MessagePayload memory payload) external {
    _collectTokens(onRamp, payload);
  }
}
