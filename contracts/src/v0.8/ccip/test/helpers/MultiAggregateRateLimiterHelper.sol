// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IPriceRegistry} from "../../interfaces/IPriceRegistry.sol";
import {Client} from "../../libraries/Client.sol";
import {MultiAggregateRateLimiter} from "../../validators/MultiAggregateRateLimiter.sol";

contract MultiAggregateRateLimiterHelper is MultiAggregateRateLimiter {
  constructor(
    MultiAggregateRateLimiter.RateLimiterConfigUpdates memory rateLimiterConfigs,
    address admin,
    address priceRegistry,
    address[] memory authorizedCallers
  ) MultiAggregateRateLimiter(rateLimiterConfigs, admin, priceRegistry, authorizedCallers) {}

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
