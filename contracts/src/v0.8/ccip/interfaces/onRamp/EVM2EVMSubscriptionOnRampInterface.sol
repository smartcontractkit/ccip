// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {CCIP} from "../../models/Models.sol";
import {BaseOnRampInterface} from "./BaseOnRampInterface.sol";

interface EVM2EVMSubscriptionOnRampInterface is BaseOnRampInterface {
  event CCIPSendRequested(CCIP.EVM2EVMSubscriptionMessage message);

  function forwardFromRouter(CCIP.EVM2AnySubscriptionMessage calldata message, address originalSender)
    external
    returns (uint64);
}
