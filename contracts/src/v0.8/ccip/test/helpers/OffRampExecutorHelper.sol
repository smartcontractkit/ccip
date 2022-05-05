// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../ramps/OffRampExecutor.sol";

contract OffRampExecutorHelper is OffRampExecutor {
  constructor(OffRampInterface offRamp, bool needFee) OffRampExecutor(offRamp, needFee) {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
