// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {EVM2EVMMultiOffRamp} from "../../offRamp/EVM2EVMMultiOffRamp.sol";
import {RateLimiter} from "../../libraries/RateLimiter.sol";
import {Internal} from "../../libraries/Internal.sol";
import {IgnoreContractSize} from "./IgnoreContractSize.sol";

contract EVM2EVMMultiOffRampHelper is EVM2EVMMultiOffRamp, IgnoreContractSize {
  constructor(
    StaticConfig memory staticConfig,
    RateLimiter.Config memory rateLimiterConfig
  ) EVM2EVMMultiOffRamp(staticConfig, rateLimiterConfig) {}

  function metadataHash(uint64 sourceChainSelector, address onRamp) external view returns (bytes32) {
    return _metadataHash(sourceChainSelector, onRamp, Internal.EVM_2_EVM_MESSAGE_HASH);
  }
}
