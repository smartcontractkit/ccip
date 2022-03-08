// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../../ramps/MessageExecutor.sol";

contract MessageExecutorHelper is MessageExecutor {
  constructor(OffRampInterface offRamp) MessageExecutor(offRamp) {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
