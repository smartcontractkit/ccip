// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {MultiAggregateRateLimiter} from "../../MultiAggregateRateLimiter.sol";
import {IPriceRegistry} from "../../interfaces/IPriceRegistry.sol";
import {Client} from "../../libraries/Client.sol";

contract MultiAggregateRateLimiterHelper is MultiAggregateRateLimiter {
  constructor(
    MultiAggregateRateLimiter.RateLimiterConfigUpdates memory rateLimiterConfigs,
    address admin
  ) MultiAggregateRateLimiter(rateLimiterConfigs, admin) {}

  function rateLimitValue(uint64 chainSelector, uint256 value) public {
    _rateLimitValue(chainSelector, value);
  }

  function getTokenValue(
    Client.EVMTokenAmount memory tokenAmount,
    IPriceRegistry priceRegistry
  ) public view returns (uint256) {
    return _getTokenValue(tokenAmount, priceRegistry);
  }
}
