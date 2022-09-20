// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../models/Models.sol";
import "./BaseOnRampInterface.sol";

interface EVM2EVMSubscriptionOnRampInterface is BaseOnRampInterface {
  event CCIPSendRequested(CCIP.EVM2EVMSubscriptionMessage message);

  function forwardFromRouter(CCIP.EVM2AnySubscriptionMessage memory message, address originalSender)
    external
    returns (uint64);
}
