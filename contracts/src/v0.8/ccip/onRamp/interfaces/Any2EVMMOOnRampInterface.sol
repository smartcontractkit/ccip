// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../utils/CCIP.sol";
import "./BaseOnRampInterface.sol";

interface Any2EVMMOOnRampInterface is BaseOnRampInterface {
  event CCIPSendRequested(CCIP.EVM2EVMMOEvent message);

  function forwardFromRouter(CCIP.EVM2AnyMOMessage memory message, address originalSender) external returns (uint64);
}
