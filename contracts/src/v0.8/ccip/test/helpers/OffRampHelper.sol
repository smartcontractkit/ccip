// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../ramps/toll/Any2EVMTollOffRamp.sol";

contract OffRampHelper is Any2EVMTollOffRamp {
  constructor(TollOffRampInterface offRamp, bool needFee) Any2EVMTollOffRamp(offRamp, needFee) {}

  function report(bytes memory executableMessages) external {
    _report(bytes32(0), 0, executableMessages);
  }
}
